---
type: pull_request
number: 1108
title: "HMS-5947: fix delete-snapshots glitchtip errors"
state: closed
author: xbhouse
created: 2025-05-12T19:21:30Z
updated: 2025-06-18T12:42:39Z
closed: 2025-06-18T12:42:39Z
base_branch: main
head_branch: 5947_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1108
---

# Pull Request #1108: HMS-5947: fix delete-snapshots glitchtip errors

**Author**: @xbhouse
**Created**: May 12, 2025 at 07:21 PM UTC
**Status**: Closed
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5947_2`

## Description

## Summary

- Adds a step to delete the publication before deleting the repo version as part of the delete-snapshots task. Without this, the task could fail, leaving snapshots stuck in a soft-deleted state
- I was only able to reproduce this with Red Hat repos, which seems like what we're seeing in stage

## Testing steps

1. Import and snapshot a few RH repos. You'll need at least 2 snapshots for a repo, and it might take at least a day to get updates (there might be a better way to force an additional snapshot of RH repos?)
2. Once you have at least 2 snapshots for a repo, update the date of the older one to be marked for deletion: `UPDATE snapshots SET created_at = (NOW() - CAST('365 days' AS INTERVAL)) WHERE uuid = '<snapshot_uuid>';`
3. On the main branch, running `go run cmd/external-repos/main.go snapshot-cleanup` should fail with this error and snapshots are soft-deleted, but still exist in the content-sources database: `[Finished Task] task failed error="error deleting rpm repository versions: 400 Bad Request: [\"The repository version cannot be deleted because it (or its publications) are currently being used to distribute content. Please update the necessary distributions first.\"]"`
4. Checkout this PR, and rerun `go run cmd/external-repos/main.go snapshot-cleanup`. The task should complete and the outdated snapshot should be removed from the database



---

## Discussion

### Comment by @jlsherrill on May 12, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5947

### Comment by @xbhouse on June 18, 2025 at 12:42 PM UTC

closing in favor of https://github.com/content-services/content-sources-backend/pull/1124

---

## Reviews

### Review by @rverdile - Commented on May 13, 2025 at 12:55 PM UTC

would you mind adding an integration test for deleting red hat snapshots? There's an example of creating two red hat snapshots in the `TestUpdateLatestSnapshotForRedHatRepo()` test

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1108*
