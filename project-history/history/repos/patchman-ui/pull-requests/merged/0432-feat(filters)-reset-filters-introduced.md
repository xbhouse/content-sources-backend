---
type: pull_request
number: 432
title: "feat(Filters): reset filters introduced"
state: merged
author: mkholjuraev
created: 2021-02-11T09:48:37Z
updated: 2021-08-09T06:56:14Z
closed: 2021-04-09T12:12:47Z
merged: 2021-04-09T12:12:47Z
base_branch: master
head_branch: changeto-reset-filters
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/432
---

# Pull Request #432: feat(Filters): reset filters introduced

**Author**: @mkholjuraev
**Created**: February 11, 2021 at 09:48 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `changeto-reset-filters`

## Description

resolves: https://issues.redhat.com/browse/SPM-716

---

## Discussion

### Comment by @codecov-io on February 11, 2021 at 09:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=h1) Report
> Merging [#432](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=desc) (9cc48e2) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/be6af2ae673826c19202159b467481d8ca61d7d2?el=desc) (be6af2a) will **decrease** coverage by `2.98%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #432      +/-   ##
==========================================
- Coverage   72.41%   69.42%   -2.99%     
==========================================
  Files          73       75       +2     
  Lines        1236     1318      +82     
  Branches      164      177      +13     
==========================================
+ Hits          895      915      +20     
- Misses        284      337      +53     
- Partials       57       66       +9     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <ø> (ø)` | |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <ø> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `98.70% <100.00%> (+0.01%)` | :arrow_up: |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |
| [src/store/Reducers/PackagesListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhY2thZ2VzTGlzdFN0b3JlLmpz) | `20.00% <100.00%> (ø)` | |
| [src/store/Reducers/SystemPackageListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbVBhY2thZ2VMaXN0U3RvcmUuanM=) | `16.66% <100.00%> (ø)` | |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `83.33% <0.00%> (-16.67%)` | :arrow_down: |
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `87.50% <0.00%> (-12.50%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `65.78% <0.00%> (-11.63%)` | :arrow_down: |
| ... and [13 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=footer). Last update [be6af2a...9cc48e2](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/432?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on April 07, 2021 at 12:14 PM UTC

https://github.com/RedHatInsights/frontend-components/pull/1035  
Looks like this can be completed now

### Comment by @jiridostal on April 09, 2021 at 12:21 PM UTC

:tada: This PR is included in version 1.17.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.17.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.17.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/432*
