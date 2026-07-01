---
type: pull_request
number: 1461
title: "RHINENG-8141: improve evaluation performance"
state: merged
author: Dugowitch
created: 2024-08-13T12:54:50Z
updated: 2024-09-26T10:10:38Z
closed: 2024-08-13T15:06:05Z
merged: 2024-08-13T15:06:05Z
base_branch: master
head_branch: RHINENG-8141
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1461
---

# Pull Request #1461: RHINENG-8141: improve evaluation performance

**Author**: @Dugowitch
**Created**: August 13, 2024 at 12:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-8141`

## Description

Refactored in a way that all data is loaded and evaluated first. Then, all updates (and just the updates) are done in a single transaction. Furthermore, advisories are now processed like package have been.

For code review see #1457.

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

### Comment by @jira-linking on August 13, 2024 at 12:54 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8141


### Comment by @codecov-commenter on August 13, 2024 at 02:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `80.38462%` with `51 lines` in your changes missing coverage. Please review.
> Project coverage is 64.97%. Comparing base [(`0a56574`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0a5657487846d0c412cbbf70960fef2de0f2ea33?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`e5890ec`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e5890ece41c30cbfb85a9908e928acbdd8cfab6e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&filepath=evaluator%2Fevaluate_advisories.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 84.61% | [18 Missing and 10 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&filepath=evaluator%2Fevaluate.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 72.72% | [8 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&filepath=evaluator%2Fevaluate_packages.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 72.00% | [6 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1461      +/-   ##
==========================================
+ Coverage   64.82%   64.97%   +0.15%     
==========================================
  Files         114      114              
  Lines        7801     7804       +3     
==========================================
+ Hits         5057     5071      +14     
+ Misses       2217     2208       -9     
+ Partials      527      525       -2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.97% <80.38%> (+0.15%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1461?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 13, 2024 at 02:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1461*
