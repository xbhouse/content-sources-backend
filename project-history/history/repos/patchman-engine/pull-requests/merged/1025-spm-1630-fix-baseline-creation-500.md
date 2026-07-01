---
type: pull_request
number: 1025
title: "SPM-1630: fix baseline creation 500"
state: merged
author: psegedy
created: 2022-07-14T10:02:57Z
updated: 2022-07-15T14:04:24Z
closed: 2022-07-15T14:04:24Z
merged: 2022-07-15T14:04:24Z
base_branch: master
head_branch: baseline_500
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1025
---

# Pull Request #1025: SPM-1630: fix baseline creation 500

**Author**: @psegedy
**Created**: July 14, 2022 at 10:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline_500`

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

### Comment by @jira-linking on July 14, 2022 at 10:02 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1630


### Comment by @codecov-commenter on July 14, 2022 at 10:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1025](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2d883c5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/af7aeaa935907531686a4c89f0c742338f3b5028?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (af7aeaa) will **decrease** coverage by `0.05%`.
> The diff coverage is `57.62%`.

```diff
@@            Coverage Diff             @@
##           master    #1025      +/-   ##
==========================================
- Coverage   61.32%   61.26%   -0.06%     
==========================================
  Files          95       95              
  Lines        5396     5414      +18     
==========================================
+ Hits         3309     3317       +8     
- Misses       1661     1673      +12     
+ Partials      426      424       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.26% <57.62%> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.20% <0.00%> (-0.13%)` | :arrow_down: |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.69% <0.00%> (-0.40%)` | :arrow_down: |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `25.35% <0.00%> (-2.35%)` | :arrow_down: |
| [base/utils/log.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9sb2cuZ28=) | `75.00% <0.00%> (-18.11%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `78.88% <87.50%> (+0.13%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.48% <100.00%> (+0.64%)` | :arrow_up: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `73.07% <100.00%> (+0.52%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `79.87% <100.00%> (+0.62%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `93.84% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `71.79% <100.00%> (-0.71%)` | :arrow_down: |
| ... and [13 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [847e4f2...2d883c5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1025?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on July 14, 2022 at 01:42 PM UTC

### Review by @MichaelMraka - Approved on July 15, 2022 at 02:02 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1025*
