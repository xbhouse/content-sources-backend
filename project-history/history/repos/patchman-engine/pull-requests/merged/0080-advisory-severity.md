---
type: pull_request
number: 80
title: "Advisory severity"
state: merged
author: semtexzv
created: 2020-01-29T10:27:16Z
updated: 2021-03-16T09:27:03Z
closed: 2020-01-30T09:44:55Z
merged: 2020-01-30T09:44:55Z
base_branch: master
head_branch: advisory_severity
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/80
---

# Pull Request #80: Advisory severity

**Author**: @semtexzv
**Created**: January 29, 2020 at 10:27 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisory_severity`

## Description

Sync advisory severity from vmaas.
- After internal discussion interprets `None` as null. Provides only index in the frontend, since strings will always be translated.

---

## Discussion

### Comment by @codecov-io on January 29, 2020 at 10:36 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=h1) Report
> Merging [#80](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/25f0c2d3bc2ed46efc07aa24d54899a2fa57bae1?src=pr&el=desc) will **increase** coverage by `0.22%`.
> The diff coverage is `77.77%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #80      +/-   ##
==========================================
+ Coverage   57.49%   57.72%   +0.22%     
==========================================
  Files          34       34              
  Lines        1327     1353      +26     
==========================================
+ Hits          763      781      +18     
- Misses        474      478       +4     
- Partials       90       94       +4
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.56% <100%> (ø)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `73.68% <100%> (+2.53%)` | :arrow_up: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `90.32% <100%> (ø)` | :arrow_up: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `41.41% <68%> (+5.51%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=footer). Last update [25f0c2d...9eb165c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/80?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 30, 2020 at 08:11 AM UTC

Just a few notes/questions.

### Review by @semtexzv - Commented on January 30, 2020 at 08:39 AM UTC

### Review by @semtexzv - Commented on January 30, 2020 at 08:39 AM UTC

### Review by @josef-hak - Approved on January 30, 2020 at 09:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/80*
