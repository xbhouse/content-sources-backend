---
type: pull_request
number: 170
title: "Split database migration from database container"
state: merged
author: josef-hak
created: 2020-03-13T18:20:08Z
updated: 2020-03-24T09:26:06Z
closed: 2020-03-18T16:25:38Z
merged: 2020-03-18T16:25:38Z
base_branch: master
head_branch: db_migrate
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/170
---

# Pull Request #170: Split database migration from database container

**Author**: @josef-hak
**Created**: March 13, 2020 at 06:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `db_migrate`

## Description

- Created new db_admin container for migration
- Updated docker-compose.test.yml
- Updated docker-compose.yml

---

## Discussion

### Comment by @codecov-io on March 13, 2020 at 06:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=h1) Report
> Merging [#170](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d1c8c5f616345998e9ac5e7836d0989cafe5e9ee&el=desc) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #170      +/-   ##
==========================================
+ Coverage   65.14%   65.16%   +0.01%     
==========================================
  Files          44       44              
  Lines        1779     1780       +1     
==========================================
+ Hits         1159     1160       +1     
  Misses        488      488              
  Partials      132      132              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `84.21% <100.00%> (+0.42%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=footer). Last update [d1c8c5f...aed5a44](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/170?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on March 18, 2020 at 07:25 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/170*
