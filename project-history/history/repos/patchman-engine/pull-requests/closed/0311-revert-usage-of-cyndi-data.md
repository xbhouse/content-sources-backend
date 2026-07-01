---
type: pull_request
number: 311
title: "Revert usage of cyndi data"
state: closed
author: semtexzv
created: 2020-08-24T13:48:48Z
updated: 2021-03-16T09:26:12Z
closed: 2020-08-26T15:03:32Z
base_branch: master
head_branch: revert-cyndi
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/311
---

# Pull Request #311: Revert usage of cyndi data

**Author**: @semtexzv
**Created**: August 24, 2020 at 01:48 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `revert-cyndi`

## Description

Revert usage of project cyndi data. Keep changes regarding primary keys.

---

## Discussion

### Comment by @codecov-commenter on August 24, 2020 at 02:06 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=h1) Report
> Merging [#311](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/06f1f4a0bf42675367147f40e97519c3bf8b3369?el=desc) will **increase** coverage by `0.15%`.
> The diff coverage is `86.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #311      +/-   ##
==========================================
+ Coverage   61.69%   61.84%   +0.15%     
==========================================
  Files          53       53              
  Lines        2600     2621      +21     
==========================================
+ Hits         1604     1621      +17     
- Misses        792      794       +2     
- Partials      204      206       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.32% <80.95%> (+0.18%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `77.08% <100.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.56% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=footer). Last update [06f1f4a...47a5319](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/311?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on August 26, 2020 at 03:03 PM UTC

Just disabling the current implementation for environments where not supported is enough. Closing.

---

## Reviews

### Review by @josef-hak - Changes Requested on August 25, 2020 at 08:05 AM UTC

### Review by @semtexzv - Commented on August 26, 2020 at 07:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/311*
