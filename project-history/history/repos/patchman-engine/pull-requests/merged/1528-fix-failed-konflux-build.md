---
type: pull_request
number: 1528
title: "fix failed konflux build"
state: merged
author: MichaelMraka
created: 2024-11-07T12:06:28Z
updated: 2024-11-07T12:53:02Z
closed: 2024-11-07T12:53:02Z
merged: 2024-11-07T12:53:02Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1528
---

# Pull Request #1528: fix failed konflux build

**Author**: @MichaelMraka
**Created**: November 07, 2024 at 12:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

## Description

This reverts commit d9dd52410d09dfeb3d1eabd4d71fd869e89f67c6.

fixes failed konflux build

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

### Comment by @jira-linking on November 07, 2024 at 12:06 PM UTC

Commits missing Jira IDs:
5a06e4668243f823db049088eb5b73280f90fc69


### Comment by @codecov-commenter on November 07, 2024 at 12:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1528?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.50%. Comparing base [(`e50db13`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e50db13a5ceabd05bb8dddcbf7db95048b638e14?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5a06e46`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5a06e4668243f823db049088eb5b73280f90fc69?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1528   +/-   ##
=======================================
  Coverage   63.50%   63.50%           
=======================================
  Files         114      114           
  Lines        9609     9609           
=======================================
  Hits         6102     6102           
  Misses       2970     2970           
  Partials      537      537           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1528/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1528/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.50% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1528?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jdobes on November 07, 2024 at 12:27 PM UTC

/retest

---

## Reviews

### Review by @jdobes - Approved on November 07, 2024 at 12:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1528*
