---
type: pull_request
number: 1575
title: "RHINENG-11956: lock the system"
state: merged
author: MichaelMraka
created: 2025-02-12T14:07:30Z
updated: 2026-03-31T21:35:32Z
closed: 2025-02-24T12:26:17Z
merged: 2025-02-24T12:26:17Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1575
---

# Pull Request #1575: RHINENG-11956: lock the system

**Author**: @MichaelMraka
**Created**: February 12, 2025 at 02:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

so other sessions won't delete it before insert

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

### Comment by @jira-linking on February 12, 2025 at 02:07 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11956


### Comment by @codecov-commenter on February 12, 2025 at 02:12 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1575?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.00%. Comparing base ([`4bf17e7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4bf17e7a8f11599d445eb3669f58a4823339eabc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e0d439f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e0d439f613bc769efabe933e9d165ef50b107844?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 994 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1575   +/-   ##
=======================================
  Coverage   58.00%   58.00%           
=======================================
  Files         143      143           
  Lines       11165    11166    +1     
=======================================
+ Hits         6476     6477    +1     
  Misses       4107     4107           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1575/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1575/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.00% <100.00%> (+<0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1575?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @MichaelMraka on February 19, 2025 at 10:01 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on February 20, 2025 at 09:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1575*
