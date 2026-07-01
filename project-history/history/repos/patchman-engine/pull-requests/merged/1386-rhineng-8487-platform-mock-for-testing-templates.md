---
type: pull_request
number: 1386
title: "RHINENG-8487: platform mock for testing templates"
state: merged
author: psegedy
created: 2024-03-04T16:14:49Z
updated: 2024-03-07T15:27:59Z
closed: 2024-03-07T15:27:59Z
merged: 2024-03-07T15:27:59Z
base_branch: master
head_branch: template_mock
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1386
---

# Pull Request #1386: RHINENG-8487: platform mock for testing templates

**Author**: @psegedy
**Created**: March 04, 2024 at 04:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `template_mock`

## Description

I wanted to use the struct by importing content-sources but it installed many dependencies and updated some others we are currently using. There was a problem with 1 openapi dependency which is incompatible with our current code. We could import structs from content-sources after we update our dependencies.

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

### Comment by @jira-linking on March 04, 2024 at 04:14 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8487


### Comment by @codecov-commenter on March 04, 2024 at 06:57 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `66.66667%` with `1 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.40%. Comparing base [(`16901ef`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/16901ef0fc50e862c639177b8a795ee82e504adb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`dd2d738`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | 66.66% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1386      +/-   ##
==========================================
+ Coverage   63.36%   63.40%   +0.04%     
==========================================
  Files         107      107              
  Lines        6537     6540       +3     
==========================================
+ Hits         4142     4147       +5     
+ Misses       1891     1889       -2     
  Partials      504      504              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.40% <66.66%> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1386?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 07, 2024 at 08:51 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1386*
