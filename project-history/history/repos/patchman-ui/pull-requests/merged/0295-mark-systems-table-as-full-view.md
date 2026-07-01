---
type: pull_request
number: 295
title: "Mark systems table as full view"
state: merged
author: karelhala
created: 2020-09-24T09:11:13Z
updated: 2020-09-29T11:37:08Z
closed: 2020-09-24T10:18:47Z
merged: 2020-09-24T10:18:47Z
base_branch: master
head_branch: inv-full-view
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/295
---

# Pull Request #295: Mark systems table as full view

**Author**: @karelhala
**Created**: September 24, 2020 at 09:11 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `inv-full-view`

## Description

### Inventory full view [1]

In order to show proper RBAC state in systems table we need to pass `fullView` prop to inventory component. This applies only to systems page as it's the only place where inventory table is just one component on screen. Without this prop the `unAuthorized` component would be in table, which is not correct based on our mocks.

[1] https://github.com/RedHatInsights/frontend-components/blob/master/packages/inventory/doc/inventory.md#using-rbac-with-inventory

---

## Discussion

### Comment by @karelhala on September 24, 2020 at 09:11 AM UTC

ping @jiridostal @mkholjuraev

### Comment by @codecov-commenter on September 24, 2020 at 09:22 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=h1) Report
> Merging [#295](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/81e5c5001622d98ad256b0e73fcd944d6fb85eee?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #295   +/-   ##
=======================================
  Coverage   67.57%   67.57%           
=======================================
  Files          56       56           
  Lines         996      996           
  Branches      119      119           
=======================================
  Hits          673      673           
  Misses        285      285           
  Partials       38       38           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=footer). Last update [81e5c50...00f8aea](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/295?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 29, 2020 at 11:37 AM UTC

:tada: This PR is included in version 0.24.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on September 24, 2020 at 09:16 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/295*
