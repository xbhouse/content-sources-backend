---
type: pull_request
number: 964
title: "SPM-1536: update ubi8-micro to latest RHEL 8.6"
state: closed
author: MichaelMraka
created: 2022-05-31T09:26:21Z
updated: 2022-06-09T16:18:24Z
closed: 2022-06-09T16:18:24Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/964
---

# Pull Request #964: SPM-1536: update ubi8-micro to latest RHEL 8.6

**Author**: @MichaelMraka
**Created**: May 31, 2022 at 09:26 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @codecov-commenter on May 31, 2022 at 09:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#964](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (93cf5a1) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a5478c51875ca0e91519e51833bdf3768d4eef01?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a5478c5) will **increase** coverage by `0.08%`.
> The diff coverage is `72.34%`.

```diff
@@            Coverage Diff             @@
##           master     #964      +/-   ##
==========================================
+ Coverage   61.25%   61.34%   +0.08%     
==========================================
  Files          91       92       +1     
  Lines        4906     4951      +45     
==========================================
+ Hits         3005     3037      +32     
- Misses       1502     1510       +8     
- Partials      399      404       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.34% <72.34%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `32.16% <25.00%> (-0.21%)` | :arrow_down: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `69.69% <69.69%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.72% <100.00%> (+0.70%)` | :arrow_up: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.69% <100.00%> (+0.35%)` | :arrow_up: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.90% <100.00%> (+0.90%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [793cd73...93cf5a1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/964?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 09, 2022 at 04:18 PM UTC

closing in favor of #974

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/964*
