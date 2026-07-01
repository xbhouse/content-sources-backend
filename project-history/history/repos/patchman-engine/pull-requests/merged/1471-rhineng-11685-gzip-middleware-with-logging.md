---
type: pull_request
number: 1471
title: "RHINENG-11685: gzip middleware with logging"
state: merged
author: psegedy
created: 2024-08-26T09:57:53Z
updated: 2024-08-28T07:40:57Z
closed: 2024-08-28T07:40:57Z
merged: 2024-08-28T07:40:57Z
base_branch: master
head_branch: gzip
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1471
---

# Pull Request #1471: RHINENG-11685: gzip middleware with logging

**Author**: @psegedy
**Created**: August 26, 2024 at 09:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `gzip`

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

### Comment by @jira-linking on August 26, 2024 at 09:57 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11685


### Comment by @codecov-commenter on August 26, 2024 at 10:02 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `20.00000%` with `8 lines` in your changes missing coverage. Please review.
> Project coverage is 64.97%. Comparing base [(`827a992`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/827a9926d50880235f713e720430311ee80973c4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`21087ae`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/21087ae59ea75a41634130b90ce21ea3c701d09d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/middlewares/gzip.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Fgzip.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9nemlwLmdv) | 0.00% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | 50.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1471   +/-   ##
=======================================
  Coverage   64.96%   64.97%           
=======================================
  Files         114      115    +1     
  Lines        7833     7837    +4     
=======================================
+ Hits         5089     5092    +3     
- Misses       2213     2214    +1     
  Partials      531      531           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.97% <20.00%> (+<0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1471?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @Dugowitch - Approved on August 26, 2024 at 10:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1471*
