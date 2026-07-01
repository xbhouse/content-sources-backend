---
type: pull_request
number: 345
title: "feat(internationalization): install dependencies and adopt intl"
state: merged
author: mkholjuraev
created: 2020-12-01T09:53:16Z
updated: 2021-08-09T06:56:01Z
closed: 2020-12-15T10:09:15Z
merged: 2020-12-15T10:09:15Z
base_branch: master
head_branch: create-language-mappings
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/345
---

# Pull Request #345: feat(internationalization): install dependencies and adopt intl

**Author**: @mkholjuraev
**Created**: December 01, 2020 at 09:53 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `create-language-mappings`

## Description

resolves: https://issues.redhat.com/browse/SPM-604

---

## Discussion

### Comment by @codecov-io on December 01, 2020 at 09:56 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=h1) Report
> Merging [#345](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=desc) (474e931) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d23144241cde57bcd83488826143b3259c3dff93?el=desc) (d231442) will **decrease** coverage by `0.47%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #345      +/-   ##
==========================================
- Coverage   73.96%   73.49%   -0.48%     
==========================================
  Files          67       69       +2     
  Lines        1210     1211       +1     
  Branches      161      158       -3     
==========================================
- Hits          895      890       -5     
- Misses        263      268       +5     
- Partials       52       53       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/Utilities/IntlProvider.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9JbnRsUHJvdmlkZXIuanM=) | `0.00% <0.00%> (ø)` | |
| [src/entry-dev.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL2VudHJ5LWRldi5qcw==) | `0.00% <ø> (ø)` | |
| [src/entry.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL2VudHJ5Lmpz) | `0.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `76.59% <0.00%> (-4.77%)` | :arrow_down: |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `73.91% <0.00%> (-0.56%)` | :arrow_down: |
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <0.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| ... and [7 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=footer). Last update [e3dcbc9...474e931](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/345?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on December 15, 2020 at 10:17 AM UTC

:tada: This PR is included in version 1.2.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.2.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.2.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/345*
