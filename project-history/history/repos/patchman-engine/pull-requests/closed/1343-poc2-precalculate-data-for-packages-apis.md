---
type: pull_request
number: 1343
title: "POC2: precalculate data for /packages APIs"
state: closed
author: MichaelMraka
created: 2023-11-20T15:39:19Z
updated: 2024-01-08T15:08:11Z
closed: 2024-01-08T15:08:11Z
base_branch: master
head_branch: poc2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1343
---

# Pull Request #1343: POC2: precalculate data for /packages APIs

**Author**: @MichaelMraka
**Created**: November 20, 2023 at 03:39 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `poc2`

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

### Comment by @jira-linking on November 20, 2023 at 03:39 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2281
https://issues.redhat.com/browse/RHINENG-2282
https://issues.redhat.com/browse/RHINENG-2283
https://issues.redhat.com/browse/RHINENG-2284


### Comment by @codecov-commenter on November 22, 2023 at 12:32 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `113 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`997b6bc`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/997b6bc3100a1312134c476be5f9309a68537f30?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.00% compared to head [(`a85afef`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.32%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [43 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 64.34% | [31 Missing and 10 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [18 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | 72.72% | [4 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | 75.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | 91.66% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/package\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXNfZXhwb3J0Lmdv) | 75.00% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [tasks/caches/refresh\_packages\_caches.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | 88.88% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1343      +/-   ##
==========================================
- Coverage   62.00%   61.32%   -0.68%     
==========================================
  Files         108      108              
  Lines        6811     6855      +44     
==========================================
- Hits         4223     4204      -19     
- Misses       2054     2113      +59     
- Partials      534      538       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.32% <53.49%> (-0.68%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1343?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on January 08, 2024 at 03:08 PM UTC

This POC did not meet expectations, the evaluation part is significantly slower and small improvement in speed of /packages API is not going to overweight it.
Closing without merge into master.  

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1343*
