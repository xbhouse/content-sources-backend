---
type: pull_request
number: 1226
title: "SPM-1915: Update /system/packages API to templates"
state: merged
author: MichaelMraka
created: 2023-05-30T11:21:25Z
updated: 2023-06-01T14:56:22Z
closed: 2023-06-01T14:56:22Z
merged: 2023-06-01T14:56:22Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1226
---

# Pull Request #1226: SPM-1915: Update /system/packages API to templates

**Author**: @MichaelMraka
**Created**: May 30, 2023 at 11:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @jira-linking on May 30, 2023 at 11:21 AM UTC

Commits missing Jira IDs:
f2ae20dce2440f120638d5f015bfa8518dda365d
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1915


### Comment by @codecov-commenter on May 30, 2023 at 11:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`60.34`**% and project coverage change: **`-0.22`** :warning:
> Comparison is base [(`2c48a45`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2c48a453ad2b5350ca9abb1dcd408fa1b06c31ba?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.12% compared to head [(`8bcbc5b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.91%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1226      +/-   ##
==========================================
- Coverage   62.12%   61.91%   -0.22%     
==========================================
  Files         104      105       +1     
  Lines        6348     6433      +85     
==========================================
+ Hits         3944     3983      +39     
- Misses       1901     1938      +37     
- Partials      503      512       +9     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.91% <60.34%> (-0.22%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `53.96% <44.82%> (-16.49%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.10% <50.00%> (-1.74%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `69.44% <50.00%> (-3.02%)` | :arrow_down: |
| [evaluator/status.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3N0YXR1cy5nbw==) | `60.00% <60.00%> (ø)` | |
| [manager/controllers/system\_packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | `58.00% <64.28%> (-3.77%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_v3.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | `76.08% <64.70%> (-6.68%)` | :arrow_down: |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.02% <69.23%> (-0.75%)` | :arrow_down: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `83.52% <80.00%> (-0.09%)` | :arrow_down: |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.72% <100.00%> (+0.08%)` | :arrow_up: |
| ... and [1 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1226?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Changes Requested on May 30, 2023 at 03:26 PM UTC

### Review by @MichaelMraka - Commented on May 31, 2023 at 01:35 PM UTC

### Review by @MichaelMraka - Commented on May 31, 2023 at 01:35 PM UTC

### Review by @psegedy - Approved on May 31, 2023 at 03:56 PM UTC

LGTM, please squash fixup commits

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1226*
