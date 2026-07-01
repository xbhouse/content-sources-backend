---
type: pull_request
number: 1248
title: "SPM-2093: mark functions as immutable"
state: merged
author: MichaelMraka
created: 2023-06-21T14:13:05Z
updated: 2023-06-26T12:28:47Z
closed: 2023-06-26T12:28:47Z
merged: 2023-06-26T12:28:46Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1248
---

# Pull Request #1248: SPM-2093: mark functions as immutable

**Author**: @MichaelMraka
**Created**: June 21, 2023 at 02:13 PM UTC
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

### Comment by @jira-linking on June 21, 2023 at 02:13 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2093


### Comment by @codecov-commenter on June 21, 2023 at 02:21 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`+0.02`** :tada:
> Comparison is base [(`733d716`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/733d71681c76aa9011b8e828e7b4199af269d03b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.57% compared to head [(`fad0be8`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.60%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1248      +/-   ##
==========================================
+ Coverage   61.57%   61.60%   +0.02%     
==========================================
  Files         105      105              
  Lines        6512     6511       -1     
==========================================
+ Hits         4010     4011       +1     
+ Misses       1983     1981       -2     
  Partials      519      519              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.60% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `77.92% <ø> (-0.10%)` | :arrow_down: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1248?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on June 22, 2023 at 09:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1248*
