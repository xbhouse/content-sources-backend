---
type: pull_request
number: 1028
title: "Refactor Template detail table to use Inventory table"
state: merged
author: leSamo
created: 2023-04-17T20:27:23Z
updated: 2023-05-08T18:14:45Z
closed: 2023-04-25T10:39:59Z
merged: 2023-04-25T10:39:59Z
base_branch: master
head_branch: template-detail-to-inventory
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1028
---

# Pull Request #1028: Refactor Template detail table to use Inventory table

**Author**: @leSamo
**Created**: April 17, 2023 at 08:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `template-detail-to-inventory`

## Description

https://issues.redhat.com/browse/SPM-1975

There should be almost no change in functionality, aside from added Tags column and sortable OS column in Template detail table.

---

## Discussion

### Comment by @codecov-commenter on April 18, 2023 at 08:13 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`11.29`**% and project coverage change: **`+0.17`** :tada:
> Comparison is base [(`a8070d7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a8070d7bc43da92bdc99429f86652677ab333021?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.12% compared to head [(`db1228e`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.29%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1028      +/-   ##
==========================================
+ Coverage   64.12%   64.29%   +0.17%     
==========================================
  Files         116      117       +1     
  Lines        2832     3014     +182     
  Branches      733      809      +76     
==========================================
+ Hits         1816     1938     +122     
- Misses       1016     1076      +60     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...tComponents/PatchSetDetail/PatchSetDetailAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbEFzc2V0cy5qcw==) | `25.00% <0.00%> (+25.00%)` | :arrow_up: |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `84.28% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `55.75% <0.00%> (-0.95%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <0.00%> (ø)` | |
| [src/Utilities/SystemsHelpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `21.73% <12.50%> (-4.93%)` | :arrow_down: |
| [src/store/Reducers/PatchSetDetailSystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0RGV0YWlsU3lzdGVtc1N0b3JlLmpz) | `42.85% <14.28%> (+7.56%)` | :arrow_up: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `55.35% <25.00%> (-1.01%)` | :arrow_down: |
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `76.47% <25.00%> (-6.87%)` | :arrow_down: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [6 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1028?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on April 19, 2023 at 10:59 PM UTC

@mkholjuraev thanks for review! I addressed all your comments. The bug with modal is not related to this PR and fixed it in #1029

---

## Reviews

### Review by @mkholjuraev - Changes Requested on April 19, 2023 at 10:49 AM UTC

Codewise looks good to me. A few comments I have faced an issue while searching for a system in the template review systems step that needs to be fixed.

<img width="1511" alt="image" src="https://user-images.githubusercontent.com/59481011/233133698-36ff2210-7cf6-46c8-96b1-c39f359d5189.png">


### Review by @leSamo - Commented on April 19, 2023 at 08:24 PM UTC

### Review by @leSamo - Commented on April 19, 2023 at 10:48 PM UTC

### Review by @leSamo - Commented on April 19, 2023 at 10:51 PM UTC

### Review by @leSamo - Commented on April 19, 2023 at 10:59 PM UTC

### Review by @AsToNlele - Approved on April 25, 2023 at 10:12 AM UTC

With the fix works great, good job! Only a small nitpick "Upto" should propably be "Up to"
![image](https://user-images.githubusercontent.com/20592948/234246130-9c242997-e427-4932-8f6f-28cab72ad042.png)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1028*
