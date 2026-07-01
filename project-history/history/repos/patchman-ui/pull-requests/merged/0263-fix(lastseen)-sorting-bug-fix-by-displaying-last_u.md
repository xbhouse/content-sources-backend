---
type: pull_request
number: 263
title: "fix(LastSeen): sorting bug fix by displaying last_upload instead of updated"
state: merged
author: mkholjuraev
created: 2020-08-31T20:44:08Z
updated: 2021-08-09T06:55:24Z
closed: 2020-09-03T11:07:10Z
merged: 2020-09-03T11:07:09Z
base_branch: master
head_branch: fix-last-seen
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/263
---

# Pull Request #263: fix(LastSeen): sorting bug fix by displaying last_upload instead of updated

**Author**: @mkholjuraev
**Created**: August 31, 2020 at 08:44 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-last-seen`

## Description

fixes: [SPM-396](https://projects.engineering.redhat.com/browse/SPM-396)

It looks a bit hackish. Improvements are appreciated ::)

---

## Discussion

### Comment by @codecov-commenter on August 31, 2020 at 08:54 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=h1) Report
> Merging [#263](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/8ad77ffd9b66ca4b5eedd355fb7176806440701a?el=desc) will **decrease** coverage by `0.58%`.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #263      +/-   ##
==========================================
- Coverage   47.93%   47.35%   -0.59%     
==========================================
  Files          55       55              
  Lines         945      963      +18     
  Branches      105      114       +9     
==========================================
+ Hits          453      456       +3     
- Misses        447      455       +8     
- Partials       45       52       +7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `86.66% <71.42%> (-13.34%)` | :arrow_down: |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `1.85% <0.00%> (-0.48%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `75.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=footer). Last update [8ad77ff...cb86d2c](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/263?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 03, 2020 at 11:14 AM UTC

:tada: This PR is included in version 0.22.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.22.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.22.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/263*
