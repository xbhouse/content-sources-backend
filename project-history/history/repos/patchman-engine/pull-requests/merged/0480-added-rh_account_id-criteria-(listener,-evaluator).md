---
type: pull_request
number: 480
title: "added rh_account_id criteria (listener, evaluator)"
state: merged
author: josef-hak
created: 2021-01-19T15:18:57Z
updated: 2021-02-04T10:49:52Z
closed: 2021-01-20T09:25:16Z
merged: 2021-01-20T09:25:16Z
base_branch: master
head_branch: fix-account-id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/480
---

# Pull Request #480: added rh_account_id criteria (listener, evaluator)

**Author**: @josef-hak
**Created**: January 19, 2021 at 03:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-account-id`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 19, 2021 at 03:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=h1) Report
> Merging [#480](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=desc) (b68bc51) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ab7dfd12d707d8b9b513ae0e75f4de962f0a24bc?el=desc) (ab7dfd1) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #480      +/-   ##
==========================================
+ Coverage   62.71%   62.73%   +0.01%     
==========================================
  Files          61       61              
  Lines        2787     2788       +1     
==========================================
+ Hits         1748     1749       +1     
  Misses        795      795              
  Partials      244      244              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.35% <100.00%> (+0.11%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=footer). Last update [ab7dfd1...b68bc51](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/480?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on January 19, 2021 at 03:51 PM UTC

I'm not sure this will be without issues. If the new evaluator receives old message without the account id, it might not get processed, and block it.

### Comment by @josef-hak on January 19, 2021 at 03:53 PM UTC

@semtexzv we can run old evaluator for some time once new listener is deployed.

---

## Reviews

### Review by @semtexzv - Approved on January 20, 2021 at 08:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/480*
