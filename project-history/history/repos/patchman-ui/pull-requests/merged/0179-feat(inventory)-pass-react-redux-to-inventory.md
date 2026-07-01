---
type: pull_request
number: 179
title: "feat(inventory): pass react-redux to inventory"
state: merged
author: karelhala
created: 2020-07-05T08:47:07Z
updated: 2020-07-09T08:14:29Z
closed: 2020-07-09T08:06:13Z
merged: 2020-07-09T08:06:13Z
base_branch: master
head_branch: pass-redux-to-inventory
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/179
---

# Pull Request #179: feat(inventory): pass react-redux to inventory

**Author**: @karelhala
**Created**: July 05, 2020 at 08:47 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `pass-redux-to-inventory`

## Description

We are working on inventory refactoring, since we'd like to fully use hooks and such we have to use App's react dependency. With this being used react-redux needs to be passed from application as well. If you are cautious about the size of your bundle I can pick only the functions really required by inventory.

---

## Discussion

### Comment by @karelhala on July 05, 2020 at 08:47 AM UTC

ping @jiridostal 

### Comment by @codecov-commenter on July 05, 2020 at 08:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=h1) Report
> Merging [#179](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/977f6a5b39f8e39684b59719c06704008650af25&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #179   +/-   ##
=======================================
  Coverage   37.04%   37.04%           
=======================================
  Files          51       51           
  Lines         845      845           
  Branches       94       94           
=======================================
  Hits          313      313           
  Misses        479      479           
  Partials       53       53           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=footer). Last update [977f6a5...9d8d802](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/179?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on July 09, 2020 at 08:14 AM UTC

:tada: This PR is included in version 0.17.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.17.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.17.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/179*
