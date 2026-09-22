"""Parse generated Issue tasklists into append-only conceptual reviews."""

from __future__ import annotations

import re
from datetime import datetime
from typing import Any, Iterable

from .common import RecallError, iso_datetime
from .github import GitHub

MARKER = re.compile(r"<!--\s*recall:version=1;date=(\d{4}-\d{2}-\d{2});problem=([^\s]+?)\s*-->")
CHECKBOX = re.compile(r"^-\s*\[([ xX])\]\s+.*?outcome:(remembered|hint|forgot)\b", re.MULTILINE)
RATING_LABELS = {
    "remembered": "Ingat dan bisa menjelaskan tanpa bantuan",
    "hint": "Perlu melihat petunjuk/solusi",
    "forgot": "Belum ingat pendekatannya",
}


def render_body(problem: dict[str, Any], day: str, metadata: dict[str, Any]) -> str:
    slug = metadata.get(problem["id"], {}).get("slug", problem.get("slug", problem["id"]))
    link = metadata.get(problem["id"], {}).get(
        "url", f"https://neetcode.io/problems/{slug}"
    )
    lines = [
        f"<!-- recall:version=1;date={day};problem={problem['id']} -->",
        f"## Recall: `{slug}`",
        "",
        "Spend 5–10 minutes before opening the solution:",
        "1. Explain the approach and why it works.",
        "2. State time and space complexity.",
        "3. Name important edge cases.",
        "",
        f"[Open problem on NeetCode]({link})",
        "",
        "Select exactly one result:",
    ]
    for rating, label in RATING_LABELS.items():
        lines.append(f"- [ ] {label} (outcome:{rating})")
    lines.append("")
    lines.append("This recall checks approach memory; it is not a coding assessment.")
    return "\n".join(lines)


def parse_issue(issue: dict[str, Any]) -> tuple[str, str, list[str]] | None:
    match = MARKER.search(str(issue.get("body", "")))
    if not match:
        return None
    day, problem_id = match.groups()
    selected = [rating for checked, rating in CHECKBOX.findall(str(issue.get("body", ""))) if checked.lower() == "x"]
    return day, problem_id, selected


def _issue_actor(issue: dict[str, Any], github: GitHub, live_issue: int | None, live_actor: str | None) -> str | None:
    issue_number = int(issue["number"])
    if live_issue == issue_number and live_actor:
        return live_actor
    return github.last_editor(issue_number)


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
    """Return new records, issues to acknowledge, and non-fatal validation messages."""
    learner = config.get("learner_login", "Jonathan0823")
    latest_by_issue: dict[str, dict[str, Any]] = {}
    for record in existing_reviews:
        latest_by_issue[str(record["issue_number"])] = record

    daily_by_issue = {
        str(entry.get("issue_number")): (day, entry.get("problem_id"))
        for day, entry in state.get("daily", {}).items()
        if entry.get("issue_number")
    }
    new_records: list[dict[str, Any]] = []
    acknowledge: list[int] = []
    errors: list[str] = []
    for issue in issues:
        parsed = parse_issue(issue)
        if parsed is None:
            continue
        day, problem_id, selected = parsed
        issue_number = int(issue["number"])
        canonical = daily_by_issue.get(str(issue_number))
        if canonical is None or canonical != (day, problem_id):
            errors.append(f"Issue #{issue_number}: canonical problem marker mismatch")
            continue
        if len(selected) == 0:
            continue
        if len(selected) > 1:
            errors.append(f"Issue #{issue_number}: select exactly one outcome")
            continue
        actor = _issue_actor(issue, github, live_issue, live_actor)
        if actor != learner:
            errors.append(f"Issue #{issue_number}: edit attribution is unavailable or unauthorized")
            continue
        rating = selected[0]
        previous = latest_by_issue.get(str(issue_number))
        if previous is None or previous["rating"] != rating:
            updated_at = issue.get("updated_at") or iso_datetime(now)
            record = {
                "version": 1,
                "event_id": f"issue:{issue_number}:{rating}:{updated_at}",
                "issue_number": issue_number,
                "problem_id": problem_id,
                "rating": rating,
                "reviewed_at": updated_at,
                "actor": actor,
                "source_updated_at": updated_at,
            }
            new_records.append(record)
            latest_by_issue[str(issue_number)] = record
        acknowledge.append(issue_number)

    pending = set(int(item) for item in state.get("pending_acknowledgements", []))
    pending.update(acknowledge)
    state["pending_acknowledgements"] = sorted(pending)
    return new_records, acknowledge, errors
