---
type: pull_request
number: 1200
title: "Test: Add check introspection time"
state: merged
author: swadeley
created: 2025-09-16T09:20:32Z
updated: 2025-09-24T11:41:58Z
closed: 2025-09-24T11:41:58Z
merged: 2025-09-24T11:41:58Z
base_branch: main
head_branch: swadeley/add_introspection_time_checks
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1200
---

# Pull Request #1200: Test: Add check introspection time

**Author**: @swadeley
**Created**: September 16, 2025 at 09:20 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/add_introspection_time_checks`

## Description

## Summary

Add check for API endpoints related to introspection time

## Testing steps

tests pass

---

## Discussion

### Comment by @marusak on September 17, 2025 at 09:13 AM UTC

The newly introduced test is failing:
<img width="719" height="804" alt="image" src="https://github.com/user-attachments/assets/7b0fd577-65cf-4eec-8534-c285077acb89" />


### Comment by @swadeley on September 17, 2025 at 03:49 PM UTC

Hi, I have added 1.5 sec as a buffer, or guard time, because the introspection time was 944 milliseconds earlier than the test start time.

### Comment by @xbhouse on September 23, 2025 at 09:19 PM UTC

this looks good to me, but i'm curious why the time check is needed? was this just something done in the IQE test that we missed? 

### Comment by @swadeley on September 24, 2025 at 11:27 AM UTC

> this looks good to me, but i'm curious why the time check is needed? was this just something done in the IQE test that we missed?

Yes.
We had that in IQE test to ensure the endpoints and the code paths behind them work as expected.


---

## Reviews

### Review by @xbhouse - Approved on September 24, 2025 at 11:34 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1200*
