---
type: pull_request
number: 457
title: "schema migration - partitioning"
state: merged
author: MichaelMraka
created: 2021-01-06T17:50:01Z
updated: 2021-04-16T11:24:23Z
closed: 2021-01-15T10:18:17Z
merged: 2021-01-15T10:18:17Z
base_branch: master
head_branch: pr8
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/457
---

# Pull Request #457: schema migration - partitioning

**Author**: @MichaelMraka
**Created**: January 06, 2021 at 05:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr8`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 06, 2021 at 06:23 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=h1) Report
> Merging [#457](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=desc) (b06c322) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b8e7872205fcf02b4dd8d6e46bc18526bdf96057?el=desc) (b8e7872) will **increase** coverage by `0.01%`.
> The diff coverage is `90.90%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #457      +/-   ##
==========================================
+ Coverage   62.72%   62.74%   +0.01%     
==========================================
  Files          61       61              
  Lines        2785     2786       +1     
==========================================
+ Hits         1747     1748       +1     
  Misses        794      794              
  Partials      244      244              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.40% <100.00%> (ø)` | |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.76% <100.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.23% <100.00%> (+0.11%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.36% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.08% <100.00%> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `77.50% <100.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `91.17% <100.00%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <100.00%> (ø)` | |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=footer). Last update [b8e7872...b06c322](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/457?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on January 14, 2021 at 11:28 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/457*
