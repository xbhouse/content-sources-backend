---
type: pull_request
number: 930
title: "SPM-1325: Update listener to read updateinfo from kafka"
state: merged
author: michalslomczynski
created: 2022-03-18T15:06:07Z
updated: 2022-03-24T12:09:48Z
closed: 2022-03-24T12:09:48Z
merged: 2022-03-24T12:09:48Z
base_branch: master
head_branch: read-updateinfo-kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/930
---

# Pull Request #930: SPM-1325: Update listener to read updateinfo from kafka

**Author**: @michalslomczynski
**Created**: March 18, 2022 at 03:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `read-updateinfo-kafka`

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

### Comment by @codecov-commenter on March 18, 2022 at 03:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#930](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0d2dd74) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/91c8b93848852d6325274f87d28faf50ab3c9313?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (91c8b93) will **increase** coverage by `0.08%`.
> The diff coverage is `85.71%`.

```diff
@@            Coverage Diff             @@
##           master     #930      +/-   ##
==========================================
+ Coverage   60.12%   60.21%   +0.08%     
==========================================
  Files          91       91              
  Lines        4800     4818      +18     
==========================================
+ Hits         2886     2901      +15     
- Misses       1523     1525       +2     
- Partials      391      392       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.21% <85.71%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.25% <85.71%> (+0.32%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [12311f2...0d2dd74](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/930?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on March 18, 2022 at 09:45 PM UTC

/retest

---

## Reviews

### Review by @josef-hak - Changes Requested on March 18, 2022 at 03:19 PM UTC

### Review by @michalslomczynski - Commented on March 18, 2022 at 03:30 PM UTC

### Review by @josef-hak - Commented on March 18, 2022 at 09:30 PM UTC

### Review by @MichaelMraka - Changes Requested on March 21, 2022 at 09:01 AM UTC

### Review by @michalslomczynski - Commented on March 21, 2022 at 10:18 AM UTC

### Review by @michalslomczynski - Commented on March 21, 2022 at 11:13 AM UTC

### Review by @MichaelMraka - Changes Requested on March 23, 2022 at 12:35 PM UTC

### Review by @michalslomczynski - Commented on March 23, 2022 at 12:41 PM UTC

### Review by @MichaelMraka - Approved on March 24, 2022 at 12:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/930*
