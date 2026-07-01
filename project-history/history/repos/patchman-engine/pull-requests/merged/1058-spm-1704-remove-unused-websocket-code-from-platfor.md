---
type: pull_request
number: 1058
title: "SPM-1704: remove unused websocket code from platform mock"
state: merged
author: MichaelMraka
created: 2022-08-15T14:46:38Z
updated: 2022-08-30T11:27:47Z
closed: 2022-08-30T11:27:47Z
merged: 2022-08-30T11:27:47Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1058
---

# Pull Request #1058: SPM-1704: remove unused websocket code from platform mock

**Author**: @MichaelMraka
**Created**: August 15, 2022 at 02:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @codecov-commenter on August 15, 2022 at 02:55 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1058](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2b33f73) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0aee69315ad80944649b71155df2ec56cfb54337?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0aee693) will **decrease** coverage by `0.07%`.
> The diff coverage is `42.30%`.

```diff
@@            Coverage Diff             @@
##           master    #1058      +/-   ##
==========================================
- Coverage   61.48%   61.41%   -0.08%     
==========================================
  Files          98       98              
  Lines        5432     5396      -36     
==========================================
- Hits         3340     3314      -26     
+ Misses       1665     1656       -9     
+ Partials      427      426       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.41% <42.30%> (-0.08%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.82% <0.00%> (+0.12%)` | :arrow_up: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `32.39% <ø> (+0.42%)` | :arrow_up: |
| [base/utils/identity.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9pZGVudGl0eS5nbw==) | `50.00% <ø> (+19.23%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.38% <ø> (-0.10%)` | :arrow_down: |
| [tasks/vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.40% <0.00%> (ø)` | |
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `26.78% <54.54%> (-8.93%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.04% <71.42%> (-1.38%)` | :arrow_down: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1058/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `74.28% <0.00%> (-2.86%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


---

## Reviews

### Review by @psegedy - Commented on August 15, 2022 at 03:13 PM UTC

there is `VMAAS_WS_ADDRESS` in `vmaas_sync/entrypoint.sh` and `conf/common.env`. I think we don't need it either

### Review by @psegedy - Approved on August 16, 2022 at 04:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1058*
