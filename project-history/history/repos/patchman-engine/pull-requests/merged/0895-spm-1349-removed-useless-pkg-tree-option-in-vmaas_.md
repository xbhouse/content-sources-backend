---
type: pull_request
number: 895
title: "SPM-1349: Removed useless pkg tree option in vmaas_sync"
state: merged
author: josef-hak
created: 2022-02-01T10:13:22Z
updated: 2022-02-01T12:39:18Z
closed: 2022-02-01T12:39:17Z
merged: 2022-02-01T12:39:17Z
base_branch: master
head_branch: fixes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/895
---

# Pull Request #895: SPM-1349: Removed useless pkg tree option in vmaas_sync

**Author**: @josef-hak
**Created**: February 01, 2022 at 10:13 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fixes`

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

### Comment by @codecov-commenter on February 01, 2022 at 10:20 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#895](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4076bc3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5c37148957af45a84198ef4ac77a8a8103294ee4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5c37148) will **decrease** coverage by `0.35%`.
> The diff coverage is `77.77%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #895      +/-   ##
==========================================
- Coverage   59.34%   58.99%   -0.36%     
==========================================
  Files          88       88              
  Lines        4627     4516     -111     
==========================================
- Hits         2746     2664      -82     
+ Misses       1496     1480      -16     
+ Partials      385      372      -13     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.99% <77.77%> (-0.36%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `41.17% <ø> (-0.50%)` | :arrow_down: |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `60.95% <77.77%> (-5.45%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [c298429...4076bc3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/895?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 01, 2022 at 12:38 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/895*
