---
type: pull_request
number: 344
title: "master -> stable"
state: merged
author: josef-hak
created: 2020-09-23T09:37:51Z
updated: 2020-09-25T11:45:09Z
closed: 2020-09-25T11:45:09Z
merged: 2020-09-25T11:45:09Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/344
---

# Pull Request #344: master -> stable

**Author**: @josef-hak
**Created**: September 23, 2020 at 09:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

- fixed tag parsing regex
- added reporter info (metrics and system_platform.reporter_id column)

---

## Discussion

### Comment by @codecov-commenter on September 23, 2020 at 09:45 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=h1) Report
> Merging [#344](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f420be6ccf3c3b16c0485cec9271bfac58c36f73?el=desc) will **decrease** coverage by `1.47%`.
> The diff coverage is `75.49%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #344      +/-   ##
==========================================
- Coverage   63.32%   61.85%   -1.48%     
==========================================
  Files          54       56       +2     
  Lines        2686     2464     -222     
==========================================
- Hits         1701     1524     -177     
+ Misses        773      721      -52     
- Partials      212      219       +7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `73.33% <ø> (-1.67%)` | :arrow_down: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.73% <0.00%> (-2.64%)` | :arrow_down: |
| [vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `60.71% <60.71%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.94% <70.00%> (-2.18%)` | :arrow_down: |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `85.18% <85.18%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `78.94% <100.00%> (+8.35%)` | :arrow_up: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <100.00%> (-4.17%)` | :arrow_down: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `66.66% <0.00%> (-7.25%)` | :arrow_down: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <0.00%> (-7.22%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `66.66% <0.00%> (-5.80%)` | :arrow_down: |
| ... and [52 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=footer). Last update [f420be6...732bc4b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/344?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on September 25, 2020 at 11:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/344*
