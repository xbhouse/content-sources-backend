---
type: pull_request
number: 801
title: "SPM-1125: Fix Clowder vars exporting"
state: merged
author: josef-hak
created: 2021-08-26T09:55:21Z
updated: 2021-08-26T18:39:42Z
closed: 2021-08-26T18:33:18Z
merged: 2021-08-26T18:33:18Z
base_branch: master
head_branch: pr
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/801
---

# Pull Request #801: SPM-1125: Fix Clowder vars exporting

**Author**: @josef-hak
**Created**: August 26, 2021 at 09:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr`

## Description

- do not format values using escape characters

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

### Comment by @codecov-commenter on August 26, 2021 at 10:23 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#801](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4b9e15a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1fe49e70a71383459f6bfc35fdb0b97b001c87f0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1fe49e7) will **decrease** coverage by `0.05%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #801      +/-   ##
==========================================
- Coverage   57.71%   57.65%   -0.06%     
==========================================
  Files          81       81              
  Lines        3883     3887       +4     
==========================================
  Hits         2241     2241              
- Misses       1323     1327       +4     
  Partials      319      319              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.65% <0.00%> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1fe49e7...4b9e15a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/801?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/801*
