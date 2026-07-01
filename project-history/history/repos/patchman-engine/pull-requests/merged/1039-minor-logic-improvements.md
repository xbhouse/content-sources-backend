---
type: pull_request
number: 1039
title: "Minor logic improvements"
state: merged
author: michalslomczynski
created: 2022-07-28T09:34:58Z
updated: 2022-08-31T11:47:31Z
closed: 2022-08-31T11:47:31Z
merged: 2022-08-31T11:47:31Z
base_branch: master
head_branch: style-improvements
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1039
---

# Pull Request #1039: Minor logic improvements

**Author**: @michalslomczynski
**Created**: July 28, 2022 at 09:34 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `style-improvements`

## Description

- First one is always true
- Second one is never nil

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

### Comment by @jira-linking on July 28, 2022 at 09:35 AM UTC

Commits missing Jira IDs:
07f37b907be87a0e37cf4154f6c38d0614fd538b


### Comment by @codecov-commenter on August 31, 2022 at 11:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.65**% // Head: **62.81**% // Increases project coverage by **`+0.15%`** :tada:
> Coverage data is based on head [(`d353657`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`19754c9`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/19754c9d8805f1bdc6cf029dd7919251f16f1632?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 11.76% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1039      +/-   ##
==========================================
+ Coverage   62.65%   62.81%   +0.15%     
==========================================
  Files          97       97              
  Lines        5463     5446      -17     
==========================================
- Hits         3423     3421       -2     
+ Misses       1603     1588      -15     
  Partials      437      437              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.81% <11.76%> (+0.15%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/database.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9kYXRhYmFzZS5nbw==) | `41.17% <0.00%> (ø)` | |
| [tasks/system\_culling/culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `66.12% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzLmdv) | `32.05% <0.00%> (+3.64%)` | :arrow_up: |
| [tasks/vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `66.10% <0.00%> (-1.14%)` | :arrow_down: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `33.33% <100.00%> (+1.36%)` | :arrow_up: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.85% <0.00%> (-4.29%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1039?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Commented on August 01, 2022 at 11:26 AM UTC

### Review by @MichaelMraka - Commented on August 01, 2022 at 03:40 PM UTC

### Review by @MichaelMraka - Approved on August 24, 2022 at 09:40 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1039*
