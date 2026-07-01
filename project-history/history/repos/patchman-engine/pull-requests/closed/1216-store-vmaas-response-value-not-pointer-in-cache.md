---
type: pull_request
number: 1216
title: "store vmaas response value not pointer in cache"
state: closed
author: psegedy
created: 2023-04-26T12:11:24Z
updated: 2023-05-02T10:52:42Z
closed: 2023-05-02T10:52:42Z
base_branch: master
head_branch: vmaas_cache_fixes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1216
---

# Pull Request #1216: store vmaas response value not pointer in cache

**Author**: @psegedy
**Created**: April 26, 2023 at 12:11 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `vmaas_cache_fixes`

## Description

It looks like that storing pointers is preventing GC from releasing memory

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

### Comment by @jira-linking on April 26, 2023 at 12:11 PM UTC

Commits missing Jira IDs:
2dbc6f6773fa66b67302d1527d1d86c209d7c55b
8fef85fa011aa70953f8ec44765992c303196557


### Comment by @codecov-commenter on April 26, 2023 at 12:21 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`90.00`**% and project coverage change: **`+0.09`** :tada:
> Comparison is base [(`466f9e9`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/466f9e9177f0753208c841f318fc9de035f596cb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.02% compared to head [(`8fef85f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.12%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1216      +/-   ##
==========================================
+ Coverage   62.02%   62.12%   +0.09%     
==========================================
  Files         103      103              
  Lines        6286     6281       -5     
==========================================
+ Hits         3899     3902       +3     
+ Misses       1887     1879       -8     
  Partials      500      500              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.12% <90.00%> (+0.09%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `22.95% <ø> (+2.66%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.55% <75.00%> (ø)` | |
| [evaluator/vmaas\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `63.63% <100.00%> (+3.63%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1216?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1216*
