---
type: pull_request
number: 1429
title: "RHINENG-21469: bulk select only items on page as default"
state: merged
author: TenSt
created: 2025-11-03T15:04:17Z
updated: 2025-11-04T12:16:09Z
closed: 2025-11-04T12:16:09Z
merged: 2025-11-04T12:16:09Z
base_branch: master
head_branch: stepan/RHINENG-21469-select-items-on-page-as-default-instead-of-all
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1429
---

# Pull Request #1429: RHINENG-21469: bulk select only items on page as default

**Author**: @TenSt
**Created**: November 03, 2025 at 03:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/RHINENG-21469-select-items-on-page-as-default-instead-of-all`

## Description

# Description

Associated Jira ticket: # ([RHINENG-21469](https://issues.redhat.com/browse/RHINENG-21469))

This PR changes the default behavior of the bulk select from "all" to "page" in the "useBulkSelectConfig" hook that used in different components.

# How to test the PR

1. Go to Content -> Advisories
2. Click on checkbox to select all items

# Before the change
All the items are selected on all pages
<img width="1661" height="669" alt="Screenshot 2025-11-03 at 15 59 28" src="https://github.com/user-attachments/assets/cebf9ad3-3904-41df-97fb-4ed796c7f4e4" />

# After the change
Only items visible on the page are selected.
<img width="1661" height="656" alt="Screenshot 2025-11-03 at 15 57 21" src="https://github.com/user-attachments/assets/32fd3204-a9f6-4f46-9926-06f473d14ae5" />

---

## Discussion

### Comment by @codecov-commenter on November 03, 2025 at 03:56 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1429?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.66%. Comparing base ([`1c9c2da`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1c9c2dab0c9e86c9dea1d3ebf8d18aafb6b74cdc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a1dc47a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a1dc47a78bfbd55881063449c0c6b72c3403a6a1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1429      +/-   ##
==========================================
- Coverage   62.67%   62.66%   -0.02%     
==========================================
  Files         126      126              
  Lines        3306     3305       -1     
  Branches      857      857              
==========================================
- Hits         2072     2071       -1     
  Misses       1113     1113              
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1429?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Approved on November 04, 2025 at 12:14 PM UTC

Thank you for the change!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1429*
