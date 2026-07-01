---
type: pull_request
number: 900
title: "SPM-1361: Added \"synced\" boolean table columns"
state: closed
author: josef-hak
created: 2022-02-03T20:04:55Z
updated: 2022-03-25T14:59:05Z
closed: 2022-03-25T14:59:05Z
base_branch: master
head_branch: synced-mark
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/900
---

# Pull Request #900: SPM-1361: Added "synced" boolean table columns

**Author**: @josef-hak
**Created**: February 03, 2022 at 08:04 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `synced-mark`

## Description

- tables: package, advisory_metadata, repo

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

### Comment by @codecov-commenter on February 03, 2022 at 10:54 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#900](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2cf1d17) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1b83f73ffd47b9e91fbdf35142d7d02eb4cce235?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1b83f73) will **decrease** coverage by `0.03%`.
> The diff coverage is `42.85%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #900      +/-   ##
==========================================
- Coverage   58.80%   58.77%   -0.04%     
==========================================
  Files          88       88              
  Lines        4530     4536       +6     
==========================================
+ Hits         2664     2666       +2     
- Misses       1494     1498       +4     
  Partials      372      372              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.77% <42.85%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `63.09% <100.00%> (+0.22%)` | :arrow_up: |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `61.22% <100.00%> (+0.26%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [31c5bb0...2cf1d17](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/900?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on February 04, 2022 at 08:52 AM UTC

/retest

### Comment by @MichaelMraka on March 25, 2022 at 02:59 PM UTC

Code has been pulled into #927 

---

## Reviews

### Review by @MichaelMraka - Approved on February 04, 2022 at 10:39 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/900*
