---
type: pull_request
number: 816
title: "Code cleanup"
state: merged
author: mkholjuraev
created: 2022-06-08T09:32:36Z
updated: 2024-04-03T09:23:19Z
closed: 2022-06-28T14:36:25Z
merged: 2022-06-28T14:36:25Z
base_branch: master
head_branch: code-cleanup
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/816
---

# Pull Request #816: Code cleanup

**Author**: @mkholjuraev
**Created**: June 08, 2022 at 09:32 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `code-cleanup`

## Description

This PR is intended to do some cleanup, mainly to remove Redux store saved in local storage. This was used to handle saving persistant params, but it is abandoned already. 

Also, there are some minor code updates like renaming, removing obsolete lines...

Please make sure to test the application as a whole.

---

## Discussion

### Comment by @codecov-commenter on June 08, 2022 at 09:39 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#816](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4a38940) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/2cf577a6b5e74b9cabf3a0103439ccc06989437d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2cf577a) will **decrease** coverage by `0.02%`.
> The diff coverage is `50.00%`.

```diff
@@            Coverage Diff             @@
##           master     #816      +/-   ##
==========================================
- Coverage   70.95%   70.92%   -0.03%     
==========================================
  Files         101      101              
  Lines        2441     2418      -23     
  Branches      629      611      -18     
==========================================
- Hits         1732     1715      -17     
+ Misses        642      638       -4     
+ Partials       67       65       -2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [.../PatchSetWizard/InputFields/ConfigurationFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Db25maWd1cmF0aW9uRmllbGRzLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `25.92% <ø> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.69% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `68.35% <0.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `26.31% <ø> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `80.00% <100.00%> (+5.92%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2cf577a...4a38940](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/816?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 24, 2022 at 10:02 AM UTC

@RedHatInsights/team-interact it would be great if this PR gets review somehow. Tests are failing due to containerizartion, so please ignore them for now.

### Comment by @jiridostal on June 30, 2022 at 01:38 PM UTC

:tada: This PR is included in version 1.48.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on June 28, 2022 at 02:36 PM UTC

 I have tried reviewing carefully. Did not see any regressions in detailed local testing. Merging...

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/816*
