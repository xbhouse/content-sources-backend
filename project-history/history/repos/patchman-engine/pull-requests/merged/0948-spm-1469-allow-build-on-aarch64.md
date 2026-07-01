---
type: pull_request
number: 948
title: "SPM-1469: allow build on aarch64"
state: merged
author: psegedy
created: 2022-05-03T10:39:19Z
updated: 2022-05-05T07:51:07Z
closed: 2022-05-05T07:51:07Z
merged: 2022-05-05T07:51:07Z
base_branch: master
head_branch: aarch64
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/948
---

# Pull Request #948: SPM-1469: allow build on aarch64

**Author**: @psegedy
**Created**: May 03, 2022 at 10:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `aarch64`

## Description

build on aarch64 requires building [librdkafka](https://github.com/edenhill/librdkafka) from source and building go app with `go build -tags dynamic` to link shared library with [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) 

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

### Comment by @codecov-commenter on May 04, 2022 at 03:29 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/948?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#948](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/948?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (90bbc13) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5f1dae0d98e90ed82c3e64ee52cf25204ac852c1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5f1dae0) will **not change** coverage.
> The diff coverage is `n/a`.

```diff
@@           Coverage Diff           @@
##           master     #948   +/-   ##
=======================================
  Coverage   60.45%   60.45%           
=======================================
  Files          92       92           
  Lines        4871     4871           
=======================================
  Hits         2945     2945           
  Misses       1534     1534           
  Partials      392      392           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/948?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/948?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e774d81...90bbc13](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/948?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on May 03, 2022 at 12:02 PM UTC

### Review by @psegedy - Commented on May 04, 2022 at 09:37 AM UTC

### Review by @psegedy - Commented on May 04, 2022 at 09:37 AM UTC

### Review by @psegedy - Commented on May 04, 2022 at 09:38 AM UTC

### Review by @psegedy - Commented on May 04, 2022 at 09:43 AM UTC

### Review by @MichaelMraka - Approved on May 04, 2022 at 02:35 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/948*
