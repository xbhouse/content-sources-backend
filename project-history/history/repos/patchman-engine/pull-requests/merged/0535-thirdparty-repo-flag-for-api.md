---
type: pull_request
number: 535
title: "thirdparty repo flag for API"
state: merged
author: MichaelMraka
created: 2021-02-15T14:03:33Z
updated: 2021-02-19T12:58:29Z
closed: 2021-02-19T12:58:29Z
merged: 2021-02-19T12:58:29Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/535
---

# Pull Request #535: thirdparty repo flag for API

**Author**: @MichaelMraka
**Created**: February 15, 2021 at 02:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [ ] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [ ] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Discussion

### Comment by @codecov-io on February 18, 2021 at 10:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=h1) Report
> Merging [#535](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=desc) (62a92f8) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/828484e30f5c06dc0e52ce824900420a2a811f21?el=desc) (828484e) will **decrease** coverage by `0.17%`.
> The diff coverage is `60.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #535      +/-   ##
==========================================
- Coverage   60.33%   60.15%   -0.18%     
==========================================
  Files          66       66              
  Lines        2894     2919      +25     
==========================================
+ Hits         1746     1756      +10     
- Misses        898      910      +12     
- Partials      250      253       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.67% <ø> (ø)` | |
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `52.77% <0.00%> (-1.51%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.82% <65.38%> (-0.67%)` | :arrow_down: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `64.70% <100.00%> (+10.85%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.00% <100.00%> (-0.08%)` | :arrow_down: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.14% <100.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=footer). Last update [ade88e3...62a92f8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/535?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 18, 2021 at 10:52 AM UTC

I like it. Just couple of small proposals made.

### Review by @MichaelMraka - Commented on February 19, 2021 at 10:41 AM UTC

### Review by @MichaelMraka - Commented on February 19, 2021 at 10:51 AM UTC

### Review by @MichaelMraka - Commented on February 19, 2021 at 10:51 AM UTC

### Review by @MichaelMraka - Commented on February 19, 2021 at 10:52 AM UTC

### Review by @josef-hak - Approved on February 19, 2021 at 11:38 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/535*
