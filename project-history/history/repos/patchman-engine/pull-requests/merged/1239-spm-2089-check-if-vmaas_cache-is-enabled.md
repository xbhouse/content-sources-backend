---
type: pull_request
number: 1239
title: "SPM-2089: check if vmaas_cache is enabled"
state: merged
author: psegedy
created: 2023-06-09T12:41:25Z
updated: 2023-06-12T08:30:34Z
closed: 2023-06-12T08:30:33Z
merged: 2023-06-12T08:30:33Z
base_branch: master
head_branch: vmaas_cache_fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1239
---

# Pull Request #1239: SPM-2089: check if vmaas_cache is enabled

**Author**: @psegedy
**Created**: June 09, 2023 at 12:41 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas_cache_fix`

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

### Comment by @jira-linking on June 09, 2023 at 12:41 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2089


### Comment by @codecov-commenter on June 09, 2023 at 12:58 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`37.50`**% and project coverage change: **`-0.04`** :warning:
> Comparison is base [(`12c1603`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/12c1603b8ce434162bd6334fb57a33498778828e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.62% compared to head [(`77689b0`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.58%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1239      +/-   ##
==========================================
- Coverage   61.62%   61.58%   -0.04%     
==========================================
  Files         105      105              
  Lines        6475     6477       +2     
==========================================
- Hits         3990     3989       -1     
- Misses       1968     1971       +3     
  Partials      517      517              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.58% <37.50%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/vmaas\_cache.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `61.76% <0.00%> (-1.88%)` | :arrow_down: |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `62.68% <50.00%> (ø)` | |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.80% <100.00%> (+0.08%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1239?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 12, 2023 at 08:11 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1239*
