---
type: pull_request
number: 728
title: "SPM-991: Fix latest_package_cache view"
state: merged
author: AlesKas
created: 2021-07-12T09:57:01Z
updated: 2021-08-04T10:27:20Z
closed: 2021-08-04T10:27:20Z
merged: 2021-08-04T10:27:20Z
base_branch: master
head_branch: SPM-991
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/728
---

# Pull Request #728: SPM-991: Fix latest_package_cache view

**Author**: @AlesKas
**Created**: July 12, 2021 at 09:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-991`

## Description

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

### Comment by @codecov-commenter on July 19, 2021 at 03:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#728](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e6e10b5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/11202b0411496bbea0d366aa1ab0942828e33ef4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (11202b0) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #728   +/-   ##
=======================================
  Coverage   57.03%   57.03%           
=======================================
  Files          81       81           
  Lines        3833     3833           
=======================================
  Hits         2186     2186           
  Misses       1329     1329           
  Partials      318      318           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.03% <0.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/refresh.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoLmdv) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [11202b0...e6e10b5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/728?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on July 29, 2021 at 01:57 PM UTC

/retest

---

## Reviews

### Review by @josef-hak - Changes Requested on July 15, 2021 at 11:31 AM UTC

some conflict appeared, please fix

### Review by @semtexzv - Approved on July 21, 2021 at 12:33 PM UTC

### Review by @josef-hak - Approved on July 19, 2021 at 12:10 PM UTC

Looks good, thanks :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/728*
