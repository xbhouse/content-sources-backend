---
type: pull_request
number: 417
title: "Refs 1581: cancel task with cause"
state: merged
author: rverdile
created: 2023-10-05T15:41:30Z
updated: 2023-10-05T18:07:34Z
closed: 2023-10-05T18:07:27Z
merged: 2023-10-05T18:07:27Z
base_branch: main
head_branch: cancel-with-cause-fix
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/417
---

# Pull Request #417: Refs 1581: cancel task with cause

**Author**: @rverdile
**Created**: October 05, 2023 at 03:41 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `cancel-with-cause-fix`

## Description

## Summary
We're seeing an issue where ListenForCancel is erroring because context is canceled. Likely happening because the task is finishing before the listener is ready to wait for a notification.

You can see the timestamps are very close.
```
Oct 5, 2023 @ 08:01:11.935 --> time right before registering listener

Oct 5, 2023 @ 08:01:12.006 --> time failed because context cancelled
```

There are many cases where this does not happen and the listener has time to finish registering, which should rule out an issue with the register command hanging.

This adds `context.CancelWithCause` so that we can identify when the context is canceled because the task is finished. If the task is finished, we simply ignore these errors.

## Testing steps
You might try to reproduce this locally, but I was not able to.


---

## Discussion

### Comment by @jlsherrill on October 05, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1581

### Comment by @jlsherrill on October 05, 2023 at 04:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on October 05, 2023 at 05:20 PM UTC

I was able to reproduce locally by adding ```time.Sleep(100 * time.Millisecond)``` right before the listen for cancel.  With this change the issues is fixed!    just one small comment and then i think this is good

---

## Reviews

### Review by @jlsherrill - Commented on October 05, 2023 at 05:09 PM UTC

### Review by @jlsherrill - Commented on October 05, 2023 at 05:19 PM UTC

### Review by @rverdile - Commented on October 05, 2023 at 05:26 PM UTC

### Review by @jlsherrill - Approved on October 05, 2023 at 05:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/417*
