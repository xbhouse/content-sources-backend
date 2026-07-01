---
type: pull_request
number: 1289
title: "RHINENG-1536: rely on sorted availableUpdates from vmaas"
state: merged
author: psegedy
created: 2023-08-11T15:48:58Z
updated: 2023-08-14T09:23:44Z
closed: 2023-08-14T09:23:44Z
merged: 2023-08-14T09:23:44Z
base_branch: master
head_branch: fix_sort
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1289
---

# Pull Request #1289: RHINENG-1536: rely on sorted availableUpdates from vmaas

**Author**: @psegedy
**Created**: August 11, 2023 at 03:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_sort`

## Description

depends on https://github.com/RedHatInsights/vmaas-lib/pull/43
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

### Comment by @jira-linking on August 11, 2023 at 03:49 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1536


### Comment by @codecov-commenter on August 11, 2023 at 03:56 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00%`** and project coverage change: **`+0.04%`** :tada:
> Comparison is base [(`9d761af`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9d761af9274bbf9921c7d7079caced72f86092df?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.70% compared to head [(`53c0659`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.74%.
> Report is 6 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1289      +/-   ##
==========================================
+ Coverage   61.70%   61.74%   +0.04%     
==========================================
  Files         106      106              
  Lines        6651     6653       +2     
==========================================
+ Hits         4104     4108       +4     
+ Misses       2013     2011       -2     
  Partials      534      534              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.74% <100.00%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.07% <100.00%> (+0.12%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1289?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1289*
