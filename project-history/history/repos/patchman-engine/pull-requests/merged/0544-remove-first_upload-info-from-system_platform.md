---
type: pull_request
number: 544
title: "remove first_upload info from system_platform"
state: merged
author: mkholjuraev
created: 2021-02-18T11:35:20Z
updated: 2021-03-10T12:33:36Z
closed: 2021-03-10T12:33:36Z
merged: 2021-03-10T12:33:36Z
base_branch: master
head_branch: drop-first-reported
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/544
---

# Pull Request #544: remove first_upload info from system_platform

**Author**: @mkholjuraev
**Created**: February 18, 2021 at 11:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `drop-first-reported`

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

### Comment by @mkholjuraev on February 24, 2021 at 12:15 PM UTC

In this PR, I wanted to remove this info from .go files. But, I could not find many lines of .go code where first_reported is used. I believe this is because why we want to get rid of it. Mainly, first_reported column is created in the database and is not used in application logic. So, I think we only need to remove it from DB in the next phase.  

### Comment by @codecov-io on March 10, 2021 at 12:07 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=h1) Report
> Merging [#544](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=desc) (af348fc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/828484e30f5c06dc0e52ce824900420a2a811f21?el=desc) (828484e) will **increase** coverage by `0.03%`.
> The diff coverage is `65.88%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #544      +/-   ##
==========================================
+ Coverage   60.33%   60.36%   +0.03%     
==========================================
  Files          66       67       +1     
  Lines        2894     2942      +48     
==========================================
+ Hits         1746     1776      +30     
- Misses        898      912      +14     
- Partials      250      254       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.67% <ø> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `52.77% <0.00%> (-1.51%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.82% <65.38%> (-0.67%)` | :arrow_down: |
| [base/mqueue/encryption.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZW5jcnlwdGlvbi5nbw==) | `80.00% <85.71%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `67.79% <100.00%> (+1.12%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `64.70% <100.00%> (+10.85%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.00% <100.00%> (-0.08%)` | :arrow_down: |
| ... and [7 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=footer). Last update [6ebd399...af348fc](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/544?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 18, 2021 at 12:15 PM UTC

Thanks, partially good :wink: 

### Review by @josef-hak - Changes Requested on March 09, 2021 at 12:38 PM UTC

### Review by @mkholjuraev - Commented on March 09, 2021 at 01:11 PM UTC

### Review by @mkholjuraev - Commented on March 09, 2021 at 02:18 PM UTC

### Review by @josef-hak - Approved on March 09, 2021 at 02:12 PM UTC

Great, now it's perfect.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/544*
