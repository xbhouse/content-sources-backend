---
type: pull_request
number: 331
title: "task(Inventroy) gate inventory calls before fetch"
state: merged
author: mkholjuraev
created: 2020-11-18T13:00:32Z
updated: 2021-08-09T06:55:51Z
closed: 2020-11-20T09:36:54Z
merged: 2020-11-20T09:36:53Z
base_branch: master
head_branch: gate-inventory
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/331
---

# Pull Request #331: task(Inventroy) gate inventory calls before fetch

**Author**: @mkholjuraev
**Created**: November 18, 2020 at 01:00 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `gate-inventory`

## Description

resolves: [SPM-608](https://issues.redhat.com/browse/SPM-608)
Not tested completely! but not breaking when RBAC result is mocked manually to false and true

---

## Discussion

### Comment by @codecov-io on November 19, 2020 at 09:10 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=h1) Report
> Merging [#331](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=desc) (5194b52) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9f23b7b78b546f2e6eef9af850e1bef0320ee79d?el=desc) (9f23b7b) will **decrease** coverage by `0.74%`.
> The diff coverage is `21.87%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #331      +/-   ##
==========================================
- Coverage   65.28%   64.54%   -0.75%     
==========================================
  Files          64       67       +3     
  Lines        1184     1210      +26     
  Branches      155      161       +6     
==========================================
+ Hits          773      781       +8     
- Misses        354      367      +13     
- Partials       57       62       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL1JvdXRlcy5qcw==) | `33.33% <0.00%> (-3.04%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `3.84% <0.00%> (-0.16%)` | :arrow_down: |
| [src/store/Reducers/SharedAppStateStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1NoYXJlZEFwcFN0YXRlU3RvcmUuanM=) | `40.00% <40.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `68.08% <100.00%> (+0.69%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `95.38% <0.00%> (ø)` | |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=footer). Last update [9f23b7b...5194b52](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/331?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on December 01, 2020 at 01:03 PM UTC

:tada: This PR is included in version 1.0.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.0.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.0.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/331*
