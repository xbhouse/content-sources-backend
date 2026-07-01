---
type: pull_request
number: 177
title: "Split store when building to specific file so we have even better treeshaking"
state: merged
author: karelhala
created: 2020-06-29T13:39:51Z
updated: 2020-07-09T08:14:26Z
closed: 2020-06-30T12:50:42Z
merged: 2020-06-30T12:50:42Z
base_branch: master
head_branch: split-store
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/177
---

# Pull Request #177: Split store when building to specific file so we have even better treeshaking

**Author**: @karelhala
**Created**: June 29, 2020 at 01:39 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `split-store`

## Description

This PR will allow importing only store without the component itself.

ping @jiridostal

---

## Discussion

### Comment by @codecov-commenter on June 29, 2020 at 01:50 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=h1) Report
> Merging [#177](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/5184c93cd6002bee7c4bfd5ee8a500ac2c351b73&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #177   +/-   ##
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


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=footer). Last update [5184c93...23351f4](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/177?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on June 30, 2020 at 12:50 PM UTC

Cool, thanks!

### Comment by @karelhala on July 01, 2020 at 12:10 PM UTC

@jiridostal is there any way how to release new version?

### Comment by @jiridostal on July 09, 2020 at 08:14 AM UTC

:tada: This PR is included in version 0.17.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.17.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.17.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/177*
