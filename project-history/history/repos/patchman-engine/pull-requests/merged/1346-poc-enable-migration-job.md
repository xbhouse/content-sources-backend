---
type: pull_request
number: 1346
title: "POC: enable migration job"
state: merged
author: psegedy
created: 2023-11-22T12:17:37Z
updated: 2023-11-23T09:21:39Z
closed: 2023-11-23T09:21:39Z
merged: 2023-11-23T09:21:39Z
base_branch: poc
head_branch: enable_migration_job
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1346
---

# Pull Request #1346: POC: enable migration job

**Author**: @psegedy
**Created**: November 22, 2023 at 12:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `poc` ← **Head**: `enable_migration_job`

## Description

if the job is not enabled, clowder won't create job resource and we cannot run it
reconciliation does not seem to be picking up disabled parameter, so it never creates the job 
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

### Comment by @jira-linking on November 22, 2023 at 12:17 PM UTC

Commits missing Jira IDs:
1dd2d075aefd00be333291d556cf423cda6ce17b


### Comment by @codecov-commenter on November 22, 2023 at 12:25 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `103 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`28f08b2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/28f08b23a52c9fdc1b222566f99824b467d71d2f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.49% compared to head [(`1dd2d07`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.03%.
> Report is 82 commits behind head on poc.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/middlewares/limits.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9saW1pdHMuZ28=) | 0.00% | [16 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 85.54% | [7 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/systems\_advisories\_view.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | 77.35% | [10 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 79.59% | [6 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/core.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | 37.50% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | 75.00% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 88.88% | [2 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/advisory\_systems\_v3.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | 33.33% | [2 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | 71.42% | [2 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| ... and [14 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##              poc    #1346      +/-   ##
==========================================
- Coverage   62.49%   62.03%   -0.47%     
==========================================
  Files         106      108       +2     
  Lines        6684     6811     +127     
==========================================
+ Hits         4177     4225      +48     
- Misses       1985     2052      +67     
- Partials      522      534      +12     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.03% <77.75%> (-0.47%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1346?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on November 23, 2023 at 08:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1346*
