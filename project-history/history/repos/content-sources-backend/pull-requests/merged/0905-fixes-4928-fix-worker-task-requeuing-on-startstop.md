---
type: pull_request
number: 905
title: "Fixes 4928: fix worker task requeuing on start/stop"
state: merged
author: dominikvagner
created: 2024-11-25T09:13:21Z
updated: 2024-12-04T11:29:55Z
closed: 2024-12-04T11:29:54Z
merged: 2024-12-04T11:29:54Z
base_branch: main
head_branch: 4928
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/905
---

# Pull Request #905: Fixes 4928: fix worker task requeuing on start/stop

**Author**: @dominikvagner
**Created**: November 25, 2024 at 09:13 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4928`

## Description

## Summary
This PR fixes the problems introduced with the changes made to how we close the pg pool, which were required to prevent leaving hanging connections. The introduced problem was when workers tried requeuing tasks on the context stop, but the pool was already closed at that point, so the requeuing was moved to when the workers start.

## Testing steps
1. Enqueue a bunch of tasks.
2. CTRL+C to close the server.
3. Start the server again.
4. Verify that there are no errors:
    - error requeuing task with task_id: 959ae228-6974-499b-b049-8600832355a7
    - error starting database transaction: closed pool

---

## Discussion

### Comment by @jlsherrill on November 25, 2024 at 09:30 AM UTC

https://issues.redhat.com/browse/HMS-4928

### Comment by @swadeley on December 02, 2024 at 12:59 AM UTC

/retest

### Comment by @swadeley on December 02, 2024 at 01:56 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on November 25, 2024 at 09:19 PM UTC

### Review by @rverdile - Commented on November 25, 2024 at 09:19 PM UTC

### Review by @dominikvagner - Commented on November 26, 2024 at 02:24 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 06:25 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 06:34 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 07:41 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 07:43 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 07:44 PM UTC

### Review by @dominikvagner - Commented on December 03, 2024 at 01:54 PM UTC

### Review by @dominikvagner - Commented on December 03, 2024 at 01:56 PM UTC

### Review by @rverdile - Approved on December 03, 2024 at 04:01 PM UTC

looks good! opened a card for the other issue

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/905*
