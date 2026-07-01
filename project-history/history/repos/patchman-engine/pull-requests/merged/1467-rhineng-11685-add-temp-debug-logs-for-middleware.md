---
type: pull_request
number: 1467
title: "RHINENG-11685: add temp debug logs for middleware"
state: merged
author: Dugowitch
created: 2024-08-21T14:14:59Z
updated: 2024-08-22T09:45:13Z
closed: 2024-08-22T09:45:13Z
merged: 2024-08-22T09:45:13Z
base_branch: master
head_branch: RHINENG-11685
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1467
---

# Pull Request #1467: RHINENG-11685: add temp debug logs for middleware

**Author**: @Dugowitch
**Created**: August 21, 2024 at 02:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-11685`

## Description

This needs to get debugged in stage because local/ephemeral uses platform mock. When the mock is used, gzipping works as expected.

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

### Comment by @jira-linking on August 21, 2024 at 02:15 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11685


### Comment by @codecov-commenter on August 21, 2024 at 02:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `33.33333%` with `10 lines` in your changes missing coverage. Please review.
> Project coverage is 65.00%. Comparing base [(`6434537`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/64345375642c681ad797d6ffe0b2d9185e9ecfbd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`a9f399e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a9f399e81c331733c95cd5fbda8f6b2b11d9d3b4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/middlewares/limits.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Flimits.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9saW1pdHMuZ28=) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | 55.55% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/timeout.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Ftimeout.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy90aW1lb3V0Lmdv) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1467      +/-   ##
==========================================
- Coverage   65.06%   65.00%   -0.07%     
==========================================
  Files         114      114              
  Lines        7827     7836       +9     
==========================================
+ Hits         5093     5094       +1     
- Misses       2208     2215       +7     
- Partials      526      527       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.00% <33.33%> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1467?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Dugowitch on August 22, 2024 at 09:13 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on August 22, 2024 at 09:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1467*
