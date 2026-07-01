---
type: pull_request
number: 1247
title: "use system_platform primary key for update"
state: merged
author: psegedy
created: 2023-06-20T09:00:24Z
updated: 2026-04-06T03:42:30Z
closed: 2023-06-20T10:50:05Z
merged: 2023-06-20T10:50:05Z
base_branch: master
head_branch: query
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1247
---

# Pull Request #1247: use system_platform primary key for update

**Author**: @psegedy
**Created**: June 20, 2023 at 09:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `query`

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

### Comment by @jira-linking on June 20, 2023 at 09:00 AM UTC

Commits missing Jira IDs:
2cec7f53b790f0795428677cd6b67ccb368ed392
4935da24c02338729f2381e59143bce7733f3177


### Comment by @codecov-commenter on June 20, 2023 at 09:42 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1247?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 61.57%. Comparing base ([`733d716`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/733d71681c76aa9011b8e828e7b4199af269d03b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`4935da2`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4935da24c02338729f2381e59143bce7733f3177?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1874 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1247      +/-   ##
==========================================
- Coverage   61.57%   61.57%   -0.01%     
==========================================
  Files         105      105              
  Lines        6512     6511       -1     
==========================================
- Hits         4010     4009       -1     
  Misses       1983     1983              
  Partials      519      519              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1247/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1247/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.57% <ø> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1247?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on June 20, 2023 at 09:33 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1247*
