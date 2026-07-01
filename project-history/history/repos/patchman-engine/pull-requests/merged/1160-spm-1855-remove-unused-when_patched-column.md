---
type: pull_request
number: 1160
title: "SPM-1855: remove unused when_patched column"
state: merged
author: MichaelMraka
created: 2023-01-19T15:10:00Z
updated: 2023-01-20T12:44:19Z
closed: 2023-01-20T12:44:19Z
merged: 2023-01-20T12:44:19Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1160
---

# Pull Request #1160: SPM-1855: remove unused when_patched column

**Author**: @MichaelMraka
**Created**: January 19, 2023 at 03:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

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

### Comment by @jira-linking on January 19, 2023 at 03:10 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1855


### Comment by @codecov-commenter on January 19, 2023 at 04:07 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.18**% // Head: **63.27**% // Increases project coverage by **`+0.08%`** :tada:
> Coverage data is based on head [(`bd051c8`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`a42c9cd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a42c9cdb8d1171a5dabefb602ca39bb253c1f9bc?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 86.84% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1160      +/-   ##
==========================================
+ Coverage   63.18%   63.27%   +0.08%     
==========================================
  Files         101      101              
  Lines        5772     5745      -27     
==========================================
- Hits         3647     3635      -12     
+ Misses       1667     1652      -15     
  Partials      458      458              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.27% <86.84%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.54% <0.00%> (+0.58%)` | :arrow_up: |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `74.79% <100.00%> (-1.90%)` | :arrow_down: |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `75.23% <100.00%> (-0.47%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1160?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on January 20, 2023 at 09:48 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1160*
