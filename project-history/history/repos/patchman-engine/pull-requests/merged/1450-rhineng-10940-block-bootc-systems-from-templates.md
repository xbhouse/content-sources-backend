---
type: pull_request
number: 1450
title: "RHINENG-10940: block bootc systems from templates"
state: merged
author: psegedy
created: 2024-07-10T12:11:02Z
updated: 2024-07-12T13:22:22Z
closed: 2024-07-12T13:22:18Z
merged: 2024-07-12T13:22:18Z
base_branch: master
head_branch: bootc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1450
---

# Pull Request #1450: RHINENG-10940: block bootc systems from templates

**Author**: @psegedy
**Created**: July 10, 2024 at 12:11 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bootc`

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

### Comment by @jira-linking on July 10, 2024 at 12:11 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10940


### Comment by @psegedy on July 10, 2024 at 06:45 PM UTC

/retest

### Comment by @psegedy on July 11, 2024 at 07:23 AM UTC

/retest

### Comment by @psegedy on July 11, 2024 at 10:34 AM UTC

still fails to deploy content-sources
```
clowdapp/content-sources-backend not ready, status conditions
```


### Comment by @psegedy on July 11, 2024 at 10:34 AM UTC

/retest

### Comment by @MichaelMraka on July 12, 2024 at 09:03 AM UTC

/retest

### Comment by @codecov-commenter on July 12, 2024 at 10:12 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `61.01695%` with `23 lines` in your changes missing coverage. Please review.
> Project coverage is 64.86%. Comparing base [(`4606ea9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4606ea9ec5fcdd8c9c134c3f2610bb9963bc5ff2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`40049b4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/40049b404fd8e7282d5d71d889af459f72ccf2b0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 4 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/baseline\_create.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fbaseline_create.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | 62.50% | [11 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX3VwZGF0ZS5nbw==) | 45.45% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/baseline\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fbaseline_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | 72.72% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 60.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1450      +/-   ##
==========================================
+ Coverage   64.81%   64.86%   +0.04%     
==========================================
  Files         114      114              
  Lines        7777     7798      +21     
==========================================
+ Hits         5041     5058      +17     
- Misses       2206     2214       +8     
+ Partials      530      526       -4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.86% <61.01%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1450?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on July 11, 2024 at 02:07 PM UTC

### Review by @MichaelMraka - Approved on July 12, 2024 at 11:18 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1450*
