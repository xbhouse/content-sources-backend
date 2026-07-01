---
type: pull_request
number: 1273
title: "SPM-2119: remove package parsing where not needed"
state: merged
author: psegedy
created: 2023-07-25T12:55:08Z
updated: 2023-07-25T15:20:25Z
closed: 2023-07-25T15:20:25Z
merged: 2023-07-25T15:20:24Z
base_branch: master
head_branch: improve_evaluator2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1273
---

# Pull Request #1273: SPM-2119: remove package parsing where not needed

**Author**: @psegedy
**Created**: July 25, 2023 at 12:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `improve_evaluator2`

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

### Comment by @jira-linking on July 25, 2023 at 12:55 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2119


### Comment by @codecov-commenter on July 25, 2023 at 01:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`83.33%`** and project coverage change: **`+0.06%`** :tada:
> Comparison is base [(`2874520`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/287452097177a6a21643eb2b204a553ee0b75a86?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.68% compared to head [(`c2a6cfe`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.74%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1273      +/-   ##
==========================================
+ Coverage   60.68%   60.74%   +0.06%     
==========================================
  Files         106      106              
  Lines        6735     6741       +6     
==========================================
+ Hits         4087     4095       +8     
+ Misses       2109     2108       -1     
+ Partials      539      538       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.74% <83.33%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.00% <0.00%> (-3.32%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `69.37% <ø> (ø)` | |
| [evaluator/remediations.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `80.00% <ø> (ø)` | |
| [manager/controllers/system\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `31.18% <ø> (ø)` | |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.94% <66.66%> (ø)` | |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.78% <86.95%> (+0.37%)` | :arrow_up: |
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.00% <100.00%> (ø)` | |
| [evaluator/evaluate\_baseline.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `92.59% <100.00%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `79.16% <100.00%> (+1.23%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `83.73% <100.00%> (+1.86%)` | :arrow_up: |
| ... and [1 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1273?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on July 25, 2023 at 02:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1273*
