---
type: pull_request
number: 110
title: "Send recalculate messages in batch"
state: merged
author: semtexzv
created: 2020-02-10T08:57:05Z
updated: 2021-03-16T09:27:30Z
closed: 2020-02-18T12:11:44Z
merged: 2020-02-18T12:11:43Z
base_branch: master
head_branch: batch-recalc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/110
---

# Pull Request #110: Send recalculate messages in batch

**Author**: @semtexzv
**Created**: February 10, 2020 at 08:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `batch-recalc`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 11, 2020 at 09:50 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=h1) Report
> Merging [#110](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/cc2ee54e5547c8f0303432ef827f11ee7e3c6912?src=pr&el=desc) will **increase** coverage by `0.21%`.
> The diff coverage is `86.36%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #110      +/-   ##
==========================================
+ Coverage   64.13%   64.35%   +0.21%     
==========================================
  Files          41       41              
  Lines        1592     1599       +7     
==========================================
+ Hits         1021     1029       +8     
+ Misses        450      449       -1     
  Partials      121      121
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `75% <ø> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `68.75% <100%> (ø)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `72.22% <75%> (+5.55%)` | :arrow_up: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `83.33% <92.3%> (+8.33%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=footer). Last update [cc2ee54...cf62d4b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/110?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on February 17, 2020 at 03:03 PM UTC

@semtexzv there are some conflicts

---

## Reviews

### Review by @josef-hak - Changes Requested on February 11, 2020 at 12:30 PM UTC

### Review by @josef-hak - Changes Requested on February 12, 2020 at 03:25 PM UTC

### Review by @josef-hak - Changes Requested on February 17, 2020 at 03:04 PM UTC

conflicts

### Review by @josef-hak - Approved on February 18, 2020 at 12:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/110*
