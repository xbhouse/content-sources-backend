---
type: pull_request
number: 969
title: "SPM-1537: replace shell scripts"
state: merged
author: psegedy
created: 2022-06-03T17:51:16Z
updated: 2022-06-09T08:37:10Z
closed: 2022-06-09T08:37:10Z
merged: 2022-06-09T08:37:09Z
base_branch: master
head_branch: wait_for_go
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/969
---

# Pull Request #969: SPM-1537: replace shell scripts

**Author**: @psegedy
**Created**: June 03, 2022 at 05:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `wait_for_go`

## Description

- replaced wait-for-services.sh
- db migration without shell scripts
- ~removed export of clowder variables as env vars~
- removed postgresql from image
- removed shell scripts depending on the functionality listed above

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

### Comment by @psegedy on June 03, 2022 at 06:11 PM UTC

doesn't work in ephemeral with clowder :/
```
time="2022-06-03T18:06:13Z" level=error msg=Panicked err="dial tcp 172.30.157.235:0: connect: no route to host" stack="goroutine 1 [running]:|runtime/debug.Stack()|\t/usr/lib/golang/src/runtime/debug/stack.go:24 +0x65|app/base/utils.LogPanics(0x1)|\t/go/src/app/base/utils/core.go:138 +0x58|panic({0x1050300, 0xc00047c7d0})|\t/usr/lib/golang/src/runtime/panic.go:1038 +0x215|app/database_admin.UpdateDB({0x7fffefe3ff85, 0x22})|\t/go/src/app/database_admin/update.go:84 +0x571|main.main()|\t/go/src/app/main.go:35 +0x10a|"
```

### Comment by @codecov-commenter on June 06, 2022 at 11:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#969](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0e93f3f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/61b0dac7628cfcf95e5702bf0cfffe8eb23439bc?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (61b0dac) will **increase** coverage by `0.19%`.
> The diff coverage is `7.40%`.

```diff
@@            Coverage Diff             @@
##           master     #969      +/-   ##
==========================================
+ Coverage   60.93%   61.12%   +0.19%     
==========================================
  Files          94       94              
  Lines        5104     5091      -13     
==========================================
+ Hits         3110     3112       +2     
+ Misses       1585     1570      -15     
  Partials      409      409              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.12% <7.40%> (+0.19%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `43.92% <ø> (+11.95%)` | :arrow_up: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `77.77% <100.00%> (+0.41%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [92714aa...0e93f3f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/969?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 08, 2022 at 10:54 AM UTC

PR checks are failing with `E       fixture 'upload_system' not found` 

### Comment by @psegedy on June 08, 2022 at 10:56 AM UTC

/retest

### Comment by @digitronik on June 09, 2022 at 07:58 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Commented on June 06, 2022 at 12:51 PM UTC

### Review by @psegedy - Commented on June 07, 2022 at 06:36 PM UTC

### Review by @MichaelMraka - Approved on June 08, 2022 at 01:47 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/969*
