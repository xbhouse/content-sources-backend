---
type: pull_request
number: 68
title: "Added advisory_account_data update to evaluation"
state: merged
author: josef-hak
created: 2020-01-20T15:00:34Z
updated: 2020-01-24T08:38:42Z
closed: 2020-01-22T13:31:59Z
merged: 2020-01-22T13:31:59Z
base_branch: master
head_branch: account_data_update
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/68
---

# Pull Request #68: Added advisory_account_data update to evaluation

**Author**: @josef-hak
**Created**: January 20, 2020 at 03:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `account_data_update`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 22, 2020 at 11:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=h1) Report
> Merging [#68](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/7b50ee7860a8705d24b24a44b64789aca13e8236?src=pr&el=desc) will **increase** coverage by `0.38%`.
> The diff coverage is `80.21%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #68      +/-   ##
==========================================
+ Coverage   57.28%   57.67%   +0.38%     
==========================================
  Files          30       30              
  Lines        1187     1212      +25     
==========================================
+ Hits          680      699      +19     
- Misses        427      434       +7     
+ Partials       80       79       -1
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0% <0%> (ø)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `82.08% <100%> (+2.77%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `69.31% <66.66%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `74.82% <87.5%> (+2.88%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=footer). Last update [7b50ee7...7e23425](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/68?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on January 21, 2020 at 01:06 PM UTC

### Review by @josef-hak - Commented on January 21, 2020 at 02:24 PM UTC

### Review by @semtexzv - Approved on January 22, 2020 at 01:31 PM UTC

While this PR is pretty complex, everything seems fine. :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/68*
