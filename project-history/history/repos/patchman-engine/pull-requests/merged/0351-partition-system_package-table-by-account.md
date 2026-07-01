---
type: pull_request
number: 351
title: "Partition system_package table by account"
state: merged
author: semtexzv
created: 2020-09-29T09:58:42Z
updated: 2021-03-16T09:28:58Z
closed: 2020-09-29T14:00:21Z
merged: 2020-09-29T14:00:21Z
base_branch: master
head_branch: sys-pkg-acc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/351
---

# Pull Request #351: Partition system_package table by account

**Author**: @semtexzv
**Created**: September 29, 2020 at 09:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sys-pkg-acc`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 29, 2020 at 10:41 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=h1) Report
> Merging [#351](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8b2f4004d5861751bcfe28658c564a1dc856e2e5?el=desc) will **increase** coverage by `0.04%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #351      +/-   ##
==========================================
+ Coverage   62.33%   62.37%   +0.04%     
==========================================
  Files          56       56              
  Lines        2506     2509       +3     
==========================================
+ Hits         1562     1565       +3     
  Misses        725      725              
  Partials      219      219              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.29% <100.00%> (+0.27%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `70.00% <100.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `59.45% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=footer). Last update [8b2f400...3bbe4dd](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/351?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on September 29, 2020 at 11:20 AM UTC

Looks good to me. Just a few questions/notes. @MichaelMraka, please check it as well. Also I don't see where the hashing and assigning to partition is done for new system_platform items.

### Review by @semtexzv - Commented on September 29, 2020 at 11:39 AM UTC

### Review by @semtexzv - Commented on September 29, 2020 at 11:40 AM UTC

### Review by @semtexzv - Commented on September 29, 2020 at 11:43 AM UTC

### Review by @MichaelMraka - Commented on September 29, 2020 at 11:43 AM UTC

### Review by @MichaelMraka - Commented on September 29, 2020 at 11:47 AM UTC

### Review by @semtexzv - Commented on September 29, 2020 at 12:26 PM UTC

### Review by @MichaelMraka - Commented on September 29, 2020 at 01:29 PM UTC

### Review by @MichaelMraka - Commented on September 29, 2020 at 01:29 PM UTC

### Review by @josef-hak - Approved on September 29, 2020 at 01:37 PM UTC

### Review by @MichaelMraka - Approved on September 29, 2020 at 01:56 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/351*
