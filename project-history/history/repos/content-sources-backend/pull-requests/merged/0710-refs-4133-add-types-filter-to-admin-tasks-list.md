---
type: pull_request
number: 710
title: "Refs 4133: add types filter to admin tasks list"
state: merged
author: rverdile
created: 2024-06-18T15:26:38Z
updated: 2024-06-24T19:07:45Z
closed: 2024-06-19T16:08:54Z
merged: 2024-06-19T16:08:54Z
base_branch: main
head_branch: filter-task-type
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/710
---

# Pull Request #710: Refs 4133: add types filter to admin tasks list

**Author**: @rverdile
**Created**: June 18, 2024 at 03:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `filter-task-type`

## Description

## Summary
Adds ability to filter admin tasks by type

## Testing steps
1. Either test using frontend PR or make a request like `GET /admin/tasks/?type=introspect,snapshot`
2. Make sure you've created/deleted some repos and templates so you have a few task types to filter.
3. These are all the task types: https://github.com/content-services/content-sources-backend/blob/main/pkg/config/tasks.go#L4-L9
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 18, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4133

### Comment by @jlsherrill on June 18, 2024 at 03:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on June 19, 2024 at 04:08 PM UTC

Works great, thank you

https://github.com/content-services/content-sources-frontend/pull/285#issuecomment-2179052469

---

## Reviews

### Review by @xbhouse - Approved on June 18, 2024 at 07:32 PM UTC

tested with API and frontend PR, looks great!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/710*
