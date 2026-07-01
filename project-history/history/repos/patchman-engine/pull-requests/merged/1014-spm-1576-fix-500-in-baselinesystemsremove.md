---
type: pull_request
number: 1014
title: "SPM-1576: fix 500 in BaselineSystemsRemove"
state: merged
author: psegedy
created: 2022-07-08T14:21:10Z
updated: 2022-07-11T09:25:15Z
closed: 2022-07-11T09:25:15Z
merged: 2022-07-11T09:25:15Z
base_branch: master
head_branch: 500_invalid_inv_id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1014
---

# Pull Request #1014: SPM-1576: fix 500 in BaselineSystemsRemove

**Author**: @psegedy
**Created**: July 08, 2022 at 02:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `500_invalid_inv_id`

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

### Comment by @jira-linking on July 08, 2022 at 02:21 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1576


### Comment by @codecov-commenter on July 08, 2022 at 02:28 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1014](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (81a4daa) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf619f46f8c3ad992a623e23e98cb50c81230a93?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf619f4) will **increase** coverage by `0.46%`.
> The diff coverage is `86.20%`.

```diff
@@            Coverage Diff             @@
##           master    #1014      +/-   ##
==========================================
+ Coverage   60.97%   61.44%   +0.46%     
==========================================
  Files          95       95              
  Lines        5374     5400      +26     
==========================================
+ Hits         3277     3318      +41     
+ Misses       1666     1657       -9     
+ Partials      431      425       -6     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.44% <86.20%> (+0.46%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `78.75% <60.00%> (+0.97%)` | :arrow_up: |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.26% <92.85%> (+4.32%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.83% <100.00%> (+0.20%)` | :arrow_up: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `72.54% <100.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `82.98% <100.00%> (+0.17%)` | :arrow_up: |
| [manager/controllers/baseline\_systems\_remove.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX3JlbW92ZS5nbw==) | `72.72% <100.00%> (+21.00%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.19% <100.00%> (+1.04%)` | :arrow_up: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [90d5e89...81a4daa](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1014?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on July 11, 2022 at 08:33 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1014*
