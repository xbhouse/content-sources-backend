---
type: pull_request
number: 452
title: "Log notices during migrations"
state: merged
author: semtexzv
created: 2021-01-04T07:23:17Z
updated: 2021-03-16T09:30:39Z
closed: 2021-01-04T13:37:38Z
merged: 2021-01-04T13:37:38Z
base_branch: master
head_branch: verbose-migrations
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/452
---

# Pull Request #452: Log notices during migrations

**Author**: @semtexzv
**Created**: January 04, 2021 at 07:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `verbose-migrations`

## Description

Print notices on the client during the migration by using custom DB instance

---

## Discussion

### Comment by @josef-hak on January 04, 2021 at 09:10 AM UTC

@semtexzv there is stylecheck fail:
~~~
database_admin/migrate.go:9:2: ST1019: package "github.com/golang-migrate/migrate/v4/database/postgres" is being imported more than once (stylecheck)
~~~


### Comment by @codecov-io on January 04, 2021 at 09:50 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=h1) Report
> Merging [#452](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=desc) (e2171d8) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/38566078bc9123c3b5ab821d7b087e420e3d7170?el=desc) (3856607) will **decrease** coverage by `0.38%`.
> The diff coverage is `57.14%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #452      +/-   ##
==========================================
- Coverage   62.87%   62.49%   -0.39%     
==========================================
  Files          59       60       +1     
  Lines        2707     2722      +15     
==========================================
- Hits         1702     1701       -1     
- Misses        769      781      +12     
- Partials      236      240       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.36% <ø> (-0.07%)` | :arrow_down: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.14% <ø> (-0.62%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.08% <ø> (-0.51%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `77.50% <ø> (-0.55%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <ø> (-0.68%)` | :arrow_down: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `65.78% <ø> (-0.88%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `90.00% <ø> (-0.33%)` | :arrow_down: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `79.16% <ø> (-0.84%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `30.48% <18.18%> (-1.91%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `50.00% <50.00%> (ø)` | |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=footer). Last update [cda5aad...e2171d8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/452?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on January 04, 2021 at 10:01 AM UTC

Fixed

---

## Reviews

### Review by @josef-hak - Changes Requested on January 04, 2021 at 09:10 AM UTC

Stylecheck.

### Review by @josef-hak - Commented on January 04, 2021 at 10:38 AM UTC

### Review by @semtexzv - Commented on January 04, 2021 at 01:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/452*
