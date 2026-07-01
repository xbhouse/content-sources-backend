---
type: pull_request
number: 450
title: "added optional env var RESET_SCHEMA to database-admin"
state: merged
author: josef-hak
created: 2020-12-17T11:49:45Z
updated: 2021-02-04T10:49:51Z
closed: 2021-01-04T11:03:17Z
merged: 2021-01-04T11:03:17Z
base_branch: master
head_branch: enable-schema-reset
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/450
---

# Pull Request #450: added optional env var RESET_SCHEMA to database-admin

**Author**: @josef-hak
**Created**: December 17, 2020 at 11:49 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `enable-schema-reset`

## Description

- allowing to drop and recreate schema (of specific version when
SCHEMA_MIGRATION is used in addition).


---

## Discussion

### Comment by @codecov-io on December 17, 2020 at 12:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=h1) Report
> Merging [#450](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=desc) (52ec637) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/38566078bc9123c3b5ab821d7b087e420e3d7170?el=desc) (3856607) will **decrease** coverage by `0.38%`.
> The diff coverage is `65.38%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #450      +/-   ##
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


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.36% <ø> (-0.07%)` | :arrow_down: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.14% <ø> (-0.62%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.08% <ø> (-0.51%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `77.50% <ø> (-0.55%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <ø> (-0.68%)` | :arrow_down: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `65.78% <ø> (-0.88%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `90.00% <ø> (-0.33%)` | :arrow_down: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `79.16% <ø> (-0.84%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `50.00% <50.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.01% <69.23%> (-0.94%)` | :arrow_down: |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=footer). Last update [9dabf7a...52ec637](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/450?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on January 04, 2021 at 05:54 AM UTC

### Review by @josef-hak - Commented on January 04, 2021 at 09:18 AM UTC

### Review by @semtexzv - Approved on January 04, 2021 at 11:02 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/450*
