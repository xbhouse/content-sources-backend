---
type: pull_request
number: 1301
title: "SPM-2110: return 400 if user adds satellite system to baseline"
state: merged
author: yungbender
created: 2023-08-25T14:30:17Z
updated: 2023-08-25T16:08:30Z
closed: 2023-08-25T16:08:30Z
merged: 2023-08-25T16:08:30Z
base_branch: master
head_branch: satellite_baseline_api
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1301
---

# Pull Request #1301: SPM-2110: return 400 if user adds satellite system to baseline

**Author**: @yungbender
**Created**: August 25, 2023 at 02:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `satellite_baseline_api`

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

### Comment by @jira-linking on August 25, 2023 at 02:30 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2110


### Comment by @codecov-commenter on August 25, 2023 at 02:45 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`90.90%`** and project coverage change: **`+0.05%`** :tada:
> Comparison is base [(`99e4b5b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/99e4b5bb1172326afa34f946059a250ae0fe368e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.01% compared to head [(`166ea5e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.07%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1301      +/-   ##
==========================================
+ Coverage   62.01%   62.07%   +0.05%     
==========================================
  Files         106      106              
  Lines        6674     6684      +10     
==========================================
+ Hits         4139     4149      +10     
  Misses       2003     2003              
  Partials      532      532              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.07% <90.90%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `79.15% <ø> (ø)` | |
| [manager/controllers/baseline\_create.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `78.48% <87.50%> (+1.76%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `80.58% <100.00%> (+0.78%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1301?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 25, 2023 at 04:02 PM UTC

### Review by @psegedy - Commented on August 25, 2023 at 04:06 PM UTC

### Review by @psegedy - Commented on August 25, 2023 at 04:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1301*
