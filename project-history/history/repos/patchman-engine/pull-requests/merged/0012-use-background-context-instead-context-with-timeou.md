---
type: pull_request
number: 12
title: "Use background context instead context with timeout"
state: merged
author: jiridostal
created: 2019-11-08T15:46:14Z
updated: 2019-11-21T08:59:27Z
closed: 2019-11-14T14:29:38Z
merged: 2019-11-14T14:29:38Z
base_branch: master
head_branch: perf-go
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/12
---

# Pull Request #12: Use background context instead context with timeout

**Author**: @jiridostal
**Created**: November 08, 2019 at 03:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `perf-go`

## Description

Using context with timeout for reading messages caused program to slow down heavily with increasing number of messages
**Before**
> {
  "timestamp": "2019-11-08T15:45:01Z",
  **"duration": 659.337115,**
  "items": 1000,
  "levelname": "info",
  "message": "batch finished",
  **"write/sec": 1.5166748196785493**
}


**After**
> {
  "timestamp": "2019-11-08T15:27:16Z",
  **"duration": 19.86282347,**
  "items": 1000,
  "levelname": "info",
  "message": "batch finished",
  **"write/sec": 50.34530974462716**
}

---

## Reviews

### Review by @semtexzv - Approved on November 13, 2019 at 10:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/12*
