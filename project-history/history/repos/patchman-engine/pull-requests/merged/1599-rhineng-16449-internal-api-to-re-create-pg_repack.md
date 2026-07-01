---
type: pull_request
number: 1599
title: "RHINENG-16449: internal api to re-create pg_repack extension"
state: merged
author: psegedy
created: 2025-03-11T16:09:08Z
updated: 2026-04-01T17:37:24Z
closed: 2025-03-12T10:39:46Z
merged: 2025-03-12T10:39:46Z
base_branch: master
head_branch: repackapi
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1599
---

# Pull Request #1599: RHINENG-16449: internal api to re-create pg_repack extension

**Author**: @psegedy
**Created**: March 11, 2025 at 04:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `repackapi`

## Description

- internal api to re-create pg_repack extension
- internal api refactoring
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

### Comment by @jira-linking on March 11, 2025 at 04:09 PM UTC

Commits missing Jira IDs:
c70197b157e0012ecf618122a10387fc75f234a9
1240b8a26e02472f28a73b5e53f77b4af226f842
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16449


### Comment by @codecov-commenter on March 11, 2025 at 04:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `88 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 57.89%. Comparing base ([`15efe07`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/15efe07d709fabc66668ce7bf874e361634e0418?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0d13b64`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0d13b64476e82fb4d54404b8ed92bdf779d8c6e0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 950 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [turnpike/controllers/pprof.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&filepath=turnpike%2Fcontrollers%2Fpprof.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dHVybnBpa2UvY29udHJvbGxlcnMvcHByb2YuZ28=) | 0.00% | [46 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [turnpike/controllers/database.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&filepath=turnpike%2Fcontrollers%2Fdatabase.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dHVybnBpa2UvY29udHJvbGxlcnMvZGF0YWJhc2UuZ28=) | 0.00% | [36 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/routes/routes.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&filepath=manager%2Froutes%2Froutes.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9yb3V0ZXMvcm91dGVzLmdv) | 0.00% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1599      +/-   ##
==========================================
- Coverage   57.98%   57.89%   -0.09%     
==========================================
  Files         143      145       +2     
  Lines       11181    11193      +12     
==========================================
- Hits         6483     6480       -3     
- Misses       4115     4130      +15     
  Partials      583      583              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.89% <0.00%> (-0.09%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1599?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @psegedy on March 11, 2025 at 04:28 PM UTC

/retest

### Comment by @psegedy on March 11, 2025 at 04:36 PM UTC

/retest

### Comment by @MichaelMraka on March 12, 2025 at 08:45 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on March 12, 2025 at 08:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1599*
