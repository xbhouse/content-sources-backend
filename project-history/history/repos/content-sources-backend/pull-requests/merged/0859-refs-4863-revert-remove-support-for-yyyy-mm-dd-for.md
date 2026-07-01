---
type: pull_request
number: 859
title: "Refs 4863: Revert \"remove support for YYYY-MM-DD formats\""
state: merged
author: xbhouse
created: 2024-10-23T12:39:43Z
updated: 2024-10-23T15:33:47Z
closed: 2024-10-23T15:33:47Z
merged: 2024-10-23T15:33:47Z
base_branch: main
head_branch: revert-date-format-change
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/859
---

# Pull Request #859: Refs 4863: Revert "remove support for YYYY-MM-DD formats"

**Author**: @xbhouse
**Created**: October 23, 2024 at 12:39 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `revert-date-format-change`

## Description

This reverts commit 533684ffc8622ffb551bbde1952aa344197e72af.

## Summary

More changes are needed in IB so this doesn't break things, so reverting this until those are in place. 

## Testing steps

Requests with date formats of YYYY-MM-DD or RFC3339 should be accepted to `/snapshots/for_date/`:

```
{
  "date": "2024-10-23T00:00:00Z",
  "repository_uuids": ["94c19c10-c5e1-4a65-8909-75e90baa4080"]
}

{
  "date": "2024-10-23",
  "repository_uuids": ["94c19c10-c5e1-4a65-8909-75e90baa4080"]
}
```




---

## Discussion

### Comment by @jlsherrill on October 23, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4863

### Comment by @jlsherrill on October 23, 2024 at 01:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @xbhouse on October 23, 2024 at 03:33 PM UTC

merging to fix image builder

---

## Reviews

### Review by @Andrewgdewar - Approved on October 23, 2024 at 03:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/859*
