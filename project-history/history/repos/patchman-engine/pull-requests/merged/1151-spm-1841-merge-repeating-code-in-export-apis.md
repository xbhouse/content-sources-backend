---
type: pull_request
number: 1151
title: "SPM-1841: merge repeating code in /export APIs"
state: merged
author: MichaelMraka
created: 2022-12-20T10:22:53Z
updated: 2022-12-20T12:38:25Z
closed: 2022-12-20T12:38:24Z
merged: 2022-12-20T12:38:24Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1151
---

# Pull Request #1151: SPM-1841: merge repeating code in /export APIs

**Author**: @MichaelMraka
**Created**: December 20, 2022 at 10:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @jira-linking on December 20, 2022 at 10:22 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1841


### Comment by @codecov-commenter on December 20, 2022 at 10:30 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.24**% // Head: **62.24**% // Decreases project coverage by **`-0.00%`** :warning:
> Coverage data is based on head [(`ae4338e`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`bd16798`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bd16798911d030fc5ea89fc677484f53583b321c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1151      +/-   ##
==========================================
- Coverage   62.24%   62.24%   -0.01%     
==========================================
  Files         100      100              
  Lines        5816     5776      -40     
==========================================
- Hits         3620     3595      -25     
+ Misses       1730     1720      -10     
+ Partials      466      461       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.24% <100.00%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `68.00% <100.00%> (-7.00%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `48.27% <100.00%> (-1.73%)` | :arrow_down: |
| [manager/controllers/package\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXNfZXhwb3J0Lmdv) | `51.85% <100.00%> (-1.09%)` | :arrow_down: |
| [manager/controllers/packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlc19leHBvcnQuZ28=) | `43.75% <100.00%> (-4.08%)` | :arrow_down: |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `55.17% <100.00%> (-0.39%)` | :arrow_down: |
| [manager/controllers/system\_packages\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | `60.60% <100.00%> (+0.60%)` | :arrow_up: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `70.58% <100.00%> (-8.58%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.42% <100.00%> (+0.34%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1151?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @michalslomczynski - Approved on December 20, 2022 at 11:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1151*
