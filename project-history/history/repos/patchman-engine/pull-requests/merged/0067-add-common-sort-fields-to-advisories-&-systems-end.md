---
type: pull_request
number: 67
title: "Add common sort fields to advisories & systems endpoints"
state: merged
author: semtexzv
created: 2020-01-20T13:28:20Z
updated: 2021-03-16T09:26:59Z
closed: 2020-01-24T09:47:13Z
merged: 2020-01-24T09:47:13Z
base_branch: master
head_branch: sorts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/67
---

# Pull Request #67: Add common sort fields to advisories & systems endpoints

**Author**: @semtexzv
**Created**: January 20, 2020 at 01:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sorts`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 23, 2020 at 08:28 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=h1) Report
> Merging [#67](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/7f846e224be89cdad3824070dd4f8fb79561b8e1?src=pr&el=desc) will **increase** coverage by `0.06%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #67      +/-   ##
==========================================
+ Coverage   57.92%   57.99%   +0.06%     
==========================================
  Files          30       30              
  Lines        1212     1214       +2     
==========================================
+ Hits          702      704       +2     
  Misses        432      432              
  Partials       78       78
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `82.08% <100%> (ø)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `65.38% <100%> (ø)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.95% <100%> (ø)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `69.23% <100%> (+2.56%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.67% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=footer). Last update [7f846e2...93d1dbc](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/67?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on January 24, 2020 at 09:40 AM UTC

Added explicit selects & openapi docs

---

## Reviews

### Review by @beav - Commented on January 21, 2020 at 01:23 PM UTC

### Review by @josef-hak - Changes Requested on January 22, 2020 at 07:41 AM UTC

### Review by @semtexzv - Commented on January 22, 2020 at 04:06 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/67*
