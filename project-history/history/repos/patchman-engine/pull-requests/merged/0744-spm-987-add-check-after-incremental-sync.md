---
type: pull_request
number: 744
title: "SPM-987: Add check after incremental sync"
state: merged
author: AlesKas
created: 2021-07-26T09:51:22Z
updated: 2021-08-04T10:26:15Z
closed: 2021-08-04T10:24:55Z
merged: 2021-08-04T10:24:55Z
base_branch: master
head_branch: SPM-987
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/744
---

# Pull Request #744: SPM-987: Add check after incremental sync

**Author**: @AlesKas
**Created**: July 26, 2021 at 09:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-987`

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

### Comment by @codecov-commenter on July 26, 2021 at 09:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#744](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5f787aa) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/de73f90996f4f3b60ce303e4b0cf75439d9f89a0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (de73f90) will **decrease** coverage by `0.28%`.
> The diff coverage is `5.55%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #744      +/-   ##
==========================================
- Coverage   57.06%   56.77%   -0.29%     
==========================================
  Files          81       81              
  Lines        3815     3850      +35     
==========================================
+ Hits         2177     2186       +9     
- Misses       1321     1344      +23     
- Partials      317      320       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.77% <5.55%> (-0.29%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `55.19% <5.55%> (-6.85%)` | :arrow_down: |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `41.66% <0.00%> (-3.79%)` | :arrow_down: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `86.20% <0.00%> (-2.95%)` | :arrow_down: |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `11.47% <0.00%> (-0.60%)` | :arrow_down: |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `75.40% <0.00%> (+1.12%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [de73f90...5f787aa](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/744?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Changes Requested on August 03, 2021 at 09:01 AM UTC

Thanks, some small improvements proposed. And please ensure the new function is used in tests or create one for that. I saw our coverage decreased by that.

### Review by @josef-hak - Commented on August 03, 2021 at 09:11 AM UTC

### Review by @AlesKas - Commented on August 03, 2021 at 09:25 AM UTC

### Review by @josef-hak - Commented on August 04, 2021 at 10:24 AM UTC

Looks good now. Thanks :+1: 

### Review by @josef-hak - Approved on August 04, 2021 at 10:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/744*
