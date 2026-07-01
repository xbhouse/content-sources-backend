---
type: pull_request
number: 1492
title: "test(HMS-10052): extend patch pagination tests"
state: merged
author: TenSt
created: 2026-02-09T15:10:37Z
updated: 2026-02-12T23:03:44Z
closed: 2026-02-12T23:03:44Z
merged: 2026-02-12T23:03:44Z
base_branch: master
head_branch: stepan/HMS-10052-extend-patch-pagination-tests
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1492
---

# Pull Request #1492: test(HMS-10052): extend patch pagination tests

**Author**: @TenSt
**Created**: February 09, 2026 at 03:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/HMS-10052-extend-patch-pagination-tests`

## Description

# Description

Associated Jira ticket: #HMS-10052

This PR extends the patch pagination tests to include advisories and packages.
Plus refactoring of the systems pagination test.

# How to test/review the PR

- Pagination tests should be green.
- The main idea behind this PR is to test all pagination elements and check the state after each change. Both top and bottom control elements are used for interactions and change.
- The `advisories` and `packages` are similar except for the `targetRows` locator.
- The `systems` page uses different table (`inventory`) which is why the locators and checks are different - it is more static, so there is no need to combine the test with the new ones.
- The `commonNonInventoryLocators` function is a helper to not define the same locators multiple times.

---

## Discussion

### Comment by @codecov-commenter on February 09, 2026 at 03:12 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1492?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`d095c7d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d095c7dfb59c922b48c8751f9f2b686f66cec33e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b08e62f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b08e62fb06d3a18e5bb0197f40129df242e8f2ed?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1492   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      899    +7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1492?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @xbhouse - Commented on February 11, 2026 at 07:26 PM UTC

### Review by @xbhouse - Commented on February 11, 2026 at 07:28 PM UTC

### Review by @xbhouse - Approved on February 12, 2026 at 05:25 PM UTC

lgtm! nice work! just needs a rebase :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1492*
