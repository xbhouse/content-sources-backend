---
type: pull_request
number: 527
title: "Refs 3372: Drop all snapshots in stage"
state: merged
author: jlsherrill
created: 2024-01-11T15:36:21Z
updated: 2024-01-12T19:50:56Z
closed: 2024-01-12T19:50:56Z
merged: 2024-01-12T19:50:56Z
base_branch: main
head_branch: 3372_del_snaps
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/527
---

# Pull Request #527: Refs 3372: Drop all snapshots in stage

**Author**: @jlsherrill
**Created**: January 11, 2024 at 03:36 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3372_del_snaps`

## Description

## Summary

Deletes all snapshots in stage, so we can migrate to a new pulp server.  Will be merged after snapshotting is disabled in stage.  The repos & versions will remain in the pulp server, but we will eventually just drop the pulp server.  The rpms will remain in the s3 bucket, but we will switch to a new bucket and delete the old one.  

## Testing steps

Create a repo with snapshotting, let it snapshot, see a snapshot is created. 
run: `go run cmd/cleanup_versions/main.go  --force`
snapshot should no longer exist.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 11, 2024 at 03:47 PM UTC

DO NOT MERGE until snapshotting is disabled in stage

### Comment by @jlsherrill on January 11, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3372

### Comment by @jlsherrill on January 11, 2024 at 04:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @Andrewgdewar - Approved on January 11, 2024 at 04:53 PM UTC

Indeed! Deletes things

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/527*
