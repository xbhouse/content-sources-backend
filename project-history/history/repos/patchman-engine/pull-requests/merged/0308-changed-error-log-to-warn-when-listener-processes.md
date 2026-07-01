---
type: pull_request
number: 308
title: "changed error log to warn when listener processes unknonw msg type"
state: merged
author: josef-hak
created: 2020-08-19T09:04:04Z
updated: 2020-08-19T11:48:40Z
closed: 2020-08-19T10:23:33Z
merged: 2020-08-19T10:23:33Z
base_branch: master
head_branch: error-to-warn
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/308
---

# Pull Request #308: changed error log to warn when listener processes unknonw msg type

**Author**: @josef-hak
**Created**: August 19, 2020 at 09:04 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `error-to-warn`

## Description

also fixed test after changing inventory_ids to uuids


---

## Discussion

### Comment by @codecov-commenter on August 19, 2020 at 10:22 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=h1) Report
> Merging [#308](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d6ffc96ccbbc57a718a77009eb482116ed4690e0&el=desc) will **decrease** coverage by `0.12%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #308      +/-   ##
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


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.14% <ø> (-0.19%)` | :arrow_down: |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `54.54% <ø> (ø)` | |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `46.83% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `77.08% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `62.50% <100.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.56% <100.00%> (+0.37%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=footer). Last update [dc8b364...1b2edb0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/308?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on August 19, 2020 at 10:02 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/308*
