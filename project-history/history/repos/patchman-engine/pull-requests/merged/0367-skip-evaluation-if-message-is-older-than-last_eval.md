---
type: pull_request
number: 367
title: "Skip evaluation if message is older than last_evaluation"
state: merged
author: semtexzv
created: 2020-10-07T10:10:23Z
updated: 2021-03-16T09:29:34Z
closed: 2020-10-08T14:58:50Z
merged: 2020-10-08T14:58:50Z
base_branch: master
head_branch: skip-eval-timestamp
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/367
---

# Pull Request #367: Skip evaluation if message is older than last_evaluation

**Author**: @semtexzv
**Created**: October 07, 2020 at 10:10 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `skip-eval-timestamp`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on October 07, 2020 at 12:16 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=h1) Report
> Merging [#367](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fee97a7f77d067ac041bc68090e9d9db66b2f89d?el=desc) will **decrease** coverage by `0.68%`.
> The diff coverage is `47.29%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #367      +/-   ##
==========================================
- Coverage   62.27%   61.59%   -0.69%     
==========================================
  Files          56       57       +1     
  Lines        2534     2588      +54     
==========================================
+ Hits         1578     1594      +16     
- Misses        730      761      +31     
- Partials      226      233       +7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.25% <0.00%> (-1.22%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.84% <20.00%> (-0.82%)` | :arrow_down: |
| ... and [10 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=footer). Last update [4c932b5...b5c65ff](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/367?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on October 08, 2020 at 01:09 PM UTC

@semtexzv ok, just ensure passing pipeline and merge it


---

## Reviews

### Review by @josef-hak - Changes Requested on October 08, 2020 at 09:04 AM UTC

Good, some notes and proposals introduced.

### Review by @semtexzv - Commented on October 08, 2020 at 09:57 AM UTC

### Review by @josef-hak - Commented on October 08, 2020 at 10:23 AM UTC

### Review by @semtexzv - Commented on October 08, 2020 at 12:48 PM UTC

### Review by @semtexzv - Commented on October 08, 2020 at 01:03 PM UTC

### Review by @josef-hak - Commented on October 08, 2020 at 01:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/367*
