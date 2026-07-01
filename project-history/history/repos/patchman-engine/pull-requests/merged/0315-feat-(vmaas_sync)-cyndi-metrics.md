---
type: pull_request
number: 315
title: "feat (vmaas_sync): Cyndi metrics"
state: merged
author: AlesKas
created: 2020-08-27T09:37:14Z
updated: 2020-10-03T07:38:50Z
closed: 2020-10-02T09:33:40Z
merged: 2020-10-02T09:33:40Z
base_branch: master
head_branch: SPM-375
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/315
---

# Pull Request #315: feat (vmaas_sync): Cyndi metrics

**Author**: @AlesKas
**Created**: August 27, 2020 at 09:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-375`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 27, 2020 at 12:27 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=h1) Report
> Merging [#315](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fd945f58e68459686e7a9b3b035259175b7217e4?el=desc) will **decrease** coverage by `0.70%`.
> The diff coverage is `37.25%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #315      +/-   ##
==========================================
- Coverage   62.44%   61.73%   -0.71%     
==========================================
  Files          56       57       +1     
  Lines        2524     2587      +63     
==========================================
+ Hits         1576     1597      +21     
- Misses        724      758      +34     
- Partials      224      232       +8     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.25% <0.00%> (-1.22%)` | :arrow_down: |
| [vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `40.42% <40.42%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `51.61% <0.00%> (-7.01%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.76% <0.00%> (-3.33%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `67.66% <0.00%> (-0.66%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <0.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `32.39% <0.00%> (+0.96%)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `80.00% <0.00%> (+1.05%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=footer). Last update [fd945f5...d9662c4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/315?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on October 02, 2020 at 07:41 AM UTC

@AlesKas do we really need
~~~
Merge confict

Merge conflict

Merge confict

Merge conflict
~~~
... in the commit message?

---

## Reviews

### Review by @josef-hak - Changes Requested on August 27, 2020 at 01:33 PM UTC

Good, but some updates are needed.

### Review by @AlesKas - Commented on August 28, 2020 at 07:49 AM UTC

### Review by @semtexzv - Commented on September 02, 2020 at 08:29 AM UTC

### Review by @semtexzv - Commented on September 03, 2020 at 12:11 PM UTC

### Review by @semtexzv - Commented on September 03, 2020 at 12:13 PM UTC

### Review by @semtexzv - Commented on September 04, 2020 at 08:11 AM UTC

### Review by @josef-hak - Changes Requested on September 11, 2020 at 07:28 AM UTC

Ok, that's much better, but still something to improve imho.

### Review by @josef-hak - Changes Requested on October 01, 2020 at 11:12 AM UTC

### Review by @AlesKas - Commented on October 01, 2020 at 11:18 AM UTC

### Review by @josef-hak - Changes Requested on October 01, 2020 at 01:23 PM UTC

Good, just a few updates, and it's done :tada: 

### Review by @AlesKas - Commented on October 01, 2020 at 01:58 PM UTC

### Review by @AlesKas - Commented on October 01, 2020 at 02:08 PM UTC

### Review by @josef-hak - Commented on October 02, 2020 at 07:46 AM UTC

### Review by @josef-hak - Changes Requested on October 02, 2020 at 08:58 AM UTC

### Review by @josef-hak - Commented on October 02, 2020 at 09:33 AM UTC

Great, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/315*
