---
type: pull_request
number: 46
title: "Implemented evaluation core function"
state: merged
author: josef-hak
created: 2020-01-06T15:54:34Z
updated: 2020-01-17T12:33:25Z
closed: 2020-01-10T08:28:57Z
merged: 2020-01-10T08:28:57Z
base_branch: master
head_branch: evaluate
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/46
---

# Pull Request #46: Implemented evaluation core function

**Author**: @josef-hak
**Created**: January 06, 2020 at 03:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `evaluate`

## Description

- Created evaluator module
- Prepared methods for getting data to process evaluation
- Ensured running tests sequentially
- Implemented main evaluation method

---

## Discussion

### Comment by @codecov-io on January 07, 2020 at 11:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@a7752ef`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `45%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=tree)

```diff
@@           Coverage Diff            @@
##             master     #46   +/-   ##
========================================
  Coverage          ?   56.2%           
========================================
  Files             ?      28           
  Lines             ?     943           
  Branches          ?       0           
========================================
  Hits              ?     530           
  Misses            ?     355           
  Partials          ?      58
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `0% <ø> (ø)` | |
| [base/models/models.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/diff?src=pr&el=tree#diff-YmFzZS9tb2RlbHMvbW9kZWxzLmdv) | `15.38% <ø> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `49.09% <ø> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0% <0%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `67.39% <100%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=footer). Last update [a7752ef...ef62948](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/46?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @beav - Commented on January 09, 2020 at 01:10 PM UTC

### Review by @semtexzv - Commented on January 09, 2020 at 10:16 AM UTC

### Review by @josef-hak - Commented on January 09, 2020 at 03:06 PM UTC

### Review by @josef-hak - Commented on January 09, 2020 at 03:19 PM UTC

### Review by @josef-hak - Commented on January 09, 2020 at 03:21 PM UTC

### Review by @semtexzv - Approved on January 10, 2020 at 08:28 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/46*
