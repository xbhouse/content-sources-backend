---
type: pull_request
number: 1223
title: "SPM-2064: re-evaluate systems removed from baseline"
state: merged
author: psegedy
created: 2023-05-23T16:03:17Z
updated: 2023-05-24T08:28:59Z
closed: 2023-05-24T08:28:59Z
merged: 2023-05-24T08:28:59Z
base_branch: master
head_branch: reevaluate
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1223
---

# Pull Request #1223: SPM-2064: re-evaluate systems removed from baseline

**Author**: @psegedy
**Created**: May 23, 2023 at 04:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `reevaluate`

## Description

systems need to be re-evaluated, otherwise we are showing incorrect Advisory installability/applicability when system is removed from template

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

### Comment by @jira-linking on May 23, 2023 at 04:03 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2064


### Comment by @codecov-commenter on May 23, 2023 at 04:11 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`+0.15`** :tada:
> Comparison is base [(`f0aa0d1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f0aa0d1d16ae0f4b9bb892a0ab2844c3ae87847b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98% compared to head [(`a2bdbef`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.14%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1223      +/-   ##
==========================================
+ Coverage   61.98%   62.14%   +0.15%     
==========================================
  Files         103      103              
  Lines        6303     6321      +18     
==========================================
+ Hits         3907     3928      +21     
+ Misses       1896     1894       -2     
+ Partials      500      499       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.14% <100.00%> (+0.15%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.63% <100.00%> (+0.08%)` | :arrow_up: |
| [manager/controllers/baseline\_systems\_remove.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX3JlbW92ZS5nbw==) | `75.00% <100.00%> (+1.47%)` | :arrow_up: |
| [manager/controllers/systems\_advisories\_view.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `80.83% <100.00%> (+5.59%)` | :arrow_up: |
| [manager/kafka/kafka.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1223?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1223*
