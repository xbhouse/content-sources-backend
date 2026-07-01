---
type: pull_request
number: 337
title: "feat(Errors) handle 403 error"
state: merged
author: mkholjuraev
created: 2020-11-20T11:40:19Z
updated: 2021-08-09T06:55:53Z
closed: 2020-12-01T09:33:10Z
merged: 2020-12-01T09:33:10Z
base_branch: master
head_branch: handle-403
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/337
---

# Pull Request #337: feat(Errors) handle 403 error

**Author**: @mkholjuraev
**Created**: November 20, 2020 at 11:40 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `handle-403`

## Description

Resolves: https://issues.redhat.com/browse/SPM-600
Tested with account without access

---

## Discussion

### Comment by @codecov-io on November 20, 2020 at 11:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=h1) Report
> Merging [#337](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=desc) (0671a96) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/7a4fa08487e948b0390a8fbd7838d470409a9410?el=desc) (7a4fa08) will **increase** coverage by `8.37%`.
> The diff coverage is `31.81%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #337      +/-   ##
==========================================
+ Coverage   65.22%   73.60%   +8.37%     
==========================================
  Files          64       68       +4     
  Lines        1182     1216      +34     
  Branches      154      161       +7     
==========================================
+ Hits          771      895     +124     
+ Misses        354      269      -85     
+ Partials       57       52       -5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/PresentationalComponents/Snippets/NoAccess.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9Ob0FjY2Vzcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1JvdXRlcy5qcw==) | `33.33% <0.00%> (-3.04%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `76.66% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `95.38% <ø> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `3.84% <0.00%> (-0.16%)` | :arrow_down: |
| [src/store/Reducers/SharedAppStateStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1NoYXJlZEFwcFN0YXRlU3RvcmUuanM=) | `40.00% <40.00%> (ø)` | |
| [...resentationalComponents/Snippets/AdvisoriesIcon.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9BZHZpc29yaWVzSWNvbi5qcw==) | `66.66% <66.66%> (ø)` | |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | `100.00% <100.00%> (ø)` | |
| ... and [14 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=footer). Last update [b73b170...0671a96](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/337?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on December 01, 2020 at 09:33 AM UTC

Looks good!


### Comment by @jiridostal on December 01, 2020 at 01:03 PM UTC

:tada: This PR is included in version 1.0.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.0.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.0.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/337*
