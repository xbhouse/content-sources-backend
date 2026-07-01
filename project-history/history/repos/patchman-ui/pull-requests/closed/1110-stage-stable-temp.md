---
type: pull_request
number: 1110
title: "Stage stable temp"
state: closed
author: LiorKGOW
created: 2023-08-23T12:18:59Z
updated: 2023-08-24T17:09:34Z
closed: 2023-08-24T17:08:52Z
base_branch: master
head_branch: stage-stable-temp
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1110
---

# Pull Request #1110: Stage stable temp

**Author**: @LiorKGOW
**Created**: August 23, 2023 at 12:18 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `stage-stable-temp`

## Description

Try to cherry-pick commits :
* fix(SPM-1623): AdvisoriesPage - change in queryParams
* Update @patternfly/react-core
* chore(deps): bump @patternfly/react-table from 4.113.0 to 4.113.3

before cherry picking them to `stage-stable`

Just for practice
**Do not merge**

---

## Discussion

### Comment by @codecov-commenter on August 23, 2023 at 12:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`66.66%`** and project coverage change: **`+0.08%`** :tada:
> Comparison is base [(`20051ce`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/20051ce810c7e2b38b2ff1ad97342313a5187cab?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57% compared to head [(`23e86cf`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.66%.
> Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1110      +/-   ##
==========================================
+ Coverage   62.57%   62.66%   +0.08%     
==========================================
  Files         119      119              
  Lines        2998     2997       -1     
  Branches      770      767       -3     
==========================================
+ Hits         1876     1878       +2     
+ Misses       1122     1119       -3     
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `82.69% <0.00%> (-1.63%)` | :arrow_down: |
| [...mponents/PatchSetWizard/InputFields/ToDateField.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Ub0RhdGVGaWVsZC5qcw==) | `72.72% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `60.41% <ø> (+2.41%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `60.52% <100.00%> (+1.06%)` | :arrow_up: |
| [src/store/Reducers/AdvisoryListStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5TGlzdFN0b3JlLmpz) | `69.56% <100.00%> (+1.38%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1110?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LiorKGOW on August 24, 2023 at 05:08 PM UTC

Closing this branch 
It was just practice before performing the actual cherry-pick to stage-stable

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1110*
