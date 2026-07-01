---
type: pull_request
number: 1212
title: "SPM-2013: process accounts from the same system_account partition together"
state: merged
author: MichaelMraka
created: 2023-04-17T11:17:26Z
updated: 2023-04-18T10:31:09Z
closed: 2023-04-18T10:31:09Z
merged: 2023-04-18T10:31:09Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1212
---

# Pull Request #1212: SPM-2013: process accounts from the same system_account partition together

**Author**: @MichaelMraka
**Created**: April 17, 2023 at 11:17 AM UTC
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

### Comment by @jira-linking on April 17, 2023 at 11:17 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2013


### Comment by @codecov-commenter on April 17, 2023 at 11:24 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`60.00`**% and project coverage change: **`-0.02`** :warning:
> Comparison is base [(`60c4f81`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/60c4f814ef08513079379044ada74a8420b4b6a8?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.04% compared to head [(`e3d9d13`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.02%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1212      +/-   ##
==========================================
- Coverage   62.04%   62.02%   -0.02%     
==========================================
  Files         103      103              
  Lines        6283     6286       +3     
==========================================
+ Hits         3898     3899       +1     
- Misses       1886     1887       +1     
- Partials      499      500       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.02% <60.00%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `75.00% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `43.83% <50.00%> (-0.46%)` | :arrow_down: |
| [tasks/caches/refresh\_advisory\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfYWR2aXNvcnlfY2FjaGVzLmdv) | `48.71% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1212?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on April 18, 2023 at 09:36 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on April 18, 2023 at 10:30 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1212*
