---
type: pull_request
number: 1310
title: "feat(RHINENG-15598): Add column management to Patch InventoryTable tables"
state: merged
author: leSamo
created: 2025-03-24T12:36:16Z
updated: 2025-04-13T22:18:33Z
closed: 2025-04-13T22:18:33Z
merged: 2025-04-13T22:18:33Z
base_branch: master
head_branch: inv-col-mgmt
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1310
---

# Pull Request #1310: feat(RHINENG-15598): Add column management to Patch InventoryTable tables

**Author**: @leSamo
**Created**: March 24, 2025 at 12:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `inv-col-mgmt`

## Description

https://issues.redhat.com/browse/RHINENG-15598

---

## Discussion

### Comment by @bastilian on March 24, 2025 at 12:38 PM UTC

@leSamo Would it not make more sense to implement the column management directly in the `InventoryTable` component, instead of implementing it individually in each application? 

### Comment by @codecov-commenter on March 24, 2025 at 12:43 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1310?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`361c657`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/361c65749ebef8e68e1157045a986f44da762cfe?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`61ca447`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/61ca447647ff82d8e223ec412af1655c803ce81a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1310   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1310/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1310/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1310/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1310?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @leSamo on March 24, 2025 at 02:03 PM UTC

@bastilian yes, that would be an option, but honestly that would only save a few lines of code here. I think that Inventory in its current state is crumbling under its own weight and we should be careful with adding new features. Future of Inventory is also quite uncertain (Inventory-centric views?).

### Comment by @bastilian on March 24, 2025 at 02:11 PM UTC

> but honestly that would only save a few lines of code here.

In each application. I think it'd be worth it. 


> Future of Inventory is also quite uncertain (Inventory-centric views?). 

We probably still have to support and maintain the `InventoryTable` for the foreseeable future.


We implemented every recent feature of the inventory table for each app in the InventoryTable itself. Plus we will need to implement the ColumnManagement in the InventoryTable anyways as well. 

### Comment by @leSamo on March 24, 2025 at 03:35 PM UTC

@bastilian Created a ticket for it: https://issues.redhat.com/browse/RHINENG-16902

### Comment by @johnsonm325 on March 24, 2025 at 08:02 PM UTC

Code looks good. If you remove all columns but one in the Advisories table, the expandable and selection columns get very wide. It does not do it on the Systems table. Maybe this is easy to fix?

![image](https://github.com/user-attachments/assets/978de4dc-5855-4204-95f1-2bf2d60a2679)


### Comment by @bastilian on March 25, 2025 at 09:41 AM UTC

> @bastilian Created a ticket for it: https://issues.redhat.com/browse/RHINENG-16902

That's great! We should now implement the column management in InventoryTable first and then  enable it for the apps using the InventoryTable, without having to implement the management in those apps themselves for their `SystemsTable`s

---

## Reviews

### Review by @johnsonm325 - Approved on March 24, 2025 at 08:02 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1310*
