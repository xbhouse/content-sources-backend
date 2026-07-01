---
type: pull_request
number: 390
title: "feat(GenericError): use frontend components generic error component"
state: merged
author: mkholjuraev
created: 2021-01-13T08:42:13Z
updated: 2021-08-09T06:55:05Z
closed: 2021-01-15T12:53:44Z
merged: 2021-01-15T12:53:44Z
base_branch: master
head_branch: use-fe-generic-error
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/390
---

# Pull Request #390: feat(GenericError): use frontend components generic error component

**Author**: @mkholjuraev
**Created**: January 13, 2021 at 08:42 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `use-fe-generic-error`

## Description

resolves: https://issues.redhat.com/browse/SPM-683

![image](https://user-images.githubusercontent.com/59481011/104427491-98595480-5583-11eb-9309-2c2315a968f4.png)


---

## Discussion

### Comment by @codecov-io on January 13, 2021 at 08:45 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=h1) Report
> Merging [#390](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=desc) (cd86835) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/6312ae1a81bc9d988083de29e65d3d7edf802a99?el=desc) (6312ae1) will **increase** coverage by `0.05%`.
> The diff coverage is `75.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #390      +/-   ##
==========================================
+ Coverage   72.80%   72.86%   +0.05%     
==========================================
  Files          69       69              
  Lines        1239     1227      -12     
  Branches      161      161              
==========================================
- Hits          902      894       -8     
+ Misses        283      279       -4     
  Partials       54       54              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `100.00% <ø> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `0.00% <ø> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `95.23% <100.00%> (-0.22%)` | :arrow_down: |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `95.74% <100.00%> (ø)` | |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=footer). Last update [6312ae1...cd86835](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/390?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on January 15, 2021 at 01:01 PM UTC

:tada: This PR is included in version 1.7.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.7.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.7.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/390*
