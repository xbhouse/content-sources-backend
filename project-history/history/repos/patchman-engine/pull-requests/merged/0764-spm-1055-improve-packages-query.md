---
type: pull_request
number: 764
title: "SPM-1055: improve /packages query"
state: merged
author: MichaelMraka
created: 2021-08-04T12:38:20Z
updated: 2021-08-05T09:45:02Z
closed: 2021-08-05T09:45:02Z
merged: 2021-08-05T09:45:02Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/764
---

# Pull Request #764: SPM-1055: improve /packages query

**Author**: @MichaelMraka
**Created**: August 04, 2021 at 12:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

mostly by removing redundant joins to system_platform and package_name
original query
Limit  (cost=609870.56..609870.61 rows=20 width=72) (actual time=3644.560..3648.030 rows=20 loops=1)
improved query
Limit  (cost=105721.03..105721.08 rows=20 width=72)

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

### Comment by @codecov-commenter on August 04, 2021 at 12:45 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#764](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a19834e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b44178eee1d1b6dd0e91849a80e044b29faef90?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4b44178) will **increase** coverage by `0.07%`.
> The diff coverage is `83.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #764      +/-   ##
==========================================
+ Coverage   57.20%   57.28%   +0.07%     
==========================================
  Files          81       81              
  Lines        3870     3870              
==========================================
+ Hits         2214     2217       +3     
+ Misses       1334     1331       -3     
  Partials      322      322              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.28% <83.33%> (+0.07%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `90.32% <100.00%> (+1.03%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [4b44178...a19834e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/764?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on August 04, 2021 at 12:50 PM UTC

/retest

### Comment by @MichaelMraka on August 05, 2021 at 09:21 AM UTC

 tag filtering fixed

---

## Reviews

### Review by @josef-hak - Commented on August 04, 2021 at 01:01 PM UTC

### Review by @josef-hak - Changes Requested on August 04, 2021 at 01:16 PM UTC

We need to update `SELECT id FROM system_platform isp...` subquery to support tag filtering.

### Review by @MichaelMraka - Commented on August 04, 2021 at 02:17 PM UTC

### Review by @josef-hak - Approved on August 05, 2021 at 09:44 AM UTC

Thanks, :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/764*
