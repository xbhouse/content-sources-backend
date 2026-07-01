---
type: pull_request
number: 1048
title: "SPM-1692: parse pkg with empty version, release, arch"
state: merged
author: psegedy
created: 2022-08-05T08:22:25Z
updated: 2022-08-05T08:52:57Z
closed: 2022-08-05T08:52:56Z
merged: 2022-08-05T08:52:56Z
base_branch: master
head_branch: nevra_regex
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1048
---

# Pull Request #1048: SPM-1692: parse pkg with empty version, release, arch

**Author**: @psegedy
**Created**: August 05, 2022 at 08:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `nevra_regex`

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

### Comment by @jira-linking on August 05, 2022 at 08:22 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1692


### Comment by @codecov-commenter on August 05, 2022 at 08:31 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1048](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0ba1b20) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **decrease** coverage by `0.06%`.
> The diff coverage is `71.05%`.

```diff
@@            Coverage Diff             @@
##           master    #1048      +/-   ##
==========================================
- Coverage   61.25%   61.18%   -0.07%     
==========================================
  Files          95       95              
  Lines        5448     5464      +16     
==========================================
+ Hits         3337     3343       +6     
- Misses       1682     1691       +9     
- Partials      429      430       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.18% <71.05%> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.69% <ø> (+0.36%)` | :arrow_up: |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.50% <ø> (ø)` | |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.14% <50.00%> (+0.33%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.41% <52.38%> (-2.19%)` | :arrow_down: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `22.85% <100.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.03% <100.00%> (-0.28%)` | :arrow_down: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.37% <100.00%> (-6.56%)` | :arrow_down: |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `75.75% <100.00%> (+1.56%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1048/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


---

## Reviews

### Review by @MichaelMraka - Approved on August 05, 2022 at 08:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1048*
