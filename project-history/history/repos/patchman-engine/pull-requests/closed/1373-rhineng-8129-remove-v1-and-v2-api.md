---
type: pull_request
number: 1373
title: "RHINENG-8129: Remove v1 and v2 API"
state: closed
author: Dugowitch
created: 2024-02-20T13:21:02Z
updated: 2024-02-20T14:59:06Z
closed: 2024-02-20T14:59:06Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1373
---

# Pull Request #1373: RHINENG-8129: Remove v1 and v2 API

**Author**: @Dugowitch
**Created**: February 20, 2024 at 01:21 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `master`

## Description

- RHINENG-8130: remove v1,v2 compatibility from /advisories

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

### Comment by @jira-linking on February 20, 2024 at 01:21 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8130
https://issues.redhat.com/browse/RHINENG-8131
https://issues.redhat.com/browse/RHINENG-8132


### Comment by @codecov-commenter on February 20, 2024 at 01:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `7 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`8fb570d`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8fb570d40d637281fe0fc7b19b68cf70e38a4447?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.97% compared to head [(`572b2c7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.60%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/advisory\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | 80.76% | [3 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | 60.00% | [0 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1373      +/-   ##
==========================================
+ Coverage   61.97%   62.60%   +0.63%     
==========================================
  Files         108      108              
  Lines        6832     6712     -120     
==========================================
- Hits         4234     4202      -32     
+ Misses       2063     1985      -78     
+ Partials      535      525      -10     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.60% <85.10%> (+0.63%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1373?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1373*
