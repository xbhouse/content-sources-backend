---
type: pull_request
number: 944
title: "SPM-1433: Add endpoint to remove patch sets from system list"
state: merged
author: michalslomczynski
created: 2022-04-21T15:50:43Z
updated: 2022-04-28T08:35:19Z
closed: 2022-04-28T08:35:19Z
merged: 2022-04-28T08:35:19Z
base_branch: master
head_branch: baseline-system-remove
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/944
---

# Pull Request #944: SPM-1433: Add endpoint to remove patch sets from system list

**Author**: @michalslomczynski
**Created**: April 21, 2022 at 03:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-system-remove`

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

### Comment by @codecov-commenter on April 26, 2022 at 03:21 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#944](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8791323) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ae8a1aae4de98327e0a7ba06a80f63dc2f030f52?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ae8a1aa) will **decrease** coverage by `0.03%`.
> The diff coverage is `55.88%`.

```diff
@@            Coverage Diff             @@
##           master     #944      +/-   ##
==========================================
- Coverage   60.37%   60.34%   -0.04%     
==========================================
  Files          91       92       +1     
  Lines        4838     4869      +31     
==========================================
+ Hits         2921     2938      +17     
- Misses       1529     1539      +10     
- Partials      388      392       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.34% <55.88%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.85% <ø> (ø)` | |
| [manager/controllers/baseline\_systems\_remove.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX3JlbW92ZS5nbw==) | `51.72% <51.72%> (ø)` | |
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.68% <75.00%> (+0.76%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `79.51% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [db727dc...8791323](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/944?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on April 26, 2022 at 09:32 AM UTC

### Review by @MichaelMraka - Commented on April 26, 2022 at 01:35 PM UTC

### Review by @MichaelMraka - Approved on April 27, 2022 at 11:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/944*
