---
type: pull_request
number: 434
title: "feat(Packages): change remediation identifier to patch-package"
state: merged
author: mkholjuraev
created: 2021-02-12T09:57:10Z
updated: 2021-08-09T06:54:51Z
closed: 2021-02-19T10:16:37Z
merged: 2021-02-19T10:16:37Z
base_branch: master
head_branch: change-patch-remediation-identifier
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/434
---

# Pull Request #434: feat(Packages): change remediation identifier to patch-package

**Author**: @mkholjuraev
**Created**: February 12, 2021 at 09:57 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `change-patch-remediation-identifier`

## Description

Resolves: https://issues.redhat.com/browse/SPM-737

- I think we have only 2 usages of remediation (remediation advisories, and remediation packages). Thus I did want to touch on existing, already working code. Instead created a special case for remediating packages.
- some unused code is removed from Packages component

---

## Discussion

### Comment by @codecov-io on February 12, 2021 at 10:01 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=h1) Report
> Merging [#434](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=desc) (bcdb08d) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/714982c751949b5d6d0383017b73e5417c7a9b84?el=desc) (714982c) will **increase** coverage by `0.15%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #434      +/-   ##
==========================================
+ Coverage   69.35%   69.51%   +0.15%     
==========================================
  Files          75       75              
  Lines        1315     1312       -3     
  Branches      175      175              
==========================================
  Hits          912      912              
+ Misses        337      334       -3     
  Partials       66       66              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `90.57% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=footer). Last update [714982c...bcdb08d](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/434?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on February 19, 2021 at 10:16 AM UTC

Nice.

### Comment by @jiridostal on February 19, 2021 at 10:25 AM UTC

:tada: This PR is included in version 1.11.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.11.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.11.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on February 18, 2021 at 12:33 PM UTC

Isn't is better to pass a constant to the remediationProvider, instead a boolean? We can create two constants - for package and for advisories so the developer can pick what he wants (also for potential future remediation types)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/434*
