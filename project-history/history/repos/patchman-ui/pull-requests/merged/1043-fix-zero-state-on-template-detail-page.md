---
type: pull_request
number: 1043
title: "Fix zero state on Template detail page"
state: merged
author: leSamo
created: 2023-05-03T14:20:52Z
updated: 2023-05-08T18:14:32Z
closed: 2023-05-04T08:39:55Z
merged: 2023-05-04T08:39:55Z
base_branch: master
head_branch: fix-zero-state-template-detail
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1043
---

# Pull Request #1043: Fix zero state on Template detail page

**Author**: @leSamo
**Created**: May 03, 2023 at 02:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-zero-state-template-detail`

## Description

Fixed a bug where Template detail page did not distinguish between template having no systems and template having no systems under currently applied global filter setting, which would incorrectly show this page
![Screenshot from 2023-05-03 16-18-59](https://user-images.githubusercontent.com/8426204/235943730-39d7552a-370a-4fa3-a52b-caa1e9bc4b60.png)

The issue is fixed by making an additional request for the systems to the backend, but this request always has no filters applied.

## What to test:
- If a template has no systems this component should be shown
![Screenshot from 2023-05-03 16-18-59](https://user-images.githubusercontent.com/8426204/235943730-39d7552a-370a-4fa3-a52b-caa1e9bc4b60.png)

- If a template has systems but currently applied global filter yields no results this component should be shown
![Screenshot from 2023-05-03 16-22-50](https://user-images.githubusercontent.com/8426204/235944842-5624f2a6-0245-49e5-be69-6cf3857f0308.png)

- Otherwise the table items should be shown
![Screenshot from 2023-05-03 16-22-45](https://user-images.githubusercontent.com/8426204/235944830-9cd1d8e9-84ac-4f61-88c1-639b26614742.png)




---

## Discussion

### Comment by @codecov-commenter on May 03, 2023 at 02:39 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`10.34`**% and project coverage change: **`+62.03`** :tada:
> Comparison is base [(`53b0acf`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/53b0acf9637e1c29afbea2c7372c5454bb2994f7?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 0.00% compared to head [(`999efa9`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.03%.

<details><summary>Additional details and impacted files</summary>


```diff
@@             Coverage Diff             @@
##           master    #1043       +/-   ##
===========================================
+ Coverage        0   62.03%   +62.03%     
===========================================
  Files           0      118      +118     
  Lines           0     3034     +3034     
  Branches        0      786      +786     
===========================================
+ Hits            0     1882     +1882     
- Misses          0     1152     +1152     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/PatchSetDetailSystemsStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0RGV0YWlsU3lzdGVtc1N0b3JlLmpz) | `30.00% <0.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `68.37% <28.57%> (ø)` | |
| [src/store/ActionTypes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |

... and [114 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1043?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @AsToNlele - Approved on May 04, 2023 at 07:17 AM UTC

Works great! Good job :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1043*
