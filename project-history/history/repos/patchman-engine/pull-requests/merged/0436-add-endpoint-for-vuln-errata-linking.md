---
type: pull_request
number: 436
title: "Add endpoint for vuln errata linking"
state: merged
author: semtexzv
created: 2020-12-03T10:00:36Z
updated: 2021-03-16T09:30:38Z
closed: 2021-01-08T13:24:42Z
merged: 2021-01-08T13:24:42Z
base_branch: master
head_branch: sys-adv-link
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/436
---

# Pull Request #436: Add endpoint for vuln errata linking

**Author**: @semtexzv
**Created**: December 03, 2020 at 10:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sys-adv-link`

## Description

Add new endpoints used by vulnerabilitites.
These endpoints:
- Don't support filtering, paging
- Provide just cross-linking between systems and advisories

---

## Discussion

### Comment by @codecov-io on December 03, 2020 at 11:07 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=h1) Report
> Merging [#436](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=desc) (a32c4e0) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf09edb207daaf77f375078e464a21b97a441987?el=desc) (bf09edb) will **increase** coverage by `0.90%`.
> The diff coverage is `80.08%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #436      +/-   ##
==========================================
+ Coverage   61.91%   62.82%   +0.90%     
==========================================
  Files          59       61       +2     
  Lines        2626     2752     +126     
==========================================
+ Hits         1626     1729     +103     
- Misses        766      781      +15     
- Partials      234      242       +8     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.86% <0.00%> (-0.39%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `50.00% <50.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `33.33% <54.54%> (+33.33%)` | :arrow_up: |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `72.22% <70.37%> (-2.78%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.89% <71.87%> (-0.60%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <75.00%> (+0.70%)` | :arrow_up: |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `80.95% <80.95%> (ø)` | |
| ... and [21 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=footer). Last update [bf09edb...a32c4e0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/436?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on December 03, 2020 at 11:10 AM UTC

Lgtm in general, just two comments or questions.

### Review by @semtexzv - Commented on January 08, 2021 at 11:53 AM UTC

### Review by @josef-hak - Commented on January 08, 2021 at 12:56 PM UTC

### Review by @semtexzv - Commented on January 08, 2021 at 01:02 PM UTC

### Review by @josef-hak - Approved on January 08, 2021 at 01:07 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/436*
