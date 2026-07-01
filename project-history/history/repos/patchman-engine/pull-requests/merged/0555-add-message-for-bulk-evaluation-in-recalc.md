---
type: pull_request
number: 555
title: "Add message for bulk evaluation in recalc"
state: merged
author: semtexzv
created: 2021-02-24T09:30:40Z
updated: 2021-03-16T09:30:32Z
closed: 2021-03-15T10:00:53Z
merged: 2021-03-15T10:00:53Z
base_branch: master
head_branch: bulk-recalc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/555
---

# Pull Request #555: Add message for bulk evaluation in recalc

**Author**: @semtexzv
**Created**: February 24, 2021 at 09:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bulk-recalc`

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

### Comment by @codecov-io on March 01, 2021 at 10:08 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=h1) Report
> Merging [#555](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=desc) (4f7d7c4) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b156adde38d602fd87a7411893b9b7d281d7c71?el=desc) (4b156ad) will **increase** coverage by `0.06%`.
> The diff coverage is `71.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #555      +/-   ##
==========================================
+ Coverage   60.38%   60.45%   +0.06%     
==========================================
  Files          67       67              
  Lines        2941     2951      +10     
==========================================
+ Hits         1776     1784       +8     
- Misses        911      913       +2     
  Partials      254      254              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.08% <42.85%> (+0.25%)` | :arrow_up: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `56.41% <100.00%> (+3.63%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <0.00%> (ø)` | |
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | | |
| [vmaas\_sync/admin\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZG1pbl9hcGkuZ28=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=footer). Last update [4b156ad...4f7d7c4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/555?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on February 25, 2021 at 10:41 AM UTC

there's and issue with TestSync

### Review by @MichaelMraka - Approved on March 01, 2021 at 11:38 AM UTC

### Review by @MichaelMraka - Commented on March 03, 2021 at 11:30 AM UTC

### Review by @MichaelMraka - Commented on March 12, 2021 at 01:51 PM UTC

### Review by @MichaelMraka - Approved on March 15, 2021 at 10:00 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/555*
