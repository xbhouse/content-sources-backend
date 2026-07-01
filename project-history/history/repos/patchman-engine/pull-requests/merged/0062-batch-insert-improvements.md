---
type: pull_request
number: 62
title: "Batch insert improvements"
state: merged
author: semtexzv
created: 2020-01-15T13:28:28Z
updated: 2021-03-16T09:27:00Z
closed: 2020-01-24T08:37:34Z
merged: 2020-01-24T08:37:34Z
base_branch: master
head_branch: reflect-batch
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/62
---

# Pull Request #62: Batch insert improvements

**Author**: @semtexzv
**Created**: January 15, 2020 at 01:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `reflect-batch`

## Description

Implements 2 changes to batch insert code.
- No longer require `[]interface{}`, allowing user to pass in any slice as an input
- Replicate `Create` behavior of scanning inserted values back into the input slice, reducing the need to perform second query to load just inserted values

---

## Discussion

### Comment by @codecov-io on January 23, 2020 at 12:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=h1) Report
> Merging [#62](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/7f846e224be89cdad3824070dd4f8fb79561b8e1?src=pr&el=desc) will **decrease** coverage by `0.34%`.
> The diff coverage is `80%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #62      +/-   ##
==========================================
- Coverage   57.92%   57.57%   -0.35%     
==========================================
  Files          30       29       -1     
  Lines        1212     1221       +9     
==========================================
+ Hits          702      703       +1     
- Misses        432      436       +4     
- Partials       78       82       +4
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `35.89% <100%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `74.82% <100%> (ø)` | :arrow_up: |
| [base/database/batch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9iYXRjaC5nbw==) | `65.28% <78.04%> (-2.94%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=footer). Last update [7f846e2...1a8fc81](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/62?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 17, 2020 at 12:35 PM UTC

Good simplification, just a few requests.

### Review by @semtexzv - Commented on January 20, 2020 at 09:40 AM UTC

### Review by @semtexzv - Commented on January 20, 2020 at 09:40 AM UTC

### Review by @josef-hak - Changes Requested on January 23, 2020 at 08:33 AM UTC

I see some confilts here.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/62*
