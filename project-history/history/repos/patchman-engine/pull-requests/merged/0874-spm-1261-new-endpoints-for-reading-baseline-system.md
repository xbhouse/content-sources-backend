---
type: pull_request
number: 874
title: "SPM-1261: new endpoints for reading baseline systems"
state: merged
author: mkholjuraev
created: 2021-12-10T10:57:28Z
updated: 2021-12-16T14:26:22Z
closed: 2021-12-16T14:26:22Z
merged: 2021-12-16T14:26:22Z
base_branch: master
head_branch: baseline-systems
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/874
---

# Pull Request #874: SPM-1261: new endpoints for reading baseline systems

**Author**: @mkholjuraev
**Created**: December 10, 2021 at 10:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-systems`

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

### Comment by @codecov-commenter on December 10, 2021 at 11:04 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#874](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c639777) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0bb4c8dbc70ff4c8eb980b523b869ceffd890b12?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0bb4c8d) will **increase** coverage by `0.16%`.
> The diff coverage is `75.60%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #874      +/-   ##
==========================================
+ Coverage   58.77%   58.93%   +0.16%     
==========================================
  Files          81       82       +1     
  Lines        4245     4286      +41     
==========================================
+ Hits         2495     2526      +31     
- Misses       1396     1402       +6     
- Partials      354      358       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.93% <75.60%> (+0.16%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `75.60% <75.60%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [0afcc80...c639777](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/874?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on December 10, 2021 at 11:46 AM UTC

/retest

### Comment by @digitronik on December 14, 2021 at 07:09 AM UTC

/retest

### Comment by @josef-hak on December 14, 2021 at 08:47 AM UTC

/retest

---

## Reviews

### Review by @josef-hak - Changes Requested on December 14, 2021 at 08:57 AM UTC

Overall good, thanks. Added some comments what to improve.

### Review by @josef-hak - Approved on December 16, 2021 at 02:26 PM UTC

Great, looks good. Thanks :+1: 


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/874*
