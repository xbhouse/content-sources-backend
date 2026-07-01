---
type: pull_request
number: 584
title: "SPM-812 fix wait-for-services criteria"
state: merged
author: josef-hak
created: 2021-03-31T15:00:16Z
updated: 2021-04-12T10:38:58Z
closed: 2021-04-01T07:54:03Z
merged: 2021-04-01T07:54:03Z
base_branch: master
head_branch: fix-wait-for-services
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/584
---

# Pull Request #584: SPM-812 fix wait-for-services criteria

**Author**: @josef-hak
**Created**: March 31, 2021 at 03:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-wait-for-services`

## Description

wait for full db schema when running tests and dirty=f
allow components running on schema dirty=t

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @codecov-io on March 31, 2021 at 03:15 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=h1) Report
> Merging [#584](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=desc) (e508517) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/64eacbd3f733d17c3b3122259f87daca1c03a660?el=desc) (64eacbd) will **not change** coverage.
> The diff coverage is `n/a`.

> :exclamation: Current head e508517 differs from pull request most recent head 19cb9b1. Consider uploading reports for the commit 19cb9b1 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #584   +/-   ##
=======================================
  Coverage   60.22%   60.22%           
=======================================
  Files          68       68           
  Lines        2972     2972           
=======================================
  Hits         1790     1790           
  Misses        924      924           
  Partials      258      258           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.22% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=footer). Last update [64eacbd...19cb9b1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/584?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on April 01, 2021 at 07:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/584*
