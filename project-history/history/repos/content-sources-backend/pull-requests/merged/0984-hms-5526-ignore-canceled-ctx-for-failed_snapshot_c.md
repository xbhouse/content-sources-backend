---
type: pull_request
number: 984
title: "HMS-5526: ignore canceled ctx for failed_snapshot_count"
state: merged
author: rverdile
created: 2025-02-18T20:10:25Z
updated: 2025-02-19T14:57:03Z
closed: 2025-02-19T14:57:00Z
merged: 2025-02-19T14:57:00Z
base_branch: main
head_branch: small-bugs
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/984
---

# Pull Request #984: HMS-5526: ignore canceled ctx for failed_snapshot_count

**Author**: @rverdile
**Created**: February 18, 2025 at 08:10 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `small-bugs`

## Description

## Summary
There are two small bugs fixed here

- For canceled snapshots, the failed_snapshot_count was not being updated because the DB query was using context that was canceled. This changes it to use context.Background().
- GetContentPath() in the pulp client was logging an error when it failed due to canceled context, but an error is expected if context is canceled
## Testing steps
1. In the UI, add and remove the popular repositories (EPEL)
2. After removing them, there will no longer be an error logged "Failed to update failed_snapshot_count: context canceled"



---

## Discussion

### Comment by @jlsherrill on February 18, 2025 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5526

---

## Reviews

### Review by @xbhouse - Approved on February 18, 2025 at 10:17 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/984*
