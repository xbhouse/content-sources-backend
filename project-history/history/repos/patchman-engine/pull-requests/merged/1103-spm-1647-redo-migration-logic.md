---
type: pull_request
number: 1103
title: "SPM-1647: redo migration logic"
state: merged
author: MichaelMraka
created: 2022-09-12T14:30:21Z
updated: 2022-09-13T09:33:49Z
closed: 2022-09-13T09:33:49Z
merged: 2022-09-13T09:33:49Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1103
---

# Pull Request #1103: SPM-1647: redo migration logic

**Author**: @MichaelMraka
**Created**: September 12, 2022 at 02:30 PM UTC
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

### Comment by @jira-linking on September 12, 2022 at 02:30 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1647


### Comment by @psegedy on September 12, 2022 at 02:38 PM UTC

```
database_admin/schema_test.go:121:25: undeclared name: `os` 
```

### Comment by @codecov-commenter on September 12, 2022 at 02:52 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.03**% // Head: **62.90**% // Decreases project coverage by **`-0.12%`** :warning:
> Coverage data is based on head [(`82f199b`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`e1b799a`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e1b799a38d84c16867c9ec46202c7230c4f9e30d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 11.76% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1103      +/-   ##
==========================================
- Coverage   63.03%   62.90%   -0.13%     
==========================================
  Files          97       97              
  Lines        5522     5535      +13     
==========================================
+ Hits         3481     3482       +1     
- Misses       1597     1609      +12     
  Partials      444      444              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.90% <11.76%> (-0.13%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `41.66% <0.00%> (-10.06%)` | :arrow_down: |
| [manager/middlewares/swagger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9zd2FnZ2VyLmdv) | `0.00% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/dbchange.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9kYmNoYW5nZS5nbw==) | `12.90% <20.00%> (-0.44%)` | :arrow_down: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.90% <100.00%> (+0.53%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1103?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on September 12, 2022 at 02:37 PM UTC

lgtm, let's wait for tests

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1103*
