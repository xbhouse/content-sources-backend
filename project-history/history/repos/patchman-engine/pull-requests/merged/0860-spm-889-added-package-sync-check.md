---
type: pull_request
number: 860
title: "SPM-889: Added package sync check"
state: merged
author: josef-hak
created: 2021-11-09T16:14:37Z
updated: 2021-11-16T12:46:33Z
closed: 2021-11-16T12:35:03Z
merged: 2021-11-16T12:35:03Z
base_branch: master
head_branch: pkg-sync
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/860
---

# Pull Request #860: SPM-889: Added package sync check

**Author**: @josef-hak
**Created**: November 09, 2021 at 04:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-sync`

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

### Comment by @codecov-commenter on November 09, 2021 at 04:25 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#860](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f3e2f0d) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2fa031524f07c1944505a3782d611c2dfa314746?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2fa0315) will **decrease** coverage by `0.29%`.
> The diff coverage is `8.69%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #860      +/-   ##
==========================================
- Coverage   58.95%   58.66%   -0.30%     
==========================================
  Files          81       81              
  Lines        4208     4229      +21     
==========================================
  Hits         2481     2481              
- Misses       1374     1394      +20     
- Partials      353      354       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.66% <8.69%> (-0.30%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `66.40% <4.54%> (-5.94%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `62.87% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [abde8d3...f3e2f0d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/860?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on November 12, 2021 at 08:55 AM UTC

/retest

### Comment by @josef-hak on November 12, 2021 at 01:11 PM UTC

/retest

### Comment by @josef-hak on November 15, 2021 at 08:19 AM UTC

/retest

### Comment by @josef-hak on November 15, 2021 at 09:46 AM UTC

/retest

### Comment by @josef-hak on November 15, 2021 at 03:42 PM UTC

/retest

### Comment by @josef-hak on November 15, 2021 at 05:24 PM UTC

/retest

### Comment by @josef-hak on November 16, 2021 at 11:45 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on November 16, 2021 at 10:23 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/860*
