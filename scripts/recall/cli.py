"""Command-line entry point used by GitHub Actions and local recovery runs."""

from __future__ import annotations

import argparse
import json
import os
from datetime import datetime, timezone
from pathlib import Path
from typing import Any

from .common import RecallError, append_reviews, load_json, load_reviews, local_date, write_json
from .github import GitHub
from .queue import ensure_daily_issue, reconcile_old_issues
from .review import parse_issue, reconcile_issues
from .scheduler import replay_state
from .sync import paths, persist_sync, sync_repository


def event_context() -> tuple[int | None, str | None, str | None]:
    event_path = os.environ.get("GITHUB_EVENT_PATH")
    fallback_actor = os.environ.get("GITHUB_ACTOR")
    if not event_path:
        return None, fallback_actor, None
    try:
        with open(event_path, encoding="utf-8") as handle:
            event = json.load(handle)
    except (OSError, json.JSONDecodeError):
        return None, fallback_actor, None
    issue = event.get("issue", {})
    issue_number = issue.get("number")
    actor = event.get("sender", {}).get("login") or fallback_actor
    return (int(issue_number) if issue_number else None, actor, issue.get("updated_at"))


def hydrate_daily_state(state: dict[str, Any], issues: list[dict[str, Any]]) -> None:
    """Recover associations after Issue creation committed later than the API call."""
    daily = state.setdefault("daily", {})
    for issue in issues:
        creator = issue.get("user", {}).get("login")
        if creator and creator != "github-actions[bot]":
            continue
        parsed = parse_issue(issue)
        if parsed is None:
            continue
        day, problem_id, _ = parsed
        daily.setdefault(
            day,
            {
                "problem_id": problem_id,
                "issue_number": int(issue["number"]),
                "created_at": issue.get("created_at"),
            },
        )


def run(mode: str, repo_root: Path) -> int:
    if mode not in {"sync", "rebuild", "daily-queue", "issue", "ack"}:
        raise RecallError(f"Unsupported mode: {mode}")
    files = paths(repo_root)
    now = datetime.now(timezone.utc)
    if os.environ.get("RECALL_NOW"):
        now = datetime.fromisoformat(os.environ["RECALL_NOW"].replace("Z", "+00:00"))

    if mode == "ack":
        state = load_json(files["state"], {})
        github = GitHub()
        for issue_number in state.get("pending_acknowledgements", []):
            github.close_issue(int(issue_number))
        print(f"Acknowledged {len(state.get('pending_acknowledgements', []))} review Issue(s).")
        return 0

    rebuild = mode == "rebuild"
    config, status, problems, state = sync_repository(repo_root, now, rebuild=rebuild)
    reviews = load_reviews(files["reviews"])
    github = GitHub()
    label = config.get("github", {}).get("issue_label", "recall")
    issues = github.list_recall_issues(label)
    hydrate_daily_state(state, issues)

    live_issue, live_actor, event_updated_at = (
        event_context() if mode == "issue" else (None, None, None)
    )
    if live_issue and live_actor and event_updated_at:
        state.setdefault("editor_provenance", {})[str(live_issue)] = {
            "actor": live_actor,
            "updated_at": event_updated_at,
        }
    new_records, _, errors = reconcile_issues(
        issues, state, reviews, github, config, now,
        live_issue=live_issue,
        live_actor=live_actor,
    )
    if new_records:
        append_reviews(files["reviews"], new_records)
        reviews.extend(new_records)
        state = replay_state(problems, reviews, state, config)
        state["pending_acknowledgements"] = sorted(
            set(state.get("pending_acknowledgements", []))
            | {int(record["issue_number"]) for record in new_records}
        )

    if mode in {"daily-queue", "rebuild"}:
        today = local_date(now, config.get("timezone", "Asia/Jakarta"))
        reconcile_old_issues(issues, today, github)
        metadata = load_json(files["metadata"], {})
        entry, problem_id, _ = ensure_daily_issue(
            state, problems, issues, github, config, now, metadata
        )
        status["last_queue_generated"] = today.isoformat()
        status["queue_problem"] = problem_id
        status["queue_issue"] = entry.get("issue_number")

    status["review_errors"] = errors
    status["pending_acknowledgements"] = len(state.get("pending_acknowledgements", []))
    persist_sync(repo_root, config, status, problems, state)
    print(
        f"mode={mode} problems={status['tracked_problems']} "
        f"new_reviews={len(new_records)} queue={status.get('queue_problem') or 'none'}"
    )
    if errors:
        print("Review warnings:")
        for error in errors:
            print(f"- {error}")
    return 0


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("mode", choices=["sync", "rebuild", "daily-queue", "issue", "ack"])
    parser.add_argument("--repo-root", type=Path, default=Path.cwd())
    args = parser.parse_args()
    try:
        return run(args.mode, args.repo_root.resolve())
    except RecallError as exc:
        print(f"recall error: {exc}")
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
