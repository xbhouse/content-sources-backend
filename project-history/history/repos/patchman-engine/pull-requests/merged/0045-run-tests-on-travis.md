---
type: pull_request
number: 45
title: "Run tests on Travis"
state: merged
author: mkoura
created: 2020-01-06T13:47:14Z
updated: 2020-01-07T10:23:32Z
closed: 2020-01-07T10:23:13Z
merged: 2020-01-07T10:23:13Z
base_branch: master
head_branch: add_travis
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/45
---

# Pull Request #45: Run tests on Travis

**Author**: @mkoura
**Created**: January 06, 2020 at 01:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add_travis`

## Description

Run unit tests on Travis. I don't have rights to enable the repo on Travis, can some of you please do it? - https://travis-ci.org/RedHatInsights/patchman-engine/

Also adds upload to codecov.io (only when running on Travis).

Test run on https://travis-ci.org/mkoura/patchman-engine/builds/633258180

---

## Discussion

### Comment by @josef-hak on January 06, 2020 at 02:27 PM UTC

@mkoura ok, I've added repo to Travis. Now there are some conflicts, could you resolve it, please? Another thing - I've added dependency on platform container in the last PR. Seed docker-compose.test.yml. Could you add it also to travis config?

### Comment by @beav on January 06, 2020 at 04:00 PM UTC

@mkoura is something else needed to get travis to run on this PR?

### Comment by @mkoura on January 06, 2020 at 04:08 PM UTC

@beav I still can see

>  This is not an active repository
>
> You don't have sufficient rights to enable this repo on Travis.
> Please contact the admin to enable it or to receive admin rights yourself. 

on https://travis-ci.org/RedHatInsights/patchman-engine/

I guess that's the reason why travis hasn't tested this PR yet.

### Comment by @josef-hak on January 06, 2020 at 07:40 PM UTC

@mkoura I see the same message about rights. @semtexzv could you help us to resolve this or could you increase our rights over this repo?

### Comment by @semtexzv on January 07, 2020 at 08:21 AM UTC

Travis was enabled.

### Comment by @josef-hak on January 07, 2020 at 09:17 AM UTC

@mkoura Travis was enabled, but [first build failed](https://travis-ci.org/RedHatInsights/patchman-engine/builds/633660537). How can we test your PR with Travis config? Could you do another push to your branch?

### Comment by @codecov-io on January 07, 2020 at 10:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@ebf014a`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #45   +/-   ##
=========================================
  Coverage          ?   56.68%           
=========================================
  Files             ?       26           
  Lines             ?      912           
  Branches          ?        0           
=========================================
  Hits              ?      517           
  Misses            ?      337           
  Partials          ?       58
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45?src=pr&el=footer). Last update [ebf014a...c874414](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/45?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on January 07, 2020 at 09:28 AM UTC

### Review by @mkoura - Commented on January 07, 2020 at 09:51 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/45*
