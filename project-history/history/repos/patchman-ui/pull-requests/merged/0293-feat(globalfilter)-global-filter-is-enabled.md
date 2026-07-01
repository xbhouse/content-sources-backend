---
type: pull_request
number: 293
title: "feat(GlobalFilter): global filter is enabled"
state: merged
author: mkholjuraev
created: 2020-09-23T11:47:28Z
updated: 2021-08-09T06:55:42Z
closed: 2020-10-01T08:27:32Z
merged: 2020-10-01T08:27:32Z
base_branch: master
head_branch: global-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/293
---

# Pull Request #293: feat(GlobalFilter): global filter is enabled

**Author**: @mkholjuraev
**Created**: September 23, 2020 at 11:47 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `global-filter`

## Description

Resolves: [SPM-478](https://issues.redhat.com/browse/SPM-478).

there is some refactoring of reducers. While working on global filter action on store, I realised that we can create a helper reducer and reuse it across our  reducers. But i am not sure if it is really beneficial.

---

## Discussion

### Comment by @codecov-commenter on September 23, 2020 at 11:58 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=h1) Report
> Merging [#293](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/81e5c5001622d98ad256b0e73fcd944d6fb85eee?el=desc) will **decrease** coverage by `0.97%`.
> The diff coverage is `68.29%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #293      +/-   ##
==========================================
- Coverage   67.57%   66.59%   -0.98%     
==========================================
  Files          56       57       +1     
  Lines         996      991       -5     
  Branches      119      126       +7     
==========================================
- Hits          673      660      -13     
- Misses        285      289       +4     
- Partials       38       42       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/SystemPackageListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbVBhY2thZ2VMaXN0U3RvcmUuanM=) | `16.66% <0.00%> (+6.66%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `95.20% <55.55%> (-3.11%)` | :arrow_down: |
| [src/store/Reducers/SystemsListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNMaXN0U3RvcmUuanM=) | `88.88% <80.00%> (-11.12%)` | :arrow_down: |
| [src/store/Reducers/AffectedSystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0FmZmVjdGVkU3lzdGVtc1N0b3JlLmpz) | `90.90% <83.33%> (-9.10%)` | :arrow_down: |
| [src/store/Reducers/AdvisoryListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5TGlzdFN0b3JlLmpz) | `90.90% <85.71%> (-9.10%)` | :arrow_down: |
| [src/store/Reducers/SystemAdvisoryListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbUFkdmlzb3J5TGlzdFN0b3JlLmpz) | `91.66% <85.71%> (-8.34%)` | :arrow_down: |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `71.87% <100.00%> (+0.90%)` | :arrow_up: |
| [src/store/Reducers/AdvisoryDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5RGV0YWlsU3RvcmUuanM=) | `100.00% <100.00%> (ø)` | |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=footer). Last update [81e5c50...72691d7](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/293?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on October 01, 2020 at 08:37 AM UTC

:tada: This PR is included in version 0.25.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.25.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.25.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/293*
