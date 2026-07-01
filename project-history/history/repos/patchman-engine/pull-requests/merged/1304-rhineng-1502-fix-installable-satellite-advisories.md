---
type: pull_request
number: 1304
title: "RHINENG-1502: fix installable satellite advisories by always using epoch"
state: merged
author: psegedy
created: 2023-08-29T14:20:10Z
updated: 2023-08-29T15:55:42Z
closed: 2023-08-29T15:55:42Z
merged: 2023-08-29T15:55:42Z
base_branch: master
head_branch: fix_installable
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1304
---

# Pull Request #1304: RHINENG-1502: fix installable satellite advisories by always using epoch

**Author**: @psegedy
**Created**: August 29, 2023 at 02:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_installable`

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

### Comment by @jira-linking on August 29, 2023 at 02:20 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1502


### Comment by @codecov-commenter on August 29, 2023 at 02:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00%`** and no project coverage change.
> Comparison is base [(`3a9ea57`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3a9ea57d1d4597b0054fa389bd38b31c8b4ade17?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.23% compared to head [(`3b885a2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.23%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1304   +/-   ##
=======================================
  Coverage   62.23%   62.23%           
=======================================
  Files         106      106           
  Lines        6684     6684           
=======================================
  Hits         4160     4160           
  Misses       1996     1996           
  Partials      528      528           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.23% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.53% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1304?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on August 29, 2023 at 03:55 PM UTC

tests passed but didn't update status in PR

---

## Reviews

### Review by @yungbender - Approved on August 29, 2023 at 03:21 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1304*
