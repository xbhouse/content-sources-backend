---
type: pull_request
number: 76
title: "Make resource limits more reasonable"
state: merged
author: semtexzv
created: 2020-01-27T13:03:20Z
updated: 2021-03-16T09:27:01Z
closed: 2020-01-27T13:15:46Z
merged: 2020-01-27T13:15:46Z
base_branch: master
head_branch: resources
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/76
---

# Pull Request #76: Make resource limits more reasonable

**Author**: @semtexzv
**Created**: January 27, 2020 at 01:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `resources`

## Description

- Listener does almost nothing, Current implementation even less so. but we need multiple pods
- Database runs on 300MB RAM right now, without any traffic

---

## Discussion

### Comment by @codecov-io on January 27, 2020 at 01:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@12ba442`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #76   +/-   ##
=========================================
  Coverage          ?   56.25%           
=========================================
  Files             ?       32           
  Lines             ?     1303           
  Branches          ?        0           
=========================================
  Hits              ?      733           
  Misses            ?      478           
  Partials          ?       92
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76?src=pr&el=footer). Last update [12ba442...f966518](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/76?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 27, 2020 at 01:09 PM UTC

### Review by @josef-hak - Commented on January 27, 2020 at 01:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/76*
