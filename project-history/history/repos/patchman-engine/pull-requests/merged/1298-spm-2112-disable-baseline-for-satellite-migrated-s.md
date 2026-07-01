---
type: pull_request
number: 1298
title: "SPM-2112: disable baseline for satellite migrated system"
state: merged
author: yungbender
created: 2023-08-22T09:07:10Z
updated: 2023-08-23T11:24:39Z
closed: 2023-08-23T11:24:39Z
merged: 2023-08-23T11:24:39Z
base_branch: master
head_branch: baseline_disable_satellite
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1298
---

# Pull Request #1298: SPM-2112: disable baseline for satellite migrated system

**Author**: @yungbender
**Created**: August 22, 2023 at 09:07 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline_disable_satellite`

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

### Comment by @jira-linking on August 22, 2023 at 09:07 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2112


### Comment by @codecov-commenter on August 22, 2023 at 09:14 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`80.26%`** and project coverage change: **`+0.30%`** :tada:
> Comparison is base [(`9d761af`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9d761af9274bbf9921c7d7079caced72f86092df?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.70% compared to head [(`2b2cec4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.00%.
> Report is 25 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1298      +/-   ##
==========================================
+ Coverage   61.70%   62.00%   +0.30%     
==========================================
  Files         106      106              
  Lines        6651     6672      +21     
==========================================
+ Hits         4104     4137      +33     
+ Misses       2013     2003      -10     
+ Partials      534      532       -2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.00% <80.26%> (+0.30%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `41.46% <0.00%> (-7.52%)` | :arrow_down: |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.55% <50.00%> (-0.23%)` | :arrow_down: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `79.15% <50.00%> (+0.46%)` | :arrow_up: |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.30% <81.81%> (+0.36%)` | :arrow_up: |
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `71.62% <100.00%> (+2.28%)` | :arrow_up: |
| [evaluator/evaluate\_baseline.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `92.59% <100.00%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `79.16% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.54% <100.00%> (+6.74%)` | :arrow_up: |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `80.76% <100.00%> (+2.82%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1298?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 23, 2023 at 11:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1298*
