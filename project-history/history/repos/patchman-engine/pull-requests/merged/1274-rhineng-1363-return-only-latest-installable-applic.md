---
type: pull_request
number: 1274
title: "RHINENG-1363: return only latest installable / applicable version of the package"
state: merged
author: MichaelMraka
created: 2023-07-28T14:51:35Z
updated: 2023-08-01T08:09:53Z
closed: 2023-08-01T08:09:52Z
merged: 2023-08-01T08:09:52Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1274
---

# Pull Request #1274: RHINENG-1363: return only latest installable / applicable version of the package

**Author**: @MichaelMraka
**Created**: July 28, 2023 at 02:51 PM UTC
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

### Comment by @jira-linking on July 28, 2023 at 02:51 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1363


### Comment by @codecov-commenter on July 28, 2023 at 02:59 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`84.21%`** and project coverage change: **`-0.11%`** :warning:
> Comparison is base [(`2874520`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/287452097177a6a21643eb2b204a553ee0b75a86?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.68% compared to head [(`a90df07`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.57%.
> Report is 24 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1274      +/-   ##
==========================================
- Coverage   60.68%   60.57%   -0.11%     
==========================================
  Files         106      106              
  Lines        6735     6631     -104     
==========================================
- Hits         4087     4017      -70     
+ Misses       2109     2083      -26     
+ Partials      539      531       -8     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.57% <84.21%> (-0.11%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.00% <0.00%> (-3.32%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `69.37% <ø> (ø)` | |
| [evaluator/remediations.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `80.00% <ø> (ø)` | |
| [manager/controllers/system\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `31.18% <ø> (ø)` | |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.94% <66.66%> (ø)` | |
| [manager/controllers/system\_packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | `60.00% <80.00%> (+1.17%)` | :arrow_up: |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.78% <86.95%> (+0.37%)` | :arrow_up: |
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.00% <100.00%> (ø)` | |
| [evaluator/evaluate\_baseline.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `92.59% <100.00%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `79.16% <100.00%> (+1.23%)` | :arrow_up: |
| ... and [3 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [5 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1274?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 31, 2023 at 08:33 AM UTC

/retest


---

## Reviews

### Review by @psegedy - Approved on July 31, 2023 at 08:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1274*
