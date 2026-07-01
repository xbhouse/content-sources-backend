---
type: pull_request
number: 451
title: "Fixes 2918: use parent context for cancel test"
state: merged
author: jlsherrill
created: 2023-10-30T13:58:57Z
updated: 2023-10-30T14:58:47Z
closed: 2023-10-30T14:56:47Z
merged: 2023-10-30T14:56:47Z
base_branch: main
head_branch: 2918
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/451
---

# Pull Request #451: Fixes 2918: use parent context for cancel test

**Author**: @jlsherrill
**Created**: October 30, 2023 at 01:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2918`

## Description

## Summary

This uses a parent context to send the cancel notification for this test.  I suspect the ordering of things executing was causing the cancel trigger to have its context triggered prior to it finishing

## Testing steps
Tests pass

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on October 30, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-2918

### Comment by @jlsherrill on October 30, 2023 at 02:24 PM UTC

/retest

### Comment by @jlsherrill on October 30, 2023 at 02:56 PM UTC

going ahead and merging, as iqe failure is unrelated

---

## Reviews

### Review by @rverdile - Approved on October 30, 2023 at 02:07 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/451*
