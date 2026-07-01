---
type: pull_request
number: 277
title: "Refs 1471: tasking fixes"
state: merged
author: jlsherrill
created: 2023-05-19T16:50:29Z
updated: 2023-06-03T15:30:28Z
closed: 2023-05-26T16:34:43Z
merged: 2023-05-26T16:34:43Z
base_branch: main
head_branch: task_improv
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/277
---

# Pull Request #277: Refs 1471: tasking fixes

**Author**: @jlsherrill
**Created**: May 19, 2023 at 04:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `task_improv`

## Description

## Summary

* moves some config values to the config parsing
* fixes heartbeat timer
* moves to a single heartbeat value
* passes queue to tasks, for future payload updating
* Adds a queue method for payload updating


## Testing steps
Ensure introspection still works

---

## Discussion

### Comment by @jlsherrill on May 22, 2023 at 07:53 PM UTC

build passed this time, maybe some sort of ordering issue, if we see it again we should file an issue

### Comment by @jlsherrill on May 22, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1471

### Comment by @jlsherrill on May 22, 2023 at 08:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on May 22, 2023 at 08:55 PM UTC

good catch, updated!

### Comment by @jlsherrill on May 26, 2023 at 04:14 PM UTC

updated

---

## Reviews

### Review by @rverdile - Commented on May 19, 2023 at 06:57 PM UTC

Tested and looks good except for needing to update `config.yaml.example`, and the failing title/tests.

The "issues" I've seen are the ones involving the values `taskToken` and `taskId` (in worker.go) not being cleared after the task finishes. But those are more annoyances than bugs.

### Review by @jlsherrill - Commented on May 22, 2023 at 07:51 PM UTC

### Review by @rverdile - Commented on May 22, 2023 at 08:49 PM UTC

There's also a logger being propagated through in `pkg/tasks/client/client.go` and `pkg/tasks/queue/pgqueue.go`.

Other than that this looks good to go.

### Review by @rverdile - Commented on May 25, 2023 at 07:25 PM UTC

logger is being propagated through `pkg/tasks/client/client.go` too, and there's a merge conflict now.

### Review by @rverdile - Approved on May 26, 2023 at 04:29 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/277*
