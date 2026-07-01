---
type: pull_request
number: 1155
title: "SPM-1848: improve packace cache computation"
state: merged
author: MichaelMraka
created: 2023-01-12T13:50:41Z
updated: 2023-01-16T10:14:54Z
closed: 2023-01-16T10:14:54Z
merged: 2023-01-16T10:14:54Z
base_branch: master
head_branch: pr5
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1155
---

# Pull Request #1155: SPM-1848: improve packace cache computation

**Author**: @MichaelMraka
**Created**: January 12, 2023 at 01:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr5`

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

### Comment by @jira-linking on January 12, 2023 at 01:50 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1848


### Comment by @codecov-commenter on January 12, 2023 at 01:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.18**% // Head: **63.06**% // Increases project coverage by **`+0.87%`** :tada:
> Coverage data is based on head [(`32e8a46`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`1109e51`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1109e515a37f3b44e2cea97dcb98d5bbc0b36bbe?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 58.18% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1155      +/-   ##
==========================================
+ Coverage   62.18%   63.06%   +0.87%     
==========================================
  Files         100      100              
  Lines        5776     5758      -18     
==========================================
+ Hits         3592     3631      +39     
+ Misses       1723     1667      -56     
+ Partials      461      460       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.06% <58.18%> (+0.87%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `76.47% <0.00%> (ø)` | |
| [tasks/caches/refresh\_packages\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `26.98% <16.66%> (+1.63%)` | :arrow_up: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `63.70% <75.60%> (+31.25%)` | :arrow_up: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `74.28% <0.00%> (+1.42%)` | :arrow_up: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `78.57% <0.00%> (+7.14%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1155?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on January 13, 2023 at 04:09 PM UTC

I have no idea how the magic in hash_partition_id works, so lgtm 😄 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1155*
