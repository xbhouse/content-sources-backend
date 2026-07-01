---
type: pull_request
number: 1100
title: "SPM-1661: rework turnpike/admin api"
state: merged
author: psegedy
created: 2022-09-08T15:43:49Z
updated: 2022-09-12T08:20:54Z
closed: 2022-09-12T08:20:53Z
merged: 2022-09-12T08:20:53Z
base_branch: master
head_branch: turnpike
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1100
---

# Pull Request #1100: SPM-1661: rework turnpike/admin api

**Author**: @psegedy
**Created**: September 08, 2022 at 03:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `turnpike`

## Description

- move admin api to its own pod
- generate openapi for admin api
- api is exposed on `/api/patch/admin`
- openapi UI is on `/api/patch/admin/openapi/index.html` on `PublicPort` or `8085` if ran locally
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

### Comment by @jira-linking on September 08, 2022 at 03:43 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1661


### Comment by @codecov-commenter on September 08, 2022 at 03:52 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.81**% // Head: **62.96**% // Increases project coverage by **`+0.14%`** :tada:
> Coverage data is based on head [(`5d2863b`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`5c02fbd`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/5c02fbdaa44a8dd146341b34181fb4ccde3f419b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 44.44% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1100      +/-   ##
==========================================
+ Coverage   62.81%   62.96%   +0.14%     
==========================================
  Files          97       97              
  Lines        5529     5535       +6     
==========================================
+ Hits         3473     3485      +12     
+ Misses       1607     1606       -1     
+ Partials      449      444       -5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.96% <44.44%> (+0.14%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <ø> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `63.85% <0.00%> (-1.58%)` | :arrow_down: |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `21.05% <0.00%> (-1.81%)` | :arrow_down: |
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `34.61% <ø> (ø)` | |
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `41.66% <0.00%> (-10.06%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `89.61% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `67.92% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `74.60% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.41% <ø> (ø)` | |
| ... and [15 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1100?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on September 12, 2022 at 08:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1100*
