---
type: pull_request
number: 885
title: "SPM-1282: Added VMAAS_CALL_USE_OPTIMISTIC_UPDATES to evaluator"
state: merged
author: josef-hak
created: 2022-01-13T16:31:10Z
updated: 2022-01-14T12:19:03Z
closed: 2022-01-14T11:29:44Z
merged: 2022-01-14T11:29:44Z
base_branch: master
head_branch: opt-upd
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/885
---

# Pull Request #885: SPM-1282: Added VMAAS_CALL_USE_OPTIMISTIC_UPDATES to evaluator

**Author**: @josef-hak
**Created**: January 13, 2022 at 04:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `opt-upd`

## Description

- when enabled, use "optimistic_updates=true" for all vmaas calls, not only for systems with third party repos

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

### Comment by @codecov-commenter on January 13, 2022 at 04:39 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#885](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (654c856) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/69c2ef7108854bf2013561f93f091391a093019c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (69c2ef7) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #885      +/-   ##
==========================================
+ Coverage   58.98%   59.00%   +0.01%     
==========================================
  Files          86       86              
  Lines        4501     4503       +2     
==========================================
+ Hits         2655     2657       +2     
  Misses       1472     1472              
  Partials      374      374              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.00% <100.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.44% <100.00%> (+0.28%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [99c2430...654c856](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/885?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 14, 2022 at 10:21 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/885*
