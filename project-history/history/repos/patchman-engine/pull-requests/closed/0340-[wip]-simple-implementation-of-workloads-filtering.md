---
type: pull_request
number: 340
title: "[WIP] Simple implementation of workloads filtering"
state: closed
author: semtexzv
created: 2020-09-18T13:22:42Z
updated: 2021-03-16T09:28:54Z
closed: 2020-09-30T10:32:26Z
base_branch: master
head_branch: simple-workloads
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/340
---

# Pull Request #340: [WIP] Simple implementation of workloads filtering

**Author**: @semtexzv
**Created**: September 18, 2020 at 01:22 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `simple-workloads`

## Description

Edited to accept what platform expects

---

## Discussion

### Comment by @codecov-commenter on September 18, 2020 at 01:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=h1) Report
> Merging [#340](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b0bb5fcbc99869111a60167b903c296616ec23b8?el=desc) will **decrease** coverage by `1.71%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #340      +/-   ##
==========================================
- Coverage   63.30%   61.59%   -1.72%     
==========================================
  Files          54       54              
  Lines        2679     2393     -286     
==========================================
- Hits         1696     1474     -222     
+ Misses        772      706      -66     
- Partials      211      213       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `87.01% <100.00%> (-1.23%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.00% <100.00%> (-2.09%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `70.00% <100.00%> (-2.73%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <100.00%> (-2.58%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `83.33% <100.00%> (-1.86%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `75.93% <100.00%> (+0.93%)` | :arrow_up: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `66.66% <0.00%> (-7.25%)` | :arrow_down: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <0.00%> (-7.22%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `66.66% <0.00%> (-5.80%)` | :arrow_down: |
| [base/utils/awscloudwatch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9hd3NjbG91ZHdhdGNoLmdv) | `13.79% <0.00%> (-5.57%)` | :arrow_down: |
| ... and [50 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=footer). Last update [b0bb5fc...edfbbe0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/340?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on September 30, 2020 at 10:32 AM UTC

Superseded by https://github.com/RedHatInsights/patchman-engine/pull/349

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/340*
