---
type: pull_request
number: 879
title: "SPM-1260: evaluate baseline"
state: merged
author: MichaelMraka
created: 2022-01-06T15:35:23Z
updated: 2022-01-10T10:14:25Z
closed: 2022-01-10T10:14:24Z
merged: 2022-01-10T10:14:24Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/879
---

# Pull Request #879: SPM-1260: evaluate baseline

**Author**: @MichaelMraka
**Created**: January 06, 2022 at 03:35 PM UTC
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

### Comment by @MichaelMraka on January 07, 2022 at 12:17 PM UTC

Please don't merge it directly I'll squash fixups first after re-review.

### Comment by @codecov-commenter on January 07, 2022 at 12:19 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#879](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ad3125a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5707aa0e90d44a9f69f594dc95541f622d95913a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5707aa0) will **increase** coverage by `0.31%`.
> The diff coverage is `76.92%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #879      +/-   ##
==========================================
+ Coverage   59.04%   59.36%   +0.31%     
==========================================
  Files          82       85       +3     
  Lines        4278     4363      +85     
==========================================
+ Hits         2526     2590      +64     
- Misses       1397     1409      +12     
- Partials      355      364       +9     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.36% <76.92%> (+0.31%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.16% <60.00%> (-0.19%)` | :arrow_down: |
| [base/database/baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9iYXNlbGluZS5nbw==) | `64.28% <64.28%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `76.19% <76.19%> (ø)` | |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `83.33% <83.33%> (ø)` | |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.33% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ac4bdf4...ad3125a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/879?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Approved on January 07, 2022 at 08:50 AM UTC

LGTM. Nice work! Just please update `clowdapp.yaml` and consider to use some cosmetic suggestions.

### Review by @MichaelMraka - Commented on January 07, 2022 at 12:13 PM UTC

### Review by @MichaelMraka - Commented on January 07, 2022 at 12:14 PM UTC

### Review by @MichaelMraka - Commented on January 07, 2022 at 12:14 PM UTC

### Review by @MichaelMraka - Commented on January 07, 2022 at 12:15 PM UTC

### Review by @MichaelMraka - Commented on January 07, 2022 at 12:15 PM UTC

### Review by @josef-hak - Approved on January 07, 2022 at 09:25 PM UTC

:+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/879*
