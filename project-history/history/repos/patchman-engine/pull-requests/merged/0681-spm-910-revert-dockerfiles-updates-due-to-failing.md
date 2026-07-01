---
type: pull_request
number: 681
title: "SPM-910 Revert dockerfiles updates due to failing osd3 build"
state: merged
author: josef-hak
created: 2021-05-21T20:19:22Z
updated: 2021-08-26T18:42:23Z
closed: 2021-05-24T08:43:59Z
merged: 2021-05-24T08:43:58Z
base_branch: master
head_branch: revert-docker
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/681
---

# Pull Request #681: SPM-910 Revert dockerfiles updates due to failing osd3 build

**Author**: @josef-hak
**Created**: May 21, 2021 at 08:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `revert-docker`

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

### Comment by @codecov-commenter on May 24, 2021 at 08:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#681](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (88895c0) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8119b3ee5755c0a82eeba830f8ae61ba0bf9df64?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8119b3e) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #681   +/-   ##
=======================================
  Coverage   58.50%   58.50%           
=======================================
  Files          72       72           
  Lines        3379     3379           
=======================================
  Hits         1977     1977           
  Misses       1118     1118           
  Partials      284      284           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.50% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8119b3e...88895c0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/681?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on May 24, 2021 at 08:33 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/681*
