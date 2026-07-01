---
type: pull_request
number: 237
title: "chore(unit tests): for api.js and DataMappers.js`"
state: merged
author: mkholjuraev
created: 2020-08-17T09:45:49Z
updated: 2021-08-09T06:55:32Z
closed: 2020-08-18T08:42:40Z
merged: 2020-08-18T08:42:39Z
base_branch: master
head_branch: unit-test-utils
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/237
---

# Pull Request #237: chore(unit tests): for api.js and DataMappers.js`

**Author**: @mkholjuraev
**Created**: August 17, 2020 at 09:45 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `unit-test-utils`

## Description

- it seems that [fix(api) res.json() is not a fucntion error fix](https://github.com/RedHatInsights/patchman-ui/commit/277395626843a20cbed002b1d8011840926e2acb) was not a complete fix. I believe this PR has a complete fix.

- I was not able to implement some tests of functions in api.js file. They can not be tested as they call another function in the same file( [Problem definition](https://stackoverflow.com/questions/52650367/jestjs-how-to-test-function-being-called-inside-another-function) )

---

## Discussion

### Comment by @codecov-commenter on August 17, 2020 at 09:56 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=h1) Report
> Merging [#237](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/277395626843a20cbed002b1d8011840926e2acb&el=desc) will **increase** coverage by `12.88%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=tree)

```diff
@@             Coverage Diff             @@
##           master     #237       +/-   ##
===========================================
+ Coverage   36.82%   49.70%   +12.88%     
===========================================
  Files          51       52        +1     
  Lines         850      851        +1     
  Branches       95       95               
===========================================
+ Hits          313      423      +110     
+ Misses        483      392       -91     
+ Partials       54       36       -18     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `85.71% <100.00%> (+68.06%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `98.30% <0.00%> (+0.84%)` | :arrow_up: |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `3.07% <0.00%> (+3.07%)` | :arrow_up: |
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `12.50% <0.00%> (+12.50%)` | :arrow_up: |
| [...sentationalComponents/Filters/PublishDateFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1B1Ymxpc2hEYXRlRmlsdGVyLmpz) | `12.50% <0.00%> (+12.50%)` | :arrow_up: |
| [...ionalComponents/AdvisoriesTable/AdvisoriesTable.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yaWVzVGFibGUvQWR2aXNvcmllc1RhYmxlLmpz) | `13.33% <0.00%> (+13.33%)` | :arrow_up: |
| [src/PresentationalComponents/Filters/TypeFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1R5cGVGaWx0ZXIuanM=) | `14.28% <0.00%> (+14.28%)` | :arrow_up: |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `17.64% <0.00%> (+17.64%)` | :arrow_up: |
| ... and [14 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=footer). Last update [2773956...fb4eeb4](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/237?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on August 18, 2020 at 11:31 AM UTC

:tada: This PR is included in version 0.19.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on August 17, 2020 at 12:42 PM UTC

### Review by @jiridostal - Approved on August 18, 2020 at 08:42 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/237*
