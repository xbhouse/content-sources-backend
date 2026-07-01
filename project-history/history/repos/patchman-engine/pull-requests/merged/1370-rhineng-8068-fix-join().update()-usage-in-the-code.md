---
type: pull_request
number: 1370
title: "RHINENG-8068: fix Join().Update() usage in the code"
state: merged
author: Dugowitch
created: 2024-02-14T10:45:08Z
updated: 2026-04-04T04:32:16Z
closed: 2024-02-15T10:03:11Z
merged: 2024-02-15T10:03:11Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1370
---

# Pull Request #1370: RHINENG-8068: fix Join().Update() usage in the code

**Author**: @Dugowitch
**Created**: February 14, 2024 at 10:45 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `master`

## Description

There are 2 occurrences of the issue.
- first solved with adding EXISTS in `.Where()`
- second solved with removing `.Joins()`

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

### Comment by @jira-linking on February 14, 2024 at 10:45 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8068


### Comment by @codecov-commenter on February 14, 2024 at 10:50 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1370?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 61.97%. Comparing base ([`2789a8a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2789a8aff28ad50a9241589c8a56ce1c4f899bf1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`19e87a0`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/19e87a07e824831e87b033dd7006c5ce6420589b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1549 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1370      +/-   ##
==========================================
- Coverage   61.98%   61.97%   -0.01%     
==========================================
  Files         108      108              
  Lines        6821     6820       -1     
==========================================
- Hits         4228     4227       -1     
  Misses       2058     2058              
  Partials      535      535              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1370/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1370/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.97% <100.00%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1370?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on February 15, 2024 at 10:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1370*
