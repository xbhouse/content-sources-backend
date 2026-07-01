---
type: pull_request
number: 1609
title: "fix(repack): give 600s timeout to acquire exclusive lock and skip repack if not succeeded"
state: merged
author: jdobes
created: 2025-03-17T15:45:34Z
updated: 2025-03-18T15:53:07Z
closed: 2025-03-18T15:53:07Z
merged: 2025-03-18T15:53:07Z
base_branch: master
head_branch: pg_repack_more_time
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1609
---

# Pull Request #1609: fix(repack): give 600s timeout to acquire exclusive lock and skip repack if not succeeded

**Author**: @jdobes
**Created**: March 17, 2025 at 03:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pg_repack_more_time`

## Description

rather do nothing than kill transactions

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

### Comment by @jira-linking on March 17, 2025 at 03:45 PM UTC

Commits missing Jira IDs:
2e7f19ce09b3ce7b445e07c6bfe4ae82330e56ea


### Comment by @codecov-commenter on March 17, 2025 at 03:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1609?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 57.87%. Comparing base [(`726d1ee`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/726d1ee1e7744aeb4fcd407b5f8de35711026dfc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`2e7f19c`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2e7f19ce09b3ce7b445e07c6bfe4ae82330e56ea?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1609      +/-   ##
==========================================
- Coverage   57.92%   57.87%   -0.05%     
==========================================
  Files         145      145              
  Lines       11193    11193              
==========================================
- Hits         6483     6478       -5     
- Misses       4127     4132       +5     
  Partials      583      583              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1609/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1609/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.87% <ø> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1609?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>🚀 New features to boost your workflow: </summary>

- ❄ [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @mtclinton on March 17, 2025 at 08:16 PM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on March 17, 2025 at 03:49 PM UTC

### Review by @MichaelMraka - Approved on March 18, 2025 at 10:06 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1609*
