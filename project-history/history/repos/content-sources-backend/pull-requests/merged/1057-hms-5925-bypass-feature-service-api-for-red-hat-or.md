---
type: pull_request
number: 1057
title: "HMS-5925: bypass feature service api for red hat org"
state: merged
author: xbhouse
created: 2025-04-03T13:58:32Z
updated: 2025-04-03T17:04:59Z
closed: 2025-04-03T17:04:59Z
merged: 2025-04-03T17:04:58Z
base_branch: main
head_branch: fix-stage
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1057
---

# Pull Request #1057: HMS-5925: bypass feature service api for red hat org

**Author**: @xbhouse
**Created**: April 03, 2025 at 01:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `fix-stage`

## Description

## Summary

Task workers in stage are crashlooping due to calling the feature service status API with an invalid org ID (-1). This adds a check to bypass that API if using an org ID of -1 (which is only used to delete unneeded repos). 

## Testing steps

1. Run `make repos-import` when the feature service is enabled. Command should succeed and repos should be imported as normal
2. Layered repos (OCP and HA) should be present when listing RH repos or viewing RH repos in the UI if using an account that has access to these repos
3. Previously `make repos-import` failed when the feature service is enabled



---

## Discussion

### Comment by @jlsherrill on April 03, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-5925

### Comment by @xbhouse on April 03, 2025 at 05:04 PM UTC

merging to try and fix stage, playwright tests are failing due to a timeout and were not affected by this change

---

## Reviews

### Review by @rverdile - Commented on April 03, 2025 at 03:59 PM UTC

### Review by @xbhouse - Commented on April 03, 2025 at 04:06 PM UTC

### Review by @rverdile - Approved on April 03, 2025 at 04:59 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1057*
