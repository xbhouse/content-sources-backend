---
type: pull_request
number: 1064
title: "SPM-1705: Add OrgID to eval recalc messages (no refactor)"
state: merged
author: michalslomczynski
created: 2022-08-18T09:48:16Z
updated: 2022-08-18T13:58:30Z
closed: 2022-08-18T13:58:30Z
merged: 2022-08-18T13:58:30Z
base_branch: master
head_branch: vmass-sync-orgid-no-refactor
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1064
---

# Pull Request #1064: SPM-1705: Add OrgID to eval recalc messages (no refactor)

**Author**: @michalslomczynski
**Created**: August 18, 2022 at 09:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmass-sync-orgid-no-refactor`

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

### Comment by @codecov-commenter on August 18, 2022 at 09:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.41**% // Head: **61.36**% // Decreases project coverage by **`-0.04%`** :warning:
> Coverage data is based on head [(`a263d20`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`4b2ead3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b2ead3ccb5a0dbdaf79148ff0e66187eb017e52?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 16.66% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1064      +/-   ##
==========================================
- Coverage   61.41%   61.36%   -0.05%     
==========================================
  Files          98       98              
  Lines        5401     5407       +6     
==========================================
+ Hits         3317     3318       +1     
- Misses       1658     1663       +5     
  Partials      426      426              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.36% <16.66%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.33% <0.00%> (-0.50%)` | :arrow_down: |
| [tasks/vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `33.33% <0.00%> (-1.29%)` | :arrow_down: |
| [tasks/vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `70.76% <100.00%> (+0.45%)` | :arrow_up: |

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1064?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 18, 2022 at 01:44 PM UTC

### Review by @psegedy - Approved on August 18, 2022 at 01:49 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1064*
