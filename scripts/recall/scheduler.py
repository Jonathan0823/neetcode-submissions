"""Repository discovery and deterministic interval scheduling."""

from __future__ import annotations

from datetime import date, datetime, timedelta
from pathlib import Path
from typing import Any

from .common import RecallError, local_date, parse_datetime, run_git

RATINGS = {"remembered", "hint", "forgot"}


def _config_value(config: dict[str, Any], key: str, default: Any) -> Any:
    return config.get("scheduler", {}).get(key, default)


def discover(root: Path, config: dict[str, Any], now: datetime) -> list[dict[str, Any]]:
    """Find problem directories containing at least one supported submission file."""
    relative_root = Path(config.get("submission_root", "Data Structures & Algorithms"))
    submission_root = root / relative_root
    extensions = set(config.get("submission_extensions", [".go", ".js", ".ts", ".py"]))
    if not submission_root.is_dir():
        raise RecallError(f"Submission root does not exist: {submission_root}")

    folders: set[Path] = set()
    for file_path in submission_root.rglob("*"):
        if file_path.is_file() and file_path.suffix.lower() in extensions:
            folders.add(file_path.parent)

    discovered: list[dict[str, Any]] = []
    for folder in sorted(folders):
        problem_id = folder.relative_to(submission_root).as_posix()
        first_submission_at = first_git_timestamp(root, folder, now)
        discovered.append(
            {
                "id": problem_id,
                "path": folder.relative_to(root).as_posix(),
                "slug": folder.name,
                "first_submission_at": first_submission_at,
                "active": True,
            }
        )
    return discovered


def first_git_timestamp(root: Path, folder: Path, now: datetime) -> str:
    try:
        output = run_git(root, "log", "--format=%aI", "--", folder.relative_to(root).as_posix())
    except RecallError:
        output = ""
    values = [line for line in output.splitlines() if line]
    if values:
        # git log is newest first; the last value is the earliest commit in the checkout.
        return values[-1]
    return now.isoformat()


def merge_problems(
    existing: list[dict[str, Any]], discovered: list[dict[str, Any]]
) -> list[dict[str, Any]]:
    by_id = {item.get("id"): dict(item) for item in existing if item.get("id")}
    discovered_ids = {item["id"] for item in discovered}
    for item in discovered:
        previous = by_id.get(item["id"], {})
        by_id[item["id"]] = {**previous, **item, "active": True}
    for problem_id, item in by_id.items():
        if problem_id not in discovered_ids:
            item["active"] = False
    return [by_id[key] for key in sorted(by_id)]


def initial_due(problem: dict[str, Any], timezone_name: str) -> date:
    submitted = parse_datetime(problem["first_submission_at"])
    return (local_date(submitted, timezone_name) + timedelta(days=1)).isoformat()


def replay_state(
    problems: list[dict[str, Any]],
    reviews: list[dict[str, Any]],
    existing_state: dict[str, Any],
    config: dict[str, Any],
) -> dict[str, Any]:
    """Rebuild derived per-problem state while retaining daily Issue associations."""
    timezone_name = config.get("timezone", "Asia/Jakarta")
    states: dict[str, dict[str, Any]] = {}
    for problem in problems:
        problem_id = problem["id"]
        old = existing_state.get("problems", {}).get(problem_id, {})
        states[problem_id] = {
            "active": bool(problem.get("active", True)),
            "last_reviewed_at": None,
            "last_rating": None,
            "review_count": 0,
            "interval_days": 0,
            "next_due": initial_due(problem, timezone_name),
            "updated_at": old.get("updated_at"),
        }

    # A correction supersedes the previous record for the same issue.
    latest_by_issue: dict[str, dict[str, Any]] = {}
    for record in reviews:
        issue_key = str(record["issue_number"])
        previous = latest_by_issue.get(issue_key)
        if previous is None or record["reviewed_at"] >= previous["reviewed_at"]:
            latest_by_issue[issue_key] = record

    ordered = sorted(latest_by_issue.values(), key=lambda item: (item["reviewed_at"], item["event_id"]))
    for record in ordered:
        problem_id = record["problem_id"]
        if problem_id not in states:
            continue
        state = states[problem_id]
        rating = record["rating"]
        previous_rating = state["last_rating"]
        previous_interval = int(state["interval_days"] or 0)
        if rating == "remembered":
            interval = 3 if previous_rating != "remembered" else min(60, max(3, previous_interval * 2))
        elif rating == "hint":
            interval = 2
        else:
            interval = 1
        reviewed = parse_datetime(record["reviewed_at"])
        state.update(
            {
                "last_reviewed_at": record["reviewed_at"],
                "last_rating": rating,
                "review_count": int(state["review_count"]) + 1,
                "interval_days": interval,
                "next_due": (local_date(reviewed, timezone_name) + timedelta(days=interval)).isoformat(),
                "updated_at": record["reviewed_at"],
            }
        )

    return {
        "version": 1,
        "scheduler_version": int(_config_value(config, "version", 1)),
        "problems": states,
        "daily": dict(existing_state.get("daily", {})),
        "editor_provenance": dict(existing_state.get("editor_provenance", {})),
        "pending_acknowledgements": list(existing_state.get("pending_acknowledgements", [])),
    }


def due_problem_ids(
    state: dict[str, Any], today: date
) -> list[str]:
    candidates = []
    for problem_id, problem_state in state.get("problems", {}).items():
        if not problem_state.get("active", True):
            continue
        next_due = problem_state.get("next_due")
        if next_due and date.fromisoformat(next_due) <= today:
            candidates.append((next_due, problem_id))
    return [problem_id for _, problem_id in sorted(candidates)]
