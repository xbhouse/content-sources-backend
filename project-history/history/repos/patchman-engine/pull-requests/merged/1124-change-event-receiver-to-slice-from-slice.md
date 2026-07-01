---
type: pull_request
number: 1124
title: "change event receiver to slice from *slice"
state: merged
author: psegedy
created: 2022-10-12T09:34:07Z
updated: 2022-10-12T12:49:46Z
closed: 2022-10-12T12:49:46Z
merged: 2022-10-12T12:49:46Z
base_branch: master
head_branch: sendslice
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1124
---

# Pull Request #1124: change event receiver to slice from *slice

**Author**: @psegedy
**Created**: October 12, 2022 at 09:34 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sendslice`

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

### Comment by @jira-linking on October 12, 2022 at 09:34 AM UTC

Commits missing Jira IDs:
0e4ee03f3497ed5f94a24f1c4ec39aa61d3cf69f


### Comment by @codecov-commenter on October 12, 2022 at 09:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.06**% // Head: **63.07**% // Increases project coverage by **`+0.01%`** :tada:
> Coverage data is based on head [(`0e4ee03`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`f4f4ac5`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f4f4ac58fb08491db4255e91026dd4e55c47a6c9?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 66.66% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1124      +/-   ##
==========================================
+ Coverage   63.06%   63.07%   +0.01%     
==========================================
  Files          99       99              
  Lines        5659     5658       -1     
==========================================
  Hits         3569     3569              
+ Misses       1634     1633       -1     
  Partials      456      456              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.07% <66.66%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `78.57% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.14% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1124?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on October 12, 2022 at 11:13 AM UTC

tests are failing due to bug in tests

---

## Reviews

### Review by @MichaelMraka - Approved on October 12, 2022 at 12:37 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1124*
