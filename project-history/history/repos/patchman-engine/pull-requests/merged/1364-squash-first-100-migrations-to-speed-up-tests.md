---
type: pull_request
number: 1364
title: "squash first 100 migrations to speed up tests"
state: merged
author: MichaelMraka
created: 2024-01-24T13:00:49Z
updated: 2026-04-03T21:21:02Z
closed: 2024-01-24T14:57:41Z
merged: 2024-01-24T14:57:41Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1364
---

# Pull Request #1364: squash first 100 migrations to speed up tests

**Author**: @MichaelMraka
**Created**: January 24, 2024 at 01:00 PM UTC
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

### Comment by @jira-linking on January 24, 2024 at 01:00 PM UTC

Commits missing Jira IDs:
3f86784a35e4c22cc19d7c0b18daaaebbee56331
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-6031


### Comment by @codecov-commenter on January 24, 2024 at 01:06 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1364?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 61.92%. Comparing base ([`9d3a4a7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9d3a4a78786ffaa604a02af518a6e8f8bc0117b2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`3f86784`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3f86784a35e4c22cc19d7c0b18daaaebbee56331?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1557 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1364   +/-   ##
=======================================
  Coverage   61.92%   61.92%           
=======================================
  Files         108      108           
  Lines        6815     6815           
=======================================
  Hits         4220     4220           
  Misses       2059     2059           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1364/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1364/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.92% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1364?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @psegedy - Commented on January 24, 2024 at 02:47 PM UTC

### Review by @psegedy - Approved on January 24, 2024 at 02:49 PM UTC

integration tests are working so I assume go-migrate can start from 100, lgtm

### Review by @MichaelMraka - Commented on January 24, 2024 at 02:54 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1364*
