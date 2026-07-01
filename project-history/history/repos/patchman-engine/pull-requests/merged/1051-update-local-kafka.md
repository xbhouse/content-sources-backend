---
type: pull_request
number: 1051
title: "Update local kafka"
state: merged
author: psegedy
created: 2022-08-09T16:00:42Z
updated: 2022-08-10T09:54:57Z
closed: 2022-08-10T09:54:57Z
merged: 2022-08-10T09:54:57Z
base_branch: master
head_branch: update_kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1051
---

# Pull Request #1051: Update local kafka

**Author**: @psegedy
**Created**: August 09, 2022 at 04:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `update_kafka`

## Description

use kafka 7.0.5 in docker-compose
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

### Comment by @jira-linking on August 10, 2022 at 07:31 AM UTC

Commits missing Jira IDs:
4afd4dc5614b265e58ead0ab28f3a49dfae50436


### Comment by @codecov-commenter on August 10, 2022 at 07:41 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1051](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (55d5196) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **decrease** coverage by `0.04%`.
> The diff coverage is `72.50%`.

```diff
@@            Coverage Diff             @@
##           master    #1051      +/-   ##
==========================================
- Coverage   61.25%   61.20%   -0.05%     
==========================================
  Files          95       95              
  Lines        5448     5462      +14     
==========================================
+ Hits         3337     3343       +6     
- Misses       1682     1690       +8     
  Partials      429      429              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.20% <72.50%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.69% <ø> (+0.36%)` | :arrow_up: |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.50% <ø> (ø)` | |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.14% <50.00%> (+0.33%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.41% <52.38%> (-2.19%)` | :arrow_down: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `22.85% <100.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.48% <100.00%> (+0.17%)` | :arrow_up: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.37% <100.00%> (-6.56%)` | :arrow_down: |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `75.75% <100.00%> (+1.56%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1051/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


---

## Reviews

### Review by @MichaelMraka - Commented on August 10, 2022 at 07:29 AM UTC

### Review by @MichaelMraka - Approved on August 10, 2022 at 07:42 AM UTC

Don't forget to squash the fixup before merging.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1051*
