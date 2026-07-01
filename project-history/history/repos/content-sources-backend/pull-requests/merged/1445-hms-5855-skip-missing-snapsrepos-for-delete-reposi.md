---
type: pull_request
number: 1445
title: "HMS-5855: skip missing snaps/repos for delete-repository-snapshots"
state: merged
author: rverdile
created: 2026-04-03T19:37:50Z
updated: 2026-04-03T20:29:37Z
closed: 2026-04-03T20:29:34Z
merged: 2026-04-03T20:29:34Z
base_branch: main
head_branch: skip-missing-snapshots
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1445
---

# Pull Request #1445: HMS-5855: skip missing snaps/repos for delete-repository-snapshots

**Author**: @rverdile
**Created**: April 03, 2026 at 07:37 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `skip-missing-snapshots`

## Description

## Summary
During the delete-repository-snapshots task, for the edge case where a snapshot or repository no longer exists, continue with the task instead of failing

## Testing steps



---

## Discussion

### Comment by @rverdile on April 03, 2026 at 07:50 PM UTC

@xbhouse added unit tests!

### Comment by @xbhouse on April 03, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-5855

### Comment by @rverdile on April 03, 2026 at 08:12 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Dismissed on April 03, 2026 at 07:45 PM UTC

### Review by @xbhouse - Approved on April 03, 2026 at 07:56 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1445*
