---
type: pull_request
number: 1173
title: "SPM-1910: create v3 API"
state: merged
author: psegedy
created: 2023-02-20T16:41:21Z
updated: 2023-02-21T10:26:44Z
closed: 2023-02-21T10:26:44Z
merged: 2023-02-21T10:26:44Z
base_branch: master
head_branch: v3api
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1173
---

# Pull Request #1173: SPM-1910: create v3 API

**Author**: @psegedy
**Created**: February 20, 2023 at 04:41 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `v3api`

## Description

/major

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

### Comment by @jira-linking on February 20, 2023 at 04:41 PM UTC

Commits missing Jira IDs:
75eb84d257fc7bb8baeb59dae006f6cfb53620ac
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1910


### Comment by @codecov-commenter on February 20, 2023 at 07:21 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.31**% // Head: **63.21**% // Decreases project coverage by **`-0.10%`** :warning:
> Coverage data is based on head [(`75eb84d`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`362bd61`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/362bd61316bbc55a1d80b15798d84d32ebc9d25a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 19.23% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1173      +/-   ##
==========================================
- Coverage   63.31%   63.21%   -0.10%     
==========================================
  Files         102      102              
  Lines        5872     5878       +6     
==========================================
- Hits         3718     3716       -2     
- Misses       1688     1696       +8     
  Partials      466      466              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.21% <19.23%> (-0.10%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `42.85% <0.00%> (+1.19%)` | :arrow_up: |
| [manager/middlewares/swagger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9zd2FnZ2VyLmdv) | `0.00% <0.00%> (ø)` | |
| [base/core/gintesting.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2dpbnRlc3RpbmcuZ28=) | `83.33% <60.00%> (-6.67%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `80.30% <100.00%> (-0.44%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1173?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 21, 2023 at 10:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1173*
