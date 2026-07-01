---
type: pull_request
number: 933
title: "SPM-1393: Fix multiple tag filters"
state: closed
author: michalslomczynski
created: 2022-03-30T09:11:17Z
updated: 2022-03-30T14:22:09Z
closed: 2022-03-30T14:22:09Z
base_branch: master
head_branch: fix-tags-add
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/933
---

# Pull Request #933: SPM-1393: Fix multiple tag filters

**Author**: @michalslomczynski
**Created**: March 30, 2022 at 09:11 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-tags-add`

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

### Comment by @codecov-commenter on March 30, 2022 at 09:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#933](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (53d4f78) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/91c8b93848852d6325274f87d28faf50ab3c9313?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (91c8b93) will **increase** coverage by `0.06%`.
> The diff coverage is `61.11%`.

```diff
@@            Coverage Diff             @@
##           master     #933      +/-   ##
==========================================
+ Coverage   60.12%   60.19%   +0.06%     
==========================================
  Files          91       91              
  Lines        4800     4833      +33     
==========================================
+ Hits         2886     2909      +23     
- Misses       1523     1532       +9     
- Partials      391      392       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.19% <61.11%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `9.70% <0.00%> (-0.30%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `85.80% <100.00%> (+0.35%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.25% <0.00%> (+0.32%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [12311f2...53d4f78](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/933?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on March 30, 2022 at 02:22 PM UTC

Supporting multiple values for the same tag makes a little sense because it can't be equal to two different values at the same time.
Based on our discussion I'm closing this ticket for now.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/933*
