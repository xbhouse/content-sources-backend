---
type: pull_request
number: 316
title: "Added \"updatable\" package list filter"
state: merged
author: josef-hak
created: 2020-08-27T13:28:25Z
updated: 2020-09-08T10:02:29Z
closed: 2020-09-08T09:14:39Z
merged: 2020-09-08T09:14:39Z
base_branch: master
head_branch: pkg-filter-upgradable
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/316
---

# Pull Request #316: Added "updatable" package list filter

**Author**: @josef-hak
**Created**: August 27, 2020 at 01:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-filter-upgradable`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 27, 2020 at 02:38 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=h1) Report
> Merging [#316](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/6687d81e32ba516fdb45e307f59a87c51ad4e4c2?el=desc) will **increase** coverage by `0.07%`.
> The diff coverage is `25.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #316      +/-   ##
==========================================
+ Coverage   61.99%   62.07%   +0.07%     
==========================================
  Files          54       54              
  Lines        2642     2642              
==========================================
+ Hits         1638     1640       +2     
+ Misses        797      796       -1     
+ Partials      207      206       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `77.77% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `67.50% <ø> (+5.00%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <ø> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `14.28% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=footer). Last update [6687d81...bee0c38](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/316?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on September 08, 2020 at 07:26 AM UTC

@semtexzv you were right, done.

---

## Reviews

### Review by @semtexzv - Commented on September 07, 2020 at 03:26 PM UTC

We can implement the same thing just by adding 
```
Updatable bool  `json:"updatable" query:"json_array_length(spkg.update_data::json) > 0"`
```
to `SystemPackagesAttrs`

### Review by @semtexzv - Commented on September 08, 2020 at 07:35 AM UTC

### Review by @josef-hak - Commented on September 08, 2020 at 07:39 AM UTC

### Review by @semtexzv - Approved on September 08, 2020 at 09:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/316*
