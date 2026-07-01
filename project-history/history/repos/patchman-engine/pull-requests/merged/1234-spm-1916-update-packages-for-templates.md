---
type: pull_request
number: 1234
title: "SPM-1916: update /packages for templates"
state: merged
author: psegedy
created: 2023-06-02T15:14:45Z
updated: 2023-06-05T11:38:06Z
closed: 2023-06-05T11:38:05Z
merged: 2023-06-05T11:38:05Z
base_branch: master
head_branch: packages_api
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1234
---

# Pull Request #1234: SPM-1916: update /packages for templates

**Author**: @psegedy
**Created**: June 02, 2023 at 03:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `packages_api`

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

### Comment by @jira-linking on June 02, 2023 at 03:14 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1916


### Comment by @codecov-commenter on June 02, 2023 at 03:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`9.09`**% and project coverage change: **`-0.16`** :warning:
> Comparison is base [(`b7a63dc`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b7a63dc6761adbd97e5e063d16f94e1eae6b468f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.77% compared to head [(`0549605`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.62%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1234      +/-   ##
==========================================
- Coverage   61.77%   61.62%   -0.16%     
==========================================
  Files         105      105              
  Lines        6456     6475      +19     
==========================================
+ Hits         3988     3990       +2     
- Misses       1953     1968      +15     
- Partials      515      517       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.62% <9.09%> (-0.16%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [tasks/caches/refresh\_packages\_caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `62.68% <7.69%> (-13.24%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `43.47% <20.00%> (-6.53%)` | :arrow_down: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1234?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 02, 2023 at 03:42 PM UTC

tests are failing because of removed `systems_updatable` filter

---

## Reviews

### Review by @MichaelMraka - Approved on June 05, 2023 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1234*
