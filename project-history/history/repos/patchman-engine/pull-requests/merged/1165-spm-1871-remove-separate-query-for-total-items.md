---
type: pull_request
number: 1165
title: "SPM-1871: remove separate query for total items"
state: merged
author: MichaelMraka
created: 2023-01-25T11:09:44Z
updated: 2023-01-31T11:20:47Z
closed: 2023-01-31T11:20:47Z
merged: 2023-01-31T11:20:47Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1165
---

# Pull Request #1165: SPM-1871: remove separate query for total items

**Author**: @MichaelMraka
**Created**: January 25, 2023 at 11:09 AM UTC
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

### Comment by @jira-linking on January 26, 2023 at 11:55 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1871


### Comment by @codecov-commenter on January 27, 2023 at 12:50 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.17**% // Head: **63.38**% // Increases project coverage by **`+0.20%`** :tada:
> Coverage data is based on head [(`795cbac`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`9ce9b5f`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/9ce9b5fd1fa8b214ba3351d9e16d4a34fcdef275?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 91.91% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1165      +/-   ##
==========================================
+ Coverage   63.17%   63.38%   +0.20%     
==========================================
  Files         102      102              
  Lines        5755     5853      +98     
==========================================
+ Hits         3636     3710      +74     
- Misses       1661     1678      +17     
- Partials      458      465       +7     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.38% <91.91%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.01% <0.00%> (-0.65%)` | :arrow_down: |
| [tasks/caches/refresh\_packages\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXNfY2FjaGVzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `66.10% <81.81%> (-0.57%)` | :arrow_down: |
| [manager/controllers/systemtags.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW10YWdzLmdv) | `80.55% <83.33%> (-3.82%)` | :arrow_down: |
| [manager/controllers/package\_versions.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3ZlcnNpb25zLmdv) | `69.04% <84.61%> (+1.40%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `75.00% <86.66%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `70.45% <92.85%> (+0.18%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.14% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `90.47% <100.00%> (+0.60%)` | :arrow_up: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `71.42% <100.00%> (+2.19%)` | :arrow_up: |
| ... and [14 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1165?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on January 27, 2023 at 06:02 PM UTC

@MichaelMraka `system/<uuid>/packages` is always returning `total=0` in tests but I haven't find the issue in the code yet.

---

## Reviews

### Review by @psegedy - Changes Requested on January 30, 2023 at 05:19 PM UTC

### Review by @MichaelMraka - Commented on January 31, 2023 at 08:32 AM UTC

### Review by @psegedy - Approved on January 31, 2023 at 09:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1165*
