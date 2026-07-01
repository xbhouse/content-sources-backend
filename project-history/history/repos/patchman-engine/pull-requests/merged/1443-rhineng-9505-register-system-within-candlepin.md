---
type: pull_request
number: 1443
title: "RHINENG-9505: register system within candlepin"
state: merged
author: MichaelMraka
created: 2024-06-25T11:33:17Z
updated: 2024-09-10T08:06:25Z
closed: 2024-09-10T08:02:13Z
merged: 2024-09-10T08:02:13Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1443
---

# Pull Request #1443: RHINENG-9505: register system within candlepin

**Author**: @MichaelMraka
**Created**: June 25, 2024 at 11:33 AM UTC
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

### Comment by @jira-linking on June 25, 2024 at 11:33 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9505


### Comment by @niyazRedhat on June 28, 2024 at 02:29 PM UTC

/retest

### Comment by @niyazRedhat on June 28, 2024 at 03:08 PM UTC

/retest

### Comment by @niyazRedhat on June 28, 2024 at 07:09 PM UTC

/retest

### Comment by @mtclinton on June 29, 2024 at 03:16 AM UTC

/retest

### Comment by @mtclinton on July 01, 2024 at 11:58 AM UTC

/retest

### Comment by @codecov-commenter on September 03, 2024 at 01:08 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `74.13793%` with `15 lines` in your changes missing coverage. Please review.
> Project coverage is 65.02%. Comparing base [(`eed2081`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/eed2081a540204510fbc93d37453e1a14f4c5bcf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8c33cc4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8c33cc43c5a086364aa720b712b53ef173589b29?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 11 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/template\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX3VwZGF0ZS5nbw==) | 71.73% | [7 Missing and 6 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_delete.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_delete.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX2RlbGV0ZS5nbw==) | 66.66% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1443      +/-   ##
==========================================
+ Coverage   65.01%   65.02%   +0.01%     
==========================================
  Files         114      114              
  Lines        7823     7877      +54     
==========================================
+ Hits         5086     5122      +36     
- Misses       2207     2218      +11     
- Partials      530      537       +7     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.02% <74.13%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1443?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on September 04, 2024 at 08:43 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Changes Requested on September 05, 2024 at 03:02 PM UTC

### Review by @MichaelMraka - Commented on September 09, 2024 at 08:49 AM UTC

### Review by @jlsherrill - Commented on September 09, 2024 at 12:05 PM UTC

### Review by @MichaelMraka - Commented on September 09, 2024 at 02:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1443*
