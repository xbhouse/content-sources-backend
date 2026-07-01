---
type: pull_request
number: 1573
title: "RHINENG-15840: sort packages before updating system_package2"
state: merged
author: psegedy
created: 2025-02-11T15:56:45Z
updated: 2025-02-11T16:38:38Z
closed: 2025-02-11T16:38:38Z
merged: 2025-02-11T16:38:38Z
base_branch: master
head_branch: system_package2_deadlock
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1573
---

# Pull Request #1573: RHINENG-15840: sort packages before updating system_package2

**Author**: @psegedy
**Created**: February 11, 2025 at 03:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `system_package2_deadlock`

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

### Comment by @jira-linking on February 11, 2025 at 03:56 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-15840


### Comment by @codecov-commenter on February 11, 2025 at 04:01 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1573?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.06%. Comparing base [(`9c8f129`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9c8f129277c84aa934d9d7ab9b9a077615907d21?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5e6b5fa`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5e6b5faaa81e83663b5ebcf8012e4cc4ce23adbe?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1573      +/-   ##
==========================================
+ Coverage   58.03%   58.06%   +0.02%     
==========================================
  Files         143      143              
  Lines       11146    11153       +7     
==========================================
+ Hits         6469     6476       +7     
  Misses       4097     4097              
  Partials      580      580              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1573/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1573/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.06% <100.00%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1573?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 11, 2025 at 04:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1573*
