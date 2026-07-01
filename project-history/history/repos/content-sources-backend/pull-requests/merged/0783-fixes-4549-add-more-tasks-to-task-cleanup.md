---
type: pull_request
number: 783
title: "Fixes 4549: add more tasks to task cleanup"
state: merged
author: rverdile
created: 2024-08-22T18:54:41Z
updated: 2024-08-28T14:37:06Z
closed: 2024-08-28T14:37:03Z
merged: 2024-08-28T14:37:02Z
base_branch: main
head_branch: cleanup-tasks
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/783
---

# Pull Request #783: Fixes 4549: add more tasks to task cleanup

**Author**: @rverdile
**Created**: August 22, 2024 at 06:54 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cleanup-tasks`

## Description

## Summary
Adds more task types to the tasks that get cleaned up nightly. Now all tasks get cleaned up if they are older than 10 days. The deletion tasks only get cleaned if they completed.

## Testing steps
1. Create a repository and template.
2. Delete a repository and template.
3. This gives a few task types to work with. Update the tasks to change the finished_at times and verify they get cleaned up.

You can use this SQL query to update a task that will always get cleaned up
```
UPDATE tasks
SET queued_at = '2024-08-08 15:24:52.793340 +00:00', started_at = '2024-08-09 15:24:52.793340 +00:00', finished_at = '2024-08-10 15:24:52.793340 +00:00'
WHERE id = '<task id>';
```

You can use this SQL query to update a task that will get cleaned up only if it succeeded. The only difference here is the status is failed.
```
UPDATE tasks
SET queued_at = '2024-08-09 15:24:52.793340 +00:00', started_at = '2024-08-09 15:24:52.793340', finished_at = '2024-08-10 15:24:52.793340 +00:00', status = 'failed'
WHERE id = '<task id>';
```
4. run `go run cmd/external-repos/main.go nightly-jobs`
5. Verify the expected tasks are cleanedup


---

## Discussion

### Comment by @jlsherrill on August 22, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4549

### Comment by @jlsherrill on August 27, 2024 at 05:49 PM UTC

tested well and looks good, just one small comment.

---

## Reviews

### Review by @jlsherrill - Commented on August 27, 2024 at 05:46 PM UTC

### Review by @jlsherrill - Commented on August 27, 2024 at 05:47 PM UTC

### Review by @rverdile - Commented on August 27, 2024 at 05:55 PM UTC

### Review by @jlsherrill - Approved on August 27, 2024 at 05:57 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/783*
