---
type: pull_request
number: 624
title: "SPM-853 Refactored \"vmaas_sync/advisory_sync.go\" process"
state: merged
author: josef-hak
created: 2021-04-20T07:03:04Z
updated: 2021-04-21T08:44:17Z
closed: 2021-04-21T06:59:50Z
merged: 2021-04-21T06:59:50Z
base_branch: master
head_branch: sync-refactor
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/624
---

# Pull Request #624: SPM-853 Refactored "vmaas_sync/advisory_sync.go" process

**Author**: @josef-hak
**Created**: April 20, 2021 at 07:03 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sync-refactor`

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

### Comment by @codecov-commenter on April 20, 2021 at 07:14 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#624](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (011d34f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e9370e580fbd83c00a884c98b56802341881ea0a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e9370e5) will **decrease** coverage by `0.00%`.
> The diff coverage is `64.95%`.

> :exclamation: Current head 011d34f differs from pull request most recent head b9be810. Consider uploading reports for the commit b9be810 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #624      +/-   ##
==========================================
- Coverage   58.73%   58.72%   -0.01%     
==========================================
  Files          72       72              
  Lines        3206     3244      +38     
==========================================
+ Hits         1883     1905      +22     
- Misses       1054     1062       +8     
- Partials      269      277       +8     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.72% <64.95%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `34.40% <0.00%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `61.93% <65.51%> (-1.32%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e9370e5...b9be810](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/624?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on April 21, 2021 at 06:59 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/624*
