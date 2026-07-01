---
repo: patchman-engine
pr_number: 1353
period: expansion
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #1353

## Problem / Motivation

The `system_package` table stored package data as JSON blobs per system — simple early on, but query performance degraded severely for large accounts (100k+ systems). RHINENG-5394 initiated migration to normalized `system_package2` tables with per-account migration to limit lock duration and timeout risk.

## Solution Approach

This PR implements per-account migration: iterate accounts, migrate package rows in batches, run VACUUM in goroutines to avoid API timeouts. Part of a multi-PR sequence (#1316 POC through #1356 drop old table) authored primarily by @psegedy and @MichaelMraka in Dec 2023.

## Key Design Decisions

- **Per-account batching**: Avoid single long transaction locking the entire table.
- **Async vacuum**: API responsiveness during migration operations.
- **POC → production path**: Several POC PRs validated admin API migration before this production implementation.

## Impact

One of the largest schema migrations in patchman-engine history. Enabled continued scale growth through 2024–2026. Cyndi filter changes (#6036) landed the same week, showing parallel inventory pipeline work.

## Review Notes

Fast merge (same day). @MichaelMraka approved. Jira RHINENG-5394 linked automatically.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
