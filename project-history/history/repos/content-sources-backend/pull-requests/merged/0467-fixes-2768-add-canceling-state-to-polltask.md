---
type: pull_request
number: 467
title: "Fixes 2768: add canceling state to PollTask"
state: merged
author: rverdile
created: 2023-11-13T20:08:04Z
updated: 2023-11-14T15:02:28Z
closed: 2023-11-14T15:02:24Z
merged: 2023-11-14T15:02:24Z
base_branch: main
head_branch: pulp-canceling-status
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/467
---

# Pull Request #467: Fixes 2768: add canceling state to PollTask

**Author**: @rverdile
**Created**: November 13, 2023 at 08:08 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `pulp-canceling-status`

## Description

## Summary
Adds "canceling" as one of the possible states while polling a pulp task. Previously it was jumping to the default case of "Pulp task with unexpected state" and logging an error.

If you search for "canceling" [here in the pulp docs](https://docs.pulpproject.org/pulpcore/restapi.html#tag/Tasks/operation/tasks_list), you can see it is a valid state.

## Testing steps
I can't find a reliable way to reproduce this, but it's a very small and simple change.

You can try

1. Creating a larger repo like epel 9, with snapshot enabled.
2. Canceling that task using the cancel endpoint (`/tasks/:uuid/cancel/`), and the `last_snapshot_task_uuid` from the create response.
3. You may or may not see the new state logged.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 13, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-2768

---

## Reviews

### Review by @jlsherrill - Approved on November 14, 2023 at 02:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/467*
