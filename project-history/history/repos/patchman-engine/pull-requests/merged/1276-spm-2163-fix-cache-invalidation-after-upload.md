---
type: pull_request
number: 1276
title: "SPM-2163: fix cache invalidation after upload"
state: merged
author: psegedy
created: 2023-07-31T14:13:53Z
updated: 2023-08-01T07:47:25Z
closed: 2023-08-01T07:47:25Z
merged: 2023-08-01T07:47:25Z
base_branch: master
head_branch: fix_cache_invalidation
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1276
---

# Pull Request #1276: SPM-2163: fix cache invalidation after upload

**Author**: @psegedy
**Created**: July 31, 2023 at 02:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_cache_invalidation`

## Description

https://gorm.io/docs/update.html#Updates-multiple-columns see the note, that's exactly what I did

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

### Comment by @jira-linking on July 31, 2023 at 02:13 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2163


### Comment by @codecov-commenter on July 31, 2023 at 02:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00%`** and project coverage change: **`-0.02%`** :warning:
> Comparison is base [(`34d3ec7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/34d3ec7f82e4476aeb1aede602bda1a490d467fd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.52% compared to head [(`0a485b7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.50%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1276      +/-   ##
==========================================
- Coverage   60.52%   60.50%   -0.02%     
==========================================
  Files         106      106              
  Lines        6624     6626       +2     
==========================================
  Hits         4009     4009              
- Misses       2083     2085       +2     
  Partials      532      532              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.50% <100.00%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.94% <100.00%> (+0.16%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1276?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 01, 2023 at 07:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1276*
