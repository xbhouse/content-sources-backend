---
type: pull_request
number: 1535
title: "RHINENG-14390: rewrite on_system_update trigger without locking"
state: merged
author: psegedy
created: 2024-11-22T16:25:02Z
updated: 2026-03-31T12:31:22Z
closed: 2024-12-04T13:03:20Z
merged: 2024-12-04T13:03:20Z
base_branch: master
head_branch: aad_insert_deadlock
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1535
---

# Pull Request #1535: RHINENG-14390: rewrite on_system_update trigger without locking

**Author**: @psegedy
**Created**: November 22, 2024 at 04:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `aad_insert_deadlock`

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

### Comment by @jira-linking on November 22, 2024 at 04:25 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-14390


### Comment by @codecov-commenter on November 26, 2024 at 01:39 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `8.69565%` with `21 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.31%. Comparing base ([`909e0ce`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/909e0cebf086fb6434f96a215803adcf0685272c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c6fa0a2`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c6fa0a2e006ebe12b390a2b38dade8eea19ef906?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1090 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [tasks/cleaning/clean\_advisory\_account\_data.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535?src=pr&el=tree&filepath=tasks%2Fcleaning%2Fclean_advisory_account_data.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2xlYW5pbmcvY2xlYW5fYWR2aXNvcnlfYWNjb3VudF9kYXRhLmdv) | 0.00% | [21 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1535      +/-   ##
==========================================
- Coverage   63.45%   63.31%   -0.14%     
==========================================
  Files         114      115       +1     
  Lines        9637     9659      +22     
==========================================
+ Hits         6115     6116       +1     
- Misses       2980     3001      +21     
  Partials      542      542              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.31% <8.69%> (-0.14%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1535?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @psegedy on November 27, 2024 at 01:27 PM UTC

/retest

### Comment by @psegedy on December 02, 2024 at 02:44 PM UTC

/retest

### Comment by @psegedy on December 02, 2024 at 03:18 PM UTC

/retest

### Comment by @psegedy on December 02, 2024 at 06:33 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on November 26, 2024 at 12:57 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1535*
