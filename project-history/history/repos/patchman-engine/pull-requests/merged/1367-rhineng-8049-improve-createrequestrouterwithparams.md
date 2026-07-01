---
type: pull_request
number: 1367
title: "RHINENG-8049: improve CreateRequestRouterWithParams func"
state: merged
author: Dugowitch
created: 2024-02-07T14:56:06Z
updated: 2024-02-13T14:25:28Z
closed: 2024-02-13T14:25:28Z
merged: 2024-02-13T14:25:28Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1367
---

# Pull Request #1367: RHINENG-8049: improve CreateRequestRouterWithParams func

**Author**: @Dugowitch
**Created**: February 07, 2024 at 02:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `master`

## Description

- unify `method` and `routerMethod`
- use `routerPath`, `param`, and `queryString` to crete url
- update all function calls
- update doTestWrongOffset to work with the new CreateRequestRouterWithParams func

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

### Comment by @jira-linking on February 07, 2024 at 02:56 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8049


### Comment by @codecov-commenter on February 13, 2024 at 09:50 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1367?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Comparison is base [(`2789a8a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2789a8aff28ad50a9241589c8a56ce1c4f899bf1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98% compared to head [(`ecc4408`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1367?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98%.
> Report is 3 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1367   +/-   ##
=======================================
  Coverage   61.98%   61.98%           
=======================================
  Files         108      108           
  Lines        6821     6821           
=======================================
  Hits         4228     4228           
  Misses       2058     2058           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1367/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1367/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.98% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1367?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on February 08, 2024 at 03:15 PM UTC

Now I even noticed we have the same pattern also in `CreateRequestRouterWithPath()` and `CreateRequestRouterWithAccount()` functions.
Could you also fix these?
Please do 3 commits - one for each function fix.

### Review by @Dugowitch - Commented on February 13, 2024 at 09:16 AM UTC

### Review by @MichaelMraka - Approved on February 13, 2024 at 02:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1367*
