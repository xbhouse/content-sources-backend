---
type: pull_request
number: 1612
title: "RHINENG-16833: disable semantic commits by renovate"
state: merged
author: psegedy
created: 2025-03-25T14:49:46Z
updated: 2025-03-26T09:59:02Z
closed: 2025-03-26T09:59:01Z
merged: 2025-03-26T09:59:01Z
base_branch: master
head_branch: renovatesemver
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1612
---

# Pull Request #1612: RHINENG-16833: disable semantic commits by renovate

**Author**: @psegedy
**Created**: March 25, 2025 at 02:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `renovatesemver`

## Description

default setting is 'auto' which causes that renovate decides to use semantic commits based on last commits. It happens that closed PRs updating major versions are re-opened due to commits having different format

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

### Comment by @jira-linking on March 25, 2025 at 02:49 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16833


### Comment by @codecov-commenter on March 25, 2025 at 02:54 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1612?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.08%. Comparing base [(`0f26381`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0f263817f5f345fb26a13ddc36a1d12c3d3cf868?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`cab2ceb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cab2cebe91b9ccf857de83dc194f73181ca8ab0c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1612   +/-   ##
=======================================
  Coverage   58.08%   58.08%           
=======================================
  Files         145      145           
  Lines       11271    11271           
=======================================
  Hits         6547     6547           
  Misses       4139     4139           
  Partials      585      585           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1612/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1612/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.08% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1612?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 25, 2025 at 03:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1612*
