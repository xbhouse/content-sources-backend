---
type: pull_request
number: 1243
title: "update dependencies, use GOMEMLIMIT and automaxprocs"
state: merged
author: psegedy
created: 2023-06-12T15:07:46Z
updated: 2023-06-13T07:53:31Z
closed: 2023-06-13T07:53:31Z
merged: 2023-06-13T07:53:31Z
base_branch: master
head_branch: go119
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1243
---

# Pull Request #1243: update dependencies, use GOMEMLIMIT and automaxprocs

**Author**: @psegedy
**Created**: June 12, 2023 at 03:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `go119`

## Description

GOMEMLIMIT needs to be set properly for Stage and Prod env since they have different memory limits

the rule of thumb is to set GOMEMLIMIT as 90% of the memory limit because it is just a soft limit which can be exceeded
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

### Comment by @jira-linking on June 12, 2023 at 03:07 PM UTC

Commits missing Jira IDs:
1b3187e993d6f0912d7171a705607e72e6214ec7
452ce94c3da0d052239265684834c644453e484b
0fc72043b94c334fe72af664e78965b09e77b153
2496929b8ae27f5fd1ee2f5c354d4733283895c6


### Comment by @codecov-commenter on June 12, 2023 at 03:18 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`56.06`**% and project coverage change: **`-0.04`** :warning:
> Comparison is base [(`12c1603`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/12c1603b8ce434162bd6334fb57a33498778828e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.62% compared to head [(`2496929`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.58%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1243      +/-   ##
==========================================
- Coverage   61.62%   61.58%   -0.04%     
==========================================
  Files         105      105              
  Lines        6475     6508      +33     
==========================================
+ Hits         3990     4008      +18     
- Misses       1968     1981      +13     
- Partials      517      519       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.58% <56.06%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/http.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `31.57% <0.00%> (-3.04%)` | :arrow_down: |
| [evaluator/vmaas\_cache.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `61.76% <0.00%> (-1.88%)` | :arrow_down: |
| [listener/listener.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `66.66% <0.00%> (-2.09%)` | :arrow_down: |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.78% <50.00%> (-1.95%)` | :arrow_down: |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `62.68% <50.00%> (ø)` | |
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `64.63% <90.00%> (+3.30%)` | :arrow_up: |
| [base/database/batch.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9iYXRjaC5nbw==) | `58.82% <100.00%> (+0.14%)` | :arrow_up: |
| [manager/controllers/baseline\_create.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `76.71% <100.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `79.79% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1243?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 13, 2023 at 07:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1243*
