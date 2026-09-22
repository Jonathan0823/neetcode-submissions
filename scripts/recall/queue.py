"""Daily Issue selection and publication."""

from __future__ import annotations

from datetime import date, datetime
from typing import Any

from .github import GitHub
from .review import parse_issue, render_body
from .scheduler import due_problem_ids


def reconcile_old_issues(
    issues: list[dict[str, Any]], today: date, github: GitHub
) -> list[int]:
    """Close unanswered prior recall Issues after their current body was inspected."""
    closed: list[int] = []
    for issue in issues:
        parsed = parse_issue(issue)
        if parsed is None or issue.get("state") != "open":
            continue
        day, _, selected = parsed
        if date.fromisoformat(day) < today and not selected:
            number = int(issue["number"])
            github.close_issue(number)
            closed.append(number)
    return closed


def ensure_daily_issue(
    state: dict[str, Any],
    problems: list[dict[str, Any]],
    issues: list[dict[str, Any]],
    github: GitHub,
    config: dict[str, Any],
    now: datetime,
    metadata: dict[str, Any],
) -> tuple[dict[str, Any], str | None, bool]:
    timezone_name = config.get("timezone", "Asia/Jakarta")
    from .common import local_date

    day = local_date(now, timezone_name).isoformat()
    existing = state.setdefault("daily", {}).get(day)
    if existing is not None:
        return existing, existing.get("problem_id"), False

    # Recover an Issue created before a commit/push was interrupted.
    known_problem_ids = {problem["id"] for problem in problems}
    for issue in issues:
        if issue.get("user", {}).get("login") != "github-actions[bot]":
            continue
        parsed = parse_issue(issue)
        if parsed and parsed[0] == day and parsed[1] in known_problem_ids:
            _, problem_id, _ = parsed
            entry = {
                "problem_id": problem_id,
                "issue_number": int(issue["number"]),
                "created_at": issue.get("created_at", now.isoformat()),
            }
            state["daily"][day] = entry
            return entry, problem_id, False

    due = due_problem_ids(state, local_date(now, timezone_name))
    problem_id = due[0] if due else None
    entry: dict[str, Any] = {
        "problem_id": problem_id,
        "issue_number": None,
        "created_at": now.isoformat(),
    }
    if problem_id is not None:
        problem = next(item for item in problems if item["id"] == problem_id)
        label = config.get("github", {}).get("issue_label", "recall")
        title = f"Recall · {problem.get('slug', problem_id)} · {day}"
        created = github.create_issue(title, render_body(problem, day, metadata), label)
        entry["issue_number"] = int(created["number"])
    state["daily"][day] = entry
    return entry, problem_id, True
