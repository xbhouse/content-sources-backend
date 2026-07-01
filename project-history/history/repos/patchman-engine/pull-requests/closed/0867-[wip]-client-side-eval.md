---
type: pull_request
number: 867
title: "[WIP] client side eval"
state: closed
author: MichaelMraka
created: 2021-12-01T15:26:53Z
updated: 2021-12-07T13:26:46Z
closed: 2021-12-07T13:26:46Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/867
---

# Pull Request #867: [WIP] client side eval

**Author**: @MichaelMraka
**Created**: December 01, 2021 at 03:26 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @codecov-commenter on December 01, 2021 at 03:35 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#867](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (267e1f3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2fa031524f07c1944505a3782d611c2dfa314746?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2fa0315) will **decrease** coverage by `0.78%`.
> The diff coverage is `20.98%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #867      +/-   ##
==========================================
- Coverage   58.95%   58.17%   -0.79%     
==========================================
  Files          81       82       +1     
  Lines        4208     4282      +74     
==========================================
+ Hits         2481     2491      +10     
- Misses       1374     1436      +62     
- Partials      353      355       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.17% <20.98%> (-0.79%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <ø> (ø)` | |
| [evaluator/clientside.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2NsaWVudHNpZGUuZ28=) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `66.40% <12.50%> (-5.94%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.63% <18.18%> (-2.71%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `75.00% <100.00%> (+0.49%)` | :arrow_up: |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `84.00% <100.00%> (+0.66%)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `89.65% <100.00%> (+0.76%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `81.78% <100.00%> (+0.28%)` | :arrow_up: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `62.87% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1f314b4...267e1f3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/867?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/867*
