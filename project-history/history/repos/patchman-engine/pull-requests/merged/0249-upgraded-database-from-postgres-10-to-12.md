---
type: pull_request
number: 249
title: "Upgraded database from postgres 10 to 12"
state: merged
author: josef-hak
created: 2020-05-05T08:50:48Z
updated: 2020-09-30T06:56:10Z
closed: 2020-09-24T14:09:54Z
merged: 2020-09-24T14:09:54Z
base_branch: master
head_branch: db-pg-12
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/249
---

# Pull Request #249: Upgraded database from postgres 10 to 12

**Author**: @josef-hak
**Created**: May 05, 2020 at 08:50 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `db-pg-12`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 16, 2020 at 09:17 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=h1) Report
> Merging [#249](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f420be6ccf3c3b16c0485cec9271bfac58c36f73?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #249   +/-   ##
=======================================
  Coverage   63.32%   63.32%           
=======================================
  Files          54       54           
  Lines        2686     2686           
=======================================
  Hits         1701     1701           
  Misses        773      773           
  Partials      212      212           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=footer). Last update [f420be6...456709f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/249?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on September 23, 2020 at 11:12 AM UTC

Did you try to upgrade pre-existing database ? tests run on fresh database every time, so we don't know whether the CI/QA envs will be upgraded correctly

### Comment by @josef-hak on September 23, 2020 at 11:58 AM UTC

@semtexzv yes, I tried that now. We will need to run ugraded db container with POSTGRESQL_UPGRADE=copy (https://hub.docker.com/r/centos/postgresql-12-centos7). I tried it in docker with exisiting volume ver10 and upgrade to version 12.

### Comment by @semtexzv on September 24, 2020 at 10:21 AM UTC

Can be merged after https://github.com/RedHatInsights/e2e-deploy/pull/2287


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/249*
