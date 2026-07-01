---
type: pull_request
number: 472
title: "Fixes 2676: add maximum retries for task"
state: merged
author: rverdile
created: 2023-11-15T18:25:46Z
updated: 2023-11-28T21:37:41Z
closed: 2023-11-28T21:37:38Z
merged: 2023-11-28T21:37:38Z
base_branch: main
head_branch: requeue-limit
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/472
---

# Pull Request #472: Fixes 2676: add maximum retries for task

**Author**: @rverdile
**Created**: November 15, 2023 at 06:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `requeue-limit`

## Description

## Summary
This adds a "retries" column to the tasks table that counts the number of times a task has been requeued. If it has been requeued more than 3 times, the task fails.

## Testing steps
1. Create a repository with a large repo (like epel) that will take some time to sync instantly.
2. Exit the server with `Ctrl + C` so the sync task requeues
3. Do this 4 times. On the fourth time, there will be a warning that the task has failed.
4. Verify the results in the database.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 15, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-2676

### Comment by @jlsherrill on November 16, 2023 at 04:17 PM UTC

code looks good, just need to do some testing.

---

## Reviews

### Review by @jlsherrill - Approved on November 28, 2023 at 06:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/472*
