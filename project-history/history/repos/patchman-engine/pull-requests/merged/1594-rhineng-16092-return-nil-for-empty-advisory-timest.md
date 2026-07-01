---
type: pull_request
number: 1594
title: "RHINENG-16092: return nil for empty advisory timestamp"
state: merged
author: psegedy
created: 2025-03-03T16:43:47Z
updated: 2025-03-04T12:32:37Z
closed: 2025-03-04T12:32:37Z
merged: 2025-03-04T12:32:37Z
base_branch: master
head_branch: time
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1594
---

# Pull Request #1594: RHINENG-16092: return nil for empty advisory timestamp

**Author**: @psegedy
**Created**: March 03, 2025 at 04:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `time`

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

### Comment by @jira-linking on March 03, 2025 at 04:43 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16092


### Comment by @codecov-commenter on March 03, 2025 at 04:55 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `66.66667%` with `1 line` in your changes missing coverage. Please review.
> Project coverage is 57.99%. Comparing base [(`8239911`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8239911c5c6fe007d739393394105c2108a53916?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`1689fa7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1689fa7214a6fc4340bfe8cb8e92e744c6149737?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/advisory\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fadvisory_detail.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | 0.00% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1594   +/-   ##
=======================================
  Coverage   57.99%   57.99%           
=======================================
  Files         143      143           
  Lines       11184    11184           
=======================================
  Hits         6486     6486           
  Misses       4115     4115           
  Partials      583      583           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.99% <66.66%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1594?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 04, 2025 at 09:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1594*
