---
type: pull_request
number: 550
title: "Add RHSMVersion to systems endpoint"
state: merged
author: AlesKas
created: 2021-02-23T08:40:51Z
updated: 2021-03-10T09:23:58Z
closed: 2021-03-09T12:42:17Z
merged: 2021-03-09T12:42:17Z
base_branch: master
head_branch: SPM-756
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/550
---

# Pull Request #550: Add RHSMVersion to systems endpoint

**Author**: @AlesKas
**Created**: February 23, 2021 at 08:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-756`

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

### Comment by @codecov-io on February 23, 2021 at 10:02 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=h1) Report
> Merging [#550](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=desc) (05621a5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/828484e30f5c06dc0e52ce824900420a2a811f21?el=desc) (828484e) will **increase** coverage by `0.05%`.
> The diff coverage is `65.21%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #550      +/-   ##
==========================================
+ Coverage   60.33%   60.38%   +0.05%     
==========================================
  Files          66       67       +1     
  Lines        2894     2941      +47     
==========================================
+ Hits         1746     1776      +30     
- Misses        898      911      +13     
- Partials      250      254       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `66.66% <ø> (+6.06%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.67% <ø> (ø)` | |
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `52.77% <0.00%> (-1.51%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.82% <65.38%> (-0.67%)` | :arrow_down: |
| [base/mqueue/encryption.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZW5jcnlwdGlvbi5nbw==) | `80.00% <80.00%> (ø)` | |
| ... and [8 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=footer). Last update [7f0cc18...05621a5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/550?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 23, 2021 at 09:39 AM UTC

Thanks, just a few things to update.

### Review by @josef-hak - Changes Requested on February 24, 2021 at 11:16 AM UTC

Just one last ask.

### Review by @AlesKas - Commented on February 24, 2021 at 11:17 AM UTC

### Review by @josef-hak - Changes Requested on February 24, 2021 at 02:57 PM UTC

Great, but now there are some conflicts probably with @MichaelMraka 's updates. So please resolve it.

### Review by @josef-hak - Approved on March 09, 2021 at 12:42 PM UTC

well done, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/550*
