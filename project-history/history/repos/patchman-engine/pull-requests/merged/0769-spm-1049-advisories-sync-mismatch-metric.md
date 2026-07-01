---
type: pull_request
number: 769
title: "SPM-1049: Advisories sync mismatch metric"
state: merged
author: AlesKas
created: 2021-08-09T08:39:16Z
updated: 2021-08-10T08:40:54Z
closed: 2021-08-10T08:40:54Z
merged: 2021-08-10T08:40:54Z
base_branch: master
head_branch: SPM-1049
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/769
---

# Pull Request #769: SPM-1049: Advisories sync mismatch metric

**Author**: @AlesKas
**Created**: August 09, 2021 at 08:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-1049`

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

### Comment by @codecov-commenter on August 09, 2021 at 08:45 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#769](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8becf7b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d23c3250201e7f1dcb01c095431a989f8f9b0b75?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d23c325) will **increase** coverage by `0.01%`.
> The diff coverage is `60.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #769      +/-   ##
==========================================
+ Coverage   57.27%   57.29%   +0.01%     
==========================================
  Files          81       81              
  Lines        3867     3871       +4     
==========================================
+ Hits         2215     2218       +3     
- Misses       1330     1331       +1     
  Partials      322      322              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.29% <60.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.28% <0.00%> (-0.29%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `62.26% <100.00%> (+0.72%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [d23c325...8becf7b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/769?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Commented on August 09, 2021 at 12:17 PM UTC

### Review by @AlesKas - Commented on August 09, 2021 at 12:18 PM UTC

### Review by @josef-hak - Commented on August 09, 2021 at 01:50 PM UTC

### Review by @josef-hak - Approved on August 10, 2021 at 08:40 AM UTC

/lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/769*
