---
type: pull_request
number: 1046
title: "HMS-5565: fix admin task fetch parsing error"
state: merged
author: rverdile
created: 2025-03-28T17:28:03Z
updated: 2025-04-01T13:57:59Z
closed: 2025-04-01T13:57:56Z
merged: 2025-04-01T13:57:55Z
base_branch: main
head_branch: dist-task-href
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1046
---

# Pull Request #1046: HMS-5565: fix admin task fetch parsing error

**Author**: @rverdile
**Created**: March 28, 2025 at 05:28 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `dist-task-href`

## Description

## Summary
Fixes bug where fetching a snapshot task resulted in a parsing error. The snapshot payload was incorrectly saving the distribution href instead of the distribution task href.

Also fixes issue where we were not polling the update distribution task from pulp

## Testing steps
1. Create a snapshot
2. Fetch that snapshot task from the admin tasks API: `/admin/tasks/<task_uuid>`
3. Previously, there would be a 500 because of a missing "logging_cid"



---

## Discussion

### Comment by @jlsherrill on March 28, 2025 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-5565

---

## Reviews

### Review by @xbhouse - Approved on March 31, 2025 at 04:52 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1046*
