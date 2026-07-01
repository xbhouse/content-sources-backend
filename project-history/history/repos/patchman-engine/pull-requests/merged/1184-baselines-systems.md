---
type: pull_request
number: 1184
title: "Baselines systems"
state: merged
author: psegedy
created: 2023-03-03T17:53:30Z
updated: 2023-03-08T12:48:53Z
closed: 2023-03-08T12:48:53Z
merged: 2023-03-08T12:48:53Z
base_branch: master
head_branch: baselines_systems
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1184
---

# Pull Request #1184: Baselines systems

**Author**: @psegedy
**Created**: March 03, 2023 at 05:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baselines_systems`

## Description

todo:
- [x] set counts in evaluator
- [x] fix 500 in baselines/systems - SPM-1928 
- [x] sort docstring
- [x] do we need to add filters?
- [x] add return after LogAndRespError
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

### Comment by @jira-linking on March 03, 2023 at 05:53 PM UTC

Commits missing Jira IDs:
4fcaa34f2395fd66200f043f90bea022affa66f5
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1911
https://issues.redhat.com/browse/SPM-1928


### Comment by @psegedy on March 03, 2023 at 06:02 PM UTC

/retest

### Comment by @codecov-commenter on March 07, 2023 at 01:42 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`72.12`**% and project coverage change: **`+0.03`** :tada:
> Comparison is base [(`ccb1698`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ccb1698df35bc79a6f480d65c98ad016b5ae8f73?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.93% compared to head [(`d937a45`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.97%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1184      +/-   ##
==========================================
+ Coverage   62.93%   62.97%   +0.03%     
==========================================
  Files         103      103              
  Lines        5957     6052      +95     
==========================================
+ Hits         3749     3811      +62     
- Misses       1741     1768      +27     
- Partials      467      473       +6     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.97% <72.12%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.50% <0.00%> (-0.05%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_v3.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | `0.00% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.45% <ø> (+3.50%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.37% <60.00%> (-0.38%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.11% <74.28%> (-0.91%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `70.94% <81.35%> (-4.06%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `77.92% <83.33%> (+0.99%)` | :arrow_up: |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `90.47% <100.00%> (-0.44%)` | :arrow_down: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `80.00% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `90.36% <100.00%> (ø)` | |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1184?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on March 07, 2023 at 09:23 AM UTC

### Review by @MichaelMraka - Approved on March 07, 2023 at 01:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1184*
