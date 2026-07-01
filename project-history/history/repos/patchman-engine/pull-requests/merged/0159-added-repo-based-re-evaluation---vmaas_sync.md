---
type: pull_request
number: 159
title: "Added repo-based re-evaluation - vmaas_sync"
state: merged
author: josef-hak
created: 2020-03-03T12:30:13Z
updated: 2020-03-12T12:22:17Z
closed: 2020-03-12T08:31:38Z
merged: 2020-03-12T08:31:38Z
base_branch: master
head_branch: repo_based_recalc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/159
---

# Pull Request #159: Added repo-based re-evaluation - vmaas_sync

**Author**: @josef-hak
**Created**: March 03, 2020 at 12:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `repo_based_recalc`

## Description

- Implemented repo-based recalculation in vmaas_sync.
- Added `ENABLE_REPO_BASED_RE_EVALUATION` env variable.
  - It's needed to be in vmaas_sync OpenShift config before merging.

---

## Discussion

### Comment by @codecov-io on March 10, 2020 at 02:15 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=h1) Report
> Merging [#159](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/daa3ed55367709983bf33d739bc0d76c38b03b67?src=pr&el=desc) will **decrease** coverage by `0.33%`.
> The diff coverage is `63.04%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #159      +/-   ##
==========================================
- Coverage   65.12%   64.78%   -0.34%     
==========================================
  Files          43       44       +1     
  Lines        1620     1701      +81     
==========================================
+ Hits         1055     1102      +47     
- Misses        452      477      +25     
- Partials      113      122       +9
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.97% <ø> (-2.6%)` | :arrow_down: |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `7.69% <0%> (-1.99%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `25% <100%> (+12.27%)` | :arrow_up: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `61.76% <30.76%> (-21.57%)` | :arrow_down: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `72.58% <72.58%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=footer). Last update [daa3ed5...0594236](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/159?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on March 10, 2020 at 09:53 AM UTC

### Review by @josef-hak - Commented on March 10, 2020 at 12:51 PM UTC

### Review by @josef-hak - Commented on March 10, 2020 at 12:55 PM UTC

### Review by @semtexzv - Commented on March 10, 2020 at 12:43 PM UTC

### Review by @semtexzv - Commented on March 10, 2020 at 01:17 PM UTC

### Review by @semtexzv - Commented on March 11, 2020 at 02:32 PM UTC

### Review by @semtexzv - Approved on March 12, 2020 at 08:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/159*
