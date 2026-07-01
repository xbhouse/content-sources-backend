---
type: pull_request
number: 962
title: "SPM-1501: add payload tracker and use org_id from kafka messages"
state: merged
author: psegedy
created: 2022-05-23T10:58:21Z
updated: 2022-06-02T10:49:20Z
closed: 2022-06-02T10:49:20Z
merged: 2022-06-02T10:49:20Z
base_branch: master
head_branch: payload_tracker
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/962
---

# Pull Request #962: SPM-1501: add payload tracker and use org_id from kafka messages

**Author**: @psegedy
**Created**: May 23, 2022 at 10:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `payload_tracker`

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

### Comment by @codecov-commenter on June 01, 2022 at 04:56 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#962](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2fc3051) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/486c1e1ed53d040254f49b1754cbf3153be002ee?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (486c1e1) will **decrease** coverage by `0.38%`.
> The diff coverage is `45.28%`.

```diff
@@            Coverage Diff             @@
##           master     #962      +/-   ##
==========================================
- Coverage   61.30%   60.91%   -0.39%     
==========================================
  Files          92       94       +2     
  Lines        4975     5102     +127     
==========================================
+ Hits         3050     3108      +58     
- Misses       1518     1585      +67     
- Partials      407      409       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.91% <45.28%> (-0.39%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <0.00%> (+22.41%)` | :arrow_up: |
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `9.52% <9.52%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `31.97% <25.00%> (-0.20%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.12% <62.50%> (-0.60%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.45% <97.82%> (+2.64%)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `68.75% <100.00%> (+2.08%)` | :arrow_up: |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <100.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `34.61% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5f4cba5...2fc3051](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/962?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 02, 2022 at 09:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/962*
