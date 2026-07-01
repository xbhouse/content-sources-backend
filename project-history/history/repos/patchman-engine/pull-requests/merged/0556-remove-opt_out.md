---
type: pull_request
number: 556
title: "Remove opt_out"
state: merged
author: semtexzv
created: 2021-02-24T09:40:57Z
updated: 2021-03-16T09:30:32Z
closed: 2021-03-01T11:23:01Z
merged: 2021-03-01T11:23:01Z
base_branch: master
head_branch: remove-opt-out
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/556
---

# Pull Request #556: Remove opt_out

**Author**: @semtexzv
**Created**: February 24, 2021 at 09:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remove-opt-out`

## Description

Remove the `opt_out` column. Change its uses to the `stale` column.

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

### Comment by @josef-hak on February 24, 2021 at 12:03 PM UTC

TestSchemaCompatiblity failed.

### Comment by @codecov-io on March 01, 2021 at 10:04 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=h1) Report
> Merging [#556](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=desc) (e48cb14) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b156adde38d602fd87a7411893b9b7d281d7c71?el=desc) (4b156ad) will **decrease** coverage by `0.02%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #556      +/-   ##
==========================================
- Coverage   60.38%   60.36%   -0.03%     
==========================================
  Files          67       67              
  Lines        2941     2942       +1     
==========================================
  Hits         1776     1776              
- Misses        911      912       +1     
  Partials      254      254              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.86% <75.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <0.00%> (ø)` | |
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | | |
| [vmaas\_sync/admin\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZG1pbl9hcGkuZ28=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=footer). Last update [4b156ad...e48cb14](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/556?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on March 01, 2021 at 11:22 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/556*
