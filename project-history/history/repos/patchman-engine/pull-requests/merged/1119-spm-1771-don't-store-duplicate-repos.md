---
type: pull_request
number: 1119
title: "SPM-1771: don't store duplicate repos"
state: merged
author: psegedy
created: 2022-10-06T10:10:40Z
updated: 2022-10-06T12:05:36Z
closed: 2022-10-06T12:05:36Z
merged: 2022-10-06T12:05:36Z
base_branch: master
head_branch: duplicate_repo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1119
---

# Pull Request #1119: SPM-1771: don't store duplicate repos

**Author**: @psegedy
**Created**: October 06, 2022 at 10:10 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `duplicate_repo`

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

### Comment by @jira-linking on October 06, 2022 at 10:10 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1771


### Comment by @codecov-commenter on October 06, 2022 at 10:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1119?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.94**% // Head: **62.94**% // Decreases project coverage by **`-0.00%`** :warning:
> Coverage data is based on head [(`09d4d88`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1119?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`3a74935`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/3a74935e7d863bf847b14f088c09137c9d60df80?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 66.66% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1119      +/-   ##
==========================================
- Coverage   62.94%   62.94%   -0.01%     
==========================================
  Files          99       99              
  Lines        5646     5651       +5     
==========================================
+ Hits         3554     3557       +3     
- Misses       1637     1638       +1     
- Partials      455      456       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.94% <66.66%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1119?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1119/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.14% <66.66%> (-0.28%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1119?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on October 06, 2022 at 11:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1119*
