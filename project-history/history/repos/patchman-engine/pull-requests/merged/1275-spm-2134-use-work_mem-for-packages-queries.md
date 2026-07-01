---
type: pull_request
number: 1275
title: "SPM-2134: use work_mem for packages queries"
state: merged
author: psegedy
created: 2023-07-31T13:59:30Z
updated: 2023-08-01T08:19:13Z
closed: 2023-08-01T08:19:13Z
merged: 2023-08-01T08:19:13Z
base_branch: master
head_branch: work_mem
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1275
---

# Pull Request #1275: SPM-2134: use work_mem for packages queries

**Author**: @psegedy
**Created**: July 31, 2023 at 01:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `work_mem`

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

### Comment by @jira-linking on July 31, 2023 at 01:59 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2134


### Comment by @codecov-commenter on July 31, 2023 at 02:06 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`52.00%`** and project coverage change: **`-0.07%`** :warning:
> Comparison is base [(`34d3ec7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/34d3ec7f82e4476aeb1aede602bda1a490d467fd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.52% compared to head [(`cc8e90e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.45%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1275      +/-   ##
==========================================
- Coverage   60.52%   60.45%   -0.07%     
==========================================
  Files         106      106              
  Lines        6624     6641      +17     
==========================================
+ Hits         4009     4015       +6     
- Misses       2083     2094      +11     
  Partials      532      532              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.45% <52.00%> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [tasks/caches/refresh\_packages\_caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `57.14% <43.75%> (-6.10%)` | :arrow_down: |
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `64.84% <100.00%> (+0.21%)` | :arrow_up: |
| [manager/controllers/packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `53.57% <100.00%> (+7.73%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1275?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 01, 2023 at 07:57 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1275*
