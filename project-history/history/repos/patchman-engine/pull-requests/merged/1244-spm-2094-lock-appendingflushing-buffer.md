---
type: pull_request
number: 1244
title: "SPM-2094: lock appending/flushing buffer"
state: merged
author: MichaelMraka
created: 2023-06-14T14:18:43Z
updated: 2023-06-14T14:45:04Z
closed: 2023-06-14T14:45:04Z
merged: 2023-06-14T14:45:04Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1244
---

# Pull Request #1244: SPM-2094: lock appending/flushing buffer

**Author**: @MichaelMraka
**Created**: June 14, 2023 at 02:18 PM UTC
**Status**: Merged
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

### Comment by @jira-linking on June 14, 2023 at 02:18 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2094


### Comment by @codecov-commenter on June 14, 2023 at 02:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1244?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`90.90`**% and project coverage change: **`+0.02`** :tada:
> Comparison is base [(`d28a3d9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d28a3d9ee309c05cc95ea6d658ec7d6331612215?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.58% compared to head [(`2eba8c3`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1244?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.60%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1244      +/-   ##
==========================================
+ Coverage   61.58%   61.60%   +0.02%     
==========================================
  Files         105      105              
  Lines        6508     6512       +4     
==========================================
+ Hits         4008     4012       +4     
  Misses       1981     1981              
  Partials      519      519              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.60% <90.90%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1244?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1244?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.40% <90.90%> (+0.25%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1244?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on June 14, 2023 at 02:34 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1244*
