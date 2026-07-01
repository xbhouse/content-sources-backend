---
type: pull_request
number: 1553
title: "chore: update go version"
state: merged
author: psegedy
created: 2025-01-10T17:43:17Z
updated: 2025-01-14T14:51:22Z
closed: 2025-01-14T14:51:02Z
merged: 2025-01-14T14:51:02Z
base_branch: master
head_branch: go122
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1553
---

# Pull Request #1553: chore: update go version

**Author**: @psegedy
**Created**: January 10, 2025 at 05:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `go122`

## Description

RHINENG-15257

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

### Comment by @jira-linking on January 10, 2025 at 05:43 PM UTC

Commits missing Jira IDs:
9863436ebdb9645a1dc122c89a77e2fef521ea94
6509f5ef50c097d4cc1da576a1389e48e1f0b7ed
d0987f3847c153fb60f90235dcd125edea08264f
d4379cb2d24cc590924da7e7de81e30eb2608871
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-15257


### Comment by @codecov-commenter on January 13, 2025 at 11:25 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `16.66667%` with `10 lines` in your changes missing coverage. Please review.
> Project coverage is 57.96%. Comparing base [(`251a057`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/251a05773b110d8503ee6069a18fbc4a959dd9a6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`d4379cb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d4379cb2d24cc590924da7e7de81e30eb2608871?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/limits.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Flimits.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9saW1pdHMuZ28=) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/mqueue/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&filepath=base%2Fmqueue%2Ftesting.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1553      +/-   ##
==========================================
- Coverage   57.97%   57.96%   -0.01%     
==========================================
  Files         143      143              
  Lines       11135    11131       -4     
==========================================
- Hits         6455     6452       -3     
+ Misses       4101     4100       -1     
  Partials      579      579              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.96% <16.66%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1553?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 13, 2025 at 12:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1553*
