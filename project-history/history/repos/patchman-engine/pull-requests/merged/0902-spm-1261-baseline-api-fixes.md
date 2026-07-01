---
type: pull_request
number: 902
title: "SPM-1261: Baseline API fixes"
state: merged
author: josef-hak
created: 2022-02-04T15:38:08Z
updated: 2022-02-08T09:28:45Z
closed: 2022-02-08T09:28:19Z
merged: 2022-02-08T09:28:19Z
base_branch: master
head_branch: fix-baseline-list
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/902
---

# Pull Request #902: SPM-1261: Baseline API fixes

**Author**: @josef-hak
**Created**: February 04, 2022 at 03:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-baseline-list`

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

### Comment by @codecov-commenter on February 07, 2022 at 08:49 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#902](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6bb5f2c) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b4017091a1581fe46192871abc1b9d70611ca25?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4b40170) will **increase** coverage by `0.07%`.
> The diff coverage is `90.47%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #902      +/-   ##
==========================================
+ Coverage   58.70%   58.77%   +0.07%     
==========================================
  Files          88       88              
  Lines        4538     4551      +13     
==========================================
+ Hits         2664     2675      +11     
- Misses       1502     1503       +1     
- Partials      372      373       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.77% <90.47%> (+0.07%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `73.33% <33.33%> (-2.86%)` | :arrow_down: |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `68.96% <100.00%> (+1.10%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `72.60% <100.00%> (ø)` | |
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `95.23% <100.00%> (+1.29%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f62df0b...6bb5f2c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/902?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 08, 2022 at 08:55 AM UTC

### Review by @mkholjuraev - Approved on February 08, 2022 at 09:09 AM UTC

I have left my comment, it is just a question. We can merge the PR, if the comment does not make sense

### Review by @josef-hak - Commented on February 08, 2022 at 09:28 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/902*
