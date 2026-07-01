---
type: pull_request
number: 941
title: "SPM-1324: let's log whole message when tracing error"
state: merged
author: MichaelMraka
created: 2022-04-13T09:14:22Z
updated: 2022-05-27T11:04:29Z
closed: 2022-05-27T11:04:28Z
merged: 2022-05-27T11:04:28Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/941
---

# Pull Request #941: SPM-1324: let's log whole message when tracing error

**Author**: @MichaelMraka
**Created**: April 13, 2022 at 09:14 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @codecov-commenter on April 14, 2022 at 10:03 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#941](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (02083da) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ae8a1aae4de98327e0a7ba06a80f63dc2f030f52?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ae8a1aa) will **increase** coverage by `0.06%`.
> The diff coverage is `91.66%`.

```diff
@@            Coverage Diff             @@
##           master     #941      +/-   ##
==========================================
+ Coverage   60.37%   60.44%   +0.06%     
==========================================
  Files          91       91              
  Lines        4838     4849      +11     
==========================================
+ Hits         2921     2931      +10     
- Misses       1529     1530       +1     
  Partials      388      388              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.44% <91.66%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.76% <90.90%> (+0.50%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `79.51% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9c561d7...02083da](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/941?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/941*
