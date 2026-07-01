---
type: pull_request
number: 1608
title: "RHINENG-16781: refactor getMissingPackages"
state: merged
author: psegedy
created: 2025-03-17T14:40:27Z
updated: 2025-03-18T12:29:57Z
closed: 2025-03-18T12:29:56Z
merged: 2025-03-18T12:29:56Z
base_branch: master
head_branch: getmissingpackages
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1608
---

# Pull Request #1608: RHINENG-16781: refactor getMissingPackages

**Author**: @psegedy
**Created**: March 17, 2025 at 02:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `getmissingpackages`

## Description

bulk load packages missing in cache but present in DB

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

### Comment by @jira-linking on March 17, 2025 at 02:40 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16781


### Comment by @codecov-commenter on March 17, 2025 at 02:46 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `76.28866%` with `23 lines` in your changes missing coverage. Please review.
> Project coverage is 58.03%. Comparing base [(`726d1ee`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/726d1ee1e7744aeb4fcd407b5f8de35711026dfc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b43bbe8`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b43bbe88f4839020eb268e9b9e53a32334de81c4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/package\_cache.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608?src=pr&el=tree&filepath=evaluator%2Fpackage_cache.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | 73.86% | [18 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1608      +/-   ##
==========================================
+ Coverage   57.92%   58.03%   +0.11%     
==========================================
  Files         145      145              
  Lines       11193    11270      +77     
==========================================
+ Hits         6483     6541      +58     
- Misses       4127     4142      +15     
- Partials      583      587       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.03% <76.28%> (+0.11%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1608?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>🚀 New features to boost your workflow: </summary>

- ❄ [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 18, 2025 at 10:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1608*
