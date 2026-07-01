---
type: pull_request
number: 71
title: "Separate kafka readers to base package"
state: merged
author: semtexzv
created: 2020-01-22T16:05:08Z
updated: 2021-03-16T09:27:00Z
closed: 2020-01-27T09:38:41Z
merged: 2020-01-27T09:38:41Z
base_branch: master
head_branch: base-kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/71
---

# Pull Request #71: Separate kafka readers to base package

**Author**: @semtexzv
**Created**: January 22, 2020 at 04:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `base-kafka`

## Description

Moves code that deals with kafka readers and parsing platform events to base package.
Blocks separating evaluator from listener.

---

## Discussion

### Comment by @codecov-io on January 22, 2020 at 04:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=h1) Report
> Merging [#71](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/7f846e224be89cdad3824070dd4f8fb79561b8e1?src=pr&el=desc) will **decrease** coverage by `0.2%`.
> The diff coverage is `62.06%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #71      +/-   ##
==========================================
- Coverage   57.92%   57.72%   -0.21%     
==========================================
  Files          30       31       +1     
  Lines        1212     1237      +25     
==========================================
+ Hits          702      714      +12     
- Misses        432      438       +6     
- Partials       78       85       +7
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `69.31% <100%> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `44.44% <100%> (ø)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `44.44% <33.33%> (-8.89%)` | :arrow_down: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50% <50%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `66.66% <66.66%> (ø)` | |
| [base/database/batch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9iYXRjaC5nbw==) | `65.28% <0%> (-2.94%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `82.08% <0%> (ø)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.67% <0%> (ø)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `66.66% <0%> (ø)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `65.38% <0%> (ø)` | :arrow_up: |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=footer). Last update [7f846e2...3e6345f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/71?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 23, 2020 at 08:51 AM UTC

Well, a lot of work done. I've proposed some improvements.

### Review by @semtexzv - Commented on January 24, 2020 at 09:41 AM UTC

### Review by @semtexzv - Commented on January 24, 2020 at 09:44 AM UTC

### Review by @semtexzv - Commented on January 27, 2020 at 08:52 AM UTC

### Review by @semtexzv - Commented on January 27, 2020 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/71*
