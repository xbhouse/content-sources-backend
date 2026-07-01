---
type: pull_request
number: 165
title: "Optimize listener message handling"
state: merged
author: semtexzv
created: 2020-03-11T12:47:10Z
updated: 2021-03-16T09:27:41Z
closed: 2020-03-13T11:38:46Z
merged: 2020-03-13T11:38:46Z
base_branch: master
head_branch: listener-tx
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/165
---

# Pull Request #165: Optimize listener message handling

**Author**: @semtexzv
**Created**: March 11, 2020 at 12:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-tx`

## Description

- Wrap code in single transaction, previously it would run in 2 transactions
- Skip sending eval messages for systems, which haven't had their `system_profile` changed.
- Skip re-writing `vmaas_json` if it haven't changed

---

## Discussion

### Comment by @codecov-io on March 12, 2020 at 09:20 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=h1) Report
> Merging [#165](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b0b9c52b614228ee4d9fae8b4c83a4d5caa97592&el=desc) will **increase** coverage by `0.20%`.
> The diff coverage is `83.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #165      +/-   ##
==========================================
+ Coverage   64.99%   65.19%   +0.20%     
==========================================
  Files          44       44              
  Lines        1757     1770      +13     
==========================================
+ Hits         1142     1154      +12     
- Misses        485      486       +1     
  Partials      130      130              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `76.60% <83.33%> (+1.29%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=footer). Last update [c756318...091e4b1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/165?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on March 12, 2020 at 12:14 PM UTC

Nice improvement, thanks. I would add test for that as mentioned, but it's up to you. Feel free to merge it then.

### Review by @josef-hak - Approved on March 13, 2020 at 11:38 AM UTC

Great, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/165*
