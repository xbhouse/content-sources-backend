---
type: pull_request
number: 882
title: "SPM-1261: Baseline update"
state: merged
author: josef-hak
created: 2022-01-10T21:51:58Z
updated: 2022-01-13T12:25:25Z
closed: 2022-01-13T12:22:31Z
merged: 2022-01-13T12:22:30Z
base_branch: master
head_branch: baseline-updates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/882
---

# Pull Request #882: SPM-1261: Baseline update

**Author**: @josef-hak
**Created**: January 10, 2022 at 09:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-updates`

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

### Comment by @codecov-commenter on January 10, 2022 at 09:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#882](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1ba408b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5707aa0e90d44a9f69f594dc95541f622d95913a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5707aa0) will **decrease** coverage by `0.12%`.
> The diff coverage is `52.68%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #882      +/-   ##
==========================================
- Coverage   59.04%   58.92%   -0.13%     
==========================================
  Files          82       86       +4     
  Lines        4278     4506     +228     
==========================================
+ Hits         2526     2655     +129     
- Misses       1397     1477      +80     
- Partials      355      374      +19     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.92% <52.68%> (-0.13%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `71.26% <71.26%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `76.19% <76.19%> (ø)` | |
| [manager/middlewares/rbac.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `62.22% <100.00%> (+2.69%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.16% <0.00%> (-0.19%)` | :arrow_down: |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `83.33% <0.00%> (ø)` | |
| [base/database/baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9iYXNlbGluZS5nbw==) | `64.28% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ac4bdf4...1ba408b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/882?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on January 12, 2022 at 09:08 AM UTC

👍🏼 

---

## Reviews

### Review by @MichaelMraka - Changes Requested on January 11, 2022 at 10:27 AM UTC

### Review by @josef-hak - Commented on January 11, 2022 at 04:05 PM UTC

### Review by @josef-hak - Commented on January 11, 2022 at 04:05 PM UTC

### Review by @MichaelMraka - Approved on January 12, 2022 at 09:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/882*
