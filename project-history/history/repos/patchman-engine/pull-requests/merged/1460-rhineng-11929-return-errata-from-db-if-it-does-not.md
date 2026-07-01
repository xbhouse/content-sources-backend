---
type: pull_request
number: 1460
title: "RHINENG-11929: return errata from db if it does not have attributes"
state: merged
author: psegedy
created: 2024-08-13T10:15:50Z
updated: 2024-08-13T14:40:07Z
closed: 2024-08-13T14:40:07Z
merged: 2024-08-13T14:40:07Z
base_branch: master
head_branch: cached_errata
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1460
---

# Pull Request #1460: RHINENG-11929: return errata from db if it does not have attributes

**Author**: @psegedy
**Created**: August 13, 2024 at 10:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `cached_errata`

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

### Comment by @jira-linking on August 13, 2024 at 10:15 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11929


### Comment by @codecov-commenter on August 13, 2024 at 01:58 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `33.33333%` with `2 lines` in your changes missing coverage. Please review.
> Project coverage is 64.85%. Comparing base [(`1d94ffb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1d94ffb7aef6cd95480354e48470d928a49fbce2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`9009cc3`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9009cc3bcbea0f3208f7ecd1a372bc7c2c76b0d7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/advisory\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fadvisory_detail.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | 33.33% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1460      +/-   ##
==========================================
- Coverage   64.86%   64.85%   -0.02%     
==========================================
  Files         114      114              
  Lines        7798     7801       +3     
==========================================
+ Hits         5058     5059       +1     
- Misses       2214     2215       +1     
- Partials      526      527       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.85% <33.33%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1460?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1460*
