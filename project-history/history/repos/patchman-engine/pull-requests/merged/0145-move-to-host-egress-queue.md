---
type: pull_request
number: 145
title: "Move to host-egress queue"
state: merged
author: semtexzv
created: 2020-02-24T11:04:53Z
updated: 2021-03-16T09:27:34Z
closed: 2020-02-28T13:06:51Z
merged: 2020-02-28T13:06:51Z
base_branch: master
head_branch: host-egress
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/145
---

# Pull Request #145: Move to host-egress queue

**Author**: @semtexzv
**Created**: February 24, 2020 at 11:04 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `host-egress`

## Description

Allows us to not call inventory at all, we can get all the info we need from kafka messages.
- Changes `mqueue.go` to only deal with kafka messages
- `events.go` now deals with standard event format for delete and eval messages.
- `upload.go` Parses messages into custom event format
- Platform mock code now uses `map[string]interface{}` instead of string interpolation

Tied with https://github.com/RedHatInsights/e2e-deploy/pull/1124


---

## Discussion

### Comment by @codecov-io on February 24, 2020 at 11:38 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=h1) Report
> Merging [#145](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e77384b3907426b4f7b189876f535a652d621eaa?src=pr&el=desc) will **decrease** coverage by `1.05%`.
> The diff coverage is `70%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #145      +/-   ##
==========================================
- Coverage   64.94%   63.88%   -1.06%     
==========================================
  Files          42       42              
  Lines        1626     1584      -42     
==========================================
- Hits         1056     1012      -44     
- Misses        450      456       +6     
+ Partials      120      116       -4
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.5% <ø> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `30.43% <ø> (ø)` | :arrow_up: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `77.5% <100%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.72% <100%> (ø)` | :arrow_up: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `83.33% <100%> (ø)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `81.25% <100%> (-5.12%)` | :arrow_down: |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `100% <100%> (ø)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.11% <60%> (-11.12%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `59.59% <63.88%> (-10.78%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=footer). Last update [e77384b...5877051](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/145?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on February 26, 2020 at 01:04 PM UTC

@semtexzv It seems that commit message is out of date.

---

## Reviews

### Review by @josef-hak - Changes Requested on February 26, 2020 at 01:17 PM UTC

Great update. It simplifies it a lot. Just added some notes.

### Review by @semtexzv - Commented on February 27, 2020 at 09:59 AM UTC

### Review by @semtexzv - Commented on February 28, 2020 at 10:05 AM UTC

### Review by @semtexzv - Commented on February 28, 2020 at 10:05 AM UTC

### Review by @josef-hak - Commented on February 28, 2020 at 12:31 PM UTC

### Review by @josef-hak - Approved on February 28, 2020 at 01:06 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/145*
