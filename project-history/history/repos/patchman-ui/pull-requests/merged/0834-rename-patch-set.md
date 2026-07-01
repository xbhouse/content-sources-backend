---
type: pull_request
number: 834
title: "Rename patch set"
state: merged
author: mkholjuraev
created: 2022-06-29T19:59:05Z
updated: 2022-06-30T13:38:53Z
closed: 2022-06-30T13:20:08Z
merged: 2022-06-30T13:20:08Z
base_branch: master
head_branch: rename-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/834
---

# Pull Request #834: Rename patch set

**Author**: @mkholjuraev
**Created**: June 29, 2022 at 07:59 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rename-patch-set`

## Description

Renames all occurances of 'Patch set' on the UI to 'Patch template'. There are some places only 'Template' word is used. As the mockups are not updated in detail, I tried to cover all places and think about correct words, where to rename... 

Only wording on the UI should change without effecting functionality.

Please refer to the [mockups](https://marvelapp.com/prototype/hddg13b/screen/86218222) for more information.

---

## Discussion

### Comment by @codecov-commenter on June 30, 2022 at 11:33 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#834](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a6f13d1) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a6e400bfedd9b7ac35fc88992d6e717ccbd9d393?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a6e400b) will **decrease** coverage by `0.02%`.
> The diff coverage is `20.00%`.

> :exclamation: Current head a6f13d1 differs from pull request most recent head a990659. Consider uploading reports for the commit a990659 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #834      +/-   ##
==========================================
- Coverage   72.03%   72.00%   -0.03%     
==========================================
  Files         107      107              
  Lines        2471     2472       +1     
  Branches      630      630              
==========================================
  Hits         1780     1780              
- Misses        631      632       +1     
  Partials       60       60              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `66.66% <ø> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `0.00% <ø> (ø)` | |
| [...s/PatchSetWizard/InputFields/SelectExistingSets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9TZWxlY3RFeGlzdGluZ1NldHMuanM=) | `90.24% <ø> (ø)` | |
| [...mponents/PatchSetWizard/InputFields/ToDateField.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Ub0RhdGVGaWVsZC5qcw==) | `72.72% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `25.92% <0.00%> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.40% <0.00%> (-0.29%)` | :arrow_down: |
| [...Components/PatchSetWizard/steps/RequestProgress.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXF1ZXN0UHJvZ3Jlc3MuanM=) | `9.09% <0.00%> (ø)` | |
| [...tComponents/PatchSetWizard/steps/ReviewPatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdQYXRjaFNldC5qcw==) | `100.00% <ø> (ø)` | |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [a6e400b...a990659](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/834?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 30, 2022 at 01:20 PM UTC

@gabipodolnikova you are awesome, thanks for the review!

### Comment by @jiridostal on June 30, 2022 at 01:38 PM UTC

:tada: This PR is included in version 1.48.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Commented on June 30, 2022 at 11:09 AM UTC

Found one possible missing rename and also the navigation item Patch set should be renamed in the config repo right? :)

### Review by @mkholjuraev - Commented on June 30, 2022 at 11:19 AM UTC

### Review by @mkholjuraev - Commented on June 30, 2022 at 11:26 AM UTC

### Review by @gabipodolnikova - Approved on June 30, 2022 at 12:17 PM UTC

LGTM! Thanks! :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/834*
