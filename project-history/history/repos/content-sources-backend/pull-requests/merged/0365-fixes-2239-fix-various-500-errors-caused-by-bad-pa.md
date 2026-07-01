---
type: pull_request
number: 365
title: "Fixes 2239: fix various 500 errors caused by bad params"
state: merged
author: rverdile
created: 2023-08-16T20:04:08Z
updated: 2023-08-22T13:05:14Z
closed: 2023-08-21T13:04:44Z
merged: 2023-08-21T13:04:44Z
base_branch: main
head_branch: 500
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/365
---

# Pull Request #365: Fixes 2239: fix various 500 errors caused by bad params

**Author**: @rverdile
**Created**: August 16, 2023 at 08:04 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `500`

## Description

## Summary
There are a few different endpoints where bad params can cause 500 errors. This should fix most if not all of those endpoints.

1. GET `/repositories/{bad-uuid}/snapshots/`
2. GET `/tasks/baduuid`
3. POST `/rpms/names/` with body ` { uuids: ["John Doe"] }`
4. GET `tasks/?offset=10&limit=10&status=%00`

Note number 4 in particular, where the status is %00. That affects other endpoints params too like `/repositories/?status=`  or `/repositories/%00/snapshots`. 

## Testing steps
1. Test the endpoints listed with the params listed. You should get no 500 errors.
2. You can try similar inputs on other similar endpoints to see if anything was missed.

---

## Discussion

### Comment by @jlsherrill on August 16, 2023 at 08:05 PM UTC

https://issues.redhat.com/browse/HMS-2239

### Comment by @swadeley on August 17, 2023 at 10:24 AM UTC

Hi, please rebase.

---

## Reviews

### Review by @swadeley - Approved on August 21, 2023 at 01:03 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/365*
