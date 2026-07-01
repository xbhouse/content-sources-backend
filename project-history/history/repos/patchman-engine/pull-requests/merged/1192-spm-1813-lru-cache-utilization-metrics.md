---
type: pull_request
number: 1192
title: "SPM-1813: lru cache utilization metrics"
state: merged
author: psegedy
created: 2023-03-16T09:43:51Z
updated: 2023-03-16T10:57:37Z
closed: 2023-03-16T10:57:37Z
merged: 2023-03-16T10:57:37Z
base_branch: master
head_branch: cache_hit_metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1192
---

# Pull Request #1192: SPM-1813: lru cache utilization metrics

**Author**: @psegedy
**Created**: March 16, 2023 at 09:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `cache_hit_metrics`

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

### Comment by @codecov-commenter on March 16, 2023 at 09:51 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`58.65`**% and project coverage change: **`-0.24`** :warning:
> Comparison is base [(`8ba642f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8ba642f2a67ceb91c0bd9343345862ab60828957?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.90% compared to head [(`d23a59d`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.67%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1192      +/-   ##
==========================================
- Coverage   62.90%   62.67%   -0.24%     
==========================================
  Files         103      103              
  Lines        6066     6200     +134     
==========================================
+ Hits         3816     3886      +70     
- Misses       1777     1831      +54     
- Partials      473      483      +10     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.67% <58.65%> (-0.24%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `44.28% <0.00%> (-25.28%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `40.47% <0.00%> (-4.77%)` | :arrow_down: |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `45.00% <26.31%> (-26.43%)` | :arrow_down: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `57.50% <50.00%> (-14.73%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.07% <51.21%> (-16.29%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `85.25% <52.63%> (-1.12%)` | :arrow_down: |
| [evaluator/vmaas\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `54.83% <60.00%> (+0.99%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `74.28% <61.90%> (-8.18%)` | :arrow_down: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `89.36% <90.47%> (-0.97%)` | :arrow_down: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [2 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1192?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 16, 2023 at 10:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1192*
