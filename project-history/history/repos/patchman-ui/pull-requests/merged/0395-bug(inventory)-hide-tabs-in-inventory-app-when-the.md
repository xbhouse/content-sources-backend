---
type: pull_request
number: 395
title: "bug(Inventory): hide tabs in Inventory app when there is not systems \u2026"
state: merged
author: mkholjuraev
created: 2021-01-18T16:21:07Z
updated: 2021-08-09T06:55:00Z
closed: 2021-01-19T09:39:21Z
merged: 2021-01-19T09:39:21Z
base_branch: master
head_branch: hide-tabs-inventory
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/395
---

# Pull Request #395: bug(Inventory): hide tabs in Inventory app when there is not systems …

**Author**: @mkholjuraev
**Created**: January 18, 2021 at 04:21 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `hide-tabs-inventory`

## Description

resolves: https://issues.redhat.com/projects/SPM/issues/SPM-691?filter=allopenissues

---

## Discussion

### Comment by @codecov-io on January 18, 2021 at 04:38 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=h1) Report
> Merging [#395](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=desc) (f36a464) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/f04525da3df2d0de151bb4bc4540e2679ab99579?el=desc) (f04525d) will **decrease** coverage by `0.03%`.
> The diff coverage is `81.81%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #395      +/-   ##
==========================================
- Coverage   72.88%   72.85%   -0.04%     
==========================================
  Files          69       69              
  Lines        1228     1234       +6     
  Branches      158      160       +2     
==========================================
+ Hits          895      899       +4     
  Misses        280      280              
- Partials       53       55       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395/diff?src=pr&el=tree#diff-c3JjL2luZGV4Lmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/SystemDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvU3lzdGVtRGV0YWlsLmpz) | `80.00% <66.66%> (-20.00%)` | :arrow_down: |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `95.74% <100.00%> (ø)` | |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=footer). Last update [f04525d...f36a464](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/395?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on January 25, 2021 at 09:07 AM UTC

:tada: This PR is included in version 1.8.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.8.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.8.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/395*
