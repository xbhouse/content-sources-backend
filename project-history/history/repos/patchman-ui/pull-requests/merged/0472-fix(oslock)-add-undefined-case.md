---
type: pull_request
number: 472
title: "fix(OSLock): add undefined case"
state: merged
author: mkholjuraev
created: 2021-03-22T11:09:15Z
updated: 2021-08-09T06:56:24Z
closed: 2021-03-26T10:31:15Z
merged: 2021-03-26T10:31:15Z
base_branch: master
head_branch: better-rhsm-condition
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/472
---

# Pull Request #472: fix(OSLock): add undefined case

**Author**: @mkholjuraev
**Created**: March 22, 2021 at 11:09 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `better-rhsm-condition`

## Description

removes incorrect OS locks 

---

## Discussion

### Comment by @codecov-io on March 22, 2021 at 11:12 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=h1) Report
> Merging [#472](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=desc) (17dc2d8) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/ec4d5a5ba0f004a5f4035cf6b1e87977a3c05b5c?el=desc) (ec4d5a5) will **not change** coverage.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #472   +/-   ##
=======================================
  Coverage   72.32%   72.32%           
=======================================
  Files          75       75           
  Lines        1319     1319           
  Branches      176      176           
=======================================
  Hits          954      954           
  Misses        316      316           
  Partials       49       49           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `90.64% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=footer). Last update [ec4d5a5...17dc2d8](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/472?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on March 26, 2021 at 10:40 AM UTC

:tada: This PR is included in version 1.14.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.14.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.14.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/472*
