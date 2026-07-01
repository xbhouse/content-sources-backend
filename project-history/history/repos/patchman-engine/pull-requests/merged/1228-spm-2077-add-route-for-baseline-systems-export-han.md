---
type: pull_request
number: 1228
title: "SPM-2077: add route for baseline systems export handler"
state: merged
author: psegedy
created: 2023-05-30T14:31:29Z
updated: 2023-05-31T14:54:14Z
closed: 2023-05-31T14:54:14Z
merged: 2023-05-31T14:54:14Z
base_branch: master
head_branch: exportbaselines2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1228
---

# Pull Request #1228: SPM-2077: add route for baseline systems export handler

**Author**: @psegedy
**Created**: May 30, 2023 at 02:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `exportbaselines2`

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

### Comment by @jira-linking on May 30, 2023 at 02:31 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2077


### Comment by @codecov-commenter on May 30, 2023 at 02:43 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`48.38`**% and project coverage change: **`-0.02`** :warning:
> Comparison is base [(`8e0407c`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8e0407cd8ad13ce3da059ed49bcbd9b75395065d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.14% compared to head [(`22708ef`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.12%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1228      +/-   ##
==========================================
- Coverage   62.14%   62.12%   -0.02%     
==========================================
  Files         103      104       +1     
  Lines        6321     6348      +27     
==========================================
+ Hits         3928     3944      +16     
- Misses       1894     1901       +7     
- Partials      499      503       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.12% <48.38%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/test\_utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100.00% <ø> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `50.49% <37.50%> (+2.04%)` | :arrow_up: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `83.61% <50.00%> (-0.33%)` | :arrow_down: |
| [manager/controllers/baseline\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `52.63% <52.63%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1228?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on May 31, 2023 at 01:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1228*
