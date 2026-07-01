---
type: pull_request
number: 1121
title: "SPM-1769: smart management check for baselines"
state: closed
author: psegedy
created: 2022-10-10T11:11:25Z
updated: 2022-10-11T12:57:26Z
closed: 2022-10-11T12:57:26Z
base_branch: master
head_branch: entitlements
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1121
---

# Pull Request #1121: SPM-1769: smart management check for baselines

**Author**: @psegedy
**Created**: October 10, 2022 at 11:11 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `entitlements`

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

### Comment by @jira-linking on October 10, 2022 at 11:11 AM UTC

Commits missing Jira IDs:
ae51a83fa3b6774c48edebd6e54d6e58d6735f99
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1769


### Comment by @codecov-commenter on October 11, 2022 at 11:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.98**% // Head: **62.75**% // Decreases project coverage by **`-0.22%`** :warning:
> Coverage data is based on head [(`13d6433`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`908b4bf`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/908b4bff53e7610ddb8e88d7ee710fafa881e18c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 34.06% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1121      +/-   ##
==========================================
- Coverage   62.98%   62.75%   -0.23%     
==========================================
  Files          99      101       +2     
  Lines        5651     5689      +38     
==========================================
+ Hits         3559     3570      +11     
- Misses       1636     1662      +26     
- Partials      456      457       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.75% <34.06%> (-0.23%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/awscloudwatch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9hd3NjbG91ZHdhdGNoLmdv) | `13.79% <ø> (ø)` | |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.14% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.42% <ø> (ø)` | |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.69% <ø> (ø)` | |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `90.90% <ø> (ø)` | |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.26% <ø> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `71.42% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.14% <ø> (ø)` | |
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `23.33% <0.00%> (+0.75%)` | :arrow_up: |
| [manager/middlewares/entitlements.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9lbnRpdGxlbWVudHMuZ28=) | `0.00% <0.00%> (ø)` | |
| ... and [7 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1121?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on October 11, 2022 at 12:57 PM UTC

closing, we will get entitlemnets from x-rh-identity

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1121*
