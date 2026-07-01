---
type: pull_request
number: 1086
title: "SPM-1728: Improve logic for running full vmaas sync"
state: merged
author: psegedy
created: 2022-08-30T11:31:53Z
updated: 2022-09-06T15:34:10Z
closed: 2022-09-06T15:34:09Z
merged: 2022-09-06T15:34:09Z
base_branch: master
head_branch: sync_timestamp
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1086
---

# Pull Request #1086: SPM-1728: Improve logic for running full vmaas sync

**Author**: @psegedy
**Created**: August 30, 2022 at 11:31 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sync_timestamp`

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

### Comment by @codecov-commenter on August 30, 2022 at 12:44 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.65**% // Head: **63.05**% // Increases project coverage by **`+0.40%`** :tada:
> Coverage data is based on head [(`5cb6513`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`19754c9`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/19754c9d8805f1bdc6cf029dd7919251f16f1632?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 31.37% of modified lines in pull request are covered.

> :exclamation: Current head 5cb6513 differs from pull request most recent head fbf189e. Consider uploading reports for the commit fbf189e to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1086      +/-   ##
==========================================
+ Coverage   62.65%   63.05%   +0.40%     
==========================================
  Files          97       97              
  Lines        5463     5422      -41     
==========================================
- Hits         3423     3419       -4     
+ Misses       1603     1569      -34     
+ Partials      437      434       -3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.05% <31.37%> (+0.40%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/database.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9kYXRhYmFzZS5nbw==) | `41.17% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `34.61% <ø> (ø)` | |
| [tasks/system\_culling/culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `66.66% <0.00%> (+0.54%)` | :arrow_up: |
| [tasks/vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzLmdv) | `32.05% <0.00%> (+3.64%)` | :arrow_up: |
| [tasks/vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `68.88% <ø> (+7.19%)` | :arrow_up: |
| [tasks/vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `70.76% <ø> (ø)` | |
| [tasks/vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `63.88% <55.00%> (-3.36%)` | :arrow_down: |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1086?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on September 02, 2022 at 09:00 AM UTC

PR / commit  should  increment minor as it's changing API.

### Comment by @MichaelMraka on September 02, 2022 at 09:02 AM UTC

And if there's no ticket for this change please create one (and put it into commit).

### Comment by @jira-linking on September 05, 2022 at 11:05 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1728


### Comment by @psegedy on September 05, 2022 at 11:14 AM UTC

added /minor and created SPM-1728

---

## Reviews

### Review by @MichaelMraka - Approved on September 06, 2022 at 09:29 AM UTC

### Review by @psegedy - Commented on September 06, 2022 at 11:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1086*
