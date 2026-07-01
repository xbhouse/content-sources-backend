---
type: pull_request
number: 396
title: "Fix tags"
state: merged
author: semtexzv
created: 2020-10-20T09:19:48Z
updated: 2021-03-16T09:29:48Z
closed: 2020-10-22T08:53:36Z
merged: 2020-10-22T08:53:36Z
base_branch: master
head_branch: fix-tags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/396
---

# Pull Request #396: Fix tags

**Author**: @semtexzv
**Created**: October 20, 2020 at 09:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-tags`

## Description

We didn't properly check, whether we really parsed a tag

---

## Discussion

### Comment by @codecov-io on October 20, 2020 at 09:29 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=h1) Report
> Merging [#396](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b146efd195079fa185b1f0856295ffa30d1e7ead?el=desc) will **increase** coverage by `0.30%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #396      +/-   ##
==========================================
+ Coverage   61.60%   61.91%   +0.30%     
==========================================
  Files          58       58              
  Lines        2602     2602              
==========================================
+ Hits         1603     1611       +8     
+ Misses        764      759       -5     
+ Partials      235      232       -3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.65% <100.00%> (+10.34%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.77% <100.00%> (+0.12%)` | :arrow_up: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `63.88% <0.00%> (+5.55%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `78.57% <0.00%> (+9.60%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=footer). Last update [b146efd...94f72a6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/396?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on October 20, 2020 at 10:19 AM UTC

just recomm. to add test, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/396*
