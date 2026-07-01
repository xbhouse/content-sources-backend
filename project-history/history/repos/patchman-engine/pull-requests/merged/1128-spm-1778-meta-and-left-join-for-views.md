---
type: pull_request
number: 1128
title: "SPM-1778: meta and left join for /views"
state: merged
author: psegedy
created: 2022-10-14T13:56:36Z
updated: 2022-10-14T15:36:37Z
closed: 2022-10-14T15:36:36Z
merged: 2022-10-14T15:36:36Z
base_branch: master
head_branch: viewsmeta
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1128
---

# Pull Request #1128: SPM-1778: meta and left join for /views

**Author**: @psegedy
**Created**: October 14, 2022 at 01:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `viewsmeta`

## Description

before 
```json
POST "http://localhost:8080/api/patch/v2/views/systems/advisories" -d '{"advisories": ["RH-2"], "limit": 10, "offset": 0}'

{
  "data": {
    "00000000-0000-0000-0000-000000000001": [
      "RH-2"
    ],
    "00000000-0000-0000-0000-000000000003": [
      "RH-2"
    ]
  }
}
```

after
```json
POST "http://localhost:8080/api/patch/v2/views/systems/advisories" -d '{"advisories": ["RH-2"], "limit": 10, "offset": 0}'

{
  "data": {
    "00000000-0000-0000-0000-000000000001": [
      "RH-2"
    ],
    "00000000-0000-0000-0000-000000000003": [
      "RH-2"
    ],
    "00000000-0000-0000-0000-000000000004": [
      ""
    ],
    "00000000-0000-0000-0000-000000000005": [
      ""
    ],
    "00000000-0000-0000-0000-000000000006": [
      ""
    ],
    "00000000-0000-0000-0000-000000000008": [
      ""
    ]
  },
  "meta": {
    "limit": 10,
    "offset": 0,
    "filter": null,
    "total_items": 6
  }
} 
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

### Comment by @jira-linking on October 14, 2022 at 01:56 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1778


### Comment by @codecov-commenter on October 14, 2022 at 02:27 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.19**% // Head: **62.21**% // Increases project coverage by **`+0.01%`** :tada:
> Coverage data is based on head [(`1591f3a`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`c27c58c`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/c27c58cfad25c21156f84e3c5841d7617fc97f34?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 79.62% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1128      +/-   ##
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
| unittests | `62.21% <79.62%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `86.07% <66.66%> (ø)` | |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `80.00% <80.39%> (-4.42%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1128?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on October 14, 2022 at 03:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1128*
