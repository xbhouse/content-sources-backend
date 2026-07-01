---
type: pull_request
number: 932
title: "SPM-1416: Add baseline description field"
state: merged
author: michalslomczynski
created: 2022-03-29T13:05:02Z
updated: 2022-04-04T11:53:09Z
closed: 2022-04-04T11:53:09Z
merged: 2022-04-04T11:53:09Z
base_branch: master
head_branch: baseline-description
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/932
---

# Pull Request #932: SPM-1416: Add baseline description field

**Author**: @michalslomczynski
**Created**: March 29, 2022 at 01:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-description`

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

### Comment by @codecov-commenter on March 29, 2022 at 01:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#932](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f2a25e6) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/37a167352c6b16d453d3d229d206a170a7731175?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (37a1673) will **increase** coverage by `0.14%`.
> The diff coverage is `64.51%`.

```diff
@@            Coverage Diff             @@
##           master     #932      +/-   ##
==========================================
+ Coverage   60.21%   60.36%   +0.14%     
==========================================
  Files          91       91              
  Lines        4818     4831      +13     
==========================================
+ Hits         2901     2916      +15     
- Misses       1525     1528       +3     
+ Partials      392      387       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.36% <64.51%> (+0.14%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `9.91% <0.00%> (-0.09%)` | :arrow_down: |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `75.80% <100.00%> (+6.31%)` | :arrow_up: |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <100.00%> (+0.91%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `79.51% <100.00%> (+7.22%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [37a1673...f2a25e6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/932?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @michalslomczynski - Commented on March 30, 2022 at 07:38 AM UTC

### Review by @MichaelMraka - Changes Requested on March 30, 2022 at 09:30 AM UTC

### Review by @michalslomczynski - Commented on March 30, 2022 at 10:24 AM UTC

### Review by @michalslomczynski - Commented on March 30, 2022 at 10:28 AM UTC

### Review by @michalslomczynski - Commented on March 30, 2022 at 02:42 PM UTC

### Review by @michalslomczynski - Commented on March 30, 2022 at 02:44 PM UTC

### Review by @michalslomczynski - Commented on March 30, 2022 at 02:46 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 07:49 AM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 08:53 AM UTC

### Review by @MichaelMraka - Changes Requested on March 31, 2022 at 09:48 AM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 12:51 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 12:54 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 12:56 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 12:58 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 12:59 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 01:18 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 01:33 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 01:47 PM UTC

### Review by @michalslomczynski - Commented on March 31, 2022 at 02:18 PM UTC

### Review by @MichaelMraka - Approved on April 01, 2022 at 02:28 PM UTC

I've put some improvements into PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/932*
