---
type: pull_request
number: 901
title: "SPM-1268: Fixed floorist sql query in deploy/clowdapp.yaml"
state: merged
author: josef-hak
created: 2022-02-04T12:22:43Z
updated: 2022-02-04T12:39:07Z
closed: 2022-02-04T12:39:06Z
merged: 2022-02-04T12:39:06Z
base_branch: master
head_branch: fix-floo-query
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/901
---

# Pull Request #901: SPM-1268: Fixed floorist sql query in deploy/clowdapp.yaml

**Author**: @josef-hak
**Created**: February 04, 2022 at 12:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-floo-query`

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

### Comment by @codecov-commenter on February 04, 2022 at 12:29 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#901](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ecf567e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1b83f73ffd47b9e91fbdf35142d7d02eb4cce235?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1b83f73) will **decrease** coverage by `0.10%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #901      +/-   ##
==========================================
- Coverage   58.80%   58.70%   -0.11%     
==========================================
  Files          88       88              
  Lines        4530     4538       +8     
==========================================
  Hits         2664     2664              
- Misses       1494     1502       +8     
  Partials      372      372              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.70% <0.00%> (-0.11%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <0.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.92% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [7b006f3...ecf567e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/901?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/901*
