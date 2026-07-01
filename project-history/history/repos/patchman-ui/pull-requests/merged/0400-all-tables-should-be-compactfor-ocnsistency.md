---
type: pull_request
number: 400
title: "All tables should be compactfor ocnsistency"
state: merged
author: jiridostal
created: 2021-01-21T11:00:24Z
updated: 2022-08-02T08:41:09Z
closed: 2021-01-25T08:59:45Z
merged: 2021-01-25T08:59:44Z
base_branch: master
head_branch: tables-compact
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/400
---

# Pull Request #400: All tables should be compactfor ocnsistency

**Author**: @jiridostal
**Created**: January 21, 2021 at 11:00 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `tables-compact`

## Description

![image](https://user-images.githubusercontent.com/6286045/105342322-7fcee700-5be0-11eb-8387-3c44026a9548.png)
![image](https://user-images.githubusercontent.com/6286045/105342342-89f0e580-5be0-11eb-9fe0-a580c63440ec.png)
![image](https://user-images.githubusercontent.com/6286045/105342369-9117f380-5be0-11eb-960f-50270732851a.png)
![image](https://user-images.githubusercontent.com/6286045/105342389-98d79800-5be0-11eb-948c-2b1115f5892b.png)


---

## Discussion

### Comment by @codecov-io on January 21, 2021 at 11:02 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=h1) Report
> Merging [#400](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=desc) (5bb3537) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/059785a9c49f1f390abfb0da911b5adb3930a645?el=desc) (059785a) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #400   +/-   ##
=======================================
  Coverage   72.85%   72.85%           
=======================================
  Files          69       69           
  Lines        1234     1234           
  Branches      160      160           
=======================================
  Hits          899      899           
  Misses        280      280           
  Partials       55       55           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/TableView/TableViewAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `95.23% <ø> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `95.74% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `76.66% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=footer). Last update [059785a...5bb3537](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/400?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @terezanovotna on January 21, 2021 at 12:59 PM UTC

Looks great! Thanks for changing it up so quickly. Feel free to tag me even here on gh when we have any UI updates

### Comment by @jiridostal on January 25, 2021 at 09:07 AM UTC

:tada: This PR is included in version 1.8.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.8.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.8.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on January 22, 2021 at 03:54 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/400*
