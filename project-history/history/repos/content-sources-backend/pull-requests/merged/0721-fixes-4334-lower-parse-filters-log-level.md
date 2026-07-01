---
type: pull_request
number: 721
title: "Fixes 4334: lower parse filters log level"
state: merged
author: rverdile
created: 2024-06-26T18:36:25Z
updated: 2024-06-27T17:42:35Z
closed: 2024-06-27T17:42:30Z
merged: 2024-06-27T17:42:30Z
base_branch: main
head_branch: 4334
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/721
---

# Pull Request #721: Fixes 4334: lower parse filters log level

**Author**: @rverdile
**Created**: June 26, 2024 at 06:36 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4334`

## Description

## Summary
Changes all the parse filter error returns to log at the same level (INFO)

## Testing steps
1. Make a request to `tasks/?exclude_red_hat_org=notabool`, where the params are the wrong type
2. In the server logs, there will be an error logged at the INFO level
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 26, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4334

---

## Reviews

### Review by @xbhouse - Approved on June 27, 2024 at 02:17 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/721*
