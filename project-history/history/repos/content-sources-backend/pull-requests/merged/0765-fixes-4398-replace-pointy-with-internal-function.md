---
type: pull_request
number: 765
title: "Fixes 4398: replace pointy with internal function"
state: merged
author: jlsherrill
created: 2024-08-06T17:20:35Z
updated: 2024-08-07T21:36:07Z
closed: 2024-08-07T21:36:06Z
merged: 2024-08-07T21:36:06Z
base_branch: main
head_branch: pointy
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/765
---

# Pull Request #765: Fixes 4398: replace pointy with internal function

**Author**: @jlsherrill
**Created**: August 06, 2024 at 05:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `pointy`

## Description

## Summary

Replace pointy.pointer with a custom function (written from the ground up!) 

## Testing steps

Tests pass

---

## Discussion

### Comment by @jlsherrill on August 06, 2024 at 05:21 PM UTC

Note to completely get rid of pointy from our go.sum, we'll need to merge this too https://github.com/content-services/yummy/pull/26 and release a new yummy

### Comment by @jlsherrill on August 06, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4398

### Comment by @jlsherrill on August 06, 2024 at 07:51 PM UTC

/retest

### Comment by @jlsherrill on August 07, 2024 at 05:31 PM UTC

updated to use new yummy, no reference to pointy anywhere

### Comment by @rverdile on August 07, 2024 at 07:05 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 07, 2024 at 07:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/765*
