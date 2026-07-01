---
type: pull_request
number: 279
title: "feat(packages): Add updatable column to package view"
state: merged
author: jiridostal
created: 2020-09-08T14:03:10Z
updated: 2022-08-02T08:40:42Z
closed: 2020-09-10T07:39:52Z
merged: 2020-09-10T07:39:52Z
base_branch: master
head_branch: package-updatable
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/279
---

# Pull Request #279: feat(packages): Add updatable column to package view

**Author**: @jiridostal
**Created**: September 08, 2020 at 02:03 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `package-updatable`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 08, 2020 at 02:26 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=h1) Report
> Merging [#279](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3c75d6b07a86e68c4cdd2cd07e1e69e31be4c42b?el=desc) will **increase** coverage by `1.03%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #279      +/-   ##
==========================================
+ Coverage   47.10%   48.14%   +1.03%     
==========================================
  Files          55       56       +1     
  Lines         968     1022      +54     
  Branches      115      127      +12     
==========================================
+ Hits          456      492      +36     
- Misses        459      472      +13     
- Partials       53       58       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/TableView/TableViewAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `75.00% <ø> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <0.00%> (ø)` | |
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <0.00%> (ø)` | |
| [src/store/Reducers/SystemPackageListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbVBhY2thZ2VMaXN0U3RvcmUuanM=) | `0.00% <0.00%> (ø)` | |
| [...c/PresentationalComponents/Filters/StatusFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1N0YXR1c0ZpbHRlci5qcw==) | `10.00% <0.00%> (ø)` | |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `2.17% <0.00%> (+0.47%)` | :arrow_up: |
| [.../PresentationalComponents/TableView/TableFooter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVGb290ZXIuanM=) | `100.00% <0.00%> (+33.33%)` | :arrow_up: |
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `100.00% <0.00%> (+50.00%)` | :arrow_up: |
| [.../PresentationalComponents/WithLoader/WithLoader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9XaXRoTG9hZGVyL1dpdGhMb2FkZXIuanM=) | `100.00% <0.00%> (+55.55%)` | :arrow_up: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=footer). Last update [3c75d6b...4fccf72](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/279?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 10, 2020 at 07:47 AM UTC

:tada: This PR is included in version 0.24.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on September 10, 2020 at 07:39 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/279*
