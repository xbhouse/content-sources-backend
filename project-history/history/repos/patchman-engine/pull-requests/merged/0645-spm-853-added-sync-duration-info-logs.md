---
type: pull_request
number: 645
title: "SPM-853 added sync duration info logs"
state: merged
author: josef-hak
created: 2021-04-28T12:16:47Z
updated: 2021-05-03T12:54:43Z
closed: 2021-04-28T17:04:15Z
merged: 2021-04-28T17:04:15Z
base_branch: master
head_branch: dur-log
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/645
---

# Pull Request #645: SPM-853 added sync duration info logs

**Author**: @josef-hak
**Created**: April 28, 2021 at 12:16 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `dur-log`

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

### Comment by @codecov-commenter on April 28, 2021 at 12:26 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#645](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a41dc6c) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/44afc7fe0cfbd63cc98d167c2938a4854af55891?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (44afc7f) will **increase** coverage by `0.04%`.
> The diff coverage is `80.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #645      +/-   ##
==========================================
+ Coverage   58.31%   58.35%   +0.04%     
==========================================
  Files          72       72              
  Lines        3260     3266       +6     
==========================================
+ Hits         1901     1906       +5     
- Misses       1081     1082       +1     
  Partials      278      278              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.35% <80.00%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `12.96% <0.00%> (-0.25%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `38.73% <50.00%> (-0.55%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `60.93% <100.00%> (+0.62%)` | :arrow_up: |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `71.79% <100.00%> (+0.49%)` | :arrow_up: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `71.42% <100.00%> (+1.21%)` | :arrow_up: |
| [vmaas\_sync/repo\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX3N5bmMuZ28=) | `33.33% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [44afc7f...a41dc6c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/645?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on April 28, 2021 at 01:07 PM UTC

### Review by @josef-hak - Commented on April 28, 2021 at 01:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/645*
