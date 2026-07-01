---
type: pull_request
number: 1374
title: "RHINENG-8130: remove v1 and v2 compatibility from /advisories"
state: merged
author: Dugowitch
created: 2024-02-20T15:24:31Z
updated: 2024-03-05T12:55:01Z
closed: 2024-02-27T16:57:18Z
merged: 2024-02-27T16:57:17Z
base_branch: master
head_branch: RHINENG-8130
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1374
---

# Pull Request #1374: RHINENG-8130: remove v1 and v2 compatibility from /advisories

**Author**: @Dugowitch
**Created**: February 20, 2024 at 03:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-8130`

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

### Comment by @jira-linking on February 20, 2024 at 03:24 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8130


### Comment by @codecov-commenter on February 20, 2024 at 03:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `5 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`8fb570d`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8fb570d40d637281fe0fc7b19b68cf70e38a4447?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.97% compared to head [(`1121b04`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.22%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/advisory\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | 80.76% | [3 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1374      +/-   ##
==========================================
+ Coverage   61.97%   62.22%   +0.25%     
==========================================
  Files         108      108              
  Lines        6832     6767      -65     
==========================================
- Hits         4234     4211      -23     
+ Misses       2063     2025      -38     
+ Partials      535      531       -4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.22% <85.71%> (+0.25%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1374?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on February 20, 2024 at 04:40 PM UTC

/retest

### Comment by @niyazRedhat on February 22, 2024 at 06:02 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Commented on February 20, 2024 at 04:42 PM UTC

please take a look if `packagesV1` and functions converting `packagesV1` to `packages` can be removed

### Review by @MichaelMraka - Approved on February 27, 2024 at 02:20 PM UTC

### Review by @psegedy - Approved on February 27, 2024 at 04:57 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1374*
