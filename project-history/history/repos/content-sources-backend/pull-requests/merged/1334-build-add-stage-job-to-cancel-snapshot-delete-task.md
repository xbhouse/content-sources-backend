---
type: pull_request
number: 1334
title: "Build: add stage job to cancel snapshot-delete tasks"
state: merged
author: rverdile
created: 2025-12-12T18:55:16Z
updated: 2025-12-12T19:35:15Z
closed: 2025-12-12T19:35:15Z
merged: 2025-12-12T19:35:15Z
base_branch: main
head_branch: add-job
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1334
---

# Pull Request #1334: Build: add stage job to cancel snapshot-delete tasks

**Author**: @rverdile
**Created**: December 12, 2025 at 06:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-job`

## Description

## Summary

## Testing steps
1. Use this to generate a bunch of running snapshot-delete tasks
```
INSERT INTO tasks (id, type, status, queued_at, started_at, finished_at, token, org_id, account_id, priority, retries, cancel_attempted)
  SELECT
      gen_random_uuid() AS id,
      'delete-snapshots' AS type,
      'running' AS status,
      clock_timestamp() - interval '5 minutes' AS queued_at,
      clock_timestamp() - interval '1 minute' AS started_at,
      NULL AS finished_at,
      gen_random_uuid() AS token,
      'test-org-' || i AS org_id,
      'test-account-' || i AS account_id,
      0 AS priority,
      0 AS retries,
      false AS cancel_attempted
  FROM generate_series(1, 100) AS i;
```
2. Use `go run cmd/jobs/main.go cancel-snapshot-delete-tasks` and verify tasks have been deleted



---

## Reviews

### Review by @xbhouse - Approved on December 12, 2025 at 07:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1334*
