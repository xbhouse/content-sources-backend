---
type: pull_request
number: 798
title: "SPM-1059: Return 404 for invalid package name"
state: merged
author: AlesKas
created: 2021-08-25T08:19:39Z
updated: 2021-09-09T08:36:29Z
closed: 2021-08-25T14:44:47Z
merged: 2021-08-25T14:44:47Z
base_branch: master
head_branch: SPM-1059
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/798
---

# Pull Request #798: SPM-1059: Return 404 for invalid package name

**Author**: @AlesKas
**Created**: August 25, 2021 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-1059`

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

### Comment by @codecov-commenter on August 25, 2021 at 08:26 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#798](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (faf16f2) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2c3a67c0528bd780b3afa488482547d1fd7270d2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2c3a67c) will **increase** coverage by `0.03%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #798      +/-   ##
==========================================
+ Coverage   57.68%   57.71%   +0.03%     
==========================================
  Files          81       81              
  Lines        3880     3883       +3     
==========================================
+ Hits         2238     2241       +3     
  Misses       1323     1323              
  Partials      319      319              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.71% <100.00%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_versions.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3ZlcnNpb25zLmdv) | `67.64% <100.00%> (+3.13%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2c3a67c...faf16f2](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/798?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @semtexzv - Approved on August 25, 2021 at 12:25 PM UTC

### Review by @josef-hak - Approved on August 25, 2021 at 02:44 PM UTC

Thanks :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/798*
