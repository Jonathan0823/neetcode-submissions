"""Shared data loading and validation helpers."""

from __future__ import annotations

import json
import os
import subprocess
from datetime import date, datetime, timezone
from pathlib import Path
from typing import Any

import yaml


class RecallError(Exception):
    """Expected, user-actionable recall engine failure."""


def utc_now() -> datetime:
    return datetime.now(timezone.utc)


def iso_datetime(value: datetime) -> str:
    return value.astimezone(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z")


def parse_datetime(value: str) -> datetime:
    try:
        return datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError as exc:
        raise RecallError(f"Invalid timestamp: {value!r}") from exc


def load_yaml(path: Path) -> dict[str, Any]:
    try:
        with path.open(encoding="utf-8") as handle:
            value = yaml.safe_load(handle) or {}
    except yaml.YAMLError as exc:
        raise RecallError(f"Invalid YAML in {path}: {exc}") from exc
    if not isinstance(value, dict):
        raise RecallError(f"Expected a mapping in {path}")
    return value


def load_json(path: Path, default: Any) -> Any:
    if not path.exists():
        return default
    try:
        with path.open(encoding="utf-8") as handle:
            return json.load(handle)
    except (OSError, json.JSONDecodeError) as exc:
        raise RecallError(f"Invalid JSON in {path}: {exc}") from exc


def write_json(path: Path, value: Any) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    temporary = path.with_suffix(path.suffix + ".tmp")
    with temporary.open("w", encoding="utf-8") as handle:
        json.dump(value, handle, indent=2, ensure_ascii=False, sort_keys=True)
        handle.write("\n")
    temporary.replace(path)


def load_reviews(path: Path) -> list[dict[str, Any]]:
    if not path.exists():
        return []
    records: list[dict[str, Any]] = []
    try:
        with path.open(encoding="utf-8") as handle:
            for line_number, line in enumerate(handle, 1):
                if not line.strip():
                    continue
                value = json.loads(line)
                if not isinstance(value, dict):
                    raise RecallError(f"Review line {line_number} is not an object")
                required = {"event_id", "issue_number", "problem_id", "rating", "reviewed_at"}
                missing = required - value.keys()
                if missing:
                    raise RecallError(f"Review line {line_number} is missing: {sorted(missing)}")
                if value["rating"] not in {"remembered", "hint", "forgot"}:
                    raise RecallError(f"Review line {line_number} has an invalid rating")
                records.append(value)
    except (OSError, json.JSONDecodeError) as exc:
        raise RecallError(f"Invalid JSONL in {path}: {exc}") from exc
    return records


def append_reviews(path: Path, records: list[dict[str, Any]]) -> None:
    if not records:
        return
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("a", encoding="utf-8") as handle:
        for record in records:
            handle.write(json.dumps(record, ensure_ascii=False, sort_keys=True) + "\n")


def run_git(root: Path, *args: str) -> str:
    result = subprocess.run(
        ["git", *args], cwd=root, text=True, capture_output=True, check=False
    )
    if result.returncode:
        raise RecallError(result.stderr.strip() or f"git {' '.join(args)} failed")
    return result.stdout.strip()


def repository_name() -> str:
    value = os.environ.get("GITHUB_REPOSITORY")
    if not value or "/" not in value:
        raise RecallError("GITHUB_REPOSITORY is required for GitHub Issue operations")
    return value


def local_date(value: datetime, timezone_name: str) -> date:
    from zoneinfo import ZoneInfo

    return value.astimezone(ZoneInfo(timezone_name)).date()
