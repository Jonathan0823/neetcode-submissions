from __future__ import annotations

import json
import subprocess
import tempfile
import unittest
from datetime import datetime, timezone
from pathlib import Path

from scripts.recall.queue import ensure_daily_issue
from scripts.recall.review import parse_issue, reconcile_issues, render_body
from scripts.recall.scheduler import replay_state
from scripts.recall.sync import sync_repository


class FakeGitHub:
    def __init__(self, editor="Jonathan0823"):
        self.editor = editor
        self.created = []
        self.closed = []

    def last_editor(self, issue_number):
        return self.editor

    def create_issue(self, title, body, label):
        issue = {"number": 42, "title": title, "body": body, "created_at": "2026-09-22T00:00:00Z"}
        self.created.append(issue)
        return issue

    def close_issue(self, issue_number):
        self.closed.append(issue_number)


class RecallTests(unittest.TestCase):
    def test_sync_recovers_multiple_missed_submission_folders(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            (root / "recall").mkdir()
            (root / "Data Structures & Algorithms" / "two-integer-sum").mkdir(parents=True)
            (root / "Data Structures & Algorithms" / "two-integer-sum" / "submission-0.go").write_text("package main\n")
            self._git(root, "initial")
            now = datetime(2026, 9, 22, tzinfo=timezone.utc)
            config = {
                "timezone": "Asia/Jakarta",
                "submission_root": "Data Structures & Algorithms",
                "submission_extensions": [".go", ".js"],
                "scheduler": {"version": 1},
            }
            (root / "recall" / "config.yml").write_text(
                "timezone: Asia/Jakarta\nsubmission_root: Data Structures & Algorithms\n"
                "submission_extensions: [.go, .js]\nscheduler: {version: 1}\n"
            )
            (root / "recall" / "reviews.jsonl").write_text("")
            _, _, problems, _ = sync_repository(root, now)
            self.assertEqual(["two-integer-sum"], [p["id"] for p in problems])

            for name in ("binary-search", "anagram-groups"):
                folder = root / "Data Structures & Algorithms" / name
                folder.mkdir()
                (folder / "submission-0.js").write_text("console.log(1);\n")
            self._git(root, "missed submissions")
            _, _, problems, _ = sync_repository(root, now)
            self.assertEqual(3, len([p for p in problems if p["active"]]))
            _, _, problems, _ = sync_repository(root, now)
            self.assertEqual(3, len([p for p in problems if p["active"]]))

    def test_replay_intervals_and_correction(self):
        problems = [{"id": "binary-search", "first_submission_at": "2026-09-01T00:00:00Z", "active": True}]
        state = {"daily": {}}
        reviews = [
            {"event_id": "issue:1:remembered", "issue_number": 1, "problem_id": "binary-search", "rating": "remembered", "reviewed_at": "2026-09-02T10:00:00Z"},
            {"event_id": "issue:1:hint", "issue_number": 1, "problem_id": "binary-search", "rating": "hint", "reviewed_at": "2026-09-02T11:00:00Z"},
        ]
        result = replay_state(problems, reviews, state, {"timezone": "Asia/Jakarta", "scheduler": {"version": 1}})
        self.assertEqual("hint", result["problems"]["binary-search"]["last_rating"])
        self.assertEqual(2, result["problems"]["binary-search"]["interval_days"])
        self.assertEqual("2026-09-04", result["problems"]["binary-search"]["next_due"])

    def test_rebuild_preserves_append_only_history(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            recall = root / "recall"
            recall.mkdir()
            folder = root / "Data Structures & Algorithms" / "binary-search"
            folder.mkdir(parents=True)
            (folder / "submission-0.go").write_text("package main\n")
            (recall / "config.yml").write_text(
                "timezone: Asia/Jakarta\nsubmission_root: Data Structures & Algorithms\n"
                "submission_extensions: [.go]\nscheduler: {version: 1}\n"
            )
            review = json.dumps({"event_id": "issue:7:remembered", "issue_number": 7,
                                 "problem_id": "binary-search", "rating": "remembered",
                                 "reviewed_at": "2026-09-10T00:00:00Z"}, sort_keys=True) + "\n"
            (recall / "reviews.jsonl").write_text(review)
            self._git(root, "fixture")
            _, _, problems, state = sync_repository(
                root, datetime(2026, 9, 22, tzinfo=timezone.utc), rebuild=True
            )
            self.assertEqual("remembered", state["problems"]["binary-search"]["last_rating"])
            self.assertEqual(review, (recall / "reviews.jsonl").read_text())
            self.assertEqual(1, len(problems))

    def test_issue_checkbox_is_idempotent_and_browser_only(self):
        problem = {"id": "binary-search", "slug": "binary-search"}
        issue = {
            "number": 7,
            "body": render_body(problem, "2026-09-22", {}),
            "updated_at": "2026-09-22T01:00:00Z",
        }
        body = issue["body"].replace("- [ ] Ingat", "- [x] Ingat")
        issue["body"] = body
        self.assertEqual(("2026-09-22", "binary-search", ["remembered"]), parse_issue(issue))
        state = {"daily": {"2026-09-22": {"problem_id": "binary-search", "issue_number": 7}}}
        github = FakeGitHub()
        config = {"learner_login": "Jonathan0823"}
        records, acknowledge, errors = reconcile_issues(
            [issue], state, [], github, config, datetime(2026, 9, 22, tzinfo=timezone.utc)
        )
        self.assertEqual(1, len(records))
        self.assertEqual([7], acknowledge)
        self.assertFalse(errors)
        records2, _, _ = reconcile_issues(
            [issue], state, records, github, config, datetime(2026, 9, 22, tzinfo=timezone.utc)
        )
        self.assertEqual([], records2)

        invalid = dict(issue)
        invalid["body"] = invalid["body"].replace("- [ ] Perlu", "- [x] Perlu")
        _, _, errors = reconcile_issues(
            [invalid], state, records, github, config, datetime(2026, 9, 22, tzinfo=timezone.utc)
        )
        self.assertIn("select exactly one outcome", errors[0])

    def test_no_due_problem_creates_no_issue(self):
        github = FakeGitHub()
        state = {"daily": {}, "problems": {"old": {"active": True, "next_due": "2026-10-01"}}}
        entry, problem_id, created = ensure_daily_issue(
            state,
            [{"id": "old", "slug": "old"}],
            [],
            github,
            {"timezone": "Asia/Jakarta", "github": {"issue_label": "recall"}},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
        )
        self.assertIsNone(problem_id)
        self.assertTrue(created)
        self.assertEqual([], github.created)
        self.assertIsNone(entry["issue_number"])

    @staticmethod
    def _git(root, message):
        subprocess.run(["git", "init", "-q"], cwd=root, check=True)
        subprocess.run(["git", "config", "user.email", "test@example.com"], cwd=root, check=True)
        subprocess.run(["git", "config", "user.name", "Test"], cwd=root, check=True)
        subprocess.run(["git", "add", "."], cwd=root, check=True)
        subprocess.run(["git", "commit", "-qm", message], cwd=root, check=True)


if __name__ == "__main__":
    unittest.main()
