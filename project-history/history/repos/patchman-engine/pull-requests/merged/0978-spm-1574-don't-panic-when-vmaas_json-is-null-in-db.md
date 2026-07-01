---
type: pull_request
number: 978
title: "SPM-1574: don't panic when vmaas_json is null in db"
state: merged
author: psegedy
created: 2022-06-15T13:39:06Z
updated: 2022-06-21T08:47:40Z
closed: 2022-06-21T08:44:53Z
merged: 2022-06-21T08:44:52Z
base_branch: master
head_branch: handle_empty_vmaas
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/978
---

# Pull Request #978: SPM-1574: don't panic when vmaas_json is null in db

**Author**: @psegedy
**Created**: June 15, 2022 at 01:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `handle_empty_vmaas`

## Description

It should not really happen in stage/prod but it is possible to create such db record since vmaas_json is text type without not null constraint. 

Skip the system eval recalc when `VmaasJSON` field (string) is `"" `
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

### Comment by @codecov-commenter on June 15, 2022 at 01:46 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#978](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (47704a9) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f59500b7d6044b526fb667512d8454c025283c98?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f59500b) will **decrease** coverage by `0.04%`.
> The diff coverage is `0.00%`.

```diff
@@            Coverage Diff             @@
##           master     #978      +/-   ##
==========================================
- Coverage   60.64%   60.59%   -0.05%     
==========================================
  Files          94       94              
  Lines        5211     5215       +4     
==========================================
  Hits         3160     3160              
- Misses       1640     1643       +3     
- Partials      411      412       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.59% <0.00%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `67.31% <0.00%> (-1.07%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8747b7c...47704a9](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/978?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 15, 2022 at 02:33 PM UTC

/retest

### Comment by @psegedy on June 20, 2022 at 04:49 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on June 17, 2022 at 12:03 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/978*
