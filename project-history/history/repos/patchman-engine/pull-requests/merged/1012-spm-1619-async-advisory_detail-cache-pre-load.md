---
type: pull_request
number: 1012
title: "SPM-1619: async advisory_detail cache pre-load"
state: merged
author: psegedy
created: 2022-07-08T13:25:09Z
updated: 2022-07-08T14:26:40Z
closed: 2022-07-08T14:26:40Z
merged: 2022-07-08T14:26:40Z
base_branch: master
head_branch: pre-load-async
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1012
---

# Pull Request #1012: SPM-1619: async advisory_detail cache pre-load

**Author**: @psegedy
**Created**: July 08, 2022 at 01:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pre-load-async`

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

### Comment by @jira-linking on July 08, 2022 at 01:25 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1619


### Comment by @codecov-commenter on July 08, 2022 at 01:33 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1012](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (95446da) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf619f46f8c3ad992a623e23e98cb50c81230a93?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf619f4) will **increase** coverage by `0.34%`.
> The diff coverage is `85.18%`.

```diff
@@            Coverage Diff             @@
##           master    #1012      +/-   ##
==========================================
+ Coverage   60.97%   61.32%   +0.34%     
==========================================
  Files          95       95              
  Lines        5374     5396      +22     
==========================================
+ Hits         3277     3309      +32     
+ Misses       1666     1661       -5     
+ Partials      431      426       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.32% <85.18%> (+0.34%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `78.75% <60.00%> (+0.97%)` | :arrow_up: |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.26% <92.85%> (+4.32%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.83% <100.00%> (+0.20%)` | :arrow_up: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `72.54% <100.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `82.98% <100.00%> (+0.17%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.19% <100.00%> (+1.04%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `79.24% <0.00%> (+2.51%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [90d5e89...95446da](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1012?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 08, 2022 at 01:34 PM UTC

@MichaelMraka what do you think, would it help with cache pre-load on pod start?

### Comment by @MichaelMraka on July 08, 2022 at 02:10 PM UTC

Looks nice. Even if some advisory is loaded as a result of parallel /advisories API hit, it should be safe to load it once again in preload.

---

## Reviews

### Review by @MichaelMraka - Approved on July 08, 2022 at 02:05 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1012*
