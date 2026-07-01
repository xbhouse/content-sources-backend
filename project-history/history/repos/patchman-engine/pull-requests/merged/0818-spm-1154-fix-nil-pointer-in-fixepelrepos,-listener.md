---
type: pull_request
number: 818
title: "SPM-1154: Fix nil pointer in fixEpelRepos, listener component"
state: merged
author: josef-hak
created: 2021-09-09T07:58:24Z
updated: 2021-09-20T13:43:50Z
closed: 2021-09-09T08:08:05Z
merged: 2021-09-09T08:08:05Z
base_branch: master
head_branch: fix-list
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/818
---

# Pull Request #818: SPM-1154: Fix nil pointer in fixEpelRepos, listener component

**Author**: @josef-hak
**Created**: September 09, 2021 at 07:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-list`

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

### Comment by @codecov-commenter on September 09, 2021 at 08:06 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#818](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fd302a4) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba2d44576563325bc06c8c2dc11d9f57b359e802?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba2d445) will **increase** coverage by `0.02%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #818      +/-   ##
==========================================
+ Coverage   57.83%   57.86%   +0.02%     
==========================================
  Files          81       81              
  Lines        3897     3899       +2     
==========================================
+ Hits         2254     2256       +2     
  Misses       1325     1325              
  Partials      318      318              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.86% <100.00%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.53% <100.00%> (+0.19%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ba2d445...fd302a4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/818?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/818*
