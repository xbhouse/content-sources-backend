---
type: pull_request
number: 1170
title: "SPM-1751: split advisories into installable/available"
state: merged
author: MichaelMraka
created: 2023-02-08T15:15:52Z
updated: 2023-03-02T17:22:47Z
closed: 2023-03-02T17:22:46Z
merged: 2023-03-02T17:22:46Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1170
---

# Pull Request #1170: SPM-1751: split advisories into installable/available

**Author**: @MichaelMraka
**Created**: February 08, 2023 at 03:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @jira-linking on February 08, 2023 at 03:16 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1751


### Comment by @codecov-commenter on March 02, 2023 at 03:33 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`65.88`**% and project coverage change: **`-0.11`** :warning:
> Comparison is base [(`ccb1698`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ccb1698df35bc79a6f480d65c98ad016b5ae8f73?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.93% compared to head [(`5ffcfb6`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.83%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1170      +/-   ##
==========================================
- Coverage   62.93%   62.83%   -0.11%     
==========================================
  Files         103      103              
  Lines        5957     5981      +24     
==========================================
+ Hits         3749     3758       +9     
- Misses       1741     1753      +12     
- Partials      467      470       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.83% <65.88%> (-0.11%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.50% <0.00%> (-0.05%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `70.94% <81.35%> (-4.06%)` | :arrow_down: |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `90.47% <100.00%> (-0.44%)` | :arrow_down: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `80.00% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `90.36% <100.00%> (ø)` | |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.22% <0.00%> (-2.78%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1170?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Commented on February 14, 2023 at 03:57 PM UTC

I'd change `available` to `applicable` since it is used in mocks.
With this PR and current mocks it looks like we don't need any other changes on /advisories endpoint. What am I missing 😄 

### Review by @MichaelMraka - Commented on February 28, 2023 at 02:34 PM UTC

### Review by @psegedy - Approved on March 02, 2023 at 05:20 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1170*
