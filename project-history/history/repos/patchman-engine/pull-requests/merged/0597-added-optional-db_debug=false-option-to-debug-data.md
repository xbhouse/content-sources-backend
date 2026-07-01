---
type: pull_request
number: 597
title: "Added optional DB_DEBUG=false option to debug database transactions"
state: merged
author: josef-hak
created: 2021-04-08T13:13:20Z
updated: 2021-04-12T10:38:40Z
closed: 2021-04-08T16:37:23Z
merged: 2021-04-08T16:37:23Z
base_branch: master
head_branch: db-debug
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/597
---

# Pull Request #597: Added optional DB_DEBUG=false option to debug database transactions

**Author**: @josef-hak
**Created**: April 08, 2021 at 01:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `db-debug`

## Description

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

### Comment by @codecov-io on April 08, 2021 at 01:21 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=h1) Report
> Merging [#597](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=desc) (1ff9eda) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/42b9ef90e774978c8c0d6aa6a142a12ca59124c5?el=desc) (42b9ef9) will **decrease** coverage by `0.02%`.
> The diff coverage is `33.33%`.

> :exclamation: Current head 1ff9eda differs from pull request most recent head 778e6fc. Consider uploading reports for the commit 778e6fc to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #597      +/-   ##
==========================================
- Coverage   59.05%   59.03%   -0.03%     
==========================================
  Files          70       70              
  Lines        3158     3161       +3     
==========================================
+ Hits         1865     1866       +1     
- Misses       1023     1024       +1     
- Partials      270      271       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.03% <33.33%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `76.47% <33.33%> (-4.18%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=footer). Last update [42b9ef9...778e6fc](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/597?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on April 08, 2021 at 04:20 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/597*
