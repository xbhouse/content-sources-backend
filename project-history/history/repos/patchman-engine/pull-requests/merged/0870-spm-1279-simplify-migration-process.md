---
type: pull_request
number: 870
title: "SPM-1279: Simplify migration process"
state: merged
author: josef-hak
created: 2021-12-06T19:53:14Z
updated: 2021-12-07T14:36:52Z
closed: 2021-12-07T11:35:20Z
merged: 2021-12-07T11:35:20Z
base_branch: master
head_branch: nologin
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/870
---

# Pull Request #870: SPM-1279: Simplify migration process

**Author**: @josef-hak
**Created**: December 06, 2021 at 07:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `nologin`

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

### Comment by @codecov-commenter on December 06, 2021 at 08:14 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#870](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (25bc502) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2fa031524f07c1944505a3782d611c2dfa314746?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2fa0315) will **decrease** coverage by `0.18%`.
> The diff coverage is `45.23%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #870      +/-   ##
==========================================
- Coverage   58.95%   58.77%   -0.19%     
==========================================
  Files          81       81              
  Lines        4208     4245      +37     
==========================================
+ Hits         2481     2495      +14     
- Misses       1374     1396      +22     
- Partials      353      354       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.77% <45.23%> (-0.19%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `66.40% <12.50%> (-5.94%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `75.00% <100.00%> (+0.49%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `71.79% <100.00%> (+0.74%)` | :arrow_up: |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `84.00% <100.00%> (+0.66%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `76.59% <100.00%> (+0.50%)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `90.00% <100.00%> (+1.11%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.43% <100.00%> (+0.43%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `81.78% <100.00%> (+0.28%)` | :arrow_up: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `62.87% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8bcafc5...25bc502](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/870?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on December 07, 2021 at 08:39 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on December 07, 2021 at 09:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/870*
