---
type: pull_request
number: 1105
title: "Build: separate delete-snapshots from process-repos"
state: merged
author: xbhouse
created: 2025-05-07T14:15:44Z
updated: 2025-05-22T19:39:44Z
closed: 2025-05-22T19:39:44Z
merged: 2025-05-22T19:39:44Z
base_branch: main
head_branch: 5947
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1105
---

# Pull Request #1105: Build: separate delete-snapshots from process-repos

**Author**: @xbhouse
**Created**: May 07, 2025 at 02:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `5947`

## Description

## Summary

Separates the snapshot-cleanup job from process-repos so it can be run on its own schedule (this sets it to run daily). This may also help narrow down and reduce the errors being reported by glitchtip (https://issues.redhat.com/browse/HMS-5947)

## Testing steps

1. Create a repo, let it snapshot, then edit the repo URL to get another snapshot
2. Adjust the date of one of the snapshots to be a year old: `UPDATE snapshots SET created_at = (NOW() - CAST('365 days' AS INTERVAL)) WHERE uuid = '<snapshot_uuid>';`
3. Run `go run cmd/external-repos/main.go process-repos`. You shouldn't see any delete-snapshots tasks enqueued
4. Run `go run cmd/external-repos/main.go snapshot-cleanup`. This should start the delete-snapshots task


---

## Reviews

### Review by @rverdile - Commented on May 12, 2025 at 06:35 PM UTC

### Review by @rverdile - Commented on May 12, 2025 at 06:38 PM UTC

### Review by @xbhouse - Commented on May 12, 2025 at 07:31 PM UTC

### Review by @rverdile - Approved on May 12, 2025 at 08:31 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1105*
