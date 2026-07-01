---
type: pull_request
number: 1391
title: "RHINENG-8717: fix total count in /views"
state: merged
author: psegedy
created: 2024-03-06T12:34:05Z
updated: 2024-03-07T15:27:24Z
closed: 2024-03-07T15:27:24Z
merged: 2024-03-07T15:27:24Z
base_branch: master
head_branch: views_meta
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1391
---

# Pull Request #1391: RHINENG-8717: fix total count in /views

**Author**: @psegedy
**Created**: March 06, 2024 at 12:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `views_meta`

## Description

Currently, total_items shows number of items after `Limit` is applied by `ListCommon`, we'd like to show real total items.
Separate `tx.Limit` from `ListCommon` so the limit can be applied to a subquery after we get the total count.
What do you think?
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

### Comment by @jira-linking on March 06, 2024 at 12:34 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8717


### Comment by @codecov-commenter on March 06, 2024 at 01:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1391?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.73%. Comparing base [(`242afaf`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/242afaf85ff5608629a75e41ffb93249347595c0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8fc837f`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1391?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1391      +/-   ##
==========================================
+ Coverage   63.65%   63.73%   +0.08%     
==========================================
  Files         107      107              
  Lines        6504     6513       +9     
==========================================
+ Hits         4140     4151      +11     
+ Misses       1875     1873       -2     
  Partials      489      489              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1391/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1391/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.73% <100.00%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1391?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 07, 2024 at 08:54 AM UTC

### Review by @Dugowitch - Approved on March 07, 2024 at 09:58 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1391*
