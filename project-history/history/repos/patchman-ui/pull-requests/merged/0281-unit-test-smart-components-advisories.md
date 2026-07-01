---
type: pull_request
number: 281
title: "Unit test smart components advisories"
state: merged
author: mkholjuraev
created: 2020-09-10T14:49:19Z
updated: 2021-08-09T06:55:20Z
closed: 2020-09-11T13:00:58Z
merged: 2020-09-11T13:00:58Z
base_branch: master
head_branch: unit-test-smart-components-advisories
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/281
---

# Pull Request #281: Unit test smart components advisories

**Author**: @mkholjuraev
**Created**: September 10, 2020 at 02:49 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `unit-test-smart-components-advisories`

## Description

resolves: [SPM-449](https://projects.engineering.redhat.com/browse/SPM-449)

---

## Discussion

### Comment by @codecov-commenter on September 11, 2020 at 08:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=h1) Report
> Merging [#281](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/ddb4b1604d80da054e9fe7d6a17f4f92410ba707?el=desc) will **increase** coverage by `16.41%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=tree)

```diff
@@             Coverage Diff             @@
##           master     #281       +/-   ##
===========================================
+ Coverage   50.05%   66.46%   +16.41%     
===========================================
  Files          56       57        +1     
  Lines         981      999       +18     
  Branches      118      118               
===========================================
+ Hits          491      664      +173     
+ Misses        441      296      -145     
+ Partials       49       39       -10     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `64.70% <64.70%> (ø)` | |
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `84.61% <0.00%> (+1.92%)` | :arrow_up: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `4.76% <0.00%> (+4.76%)` | :arrow_up: |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `70.96% <0.00%> (+9.67%)` | :arrow_up: |
| [src/store/Reducers/SystemPackageListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbVBhY2thZ2VMaXN0U3RvcmUuanM=) | `10.00% <0.00%> (+10.00%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `42.85% <0.00%> (+42.85%)` | :arrow_up: |
| [src/PresentationalComponents/Filters/TypeFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1R5cGVGaWx0ZXIuanM=) | `71.42% <0.00%> (+57.14%)` | :arrow_up: |
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `75.00% <0.00%> (+62.50%)` | :arrow_up: |
| ... and [4 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=footer). Last update [ddb4b16...e6d4afb](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/281?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 15, 2020 at 12:27 PM UTC

:tada: This PR is included in version 0.24.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/281*
