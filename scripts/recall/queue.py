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
    known_problem_ids = {problem["id"] for problem in problems}
    daily = state.setdefault("daily", {})
    daily_issues = state.setdefault("daily_issues", {})
    existing = daily.get(day)
    if existing is not None:
        existing_issue = next(
            (issue for issue in issues if int(issue.get("number", -1)) == existing.get("issue_number")),
            None,
        )
        completed = (
            existing_issue is not None
            and existing_issue.get("state") == "closed"
            and state.get("problems", {}).get(existing.get("problem_id"), {}).get("last_reviewed_at")
        )
        if not completed:
            return existing, existing.get("problem_id"), False
        if existing.get("issue_number"):
            daily_issues[str(existing["issue_number"])] = {
                "day": day,
                "problem_id": existing.get("problem_id"),
            }

    # Recover an Issue created before a commit/push was interrupted.
    for issue in issues:
        if issue.get("user", {}).get("login") != "github-actions[bot]":
            continue
        parsed = parse_issue(issue)
        if parsed is None or parsed[0] != day or parsed[1] not in known_problem_ids:
            continue
        _, problem_id, _ = parsed
        issue_number = int(issue["number"])
        daily_issues[str(issue_number)] = {"day": day, "problem_id": problem_id}
        completed = (
            issue.get("state") == "closed"
            and state.get("problems", {}).get(problem_id, {}).get("last_reviewed_at")
        )
        if not completed:
            entry = {
                "problem_id": problem_id,
                "issue_number": issue_number,
                "created_at": issue.get("created_at", now.isoformat()),
            }
            state["daily"][day] = entry
            return entry, problem_id, False

    excluded_today = {
        entry.get("problem_id")
        for entry in daily_issues.values()
        if entry.get("day") == day
    }
    due = [
        problem_id
        for problem_id in due_problem_ids(state, local_date(now, timezone_name))
        if problem_id not in excluded_today
    ]
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
    if entry.get("issue_number"):
        daily_issues[str(entry["issue_number"])] = {
            "day": day,
            "problem_id": problem_id,
        }
    return entry, problem_id, True
