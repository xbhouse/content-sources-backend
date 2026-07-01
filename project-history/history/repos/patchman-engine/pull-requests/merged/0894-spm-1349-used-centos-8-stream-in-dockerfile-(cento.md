---
type: pull_request
number: 894
title: "SPM-1349: Used CentOS 8 Stream in Dockerfile (Centos 8 EOL)"
state: merged
author: josef-hak
created: 2022-01-31T23:07:18Z
updated: 2022-02-01T09:26:22Z
closed: 2022-02-01T09:26:21Z
merged: 2022-02-01T09:26:21Z
base_branch: master
head_branch: centos
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/894
---

# Pull Request #894: SPM-1349: Used CentOS 8 Stream in Dockerfile (Centos 8 EOL)

**Author**: @josef-hak
**Created**: January 31, 2022 at 11:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `centos`

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

### Comment by @josef-hak on February 01, 2022 at 08:11 AM UTC

/retest

### Comment by @codecov-commenter on February 01, 2022 at 08:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#894](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9fc71fe) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4d1f24e7a6654d7f07b1586895cc0008eadb50a2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4d1f24e) will **increase** coverage by `0.08%`.
> The diff coverage is `62.50%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #894      +/-   ##
==========================================
+ Coverage   59.26%   59.34%   +0.08%     
==========================================
  Files          88       88              
  Lines        4608     4627      +19     
==========================================
+ Hits         2731     2746      +15     
- Misses       1493     1496       +3     
- Partials      384      385       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.34% <62.50%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `72.60% <ø> (ø)` | |
| [manager/middlewares/swagger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9zd2FnZ2VyLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `41.66% <0.00%> (ø)` | |
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `65.21% <71.42%> (+65.21%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b148ecf...9fc71fe](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/894?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 01, 2022 at 09:09 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/894*
