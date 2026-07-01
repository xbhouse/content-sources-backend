---
type: pull_request
number: 469
title: "Fixes 3033: calculate timestamp to snapshot based off now()"
state: merged
author: jlsherrill
created: 2023-11-14T22:57:47Z
updated: 2023-11-24T05:00:28Z
closed: 2023-11-23T00:49:56Z
merged: 2023-11-23T00:49:56Z
base_branch: main
head_branch: 3033
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/469
---

# Pull Request #469: Fixes 3033: calculate timestamp to snapshot based off now()

**Author**: @jlsherrill
**Created**: November 14, 2023 at 10:57 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3033`

## Description

## Summary
and not current_date

## Testing steps
1.  ```make repos-import```
2.  adjust SnapshotInterval in value_constraints.go to 0
3. run  ```go run ./cmd/external-repos/main.go nightly-jobs```   Notice snapshot tasks get created
4. wait for them to finish
5. run  ```go run ./cmd/external-repos/main.go nightly-jobs```  again, Notice snapshot tasks get created. Without this pr, you wouldn't see them

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 14, 2023 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-3033

### Comment by @jlsherrill on November 15, 2023 at 02:24 AM UTC

/retest

### Comment by @swadeley on November 20, 2023 at 01:30 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on November 16, 2023 at 06:44 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/469*
