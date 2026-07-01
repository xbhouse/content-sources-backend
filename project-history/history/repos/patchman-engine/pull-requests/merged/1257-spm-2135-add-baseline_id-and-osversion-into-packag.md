---
type: pull_request
number: 1257
title: "SPM-2135: add baseline_id and osversion into /package/systems"
state: merged
author: MichaelMraka
created: 2023-07-04T14:42:19Z
updated: 2023-07-10T14:38:32Z
closed: 2023-07-10T14:38:32Z
merged: 2023-07-10T14:38:31Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1257
---

# Pull Request #1257: SPM-2135: add baseline_id and osversion into /package/systems

**Author**: @MichaelMraka
**Created**: July 04, 2023 at 02:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @jira-linking on July 04, 2023 at 02:42 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2135


### Comment by @codecov-commenter on July 04, 2023 at 02:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.05`** :warning:
> Comparison is base [(`c6dca18`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c6dca18aa76c0ebf975e231e024507e9b11818b6?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.34% compared to head [(`7081451`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.29%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1257      +/-   ##
==========================================
- Coverage   61.34%   61.29%   -0.05%     
==========================================
  Files         106      106              
  Lines        6540     6542       +2     
==========================================
- Hits         4012     4010       -2     
- Misses       2008     2011       +3     
- Partials      520      521       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.29% <0.00%> (-0.05%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.94% <0.00%> (-0.47%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `62.22% <ø> (ø)` | |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1257?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Changes Requested on July 10, 2023 at 07:55 AM UTC

### Review by @MichaelMraka - Commented on July 10, 2023 at 11:49 AM UTC

### Review by @psegedy - Approved on July 10, 2023 at 02:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1257*
