---
type: pull_request
number: 936
title: "HMS-5259: cleanup canceled tasks"
state: merged
author: rverdile
created: 2025-01-09T21:13:13Z
updated: 2025-01-22T15:55:09Z
closed: 2025-01-13T12:37:01Z
merged: 2025-01-13T12:37:01Z
base_branch: main
head_branch: cleanup-canceled-tasks
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/936
---

# Pull Request #936: HMS-5259: cleanup canceled tasks

**Author**: @rverdile
**Created**: January 09, 2025 at 09:13 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cleanup-canceled-tasks`

## Description

## Summary
Adds canceled tasks to the query that cleans up all the old the tasks. Cleans up canceled with a queued_at time older than 20 days, not finished_at because canceled tasks do not have a finished_at time

## Testing steps
1. Create and then delete a repository so that the snapshot and/or update-latest-snapshot tasks get canceled
2. Manually edit the queued_at time so it is older than 20 days
3. Run nightly-jobs and the task should be cleaned up


---

## Discussion

### Comment by @jlsherrill on January 09, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5259

---

## Reviews

### Review by @dominikvagner - Approved on January 13, 2025 at 08:51 AM UTC

looks good and works as expected! 🎉 👏🏼 
approved ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/936*
