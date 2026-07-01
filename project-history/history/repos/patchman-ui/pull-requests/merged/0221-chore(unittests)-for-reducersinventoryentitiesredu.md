---
type: pull_request
number: 221
title: "chore(UnitTests): for Reducers/InventoryEntitiesReducer + SystemDetai\u2026"
state: merged
author: mkholjuraev
created: 2020-08-10T21:06:06Z
updated: 2021-08-09T06:55:35Z
closed: 2020-08-17T12:40:36Z
merged: 2020-08-17T12:40:36Z
base_branch: master
head_branch: unit-tests-for-reducers
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/221
---

# Pull Request #221: chore(UnitTests): for Reducers/InventoryEntitiesReducer + SystemDetai…

**Author**: @mkholjuraev
**Created**: August 10, 2020 at 09:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `unit-tests-for-reducers`

## Description

[SPM-400](https://projects.engineering.redhat.com/browse/SPM-400)

---

## Discussion

### Comment by @codecov-commenter on August 10, 2020 at 09:16 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=h1) Report
> Merging [#221](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a04c4ccb883b8a05ae37ef227f2893df82d09179&el=desc) will **increase** coverage by `7.31%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #221      +/-   ##
==========================================
+ Coverage   36.91%   44.22%   +7.31%     
==========================================
  Files          51       51              
  Lines         848      848              
  Branches       95       95              
==========================================
+ Hits          313      375      +62     
+ Misses        481      427      -54     
+ Partials       54       46       -8     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `3.07% <0.00%> (+3.07%)` | :arrow_up: |
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `12.50% <0.00%> (+12.50%)` | :arrow_up: |
| [...sentationalComponents/Filters/PublishDateFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1B1Ymxpc2hEYXRlRmlsdGVyLmpz) | `12.50% <0.00%> (+12.50%)` | :arrow_up: |
| [...ionalComponents/AdvisoriesTable/AdvisoriesTable.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yaWVzVGFibGUvQWR2aXNvcmllc1RhYmxlLmpz) | `13.33% <0.00%> (+13.33%)` | :arrow_up: |
| [src/PresentationalComponents/Filters/TypeFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1R5cGVGaWx0ZXIuanM=) | `14.28% <0.00%> (+14.28%)` | :arrow_up: |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `17.64% <0.00%> (+17.64%)` | :arrow_up: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `21.42% <0.00%> (+21.42%)` | :arrow_up: |
| [...resentationalComponents/Snippets/SystemUpToDate.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9TeXN0ZW1VcFRvRGF0ZS5qcw==) | `33.33% <0.00%> (+33.33%)` | :arrow_up: |
| [.../PresentationalComponents/Snippets/NoSystemData.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9Ob1N5c3RlbURhdGEuanM=) | `50.00% <0.00%> (+50.00%)` | :arrow_up: |
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `50.00% <0.00%> (+50.00%)` | :arrow_up: |
| ... and [9 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=footer). Last update [a04c4cc...46f2767](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/221?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on August 18, 2020 at 11:31 AM UTC

:tada: This PR is included in version 0.19.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/221*
