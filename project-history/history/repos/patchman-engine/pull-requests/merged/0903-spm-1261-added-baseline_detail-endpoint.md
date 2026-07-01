---
type: pull_request
number: 903
title: "SPM-1261: Added baseline_detail endpoint"
state: merged
author: josef-hak
created: 2022-02-07T18:24:41Z
updated: 2022-02-08T09:21:25Z
closed: 2022-02-08T09:21:24Z
merged: 2022-02-08T09:21:24Z
base_branch: master
head_branch: baseline-detail
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/903
---

# Pull Request #903: SPM-1261: Added baseline_detail endpoint

**Author**: @josef-hak
**Created**: February 07, 2022 at 06:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-detail`

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

### Comment by @codecov-commenter on February 07, 2022 at 06:31 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#903](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (cf1ab3f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b4017091a1581fe46192871abc1b9d70611ca25?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4b40170) will **increase** coverage by `0.20%`.
> The diff coverage is `79.06%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #903      +/-   ##
==========================================
+ Coverage   58.70%   58.90%   +0.20%     
==========================================
  Files          88       89       +1     
  Lines        4538     4580      +42     
==========================================
+ Hits         2664     2698      +34     
- Misses       1502     1508       +6     
- Partials      372      374       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.90% <79.06%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `87.17% <87.17%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f62df0b...cf1ab3f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/903?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 08, 2022 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/903*
