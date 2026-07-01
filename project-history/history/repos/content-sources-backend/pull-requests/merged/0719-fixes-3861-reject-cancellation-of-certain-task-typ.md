---
type: pull_request
number: 719
title: "Fixes 3861: reject cancellation of certain task types"
state: merged
author: xbhouse
created: 2024-06-25T20:21:03Z
updated: 2024-07-01T13:42:10Z
closed: 2024-07-01T13:42:10Z
merged: 2024-07-01T13:42:10Z
base_branch: main
head_branch: 3861
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/719
---

# Pull Request #719: Fixes 3861: reject cancellation of certain task types

**Author**: @xbhouse
**Created**: June 25, 2024 at 08:21 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3861`

## Description

## Summary

- Adds a list of cancellable tasks to check against and rejects cancelling any tasks not included in the list. Currently includes only snapshot and introspect tasks
- Unrelated but attempting to cancel a RH task resulted in a panic due to dereferencing an error that doesn't exist so this includes that fix also

## Testing steps

- Start a few different tasks in addition to introspect and snapshot and take note of the task uuid. List of all task types [here](https://github.com/content-services/content-sources-backend/blob/main/pkg/config/tasks.go)
- Make a request to cancel any task that isn't snapshot or introspect: `POST /tasks/:uuid/cancel/`. You should get a 400
- Cancelling snapshot or introspect tasks works as before

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 25, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3861

### Comment by @swadeley on June 27, 2024 at 08:42 AM UTC

/retest

### Comment by @swadeley on June 27, 2024 at 07:33 PM UTC

go lint fails

---

## Reviews

### Review by @rverdile - Commented on June 26, 2024 at 01:46 PM UTC

### Review by @xbhouse - Commented on June 26, 2024 at 02:19 PM UTC

### Review by @xbhouse - Commented on June 26, 2024 at 03:21 PM UTC

### Review by @rverdile - Commented on June 26, 2024 at 03:32 PM UTC

### Review by @xbhouse - Commented on June 26, 2024 at 03:46 PM UTC

### Review by @rverdile - Commented on June 26, 2024 at 06:40 PM UTC

code looks good!! 

could you add a test to pgqueue_test.go to check that SendCancelNotification() errors for an un-cancelable type?

### Review by @rverdile - Commented on June 28, 2024 at 08:06 PM UTC

### Review by @rverdile - Approved on June 28, 2024 at 08:38 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/719*
