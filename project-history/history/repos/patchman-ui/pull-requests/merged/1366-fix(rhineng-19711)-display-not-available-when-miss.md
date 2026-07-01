---
type: pull_request
number: 1366
title: "fix(RHINENG-19711): display Not available when missing advisory date"
state: merged
author: darkeriossss
created: 2025-08-12T13:29:58Z
updated: 2025-08-14T08:09:00Z
closed: 2025-08-14T08:08:57Z
merged: 2025-08-14T08:08:57Z
base_branch: master
head_branch: unavailable-date
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1366
---

# Pull Request #1366: fix(RHINENG-19711): display Not available when missing advisory date

**Author**: @darkeriossss
**Created**: August 12, 2025 at 01:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `unavailable-date`

## Description

# Description

Associated Jira ticket: [RHINENG-19711](https://issues.redhat.com/browse/RHINENG-19711)

Display Not available instead of 1.1.1970 when missing date in an advisory


# How to test the PR

1. filter advisories by type other
2. date should show Not available instead of 1.1.1970

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description


---

## Discussion

### Comment by @codecov-commenter on August 12, 2025 at 01:32 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1366?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.54%. Comparing base ([`4dac963`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4dac963fe4487312242da83dce2ad0777291ade3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a996a93`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a996a933a8552592fb1d7e5544fad9efc56f8200?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1366      +/-   ##
==========================================
+ Coverage   62.53%   62.54%   +0.01%     
==========================================
  Files         126      126              
  Lines        3334     3335       +1     
  Branches      864      871       +7     
==========================================
+ Hits         2085     2086       +1     
  Misses       1128     1128              
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1366?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @darkeriossss on August 13, 2025 at 07:57 AM UTC

/retest

---

## Reviews

### Review by @bastilian - Approved on August 13, 2025 at 03:19 PM UTC

LGTM! Thank you @xmicha82!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1366*
