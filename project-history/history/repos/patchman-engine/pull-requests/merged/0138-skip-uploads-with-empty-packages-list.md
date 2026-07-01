---
type: pull_request
number: 138
title: "Skip uploads with empty packages list"
state: merged
author: josef-hak
created: 2020-02-18T14:37:49Z
updated: 2020-02-20T12:42:35Z
closed: 2020-02-19T14:25:58Z
merged: 2020-02-19T14:25:58Z
base_branch: master
head_branch: skip_nopkgs_upload
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/138
---

# Pull Request #138: Skip uploads with empty packages list

**Author**: @josef-hak
**Created**: February 18, 2020 at 02:37 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `skip_nopkgs_upload`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 18, 2020 at 02:51 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=h1) Report
> Merging [#138](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eb4f5c118991b1bf33ad181816da6f0edcad73d6?src=pr&el=desc) will **increase** coverage by `0.15%`.
> The diff coverage is `75%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #138      +/-   ##
==========================================
+ Coverage   64.46%   64.61%   +0.15%     
==========================================
  Files          41       41              
  Lines        1604     1611       +7     
==========================================
+ Hits         1034     1041       +7     
- Misses        449      450       +1     
+ Partials      121      120       -1
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.5% <ø> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `70.37% <75%> (+1.62%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=footer). Last update [eb4f5c1...ae1f5d5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/138?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on February 19, 2020 at 02:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/138*
