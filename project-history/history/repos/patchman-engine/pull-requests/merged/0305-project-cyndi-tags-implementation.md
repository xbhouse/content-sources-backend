---
type: pull_request
number: 305
title: "Project cyndi tags implementation"
state: merged
author: semtexzv
created: 2020-08-17T13:10:18Z
updated: 2021-03-16T09:26:09Z
closed: 2020-08-18T12:11:09Z
merged: 2020-08-18T12:11:09Z
base_branch: master
head_branch: cyndi-tags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/305
---

# Pull Request #305: Project cyndi tags implementation

**Author**: @semtexzv
**Created**: August 17, 2020 at 01:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `cyndi-tags`

## Description

- Change dev data to use UUIDS (allows for joining with inventory.hosts table)
- Use tags stored in inventory.hosts table for filtering

---

## Discussion

### Comment by @codecov-commenter on August 17, 2020 at 01:43 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=h1) Report
> Merging [#305](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d6ffc96ccbbc57a718a77009eb482116ed4690e0&el=desc) will **decrease** coverage by `0.12%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #305      +/-   ##
==========================================
- Coverage   61.98%   61.85%   -0.13%     
==========================================
  Files          52       52              
  Lines        2612     2593      -19     
==========================================
- Hits         1619     1604      -15     
+ Misses        787      785       -2     
+ Partials      206      204       -2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.14% <ø> (-0.19%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `77.08% <100.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.56% <100.00%> (+0.37%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=footer). Last update [0118ec9...3514460](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/305?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/305*
