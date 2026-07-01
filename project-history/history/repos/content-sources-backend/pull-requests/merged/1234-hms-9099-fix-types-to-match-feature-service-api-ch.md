---
type: pull_request
number: 1234
title: "HMS-9099: fix types to match feature service api change"
state: merged
author: Starle21
created: 2025-10-14T12:17:44Z
updated: 2025-10-15T12:30:17Z
closed: 2025-10-15T12:30:11Z
merged: 2025-10-15T12:30:11Z
base_branch: main
head_branch: HMS-9099
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1234
---

# Pull Request #1234: HMS-9099: fix types to match feature service api change

**Author**: @Starle21
**Created**: October 14, 2025 at 12:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-9099`

## Description

## Summary
Outside service called feature-service changed their api. 

Previously it returned the field Rules as a single object, 
now it returns an array of objects. 

Our types were updated to match this change.

## Testing steps
Do an api request to `/admin/features/` path
It should return list of features without any errors.


---

## Discussion

### Comment by @xbhouse on October 14, 2025 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-9099

### Comment by @Starle21 on October 14, 2025 at 03:10 PM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on October 15, 2025 at 08:14 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1234*
