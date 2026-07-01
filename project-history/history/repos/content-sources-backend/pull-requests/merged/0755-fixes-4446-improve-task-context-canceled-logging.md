---
type: pull_request
number: 755
title: "Fixes 4446: improve task context canceled logging"
state: merged
author: rverdile
created: 2024-07-24T15:49:25Z
updated: 2024-07-29T14:59:12Z
closed: 2024-07-29T14:59:06Z
merged: 2024-07-29T14:59:06Z
base_branch: main
head_branch: cleanup-cancel-logs
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/755
---

# Pull Request #755: Fixes 4446: improve task context canceled logging

**Author**: @rverdile
**Created**: July 24, 2024 at 03:49 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cleanup-cancel-logs`

## Description

## Summary
We are seeing what's likely expected error logs from tasks that are cancelled. We want to avoid logging these errors because they are expected when a task is cancelled and can be ignored.

## Testing steps
1. `make repos-import`
2. `go run cmd/external-repos/main.go nightly-jobs`. Turn off snapshotting to save time.
3. Keep exiting the server with CTRL+C and restarting it
4. You shouldn't see error logs, expect maybe `error requeuing task with task_id: 00000000-0000-0000-0000-000000000000`. Ignore that one.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 24, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4446

### Comment by @jlsherrill on July 26, 2024 at 06:11 PM UTC

/retest

### Comment by @rverdile on July 29, 2024 at 01:33 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on July 26, 2024 at 06:17 PM UTC

ack, reproduced the issue and this does seem to solve it.  

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/755*
