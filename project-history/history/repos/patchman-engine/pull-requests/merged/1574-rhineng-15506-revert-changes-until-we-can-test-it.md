---
type: pull_request
number: 1574
title: "RHINENG-15506: revert changes until we can test it "
state: merged
author: MichaelMraka
created: 2025-02-11T16:15:17Z
updated: 2025-02-12T08:28:43Z
closed: 2025-02-12T08:28:43Z
merged: 2025-02-12T08:28:43Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1574
---

# Pull Request #1574: RHINENG-15506: revert changes until we can test it 

**Author**: @MichaelMraka
**Created**: February 11, 2025 at 04:15 PM UTC
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

### Comment by @jira-linking on February 11, 2025 at 04:15 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-15506


### Comment by @codecov-commenter on February 11, 2025 at 04:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `52.00000%` with `24 lines` in your changes missing coverage. Please review.
> Project coverage is 57.94%. Comparing base [(`9c8f129`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9c8f129277c84aa934d9d7ab9b9a077615907d21?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`f90fd5a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f90fd5acda0e8ae0448cc06c137de736f014702f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [platform/candlepin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?src=pr&el=tree&filepath=platform%2Fcandlepin.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-cGxhdGZvcm0vY2FuZGxlcGluLmdv) | 0.00% | [16 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX3VwZGF0ZS5nbw==) | 72.41% | [5 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1574      +/-   ##
==========================================
- Coverage   58.03%   57.94%   -0.09%     
==========================================
  Files         143      143              
  Lines       11146    11158      +12     
==========================================
- Hits         6469     6466       -3     
- Misses       4097     4110      +13     
- Partials      580      582       +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.94% <52.00%> (-0.09%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1574?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on February 11, 2025 at 04:43 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1574*
