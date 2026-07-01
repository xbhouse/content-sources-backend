---
type: pull_request
number: 1360
title: "RHINENG-6802: fix repo based reevalution timestamp"
state: merged
author: psegedy
created: 2024-01-05T11:29:45Z
updated: 2024-01-05T16:05:25Z
closed: 2024-01-05T16:05:24Z
merged: 2024-01-05T16:05:24Z
base_branch: master
head_branch: re-eval
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1360
---

# Pull Request #1360: RHINENG-6802: fix repo based reevalution timestamp

**Author**: @psegedy
**Created**: January 05, 2024 at 11:29 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `re-eval`

## Description

- last page of /repos response does not necessarily has the latest `latest_repo_change` timestamp
- add 1s to timestamp due to truncation of milliseconds in a format we use in DB 
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

### Comment by @jira-linking on January 05, 2024 at 11:29 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-6802


### Comment by @codecov-commenter on January 05, 2024 at 12:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `4 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`c98bc06`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c98bc06cea4542a6991d4b8534a11636e40fadd1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.09% compared to head [(`7891472`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.09%.
> Report is 3 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 60.00% | [2 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1360      +/-   ##
==========================================
- Coverage   62.09%   62.09%   -0.01%     
==========================================
  Files         108      108              
  Lines        6801     6811      +10     
==========================================
+ Hits         4223     4229       +6     
- Misses       2044     2046       +2     
- Partials      534      536       +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.09% <69.23%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1360?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 05, 2024 at 02:06 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1360*
