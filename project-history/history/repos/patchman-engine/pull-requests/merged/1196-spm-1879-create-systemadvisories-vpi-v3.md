---
type: pull_request
number: 1196
title: "SPM-1879: create /system/advisories VPI v3 "
state: merged
author: MichaelMraka
created: 2023-03-17T15:44:10Z
updated: 2023-03-17T16:47:13Z
closed: 2023-03-17T16:47:12Z
merged: 2023-03-17T16:47:12Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1196
---

# Pull Request #1196: SPM-1879: create /system/advisories VPI v3 

**Author**: @MichaelMraka
**Created**: March 17, 2023 at 03:44 PM UTC
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

### Comment by @codecov-commenter on March 17, 2023 at 03:54 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`52.63`**% and project coverage change: **`-0.77`** :warning:
> Comparison is base [(`8ba642f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8ba642f2a67ceb91c0bd9343345862ab60828957?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.90% compared to head [(`44a2141`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.14%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1196      +/-   ##
==========================================
- Coverage   62.90%   62.14%   -0.77%     
==========================================
  Files         103      103              
  Lines        6066     6205     +139     
==========================================
+ Hits         3816     3856      +40     
- Misses       1777     1857      +80     
- Partials      473      492      +19     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.14% <52.63%> (-0.77%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `44.28% <0.00%> (-25.28%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `40.47% <0.00%> (-4.77%)` | :arrow_down: |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `45.00% <26.31%> (-26.43%)` | :arrow_down: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `63.82% <33.33%> (-26.50%)` | :arrow_down: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `57.50% <50.00%> (-14.73%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `68.57% <52.38%> (-13.89%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `85.01% <52.63%> (-1.37%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.07% <53.48%> (-16.29%)` | :arrow_down: |
| [evaluator/vmaas\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `54.83% <60.00%> (+0.99%)` | :arrow_up: |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [3 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1196?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on March 17, 2023 at 04:47 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1196*
