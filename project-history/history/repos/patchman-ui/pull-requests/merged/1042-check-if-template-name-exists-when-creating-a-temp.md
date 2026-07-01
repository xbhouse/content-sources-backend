---
type: pull_request
number: 1042
title: "Check if template name exists when creating a template"
state: merged
author: leSamo
created: 2023-05-02T18:12:53Z
updated: 2023-05-08T18:14:33Z
closed: 2023-05-03T12:27:42Z
merged: 2023-05-03T12:27:42Z
base_branch: master
head_branch: template-name-check
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1042
---

# Pull Request #1042: Check if template name exists when creating a template

**Author**: @leSamo
**Created**: May 02, 2023 at 06:12 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `template-name-check`

## Description

https://issues.redhat.com/browse/SPM-1952

Inside the template wizard:
- if user enters template name which is already used, they cannot continue to the next step
- if user enters template name which does not match regex `^[a-z0-9][a-z0-9_-]*$`, they cannot continue to the next step

---

## Discussion

### Comment by @codecov-commenter on May 02, 2023 at 06:23 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`43.18`**% and project coverage change: **`-2.00`** :warning:
> Comparison is base [(`1f8a079`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1f8a07987ce727aadf49d1ab39917fad2908d8e0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.89% compared to head [(`1cf564d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.90%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1042      +/-   ##
==========================================
- Coverage   63.89%   61.90%   -2.00%     
==========================================
  Files         118      118              
  Lines        2867     3000     +133     
  Branches      737      774      +37     
==========================================
+ Hits         1832     1857      +25     
- Misses       1035     1143     +108     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...mponents/PatchSetWizard/InputFields/ToDateField.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Ub0RhdGVGaWVsZC5qcw==) | `72.72% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.40% <0.00%> (-0.29%)` | :arrow_down: |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.69% <0.00%> (ø)` | |
| [...martComponents/PatchSetWizard/steps/ContentStep.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db250ZW50U3RlcC5qcw==) | `18.18% <0.00%> (-4.05%)` | :arrow_down: |
| [src/store/Reducers/SpecificPatchSetReducer.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1NwZWNpZmljUGF0Y2hTZXRSZWR1Y2VyLmpz) | `20.00% <0.00%> (-6.32%)` | :arrow_down: |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `35.48% <11.11%> (-8.00%)` | :arrow_down: |
| [src/store/Actions/Actions.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `71.42% <25.00%> (-1.98%)` | :arrow_down: |
| [...Components/PatchSetWizard/InputFields/NameField.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9OYW1lRmllbGQuanM=) | `83.33% <92.30%> (+10.60%)` | :arrow_up: |
| [.../PatchSetWizard/InputFields/ConfigurationFields.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Db25maWd1cmF0aW9uRmllbGRzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/ActionTypes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |

... and [3 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1042?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @gkarat - Approved on May 03, 2023 at 12:04 PM UTC

LGTM! Tried locally under stage-preview environment. Haven't spotted any code issues

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1042*
