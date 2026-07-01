---
type: pull_request
number: 891
title: "SPM-1331: Mark db systems with host_type=edge as stale (db migr.)"
state: merged
author: josef-hak
created: 2022-01-21T21:58:27Z
updated: 2022-01-24T18:00:43Z
closed: 2022-01-24T18:00:42Z
merged: 2022-01-24T18:00:42Z
base_branch: master
head_branch: rm-edge-db
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/891
---

# Pull Request #891: SPM-1331: Mark db systems with host_type=edge as stale (db migr.)

**Author**: @josef-hak
**Created**: January 21, 2022 at 09:58 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rm-edge-db`

## Description

- the same approach used for SPM-626
- then it will be deleted automatically using system culling job (vmaas-sync)

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

### Comment by @codecov-commenter on January 21, 2022 at 10:06 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#891](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b3a333a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8c172b594e55e35e8827d0de51a3e91ded395438?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8c172b5) will **not change** coverage.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #891   +/-   ##
=======================================
  Coverage   59.24%   59.24%           
=======================================
  Files          88       88           
  Lines        4601     4601           
=======================================
  Hits         2726     2726           
  Misses       1492     1492           
  Partials      383      383           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.24% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.53% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [31a1085...b3a333a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/891?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 24, 2022 at 03:24 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/891*
