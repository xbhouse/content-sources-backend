---
type: pull_request
number: 1363
title: "RHINENG-6802: skip re-calc when latest_repo_change is nil"
state: merged
author: psegedy
created: 2024-01-19T09:38:22Z
updated: 2024-01-19T12:15:07Z
closed: 2024-01-19T12:15:06Z
merged: 2024-01-19T12:15:06Z
base_branch: master
head_branch: vmaas_sync_nil
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1363
---

# Pull Request #1363: RHINENG-6802: skip re-calc when latest_repo_change is nil

**Author**: @psegedy
**Created**: January 19, 2024 at 09:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas_sync_nil`

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

### Comment by @jira-linking on January 19, 2024 at 09:38 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-6802


### Comment by @codecov-commenter on January 19, 2024 at 09:44 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `17 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`405aa35`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/405aa3598abde5524412fb493f1801270e77293f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.91% compared to head [(`8c03201`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.89%.
> Report is 8 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [10 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 73.68% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [tasks/vmaas\_sync/repo\_based.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | 0.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1363      +/-   ##
==========================================
- Coverage   61.91%   61.89%   -0.03%     
==========================================
  Files         108      108              
  Lines        6804     6815      +11     
==========================================
+ Hits         4213     4218       +5     
- Misses       2056     2061       +5     
- Partials      535      536       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.89% <51.42%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1363?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 19, 2024 at 11:19 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1363*
