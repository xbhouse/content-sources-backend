---
type: pull_request
number: 1251
title: "SPM-2126: add metrics for package_refresh job"
state: merged
author: psegedy
created: 2023-06-27T15:18:16Z
updated: 2023-06-30T08:47:27Z
closed: 2023-06-30T08:47:26Z
merged: 2023-06-30T08:47:26Z
base_branch: master
head_branch: asdfrefreshjob
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1251
---

# Pull Request #1251: SPM-2126: add metrics for package_refresh job

**Author**: @psegedy
**Created**: June 27, 2023 at 03:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `asdfrefreshjob`

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

### Comment by @jira-linking on June 27, 2023 at 03:18 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2126


### Comment by @codecov-commenter on June 27, 2023 at 03:26 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.21`** :warning:
> Comparison is base [(`16f6f22`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/16f6f22d88ceadcd0073738de95496775ab79206?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.51% compared to head [(`1e20e99`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.31%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1251      +/-   ##
==========================================
- Coverage   61.51%   61.31%   -0.21%     
==========================================
  Files         105      106       +1     
  Lines        6515     6537      +22     
==========================================
  Hits         4008     4008              
- Misses       1988     2010      +22     
  Partials      519      519              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.31% <0.00%> (-0.21%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [tasks/caches/caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL2NhY2hlcy5nbw==) | `18.18% <0.00%> (-4.05%)` | :arrow_down: |
| [tasks/caches/metrics.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL21ldHJpY3MuZ28=) | `0.00% <0.00%> (ø)` | |
| [tasks/caches/refresh\_packages\_caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1251?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 29, 2023 at 09:48 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1251*
