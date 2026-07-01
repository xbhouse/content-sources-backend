---
type: pull_request
number: 53
title: "Add basic sorting"
state: merged
author: semtexzv
created: 2020-01-13T13:33:23Z
updated: 2020-01-16T09:05:13Z
closed: 2020-01-16T09:05:13Z
merged: 2020-01-16T09:05:13Z
base_branch: master
head_branch: sorting
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/53
---

# Pull Request #53: Add basic sorting

**Author**: @semtexzv
**Created**: January 13, 2020 at 01:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sorting`

## Description

- Adds a method which by default applies sort by internal ID, and allows adding new fields to sort by easily

---

## Discussion

### Comment by @codecov-io on January 13, 2020 at 01:40 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=h1) Report
> Merging [#53](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/949d233a2a1ad7d2683d69cc67f3c0c59d73706f?src=pr&el=desc) will **increase** coverage by `0.12%`.
> The diff coverage is `76.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #53      +/-   ##
==========================================
+ Coverage   58.16%   58.29%   +0.12%     
==========================================
  Files          29       29              
  Lines        1114     1151      +37     
==========================================
+ Hits          648      671      +23     
- Misses        392      401       +9     
- Partials       74       79       +5
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `66.03% <25%> (-3.36%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `79.66% <57.14%> (-3.68%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `81.25% <66.66%> (-3.75%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `66.66% <86.66%> (+33.33%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75% <88%> (-1.2%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=footer). Last update [949d233...7ba1250](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/53?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on January 14, 2020 at 08:38 AM UTC

@semtexzv please, resolve conflicts and re-request review then.

---

## Reviews

### Review by @josef-hak - Changes Requested on January 13, 2020 at 02:21 PM UTC

Pretty good solution to add such generic method for sorting. I would just slightly improve it.

### Review by @semtexzv - Commented on January 14, 2020 at 08:02 AM UTC

### Review by @josef-hak - Commented on January 14, 2020 at 08:31 AM UTC

### Review by @semtexzv - Commented on January 15, 2020 at 12:04 PM UTC

### Review by @josef-hak - Commented on January 16, 2020 at 08:57 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/53*
