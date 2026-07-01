---
type: pull_request
number: 1291
title: "ESSNTL-4817: fix rbac with multiple inventory and patch permissions"
state: merged
author: psegedy
created: 2023-08-14T13:31:49Z
updated: 2023-08-14T14:01:46Z
closed: 2023-08-14T14:01:46Z
merged: 2023-08-14T14:01:46Z
base_branch: master
head_branch: groups2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1291
---

# Pull Request #1291: ESSNTL-4817: fix rbac with multiple inventory and patch permissions

**Author**: @psegedy
**Created**: August 14, 2023 at 01:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `groups2`

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

### Comment by @jira-linking on August 14, 2023 at 01:31 PM UTC

Commits missing Jira IDs:
5cf54f33b9adddcf0d512d886af87af6161f13a7


### Comment by @codecov-commenter on August 14, 2023 at 01:39 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`93.75%`** and project coverage change: **`+0.13%`** :tada:
> Comparison is base [(`9d761af`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9d761af9274bbf9921c7d7079caced72f86092df?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.70% compared to head [(`5cf54f3`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.83%.
> Report is 13 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1291      +/-   ##
==========================================
+ Coverage   61.70%   61.83%   +0.13%     
==========================================
  Files         106      106              
  Lines        6651     6674      +23     
==========================================
+ Hits         4104     4127      +23     
  Misses       2013     2013              
  Partials      534      534              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.83% <93.75%> (+0.13%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.30% <85.71%> (+0.36%)` | :arrow_up: |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `80.26% <100.00%> (+2.32%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1291?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1291*
