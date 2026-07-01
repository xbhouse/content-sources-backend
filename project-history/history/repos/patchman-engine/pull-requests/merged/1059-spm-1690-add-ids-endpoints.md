---
type: pull_request
number: 1059
title: "SPM-1690: add /ids endpoints"
state: merged
author: psegedy
created: 2022-08-15T15:52:46Z
updated: 2022-09-01T13:34:02Z
closed: 2022-09-01T13:34:02Z
merged: 2022-09-01T13:34:02Z
base_branch: master
head_branch: ids_endpoints
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1059
---

# Pull Request #1059: SPM-1690: add /ids endpoints

**Author**: @psegedy
**Created**: August 15, 2022 at 03:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `ids_endpoints`

## Description

add /ids endpoints returning only advisory/system ids

TODO:
- [x] unit tests
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

### Comment by @codecov-commenter on August 16, 2022 at 09:05 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.59**% // Head: **61.66**% // Increases project coverage by **`+0.07%`** :tada:
> Coverage data is based on head [(`dd1d378`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`485b23a`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/485b23a4a62f19ee1a8b48ccfab58dcff7291c94?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 67.70% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1059      +/-   ##
==========================================
+ Coverage   61.59%   61.66%   +0.07%     
==========================================
  Files          97       97              
  Lines        5411     5494      +83     
==========================================
+ Hits         3333     3388      +55     
- Misses       1651     1668      +17     
- Partials      427      438      +11     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.66% <67.70%> (+0.07%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `74.62% <25.00%> (-3.16%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `67.92% <57.57%> (-3.88%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.41% <59.37%> (-2.59%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `74.60% <66.66%> (-2.95%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `89.61% <76.00%> (-4.24%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.25% <80.00%> (-0.19%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.00% <85.00%> (-0.86%)` | :arrow_down: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1059?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on August 30, 2022 at 07:41 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on August 24, 2022 at 09:46 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1059*
