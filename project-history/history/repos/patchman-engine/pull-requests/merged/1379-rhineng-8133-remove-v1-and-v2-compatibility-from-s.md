---
type: pull_request
number: 1379
title: "RHINENG-8133: remove v1 and v2 compatibility from /systems"
state: merged
author: Dugowitch
created: 2024-02-27T19:13:38Z
updated: 2024-03-05T12:55:40Z
closed: 2024-03-05T12:18:24Z
merged: 2024-03-05T12:18:24Z
base_branch: master
head_branch: RHINENG-8133
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1379
---

# Pull Request #1379: RHINENG-8133: remove v1 and v2 compatibility from /systems

**Author**: @Dugowitch
**Created**: February 27, 2024 at 07:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-8133`

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

### Comment by @jira-linking on February 27, 2024 at 07:13 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8133


### Comment by @codecov-commenter on March 05, 2024 at 09:07 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `85.33333%` with `11 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.29%. Comparing base [(`90edf01`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/90edf01513b2831e467b135687870694d6247884?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`7d12ff1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 24 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | 73.68% | [6 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | 87.50% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1379      +/-   ##
==========================================
+ Coverage   62.43%   63.29%   +0.86%     
==========================================
  Files         108      107       -1     
  Lines        6772     6552     -220     
==========================================
- Hits         4228     4147      -81     
+ Misses       2016     1897     -119     
+ Partials      528      508      -20     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.29% <85.33%> (+0.86%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1379?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Changes Requested on February 28, 2024 at 12:54 PM UTC

### Review by @Dugowitch - Commented on March 05, 2024 at 09:05 AM UTC

### Review by @psegedy - Approved on March 05, 2024 at 11:47 AM UTC

LGTM! Great job

### Review by @psegedy - Commented on March 05, 2024 at 11:49 AM UTC

### Review by @psegedy - Commented on March 05, 2024 at 12:18 PM UTC

### Review by @Dugowitch - Commented on March 05, 2024 at 12:19 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1379*
