---
type: pull_request
number: 1611
title: "RHINENG-16748: replace json.Marshal/Unmarshal with sonic"
state: merged
author: psegedy
created: 2025-03-18T12:55:26Z
updated: 2026-04-01T17:45:22Z
closed: 2025-03-19T10:48:29Z
merged: 2025-03-19T10:48:29Z
base_branch: master
head_branch: sonic
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1611
---

# Pull Request #1611: RHINENG-16748: replace json.Marshal/Unmarshal with sonic

**Author**: @psegedy
**Created**: March 18, 2025 at 12:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sonic`

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

### Comment by @jira-linking on March 18, 2025 at 12:55 PM UTC

Commits missing Jira IDs:
f123806bdf8ce3ca6bcb815b2f8f49b3f8746acf
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16748


### Comment by @codecov-commenter on March 18, 2025 at 01:01 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `64.70588%` with `18 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 58.08%. Comparing base ([`b4dbe06`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b4dbe0617a68431a38dde5d29f39ef862c0bcfbb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e30c6fa`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e30c6fae46cf2cc4fff51ad4589ef5b1d38d0c7d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 929 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [platform/platform.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=platform%2Fplatform.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-cGxhdGZvcm0vcGxhdGZvcm0uZ28=) | 0.00% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/events.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=listener%2Fevents.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | 0.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/mqueue/message.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Fmqueue%2Fmessage.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/mqueue/payload\_tracker\_event.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Fmqueue%2Fpayload_tracker_event.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/rbac/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Frbac%2Frbac.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9yYmFjL3JiYWMuZ28=) | 50.00% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/types/timestamp.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Ftypes%2Ftimestamp.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS90eXBlcy90aW1lc3RhbXAuZ28=) | 66.66% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/core.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Futils%2Fcore.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=base%2Futils%2Fvmaas.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/notifications.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=evaluator%2Fnotifications.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | 0.00% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/templates.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&filepath=listener%2Ftemplates.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdGVtcGxhdGVzLmdv) | 0.00% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| ... and [2 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1611      +/-   ##
==========================================
+ Coverage   58.03%   58.08%   +0.04%     
==========================================
  Files         145      145              
  Lines       11270    11271       +1     
==========================================
+ Hits         6541     6547       +6     
+ Misses       4142     4139       -3     
+ Partials      587      585       -2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.08% <64.70%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1611?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 19, 2025 at 08:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1611*
