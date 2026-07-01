---
type: pull_request
number: 899
title: "SPM-1349: Used simplified vmaas client in listener (structs only)"
state: merged
author: josef-hak
created: 2022-02-03T18:21:59Z
updated: 2022-02-04T10:40:29Z
closed: 2022-02-04T10:40:29Z
merged: 2022-02-04T10:40:29Z
base_branch: master
head_branch: rm-vmaas-listener
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/899
---

# Pull Request #899: SPM-1349: Used simplified vmaas client in listener (structs only)

**Author**: @josef-hak
**Created**: February 03, 2022 at 06:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rm-vmaas-listener`

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

### Comment by @codecov-commenter on February 03, 2022 at 06:29 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#899](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ebe8cdd) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1b83f73ffd47b9e91fbdf35142d7d02eb4cce235?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1b83f73) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #899   +/-   ##
=======================================
  Coverage   58.80%   58.80%           
=======================================
  Files          88       88           
  Lines        4530     4530           
=======================================
  Hits         2664     2664           
  Misses       1494     1494           
  Partials      372      372           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.80% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.92% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [31c5bb0...ebe8cdd](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/899?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on February 03, 2022 at 07:19 PM UTC

/retest

### Comment by @josef-hak on February 03, 2022 at 10:49 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on February 04, 2022 at 10:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/899*
