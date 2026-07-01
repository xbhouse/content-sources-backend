---
type: pull_request
number: 1046
title: "SPM-1689: Add OrgID error handler"
state: merged
author: michalslomczynski
created: 2022-08-02T15:02:05Z
updated: 2022-08-03T08:08:36Z
closed: 2022-08-03T08:08:36Z
merged: 2022-08-03T08:08:36Z
base_branch: master
head_branch: notification-checks
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1046
---

# Pull Request #1046: SPM-1689: Add OrgID error handler

**Author**: @michalslomczynski
**Created**: August 02, 2022 at 03:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `notification-checks`

## Description

We send some messages with empty OrgID and AccountID. Let's abort in such case, to not spam notifications team, until we find out why incoming Kafka messages do not have these fields.

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

### Comment by @jira-linking on August 02, 2022 at 03:02 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1689


### Comment by @codecov-commenter on August 03, 2022 at 07:46 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1046](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a810a27) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **decrease** coverage by `0.09%`.
> The diff coverage is `65.62%`.

```diff
@@            Coverage Diff             @@
##           master    #1046      +/-   ##
==========================================
- Coverage   61.25%   61.16%   -0.10%     
==========================================
  Files          95       95              
  Lines        5448     5461      +13     
==========================================
+ Hits         3337     3340       +3     
- Misses       1682     1691       +9     
- Partials      429      430       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.16% <65.62%> (-0.10%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.69% <ø> (+0.36%)` | :arrow_up: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `76.81% <0.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.41% <50.00%> (-2.19%)` | :arrow_down: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `22.85% <100.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.03% <100.00%> (-0.28%)` | :arrow_down: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.37% <100.00%> (-6.56%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5e6419c...a810a27](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1046?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @michalslomczynski - Commented on August 03, 2022 at 07:13 AM UTC

### Review by @MichaelMraka - Approved on August 03, 2022 at 07:48 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1046*
