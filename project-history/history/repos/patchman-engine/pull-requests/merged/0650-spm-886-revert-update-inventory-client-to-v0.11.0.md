---
type: pull_request
number: 650
title: "SPM-886 Revert \"update inventory client to v0.11.0\""
state: merged
author: josef-hak
created: 2021-04-29T09:05:12Z
updated: 2021-05-03T12:54:49Z
closed: 2021-04-29T09:44:36Z
merged: 2021-04-29T09:44:36Z
base_branch: master
head_branch: inv-revert
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/650
---

# Pull Request #650: SPM-886 Revert "update inventory client to v0.11.0"

**Author**: @josef-hak
**Created**: April 29, 2021 at 09:05 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `inv-revert`

## Description

 - This reverts commit 0f65578eae51abefcc615639a4ffe2d9c7bb4f81.
Revert "update code to new inventory client"
 - This reverts commit d57d13850647e588193416035e401c0e6e7735f7.
Revert "update tests to new inventory client"
 - This reverts commit e0a0926c2f2604ed40ba1d933f17e4d797e443ce.

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

### Comment by @codecov-commenter on April 29, 2021 at 09:14 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#650](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8d0f2df) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5c8a822fdd1a6a2383e2949f0a2b17aa6d91d36e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5c8a822) will **increase** coverage by `0.34%`.
> The diff coverage is `68.75%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #650      +/-   ##
==========================================
+ Coverage   58.25%   58.59%   +0.34%     
==========================================
  Files          73       72       -1     
  Lines        3282     3251      -31     
==========================================
- Hits         1912     1905       -7     
+ Misses       1092     1068      -24     
  Partials      278      278              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.59% <68.75%> (+0.34%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.00% <0.00%> (+0.73%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.66% <83.33%> (-0.12%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `60.93% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5c8a822...8d0f2df](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/650?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @semtexzv - Approved on April 29, 2021 at 09:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/650*
