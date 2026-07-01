---
type: pull_request
number: 1159
title: "SPM-1862: set timeout for requests"
state: merged
author: psegedy
created: 2023-01-19T11:58:41Z
updated: 2023-01-20T09:41:13Z
closed: 2023-01-20T09:41:13Z
merged: 2023-01-20T09:41:13Z
base_branch: master
head_branch: timouts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1159
---

# Pull Request #1159: SPM-1862: set timeout for requests

**Author**: @psegedy
**Created**: January 19, 2023 at 11:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `timouts`

## Description

- AbortWithStatusJSON only sets `c.Abort()`
- it looks like it is needed to call `c.Done()` so the DB context will be cancelled as well
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

### Comment by @jira-linking on January 19, 2023 at 11:58 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1862


### Comment by @codecov-commenter on January 19, 2023 at 01:02 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.18**% // Head: **63.05**% // Decreases project coverage by **`-0.13%`** :warning:
> Coverage data is based on head [(`d9cfeb5`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`a42c9cd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a42c9cdb8d1171a5dabefb602ca39bb253c1f9bc?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 10.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1159      +/-   ##
==========================================
- Coverage   63.18%   63.05%   -0.13%     
==========================================
  Files         101      102       +1     
  Lines        5772     5782      +10     
==========================================
- Hits         3647     3646       -1     
- Misses       1667     1678      +11     
  Partials      458      458              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.05% <10.00%> (-0.13%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/timeout.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy90aW1lb3V0Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `63.30% <100.00%> (+0.26%)` | :arrow_up: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.22% <0.00%> (-2.78%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1159?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on January 19, 2023 at 03:18 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1159*
