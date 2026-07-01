---
type: pull_request
number: 816
title: "SPM-1132: Return 400 when using incorrect filter"
state: merged
author: AlesKas
created: 2021-09-08T11:59:00Z
updated: 2021-09-09T08:36:34Z
closed: 2021-09-08T13:55:53Z
merged: 2021-09-08T13:55:53Z
base_branch: master
head_branch: SPM-1132
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/816
---

# Pull Request #816: SPM-1132: Return 400 when using incorrect filter

**Author**: @AlesKas
**Created**: September 08, 2021 at 11:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-1132`

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

### Comment by @codecov-commenter on September 08, 2021 at 12:06 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#816](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bb29d3c) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/9ef3bfd49d0d784908c4fa909cde7fa033aa5e4e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9ef3bfd) will **increase** coverage by `0.15%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #816      +/-   ##
==========================================
+ Coverage   57.68%   57.83%   +0.15%     
==========================================
  Files          81       81              
  Lines        3890     3897       +7     
==========================================
+ Hits         2244     2254      +10     
+ Misses       1327     1325       -2     
+ Partials      319      318       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.83% <100.00%> (+0.15%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `81.16% <100.00%> (+1.99%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9ef3bfd...bb29d3c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/816?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Approved on September 08, 2021 at 01:55 PM UTC

Looks good. Thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/816*
