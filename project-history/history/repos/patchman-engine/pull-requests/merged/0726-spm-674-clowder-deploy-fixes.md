---
type: pull_request
number: 726
title: "SPM-674: Clowder deploy fixes"
state: merged
author: josef-hak
created: 2021-07-07T10:06:01Z
updated: 2021-08-26T18:42:13Z
closed: 2021-07-12T08:56:25Z
merged: 2021-07-12T08:56:25Z
base_branch: master
head_branch: puptoo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/726
---

# Pull Request #726: SPM-674: Clowder deploy fixes

**Author**: @josef-hak
**Created**: July 07, 2021 at 10:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `puptoo`

## Description

because PR CI fails waiting for this service to be ready

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

### Comment by @josef-hak on July 07, 2021 at 11:17 AM UTC

/retest


### Comment by @codecov-commenter on July 07, 2021 at 11:25 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#726](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (cf5e59a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5d0c80cffcfedfb267f25f03a8d45bb58d796413?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5d0c80c) will **decrease** coverage by `0.04%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #726      +/-   ##
==========================================
- Coverage   56.17%   56.12%   -0.05%     
==========================================
  Files          80       80              
  Lines        3692     3695       +3     
==========================================
  Hits         2074     2074              
- Misses       1305     1308       +3     
  Partials      313      313              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.12% <0.00%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5d0c80c...cf5e59a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/726?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on July 07, 2021 at 11:29 AM UTC

/retest

---

## Reviews

### Review by @semtexzv - Approved on July 12, 2021 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/726*
