---
type: pull_request
number: 266
title: "Experimental package data implementation"
state: merged
author: semtexzv
created: 2020-06-08T07:20:45Z
updated: 2021-03-16T09:28:27Z
closed: 2020-06-16T10:06:21Z
merged: 2020-06-16T10:06:21Z
base_branch: master
head_branch: packages
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/266
---

# Pull Request #266: Experimental package data implementation

**Author**: @semtexzv
**Created**: June 08, 2020 at 07:20 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `packages`

## Description

- Stores packages as JSONB column
- `/packages` endpoint for account-level data sped up by GIN index
- `/systems/:system_id/packages` for per-system package data. 

---

## Discussion

### Comment by @codecov-commenter on June 13, 2020 at 08:39 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=h1) Report
> Merging [#266](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/70511e2f434004fb8b51f26e1adaaf01bf886eeb&el=desc) will **decrease** coverage by `0.18%`.
> The diff coverage is `64.35%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #266      +/-   ##
==========================================
- Coverage   62.32%   62.14%   -0.19%     
==========================================
  Files          47       49       +2     
  Lines        2243     2327      +84     
==========================================
+ Hits         1398     1446      +48     
- Misses        679      704      +25     
- Partials      166      177      +11     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `58.33% <33.33%> (-27.39%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `57.14% <57.14%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.09% <63.63%> (-0.89%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `66.66% <66.66%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.87% <76.92%> (-0.70%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=footer). Last update [70511e2...b25f9c0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/266?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on June 15, 2020 at 09:00 AM UTC

Nice, I like it. I have several notes to it but conceptually it looks good.

### Review by @semtexzv - Commented on June 16, 2020 at 07:41 AM UTC

### Review by @semtexzv - Commented on June 16, 2020 at 07:44 AM UTC

### Review by @semtexzv - Commented on June 16, 2020 at 07:50 AM UTC

### Review by @semtexzv - Commented on June 16, 2020 at 07:54 AM UTC

### Review by @semtexzv - Commented on June 16, 2020 at 08:05 AM UTC

### Review by @josef-hak - Approved on June 16, 2020 at 08:22 AM UTC

Great work, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/266*
