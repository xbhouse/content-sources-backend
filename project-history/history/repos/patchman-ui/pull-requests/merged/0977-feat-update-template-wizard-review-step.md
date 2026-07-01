---
type: pull_request
number: 977
title: "feat: Update template wizard review step"
state: merged
author: leSamo
created: 2023-03-02T00:32:50Z
updated: 2023-05-08T18:17:01Z
closed: 2023-03-02T10:59:18Z
merged: 2023-03-02T10:59:18Z
base_branch: master
head_branch: update-wizard-review
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/977
---

# Pull Request #977: feat: Update template wizard review step

**Author**: @leSamo
**Created**: March 02, 2023 at 12:32 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-wizard-review`

## Description

# Description

Update template wizard review step (3rd step) according to the [mockups](https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/Qbnd821).

This change is WIP and will be split into multiple PRs.

## Before
![Screenshot from 2023-03-02 11-22-15](https://user-images.githubusercontent.com/8426204/222401501-1488e716-80d3-465a-9e02-11667e9d7064.png)

## After
![Screenshot from 2023-03-02 11-23-01](https://user-images.githubusercontent.com/8426204/222401489-23b537a1-b359-4d6f-9d93-d334dd62e769.png)

## Mockup
![Screenshot from 2023-03-02 11-22-31](https://user-images.githubusercontent.com/8426204/222401630-10a7a085-6070-46aa-8606-179f33bd6d1b.png)


# How to test the PR

Check if the code is acceptable. 

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on March 02, 2023 at 12:41 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`+0.01`** :tada:
> Comparison is base [(`c960cc2`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/c960cc22157c18c7ccd951a7af35a9d1aa700ba5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 70.61% compared to head [(`d8dfcb8`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 70.62%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #977      +/-   ##
==========================================
+ Coverage   70.61%   70.62%   +0.01%     
==========================================
  Files         113      113              
  Lines        2695     2696       +1     
  Branches      687      687              
==========================================
+ Hits         1903     1904       +1     
  Misses        792      792              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `29.03% <100.00%> (ø)` | |
| [...tComponents/PatchSetWizard/steps/ReviewPatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdQYXRjaFNldC5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `72.83% <100.00%> (+0.33%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/977?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Fewwy on March 02, 2023 at 10:25 AM UTC

The sketch shows that the "X systems"(the bottom right thingy) should be links. Currently, it's not
![image](https://user-images.githubusercontent.com/62722417/222401546-e79370bb-9c44-4aef-bf5c-e596d68625fc.png)


### Comment by @mkholjuraev on March 02, 2023 at 11:16 AM UTC

:tada: This PR is included in version 1.62.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/977*
