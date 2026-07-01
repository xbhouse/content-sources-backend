---
type: pull_request
number: 247
title: "Handle 'updated' events"
state: merged
author: semtexzv
created: 2020-05-05T08:30:29Z
updated: 2021-03-16T09:28:17Z
closed: 2020-05-07T09:43:22Z
merged: 2020-05-07T09:43:22Z
base_branch: master
head_branch: handle-updates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/247
---

# Pull Request #247: Handle 'updated' events

**Author**: @semtexzv
**Created**: May 05, 2020 at 08:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `handle-updates`

## Description

Handle 'updated' messages from inventory to update display_name

---

## Discussion

### Comment by @codecov-io on May 06, 2020 at 08:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=h1) Report
> Merging [#247](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/81f008b622cfc2d775cf0d2c318e020c800e3dec&el=desc) will **decrease** coverage by `0.39%`.
> The diff coverage is `45.45%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #247      +/-   ##
==========================================
- Coverage   62.92%   62.52%   -0.40%     
==========================================
  Files          47       47              
  Lines        2109     2135      +26     
==========================================
+ Hits         1327     1335       +8     
- Misses        629      644      +15     
- Partials      153      156       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `54.68% <43.75%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `73.68% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=footer). Last update [81f008b...9fecc44](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/247?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on May 05, 2020 at 11:18 AM UTC

ok, seems good, just typo found

### Review by @josef-hak - Approved on May 07, 2020 at 09:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/247*
