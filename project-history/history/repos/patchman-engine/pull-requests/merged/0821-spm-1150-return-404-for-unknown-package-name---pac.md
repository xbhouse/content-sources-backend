---
type: pull_request
number: 821
title: "SPM-1150: return 404 for unknown package name - package systems"
state: merged
author: AlesKas
created: 2021-09-09T08:49:15Z
updated: 2021-09-09T19:49:57Z
closed: 2021-09-09T19:49:57Z
merged: 2021-09-09T19:49:57Z
base_branch: master
head_branch: SPM-1150
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/821
---

# Pull Request #821: SPM-1150: return 404 for unknown package name - package systems

**Author**: @AlesKas
**Created**: September 09, 2021 at 08:49 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-1150`

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

### Comment by @codecov-commenter on September 09, 2021 at 08:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#821](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2810ae2) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e14ebd665e4977c9a1c1a7d321531b3a661ff5e7?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e14ebd6) will **increase** coverage by `0.03%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #821      +/-   ##
==========================================
+ Coverage   57.86%   57.89%   +0.03%     
==========================================
  Files          81       81              
  Lines        3899     3902       +3     
==========================================
+ Hits         2256     2259       +3     
  Misses       1325     1325              
  Partials      318      318              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.89% <100.00%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `75.67% <100.00%> (+2.14%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e14ebd6...2810ae2](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/821?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Changes Requested on September 09, 2021 at 09:27 AM UTC

### Review by @AlesKas - Commented on September 09, 2021 at 09:29 AM UTC

### Review by @semtexzv - Approved on September 09, 2021 at 12:36 PM UTC

### Review by @josef-hak - Approved on September 09, 2021 at 07:49 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/821*
