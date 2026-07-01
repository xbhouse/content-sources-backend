---
type: pull_request
number: 1102
title: "SPM-1730: fix wrong usage of errors.Wrap"
state: merged
author: psegedy
created: 2022-09-09T13:43:46Z
updated: 2022-09-09T15:23:49Z
closed: 2022-09-09T15:23:49Z
merged: 2022-09-09T15:23:49Z
base_branch: master
head_branch: wrong_errors
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1102
---

# Pull Request #1102: SPM-1730: fix wrong usage of errors.Wrap

**Author**: @psegedy
**Created**: September 09, 2022 at 01:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `wrong_errors`

## Description

can't use errors.Wrap on err=nil 🤦 
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

### Comment by @jira-linking on September 09, 2022 at 01:43 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1730


### Comment by @codecov-commenter on September 09, 2022 at 01:52 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.81**% // Head: **63.03**% // Increases project coverage by **`+0.22%`** :tada:
> Coverage data is based on head [(`2c90da3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`5c02fbd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5c02fbdaa44a8dd146341b34181fb4ccde3f419b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 47.78% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1102      +/-   ##
==========================================
+ Coverage   62.81%   63.03%   +0.22%     
==========================================
  Files          97       97              
  Lines        5529     5522       -7     
==========================================
+ Hits         3473     3481       +8     
+ Misses       1607     1597      -10     
+ Partials      449      444       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.03% <47.78%> (+0.22%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <ø> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `63.85% <0.00%> (-1.58%)` | :arrow_down: |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `21.05% <0.00%> (-1.81%)` | :arrow_down: |
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `34.61% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `89.61% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.00% <ø> (ø)` | |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.21% <ø> (+0.38%)` | :arrow_up: |
| [tasks/vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `68.88% <ø> (+7.19%)` | :arrow_up: |
| ... and [13 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1102?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on September 09, 2022 at 01:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1102*
