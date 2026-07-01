---
type: pull_request
number: 606
title: "Fixes 3775: handle nil response for domain updating"
state: merged
author: jlsherrill
created: 2024-03-14T12:17:13Z
updated: 2024-03-14T14:30:30Z
closed: 2024-03-14T14:26:40Z
merged: 2024-03-14T14:26:39Z
base_branch: main
head_branch: 3775
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/606
---

# Pull Request #606: Fixes 3775: handle nil response for domain updating

**Author**: @jlsherrill
**Created**: March 14, 2024 at 12:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3775`

## Description

## Summary

fixes this glitchtip, which occurs when updating a domain fails:

:runtime error: invalid memory address or nil pointer dereference
Module "github.com/content-services/content-sources-backend/pkg/tasks/worker", line 221, in (*worker).recoverOnPanic
Module "github.com/content-services/content-sources-backend/pkg/pulp_client", line 66, in (*pulpDaoImpl).UpdateDomainIfNeeded
Module "github.com/content-services/content-sources-backend/pkg/tasks", line 89, in (*SnapshotRepository).Run
Module "github.com/content-services/content-sources-backend/pkg/tasks", line 54, in SnapshotHandler
Module "github.com/content-services/content-sources-backend/pkg/tasks/worker", line 174, in (*worker).process

## Testing steps

tests pass
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 14, 2024 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-3775

### Comment by @swadeley on March 14, 2024 at 01:49 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on March 14, 2024 at 01:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/606*
