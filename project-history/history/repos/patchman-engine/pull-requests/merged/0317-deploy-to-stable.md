---
type: pull_request
number: 317
title: "Deploy to stable"
state: merged
author: semtexzv
created: 2020-08-28T11:51:48Z
updated: 2020-08-28T12:48:33Z
closed: 2020-08-28T12:48:33Z
merged: 2020-08-28T12:48:33Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/317
---

# Pull Request #317: Deploy to stable

**Author**: @semtexzv
**Created**: August 28, 2020 at 11:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

- Contains mostly fixes to packages endpoint

---

## Discussion

### Comment by @codecov-commenter on August 28, 2020 at 11:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=h1) Report
> Merging [#317](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/56f040a828c66977c25913e52837d26550e40c73?el=desc) will **decrease** coverage by `0.00%`.
> The diff coverage is `76.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #317      +/-   ##
==========================================
- Coverage   61.83%   61.83%   -0.01%     
==========================================
  Files          52       53       +1     
  Lines        2602     2615      +13     
==========================================
+ Hits         1609     1617       +8     
- Misses        787      793       +6     
+ Partials      206      205       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.14% <ø> (-0.19%)` | :arrow_down: |
| [manager/controllers/status.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zdGF0dXMuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `54.54% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `75.38% <77.77%> (-0.81%)` | :arrow_down: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `73.91% <81.81%> (+11.41%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.16% <100.00%> (+1.41%)` | :arrow_up: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `46.83% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `77.08% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `62.50% <100.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=footer). Last update [56f040a...7fafda7](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/317?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/317*
