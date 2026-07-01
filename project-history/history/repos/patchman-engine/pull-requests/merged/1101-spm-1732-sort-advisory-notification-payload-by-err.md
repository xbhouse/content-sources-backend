---
type: pull_request
number: 1101
title: "SPM-1732: Sort advisory notification payload by errata name"
state: merged
author: michalslomczynski
created: 2022-09-09T12:28:23Z
updated: 2022-09-12T08:42:38Z
closed: 2022-09-12T08:42:38Z
merged: 2022-09-12T08:42:38Z
base_branch: master
head_branch: advisory-sort
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1101
---

# Pull Request #1101: SPM-1732: Sort advisory notification payload by errata name

**Author**: @michalslomczynski
**Created**: September 09, 2022 at 12:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisory-sort`

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

### Comment by @jira-linking on September 09, 2022 at 12:28 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1732


### Comment by @codecov-commenter on September 09, 2022 at 12:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.81**% // Head: **63.05**% // Increases project coverage by **`+0.24%`** :tada:
> Coverage data is based on head [(`ebd3918`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`5c02fbd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5c02fbdaa44a8dd146341b34181fb4ccde3f419b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 50.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1101      +/-   ##
==========================================
+ Coverage   62.81%   63.05%   +0.24%     
==========================================
  Files          97       97              
  Lines        5529     5525       -4     
==========================================
+ Hits         3473     3484      +11     
+ Misses       1607     1597      -10     
+ Partials      449      444       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.05% <50.00%> (+0.24%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <ø> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `63.85% <0.00%> (-1.58%)` | :arrow_down: |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `21.05% <0.00%> (-1.81%)` | :arrow_down: |
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `34.61% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `89.61% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `67.92% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `74.60% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.41% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.00% <ø> (ø)` | |
| ... and [14 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1101?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on September 09, 2022 at 01:52 PM UTC

### Review by @MichaelMraka - Approved on September 12, 2022 at 08:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1101*
