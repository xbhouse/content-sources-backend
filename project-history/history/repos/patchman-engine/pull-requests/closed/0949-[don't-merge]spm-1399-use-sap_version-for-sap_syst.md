---
type: pull_request
number: 949
title: "[DON'T MERGE]SPM-1399: use sap_version for sap_system filter"
state: closed
author: psegedy
created: 2022-05-04T15:57:20Z
updated: 2022-05-31T12:50:11Z
closed: 2022-05-31T12:50:11Z
base_branch: master
head_branch: sap_system
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/949
---

# Pull Request #949: [DON'T MERGE]SPM-1399: use sap_version for sap_system filter

**Author**: @psegedy
**Created**: May 04, 2022 at 03:57 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `sap_system`

## Description

I'd wait with merging till I get a reply in ESSNTL-1419 because there are systems with `sap_system: true` and without `sap_version` in Stage env
 
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

### Comment by @codecov-commenter on May 05, 2022 at 08:33 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#949](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e9cf14b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5f1dae0d98e90ed82c3e64ee52cf25204ac852c1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5f1dae0) will **increase** coverage by `0.04%`.
> The diff coverage is `100.00%`.

```diff
@@            Coverage Diff             @@
##           master     #949      +/-   ##
==========================================
+ Coverage   60.45%   60.50%   +0.04%     
==========================================
  Files          92       92              
  Lines        4871     4877       +6     
==========================================
+ Hits         2945     2951       +6     
  Misses       1534     1534              
  Partials      392      392              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.50% <100.00%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `61.90% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `85.41% <100.00%> (+0.26%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ad07cf3...e9cf14b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/949?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on May 31, 2022 at 12:50 PM UTC

closing since platform decided to not remove sap_system filed

---

## Reviews

### Review by @michalslomczynski - Changes Requested on May 30, 2022 at 12:15 PM UTC

Does that need update?
https://github.com/psegedy/patchman-engine/blob/46990204276920b9efc992b7c03ca3aefcd539c5/manager/controllers/utils.go#L412

### Review by @michalslomczynski - Changes Requested on May 30, 2022 at 03:03 PM UTC

I added test that must pass and left suggestions how it can be done.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/949*
