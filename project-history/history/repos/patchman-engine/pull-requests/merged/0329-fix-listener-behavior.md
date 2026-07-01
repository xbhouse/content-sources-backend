---
type: pull_request
number: 329
title: "Fix listener behavior"
state: merged
author: semtexzv
created: 2020-09-10T08:28:40Z
updated: 2021-03-16T09:28:51Z
closed: 2020-09-10T10:14:00Z
merged: 2020-09-10T10:14:00Z
base_branch: master
head_branch: fix-listener
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/329
---

# Pull Request #329: Fix listener behavior

**Author**: @semtexzv
**Created**: September 10, 2020 at 08:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-listener`

## Description

- Process uploads on update events.
- Cleanup the listener code


---

## Discussion

### Comment by @codecov-commenter on September 10, 2020 at 09:16 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=h1) Report
> Merging [#329](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d34e643f79854d50152b49a26cdcd704324fc62f?el=desc) will **increase** coverage by `0.41%`.
> The diff coverage is `73.91%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #329      +/-   ##
==========================================
+ Coverage   62.39%   62.81%   +0.41%     
==========================================
  Files          54       54              
  Lines        2694     2681      -13     
==========================================
+ Hits         1681     1684       +3     
+ Misses        801      789      -12     
+ Partials      212      208       -4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `50.87% <0.00%> (+4.04%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.12% <75.55%> (+1.97%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=footer). Last update [d34e643...287e02f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/329?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on September 10, 2020 at 09:57 AM UTC

Uh, I did not expect so many updates. But guessing it's code cleaning. Ok.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/329*
