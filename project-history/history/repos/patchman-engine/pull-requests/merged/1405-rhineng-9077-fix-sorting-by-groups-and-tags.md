---
type: pull_request
number: 1405
title: "RHINENG-9077: fix sorting by groups and tags"
state: merged
author: psegedy
created: 2024-04-04T17:30:27Z
updated: 2024-04-05T13:32:40Z
closed: 2024-04-05T13:32:40Z
merged: 2024-04-05T13:32:40Z
base_branch: master
head_branch: groups_sort
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1405
---

# Pull Request #1405: RHINENG-9077: fix sorting by groups and tags

**Author**: @psegedy
**Created**: April 04, 2024 at 05:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `groups_sort`

## Description

change gorm column name of groups and tags so we are able to sort by `sort=groups` and `sort=tags` (`sort=tags` isn't documented but I changed it) it was possible to sort with `sort=groups_json` (breaks sorting in prod) and `sort=tags_json`

I would expect that sorting should work according to attribute name in json, I tried to changed that to use the json attribute name instead of gorm column name but that broke osname, osmajor, osminor sorts...
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

### Comment by @jira-linking on April 04, 2024 at 05:30 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9077


### Comment by @codecov-commenter on April 04, 2024 at 05:37 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1405?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.62%. Comparing base [(`ccba24a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ccba24a45b67a36f4ef6b0d4323682911cb5af4b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`3d77339`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1405?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 5 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1405      +/-   ##
==========================================
+ Coverage   63.59%   63.62%   +0.02%     
==========================================
  Files         112      112              
  Lines        6774     6779       +5     
==========================================
+ Hits         4308     4313       +5     
  Misses       1958     1958              
  Partials      508      508              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1405/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1405/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.62% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1405?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on April 05, 2024 at 08:05 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1405*
