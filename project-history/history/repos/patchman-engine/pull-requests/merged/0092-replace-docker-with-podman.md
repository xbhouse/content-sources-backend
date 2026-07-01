---
type: pull_request
number: 92
title: "Replace docker with podman"
state: merged
author: josef-hak
created: 2020-01-31T10:04:23Z
updated: 2020-02-03T07:00:30Z
closed: 2020-01-31T15:37:26Z
merged: 2020-01-31T15:37:26Z
base_branch: master
head_branch: podman
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/92
---

# Pull Request #92: Replace docker with podman

**Author**: @josef-hak
**Created**: January 31, 2020 at 10:04 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `podman`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 31, 2020 at 10:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=h1) Report
> Merging [#92](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b7f6eec7009dde5a49bbbbf466a09a2786c8b190?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=tree)

```diff
@@          Coverage Diff           @@
##           master     #92   +/-   ##
======================================
  Coverage    57.6%   57.6%           
======================================
  Files          34      34           
  Lines        1401    1401           
======================================
  Hits          807     807           
  Misses        497     497           
  Partials       97      97
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=footer). Last update [b7f6eec...5093f15](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/92?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @beav on January 31, 2020 at 03:31 PM UTC

I tested this with docker-compose and it is still compatible.

### Comment by @josef-hak on January 31, 2020 at 03:37 PM UTC

@beav yes, because podman-compose (until now) can only subset of features of docker-compose. But anyway I changed guides in README to podman. Usage of docker can be still supported but let's say "unofficially" :smile: 

---

## Reviews

### Review by @MichaelMraka - Approved on January 31, 2020 at 10:30 AM UTC

ack :)

### Review by @beav - Approved on January 31, 2020 at 03:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/92*
