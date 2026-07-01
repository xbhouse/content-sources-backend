---
type: pull_request
number: 1200
title: "SPM-1985: show stale systems on baselines"
state: merged
author: psegedy
created: 2023-03-23T10:57:57Z
updated: 2023-03-27T09:20:00Z
closed: 2023-03-27T09:20:00Z
merged: 2023-03-27T09:19:59Z
base_branch: master
head_branch: baselines_stale
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1200
---

# Pull Request #1200: SPM-1985: show stale systems on baselines

**Author**: @psegedy
**Created**: March 23, 2023 at 10:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baselines_stale`

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

### Comment by @codecov-commenter on March 23, 2023 at 11:08 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`60.00`**% and project coverage change: **`-0.01`** :warning:
> Comparison is base [(`b42e944`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b42e9449d607af037aa9e0a69bce60a436608fbb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.11% compared to head [(`bc3ea2d`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.10%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1200      +/-   ##
==========================================
- Coverage   62.11%   62.10%   -0.01%     
==========================================
  Files         103      103              
  Lines        6205     6210       +5     
==========================================
+ Hits         3854     3857       +3     
- Misses       1859     1860       +1     
- Partials      492      493       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.10% <60.00%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `84.00% <ø> (-0.22%)` | :arrow_down: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `62.32% <55.55%> (-0.99%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `56.57% <100.00%> (-0.57%)` | :arrow_down: |

... and [1 file with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1200?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 23, 2023 at 02:28 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1200*
