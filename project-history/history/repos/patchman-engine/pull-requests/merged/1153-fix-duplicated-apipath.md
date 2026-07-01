---
type: pull_request
number: 1153
title: "fix duplicated apiPath"
state: merged
author: psegedy
created: 2023-01-06T10:59:40Z
updated: 2023-01-09T12:04:18Z
closed: 2023-01-09T12:04:18Z
merged: 2023-01-09T12:04:18Z
base_branch: master
head_branch: path
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1153
---

# Pull Request #1153: fix duplicated apiPath

**Author**: @psegedy
**Created**: January 06, 2023 at 10:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `path`

## Description

affects ephemeral
```
patchman-admin-mzw2x                        env-ephemeral-x9qw1o-xquihayy.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com        /api/patch/                       patchman-admin                        auth       edge/Redirect   None
patchman-manager-hqcz8                      HostAlreadyClaimed                                                             /api/patch/                       patchman-manager                      auth       edge/Redirect   None
```
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

### Comment by @jira-linking on January 06, 2023 at 10:59 AM UTC

Commits missing Jira IDs:
bae50d371e4e513a5ade6f88fb0bcd9811c1549b


### Comment by @codecov-commenter on January 06, 2023 at 11:07 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1153?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.18**% // Head: **62.24**% // Increases project coverage by **`+0.05%`** :tada:
> Coverage data is based on head [(`bae50d3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1153?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`1109e51`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1109e515a37f3b44e2cea97dcb98d5bbc0b36bbe?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch has no changes to coverable lines.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1153      +/-   ##
==========================================
+ Coverage   62.18%   62.24%   +0.05%     
==========================================
  Files         100      100              
  Lines        5776     5776              
==========================================
+ Hits         3592     3595       +3     
+ Misses       1723     1720       -3     
  Partials      461      461              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.24% <ø> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1153?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1153?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.14% <0.00%> (+4.28%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1153?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 09, 2023 at 08:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1153*
