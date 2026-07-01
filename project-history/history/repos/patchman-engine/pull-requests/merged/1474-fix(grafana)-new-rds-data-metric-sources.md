---
type: pull_request
number: 1474
title: "fix(grafana): new rds data metric sources"
state: merged
author: vkrizan
created: 2024-09-05T12:29:49Z
updated: 2026-04-04T01:54:14Z
closed: 2024-09-09T08:07:36Z
merged: 2024-09-09T08:07:36Z
base_branch: master
head_branch: fix-grafana-rds-source
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1474
---

# Pull Request #1474: fix(grafana): new rds data metric sources

**Author**: @vkrizan
**Created**: September 05, 2024 at 12:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-grafana-rds-source`

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

### Comment by @jira-linking on September 05, 2024 at 12:29 PM UTC

Commits missing Jira IDs:
5a4723ad6cde49eadd0c9447443942c7a73d3b10
33e3b7f388f3546452bd9b3f0d7773c9ae6c5e4b


### Comment by @codecov-commenter on September 05, 2024 at 12:34 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1474?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 65.01%. Comparing base ([`ebaa313`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ebaa3139e0f79d8eba153b8a09416c2160e1838d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`33e3b7f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/33e3b7f388f3546452bd9b3f0d7773c9ae6c5e4b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1232 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1474   +/-   ##
=======================================
  Coverage   65.01%   65.01%           
=======================================
  Files         114      114           
  Lines        7823     7823           
=======================================
  Hits         5086     5086           
  Misses       2207     2207           
  Partials      530      530           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1474/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1474/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1474?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Changes Requested on September 06, 2024 at 07:47 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1474*
