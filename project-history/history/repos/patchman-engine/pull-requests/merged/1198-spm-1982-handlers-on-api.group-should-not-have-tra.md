---
type: pull_request
number: 1198
title: "SPM-1982: handlers on api.group should not have trailing /"
state: merged
author: psegedy
created: 2023-03-22T09:58:02Z
updated: 2023-03-22T12:58:27Z
closed: 2023-03-22T12:58:26Z
merged: 2023-03-22T12:58:26Z
base_branch: master
head_branch: trailing_slash
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1198
---

# Pull Request #1198: SPM-1982: handlers on api.group should not have trailing /

**Author**: @psegedy
**Created**: March 22, 2023 at 09:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `trailing_slash`

## Description

all apis with trailing slash should be redirected to an api without trailing slash

e.g. `/systems/?page=1` should be redirected to `/systems?page=1`

but all apis without trailing slash should not be redirected

e.g. `/systems?page=1` should not be redirected to `/systems/?page=1`
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

### Comment by @codecov-commenter on March 22, 2023 at 10:07 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`73.33`**% and project coverage change: **`-0.06`** :warning:
> Comparison is base [(`d2286bb`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d2286bb98ba696f8c78556fb8a4393b16e7cad90?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.15% compared to head [(`83d3fcd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.09%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1198      +/-   ##
==========================================
- Coverage   62.15%   62.09%   -0.06%     
==========================================
  Files         103      103              
  Lines        6201     6205       +4     
==========================================
- Hits         3854     3853       -1     
- Misses       1856     1860       +4     
- Partials      491      492       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.09% <73.33%> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.07% <66.66%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `45.00% <66.66%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.46% <66.66%> (-1.39%)` | :arrow_down: |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `56.66% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `85.01% <100.00%> (ø)` | |

... and [1 file with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1198?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on March 22, 2023 at 12:36 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on March 22, 2023 at 10:06 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1198*
