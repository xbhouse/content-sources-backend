---
type: pull_request
number: 441
title: "Spec check workflow"
state: closed
author: digitronik
created: 2020-12-09T11:58:48Z
updated: 2020-12-11T08:15:30Z
closed: 2020-12-11T08:15:29Z
base_branch: master
head_branch: spec_check_workflow
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/441
---

# Pull Request #441: Spec check workflow

**Author**: @digitronik
**Created**: December 09, 2020 at 11:58 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `spec_check_workflow`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on December 10, 2020 at 08:01 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=h1) Report
> Merging [#441](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=desc) (3eb876f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf09edb207daaf77f375078e464a21b97a441987?el=desc) (bf09edb) will **increase** coverage by `0.76%`.
> The diff coverage is `81.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #441      +/-   ##
==========================================
+ Coverage   61.91%   62.68%   +0.76%     
==========================================
  Files          59       60       +1     
  Lines        2626     2731     +105     
==========================================
+ Hits         1626     1712      +86     
- Misses        766      779      +13     
- Partials      234      240       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.86% <0.00%> (-0.39%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `50.00% <50.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `33.33% <60.00%> (+33.33%)` | :arrow_up: |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `72.22% <70.37%> (-2.78%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.89% <71.87%> (-0.60%)` | :arrow_down: |
| [base/utils/identity.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9pZGVudGl0eS5nbw==) | `50.00% <100.00%> (+6.25%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.40% <100.00%> (ø)` | |
| ... and [19 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=footer). Last update [bf09edb...3eb876f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/441?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @digitronik on December 11, 2020 at 08:15 AM UTC

lots of conflicts  :) 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/441*
