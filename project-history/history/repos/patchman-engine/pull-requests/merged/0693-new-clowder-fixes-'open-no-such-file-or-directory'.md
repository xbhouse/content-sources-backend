---
type: pull_request
number: 693
title: "new clowder fixes 'open : no such file or directory'"
state: merged
author: MichaelMraka
created: 2021-06-03T11:55:59Z
updated: 2021-06-03T13:22:17Z
closed: 2021-06-03T13:22:17Z
merged: 2021-06-03T13:22:17Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/693
---

# Pull Request #693: new clowder fixes 'open : no such file or directory'

**Author**: @MichaelMraka
**Created**: June 03, 2021 at 11:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @codecov-commenter on June 03, 2021 at 12:02 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#693](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f35486f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/3302a2555d095ae635e24aa3cb5cfabf134275ce?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3302a25) will **decrease** coverage by `0.18%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #693      +/-   ##
==========================================
- Coverage   57.84%   57.65%   -0.19%     
==========================================
  Files          73       73              
  Lines        3428     3441      +13     
==========================================
+ Hits         1983     1984       +1     
- Misses       1162     1174      +12     
  Partials      283      283              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.65% <ø> (-0.19%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.39% <0.00%> (-0.31%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `75.00% <0.00%> (+2.27%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3302a25...f35486f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/693?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Approved on June 03, 2021 at 01:22 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/693*
