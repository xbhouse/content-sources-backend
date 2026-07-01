---
type: pull_request
number: 950
title: "SPM-1469: allow running tests with -tags dynamic"
state: closed
author: psegedy
created: 2022-05-05T09:37:34Z
updated: 2022-05-17T07:47:50Z
closed: 2022-05-17T07:47:49Z
base_branch: master
head_branch: tests_arm
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/950
---

# Pull Request #950: SPM-1469: allow running tests with -tags dynamic

**Author**: @psegedy
**Created**: May 05, 2022 at 09:37 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `tests_arm`

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

### Comment by @codecov-commenter on May 05, 2022 at 09:44 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/950?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#950](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/950?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1879745) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5f1dae0d98e90ed82c3e64ee52cf25204ac852c1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5f1dae0) will **not change** coverage.
> The diff coverage is `n/a`.

```diff
@@           Coverage Diff           @@
##           master     #950   +/-   ##
=======================================
  Coverage   60.45%   60.45%           
=======================================
  Files          92       92           
  Lines        4871     4871           
=======================================
  Hits         2945     2945           
  Misses       1534     1534           
  Partials      392      392           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/950?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/950?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ad07cf3...1879745](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/950?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on May 10, 2022 at 03:50 PM UTC

/retest

### Comment by @psegedy on May 10, 2022 at 04:13 PM UTC

should not be needed after https://github.com/RedHatInsights/patchman-engine/pull/955

### Comment by @MichaelMraka on May 17, 2022 at 07:47 AM UTC

#955 merged, closing this one.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/950*
