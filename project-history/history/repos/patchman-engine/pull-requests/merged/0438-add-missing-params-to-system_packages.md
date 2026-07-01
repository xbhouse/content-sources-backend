---
type: pull_request
number: 438
title: "Add missing params to system_packages"
state: merged
author: semtexzv
created: 2020-12-04T08:06:44Z
updated: 2021-03-16T09:30:42Z
closed: 2020-12-07T13:29:16Z
merged: 2020-12-07T13:29:16Z
base_branch: master
head_branch: fix-params
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/438
---

# Pull Request #438: Add missing params to system_packages

**Author**: @semtexzv
**Created**: December 04, 2020 at 08:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-params`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on December 04, 2020 at 08:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=h1) Report
> Merging [#438](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=desc) (c18f37b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/21cf1e007189b5f252be8a032afc117a5764478c?el=desc) (21cf1e0) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #438      +/-   ##
==========================================
+ Coverage   62.68%   62.70%   +0.01%     
==========================================
  Files          60       60              
  Lines        2731     2732       +1     
==========================================
+ Hits         1712     1713       +1     
  Misses        779      779              
  Partials      240      240              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `77.50% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `65.78% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.01% <100.00%> (+0.11%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=footer). Last update [21cf1e0...c18f37b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/438?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on December 04, 2020 at 08:10 AM UTC

I see the same problem in package_systems. There are probably more missing params as the annotation is pretty short in this handler.

### Review by @josef-hak - Approved on December 07, 2020 at 01:29 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/438*
