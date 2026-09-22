from __future__ import annotations

import json
import subprocess
import tempfile
import unittest
from datetime import datetime, timezone
from pathlib import Path

from scripts.recall.queue import ensure_daily_issue
from scripts.recall.review import (
    parse_issue,
    parse_issue_sections,
    reconcile_issues,
    render_body,
    render_daily_body,
)
from scripts.recall.scheduler import replay_state
from scripts.recall.sync import sync_repository


class FakeGitHub:
    def __init__(self, editor="Jonathan0823"):
        self.editor = editor
        self.created = []
        self.closed = []
        self.next_issue = 43

    def create_issue(self, title, body, label):
        issue = {"number": self.next_issue, "title": title, "body": body, "created_at": "2026-09-22T00:00:00Z"}
        self.next_issue += 1
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
        body = issue["body"].replace(
            "- [ ] I remember", "- [x] I remember"
        )
        issue["body"] = body
        self.assertEqual(("2026-09-22", "binary-search", ["remembered"]), parse_issue(issue))
        state = {
            "daily": {"2026-09-22": {"problem_id": "binary-search", "issue_number": 7}},
            "editor_provenance": {
                "7": {"actor": "Jonathan0823", "updated_at": "2026-09-22T01:00:00Z"}
            },
        }
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
        invalid["body"] = invalid["body"].replace(
            "- [ ] I need to see", "- [x] I need to see"
        )
        _, _, errors = reconcile_issues(
            [invalid], state, records, github, config, datetime(2026, 9, 22, tzinfo=timezone.utc)
        )
        self.assertIn("select exactly one outcome", errors[0])

        stale = dict(issue)
        stale["updated_at"] = "2026-09-22T02:00:00Z"
        stale_records, _, stale_errors = reconcile_issues(
            [stale], state, [], github, config, datetime(2026, 9, 22, tzinfo=timezone.utc)
        )
        self.assertEqual([], stale_records)
        self.assertIn("attribution is unavailable", stale_errors[0])

    def test_daily_recovery_requires_bot_and_known_problem(self):
        github = FakeGitHub()
        known = {"id": "binary-search", "slug": "binary-search"}
        valid = {
            "number": 42,
            "user": {"login": "github-actions[bot]"},
            "body": render_body(known, "2026-09-22", {}),
            "created_at": "2026-09-22T00:00:00Z",
        }
        unknown = {
            **valid,
            "number": 43,
            "body": render_body({"id": "not-tracked", "slug": "not-tracked"}, "2026-09-22", {}),
        }
        untrusted = {**valid, "number": 44, "user": {"login": "someone-else"}}
        state = {"daily": {}, "problems": {"binary-search": {"next_due": "2026-10-01"}}}
        entry, problem_id, _ = ensure_daily_issue(
            state,
            [known],
            [unknown, untrusted, valid],
            github,
            {"timezone": "Asia/Jakarta"},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
        )
        self.assertEqual(["binary-search"], problem_id)
        self.assertEqual(42, entry["issue_number"])
        self.assertEqual([], github.created)

    def test_closed_completed_issue_stays_the_only_issue_for_the_day(self):
        github = FakeGitHub()
        completed = {"id": "anagram-groups", "slug": "anagram-groups"}
        next_problem = {"id": "binary-search", "slug": "binary-search"}
        state = {
            "daily": {
                "2026-09-22": {
                    "problem_id": "anagram-groups",
                    "issue_number": 42,
                }
            },
            "problems": {
                "anagram-groups": {"last_reviewed_at": "2026-09-22T03:00:00Z", "next_due": "2026-09-25"},
                "binary-search": {"next_due": "2026-09-08"},
            },
        }
        closed_issue = {
            "number": 42,
            "state": "closed",
            "user": {"login": "github-actions[bot]"},
            "body": render_body(completed, "2026-09-22", {}),
        }
        entry, problem_id, created = ensure_daily_issue(
            state,
            [completed, next_problem],
            [closed_issue],
            github,
            {"timezone": "Asia/Jakarta"},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
        )
        self.assertEqual(["anagram-groups"], problem_id)
        self.assertEqual(42, entry["issue_number"])
        self.assertFalse(created)
        self.assertEqual({"42"}, set(state["daily_issues"]))

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
        self.assertEqual([], problem_id)
        self.assertTrue(created)
        self.assertEqual([], github.created)
        self.assertIsNone(entry["issue_number"])

    def test_daily_queue_caps_at_three_and_preserves_sections(self):
        github = FakeGitHub()
        problems = [
            {"id": f"problem-{index}", "slug": f"problem-{index}"}
            for index in range(4)
        ]
        state = {
            "daily": {},
            "daily_issues": {},
            "problems": {
                problem["id"]: {"active": True, "next_due": "2026-09-22"}
                for problem in problems
            },
        }
        entry, problem_ids, created = ensure_daily_issue(
            state,
            problems,
            [],
            github,
            {"timezone": "Asia/Jakarta", "recall": {"daily_limit": 3}},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
        )
        self.assertTrue(created)
        self.assertEqual(["problem-0", "problem-1", "problem-2"], problem_ids)
        self.assertEqual(43, entry["issue_number"])
        self.assertEqual(3, len(parse_issue_sections(github.created[0])[1]))

        rerun, rerun_ids, rerun_created = ensure_daily_issue(
            state,
            problems,
            github.created,
            github,
            {"timezone": "Asia/Jakarta", "recall": {"daily_limit": 3}},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
        )
        self.assertEqual(entry, rerun)
        self.assertEqual(problem_ids, rerun_ids)
        self.assertFalse(rerun_created)
        self.assertEqual(1, len(github.created))

        forced, forced_ids, forced_created = ensure_daily_issue(
            state,
            problems,
            github.created,
            github,
            {"timezone": "Asia/Jakarta", "recall": {"daily_limit": 3}},
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            {},
            force_new=True,
        )
        self.assertTrue(forced_created)
        self.assertEqual(problem_ids, forced_ids)
        self.assertEqual(44, forced["issue_number"])
        self.assertEqual(2, len(github.created))

    def test_partial_multi_problem_review_is_independent_and_stays_open(self):
        problems = [
            {"id": "binary-search", "slug": "binary-search"},
            {"id": "two-integer-sum", "slug": "two-integer-sum"},
        ]
        body = render_daily_body(problems, "2026-09-22", {})
        first_label = "I remember it and can explain it without help"
        body = body.replace(f"- [ ] {first_label}", f"- [x] {first_label}", 1)
        issue = {
            "number": 42,
            "body": body,
            "updated_at": "2026-09-22T01:00:00Z",
        }
        state = {
            "daily": {
                "2026-09-22": {
                    "problem_ids": ["binary-search", "two-integer-sum"],
                    "issue_number": 42,
                }
            },
            "daily_issues": {
                "42": {
                    "day": "2026-09-22",
                    "problem_ids": ["binary-search", "two-integer-sum"],
                }
            },
            "editor_provenance": {
                "42": {"actor": "Jonathan0823", "updated_at": issue["updated_at"]}
            },
        }
        config = {"learner_login": "Jonathan0823"}
        records, acknowledge, errors = reconcile_issues(
            [issue], state, [], FakeGitHub(), config,
            datetime(2026, 9, 22, tzinfo=timezone.utc),
        )
        self.assertEqual(["binary-search"], [record["problem_id"] for record in records])
        self.assertEqual([], acknowledge)
        self.assertFalse(errors)

        second_label = "I need to see a hint or the solution"
        old_checkbox = f"- [ ] {second_label}"
        new_checkbox = f"- [x] {second_label}"
        before, after = issue["body"].rsplit(old_checkbox, 1)
        issue["body"] = before + new_checkbox + after
        issue["updated_at"] = "2026-09-22T02:00:00Z"
        state["editor_provenance"]["42"] = {
            "actor": "Jonathan0823", "updated_at": issue["updated_at"]
        }
        records, acknowledge, errors = reconcile_issues(
            [issue], state, records, FakeGitHub(), config,
            datetime(2026, 9, 22, tzinfo=timezone.utc),
            live_issue=42,
            live_actor="Jonathan0823",
        )
        self.assertEqual(["two-integer-sum"], [record["problem_id"] for record in records])
        self.assertEqual([42], acknowledge)
        self.assertFalse(errors)

    def test_workflow_retries_by_regenerating_from_origin(self):
        workflow = Path(".github/workflows/recall.yml").read_text()
        self.assertIn("git reset --hard origin/main", workflow)
        self.assertIn("python -m scripts.recall.cli \"$RECALL_MODE\"", workflow)
        self.assertIn("RECALL_FORCE_NEW_ISSUE", workflow)
        self.assertNotIn("git rebase origin/main", workflow)
        self.assertIn("Unable to push after three regeneration attempts", workflow)

    @staticmethod
    def _git(root, message):
        subprocess.run(["git", "init", "-q"], cwd=root, check=True)
        subprocess.run(["git", "config", "user.email", "test@example.com"], cwd=root, check=True)
        subprocess.run(["git", "config", "user.name", "Test"], cwd=root, check=True)
        subprocess.run(["git", "add", "."], cwd=root, check=True)
        subprocess.run(["git", "commit", "-qm", message], cwd=root, check=True)


if __name__ == "__main__":
    unittest.main()
