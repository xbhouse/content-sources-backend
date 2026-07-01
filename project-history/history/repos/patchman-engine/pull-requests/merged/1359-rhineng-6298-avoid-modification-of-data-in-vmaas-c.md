---
type: pull_request
number: 1359
title: "RHINENG-6298: avoid modification of data in vmaas cache"
state: merged
author: psegedy
created: 2024-01-03T16:44:37Z
updated: 2024-01-04T10:39:21Z
closed: 2024-01-04T10:39:21Z
merged: 2024-01-04T10:39:21Z
base_branch: master
head_branch: deepcopy_vmaas_cache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1359
---

# Pull Request #1359: RHINENG-6298: avoid modification of data in vmaas cache

**Author**: @psegedy
**Created**: January 03, 2024 at 04:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `deepcopy_vmaas_cache`

## Description

Vmaas cache can cause incorrect evaluation of systems with same package list

1. upload 1st system
2. assign template to 1st system
3. vmaas response is modified and proper installable/applicable advisories are set according to template, response is stored into cache
4. upload 2nd system with the same package list as 1st system
5. set of advisories is obtained from cache but it shows incorrect installable advisories because 2nd system is not in template and it uses cached data for 1st system

I used github.com/jinzhu/copier because our struct contains pointers and slices, we need to make a deepcopy, it isn't sufficient to just copy struct or store struct directly to the cache instead of pointer to the struct.

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

### Comment by @jira-linking on January 03, 2024 at 04:44 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-6298


### Comment by @psegedy on January 03, 2024 at 06:54 PM UTC

/retest

### Comment by @codecov-commenter on January 03, 2024 at 07:13 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `4 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`c98bc06`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c98bc06cea4542a6991d4b8534a11636e40fadd1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.09% compared to head [(`cc7afed`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.09%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 60.00% | [2 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1359   +/-   ##
=======================================
  Coverage   62.09%   62.09%           
=======================================
  Files         108      108           
  Lines        6801     6809    +8     
=======================================
+ Hits         4223     4228    +5     
- Misses       2044     2045    +1     
- Partials      534      536    +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.09% <60.00%> (+<0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1359?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 04, 2024 at 09:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1359*
