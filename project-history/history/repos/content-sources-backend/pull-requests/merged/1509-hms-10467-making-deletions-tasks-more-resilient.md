---
type: pull_request
number: 1509
title: "HMS-10467: Making deletions tasks more resilient"
state: merged
author: Starle21
created: 2026-06-01T10:44:37Z
updated: 2026-06-12T12:50:12Z
closed: 2026-06-12T12:50:12Z
merged: 2026-06-12T12:50:12Z
base_branch: main
head_branch: resilient-deletion
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1509
---

# Pull Request #1509: HMS-10467: Making deletions tasks more resilient

**Author**: @Starle21
**Created**: June 01, 2026 at 10:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `resilient-deletion`

## Description

## Summary
Previously, if any of the 3 deletion tasks:
`delete_snapshots`
`delete_templates`
`delete_repository_snapshots`
failed to delete a resource, it was not exactly clear where and why it happened.

To make it more resilient, this was changed:
- detailed failure messages were added to identify easily which op caused a failure
- logging warnings if a resource was not found or nil
- loops now return errors at the end, so all the resources in a loop get a chance to be deleted

## Testing steps

Test manually by creating templates, repositories with snapshots and without snapshots and then deleting them.
Everything works as before, but you get much better idea what is happening if anything fails.
Automated tests pass.


---

## Discussion

### Comment by @xbhouse on June 01, 2026 at 11:00 AM UTC

https://issues.redhat.com/browse/HMS-10467

---

## Reviews

### Review by @rverdile - Commented on June 04, 2026 at 05:39 PM UTC

### Review by @rverdile - Commented on June 08, 2026 at 02:18 PM UTC

To small comments and then it's good to go!

### Review by @Starle21 - Commented on June 11, 2026 at 03:10 PM UTC

### Review by @rverdile - Approved on June 12, 2026 at 12:49 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1509*
