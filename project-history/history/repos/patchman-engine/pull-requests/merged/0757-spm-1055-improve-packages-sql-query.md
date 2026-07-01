---
type: pull_request
number: 757
title: "SPM-1055: Improve /packages sql query"
state: merged
author: josef-hak
created: 2021-08-02T16:06:19Z
updated: 2021-08-26T18:42:02Z
closed: 2021-08-03T10:13:10Z
merged: 2021-08-03T10:13:09Z
base_branch: master
head_branch: packages-optim
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/757
---

# Pull Request #757: SPM-1055: Improve /packages sql query

**Author**: @josef-hak
**Created**: August 02, 2021 at 04:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `packages-optim`

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

### Comment by @josef-hak on August 02, 2021 at 04:14 PM UTC

/retest

### Comment by @codecov-commenter on August 02, 2021 at 04:29 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#757](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2855592) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4081308cd2111c4673aac4612db4f900feb551ad?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4081308) will **decrease** coverage by `0.02%`.
> The diff coverage is `40.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #757      +/-   ##
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

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.28% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [4081308...2855592](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/757?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on August 02, 2021 at 05:18 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on August 03, 2021 at 10:05 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/757*
