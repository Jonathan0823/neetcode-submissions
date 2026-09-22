"""Small GitHub REST adapter backed by the preinstalled `gh` CLI."""

from __future__ import annotations

import json
import subprocess
from typing import Any

from .common import RecallError, repository_name


class GitHub:
    def __init__(self, repository: str | None = None) -> None:
        self.repository = repository or repository_name()

    def _api(self, endpoint: str, *, method: str = "GET", payload: dict[str, Any] | None = None) -> Any:
        command = ["gh", "api", endpoint, "--method", method]
        input_data = None
        if payload is not None:
            command += ["--input", "-"]
            input_data = json.dumps(payload)
        result = subprocess.run(
            command,
            input=input_data,
            text=True,
            capture_output=True,
            check=False,
        )
        if result.returncode:
            raise RecallError(result.stderr.strip() or f"GitHub API request failed: {endpoint}")
        try:
            return json.loads(result.stdout) if result.stdout.strip() else None
        except json.JSONDecodeError as exc:
            raise RecallError(f"GitHub returned invalid JSON for {endpoint}") from exc

    def list_recall_issues(self, label: str) -> list[dict[str, Any]]:
        # The machine marker is authoritative; do not require a pre-created label.
        endpoint = f"repos/{self.repository}/issues?state=all&per_page=100"
        command = ["gh", "api", endpoint, "--paginate", "--slurp"]
        result = subprocess.run(command, text=True, capture_output=True, check=False)
        if result.returncode:
            raise RecallError(result.stderr.strip() or "Unable to list GitHub Issues")
        try:
            pages = json.loads(result.stdout) if result.stdout.strip() else []
        except json.JSONDecodeError as exc:
            raise RecallError("GitHub returned invalid Issue JSON") from exc
        issues = [item for page in pages for item in page] if pages and isinstance(pages[0], list) else pages
        return [item for item in issues if isinstance(item, dict) and "pull_request" not in item]

    def create_issue(self, title: str, body: str, label: str) -> dict[str, Any]:
        return self._api(
            f"repos/{self.repository}/issues",
            method="POST",
            payload={"title": title, "body": body},
        )

    def close_issue(self, issue_number: int) -> None:
        self._api(
            f"repos/{self.repository}/issues/{issue_number}",
            method="PATCH",
            payload={"state": "closed"},
        )

    def last_editor(self, issue_number: int) -> str | None:
        """Best-effort attribution for recovery after an issues.edited event was missed."""
        try:
            events = self._api(
                f"repos/{self.repository}/issues/{issue_number}/timeline?per_page=100"
            )
        except RecallError:
            return None
        if not isinstance(events, list):
            return None
        for event in reversed(events):
            actor = event.get("actor") or event.get("user") or {}
            login = actor.get("login") if isinstance(actor, dict) else None
            if login:
                return login
        return None
