---
type: pull_request
number: 215
title: "Ensured gracefull containers termination"
state: merged
author: josef-hak
created: 2020-04-07T16:01:57Z
updated: 2020-04-09T11:55:56Z
closed: 2020-04-09T11:42:31Z
merged: 2020-04-09T11:42:31Z
base_branch: master
head_branch: signals
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/215
---

# Pull Request #215: Ensured gracefull containers termination

**Author**: @josef-hak
**Created**: April 07, 2020 at 04:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `signals`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on April 08, 2020 at 01:28 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=h1) Report
> Merging [#215](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f0511854e00faec203e40a0069a6f9ccb544d21&el=desc) will **not change** coverage by `%`.
> The diff coverage is `60.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #215   +/-   ##
=======================================
  Coverage   63.02%   63.02%           
=======================================
  Files          46       46           
  Lines        1993     1993           
=======================================
  Hits         1256     1256           
  Misses        590      590           
  Partials      147      147           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.76% <100.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `75.55% <100.00%> (ø)` | |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `54.54% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=footer). Last update [1f05118...221234f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/215?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on April 08, 2020 at 09:36 AM UTC

### Review by @josef-hak - Commented on April 08, 2020 at 10:06 AM UTC

### Review by @semtexzv - Commented on April 09, 2020 at 09:43 AM UTC

Great changes, but there are still some spaces when we want to derive context from the global one in order to properly cancel operations when receiving signal

### Review by @josef-hak - Commented on April 09, 2020 at 10:06 AM UTC

### Review by @josef-hak - Commented on April 09, 2020 at 10:08 AM UTC

### Review by @josef-hak - Commented on April 09, 2020 at 10:08 AM UTC

### Review by @semtexzv - Approved on April 09, 2020 at 11:35 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/215*
