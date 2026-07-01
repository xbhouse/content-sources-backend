---
type: pull_request
number: 1106
title: "SPM-1744: simplify semantic release"
state: merged
author: MichaelMraka
created: 2022-09-15T12:46:38Z
updated: 2022-09-15T14:34:02Z
closed: 2022-09-15T14:34:02Z
merged: 2022-09-15T14:34:01Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1106
---

# Pull Request #1106: SPM-1744: simplify semantic release

**Author**: @MichaelMraka
**Created**: September 15, 2022 at 12:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

updating .go files breaks image caching

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

### Comment by @jira-linking on September 15, 2022 at 12:46 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1744


### Comment by @codecov-commenter on September 15, 2022 at 12:55 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.87**% // Head: **62.90**% // Increases project coverage by **`+0.03%`** :tada:
> Coverage data is based on head [(`8a9d5f1`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`146efec`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/146efeca7ce7bae151b4bc75c056227eb7c7fde5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 80.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1106      +/-   ##
==========================================
+ Coverage   62.87%   62.90%   +0.03%     
==========================================
  Files          97       97              
  Lines        5535     5548      +13     
==========================================
+ Hits         3480     3490      +10     
+ Misses       1611     1609       -2     
- Partials      444      449       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.90% <80.00%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `9.96% <0.00%> (ø)` | |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `75.80% <ø> (ø)` | |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <ø> (ø)` | |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.41% <76.19%> (+0.25%)` | :arrow_up: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.69% <100.00%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.26% <100.00%> (ø)` | |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.90% <100.00%> (ø)` | |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `71.87% <100.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `81.31% <100.00%> (ø)` | |
| ... and [4 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1106?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on September 15, 2022 at 01:54 PM UTC

LGTM, please rebase it, I merged a PR which modified manager.go admin_api.go files

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1106*
