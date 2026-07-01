---
type: pull_request
number: 314
title: "added support for custom rbac roles"
state: merged
author: josef-hak
created: 2020-08-26T13:41:42Z
updated: 2020-09-08T10:03:20Z
closed: 2020-09-07T11:31:35Z
merged: 2020-09-07T11:31:35Z
base_branch: master
head_branch: custom-rbac-roles
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/314
---

# Pull Request #314: added support for custom rbac roles

**Author**: @josef-hak
**Created**: August 26, 2020 at 01:41 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `custom-rbac-roles`

## Description

related to rbac-config patch update: https://github.com/RedHatInsights/rbac-config/pull/84

---

## Discussion

### Comment by @codecov-commenter on August 26, 2020 at 01:53 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=h1) Report
> Merging [#314](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/06f1f4a0bf42675367147f40e97519c3bf8b3369?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #314   +/-   ##
=======================================
  Coverage   61.69%   61.69%           
=======================================
  Files          53       53           
  Lines        2600     2600           
=======================================
  Hits         1604     1604           
  Misses        792      792           
  Partials      204      204           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=footer). Last update [06f1f4a...3effced](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/314?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on August 27, 2020 at 01:33 PM UTC

We must synchronize pushing these changes to changes in RBAC config.

### Comment by @josef-hak on August 27, 2020 at 02:20 PM UTC

@semtexzv I don't think we need to synchronize it anyhow. Even if we can not expect "all:all" from RBAC service we still accept already supported "*:*". Actually shouldn't I leave `"*":"*"` row in patch rbac-config?  

---

## Reviews

### Review by @semtexzv - Approved on August 27, 2020 at 01:33 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/314*
