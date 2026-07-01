---
type: pull_request
number: 1271
title: "ESSNTL-4817: rewrite filter parsing"
state: merged
author: MichaelMraka
created: 2023-07-24T14:53:40Z
updated: 2023-07-28T14:50:04Z
closed: 2023-07-28T14:50:04Z
merged: 2023-07-28T14:50:04Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1271
---

# Pull Request #1271: ESSNTL-4817: rewrite filter parsing

**Author**: @MichaelMraka
**Created**: July 24, 2023 at 02:53 PM UTC
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

### Comment by @jira-linking on July 24, 2023 at 02:53 PM UTC

Commits missing Jira IDs:
4e16789e6ec11d90751b4425b4e5720c488a6bbd
fa8f44cff500d86e00a5082a07ce2219efda816b
5768be1032a71720f8f284ad31b2d9bf06cd379f
06acc69a032a104223f61ff55e271a90fea13a74
2c2c96fb12d7795b4015573d7c064062dd38b2de
f4a9263f53c01845d4730e4ded32f7a0cc89ebd6
8b9a49fd75f59f09b028474b0558dfb8e17ba69c
3bf8acbb2c6e99cf6a47d4b4a3bff4d0e64c4570
a39161580cdc0d8b5e04fec4a4ce5fbc5b3cd419
dabc67dc18a4085f987a8dd3203086411e132e64
77ac9e2f55288bf766f82f8d26dec48a3167b62c
ecfcdc78970fde7ffbced73d24ca496294755718
cb7a3b9cd92ee14dc5b77ed2a9fb40db8ad1a704


### Comment by @codecov-commenter on July 25, 2023 at 12:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`94.50%`** and project coverage change: **`-0.63%`** :warning:
> Comparison is base [(`e0b99e8`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e0b99e8cc146ef1874b7cedfc31ae8376ede4d9a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.11% compared to head [(`f80b87b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 60.48%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1271      +/-   ##
==========================================
- Coverage   61.11%   60.48%   -0.63%     
==========================================
  Files         106      106              
  Lines        6687     6631      -56     
==========================================
- Hits         4087     4011      -76     
- Misses       2061     2086      +25     
+ Partials      539      534       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.48% <94.50%> (-0.63%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `55.00% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `53.73% <50.00%> (-0.96%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `48.59% <66.66%> (-0.45%)` | :arrow_down: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `79.72% <97.50%> (+2.03%)` | :arrow_up: |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `7.69% <100.00%> (+1.60%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `67.79% <100.00%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `48.97% <100.00%> (-4.09%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.85% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `41.86% <100.00%> (ø)` | |
| [manager/controllers/baselines.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `83.78% <100.00%> (ø)` | |
| ... and [10 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [4 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1271?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 25, 2023 at 03:12 PM UTC

some tests seem to be broken, 1 test seems to be failing due to changed openapi scheme, others might be affected by potential bug in https://github.com/RedHatInsights/patchman-engine/pull/1264 

### Comment by @psegedy on July 25, 2023 at 03:17 PM UTC

/retest

### Comment by @MichaelMraka on July 27, 2023 at 09:16 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on July 25, 2023 at 03:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1271*
