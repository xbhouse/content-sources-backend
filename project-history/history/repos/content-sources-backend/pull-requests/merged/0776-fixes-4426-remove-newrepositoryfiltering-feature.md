---
type: pull_request
number: 776
title: "Fixes 4426: Remove NewRepositoryFiltering feature"
state: merged
author: Andrewgdewar
created: 2024-08-14T20:38:08Z
updated: 2024-08-30T13:57:41Z
closed: 2024-08-30T13:57:41Z
merged: 2024-08-30T13:57:41Z
base_branch: main
head_branch: HMS-4426
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/776
---

# Pull Request #776: Fixes 4426: Remove NewRepositoryFiltering feature

**Author**: @Andrewgdewar
**Created**: August 14, 2024 at 08:38 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-4426`

## Description

## Summary

Associated ticket [here](https://issues.redhat.com/browse/HMS-4426). 

This removes the NewRepositoryFiltering feature option. 

## Testing steps
- Call the features API, the NewRepositoryFiltering variable should no longer appear.


---

## Discussion

### Comment by @jlsherrill on August 14, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4426

### Comment by @xbhouse on August 20, 2024 at 04:09 PM UTC

i think a few test cases are failing in dao/snapshots_test.go because all the repo types (external, redhat, and upload) are being listed now, so you may need to adjust the values being asserted 

### Comment by @swadeley on August 27, 2024 at 12:26 PM UTC

Hi @Andrewgdewar 

When you have time to fix the unit tests, please rebase to pick up change in https://github.com/content-services/content-sources-backend/pull/788

Thank you

---

## Reviews

### Review by @jlsherrill - Commented on August 16, 2024 at 12:53 PM UTC

### Review by @jlsherrill - Commented on August 16, 2024 at 12:53 PM UTC

### Review by @xbhouse - Commented on August 20, 2024 at 01:50 PM UTC

### Review by @jlsherrill - Commented on August 27, 2024 at 01:00 PM UTC

### Review by @jlsherrill - Commented on August 27, 2024 at 01:23 PM UTC

### Review by @xbhouse - Commented on August 27, 2024 at 01:25 PM UTC

### Review by @Andrewgdewar - Commented on August 28, 2024 at 02:31 PM UTC

### Review by @xbhouse - Approved on August 28, 2024 at 03:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/776*
