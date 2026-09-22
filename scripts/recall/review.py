"""Parse generated Issue tasklists into append-only conceptual reviews."""

from __future__ import annotations

import re
from datetime import datetime
from typing import Any, Iterable

from .common import RecallError, daily_problem_ids, iso_datetime
from .github import GitHub

MARKER = re.compile(
    r"<!--\s*recall:version=(?:1|2);date=(\d{4}-\d{2}-\d{2});problem=([^\s]+?)\s*-->"
)
CHECKBOX = re.compile(
    r"^-\s*\[([ xX])\]\s+.*?outcome:(remembered|hint|forgot)\b", re.MULTILINE
)
RATING_LABELS = {
    "remembered": "I remember it and can explain it without help",
    "hint": "I need to see a hint or the solution",
    "forgot": "I cannot recall the approach yet",
}


def render_daily_body(
    problems: list[dict[str, Any]], day: str, metadata: dict[str, Any]
) -> str:
    lines = [
        "# Daily conceptual recall",
        "",
        "Complete any selected section in 5–10 minutes. Select exactly one result per problem.",
        "",
    ]
    for index, problem in enumerate(problems, 1):
        problem_metadata = metadata.get(problem["id"], {})
        slug = problem_metadata.get("slug", problem.get("slug", problem["id"]))
        link = problem_metadata.get(
            "url", f"https://neetcode.io/problems/{slug}"
        )
        lines.extend(
            [
                f"<!-- recall:version=2;date={day};problem={problem['id']} -->",
                f"## Recall {index}: `{slug}`",
                "",
                "Before opening the solution:",
                "1. Explain the approach and why it works.",
                "2. State time and space complexity.",
                "3. Name important edge cases.",
                "",
                f"[Open problem on NeetCode]({link})",
                "",
            ]
        )
        for rating, label in RATING_LABELS.items():
            lines.append(f"- [ ] {label} (outcome:{rating})")
        lines.append("")
    lines.append("This recall checks approach memory; it is not a coding assessment.")
    return "\n".join(lines)


def render_body(problem: dict[str, Any], day: str, metadata: dict[str, Any]) -> str:
    """Render the legacy single-problem shape through the multi-item renderer."""
    return render_daily_body([problem], day, metadata)


def parse_issue_sections(
    issue: dict[str, Any],
) -> tuple[str, list[dict[str, Any]]] | None:
    """Return the canonical day and each problem section in an Issue."""
    body = str(issue.get("body", ""))
    matches = list(MARKER.finditer(body))
    if not matches:
        return None
    day = matches[0].group(1)
    sections: list[dict[str, Any]] = []
    for index, match in enumerate(matches):
        section_day, problem_id = match.groups()
        if section_day != day:
            raise RecallError(f"Issue #{issue.get('number')}: mixed recall dates")
        end = matches[index + 1].start() if index + 1 < len(matches) else len(body)
        section_body = body[match.end() : end]
        selected = [
            rating
            for checked, rating in CHECKBOX.findall(section_body)
            if checked.lower() == "x"
        ]
        sections.append({"problem_id": problem_id, "selected": selected})
    return day, sections


def parse_issue(issue: dict[str, Any]) -> tuple[str, str, list[str]] | None:
    """Preserve the v1 helper shape for callers that inspect one section."""
    parsed = parse_issue_sections(issue)
    if parsed is None:
        return None
    day, sections = parsed
    first = sections[0]
    return day, first["problem_id"], first["selected"]


def _issue_actor(
    issue: dict[str, Any],
    state: dict[str, Any],
    live_issue: int | None,
    live_actor: str | None,
) -> str | None:
    issue_number = int(issue["number"])
    if live_issue == issue_number and live_actor:
        return live_actor
    provenance = state.get("editor_provenance", {}).get(str(issue_number), {})
    if provenance.get("updated_at") == issue.get("updated_at"):
        return provenance.get("actor")
    return None


def _canonical_daily_issues(state: dict[str, Any]) -> dict[str, tuple[str, list[str]]]:
    canonical: dict[str, tuple[str, list[str]]] = {}
    for issue_number, entry in state.get("daily_issues", {}).items():
        ids = daily_problem_ids(entry)
        if entry.get("day") and ids:
            canonical[str(issue_number)] = (str(entry["day"]), ids)
    for day, entry in state.get("daily", {}).items():
        issue_number = entry.get("issue_number")
        ids = daily_problem_ids(entry)
        if issue_number and ids:
            canonical.setdefault(str(issue_number), (str(day), ids))
    return canonical


def reconcile_issues(
    issues: Iterable[dict[str, Any]],
    state: dict[str, Any],
    existing_reviews: list[dict[str, Any]],
    github: GitHub,
    config: dict[str, Any],
    now: datetime,
    *,
    live_issue: int | None = None,
    live_actor: str | None = None,
) -> tuple[list[dict[str, Any]], list[int], list[str]]:
    """Return new records, Issues to acknowledge, and validation messages."""
    learner = config.get("learner_login", "Jonathan0823")
    latest_by_review: dict[tuple[str, str], dict[str, Any]] = {}
    for record in existing_reviews:
        latest_by_review[(str(record["issue_number"]), record["problem_id"])] = record

    canonical = _canonical_daily_issues(state)
    new_records: list[dict[str, Any]] = []
    acknowledge: list[int] = []
    errors: list[str] = []
    for issue in issues:
        parsed = parse_issue_sections(issue)
        if parsed is None:
            continue
        day, sections = parsed
        issue_number = int(issue["number"])
        expected = canonical.get(str(issue_number))
        actual_ids = [section["problem_id"] for section in sections]
        if expected is None or expected != (day, actual_ids):
            errors.append(f"Issue #{issue_number}: canonical problem marker mismatch")
            continue

        selected_sections = [section for section in sections if section["selected"]]
        invalid = [section for section in selected_sections if len(section["selected"]) != 1]
        if invalid:
            errors.append(f"Issue #{issue_number}: select exactly one outcome per problem")
            continue
        if not selected_sections:
            continue

        actor = _issue_actor(issue, state, live_issue, live_actor)
        issue_records: list[dict[str, Any]] = []
        unauthorized_change = False
        updated_at = issue.get("updated_at") or iso_datetime(now)
        for section in selected_sections:
            problem_id = section["problem_id"]
            rating = section["selected"][0]
            previous = latest_by_review.get((str(issue_number), problem_id))
            if previous is not None and previous["rating"] == rating:
                continue
            if actor != learner:
                unauthorized_change = True
                break
            issue_records.append(
                {
                    "version": 2,
                    "event_id": f"issue:{issue_number}:{problem_id}:{rating}:{updated_at}",
                    "issue_number": issue_number,
                    "problem_id": problem_id,
                    "rating": rating,
                    "reviewed_at": updated_at,
                    "actor": actor,
                    "source_updated_at": updated_at,
                }
            )
        if unauthorized_change:
            errors.append(
                f"Issue #{issue_number}: edit attribution is unavailable or unauthorized"
            )
            continue
        for record in issue_records:
            new_records.append(record)
            latest_by_review[(str(issue_number), record["problem_id"])] = record

        # Keep partial Issues open; the next edit can complete the remaining sections.
        if len(selected_sections) == len(sections):
            acknowledge.append(issue_number)

    pending = set(int(item) for item in state.get("pending_acknowledgements", []))
    pending.update(acknowledge)
    state["pending_acknowledgements"] = sorted(pending)
    return new_records, acknowledge, errors
