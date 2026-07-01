---
type: pull_request
number: 1114
title: "Hotfix force schema fix"
state: closed
author: MichaelMraka
created: 2022-09-27T11:54:57Z
updated: 2022-10-03T12:15:33Z
closed: 2022-10-03T12:15:33Z
base_branch: master
head_branch: hotfix-force-schema-fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1114
---

# Pull Request #1114: Hotfix force schema fix

**Author**: @MichaelMraka
**Created**: September 27, 2022 at 11:54 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `hotfix-force-schema-fix`

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

### Comment by @jira-linking on September 27, 2022 at 11:54 AM UTC

Commits missing Jira IDs:
573588545fda72e44e9e9120bdb419ff9741de04
96403d77a09ffc98d2d1362ccbd3cc713fd21e7f
afbac60107678c768cb8b8a472f29c4c624ce46e
7496f3c5aa08a52a01e343885eaafd1566f1896f
8c79dff2aed98fcbc204d7a4f85a722b086de6b7
24555b3e5967f8d1ed238f931650e365210ee7ac
0056b331c300bd9fc96ce4dbe79b951f5024ed37
259c7bfb46d717552620fcc22aba5cdaa1d4d47f
b8ec3af09bfd62621a530ba06fbd3b5464e59ba9
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1661
https://issues.redhat.com/browse/SPM-1745
https://issues.redhat.com/browse/SPM-1737
https://issues.redhat.com/browse/SPM-1753
https://issues.redhat.com/browse/SPM-1746
https://issues.redhat.com/browse/SPM-1744


### Comment by @codecov-commenter on September 27, 2022 at 12:09 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.81**% // Head: **63.05**% // Increases project coverage by **`+0.24%`** :tada:
> Coverage data is based on head [(`99b55fe`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`86a6982`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/86a69825cf90e62955de8c81c59c07bc1ee165be?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1114      +/-   ##
==========================================
+ Coverage   62.81%   63.05%   +0.24%     
==========================================
  Files          98       97       -1     
  Lines        5615     5587      -28     
==========================================
- Hits         3527     3523       -4     
+ Misses       1635     1613      -22     
+ Partials      453      451       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.05% <100.00%> (+0.24%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.38% <ø> (-0.05%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `84.37% <ø> (+13.32%)` | :arrow_up: |
| [tasks/caches/caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL2NhY2hlcy5nbw==) | `37.50% <ø> (+19.85%)` | :arrow_up: |
| [tasks/caches/refresh\_advisory\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfYWR2aXNvcnlfY2FjaGVzLmdv) | `46.66% <100.00%> (+3.80%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1114?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1114*
