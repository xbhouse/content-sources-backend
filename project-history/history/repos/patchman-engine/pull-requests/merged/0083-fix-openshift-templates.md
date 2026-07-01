---
type: pull_request
number: 83
title: "Fix openshift templates"
state: merged
author: semtexzv
created: 2020-01-29T16:29:59Z
updated: 2021-03-16T09:27:05Z
closed: 2020-01-30T07:10:51Z
merged: 2020-01-30T07:10:51Z
base_branch: master
head_branch: templates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/83
---

# Pull Request #83: Fix openshift templates

**Author**: @semtexzv
**Created**: January 29, 2020 at 04:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `templates`

## Description

- Env variables must be string
- missed default consumer count from listener

---

## Discussion

### Comment by @codecov-io on January 29, 2020 at 04:40 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@d07f394`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #83   +/-   ##
=========================================
  Coverage          ?   57.42%           
=========================================
  Files             ?       34           
  Lines             ?     1334           
  Branches          ?        0           
=========================================
  Hits              ?      766           
  Misses            ?      478           
  Partials          ?       90
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83?src=pr&el=footer). Last update [d07f394...c0f93fe](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/83?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 30, 2020 at 07:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/83*
