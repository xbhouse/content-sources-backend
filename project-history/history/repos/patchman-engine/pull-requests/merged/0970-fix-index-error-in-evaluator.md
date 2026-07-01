---
type: pull_request
number: 970
title: "fix: index error in evaluator"
state: merged
author: psegedy
created: 2022-06-07T08:28:35Z
updated: 2022-06-07T16:25:22Z
closed: 2022-06-07T16:25:22Z
merged: 2022-06-07T16:25:22Z
base_branch: master
head_branch: eval_bug
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/970
---

# Pull Request #970: fix: index error in evaluator

**Author**: @psegedy
**Created**: June 07, 2022 at 08:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval_bug`

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

### Comment by @codecov-commenter on June 07, 2022 at 08:35 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#970](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (900a4d1) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a6aa827d8a6c0c781ad496b3f30a60402f015544?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a6aa827) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

```diff
@@            Coverage Diff             @@
##           master     #970      +/-   ##
==========================================
+ Coverage   60.91%   60.93%   +0.01%     
==========================================
  Files          94       94              
  Lines        5102     5104       +2     
==========================================
+ Hits         3108     3110       +2     
  Misses       1585     1585              
  Partials      409      409              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.93% <100.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.37% <100.00%> (+0.25%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ae7cbc0...900a4d1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/970?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 07, 2022 at 09:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/970*
