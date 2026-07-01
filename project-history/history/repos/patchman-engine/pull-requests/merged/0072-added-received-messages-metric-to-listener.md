---
type: pull_request
number: 72
title: "Added received messages metric to listener"
state: merged
author: josef-hak
created: 2020-01-22T17:36:12Z
updated: 2020-01-30T17:25:02Z
closed: 2020-01-28T08:50:45Z
merged: 2020-01-28T08:50:45Z
base_branch: master
head_branch: listener_metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/72
---

# Pull Request #72: Added received messages metric to listener

**Author**: @josef-hak
**Created**: January 22, 2020 at 05:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener_metrics`

## Description

- received messages metric
- evaluation duration metric
- updates count metric, etc

---

## Discussion

### Comment by @codecov-io on January 23, 2020 at 11:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@12ba442`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `32.07%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #72   +/-   ##
=========================================
  Coverage          ?   56.21%           
=========================================
  Files             ?       34           
  Lines             ?     1327           
  Branches          ?        0           
=========================================
  Hits              ?      746           
  Misses            ?      486           
  Partials          ?       95
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `60% <0%> (ø)` | |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `100% <100%> (ø)` | |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `29.41% <11.11%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `22.22% <22.22%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `67.02% <29.41%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `75.81% <46.66%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=footer). Last update [12ba442...4fafa26](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/72?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on January 24, 2020 at 09:35 AM UTC

A few issues

### Review by @josef-hak - Commented on January 27, 2020 at 12:04 PM UTC

### Review by @josef-hak - Commented on January 27, 2020 at 12:25 PM UTC

### Review by @semtexzv - Approved on January 28, 2020 at 08:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/72*
