---
type: pull_request
number: 780
title: "Fixes 4478: support cancel pending tasks"
state: merged
author: rverdile
created: 2024-08-19T14:06:41Z
updated: 2024-08-22T08:30:22Z
closed: 2024-08-22T08:20:42Z
merged: 2024-08-22T08:20:42Z
base_branch: main
head_branch: rework-cancel
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/780
---

# Pull Request #780: Fixes 4478: support cancel pending tasks

**Author**: @rverdile
**Created**: August 19, 2024 at 02:06 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `rework-cancel`

## Description

## Summary
Adds support for cancelling pending tasks. Anywhere in our code that was cancelling a running task will now cancel a running or pending task.

## Testing steps
1. Add all the EPEL repos, so their snapshot tasks will use all 3 workers. These snapshot tasks will be running.
2. Create another repo with snapshot. This snapshot task will be pending.
3. Delete all the repos you just created.
4. Both the running and pending snapshot tasks should be canceled.



---

## Discussion

### Comment by @jlsherrill on August 19, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-4478

### Comment by @rverdile on August 19, 2024 at 02:59 PM UTC

/retest

### Comment by @xbhouse on August 20, 2024 at 04:40 PM UTC

i see some logs like this after adding 4 repos and deleting them all while pending or running:
`DBG CancelTask: Status Conflict`
`WRN error canceling task: error deleting heartbeat: heartbeat not found. was this task requeued recently?`

it does seem like all the tasks do get cancelled though 🤔 

### Comment by @xbhouse on August 20, 2024 at 04:49 PM UTC

code looks good to me! just needs a rebase

### Comment by @rverdile on August 20, 2024 at 05:29 PM UTC

those logs are okay. they're just warnings and in this case they don't indicate that something went wrong

### Comment by @rverdile on August 20, 2024 at 07:17 PM UTC

/retest

### Comment by @swadeley on August 21, 2024 at 07:48 AM UTC

Hi @rverdile 

I see some `update-latest-snapshot` tasks pending.

### Comment by @rverdile on August 21, 2024 at 01:48 PM UTC

thanks @swadeley, that should be fixed now

### Comment by @swadeley on August 21, 2024 at 02:40 PM UTC

/retest

### Comment by @swadeley on August 22, 2024 at 07:24 AM UTC

/retest

### Comment by @swadeley on August 22, 2024 at 08:20 AM UTC

Hi

all tasks cancelled as expected.

Thank you.

---

## Reviews

### Review by @xbhouse - Approved on August 20, 2024 at 07:34 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/780*
