---
type: pull_request
number: 1032
title: "Split first step of Template Wizard into two steps"
state: merged
author: leSamo
created: 2023-04-25T14:44:56Z
updated: 2023-05-08T18:14:43Z
closed: 2023-04-26T11:29:28Z
merged: 2023-04-26T11:29:28Z
base_branch: master
head_branch: split-template-wizard-steps
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1032
---

# Pull Request #1032: Split first step of Template Wizard into two steps

**Author**: @leSamo
**Created**: April 25, 2023 at 02:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `split-template-wizard-steps`

## Description

Part of: https://issues.redhat.com/browse/SPM-1976
Mockups: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/zxrAJZy#Inspect

---

## Discussion

### Comment by @codecov-commenter on April 25, 2023 at 02:56 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`38.46`**% and project coverage change: **`-1.66`** :warning:
> Comparison is base [(`2b1ed44`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/2b1ed4460f5a572a3782caf553cb0ede8b7ecea0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.22% compared to head [(`e1fc191`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1032      +/-   ##
==========================================
- Coverage   64.22%   62.57%   -1.66%     
==========================================
  Files         117      118       +1     
  Lines        2840     2979     +139     
  Branches      735      780      +45     
==========================================
+ Hits         1824     1864      +40     
- Misses       1016     1115      +99     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `5.76% <ø> (+1.11%)` | :arrow_up: |
| [...martComponents/PatchSetWizard/steps/ContentStep.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db250ZW50U3RlcC5qcw==) | `22.22% <22.22%> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.69% <50.00%> (+0.54%)` | :arrow_up: |
| [.../PatchSetWizard/InputFields/ConfigurationFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Db25maWd1cmF0aW9uRmllbGRzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `43.47% <100.00%> (+4.34%)` | :arrow_up: |

... and [11 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1032?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @AsToNlele - Approved on April 26, 2023 at 07:32 AM UTC

Looking good! Great job :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1032*
