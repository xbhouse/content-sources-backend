---
type: pull_request
number: 1475
title: "RHINENG-10032: implement pg_repack job and admin API"
state: merged
author: Dugowitch
created: 2024-09-06T12:01:20Z
updated: 2024-09-12T12:44:42Z
closed: 2024-09-12T12:44:42Z
merged: 2024-09-12T12:44:42Z
base_branch: master
head_branch: RHINENG-10032
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1475
---

# Pull Request #1475: RHINENG-10032: implement pg_repack job and admin API

**Author**: @Dugowitch
**Created**: September 06, 2024 at 12:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-10032`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [x] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [x] Database Security
- [ ] File Management
- [ ] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @jira-linking on September 06, 2024 at 12:01 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10032


### Comment by @codecov-commenter on September 06, 2024 at 02:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1475?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.06%. Comparing base [(`ebaa313`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ebaa3139e0f79d8eba153b8a09416c2160e1838d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`812cde5`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/812cde59ccc35a9e8c8fae384f6ad35b942b3c38?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 16 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1475      +/-   ##
==========================================
+ Coverage   65.01%   65.06%   +0.04%     
==========================================
  Files         114      114              
  Lines        7823     7877      +54     
==========================================
+ Hits         5086     5125      +39     
- Misses       2207     2215       +8     
- Partials      530      537       +7     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1475/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1475/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.06% <ø> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1475?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Dismissed on September 06, 2024 at 03:01 PM UTC

### Review by @Dugowitch - Commented on September 09, 2024 at 09:28 AM UTC

### Review by @Dugowitch - Commented on September 09, 2024 at 09:47 AM UTC

### Review by @Dugowitch - Commented on September 09, 2024 at 11:00 AM UTC

### Review by @Dugowitch - Commented on September 10, 2024 at 01:01 PM UTC

### Review by @Dugowitch - Commented on September 10, 2024 at 01:58 PM UTC

### Review by @psegedy - Approved on September 12, 2024 at 12:44 PM UTC

LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1475*
