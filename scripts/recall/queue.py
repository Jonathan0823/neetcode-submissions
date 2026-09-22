"""Daily Issue selection and publication."""

from __future__ import annotations

from datetime import date, datetime
from typing import Any

from .common import RecallError, daily_problem_ids, local_date
from .github import GitHub
from .review import parse_issue_sections, render_daily_body
from .scheduler import due_problem_ids


def _normalise_daily_state(state: dict[str, Any]) -> None:
    daily = state.setdefault("daily", {})
    for entry in daily.values():
        ids = daily_problem_ids(entry)
        entry["problem_ids"] = ids
        entry["problem_id"] = ids[0] if ids else None

    daily_issues = state.setdefault("daily_issues", {})
    for entry in daily_issues.values():
        ids = daily_problem_ids(entry)
        entry["problem_ids"] = ids
        entry["problem_id"] = ids[0] if ids else None
    for day, entry in daily.items():
        if entry.get("issue_number"):
            daily_issues.setdefault(
                str(entry["issue_number"]),
                {
                    "day": day,
                    "problem_ids": list(entry["problem_ids"]),
                    "problem_id": entry.get("problem_id"),
                },
            )


def _daily_limit(config: dict[str, Any]) -> int:
    value = config.get("recall", {}).get("daily_limit", 3)
    if not isinstance(value, int) or isinstance(value, bool) or value < 1:
        raise RecallError("recall.daily_limit must be a positive integer")
    return value


def _entry(day: str, problem_ids: list[str], issue_number: int | None, created_at: str) -> dict[str, Any]:
    return {
        "problem_ids": problem_ids,
        "problem_id": problem_ids[0] if problem_ids else None,
        "issue_number": issue_number,
        "created_at": created_at,
    }


def reconcile_old_issues(
    issues: list[dict[str, Any]], today: date, github: GitHub
) -> list[int]:
    """Close unanswered prior recall Issues after their current body was inspected."""
    closed: list[int] = []
    for issue in issues:
        parsed = parse_issue_sections(issue)
        if parsed is None or issue.get("state") != "open":
            continue
        day, sections = parsed
        if date.fromisoformat(day) < today and not any(
            section["selected"] for section in sections
        ):
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
) -> tuple[dict[str, Any], list[str], bool]:
    """Create or recover one Issue containing today's due recall sections."""
    timezone_name = config.get("timezone", "Asia/Jakarta")
    day = local_date(now, timezone_name).isoformat()
    known_problem_ids = {problem["id"] for problem in problems}
    _normalise_daily_state(state)
    daily = state["daily"]
    daily_issues = state["daily_issues"]

    existing = daily.get(day)
    if existing is not None:
        # Preserve the day's selection, including an empty queue, on every rerun.
        return existing, daily_problem_ids(existing), False

    # Recover Issues created before their state commits reached the repository.
    recovered: tuple[dict[str, Any], list[str], tuple[str, int]] | None = None
    for issue in issues:
        if issue.get("user", {}).get("login") != "github-actions[bot]":
            continue
        parsed = parse_issue_sections(issue)
        if parsed is None or parsed[0] != day:
            continue
        _, sections = parsed
        problem_ids = [section["problem_id"] for section in sections]
        if not problem_ids or any(problem_id not in known_problem_ids for problem_id in problem_ids):
            continue
        issue_number = int(issue["number"])
        created_at = str(issue.get("created_at", now.isoformat()))
        daily_issues[str(issue_number)] = {
            "day": day,
            "problem_ids": problem_ids,
            "problem_id": problem_ids[0],
        }
        candidate = (
            _entry(day, problem_ids, issue_number, created_at),
            problem_ids,
            (created_at, issue_number),
        )
        if recovered is None or candidate[2] >= recovered[2]:
            recovered = candidate
    if recovered is not None:
        daily[day] = recovered[0]
        return recovered[0], recovered[1], False

    excluded_today = {
        problem_id
        for entry in daily_issues.values()
        if entry.get("day") == day
        for problem_id in daily_problem_ids(entry)
    }
    due = [
        problem_id
        for problem_id in due_problem_ids(state, local_date(now, timezone_name))
        if problem_id not in excluded_today
    ][: _daily_limit(config)]
    entry = _entry(day, due, None, now.isoformat())
    if due:
        selected = [next(problem for problem in problems if problem["id"] == problem_id) for problem_id in due]
        label = config.get("github", {}).get("issue_label", "recall")
        title = f"Recall · {len(due)} problem(s) · {day}"
        created = github.create_issue(title, render_daily_body(selected, day, metadata), label)
        entry["issue_number"] = int(created["number"])
        daily_issues[str(created["number"])] = {
            "day": day,
            "problem_ids": due,
            "problem_id": due[0],
        }
    daily[day] = entry
    return entry, due, True
