---
type: pull_request
number: 777
title: "SPM-1091: Fix advisory type index in evaluate.go"
state: merged
author: josef-hak
created: 2021-08-10T14:42:36Z
updated: 2021-08-26T18:41:56Z
closed: 2021-08-10T15:00:09Z
merged: 2021-08-10T15:00:09Z
base_branch: master
head_branch: mis-pkg
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/777
---

# Pull Request #777: SPM-1091: Fix advisory type index in evaluate.go

**Author**: @josef-hak
**Created**: August 10, 2021 at 02:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `mis-pkg`

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

### Comment by @codecov-commenter on August 10, 2021 at 02:49 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#777](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2698555) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2693a3b2530d6b8258eee693c2f97885ccd5fd31?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2693a3b) will **not change** coverage.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #777   +/-   ##
=======================================
  Coverage   57.45%   57.45%           
=======================================
  Files          81       81           
  Lines        3843     3843           
=======================================
  Hits         2208     2208           
  Misses       1318     1318           
  Partials      317      317           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.45% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.66% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2693a3b...2698555](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/777?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 10, 2021 at 02:48 PM UTC

### Review by @semtexzv - Approved on August 10, 2021 at 02:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/777*
