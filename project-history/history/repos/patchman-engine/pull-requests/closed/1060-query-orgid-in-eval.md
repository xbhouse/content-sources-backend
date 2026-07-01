---
type: pull_request
number: 1060
title: "Query orgID in eval"
state: closed
author: michalslomczynski
created: 2022-08-16T07:55:38Z
updated: 2022-08-17T09:17:39Z
closed: 2022-08-17T09:17:39Z
base_branch: master
head_branch: eval-orgID
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1060
---

# Pull Request #1060: Query orgID in eval

**Author**: @michalslomczynski
**Created**: August 16, 2022 at 07:55 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `eval-orgID`

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

### Comment by @codecov-commenter on August 16, 2022 at 08:03 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.41**% // Head: **61.39**% // Decreases project coverage by **`-0.02%`** :warning:
> Coverage data is based on head [(`00d4c44`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`4b2ead3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b2ead3ccb5a0dbdaf79148ff0e66187eb017e52?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 75.00% of modified lines in pull request are covered.

> :exclamation: Current head 00d4c44 differs from pull request most recent head 11a8b2e. Consider uploading reports for the commit 11a8b2e to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1060      +/-   ##
==========================================
- Coverage   61.41%   61.39%   -0.03%     
==========================================
  Files          98       98              
  Lines        5401     5408       +7     
==========================================
+ Hits         3317     3320       +3     
- Misses       1658     1661       +3     
- Partials      426      427       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.39% <75.00%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `73.77% <75.00%> (+3.40%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.07% <0.00%> (-1.31%)` | :arrow_down: |

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1060?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on August 17, 2022 at 09:17 AM UTC

Declined in favor of  #1061 

---

## Reviews

### Review by @psegedy - Changes Requested on August 16, 2022 at 08:38 AM UTC

### Review by @michalslomczynski - Commented on August 16, 2022 at 09:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1060*
