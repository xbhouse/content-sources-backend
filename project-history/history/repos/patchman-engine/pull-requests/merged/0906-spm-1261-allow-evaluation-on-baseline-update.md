---
type: pull_request
number: 906
title: "SPM-1261: Allow evaluation on baseline update"
state: merged
author: josef-hak
created: 2022-02-10T21:31:57Z
updated: 2022-02-15T20:15:30Z
closed: 2022-02-15T20:15:29Z
merged: 2022-02-15T20:15:29Z
base_branch: master
head_branch: manager-kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/906
---

# Pull Request #906: SPM-1261: Allow evaluation on baseline update

**Author**: @josef-hak
**Created**: February 10, 2022 at 09:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `manager-kafka`

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

### Comment by @codecov-commenter on February 10, 2022 at 09:40 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#906](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (76496ed) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0a6d789f977bd08adecd02329efb8e457fd49b7c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0a6d789) will **decrease** coverage by `0.02%`.
> The diff coverage is `55.88%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #906      +/-   ##
==========================================
- Coverage   58.98%   58.96%   -0.03%     
==========================================
  Files          89       90       +1     
  Lines        4594     4628      +34     
==========================================
+ Hits         2710     2729      +19     
- Misses       1509     1519      +10     
- Partials      375      380       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.96% <55.88%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `56.66% <ø> (ø)` | |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `60.00% <0.00%> (-15.00%)` | :arrow_down: |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `67.85% <50.00%> (-0.67%)` | :arrow_down: |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `57.14% <57.14%> (ø)` | |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `70.00% <100.00%> (+1.03%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `73.33% <100.00%> (+0.36%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [841213d...76496ed](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/906?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on February 15, 2022 at 07:56 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Changes Requested on February 11, 2022 at 08:44 AM UTC

### Review by @josef-hak - Commented on February 11, 2022 at 09:09 AM UTC

### Review by @MichaelMraka - Changes Requested on February 15, 2022 at 08:30 AM UTC

### Review by @MichaelMraka - Approved on February 15, 2022 at 11:06 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/906*
