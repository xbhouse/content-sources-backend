---
type: pull_request
number: 1035
title: "SPM-1616: improve waiting for db upgrade"
state: merged
author: psegedy
created: 2022-07-21T16:05:05Z
updated: 2022-07-22T15:42:50Z
closed: 2022-07-22T15:42:50Z
merged: 2022-07-22T15:42:50Z
base_branch: master
head_branch: init_container_pt2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1035
---

# Pull Request #1035: SPM-1616: improve waiting for db upgrade

**Author**: @psegedy
**Created**: July 21, 2022 at 04:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `init_container_pt2`

## Description

- add env vars to skip updating users and db config
- start migration only if the latest migration is not applied
- add init containers for other pods to wait for migration

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @jira-linking on July 21, 2022 at 04:05 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1616


### Comment by @codecov-commenter on July 22, 2022 at 09:03 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1035?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1035](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1035?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (205c25a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f7eb3f6b89196c9f0d885734370cb8fdeedfd8c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1f7eb3f) will **not change** coverage.
> The diff coverage is `n/a`.

```diff
@@           Coverage Diff           @@
##           master    #1035   +/-   ##
=======================================
  Coverage   61.09%   61.09%           
=======================================
  Files          95       95           
  Lines        5416     5416           
=======================================
  Hits         3309     3309           
  Misses       1680     1680           
  Partials      427      427           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.09% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1035?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1035?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [41211fc...205c25a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1035?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on July 22, 2022 at 12:32 PM UTC

/retest

### Comment by @psegedy on July 22, 2022 at 01:10 PM UTC

/retest

### Comment by @psegedy on July 22, 2022 at 01:13 PM UTC

hmm looks like there is something wrong in init containers, they keep running, need to debug it in ephemeral

### Comment by @psegedy on July 22, 2022 at 01:35 PM UTC

```
exec: \"./database_admin/check-upgraded.sh\": permission denied"
```
🤦 

### Comment by @psegedy on July 22, 2022 at 02:42 PM UTC

@michalslomczynski it should be fixed now

---

## Reviews

### Review by @michalslomczynski - Approved on July 22, 2022 at 03:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1035*
