---
type: pull_request
number: 327
title: "chore(optimization) some shared presentational components"
state: merged
author: mkholjuraev
created: 2020-11-11T22:50:51Z
updated: 2021-08-09T06:55:46Z
closed: 2020-11-19T08:56:13Z
merged: 2020-11-19T08:56:13Z
base_branch: master
head_branch: presentational-components-datamappers-helpers
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/327
---

# Pull Request #327: chore(optimization) some shared presentational components

**Author**: @mkholjuraev
**Created**: November 11, 2020 at 10:50 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `presentational-components-datamappers-helpers`

## Description

A partl of: [SPM-614](https://issues.redhat.com/browse/SPM-614) 

---

## Discussion

### Comment by @codecov-io on November 11, 2020 at 10:56 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=h1) Report
> Merging [#327](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=desc) (7e16a42) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/cc37812b4beb85cefede9c4154ca86e5605a9986?el=desc) (cc37812) will **increase** coverage by `1.43%`.
> The diff coverage is `80.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #327      +/-   ##
==========================================
+ Coverage   63.85%   65.29%   +1.43%     
==========================================
  Files          64       66       +2     
  Lines        1220     1187      -33     
  Branches      157      154       -3     
==========================================
- Hits          779      775       -4     
+ Misses        381      355      -26     
+ Partials       60       57       -3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `76.66% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `95.38% <ø> (ø)` | |
| [...resentationalComponents/Snippets/AdvisoriesIcon.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9BZHZpc29yaWVzSWNvbi5qcw==) | `66.66% <66.66%> (ø)` | |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | `100.00% <100.00%> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `96.07% <0.00%> (-0.85%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <0.00%> (ø)` | |
| [...sentationalComponents/Filters/PublishDateFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1B1Ymxpc2hEYXRlRmlsdGVyLmpz) | `100.00% <0.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=footer). Last update [cc37812...7e16a42](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/327?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on December 01, 2020 at 01:03 PM UTC

:tada: This PR is included in version 1.0.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.0.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.0.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on November 13, 2020 at 10:15 AM UTC

Left a few comments, other than that it looks great!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/327*
