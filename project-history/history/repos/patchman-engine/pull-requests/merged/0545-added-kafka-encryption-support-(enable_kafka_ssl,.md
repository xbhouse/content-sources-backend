---
type: pull_request
number: 545
title: "added kafka encryption support (ENABLE_KAFKA_SSL, KAFKA_SSL_CERT)"
state: merged
author: josef-hak
created: 2021-02-19T10:24:26Z
updated: 2021-02-23T08:15:47Z
closed: 2021-02-19T12:15:43Z
merged: 2021-02-19T12:15:43Z
base_branch: master
head_branch: kafka-encryption
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/545
---

# Pull Request #545: added kafka encryption support (ENABLE_KAFKA_SSL, KAFKA_SSL_CERT)

**Author**: @josef-hak
**Created**: February 19, 2021 at 10:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka-encryption`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 19, 2021 at 10:56 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=h1) Report
> Merging [#545](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=desc) (f7c787e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0b14433a3003f25cb40aa874fa5db28e3bcc0397?el=desc) (0b14433) will **decrease** coverage by `0.17%`.
> The diff coverage is `29.41%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #545      +/-   ##
==========================================
- Coverage   60.21%   60.03%   -0.18%     
==========================================
  Files          66       67       +1     
  Lines        2893     2910      +17     
==========================================
+ Hits         1742     1747       +5     
- Misses        902      914      +12     
  Partials      249      249              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/encryption.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZW5jcnlwdGlvbi5nbw==) | `20.00% <20.00%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `67.79% <100.00%> (+1.12%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=footer). Last update [0b14433...f7c787e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/545?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on February 19, 2021 at 10:34 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/545*
