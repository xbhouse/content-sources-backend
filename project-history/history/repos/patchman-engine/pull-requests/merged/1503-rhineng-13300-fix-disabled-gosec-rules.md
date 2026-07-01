---
type: pull_request
number: 1503
title: "RHINENG-13300: fix disabled gosec rules"
state: merged
author: MichaelMraka
created: 2024-10-07T10:53:59Z
updated: 2024-10-09T12:21:34Z
closed: 2024-10-09T12:21:34Z
merged: 2024-10-09T12:21:34Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1503
---

# Pull Request #1503: RHINENG-13300: fix disabled gosec rules

**Author**: @MichaelMraka
**Created**: October 07, 2024 at 10:53 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @jira-linking on October 07, 2024 at 10:54 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-13300


### Comment by @codecov-commenter on October 07, 2024 at 10:59 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `33.33333%` with `2 lines` in your changes missing coverage. Please review.
> Project coverage is 65.06%. Comparing base [(`6b519b4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6b519b4b92c4842952000ab7d66560c614a2873b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`a8ae9f2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a8ae9f2d4cf8d02c52de9cfcf975f282431a12e3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/http.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503?src=pr&el=tree&filepath=base%2Futils%2Fhttp.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1503      +/-   ##
==========================================
- Coverage   65.07%   65.06%   -0.01%     
==========================================
  Files         114      114              
  Lines        7880     7881       +1     
==========================================
  Hits         5128     5128              
- Misses       2215     2216       +1     
  Partials      537      537              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.06% <33.33%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1503?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on October 08, 2024 at 02:05 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1503*
