---
type: pull_request
number: 687
title: "SPM-674 Support Kafka dependent components"
state: merged
author: josef-hak
created: 2021-05-27T05:17:31Z
updated: 2021-08-26T18:41:18Z
closed: 2021-05-31T08:05:13Z
merged: 2021-05-31T08:05:13Z
base_branch: master
head_branch: clowder2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/687
---

# Pull Request #687: SPM-674 Support Kafka dependent components

**Author**: @josef-hak
**Created**: May 27, 2021 at 05:17 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `clowder2`

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

### Comment by @codecov-commenter on May 27, 2021 at 05:25 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#687](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8131ad7) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/def933fc606d683d67ec6681f8ff8eb86ee6cee2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (def933f) will **decrease** coverage by `0.37%`.
> The diff coverage is `29.03%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #687      +/-   ##
==========================================
- Coverage   58.13%   57.76%   -0.38%     
==========================================
  Files          73       73              
  Lines        3404     3426      +22     
==========================================
  Hits         1979     1979              
- Misses       1141     1163      +22     
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.76% <29.03%> (-0.38%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |
| [base/utils/log.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9sb2cuZ28=) | `93.10% <100.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.16% <100.00%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `70.37% <100.00%> (ø)` | |
| [manager/middlewares/rbac.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `59.52% <100.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `40.86% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [def933f...8131ad7](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/687?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Commented on May 27, 2021 at 01:40 PM UTC

### Review by @semtexzv - Approved on May 31, 2021 at 07:59 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/687*
