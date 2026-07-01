---
type: pull_request
number: 84
title: "Fix advisory counts per type in advisory/id/systems"
state: merged
author: semtexzv
created: 2020-01-29T16:33:30Z
updated: 2021-03-16T09:27:06Z
closed: 2020-01-30T10:08:39Z
merged: 2020-01-30T10:08:39Z
base_branch: master
head_branch: advisory-counts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/84
---

# Pull Request #84: Fix advisory counts per type in advisory/id/systems

**Author**: @semtexzv
**Created**: January 29, 2020 at 04:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisory-counts`

## Description

Data provided in `/advisoies/:id/systems` was invalid.

---

## Discussion

### Comment by @codecov-io on January 30, 2020 at 09:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@d07f394`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #84   +/-   ##
=========================================
  Coverage          ?   57.74%           
=========================================
  Files             ?       34           
  Lines             ?     1363           
  Branches          ?        0           
=========================================
  Hits              ?      787           
  Misses            ?      482           
  Partials          ?       94
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `80.51% <100%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=footer). Last update [d07f394...ed52c89](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/84?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 30, 2020 at 10:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/84*
