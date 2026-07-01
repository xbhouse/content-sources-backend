---
type: pull_request
number: 966
title: "SPM-1546: Read static data from memory instead of DB"
state: merged
author: michalslomczynski
created: 2022-05-31T12:30:37Z
updated: 2022-06-09T08:22:31Z
closed: 2022-06-09T08:22:30Z
merged: 2022-06-09T08:22:30Z
base_branch: master
head_branch: static-data
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/966
---

# Pull Request #966: SPM-1546: Read static data from memory instead of DB

**Author**: @michalslomczynski
**Created**: May 31, 2022 at 12:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `static-data`

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

### Comment by @codecov-commenter on May 31, 2022 at 01:31 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#966](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f0f5e01) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/486c1e1ed53d040254f49b1754cbf3153be002ee?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (486c1e1) will **increase** coverage by `0.05%`.
> The diff coverage is `88.23%`.

```diff
@@            Coverage Diff             @@
##           master     #966      +/-   ##
==========================================
+ Coverage   61.30%   61.36%   +0.05%     
==========================================
  Files          92       92              
  Lines        4975     4972       -3     
==========================================
+ Hits         3050     3051       +1     
+ Misses       1518     1516       -2     
+ Partials      407      405       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.36% <88.23%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `77.41% <77.77%> (+0.06%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.72% <100.00%> (ø)` | |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `82.60% <100.00%> (+12.91%)` | :arrow_up: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `74.50% <100.00%> (-0.50%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5f4cba5...f0f5e01](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/966?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 02, 2022 at 10:55 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Commented on June 01, 2022 at 09:15 AM UTC

Code looks ok but let's wait for passing CI tests.

It looks like problem with tests / CI env  not the code.

### Review by @MichaelMraka - Approved on June 03, 2022 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/966*
