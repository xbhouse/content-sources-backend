---
type: pull_request
number: 319
title: "Link between advisories and packages"
state: merged
author: semtexzv
created: 2020-09-07T13:39:28Z
updated: 2021-03-16T09:28:49Z
closed: 2020-09-10T06:37:58Z
merged: 2020-09-10T06:37:58Z
base_branch: master
head_branch: pkg-advisories-link
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/319
---

# Pull Request #319: Link between advisories and packages

**Author**: @semtexzv
**Created**: September 07, 2020 at 01:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-advisories-link`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 08, 2020 at 10:45 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=h1) Report
> Merging [#319](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/9b05081bf725b31c9512fef6db52ebcaa6558bd6?el=desc) will **increase** coverage by `0.37%`.
> The diff coverage is `77.77%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #319      +/-   ##
==========================================
+ Coverage   62.07%   62.44%   +0.37%     
==========================================
  Files          54       54              
  Lines        2642     2692      +50     
==========================================
+ Hits         1640     1681      +41     
- Misses        796      800       +4     
- Partials      206      211       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <ø> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.41% <70.00%> (+0.12%)` | :arrow_up: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `78.57% <100.00%> (+0.79%)` | :arrow_up: |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `81.42% <100.00%> (+1.12%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `75.00% <0.00%> (-0.39%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `88.23% <0.00%> (-0.10%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <0.00%> (+1.75%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=footer). Last update [9b05081...816320e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/319?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on September 08, 2020 at 11:25 AM UTC

Just two details to improve found. Good work.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/319*
