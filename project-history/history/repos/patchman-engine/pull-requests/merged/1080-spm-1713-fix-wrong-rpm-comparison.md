---
type: pull_request
number: 1080
title: "SPM-1713: fix wrong rpm comparison"
state: merged
author: psegedy
created: 2022-08-24T15:34:45Z
updated: 2022-08-25T11:46:19Z
closed: 2022-08-25T11:46:18Z
merged: 2022-08-25T11:46:18Z
base_branch: master
head_branch: vercmp_test
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1080
---

# Pull Request #1080: SPM-1713: fix wrong rpm comparison

**Author**: @psegedy
**Created**: August 24, 2022 at 03:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vercmp_test`

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

### Comment by @psegedy on August 24, 2022 at 04:12 PM UTC

@MichaelMraka I hope it's okay to omit Arch in comparison

### Comment by @codecov-commenter on August 24, 2022 at 04:19 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1080?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.49**% // Head: **61.50**% // Increases project coverage by **`+0.01%`** :tada:
> Coverage data is based on head [(`c59e6ff`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1080?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`4228083`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/42280837f196ab0366ae2a3514d0639564973d9d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1080      +/-   ##
==========================================
+ Coverage   61.49%   61.50%   +0.01%     
==========================================
  Files          97       97              
  Lines        5443     5450       +7     
==========================================
+ Hits         3347     3352       +5     
- Misses       1664     1666       +2     
  Partials      432      432              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.50% <100.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1080?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1080/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `55.31% <100.00%> (+2.81%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1080?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 25, 2022 at 07:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1080*
