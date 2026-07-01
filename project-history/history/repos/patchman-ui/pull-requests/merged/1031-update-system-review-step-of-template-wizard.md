---
type: pull_request
number: 1031
title: "Update System review step of Template wizard"
state: merged
author: leSamo
created: 2023-04-21T22:42:21Z
updated: 2023-05-08T18:14:42Z
closed: 2023-04-27T12:30:39Z
merged: 2023-04-27T12:30:39Z
base_branch: master
head_branch: update-system-review-step
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1031
---

# Pull Request #1031: Update System review step of Template wizard

**Author**: @leSamo
**Created**: April 21, 2023 at 10:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `update-system-review-step`

## Description

Mockups: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/DP7Ka9y

- Shorten the first sentence according to mockups
- Move the second sentence inside a warning alert, not exactly according to mockups, but full implementation is blocked, so this is a workaround
- Add "Last updated" column to the table according to mockups

## Before:
![Screenshot from 2023-04-22 00-39-19](https://user-images.githubusercontent.com/8426204/233744752-5c013d91-4fd3-4931-9634-8389e3ee89c3.png)

## After:
![Screenshot from 2023-04-22 00-38-47](https://user-images.githubusercontent.com/8426204/233744754-2c80386e-5c5d-4312-b729-e589d8b22a89.png)


---

## Discussion

### Comment by @codecov-commenter on April 21, 2023 at 10:48 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch and project coverage have no change.
> Comparison is base [(`af13391`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/af133912cd0bc0f85f837c4de5d41423bd2b69c4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.92% compared to head [(`6cae086`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.92%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1031   +/-   ##
=======================================
  Coverage   63.92%   63.92%           
=======================================
  Files         118      118           
  Lines        2866     2866           
  Branches      736      736           
=======================================
  Hits         1832     1832           
  Misses       1034     1034           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `43.47% <ø> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.65% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `57.29% <ø> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1031?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @AsToNlele - Approved on April 27, 2023 at 08:26 AM UTC

LGTM! :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1031*
