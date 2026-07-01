---
type: pull_request
number: 1456
title: "RHINENG-8141: improve evaluation performance"
state: closed
author: Dugowitch
created: 2024-07-31T08:31:53Z
updated: 2024-08-07T15:48:41Z
closed: 2024-08-07T15:48:41Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1456
---

# Pull Request #1456: RHINENG-8141: improve evaluation performance

**Author**: @Dugowitch
**Created**: July 31, 2024 at 08:31 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `master`

## Description

Refactored in a way that all data is loaded and evaluated first. Then, all updates (and just the updates) are done in a single transaction. Furthermore, advisories are now processed like package have been.

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

### Comment by @jira-linking on July 31, 2024 at 08:31 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8141


### Comment by @codecov-commenter on July 31, 2024 at 08:37 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `81.78439%` with `49 lines` in your changes missing coverage. Please review.
> Project coverage is 65.09%. Comparing base [(`aba05cd`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/aba05cdea7a1897f9f4d71e2cb5afcac6696e4dc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`6c2f2bb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6c2f2bb3b3e2bfa0f29b48f3f4ed4e620c3e4993?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&filepath=evaluator%2Fevaluate_advisories.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 85.55% | [17 Missing and 9 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&filepath=evaluator%2Fevaluate.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 72.72% | [8 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&filepath=evaluator%2Fevaluate_packages.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 72.00% | [6 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1456      +/-   ##
==========================================
+ Coverage   64.83%   65.09%   +0.25%     
==========================================
  Files         114      114              
  Lines        7798     7810      +12     
==========================================
+ Hits         5056     5084      +28     
+ Misses       2216     2203      -13     
+ Partials      526      523       -3     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.09% <81.78%> (+0.25%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1456?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on August 07, 2024 at 11:07 AM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1456*
