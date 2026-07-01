---
type: pull_request
number: 417
title: "Indicate engine version in docs and API"
state: merged
author: MichaelMraka
created: 2020-11-06T13:24:58Z
updated: 2020-11-09T08:39:30Z
closed: 2020-11-09T08:39:30Z
merged: 2020-11-09T08:39:30Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/417
---

# Pull Request #417: Indicate engine version in docs and API

**Author**: @MichaelMraka
**Created**: November 06, 2020 at 01:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

## Description

engine version is shown in metrics as
```
# HELP patchman_engine_core_info Patchman project deployment information
# TYPE patchman_engine_core_info gauge
patchman_engine_core_info{version="v1.4.2"} 1

```

---

## Reviews

### Review by @josef-hak - Changes Requested on November 06, 2020 at 01:36 PM UTC

Looks good, just a few notes how to improve that. Thanks.

### Review by @MichaelMraka - Commented on November 06, 2020 at 03:11 PM UTC

### Review by @MichaelMraka - Commented on November 06, 2020 at 03:11 PM UTC

### Review by @MichaelMraka - Commented on November 06, 2020 at 03:11 PM UTC

### Review by @MichaelMraka - Commented on November 06, 2020 at 03:12 PM UTC

### Review by @MichaelMraka - Commented on November 06, 2020 at 03:17 PM UTC

### Review by @josef-hak - Approved on November 09, 2020 at 08:39 AM UTC

great, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/417*
