"""Synchronize submission folders and replay authoritative review history."""

from __future__ import annotations

from datetime import datetime
from pathlib import Path
from typing import Any

from .common import load_json, load_reviews, load_yaml, write_json
from .scheduler import discover, merge_problems, replay_state


def paths(repo_root: Path) -> dict[str, Path]:
    recall = repo_root / "recall"
    return {
        "config": recall / "config.yml",
        "problems": recall / "problems.json",
        "reviews": recall / "reviews.jsonl",
        "state": recall / "state.json",
        "metadata": recall / "metadata.json",
        "status": recall / "status.json",
    }


def sync_repository(
    repo_root: Path, now: datetime, *, rebuild: bool = False
) -> tuple[dict[str, Any], dict[str, Any], list[dict[str, Any]], dict[str, Any]]:
    files = paths(repo_root)
    config = load_yaml(files["config"])
    existing_problems = [] if rebuild else load_json(files["problems"], [])
    existing_state = {} if rebuild else load_json(files["state"], {})
    reviews = load_reviews(files["reviews"])
    discovered = discover(repo_root, config, now)
    problems = merge_problems(existing_problems, discovered)
    state = replay_state(problems, reviews, existing_state, config)
    active_count = sum(1 for problem in discovered if problem.get("active"))
    tracked_count = sum(1 for problem in problems if problem.get("active"))
    status = load_json(files["status"], {})
    status.update(
        {
            "version": 1,
            "last_sync": now.isoformat(),
            "submission_folders": active_count,
            "tracked_problems": tracked_count,
            "unsynced": max(0, active_count - tracked_count),
            "scheduler_version": state["scheduler_version"],
        }
    )
    if rebuild:
        status["rebuild_at"] = now.isoformat()
    return config, status, problems, state


def persist_sync(
    repo_root: Path,
    config: dict[str, Any],
    status: dict[str, Any],
    problems: list[dict[str, Any]],
    state: dict[str, Any],
) -> None:
    files = paths(repo_root)
    write_json(files["problems"], problems)
    write_json(files["state"], state)
    write_json(files["status"], status)
    if not files["metadata"].exists():
        write_json(files["metadata"], {})
    if not files["reviews"].exists():
        files["reviews"].touch()
