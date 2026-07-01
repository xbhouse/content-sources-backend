---
type: pull_request
number: 1162
title: "VULN-2519: don't send empty repository_list to vmaas"
state: merged
author: psegedy
created: 2023-01-19T16:52:30Z
updated: 2023-01-25T17:04:44Z
closed: 2023-01-25T17:04:43Z
merged: 2023-01-25T17:04:43Z
base_branch: master
head_branch: null_repolist
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1162
---

# Pull Request #1162: VULN-2519: don't send empty repository_list to vmaas

**Author**: @psegedy
**Created**: January 19, 2023 at 04:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `null_repolist`

## Description

- simplify repository_list, it can be []string instead of *[]string
- don't send request to vmaas when system has no repos

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

### Comment by @jira-linking on January 19, 2023 at 04:52 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/VULN-2519


### Comment by @codecov-commenter on January 19, 2023 at 05:01 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.17**% // Head: **63.14**% // Decreases project coverage by **`-0.04%`** :warning:
> Coverage data is based on head [(`f70d44b`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`9ce9b5f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/9ce9b5fd1fa8b214ba3351d9e16d4a34fcdef275?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 50.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1162      +/-   ##
==========================================
- Coverage   63.17%   63.14%   -0.04%     
==========================================
  Files         102      102              
  Lines        5755     5758       +3     
==========================================
  Hits         3636     3636              
- Misses       1661     1663       +2     
- Partials      458      459       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.14% <50.00%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.01% <0.00%> (-0.65%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.14% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1162?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 25, 2023 at 11:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1162*
