---
type: pull_request
number: 1076
title: "SPM-1564 add metrics for request durations per API endpoint"
state: merged
author: tkasparek
created: 2022-08-23T14:36:17Z
updated: 2022-09-01T13:34:30Z
closed: 2022-09-01T13:34:30Z
merged: 2022-09-01T13:34:30Z
base_branch: master
head_branch: request_durations
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1076
---

# Pull Request #1076: SPM-1564 add metrics for request durations per API endpoint

**Author**: @tkasparek
**Created**: August 23, 2022 at 02:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `request_durations`

## Description

Example:
```
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="1"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="1.5"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="1.75"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="2"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="2.5"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="3"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="3.5"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="4"} 10
patchman_engine_manager_request_durations_bucket{endpoint="GET/api/patch/v2/advisories/",le="+Inf"} 10
patchman_engine_manager_request_durations_sum{endpoint="GET/api/patch/v2/advisories/"} 0.021635808
patchman_engine_manager_request_durations_count{endpoint="GET/api/patch/v2/advisories/"} 10
```
Buckets are same as in Vulnerability in order to help us discover API endpoints which are close to 2s (latency target).
Tested with 0.005 bucket to see it works.

---

## Reviews

### Review by @psegedy - Approved on August 23, 2022 at 03:31 PM UTC

### Review by @MichaelMraka - Approved on August 24, 2022 at 01:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1076*
