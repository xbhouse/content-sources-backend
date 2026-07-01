---
type: pull_request
number: 86
title: "Make advisories sortable by type"
state: merged
author: semtexzv
created: 2020-01-29T16:59:42Z
updated: 2020-01-30T07:07:37Z
closed: 2020-01-30T07:07:37Z
merged: 2020-01-30T07:07:37Z
base_branch: master
head_branch: advisories-sort
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/86
---

# Pull Request #86: Make advisories sortable by type

**Author**: @semtexzv
**Created**: January 29, 2020 at 04:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisories-sort`

## Description

- Remove description, summary from sort fields
- Add type to sort keys for advisories
- Add `id` sort key to openapi doc

---

## Discussion

### Comment by @codecov-io on January 29, 2020 at 05:09 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@d07f394`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #86   +/-   ##
=========================================
  Coverage          ?   57.42%           
=========================================
  Files             ?       34           
  Lines             ?     1334           
  Branches          ?        0           
=========================================
  Hits              ?      766           
  Misses            ?      478           
  Partials          ?       90
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `79.72% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.71% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.56% <100%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.15% <100%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=footer). Last update [d07f394...d53fae8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/86?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 30, 2020 at 07:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/86*
