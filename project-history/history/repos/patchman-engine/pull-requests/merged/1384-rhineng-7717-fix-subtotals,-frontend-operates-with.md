---
type: pull_request
number: 1384
title: "RHINENG-7717: fix subtotals, frontend operates with installable not a\u2026"
state: merged
author: psegedy
created: 2024-03-01T12:36:18Z
updated: 2024-03-04T11:31:38Z
closed: 2024-03-04T11:31:38Z
merged: 2024-03-04T11:31:38Z
base_branch: master
head_branch: subtotals
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1384
---

# Pull Request #1384: RHINENG-7717: fix subtotals, frontend operates with installable not a…

**Author**: @psegedy
**Created**: March 01, 2024 at 12:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `subtotals`

## Description

…pplicable packages

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

### Comment by @jira-linking on March 01, 2024 at 12:36 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-7717


### Comment by @codecov-commenter on March 01, 2024 at 12:42 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1384?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 62.50%. Comparing base [(`90edf01`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/90edf01513b2831e467b135687870694d6247884?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`3912a6c`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1384?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 16 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1384      +/-   ##
==========================================
+ Coverage   62.43%   62.50%   +0.07%     
==========================================
  Files         108      108              
  Lines        6772     6687      -85     
==========================================
- Hits         4228     4180      -48     
+ Misses       2016     1982      -34     
+ Partials      528      525       -3     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1384/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1384/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.50% <ø> (+0.07%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1384?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on March 01, 2024 at 02:21 PM UTC

test is failing because there is `limit=-1` used in tests which is not allowed

### Comment by @niyazRedhat on March 04, 2024 at 06:53 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on March 01, 2024 at 02:37 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1384*
