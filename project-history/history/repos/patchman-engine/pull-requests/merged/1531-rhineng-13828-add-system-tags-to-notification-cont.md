---
type: pull_request
number: 1531
title: "RHINENG-13828: add system tags to notification context"
state: merged
author: psegedy
created: 2024-11-11T17:31:06Z
updated: 2024-11-12T11:05:24Z
closed: 2024-11-12T11:05:24Z
merged: 2024-11-12T11:05:24Z
base_branch: master
head_branch: event_system_tag
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1531
---

# Pull Request #1531: RHINENG-13828: add system tags to notification context

**Author**: @psegedy
**Created**: November 11, 2024 at 05:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `event_system_tag`

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

### Comment by @jira-linking on November 11, 2024 at 05:31 PM UTC

Commits missing Jira IDs:
c6a7d6a4654c91f5c6ddb59f8c8e4a10e190e187
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-13828


### Comment by @codecov-commenter on November 11, 2024 at 05:36 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `57.14286%` with `12 lines` in your changes missing coverage. Please review.
> Project coverage is 63.46%. Comparing base [(`7375035`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7375035ee23b1d106e06e0f9c1f1de7fc68b7aaf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`c6a7d6a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c6a7d6a4654c91f5c6ddb59f8c8e4a10e190e187?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/notifications.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531?src=pr&el=tree&filepath=evaluator%2Fnotifications.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | 55.55% | [8 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1531      +/-   ##
==========================================
- Coverage   63.50%   63.46%   -0.04%     
==========================================
  Files         114      114              
  Lines        9609     9632      +23     
==========================================
+ Hits         6102     6113      +11     
- Misses       2970     2978       +8     
- Partials      537      541       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.46% <57.14%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1531?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on November 12, 2024 at 10:30 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1531*
