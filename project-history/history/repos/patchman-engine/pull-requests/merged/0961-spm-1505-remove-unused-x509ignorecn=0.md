---
type: pull_request
number: 961
title: "SPM-1505: remove unused x509ignoreCN=0"
state: merged
author: MichaelMraka
created: 2022-05-17T09:32:57Z
updated: 2022-05-17T10:52:31Z
closed: 2022-05-17T10:52:30Z
merged: 2022-05-17T10:52:30Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/961
---

# Pull Request #961: SPM-1505: remove unused x509ignoreCN=0

**Author**: @MichaelMraka
**Created**: May 17, 2022 at 09:32 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

## Description

go 1.17 no more recognize this option

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

### Comment by @codecov-commenter on May 17, 2022 at 09:40 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#961](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (dc6871e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f1a89b80985fe10be4b8a756c3dcf9face1e70cd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f1a89b8) will **increase** coverage by `0.60%`.
> The diff coverage is `n/a`.

```diff
@@            Coverage Diff             @@
##           master     #961      +/-   ##
==========================================
+ Coverage   60.60%   61.20%   +0.60%     
==========================================
  Files          92       91       -1     
  Lines        4871     4823      -48     
==========================================
  Hits         2952     2952              
+ Misses       1523     1475      -48     
  Partials      396      396              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.20% <ø> (+0.60%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `70.83% <ø> (+14.16%)` | :arrow_up: |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `65.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [4840eda...dc6871e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/961?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @michalslomczynski - Approved on May 17, 2022 at 10:48 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/961*
