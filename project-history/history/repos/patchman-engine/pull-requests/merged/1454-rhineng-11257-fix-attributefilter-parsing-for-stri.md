---
type: pull_request
number: 1454
title: "RHINENG-11257: fix attributeFilter parsing for string value"
state: merged
author: psegedy
created: 2024-07-11T07:21:08Z
updated: 2024-07-12T07:50:10Z
closed: 2024-07-12T07:50:10Z
merged: 2024-07-12T07:50:10Z
base_branch: master
head_branch: rbacschema
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1454
---

# Pull Request #1454: RHINENG-11257: fix attributeFilter parsing for string value

**Author**: @psegedy
**Created**: July 11, 2024 at 07:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rbacschema`

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

### Comment by @jira-linking on July 11, 2024 at 07:21 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11257


### Comment by @codecov-commenter on July 11, 2024 at 07:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `80.00000%` with `6 lines` in your changes missing coverage. Please review.
> Project coverage is 64.81%. Comparing base [(`e6fd838`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e6fd8381d52cde6f7bbe546c901110bd0f36856e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`893ec7d`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/893ec7d2f1e9554fbdd92e254a58c7602d66ccc5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 26 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | 73.33% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/rbac/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?src=pr&el=tree&filepath=base%2Frbac%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9yYmFjL3JiYWMuZ28=) | 86.66% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1454      +/-   ##
==========================================
+ Coverage   63.91%   64.81%   +0.90%     
==========================================
  Files         113      114       +1     
  Lines        6961     7777     +816     
==========================================
+ Hits         4449     5041     +592     
- Misses       1990     2206     +216     
- Partials      522      530       +8     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.81% <80.00%> (+0.90%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1454?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 11, 2024 at 10:50 AM UTC

content-source/pulp deploy fails


### Comment by @psegedy on July 11, 2024 at 10:50 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on July 11, 2024 at 08:14 AM UTC

### Review by @MichaelMraka - Approved on July 11, 2024 at 02:03 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1454*
