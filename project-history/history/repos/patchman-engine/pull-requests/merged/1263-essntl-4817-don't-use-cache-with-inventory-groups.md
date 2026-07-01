---
type: pull_request
number: 1263
title: "ESSNTL-4817: don't use cache with inventory groups"
state: merged
author: psegedy
created: 2023-07-17T11:19:59Z
updated: 2023-07-18T09:26:39Z
closed: 2023-07-18T09:26:39Z
merged: 2023-07-18T09:26:39Z
base_branch: master
head_branch: groups_cache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1263
---

# Pull Request #1263: ESSNTL-4817: don't use cache with inventory groups

**Author**: @psegedy
**Created**: July 17, 2023 at 11:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `groups_cache`

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

### Comment by @jira-linking on July 17, 2023 at 11:20 AM UTC

Commits missing Jira IDs:
025108aee545edc0b5f47eb043ebbf66c31205d0
5406a54cac21690ae85130bdd0acb6395d467066


### Comment by @codecov-commenter on July 17, 2023 at 11:27 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`33.33`**% and no project coverage change.
> Comparison is base [(`b3e2e9a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b3e2e9afbead26df36dd52101238446ff4f301a5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.52% compared to head [(`e719593`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.52%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1263   +/-   ##
=======================================
  Coverage   61.52%   61.52%           
=======================================
  Files         106      106           
  Lines        6594     6594           
=======================================
  Hits         4057     4057           
  Misses       2014     2014           
  Partials      523      523           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.52% <33.33%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `46.34% <0.00%> (ø)` | |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `63.23% <0.00%> (ø)` | |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.54% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1263?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on July 18, 2023 at 09:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1263*
