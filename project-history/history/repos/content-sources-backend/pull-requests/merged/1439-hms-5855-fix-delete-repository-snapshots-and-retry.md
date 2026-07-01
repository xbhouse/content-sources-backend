---
type: pull_request
number: 1439
title: "HMS-5855: fix delete-repository-snapshots and retry"
state: merged
author: rverdile
created: 2026-04-02T19:36:00Z
updated: 2026-04-03T13:34:35Z
closed: 2026-04-03T13:34:31Z
merged: 2026-04-03T13:34:31Z
base_branch: main
head_branch: fix-delete-repo-snapshots
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1439
---

# Pull Request #1439: HMS-5855: fix delete-repository-snapshots and retry

**Author**: @rverdile
**Created**: April 02, 2026 at 07:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-delete-repo-snapshots`

## Description

## Summary
The delete-repository-snapshots task needs to also delete soft deleted snapshots and the failed ones need to be retried.

## Testing steps
1. Create a repository with snapshot
2. Set the snapshot "deleted_at" field
3. Delete the repository
4. Without the fix, the deletion would fail. 
5. You can test the job by first testing without the fix.



---

## Discussion

### Comment by @xbhouse on April 02, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-5855

---

## Reviews

### Review by @xbhouse - Approved on April 03, 2026 at 01:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1439*
