---
type: pull_request
number: 152
title: "Improved testing log hook, added uploadHandler tests"
state: merged
author: josef-hak
created: 2020-02-27T14:57:13Z
updated: 2020-03-09T11:46:17Z
closed: 2020-03-02T12:01:10Z
merged: 2020-03-02T12:01:10Z
base_branch: master
head_branch: rm_duplicite_checking
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/152
---

# Pull Request #152: Improved testing log hook, added uploadHandler tests

**Author**: @josef-hak
**Created**: February 27, 2020 at 02:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rm_duplicite_checking`

## Description

- added uploadHandler tests for error states

---

## Discussion

### Comment by @codecov-io on February 27, 2020 at 03:06 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=h1) Report
> Merging [#152](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e77384b3907426b4f7b189876f535a652d621eaa?src=pr&el=desc) will **increase** coverage by `1.35%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #152      +/-   ##
==========================================
+ Coverage   64.94%   66.29%   +1.35%     
==========================================
  Files          42       42              
  Lines        1626     1629       +3     
==========================================
+ Hits         1056     1080      +24     
+ Misses        450      435      -15     
+ Partials      120      114       -6
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `85.71% <100%> (+15.34%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `62.5% <100%> (+17.04%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=footer). Last update [e77384b...d093ad6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/152?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on March 02, 2020 at 12:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/152*
