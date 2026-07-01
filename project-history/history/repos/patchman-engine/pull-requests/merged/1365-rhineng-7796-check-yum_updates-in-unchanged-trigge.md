---
type: pull_request
number: 1365
title: "RHINENG-7796: check yum_updates in unchanged trigger"
state: merged
author: psegedy
created: 2024-01-25T12:38:04Z
updated: 2024-01-25T15:40:38Z
closed: 2024-01-25T15:40:38Z
merged: 2024-01-25T15:40:38Z
base_branch: master
head_branch: upload_yum
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1365
---

# Pull Request #1365: RHINENG-7796: check yum_updates in unchanged trigger

**Author**: @psegedy
**Created**: January 25, 2024 at 12:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `upload_yum`

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

### Comment by @jira-linking on January 25, 2024 at 12:38 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-7796


### Comment by @codecov-commenter on January 25, 2024 at 03:12 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1365?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Comparison is base [(`79e2763`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/79e2763cbf3c87eb58344005ffdc7782415d54e1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.87% compared to head [(`d1c683f`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1365?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.92%.
> Report is 1 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1365      +/-   ##
==========================================
+ Coverage   61.87%   61.92%   +0.04%     
==========================================
  Files         108      108              
  Lines        6815     6815              
==========================================
+ Hits         4217     4220       +3     
+ Misses       2062     2059       -3     
  Partials      536      536              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1365/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1365/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.92% <ø> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1365?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 25, 2024 at 03:39 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1365*
