---
type: pull_request
number: 1354
title: "RHINENG-5394: run vacuum in goroutine so api won't time out"
state: merged
author: psegedy
created: 2023-12-11T11:00:10Z
updated: 2023-12-11T12:34:45Z
closed: 2023-12-11T12:34:45Z
merged: 2023-12-11T12:34:44Z
base_branch: master
head_branch: migration4
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1354
---

# Pull Request #1354: RHINENG-5394: run vacuum in goroutine so api won't time out

**Author**: @psegedy
**Created**: December 11, 2023 at 11:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `migration4`

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

### Comment by @jira-linking on December 11, 2023 at 11:00 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-5394


### Comment by @codecov-commenter on December 11, 2023 at 12:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1354?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Comparison is base [(`997b6bc`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/997b6bc3100a1312134c476be5f9309a68537f30?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.00% compared to head [(`a4294a5`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1354?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.03%.
> Report is 4 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1354      +/-   ##
==========================================
+ Coverage   62.00%   62.03%   +0.02%     
==========================================
  Files         108      108              
  Lines        6811     6811              
==========================================
+ Hits         4223     4225       +2     
+ Misses       2054     2052       -2     
  Partials      534      534              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1354/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1354/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.03% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1354?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on December 11, 2023 at 12:16 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1354*
