---
type: pull_request
number: 713
title: "SPM-831: Added advisory caches refresh job to vmaas-sync"
state: merged
author: josef-hak
created: 2021-06-21T19:17:23Z
updated: 2021-08-26T18:42:19Z
closed: 2021-06-22T08:23:58Z
merged: 2021-06-22T08:23:58Z
base_branch: master
head_branch: recalc-caches
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/713
---

# Pull Request #713: SPM-831: Added advisory caches refresh job to vmaas-sync

**Author**: @josef-hak
**Created**: June 21, 2021 at 07:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `recalc-caches`

## Description

- ADVISORY_CACHES_WAIT_SECONDS
- ENABLE_REFRESH_ADVISORY_CACHES

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

### Comment by @codecov-commenter on June 21, 2021 at 07:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#713](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c3818ee) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b2c7f8171248a1ec1f5e4b3d258e74af6c592fe2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b2c7f81) will **decrease** coverage by `0.08%`.
> The diff coverage is `32.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #713      +/-   ##
==========================================
- Coverage   56.36%   56.28%   -0.09%     
==========================================
  Files          79       80       +1     
  Lines        3658     3683      +25     
==========================================
+ Hits         2062     2073      +11     
- Misses       1287     1297      +10     
- Partials      309      313       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.28% <32.00%> (-0.09%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/refresh\_advisory\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoX2Fkdmlzb3J5X2NhY2hlcy5nbw==) | `38.88% <38.88%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `41.17% <50.00%> (+0.15%)` | :arrow_up: |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `39.47% <0.00%> (+7.89%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b2c7f81...c3818ee](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/713?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 22, 2021 at 08:22 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/713*
