---
type: pull_request
number: 274
title: "fix(SystemPackages): Empty state added for empty table in Systems packages page"
state: merged
author: mkholjuraev
created: 2020-09-03T10:59:43Z
updated: 2021-08-09T06:55:18Z
closed: 2020-09-15T12:12:00Z
merged: 2020-09-15T12:12:00Z
base_branch: master
head_branch: empty-list-system-packages
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/274
---

# Pull Request #274: fix(SystemPackages): Empty state added for empty table in Systems packages page

**Author**: @mkholjuraev
**Created**: September 03, 2020 at 10:59 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `empty-list-system-packages`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 03, 2020 at 11:11 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=h1) Report
> Merging [#274](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/06a89f0acca7f84ee64f99ddc78ead0c5b43fadc?el=desc) will **increase** coverage by `0.63%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #274      +/-   ##
==========================================
+ Coverage   66.46%   67.09%   +0.63%     
==========================================
  Files          57       57              
  Lines         999     1003       +4     
  Branches      118      119       +1     
==========================================
+ Hits          664      673       +9     
+ Misses        296      292       -4     
+ Partials       39       38       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `1.69% <ø> (ø)` | |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `100.00% <100.00%> (+25.00%)` | :arrow_up: |
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=footer). Last update [06a89f0...a12849d](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/274?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 15, 2020 at 12:27 PM UTC

:tada: This PR is included in version 0.24.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on September 08, 2020 at 02:10 PM UTC

Can you please resolve conflicts?

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/274*
