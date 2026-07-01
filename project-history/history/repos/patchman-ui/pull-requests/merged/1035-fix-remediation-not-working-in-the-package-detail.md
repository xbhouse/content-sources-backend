---
type: pull_request
number: 1035
title: "Fix remediation not working in the package detail page"
state: merged
author: leSamo
created: 2023-04-28T20:16:17Z
updated: 2023-05-08T18:14:36Z
closed: 2023-05-03T14:29:16Z
merged: 2023-05-03T14:29:16Z
base_branch: master
head_branch: fix-package-remediations
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1035
---

# Pull Request #1035: Fix remediation not working in the package detail page

**Author**: @leSamo
**Created**: April 28, 2023 at 08:16 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-package-remediations`

## Description

https://issues.redhat.com/browse/SPM-2046

## Before:
- selecting systems one by one in the package detail page and then clicking remediate would crash the remediation wizard
- bulk selecting page/all worked fine

## After:
- selecting systems and remediating should work correctly

## To test:
- try to remediate some systems in the package detail page
  - try selecting systems one by one, bulk selecting, deselecting one by one and various combinations of these

---

## Discussion

### Comment by @codecov-commenter on April 28, 2023 at 08:21 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`50.00`**% and project coverage change: **`-0.43`** :warning:
> Comparison is base [(`af13391`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/af133912cd0bc0f85f837c4de5d41423bd2b69c4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.92% compared to head [(`d827155`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.49%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1035      +/-   ##
==========================================
- Coverage   63.92%   63.49%   -0.43%     
==========================================
  Files         118      118              
  Lines        2866     2909      +43     
  Branches      736      757      +21     
==========================================
+ Hits         1832     1847      +15     
- Misses       1034     1062      +28     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `64.44% <42.85%> (-12.03%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <100.00%> (ø)` | |

... and [4 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1035?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @AsToNlele on May 03, 2023 at 10:39 AM UTC

Hi, on one occasion when selecting 3 systems and then selecing a page I got this error:
![image](https://user-images.githubusercontent.com/20592948/235893455-d4c44db7-7fb8-4427-9b14-e3b8356ffc68.png)

Unfortunately I wasn't to able to get this error multiple times

---

## Reviews

### Review by @gkarat - Approved on May 03, 2023 at 11:55 AM UTC

Code-wise LGTM. Tried locally under stage-preview, it worked for me with different selection combinations. I also couldn't reproduce https://github.com/RedHatInsights/patchman-ui/pull/1035#issuecomment-1532803890 for some reason

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1035*
