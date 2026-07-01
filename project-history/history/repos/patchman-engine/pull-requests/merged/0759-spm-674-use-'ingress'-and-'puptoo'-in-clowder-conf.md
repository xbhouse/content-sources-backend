---
type: pull_request
number: 759
title: "SPM-674: Use 'ingress' and 'puptoo' in Clowder config"
state: merged
author: josef-hak
created: 2021-08-03T11:24:15Z
updated: 2021-08-26T18:41:14Z
closed: 2021-08-03T11:45:50Z
merged: 2021-08-03T11:45:50Z
base_branch: master
head_branch: clowder-ing-pup
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/759
---

# Pull Request #759: SPM-674: Use 'ingress' and 'puptoo' in Clowder config

**Author**: @josef-hak
**Created**: August 03, 2021 at 11:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `clowder-ing-pup`

## Description

- these services failed before when being deployed with

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

### Comment by @codecov-commenter on August 03, 2021 at 11:31 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#759](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8e0422d) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/51a44e7fb8c70ea82c3dee33edb12b8c73363ebd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (51a44e7) will **decrease** coverage by `0.02%`.
> The diff coverage is `40.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #759      +/-   ##
==========================================
- Coverage   57.24%   57.21%   -0.03%     
==========================================
  Files          81       81              
  Lines        3850     3852       +2     
==========================================
  Hits         2204     2204              
- Misses       1328     1330       +2     
  Partials      318      318              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.21% <40.00%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.28% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [796b79d...8e0422d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/759?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 03, 2021 at 11:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/759*
