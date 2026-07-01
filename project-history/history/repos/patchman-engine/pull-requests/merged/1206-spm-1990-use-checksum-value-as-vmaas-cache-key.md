---
type: pull_request
number: 1206
title: "SPM-1990: use checksum value as vmaas-cache key"
state: merged
author: psegedy
created: 2023-03-31T10:06:59Z
updated: 2023-04-03T14:11:53Z
closed: 2023-04-03T14:11:53Z
merged: 2023-04-03T14:11:53Z
base_branch: master
head_branch: vmaas_cache_2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1206
---

# Pull Request #1206: SPM-1990: use checksum value as vmaas-cache key

**Author**: @psegedy
**Created**: March 31, 2023 at 10:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas_cache_2`

## Description

- checksum is assigned to a new variable after every upload, so it has a different pointer and there is no cache hit
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

### Comment by @jira-linking on March 31, 2023 at 10:07 AM UTC

Commits missing Jira IDs:
06bd3cb0a2c5c6fabf6a58c52a173364b9ca026d
b59c6ca4256361af893af6ffb625c7cd34b1bbf0
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1990


### Comment by @codecov-commenter on March 31, 2023 at 02:40 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`89.28`**% and project coverage change: **`+0.02`** :tada:
> Comparison is base [(`b80f76c`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b80f76ce4f153bb831a0752f1c203d5efd382332?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.12% compared to head [(`b59c6ca`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.15%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1206      +/-   ##
==========================================
+ Coverage   62.12%   62.15%   +0.02%     
==========================================
  Files         103      103              
  Lines        6247     6246       -1     
==========================================
+ Hits         3881     3882       +1     
+ Misses       1869     1867       -2     
  Partials      497      497              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.15% <89.28%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [tasks/vmaas\_sync/dbchange.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9kYmNoYW5nZS5nbw==) | `61.90% <ø> (ø)` | |
| [tasks/vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `71.83% <ø> (ø)` | |
| [evaluator/vmaas\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `60.00% <66.66%> (+5.16%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `81.87% <94.44%> (ø)` | |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `80.88% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1206?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on April 03, 2023 at 01:34 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1206*
