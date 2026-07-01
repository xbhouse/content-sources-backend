---
type: pull_request
number: 1326
title: "HMS-9846: run snapshot cleanup in deployment"
state: merged
author: rverdile
created: 2025-12-05T15:51:20Z
updated: 2025-12-05T16:52:55Z
closed: 2025-12-05T16:52:51Z
merged: 2025-12-05T16:52:51Z
base_branch: main
head_branch: snap-cleanup
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1326
---

# Pull Request #1326: HMS-9846: run snapshot cleanup in deployment

**Author**: @rverdile
**Created**: December 05, 2025 at 03:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `snap-cleanup`

## Description

## Summary
There were missing environment variables in the cleanup deployment. FEATURES_SNAPSHOTS_ENABLED specifically was causing snapshot cleanup to be skipped

## Testing steps
1. You can recreate the behavior locally with `FEATURES_SNAPSHOTS_ENABLED=false go run cmd/external-repos/main.go cleanup --exclude pulp_orphan` There's a warning that says snapshotting is disabled.

---

## Discussion

### Comment by @xbhouse on December 05, 2025 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-9846

---

## Reviews

### Review by @xbhouse - Approved on December 05, 2025 at 04:05 PM UTC

oof, good catch! ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1326*
