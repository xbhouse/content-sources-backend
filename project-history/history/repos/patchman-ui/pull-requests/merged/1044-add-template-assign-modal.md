---
type: pull_request
number: 1044
title: "Add template assign modal"
state: merged
author: leSamo
created: 2023-05-05T01:12:39Z
updated: 2023-05-09T16:05:30Z
closed: 2023-05-05T06:59:35Z
merged: 2023-05-05T06:59:35Z
base_branch: master
head_branch: assign-modal
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1044
---

# Pull Request #1044: Add template assign modal

**Author**: @leSamo
**Created**: May 05, 2023 at 01:12 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `assign-modal`

## Description

Add a new modal when assigning systems to template from System list page

![Screenshot from 2023-05-05 03-11-24](https://user-images.githubusercontent.com/8426204/236359872-c519101c-6937-4bfb-811e-8778980e3c28.png)

Mockups: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/Om3zKYq#Inspect

So far it's not 100% according to the mockups, but it's good a large step in the right direction.




---

## Discussion

### Comment by @codecov-commenter on May 05, 2023 at 01:22 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`57.14`**% and project coverage change: **`-0.14`** :warning:
> Comparison is base [(`a1001c2`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a1001c222346b199054ccf818e2457657d525bf8?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.23% compared to head [(`c3e33f6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.10%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1044      +/-   ##
==========================================
- Coverage   63.23%   63.10%   -0.14%     
==========================================
  Files         118      119       +1     
  Lines        2927     2957      +30     
  Branches      757      759       +2     
==========================================
+ Hits         1851     1866      +15     
- Misses       1076     1091      +15     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...ionalComponents/PatchSetWrapper/PatchSetWrapper.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9QYXRjaFNldFdyYXBwZXIvUGF0Y2hTZXRXcmFwcGVyLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/Modals/useUnassignSystemsHook.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvdXNlVW5hc3NpZ25TeXN0ZW1zSG9vay5qcw==) | `100.00% <ø> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.69% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `83.09% <0.00%> (-1.19%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <0.00%> (ø)` | |
| [src/Utilities/constants.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Modals/AssignSystemsModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvQXNzaWduU3lzdGVtc01vZGFsLmpz) | `47.61% <47.61%> (ø)` | |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.85% <50.00%> (-0.21%)` | :arrow_down: |
| [src/Utilities/usePatchSetState.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VQYXRjaFNldFN0YXRlLmpz) | `57.14% <66.66%> (-1.20%)` | :arrow_down: |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `44.44% <75.00%> (+44.44%)` | :arrow_up: |
| ... and [2 more](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1044?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 09, 2023 at 04:05 PM UTC

:tada: This PR is included in version 1.63.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on May 05, 2023 at 06:48 AM UTC

LGTM, good job! :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1044*
