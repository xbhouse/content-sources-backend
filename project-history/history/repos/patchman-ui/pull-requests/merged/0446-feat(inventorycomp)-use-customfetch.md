---
type: pull_request
number: 446
title: "feat(InventoryComp): use customFetch"
state: merged
author: mkholjuraev
created: 2021-02-22T12:55:21Z
updated: 2021-08-09T06:56:59Z
closed: 2021-06-01T20:26:36Z
merged: 2021-06-01T20:26:36Z
base_branch: master
head_branch: use-customFetch
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/446
---

# Pull Request #446: feat(InventoryComp): use customFetch

**Author**: @mkholjuraev
**Created**: February 22, 2021 at 12:55 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `use-customFetch`

## Description

Resolves: https://issues.redhat.com/browse/SPM-754

I am not sure if I have achieved everything that is wanted. As per slack conversation, we could get rid of InventoryEntitiesReducer. But when I removed the store, the Inventory table keeps loading with the error: 'page is not identified'

![image](https://user-images.githubusercontent.com/59481011/108910155-c1124680-7625-11eb-8f9e-e98a21df4d2c.png)


---

## Discussion

### Comment by @codecov-io on February 22, 2021 at 12:58 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=h1) Report
> Merging [#446](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=desc) (b2ec9cc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/18eaa9528dc74bfe7553b540ec4e90610f25c66e?el=desc) (18eaa95) will **decrease** coverage by `0.14%`.
> The diff coverage is `25.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #446      +/-   ##
==========================================
- Coverage   72.16%   72.01%   -0.15%     
==========================================
  Files          75       75              
  Lines        1304     1308       +4     
  Branches      173      174       +1     
==========================================
+ Hits          941      942       +1     
- Misses        315      317       +2     
- Partials       48       49       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `61.53% <25.00%> (-2.40%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=footer). Last update [18eaa95...b2ec9cc](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/446?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on June 01, 2021 at 08:36 PM UTC

:tada: This PR is included in version 1.20.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.20.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.20.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/446*
