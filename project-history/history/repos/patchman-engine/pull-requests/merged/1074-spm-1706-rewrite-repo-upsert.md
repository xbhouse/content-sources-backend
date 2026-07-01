---
type: pull_request
number: 1074
title: "SPM-1706: rewrite repo upsert"
state: merged
author: psegedy
created: 2022-08-22T12:27:04Z
updated: 2022-08-29T12:04:34Z
closed: 2022-08-29T12:04:34Z
merged: 2022-08-29T12:04:34Z
base_branch: master
head_branch: upsert-repo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1074
---

# Pull Request #1074: SPM-1706: rewrite repo upsert

**Author**: @psegedy
**Created**: August 22, 2022 at 12:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `upsert-repo`

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

### Comment by @codecov-commenter on August 23, 2022 at 11:17 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.59**% // Head: **61.54**% // Decreases project coverage by **`-0.04%`** :warning:
> Coverage data is based on head [(`27190e3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`485b23a`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/485b23a4a62f19ee1a8b48ccfab58dcff7291c94?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 85.71% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1074      +/-   ##
==========================================
- Coverage   61.59%   61.54%   -0.05%     
==========================================
  Files          97       97              
  Lines        5411     5451      +40     
==========================================
+ Hits         3333     3355      +22     
- Misses       1651     1664      +13     
- Partials      427      432       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.54% <85.71%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.61% <85.71%> (+0.57%)` | :arrow_up: |
| [tasks/vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `61.69% <0.00%> (-3.40%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1074?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 24, 2022 at 12:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1074*
