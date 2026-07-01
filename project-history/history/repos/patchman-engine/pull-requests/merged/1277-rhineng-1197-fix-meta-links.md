---
type: pull_request
number: 1277
title: "RHINENG-1197: fix meta links "
state: merged
author: MichaelMraka
created: 2023-08-01T13:54:42Z
updated: 2023-08-04T08:12:59Z
closed: 2023-08-04T08:12:59Z
merged: 2023-08-04T08:12:59Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1277
---

# Pull Request #1277: RHINENG-1197: fix meta links 

**Author**: @MichaelMraka
**Created**: August 01, 2023 at 01:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @jira-linking on August 01, 2023 at 01:54 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1197


### Comment by @codecov-commenter on August 01, 2023 at 02:11 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`84.82%`** and project coverage change: **`-0.08%`** :warning:
> Comparison is base [(`1790976`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1790976fadceacb3c3e7e39dcb61ee0634adef43?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.54% compared to head [(`48fc1ea`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.47%.
> Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1277      +/-   ##
==========================================
- Coverage   60.54%   60.47%   -0.08%     
==========================================
  Files         106      106              
  Lines        6631     6646      +15     
==========================================
+ Hits         4015     4019       +4     
- Misses       2085     2096      +11     
  Partials      531      531              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.47% <84.82%> (-0.08%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [tasks/caches/refresh\_packages\_caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `57.14% <50.00%> (-6.10%)` | :arrow_down: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.69% <88.09%> (-1.03%)` | :arrow_down: |
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `64.84% <100.00%> (+0.21%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `67.79% <100.00%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `48.97% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.85% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `41.86% <100.00%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `48.59% <100.00%> (ø)` | |
| [manager/controllers/baselines.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `83.78% <100.00%> (ø)` | |
| ... and [10 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1277?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on August 01, 2023 at 03:49 PM UTC

/retest

### Comment by @psegedy on August 02, 2023 at 11:35 AM UTC

/retest

### Comment by @psegedy on August 02, 2023 at 04:00 PM UTC

/retest

### Comment by @MichaelMraka on August 03, 2023 at 09:54 AM UTC

There's couple of invalid filter tests, e.g.
TestSystemsFilterNotExisting
TestSystemsIDsFilterNotExisting

---

## Reviews

### Review by @psegedy - Commented on August 02, 2023 at 01:02 PM UTC

### Review by @psegedy - Commented on August 02, 2023 at 01:04 PM UTC

we should make sure we don't break any checks for invalid filters if we have them now, I'd rely on unit tests and QE to provide an answer whether the functionality is correct or not

if it doesn't break any current checks for invalid filters, then LGTM

PR checks are broken until Brandon and app-sre make necessary changes

### Review by @MichaelMraka - Commented on August 03, 2023 at 09:51 AM UTC

### Review by @psegedy - Approved on August 04, 2023 at 07:36 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1277*
