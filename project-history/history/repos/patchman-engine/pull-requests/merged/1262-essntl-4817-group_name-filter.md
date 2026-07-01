---
type: pull_request
number: 1262
title: "ESSNTL-4817: group_name filter"
state: merged
author: psegedy
created: 2023-07-14T17:00:50Z
updated: 2023-07-18T09:14:36Z
closed: 2023-07-18T09:14:36Z
merged: 2023-07-18T09:14:36Z
base_branch: master
head_branch: groups_filter
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1262
---

# Pull Request #1262: ESSNTL-4817: group_name filter

**Author**: @psegedy
**Created**: July 14, 2023 at 05:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `groups_filter`

## Description

usage: `?filter[group_name][in]=group1&filter[group_name][in]=group2`

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

### Comment by @jira-linking on July 14, 2023 at 05:00 PM UTC

Commits missing Jira IDs:
b081bfac95009309ea69263ee6d556ae527936e5
cb9175aa2b0a4ddcf43028e2225b6d2933d44dc6
32f3202484d9bee6aab1cf6e10d01465d11d0da0
f50bc6210f461abc3e56d250a8ed1c615fcc1287
0aa75a913097449433050fb9bd1e36c88c27987b
7453eb97145337249d2585122ba21e8a4ecd2a35
f254394653a4371bbc797f15ecc6a08d8d92a2a3


### Comment by @codecov-commenter on July 14, 2023 at 05:08 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`72.22`**% and project coverage change: **`-0.11`** :warning:
> Comparison is base [(`b3e2e9a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b3e2e9afbead26df36dd52101238446ff4f301a5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.52% compared to head [(`47a9dc5`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.42%.

> :exclamation: Current head 47a9dc5 differs from pull request most recent head 3733e2a. Consider uploading reports for the commit 3733e2a to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1262      +/-   ##
==========================================
- Coverage   61.52%   61.42%   -0.11%     
==========================================
  Files         106      106              
  Lines        6594     6654      +60     
==========================================
+ Hits         4057     4087      +30     
- Misses       2014     2034      +20     
- Partials      523      533      +10     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.42% <72.22%> (-0.11%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `24.63% <0.00%> (-2.35%)` | :arrow_down: |
| [manager/controllers/baseline\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `55.00% <ø> (ø)` | |
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `46.34% <50.00%> (ø)` | |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `80.28% <70.21%> (-3.24%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.54% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.85% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `41.86% <100.00%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `50.96% <100.00%> (ø)` | |
| [manager/controllers/baselines.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `83.78% <100.00%> (ø)` | |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `62.63% <100.00%> (ø)` | |
| ... and [6 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1262?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on July 18, 2023 at 09:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1262*
