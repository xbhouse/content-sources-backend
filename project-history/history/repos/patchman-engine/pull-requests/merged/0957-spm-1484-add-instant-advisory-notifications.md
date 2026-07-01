---
type: pull_request
number: 957
title: "SPM-1484: Add instant advisory notifications"
state: merged
author: michalslomczynski
created: 2022-05-11T12:02:35Z
updated: 2022-05-27T11:04:04Z
closed: 2022-05-27T11:04:03Z
merged: 2022-05-27T11:04:03Z
base_branch: master
head_branch: advisory_notifications
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/957
---

# Pull Request #957: SPM-1484: Add instant advisory notifications

**Author**: @michalslomczynski
**Created**: May 11, 2022 at 12:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisory_notifications`

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

### Comment by @codecov-commenter on May 11, 2022 at 01:32 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#957](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (683eba8) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2c8e5bfd7b37d19a7b040a38cc3f27c0300b36b7?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2c8e5bf) will **increase** coverage by `0.80%`.
> The diff coverage is `87.17%`.

```diff
@@            Coverage Diff             @@
##           master     #957      +/-   ##
==========================================
+ Coverage   60.60%   61.40%   +0.80%     
==========================================
  Files          92       92              
  Lines        4871     4858      -13     
==========================================
+ Hits         2952     2983      +31     
+ Misses       1523     1477      -46     
- Partials      396      398       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.40% <87.17%> (+0.80%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `83.33% <83.33%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.72% <100.00%> (+0.70%)` | :arrow_up: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `77.37% <100.00%> (+1.03%)` | :arrow_up: |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `65.00% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_confluentic.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfY29uZmx1ZW50aWMuZ28=) | | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `70.83% <0.00%> (+14.16%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2c8e5bf...683eba8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/957?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on May 19, 2022 at 10:09 AM UTC

Maybe `last_updated` column can be dropped. Will be sure after aggregation in notifications-backend.

### Comment by @michalslomczynski on May 25, 2022 at 04:01 PM UTC

Stripped DB since we do not need it anymore and pushed each advisory to separate event because notifications-backend prefers this way. Should be final version and ready for review now @MichaelMraka :slightly_smiling_face: Please also take a look on advisories type comment, however it can be handled in separate PR in my opinion since it could affect existing endpoints out of this scope etc.

---

## Reviews

### Review by @MichaelMraka - Changes Requested on May 25, 2022 at 09:46 AM UTC

### Review by @michalslomczynski - Commented on May 25, 2022 at 12:40 PM UTC

### Review by @michalslomczynski - Commented on May 25, 2022 at 12:52 PM UTC

### Review by @michalslomczynski - Commented on May 25, 2022 at 02:52 PM UTC

### Review by @MichaelMraka - Approved on May 27, 2022 at 09:39 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/957*
