---
type: pull_request
number: 548
title: "Fixes 2017: add retry for failed tasks"
state: merged
author: rverdile
created: 2024-01-24T21:03:43Z
updated: 2024-02-01T16:53:28Z
closed: 2024-02-01T16:53:25Z
merged: 2024-02-01T16:53:25Z
base_branch: main
head_branch: retry-failed-tasks
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/548
---

# Pull Request #548: Fixes 2017: add retry for failed tasks

**Author**: @rverdile
**Created**: January 24, 2024 at 09:03 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `retry-failed-tasks`

## Description

## Summary
Adds a periodic requeue for delete-repository-snapshot tasks that have failed. Limited up to 3 retries.
Adds backoff timer for retries, with config option (tasking.retry_wait_upper_bound) for maximum wait time (wait time for final retry). 

## Testing steps
1. Add the `retry_wait_upper_bound` field to your config, with a low value like `1s`.
2. Create a snapshot
3. Get the delete-repository-snapshot task to fail. One way to do that is to do `podman stop cs_pulp_api_1`, so requests to the pulp server get a 502.
4. Observe that the task will get requeued and dequeued 3 times before permanently finishing in a failed state.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @rverdile on January 24, 2024 at 09:08 PM UTC

It occurs to me that this may not be useful without some kind of delay between the requeues, so I'm gonna put this in draft and keep thinking about that.

### Comment by @jlsherrill on January 24, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-2017

### Comment by @jlsherrill on January 25, 2024 at 07:03 PM UTC

I could see only requeing tasks that failed more than ~4 hours ago or something like that?   these tasks aren't really meant to be user visible, so its okay if they take a while.

### Comment by @rverdile on January 25, 2024 at 07:28 PM UTC

> I could see only requeing tasks that failed more than ~4 hours ago or something like that? these tasks aren't really meant to be user visible, so its okay if they take a while.

Yeah, maybe it's more complicated than it needs to be, but I've been trying to do something with exponential backoff.

### Comment by @jlsherrill on January 25, 2024 at 07:29 PM UTC

oh yeah, that does sound complicated!  But if you're close that sounds promising.  If not, i think simple is fine to start with.  Either way!

### Comment by @jlsherrill on January 30, 2024 at 08:28 PM UTC

It might be worth adding an explicit Error Log message when a task failed and is out of retries.  I could see this be worth investigating.  Thoughts?

---

## Reviews

### Review by @jlsherrill - Commented on January 30, 2024 at 08:26 PM UTC

### Review by @rverdile - Commented on January 30, 2024 at 09:03 PM UTC

### Review by @jlsherrill - Commented on January 31, 2024 at 10:02 PM UTC

### Review by @rverdile - Commented on January 31, 2024 at 10:04 PM UTC

### Review by @jlsherrill - Commented on January 31, 2024 at 10:06 PM UTC

### Review by @rverdile - Commented on January 31, 2024 at 10:09 PM UTC

### Review by @jlsherrill - Approved on February 01, 2024 at 03:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/548*
