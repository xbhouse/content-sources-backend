---
type: pull_request
number: 1322
title: "feat(RHINENG-15545): Consume last seen column from Inventory"
state: merged
author: leSamo
created: 2025-04-25T12:09:09Z
updated: 2025-04-28T09:39:53Z
closed: 2025-04-28T09:39:53Z
merged: 2025-04-28T09:39:53Z
base_branch: master
head_branch: add-tooltip
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1322
---

# Pull Request #1322: feat(RHINENG-15545): Consume last seen column from Inventory

**Author**: @leSamo
**Created**: April 25, 2025 at 12:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-tooltip`

## Description



https://issues.redhat.com/browse/RHINENG-15545

The tooltip was added to the Inventory last seen column, but we were overwriting the column with our title due to intl, but it's not needed anymore. The column header will have tooltip only after RedHatInsights/insights-inventory-frontend#2429 has been merged.



---

## Discussion

### Comment by @codecov-commenter on April 25, 2025 at 12:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1322?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`b52f563`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b52f5630cd36ea1471c0cc6eeae9667a5bb89f19?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`21b5d3e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/21b5d3e179b6d2f69c3f71f40c69446b705e7442?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1322   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1322/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1322/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1322/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1322?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @bastilian - Approved on April 28, 2025 at 08:15 AM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1322*
