---
type: pull_request
number: 931
title: "SPM-1326: Handle updateinfo in evaluator"
state: merged
author: michalslomczynski
created: 2022-03-21T15:28:08Z
updated: 2022-06-22T14:28:59Z
closed: 2022-06-22T14:28:59Z
merged: 2022-06-22T14:28:59Z
base_branch: master
head_branch: updateinfo-evaluator
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/931
---

# Pull Request #931: SPM-1326: Handle updateinfo in evaluator

**Author**: @michalslomczynski
**Created**: March 21, 2022 at 03:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `updateinfo-evaluator`

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

### Comment by @codecov-commenter on March 28, 2022 at 02:32 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#931](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6739517) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f59500b7d6044b526fb667512d8454c025283c98?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f59500b) will **increase** coverage by `0.08%`.
> The diff coverage is `66.96%`.

```diff
@@            Coverage Diff             @@
##           master     #931      +/-   ##
==========================================
+ Coverage   60.64%   60.72%   +0.08%     
==========================================
  Files          94       95       +1     
  Lines        5211     5281      +70     
==========================================
+ Hits         3160     3207      +47     
- Misses       1640     1650      +10     
- Partials      411      424      +13     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.72% <66.96%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.73% <54.23%> (-2.65%)` | :arrow_down: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.77% <77.77%> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.50% <100.00%> (+10.83%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.74% <100.00%> (+0.29%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8747b7c...6739517](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/931?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on April 11, 2022 at 02:56 PM UTC

I've pushed my fixes for mentioned issues.

### Comment by @michalslomczynski on April 12, 2022 at 11:10 AM UTC

Great changes. Will you merge?

### Comment by @psegedy on June 20, 2022 at 04:49 PM UTC

/retest

---

## Reviews

### Review by @michalslomczynski - Commented on March 22, 2022 at 03:09 PM UTC

I will fix tests interference tomorrow.

### Review by @MichaelMraka - Changes Requested on March 25, 2022 at 03:02 PM UTC

### Review by @michalslomczynski - Commented on March 28, 2022 at 07:10 AM UTC

### Review by @MichaelMraka - Commented on April 11, 2022 at 02:14 PM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 09:49 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:04 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:04 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:05 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:25 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:55 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 10:55 AM UTC

### Review by @MichaelMraka - Commented on April 12, 2022 at 11:45 AM UTC

### Review by @michalslomczynski - Commented on April 12, 2022 at 11:52 AM UTC

### Review by @MichaelMraka - Commented on April 12, 2022 at 12:54 PM UTC

### Review by @psegedy - Commented on June 16, 2022 at 09:52 AM UTC

### Review by @MichaelMraka - Commented on June 16, 2022 at 11:18 AM UTC

### Review by @psegedy - Commented on June 16, 2022 at 01:59 PM UTC

### Review by @MichaelMraka - Commented on June 17, 2022 at 08:12 AM UTC

### Review by @MichaelMraka - Commented on June 17, 2022 at 09:12 AM UTC

### Review by @psegedy - Commented on June 17, 2022 at 11:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/931*
