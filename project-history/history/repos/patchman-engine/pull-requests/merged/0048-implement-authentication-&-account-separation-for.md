---
type: pull_request
number: 48
title: "Implement authentication & account separation for hosts"
state: merged
author: semtexzv
created: 2020-01-07T14:06:01Z
updated: 2021-03-16T09:26:50Z
closed: 2020-01-14T08:36:01Z
merged: 2020-01-14T08:36:01Z
base_branch: master
head_branch: auth-acccounts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/48
---

# Pull Request #48: Implement authentication & account separation for hosts

**Author**: @semtexzv
**Created**: January 07, 2020 at 02:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `auth-acccounts`

## Description

This PR adds middleware for parsing `x-rh-identity` header, and SQL clauses to restrict 
systems returned based on account_id provided in the identity header.

---

## Discussion

### Comment by @codecov-io on January 07, 2020 at 02:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=h1) Report
> Merging [#48](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ff6dc5ec4ec07879349dbc2feec90eaa385a8cce?src=pr&el=desc) will **increase** coverage by `2.31%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #48      +/-   ##
==========================================
+ Coverage    56.2%   58.51%   +2.31%     
==========================================
  Files          28       27       -1     
  Lines         943      945       +2     
==========================================
+ Hits          530      553      +23     
+ Misses        355      334      -21     
  Partials       58       58
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `70.58% <100%> (+3.19%)` | :arrow_up: |
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100% <100%> (ø)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `90.62% <100%> (+1.33%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85% <100%> (+1.07%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.19% <100%> (+1.19%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `0% <0%> (ø)` | :arrow_up: |
| [base/models/models.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-YmFzZS9tb2RlbHMvbW9kZWxzLmdv) | | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | | |
| [listener/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZhbHVhdGUuZ28=) | `0% <0%> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=footer). Last update [ff6dc5e...3ea24c6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/48?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on January 09, 2020 at 09:07 AM UTC

@semtexzv please, solve these conflicts

---

## Reviews

### Review by @josef-hak - Changes Requested on January 13, 2020 at 09:35 AM UTC

Cool, thanks for update. I've added only couple of ideas how to improve it.

### Review by @semtexzv - Commented on January 13, 2020 at 01:22 PM UTC

### Review by @semtexzv - Commented on January 13, 2020 at 01:22 PM UTC

### Review by @semtexzv - Commented on January 13, 2020 at 01:23 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/48*
