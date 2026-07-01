---
type: pull_request
number: 1286
title: "SPM-2105: store in DB if system was uploaded with cache refreshed"
state: merged
author: yungbender
created: 2023-08-10T13:57:53Z
updated: 2023-08-11T16:13:55Z
closed: 2023-08-11T16:13:55Z
merged: 2023-08-11T16:13:55Z
base_branch: master
head_branch: build_packagecache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1286
---

# Pull Request #1286: SPM-2105: store in DB if system was uploaded with cache refreshed

**Author**: @yungbender
**Created**: August 10, 2023 at 01:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `build_packagecache`

## Description

- Past tense of `build_pkgcache` sounds better to me in db, so I named it `built_pkgcache`

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

### Comment by @jira-linking on August 10, 2023 at 01:57 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2105


### Comment by @codecov-commenter on August 11, 2023 at 09:26 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`85.71%`** and project coverage change: **`+0.08%`** :tada:
> Comparison is base [(`9d761af`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9d761af9274bbf9921c7d7079caced72f86092df?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.70% compared to head [(`543d690`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.79%.
> Report is 6 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1286      +/-   ##
==========================================
+ Coverage   61.70%   61.79%   +0.08%     
==========================================
  Files         106      106              
  Lines        6651     6666      +15     
==========================================
+ Hits         4104     4119      +15     
  Misses       2013     2013              
  Partials      534      534              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.79% <85.71%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.30% <85.71%> (+0.36%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1286?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 10, 2023 at 02:37 PM UTC

### Review by @psegedy - Changes Requested on August 10, 2023 at 02:40 PM UTC

### Review by @psegedy - Approved on August 11, 2023 at 04:13 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1286*
