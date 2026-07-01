---
type: pull_request
number: 1220
title: "RHINENG-390: skip hosts without epoch in package_list"
state: closed
author: psegedy
created: 2023-05-16T11:27:44Z
updated: 2023-05-17T09:35:45Z
closed: 2023-05-17T09:35:45Z
base_branch: master
head_branch: epoch
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1220
---

# Pull Request #1220: RHINENG-390: skip hosts without epoch in package_list

**Author**: @psegedy
**Created**: May 16, 2023 at 11:27 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `epoch`

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

### Comment by @jira-linking on May 16, 2023 at 11:27 AM UTC

Commits missing Jira IDs:
2f87d41bfc3d36c29b2ebce9e57b0b4410ea8b66
5b87fc8a5ef46e005e03324c8032e14d8d9d97cd
93bd696b7cdb098da8455a0fe5409a70d2364190
e1ebfc6073ff8b4125a669f1ecaa265ba891abb1


### Comment by @codecov-commenter on May 16, 2023 at 06:46 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`52.63`**% and project coverage change: **`-0.01`** :warning:
> Comparison is base [(`f0aa0d1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f0aa0d1d16ae0f4b9bb892a0ab2844c3ae87847b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98% compared to head [(`3502e86`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1220      +/-   ##
==========================================
- Coverage   61.98%   61.98%   -0.01%     
==========================================
  Files         103      103              
  Lines        6303     6318      +15     
==========================================
+ Hits         3907     3916       +9     
- Misses       1896     1899       +3     
- Partials      500      503       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.98% <52.63%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `55.31% <25.00%> (ø)` | |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.91% <33.33%> (-0.64%)` | :arrow_down: |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.00% <71.42%> (-0.14%)` | :arrow_down: |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.97% <100.00%> (+0.19%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1220?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1220*
