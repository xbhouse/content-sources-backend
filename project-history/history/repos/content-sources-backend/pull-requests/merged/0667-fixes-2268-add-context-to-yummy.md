---
type: pull_request
number: 667
title: "Fixes 2268: add context to yummy"
state: merged
author: rverdile
created: 2024-05-10T17:27:35Z
updated: 2024-05-15T19:37:49Z
closed: 2024-05-15T15:09:54Z
merged: 2024-05-15T15:09:54Z
base_branch: main
head_branch: yummy-ctx
labels: []
url: https://github.com/content-services/content-sources-backend/pull/667
---

# Pull Request #667: Fixes 2268: add context to yummy

**Author**: @rverdile
**Created**: May 10, 2024 at 05:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `yummy-ctx`

## Description

## Summary
Cannot be merged until the changes in yummy have been merged and a new release is made

Yummy is updated to include a context variable as a parameter in each method call. This updates all the yummy calls here. It also removed the mock yummy interface defined here, in favor of one exported by yummy.

## Testing steps
1. Introspect a repository
2. It should not fail

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 10, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-2286

### Comment by @rverdile on May 15, 2024 at 07:37 PM UTC

fixed by: https://issues.redhat.com/browse/HMS-2268

---

## Reviews

### Review by @xbhouse - Approved on May 15, 2024 at 02:24 PM UTC

looks good! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/667*
