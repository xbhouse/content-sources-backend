---
type: pull_request
number: 1636
title: "HMS-5960: don't remove templates in listener"
state: merged
author: psegedy
created: 2025-04-14T14:26:49Z
updated: 2026-04-03T15:00:37Z
closed: 2025-04-14T15:16:34Z
merged: 2025-04-14T15:16:34Z
base_branch: master
head_branch: fixenv
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1636
---

# Pull Request #1636: HMS-5960: don't remove templates in listener

**Author**: @psegedy
**Created**: April 14, 2025 at 02:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fixenv`

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

### Comment by @jira-linking on April 14, 2025 at 02:26 PM UTC

Commits missing Jira IDs:
1ae9958673fa91aa65177792ef94d525dab0f8f6
b258666ded06747caa700ef7a06b9954e9c84e7a


### Comment by @codecov-commenter on April 14, 2025 at 02:32 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `42.85714%` with `8 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 58.16%. Comparing base ([`63cd599`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/63cd5991d44ef2632c26dbdec18d8a94a9c1381f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b258666`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b258666ded06747caa700ef7a06b9954e9c84e7a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 883 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 11.11% | [7 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1636      +/-   ##
==========================================
- Coverage   58.19%   58.16%   -0.04%     
==========================================
  Files         146      146              
  Lines       11330    11333       +3     
==========================================
- Hits         6594     6592       -2     
- Misses       4154     4159       +5     
  Partials      582      582              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.16% <42.85%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1636?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on April 14, 2025 at 02:37 PM UTC

👍🏼 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1636*
