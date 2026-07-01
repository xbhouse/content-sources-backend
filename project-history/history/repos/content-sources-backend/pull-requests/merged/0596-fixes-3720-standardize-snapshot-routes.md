---
type: pull_request
number: 596
title: "Fixes 3720: standardize snapshot routes"
state: merged
author: jlsherrill
created: 2024-03-05T17:25:06Z
updated: 2024-03-05T22:30:27Z
closed: 2024-03-05T22:00:44Z
merged: 2024-03-05T22:00:44Z
base_branch: main
head_branch: 3720
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/596
---

# Pull Request #596: Fixes 3720: standardize snapshot routes

**Author**: @jlsherrill
**Created**: March 05, 2024 at 05:25 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3720`

## Description

## Summary

This changes 3 apis:

/repositories/snapshots/for_date/
/repositories/{uuid}/snapshots/{snapshot_uuid}/config.repo

to
/snapshots/for_date
/snapshots/{snapshot_uuid}/config.repo


## Testing steps

make requests to the new 2 apis

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 05, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-3720

---

## Reviews

### Review by @Andrewgdewar - Approved on March 05, 2024 at 09:39 PM UTC

Tested with associated front-end PR, works as intended. 
IQE PR updated, and can be found [here](https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/merge_requests/340).


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/596*
