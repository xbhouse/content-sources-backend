---
type: pull_request
number: 1525
title: "chore: use replicas instead of minReplicas"
state: merged
author: psegedy
created: 2024-11-04T12:57:00Z
updated: 2024-11-04T15:18:03Z
closed: 2024-11-04T15:18:03Z
merged: 2024-11-04T15:18:03Z
base_branch: master
head_branch: replicas
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1525
---

# Pull Request #1525: chore: use replicas instead of minReplicas

**Author**: @psegedy
**Created**: November 04, 2024 at 12:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `replicas`

## Description

minReplicas is deprecated https://redhatinsights.github.io/clowder/clowder/dev/api_reference.html\#k8s-api-github-com-redhatinsights-clowder-apis-cloud-redhat-com-v1alpha1-deployment

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

### Comment by @jira-linking on November 04, 2024 at 12:57 PM UTC

Commits missing Jira IDs:
54caa2976fa46c798ba19a288c72d2396fd8ebd0


### Comment by @codecov-commenter on November 04, 2024 at 01:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1525?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.47%. Comparing base [(`b0e0764`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b0e0764c0178436c19a96e50b1c5eb4643e5379d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`54caa29`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/54caa2976fa46c798ba19a288c72d2396fd8ebd0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1525      +/-   ##
==========================================
- Coverage   63.50%   63.47%   -0.04%     
==========================================
  Files         114      114              
  Lines        9609     9609              
==========================================
- Hits         6102     6099       -3     
- Misses       2970     2973       +3     
  Partials      537      537              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1525/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1525/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.47% <ø> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1525?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on November 04, 2024 at 01:36 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1525*
