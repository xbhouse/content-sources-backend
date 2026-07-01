---
type: pull_request
number: 1132
title: "udpate dependencies"
state: merged
author: psegedy
created: 2022-10-17T15:10:43Z
updated: 2022-10-20T11:44:12Z
closed: 2022-10-20T11:44:11Z
merged: 2022-10-20T11:44:11Z
base_branch: master
head_branch: deps
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1132
---

# Pull Request #1132: udpate dependencies

**Author**: @psegedy
**Created**: October 17, 2022 at 03:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `deps`

## Description

- update go version in go.mod
- update all modules except `gocsv`, latest version does not work properly
- golangci-lint update
- go.mod now includes all indirect dependencies
- sadly, it looks like go.sum must be this big

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

### Comment by @jira-linking on October 17, 2022 at 03:10 PM UTC

Commits missing Jira IDs:
86b3be75fb9b0a64c99821c65365ccad3c39f653
afe21c4d20685aa4bfb97ec42770fa6502341d7f


### Comment by @codecov-commenter on October 19, 2022 at 10:38 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.19**% // Head: **62.21**% // Increases project coverage by **`+0.01%`** :tada:
> Coverage data is based on head [(`afe21c4`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`c27c58c`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/c27c58cfad25c21156f84e3c5841d7617fc97f34?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 80.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1132      +/-   ##
==========================================
+ Coverage   62.19%   62.21%   +0.01%     
==========================================
  Files          99       99              
  Lines        5738     5761      +23     
==========================================
+ Hits         3569     3584      +15     
- Misses       1713     1717       +4     
- Partials      456      460       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.21% <80.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/awscloudwatch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9hd3NjbG91ZHdhdGNoLmdv) | `13.79% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.07% <66.66%> (ø)` | |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `80.00% <80.39%> (-4.42%)` | :arrow_down: |
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `41.66% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1132?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on October 20, 2022 at 09:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1132*
