---
type: pull_request
number: 877
title: "SPM-1161: new endpoint for updating a baseline"
state: closed
author: mkholjuraev
created: 2021-12-16T10:37:47Z
updated: 2022-01-10T21:56:27Z
closed: 2022-01-10T21:56:27Z
base_branch: master
head_branch: update-baseline
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/877
---

# Pull Request #877: SPM-1161: new endpoint for updating a baseline

**Author**: @mkholjuraev
**Created**: December 16, 2021 at 10:37 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `update-baseline`

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

### Comment by @codecov-commenter on December 16, 2021 at 11:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#877](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c29ccb3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5707aa0e90d44a9f69f594dc95541f622d95913a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5707aa0) will **decrease** coverage by `0.19%`.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #877      +/-   ##
==========================================
- Coverage   59.04%   58.84%   -0.20%     
==========================================
  Files          82       86       +4     
  Lines        4278     4486     +208     
==========================================
+ Hits         2526     2640     +114     
- Misses       1397     1474      +77     
- Partials      355      372      +17     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.84% <50.00%> (-0.20%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `70.14% <70.14%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `76.19% <76.19%> (ø)` | |
| [manager/middlewares/rbac.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `62.22% <100.00%> (+2.69%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.16% <0.00%> (-0.19%)` | :arrow_down: |
| [base/database/baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9iYXNlbGluZS5nbw==) | `64.28% <0.00%> (ø)` | |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `83.33% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ac4bdf4...c29ccb3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/877?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on January 10, 2022 at 09:56 PM UTC

Included to #882 , added some improvements

---

## Reviews

### Review by @josef-hak - Changes Requested on December 16, 2021 at 02:27 PM UTC

some conflicts after previous PR merge

### Review by @josef-hak - Changes Requested on December 16, 2021 at 03:19 PM UTC

### Review by @mkholjuraev - Commented on December 20, 2021 at 10:49 PM UTC

### Review by @josef-hak - Commented on December 21, 2021 at 02:29 PM UTC

### Review by @mkholjuraev - Commented on December 22, 2021 at 11:01 AM UTC

### Review by @josef-hak - Changes Requested on January 05, 2022 at 10:24 AM UTC

### Review by @josef-hak - Changes Requested on January 05, 2022 at 10:46 AM UTC

### Review by @josef-hak - Changes Requested on January 06, 2022 at 08:50 PM UTC

Uh, I found yet some important things to improve. But we are almost there :wink: 

### Review by @mkholjuraev - Commented on January 07, 2022 at 09:39 AM UTC

### Review by @josef-hak - Commented on January 07, 2022 at 09:56 AM UTC

### Review by @mkholjuraev - Commented on January 07, 2022 at 10:23 AM UTC

### Review by @josef-hak - Commented on January 07, 2022 at 03:16 PM UTC

### Review by @josef-hak - Changes Requested on January 10, 2022 at 12:58 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/877*
