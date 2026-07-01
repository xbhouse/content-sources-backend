---
type: pull_request
number: 634
title: "Fixes 3940: handle pending snapshot task"
state: merged
author: xbhouse
created: 2024-04-11T15:47:11Z
updated: 2024-04-12T18:30:22Z
closed: 2024-04-12T18:19:44Z
merged: 2024-04-12T18:19:44Z
base_branch: main
head_branch: handle-pending-task
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/634
---

# Pull Request #634: Fixes 3940: handle pending snapshot task

**Author**: @xbhouse
**Created**: April 11, 2024 at 03:47 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `handle-pending-task`

## Description

## Summary

Fixes case where repository status is "Unknown" if a snapshot task never reaches a running state 

## Testing steps

- When a snapshot task is `pending` or `running`, the status in the repository response should show `Pending` and the UI should show it's `In progress`
- These `Pending` repositories should be displayed when filtering on status `Pending`
- I tested this by adding a repository and snapshotting, then manually changing the snapshot task status to `pending` in the db and verifying that the status shown in the repository status is `Pending`, but there is probably a better way to simulate this case
- EDIT: see [Ryan's comment](https://github.com/content-services/content-sources-backend/pull/634#issuecomment-2051854053) for a better way to test this

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 11, 2024 at 04:09 PM UTC

https://issues.redhat.com/browse/HMS-3488

### Comment by @jlsherrill on April 11, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-3940

### Comment by @rverdile on April 12, 2024 at 02:19 PM UTC

To test without modifying the database, start a bunch of snapshot tasks for different repos so that some would be left in the Pending state. Then if you fetch the repository associated to one of those Pending tasks, you'll no longer see "Unknown" as the repository status.

### Comment by @xbhouse on April 12, 2024 at 02:24 PM UTC

> To test without modifying the database, start a bunch of snapshot tasks for different repos so that some would be left in the Pending state. Then if you fetch the repository associated to one of those Pending tasks, you'll no longer see "Unknown" as the repository status.

ah yes, makes sense! thank you! 

### Comment by @swadeley on April 12, 2024 at 06:19 PM UTC

lgtm, thank you

---

## Reviews

### Review by @rverdile - Approved on April 12, 2024 at 01:59 PM UTC

tested and working!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/634*
