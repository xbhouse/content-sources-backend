---
type: pull_request
number: 1208
title: "SPM-2013: store applicable but not installable advisories to aad"
state: merged
author: psegedy
created: 2023-04-12T14:02:50Z
updated: 2023-04-13T08:20:24Z
closed: 2023-04-13T08:20:24Z
merged: 2023-04-13T08:20:24Z
base_branch: master
head_branch: negative_applicable_count
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1208
---

# Pull Request #1208: SPM-2013: store applicable but not installable advisories to aad

**Author**: @psegedy
**Created**: April 12, 2023 at 02:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `negative_applicable_count`

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

### Comment by @jira-linking on April 12, 2023 at 02:02 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2013


### Comment by @codecov-commenter on April 12, 2023 at 02:13 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`72.34`**% and project coverage change: **`-0.06`** :warning:
> Comparison is base [(`b80f76c`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b80f76ce4f153bb831a0752f1c203d5efd382332?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.12% compared to head [(`42f44fc`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.07%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1208      +/-   ##
==========================================
- Coverage   62.12%   62.07%   -0.06%     
==========================================
  Files         103      103              
  Lines        6247     6257      +10     
==========================================
+ Hits         3881     3884       +3     
- Misses       1869     1875       +6     
- Partials      497      498       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.07% <72.34%> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.07% <ø> (ø)` | |
| [tasks/vmaas\_sync/dbchange.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9kYmNoYW5nZS5nbw==) | `61.90% <ø> (ø)` | |
| [tasks/vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `71.83% <ø> (ø)` | |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `69.37% <25.00%> (-2.24%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `83.93% <55.55%> (-0.74%)` | :arrow_down: |
| [evaluator/vmaas\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3ZtYWFzX2NhY2hlLmdv) | `60.00% <66.66%> (+5.16%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `81.87% <94.44%> (ø)` | |
| [base/core/gintesting.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2dpbnRlc3RpbmcuZ28=) | `84.61% <100.00%> (ø)` | |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `80.88% <100.00%> (ø)` | |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `56.66% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1208?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on April 12, 2023 at 08:16 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on April 12, 2023 at 02:19 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1208*
