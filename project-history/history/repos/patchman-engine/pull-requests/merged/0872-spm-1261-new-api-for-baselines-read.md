---
type: pull_request
number: 872
title: "SPM-1261: new api for baselines read"
state: merged
author: mkholjuraev
created: 2021-12-09T10:40:13Z
updated: 2021-12-10T13:11:50Z
closed: 2021-12-10T13:11:50Z
merged: 2021-12-10T13:11:50Z
base_branch: master
head_branch: baseline
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/872
---

# Pull Request #872: SPM-1261: new api for baselines read

**Author**: @mkholjuraev
**Created**: December 09, 2021 at 10:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline`

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


@Josca The new Api is returning correct response, but I was not able to generate new OpenApi docs using './scripts/generate_docs.sh'. This command is always giving current docs.

---

## Discussion

### Comment by @codecov-commenter on December 09, 2021 at 10:47 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#872](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (49d0da7) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0bb4c8dbc70ff4c8eb980b523b869ceffd890b12?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0bb4c8d) will **increase** coverage by `0.36%`.
> The diff coverage is `79.22%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #872      +/-   ##
==========================================
+ Coverage   58.77%   59.13%   +0.36%     
==========================================
  Files          81       83       +2     
  Lines        4245     4322      +77     
==========================================
+ Hits         2495     2556      +61     
- Misses       1396     1405       +9     
- Partials      354      361       +7     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.13% <79.22%> (+0.36%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `70.73% <70.73%> (ø)` | |
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `88.88% <88.88%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [0afcc80...49d0da7](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/872?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on December 10, 2021 at 11:30 AM UTC

/retest

---

## Reviews

### Review by @josef-hak - Changes Requested on December 09, 2021 at 11:44 AM UTC

Good work. Just some changes required.

### Review by @josef-hak - Approved on December 10, 2021 at 01:11 PM UTC

There are some weird errors in vmaas and invenotry. I guess it's not our side problem, merging.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/872*
