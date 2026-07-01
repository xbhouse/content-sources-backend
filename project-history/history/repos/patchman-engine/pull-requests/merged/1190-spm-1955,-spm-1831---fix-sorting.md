---
type: pull_request
number: 1190
title: "SPM-1955, SPM-1831 - fix sorting"
state: merged
author: psegedy
created: 2023-03-14T12:24:47Z
updated: 2023-03-14T14:45:25Z
closed: 2023-03-14T14:45:25Z
merged: 2023-03-14T14:45:25Z
base_branch: master
head_branch: collation
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1190
---

# Pull Request #1190: SPM-1955, SPM-1831 - fix sorting

**Author**: @psegedy
**Created**: March 14, 2023 at 12:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `collation`

## Description

- use collation to use natural sort for alphanumerical strings (OS) - see commit message
- make nulls last when sorting, based on SPM-1831 it is the common way of sorting by all apps
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

### Comment by @jira-linking on March 14, 2023 at 12:24 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1955
https://issues.redhat.com/browse/SPM-1831


### Comment by @codecov-commenter on March 14, 2023 at 01:48 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1190?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and no project coverage change.
> Comparison is base [(`d7077e5`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d7077e58dad3c39ac5cc160d3ea30be27624f694?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.90% compared to head [(`d73fedf`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1190?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.90%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1190   +/-   ##
=======================================
  Coverage   62.90%   62.90%           
=======================================
  Files         103      103           
  Lines        6066     6066           
=======================================
  Hits         3816     3816           
  Misses       1777     1777           
  Partials      473      473           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.90% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1190?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1190?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.37% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1190?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 14, 2023 at 02:39 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1190*
