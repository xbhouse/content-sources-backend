---
type: pull_request
number: 1425
title: "RHINENG-10030: handle uuid errors"
state: merged
author: MichaelMraka
created: 2024-05-13T14:05:58Z
updated: 2024-05-15T08:03:47Z
closed: 2024-05-15T08:03:47Z
merged: 2024-05-15T08:03:47Z
base_branch: master
head_branch: pr4
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1425
---

# Pull Request #1425: RHINENG-10030: handle uuid errors

**Author**: @MichaelMraka
**Created**: May 13, 2024 at 02:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr4`

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

### Comment by @jira-linking on May 13, 2024 at 02:06 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10030


### Comment by @codecov-commenter on May 14, 2024 at 11:35 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `81.25000%` with `3 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.78%. Comparing base [(`51311b6`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/51311b65557236abd9139b89858d4cf0059677a0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`1b18379`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 11 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/template\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zLmdv) | 80.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1425      +/-   ##
==========================================
+ Coverage   63.52%   63.78%   +0.25%     
==========================================
  Files         113      113              
  Lines        6863     6892      +29     
==========================================
+ Hits         4360     4396      +36     
+ Misses       1985     1981       -4     
+ Partials      518      515       -3     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.78% <81.25%> (+0.25%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1425?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on May 14, 2024 at 03:05 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1425*
