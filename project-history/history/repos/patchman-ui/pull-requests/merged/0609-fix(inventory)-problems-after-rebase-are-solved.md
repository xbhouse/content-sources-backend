---
type: pull_request
number: 609
title: "fix(Inventory): problems after rebase are solved"
state: merged
author: mkholjuraev
created: 2021-07-28T16:29:10Z
updated: 2022-05-17T08:50:27Z
closed: 2021-07-29T07:33:05Z
merged: 2021-07-29T07:33:05Z
base_branch: master
head_branch: selection
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/609
---

# Pull Request #609: fix(Inventory): problems after rebase are solved

**Author**: @mkholjuraev
**Created**: July 28, 2021 at 04:29 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `selection`

## Description

The issues after rebase:
1. `filterSelectedRowIDs` function in Helpers.js was renamed to `removeUndefinedObjectKeys`. But this change was got lost during rebase and it was not reflected in Inventory Components.
2. `useEffect` in PackageSystems page was changed to async function and it broke the implementation of `clearInventoryReducer`. Thus if you select a system on the Package Systems page it was present also in Systems Page. 
3. If you select All systems, all of the systems on the first page is selected even if it was selected. It was because we did not call the `inventoryModifier` reducer to remap with updated systems selected.
4. The last problem that occurred was the incorrect `useSelector` for selected rows in all Inventory pages. Previously `selectedRows` was a part of the relevant reducer for the component. But after rebase, it is now stored in `entities` reducer

---

## Discussion

### Comment by @codecov-commenter on July 28, 2021 at 04:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#609](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d0499c6) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/7c10867bb1baffea79e4d8ebd6f6f6fa0253b240?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7c10867) will **decrease** coverage by `0.05%`.
> The diff coverage is `8.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #609      +/-   ##
==========================================
- Coverage   52.87%   52.81%   -0.06%     
==========================================
  Files          79       79              
  Lines        1876     1880       +4     
  Branches      402      402              
==========================================
+ Hits          992      993       +1     
- Misses        804      807       +3     
  Partials       80       80              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `40.00% <0.00%> (-1.67%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.31% <100.00%> (+0.10%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [7c10867...d0499c6](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/609?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on July 29, 2021 at 07:33 AM UTC

did not try the change, but let's merge it and observe in CI.

### Comment by @jiridostal on July 29, 2021 at 07:48 AM UTC

:tada: This PR is included in version 1.27.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.27.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.27.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/609*
