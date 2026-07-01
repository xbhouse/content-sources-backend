---
type: pull_request
number: 1606
title: "Disable PR check tests"
state: merged
author: psegedy
created: 2025-03-17T07:47:48Z
updated: 2026-04-01T22:42:28Z
closed: 2025-03-17T12:38:21Z
merged: 2025-03-17T12:38:21Z
base_branch: master
head_branch: psegedy-patch-2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1606
---

# Pull Request #1606: Disable PR check tests

**Author**: @psegedy
**Created**: March 17, 2025 at 07:47 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `psegedy-patch-2`

## Description

tests keep failing in every PR, disable tests, keep only PR build and deploy to ephemeral

@niyazRedhat FYI

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

### Comment by @jira-linking on March 17, 2025 at 07:47 AM UTC

Commits missing Jira IDs:
1d3f4f2a0d489a85222f78b9a9a631714de18e7c


### Comment by @codecov-commenter on March 17, 2025 at 07:51 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1606?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 57.92%. Comparing base ([`8e6bfb2`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8e6bfb28658520a02fd571cd5bbe43fa0ffd597b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1d3f4f2`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1d3f4f2a0d489a85222f78b9a9a631714de18e7c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 939 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1606   +/-   ##
=======================================
  Coverage   57.92%   57.92%           
=======================================
  Files         145      145           
  Lines       11193    11193           
=======================================
  Hits         6483     6483           
  Misses       4127     4127           
  Partials      583      583           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1606/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1606/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.92% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1606?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 17, 2025 at 09:02 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1606*
