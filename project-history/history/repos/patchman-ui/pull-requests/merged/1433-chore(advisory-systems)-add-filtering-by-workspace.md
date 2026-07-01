---
type: pull_request
number: 1433
title: "chore(advisory-systems): add filtering by workspace"
state: merged
author: TenSt
created: 2025-11-07T14:23:01Z
updated: 2025-11-10T14:25:58Z
closed: 2025-11-10T14:25:57Z
merged: 2025-11-10T14:25:57Z
base_branch: master
head_branch: stepan/RHINENG-5432-add-filtering-by-workspace-to-advisories-systems-table
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1433
---

# Pull Request #1433: chore(advisory-systems): add filtering by workspace

**Author**: @TenSt
**Created**: November 07, 2025 at 02:23 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/RHINENG-5432-add-filtering-by-workspace-to-advisories-systems-table`

## Description

# Description

Associated Jira ticket: [RHINENG-5432](https://issues.redhat.com/browse/RHINENG-5432)

This PR enables filtering by workspace for the systems list on the Advisory details page.


# How to test the PR

1. Make sure you have one or more workspaces
2. Go to Content -> Advisories
3. Click on any advisory
4. If the filter dropbox select "Workspace" and select some workspace from the list

# Before the change
"Workspace" option was not enabled in the filter dropbox.
<img width="1417" height="868" alt="Screenshot 2025-11-07 at 15 22 06" src="https://github.com/user-attachments/assets/354af5f5-2ec8-47e8-aad4-b8ddf070c8ff" />

# After the change
"Workspace" option is enabled in the filter dropbox.
<img width="1417" height="868" alt="Screenshot 2025-11-07 at 15 21 57" src="https://github.com/user-attachments/assets/ca78af69-da7a-4f5c-854d-9fadb967f281" />

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [x] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 07, 2025 at 02:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1433?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.36%. Comparing base ([`d1d2aae`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d1d2aaed620ec9d3d58790d82c30df1ac57736b6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b6f65ea`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b6f65eaf330876288641543bafb2ce699a91d22e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1433   +/-   ##
=======================================
  Coverage   62.36%   62.36%           
=======================================
  Files         128      128           
  Lines        3329     3329           
  Branches      894      894           
=======================================
  Hits         2076     2076           
  Misses       1132     1132           
  Partials      121      121           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1433?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Approved on November 09, 2025 at 08:46 PM UTC

Looks great! Thank you for getting this change in ✅ 

Just a thought: I've been looking at this, and it might be helpful for clarity if we explicitly pull `groups` out of the `rest` prop https://github.com/TenSt/patchman-ui/blob/9b2d117b402c78c3b5519621aa980034d7594e25/src/Utilities/DataMappers.js#L218. That way, it's immediately obvious to anyone that this table is also handling workspace data

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1433*
