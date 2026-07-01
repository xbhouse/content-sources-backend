---
type: pull_request
number: 1587
title: "task: Add required FIPS dependency"
state: merged
author: SteveHNH
created: 2025-02-26T17:26:27Z
updated: 2026-04-03T06:23:03Z
closed: 2025-02-27T10:20:25Z
merged: 2025-02-27T10:20:25Z
base_branch: master
head_branch: add_additional_fips_dep
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1587
---

# Pull Request #1587: task: Add required FIPS dependency

**Author**: @SteveHNH
**Created**: February 26, 2025 at 05:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add_additional_fips_dep`

## Description

When running in RHEL9, the /usr/lib64/ossl-modules directory is also required to be copied into the micro image. Without it, FIPS errors occur during the database init phase of the startup.

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

### Comment by @jira-linking on February 26, 2025 at 05:26 PM UTC

Commits missing Jira IDs:
5dc0e76c9bc0983f0810f12edc54b299b372bf85


### Comment by @codecov-commenter on February 26, 2025 at 05:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1587?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 57.98%. Comparing base ([`a6a26d4`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a6a26d4c418d16c3f639cc162c9333cddbc80761?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`5dc0e76`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5dc0e76c9bc0983f0810f12edc54b299b372bf85?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 972 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1587   +/-   ##
=======================================
  Coverage   57.98%   57.98%           
=======================================
  Files         143      143           
  Lines       11172    11172           
=======================================
  Hits         6478     6478           
  Misses       4112     4112           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1587/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1587/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.98% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1587?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on February 27, 2025 at 10:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1587*
