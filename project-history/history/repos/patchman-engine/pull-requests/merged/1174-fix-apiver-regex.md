---
type: pull_request
number: 1174
title: "fix apiver regex"
state: merged
author: psegedy
created: 2023-02-21T17:00:05Z
updated: 2023-02-22T10:14:52Z
closed: 2023-02-22T10:14:52Z
merged: 2023-02-22T10:14:52Z
base_branch: master
head_branch: v3fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1174
---

# Pull Request #1174: fix apiver regex

**Author**: @psegedy
**Created**: February 21, 2023 at 05:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `v3fix`

## Description

basePath = /api/patch/v3

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

### Comment by @jira-linking on February 21, 2023 at 05:00 PM UTC

Commits missing Jira IDs:
3b2c15f9ad13cea364f9326e5705d663bb55bbd0


### Comment by @codecov-commenter on February 21, 2023 at 05:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1174?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.16**% // Head: **63.16**% // No change to project coverage :thumbsup:
> Coverage data is based on head [(`3b2c15f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1174?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`8814bac`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8814bacdb19a122c94d34739fe995145091db2f8?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch has no changes to coverable lines.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1174   +/-   ##
=======================================
  Coverage   63.16%   63.16%           
=======================================
  Files         102      102           
  Lines        5878     5878           
=======================================
  Hits         3713     3713           
  Misses       1699     1699           
  Partials      466      466           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1174?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/swagger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1174?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9zd2FnZ2VyLmdv) | `0.00% <ø> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1174?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on February 22, 2023 at 10:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1174*
