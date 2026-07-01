---
type: pull_request
number: 1434
title: "chore(package-systems): add filtering by workspace"
state: merged
author: TenSt
created: 2025-11-07T14:26:17Z
updated: 2025-11-11T09:54:15Z
closed: 2025-11-11T09:54:15Z
merged: 2025-11-11T09:54:15Z
base_branch: master
head_branch: stepan/RHINENG-5433-add-filtering-by-workspace-to-packages-systems-table
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1434
---

# Pull Request #1434: chore(package-systems): add filtering by workspace

**Author**: @TenSt
**Created**: November 07, 2025 at 02:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/RHINENG-5433-add-filtering-by-workspace-to-packages-systems-table`

## Description

Associated Jira ticket: [RHINENG-5433](https://issues.redhat.com/browse/RHINENG-5433)

This PR enables filtering by workspace for the systems list on the Package details page.


# How to test the PR

1. Make sure you have one or more workspaces
2. Go to Content -> Advisories
3. Click on any advisory
4. If the filter dropbox select "Workspace" and select some workspace from the list

# Before the change
"Workspace" option was not enabled in the filter dropbox.
<img width="1352" height="688" alt="Screenshot 2025-11-07 at 15 25 54" src="https://github.com/user-attachments/assets/a6f6db6f-f65b-4f26-8a05-5dedb442acd2" />

# After the change
"Workspace" option is enabled in the filter dropbox.
<img width="1352" height="688" alt="Screenshot 2025-11-07 at 15 25 39" src="https://github.com/user-attachments/assets/15326694-abd5-4ab0-83d9-2f1bf78dc696" />

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [x] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 07, 2025 at 02:28 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1434?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.35%. Comparing base ([`012e8fb`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/012e8fb301346433809fc4e61da4879d24c73544?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b8e3f0a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b8e3f0a14da4ace33e3dbc4918f21438a7e8a9cd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1434      +/-   ##
==========================================
- Coverage   62.36%   62.35%   -0.01%     
==========================================
  Files         128      128              
  Lines        3329     3323       -6     
  Branches      894      894              
==========================================
- Hits         2076     2072       -4     
+ Misses       1132     1130       -2     
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1434?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @katarinazaprazna on November 09, 2025 at 09:36 PM UTC

The "Workspaces" column was empty because the PR was adding the column header, not the corresponding row data. Adding the `groups` data here https://github.com/TenSt/patchman-ui/blob/6edaee74e371c8e973fc65aa6ca272dcc45a6f27/src/Utilities/DataMappers.js#L191 resolves the issue :)

```
export const createPackageSystemsRows = (rows, selectedRows = {}) => {
  const data =
    rows &&
    rows.map((row) => ({
      id: row.id,
      key: Math.random().toString() + row.id,
      display_name: row.display_name,
      installed_evra: row.installed_evra,
      available_evra: row.updatable ? row.available_evra : row.installed_evra,
      disableSelection: !row.updatable,
      updatable: row.updatable,
      update_status: row.update_status,
      selected: selectedRows[row.id] !== undefined,
      tags: row.tags,
      groups: row.groups, // <--- Workspaces
      os: {
        osName: row.os?.osName || row.os || 'N/A',
        rhsm: row.rhsm,
      },
      baseline_name: row.baseline_name, // ToBeDeprecated
      template_name: row.template_name,
      template_uuid: row.template_uuid,
      baseline_id: row.baseline_id,
      satellite_managed: row.satellite_managed,
    }));
  return data || [];
};
```

The previous PR displayed the table correctly as the table already had the row data included

### Comment by @TenSt on November 10, 2025 at 04:55 PM UTC

/retest

---

## Reviews

### Review by @katarinazaprazna - Approved on November 11, 2025 at 08:02 AM UTC

Looks great!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1434*
