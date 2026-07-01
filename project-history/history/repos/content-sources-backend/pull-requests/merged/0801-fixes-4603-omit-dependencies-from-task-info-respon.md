---
type: pull_request
number: 801
title: "Fixes 4603: omit dependencies from task info response"
state: merged
author: rverdile
created: 2024-09-06T13:49:26Z
updated: 2024-09-11T19:30:30Z
closed: 2024-09-11T19:00:10Z
merged: 2024-09-11T19:00:10Z
base_branch: main
head_branch: omit-dependencies
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/801
---

# Pull Request #801: Fixes 4603: omit dependencies from task info response

**Author**: @rverdile
**Created**: September 06, 2024 at 01:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `omit-dependencies`

## Description

## Summary
if dependencies or dependents is empty, they'll now be omitted from the response.

## Testing steps
1. Create a repo
2. The snapshot task for that repo should show a dependent
3. The introspect task for that repo should have dependents and dependencies omitted
4. The last_snapshot_task field in the repository response for that repo should have dependents and dependencies omitted 

---

## Discussion

### Comment by @jlsherrill on September 06, 2024 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-4603

---

## Reviews

### Review by @xbhouse - Approved on September 10, 2024 at 06:59 PM UTC

looks good! i think a rebase will fix these test failures too

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/801*
