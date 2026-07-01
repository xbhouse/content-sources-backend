---
type: pull_request
number: 385
title: "system_platform and system_advisories partitioning"
state: merged
author: MichaelMraka
created: 2020-10-13T14:34:52Z
updated: 2020-11-09T08:44:33Z
closed: 2020-11-09T08:44:33Z
merged: 2020-11-09T08:44:33Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/385
---

# Pull Request #385: system_platform and system_advisories partitioning

**Author**: @MichaelMraka
**Created**: October 13, 2020 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on October 28, 2020 at 06:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=h1) Report
> Merging [#385](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf09edb207daaf77f375078e464a21b97a441987?el=desc) will **increase** coverage by `0.01%`.
> The diff coverage is `87.50%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #385      +/-   ##
==========================================
+ Coverage   61.91%   61.93%   +0.01%     
==========================================
  Files          59       59              
  Lines        2626     2627       +1     
==========================================
+ Hits         1626     1627       +1     
  Misses        766      766              
  Partials      234      234              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.40% <100.00%> (ø)` | |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.76% <100.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.65% <100.00%> (+0.11%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <100.00%> (ø)` | |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `51.61% <100.00%> (ø)` | |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `89.28% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `63.88% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=footer). Last update [bf09edb...798fe04](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/385?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @MichaelMraka on October 30, 2020 at 11:03 AM UTC

rebased to v1.4.1

### Comment by @MichaelMraka on October 30, 2020 at 12:27 PM UTC

Oops, I've just found out we have to move sequence on new system_platform to where it was on old table before migration.

### Comment by @josef-hak on October 30, 2020 at 02:03 PM UTC

@MichaelMraka oh, it's seen that was something what was not covered by tests, hmmm :thinking: 

### Comment by @MichaelMraka on October 30, 2020 at 02:36 PM UTC

Just comparing sequences is not enough.
This could be found only if we had some pre-filled data in schema comparison/migration test.

---

## Reviews

### Review by @semtexzv - Commented on October 14, 2020 at 08:56 AM UTC

### Review by @josef-hak - Approved on October 30, 2020 at 11:59 AM UTC

Greate job @mim, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/385*
