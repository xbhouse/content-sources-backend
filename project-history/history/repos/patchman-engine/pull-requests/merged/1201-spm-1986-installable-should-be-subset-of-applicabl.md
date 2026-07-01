---
type: pull_request
number: 1201
title: "SPM-1986: installable should be subset of applicable"
state: merged
author: psegedy
created: 2023-03-28T12:30:11Z
updated: 2023-03-28T16:56:13Z
closed: 2023-03-28T16:56:13Z
merged: 2023-03-28T16:56:13Z
base_branch: master
head_branch: applicable
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1201
---

# Pull Request #1201: SPM-1986: installable should be subset of applicable

**Author**: @psegedy
**Created**: March 28, 2023 at 12:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `applicable`

## Description

- applicable advisory is any advisory which can be applied to a system
- installable advisory is an applicable advisory which can be installed on a system based on the current template `to_date` 
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

### Comment by @jira-linking on March 28, 2023 at 12:30 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1986


### Comment by @codecov-commenter on March 28, 2023 at 12:40 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`82.60`**% and project coverage change: **`+0.16`** :tada:
> Comparison is base [(`2638d25`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2638d25796fbe11560a8925fc8515954720c9ebd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.94% compared to head [(`5d6b225`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.10%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1201      +/-   ##
==========================================
+ Coverage   61.94%   62.10%   +0.16%     
==========================================
  Files         103      103              
  Lines        6231     6239       +8     
==========================================
+ Hits         3860     3875      +15     
+ Misses       1875     1868       -7     
  Partials      496      496              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.10% <82.60%> (+0.16%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.46% <0.00%> (-0.05%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `71.97% <90.00%> (+1.02%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.55% <100.00%> (+2.43%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.07% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1201?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @michalslomczynski - Approved on March 28, 2023 at 03:54 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1201*
