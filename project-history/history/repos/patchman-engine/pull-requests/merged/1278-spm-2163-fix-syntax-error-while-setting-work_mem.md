---
type: pull_request
number: 1278
title: "SPM-2163: fix syntax error while setting work_mem"
state: merged
author: psegedy
created: 2023-08-02T15:15:49Z
updated: 2026-04-06T04:30:18Z
closed: 2023-08-04T07:36:29Z
merged: 2023-08-04T07:36:29Z
base_branch: master
head_branch: work_mem_2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1278
---

# Pull Request #1278: SPM-2163: fix syntax error while setting work_mem

**Author**: @psegedy
**Created**: August 02, 2023 at 03:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `work_mem_2`

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

### Comment by @jira-linking on August 02, 2023 at 03:15 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2163


### Comment by @codecov-commenter on August 02, 2023 at 03:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1278?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 61.76%. Comparing base ([`1790976`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1790976fadceacb3c3e7e39dcb61ee0634adef43?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`dc8fa91`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dc8fa917f19371bc5fdbbe8f00eca9190bfbc8a6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1778 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1278      +/-   ##
==========================================
+ Coverage   60.54%   61.76%   +1.21%     
==========================================
  Files         106      106              
  Lines        6631     6648      +17     
==========================================
+ Hits         4015     4106      +91     
+ Misses       2085     2009      -76     
- Partials      531      533       +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1278/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1278/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.76% <100.00%> (+1.21%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1278?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on August 03, 2023 at 09:47 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1278*
