---
type: pull_request
number: 1595
title: "RHINENG-16331: update backoff to v2"
state: merged
author: psegedy
created: 2025-03-03T17:43:39Z
updated: 2026-04-07T02:51:06Z
closed: 2025-03-05T12:09:29Z
merged: 2025-03-05T12:09:29Z
base_branch: master
head_branch: updatedeps
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1595
---

# Pull Request #1595: RHINENG-16331: update backoff to v2

**Author**: @psegedy
**Created**: March 03, 2025 at 05:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `updatedeps`

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

### Comment by @jira-linking on March 03, 2025 at 05:43 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16331


### Comment by @niyazRedhat on March 04, 2025 at 12:26 PM UTC

/retest

### Comment by @codecov-commenter on March 04, 2025 at 01:05 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `84.21053%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 58.00%. Comparing base ([`8239911`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8239911c5c6fe007d739393394105c2108a53916?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0244d7c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0244d7c7025fae0a90dd88921705cc3d9f7a5d30?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 969 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/http.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595?src=pr&el=tree&filepath=base%2Futils%2Fhttp.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | 66.66% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1595      +/-   ##
==========================================
+ Coverage   57.99%   58.00%   +0.01%     
==========================================
  Files         143      143              
  Lines       11184    11180       -4     
==========================================
- Hits         6486     6485       -1     
+ Misses       4115     4113       -2     
+ Partials      583      582       -1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.00% <84.21%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1595?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 04, 2025 at 09:48 AM UTC

### Review by @MichaelMraka - Approved on March 05, 2025 at 09:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1595*
