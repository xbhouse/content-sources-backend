---
type: pull_request
number: 959
title: "fix: Pagination fixes (loading skeletons, margins)"
state: merged
author: leSamo
created: 2023-02-14T23:58:11Z
updated: 2023-05-08T18:17:07Z
closed: 2023-02-15T23:22:54Z
merged: 2023-02-15T23:22:54Z
base_branch: master
head_branch: fix-paginations
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/959
---

# Pull Request #959: fix: Pagination fixes (loading skeletons, margins)

**Author**: @leSamo
**Created**: February 14, 2023 at 11:58 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-paginations`

## Description

# Description

- fix bottom pagination using double padding (32px instead of 16px)

### Before
![Screenshot from 2023-02-15 00-47-04](https://user-images.githubusercontent.com/8426204/218888644-db0cc2a7-b62d-4e52-a5af-630f7cd4546f.png)

### After
![Screenshot from 2023-02-15 00-46-55](https://user-images.githubusercontent.com/8426204/218888650-d50367e2-1292-4dbb-9eb2-dd1a980a6039.png)

- fix paginations not being replaced by loading skeletons when data is loading (this fix applies to all pages except Inventory tables, which will be addressed in a separate PR); top pagination displayed stale data for the duration of the loading, bottom pagination was removed for the duration of the loading

### Before
![Screenshot from 2023-02-15 00-55-45](https://user-images.githubusercontent.com/8426204/218890188-866b0986-7633-4ab8-be91-263303743e18.png)

### After
![Screenshot from 2023-02-15 00-55-22](https://user-images.githubusercontent.com/8426204/218890220-d66e38e4-6c9f-40f9-a9fa-b7be2535d033.png)

# How to test the PR
Bottom pagination should have correct 16px padding. Top and bottom paginations of all tables except Systems (Inventory) tables should be replaced by 200px loading skeleton when table data is loading.

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 15, 2023 at 12:05 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.35**% // Head: **72.41**% // Increases project coverage by **`+0.06%`** :tada:
> Coverage data is based on head [(`682ee4a`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`89bb482`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/89bb482e0410ddb3417b969d4b0379b08944535d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 73.33% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #959      +/-   ##
==========================================
+ Coverage   72.35%   72.41%   +0.06%     
==========================================
  Files         110      110              
  Lines        2615     2621       +6     
  Branches      658      662       +4     
==========================================
+ Hits         1892     1898       +6     
  Misses        723      723              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `85.07% <ø> (ø)` | |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `30.00% <33.33%> (+3.68%)` | :arrow_up: |
| [.../PresentationalComponents/TableView/TableFooter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVGb290ZXIuanM=) | `100.00% <100.00%> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/959?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on February 15, 2023 at 10:20 PM UTC

@mkholjuraev I remember UXD reporting tickets for these two issues when we had them in Vulnerability. No one complained about Patchman I just throught I might as well port the fixes here too to improve the UX a bit :grinning: 

### Comment by @mkholjuraev on February 15, 2023 at 11:39 PM UTC

:tada: This PR is included in version 1.59.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.59.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.59.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on February 15, 2023 at 09:47 PM UTC

interesting, is this a UX ask? should this be done across apps?

### Review by @mkholjuraev - Approved on February 15, 2023 at 10:25 PM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/959*
