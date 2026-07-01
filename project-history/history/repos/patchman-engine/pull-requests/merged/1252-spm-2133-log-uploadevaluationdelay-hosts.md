---
type: pull_request
number: 1252
title: "SPM-2133: log uploadEvaluationDelay hosts"
state: merged
author: MichaelMraka
created: 2023-06-28T15:00:52Z
updated: 2023-06-30T08:47:57Z
closed: 2023-06-30T08:47:56Z
merged: 2023-06-30T08:47:56Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1252
---

# Pull Request #1252: SPM-2133: log uploadEvaluationDelay hosts

**Author**: @MichaelMraka
**Created**: June 28, 2023 at 03:00 PM UTC
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

### Comment by @jira-linking on June 28, 2023 at 03:00 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2133


### Comment by @MichaelMraka on June 29, 2023 at 01:40 PM UTC

/retest

### Comment by @codecov-commenter on June 29, 2023 at 01:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`33.33`**% and project coverage change: **`+0.03`** :tada:
> Comparison is base [(`16f6f22`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/16f6f22d88ceadcd0073738de95496775ab79206?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.51% compared to head [(`0535a1b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.55%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1252      +/-   ##
==========================================
+ Coverage   61.51%   61.55%   +0.03%     
==========================================
  Files         105      105              
  Lines        6515     6518       +3     
==========================================
+ Hits         4008     4012       +4     
+ Misses       1988     1986       -2     
- Partials      519      520       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.55% <33.33%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.49% <33.33%> (-0.30%)` | :arrow_down: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1252?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on June 29, 2023 at 09:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1252*
