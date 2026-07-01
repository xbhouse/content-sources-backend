---
type: pull_request
number: 75
title: "Fix CI deployment"
state: closed
author: semtexzv
created: 2020-01-27T12:45:34Z
updated: 2021-03-16T09:27:01Z
closed: 2020-01-27T13:00:38Z
base_branch: master
head_branch: ci-deploy
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/75
---

# Pull Request #75: Fix CI deployment

**Author**: @semtexzv
**Created**: January 27, 2020 at 12:45 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `ci-deploy`

## Description

Fix missing configuration options on which CI deployment was crashing.
Limit number of listener pods (Incorrect resource limits, resolving in progress). 

This is just a simple fix for resolving current issues on CI. Actual fix which adresses listener performance & resource limits is in progress.

---

## Discussion

### Comment by @codecov-io on January 27, 2020 at 12:54 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@12ba442`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #75   +/-   ##
=========================================
  Coverage          ?   56.25%           
=========================================
  Files             ?       32           
  Lines             ?     1303           
  Branches          ?        0           
=========================================
  Hits              ?      733           
  Misses            ?      478           
  Partials          ?       92
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75?src=pr&el=footer). Last update [12ba442...4c15de5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/75?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/75*
