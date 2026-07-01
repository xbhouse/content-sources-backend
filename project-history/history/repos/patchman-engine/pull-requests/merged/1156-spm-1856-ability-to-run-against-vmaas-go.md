---
type: pull_request
number: 1156
title: "SPM-1856: ability to run against vmaas-go"
state: merged
author: psegedy
created: 2023-01-16T14:53:10Z
updated: 2023-01-17T11:35:03Z
closed: 2023-01-17T11:35:03Z
merged: 2023-01-17T11:35:03Z
base_branch: master
head_branch: vmaas_go
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1156
---

# Pull Request #1156: SPM-1856: ability to run against vmaas-go

**Author**: @psegedy
**Created**: January 16, 2023 at 02:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas_go`

## Description

- use vmaas-go by setting USE_VMAAS_GO env var for evaluator
- sort AvailableUpdates because sorted values are needed for merging with yum_updates
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

### Comment by @jira-linking on January 16, 2023 at 03:08 PM UTC

Commits missing Jira IDs:
2949db6f29da08c326043df983d154df22005b36
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1856


### Comment by @codecov-commenter on January 16, 2023 at 03:16 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.18**% // Head: **63.20**% // Increases project coverage by **`+1.01%`** :tada:
> Coverage data is based on head [(`a300ad6`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`1109e51`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1109e515a37f3b44e2cea97dcb98d5bbc0b36bbe?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 81.43% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1156      +/-   ##
==========================================
+ Coverage   62.18%   63.20%   +1.01%     
==========================================
  Files         100      101       +1     
  Lines        5776     5770       -6     
==========================================
+ Hits         3592     3647      +55     
+ Misses       1723     1666      -57     
+ Partials      461      457       -4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.20% <81.43%> (+1.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `76.47% <0.00%> (ø)` | |
| [manager/middlewares/db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9kYi5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/caches/refresh\_packages\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `26.98% <16.66%> (+1.63%)` | :arrow_up: |
| [tasks/vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `68.88% <25.00%> (ø)` | |
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `41.66% <40.00%> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `69.23% <66.66%> (+1.23%)` | :arrow_up: |
| [manager/controllers/package\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXNfZXhwb3J0Lmdv) | `53.57% <66.66%> (+1.71%)` | :arrow_up: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `63.04% <73.33%> (+30.59%)` | :arrow_up: |
| ... and [41 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1156?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on January 17, 2023 at 11:34 AM UTC

There's an issue in PR check, the failures are the same as in https://github.com/RedHatInsights/patchman-engine/pull/1157

---

## Reviews

### Review by @MichaelMraka - Approved on January 16, 2023 at 04:13 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1156*
