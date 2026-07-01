---
type: pull_request
number: 162
title: "Increase evaluation performance"
state: merged
author: semtexzv
created: 2020-03-09T16:08:34Z
updated: 2020-03-10T09:51:31Z
closed: 2020-03-10T09:51:31Z
merged: 2020-03-10T09:51:31Z
base_branch: master
head_branch: eval-perf
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/162
---

# Pull Request #162: Increase evaluation performance

**Author**: @semtexzv
**Created**: March 09, 2020 at 04:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval-perf`

## Description

- Adds a missing private key to `advisory_account_data` (main increase)
- Removes `id` column from `system_advisories`, since it is a join table
- Remove batching from kafka writer
- Misc removal of pointer-to-array antipattern


---

## Discussion

### Comment by @codecov-io on March 10, 2020 at 07:44 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=h1) Report
> Merging [#162](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fa420493abfeba7c5a00d50ea4a2c5e99b8980d5?src=pr&el=desc) will **increase** coverage by `0.04%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #162      +/-   ##
==========================================
+ Coverage   65.08%   65.12%   +0.04%     
==========================================
  Files          43       43              
  Lines        1618     1620       +2     
==========================================
+ Hits         1053     1055       +2     
  Misses        452      452              
  Partials      113      113
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `78.04% <100%> (+0.54%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `76.31% <100%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.72% <100%> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79% <100%> (+0.21%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `91.89% <100%> (ø)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `88.88% <100%> (ø)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `78.57% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=footer). Last update [fa42049...401bf64](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/162?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on March 10, 2020 at 09:39 AM UTC

### Review by @semtexzv - Commented on March 10, 2020 at 09:47 AM UTC

### Review by @semtexzv - Commented on March 10, 2020 at 09:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/162*
