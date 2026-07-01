---
type: pull_request
number: 679
title: "SPM-827 Isolate \"segmentio/kafka-go\" usage to mqueue only"
state: merged
author: josef-hak
created: 2021-05-19T11:31:00Z
updated: 2021-05-19T12:26:52Z
closed: 2021-05-19T12:24:26Z
merged: 2021-05-19T12:24:26Z
base_branch: master
head_branch: isolate-kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/679
---

# Pull Request #679: SPM-827 Isolate "segmentio/kafka-go" usage to mqueue only

**Author**: @josef-hak
**Created**: May 19, 2021 at 11:31 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `isolate-kafka`

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

### Comment by @codecov-commenter on May 19, 2021 at 12:14 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#679](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e6f9328) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/92f43c2e6aca29d2cee9dca6488003402a435421?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (92f43c2) will **decrease** coverage by `0.09%`.
> The diff coverage is `62.35%`.

> :exclamation: Current head e6f9328 differs from pull request most recent head 6221ab5. Consider uploading reports for the commit 6221ab5 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #679      +/-   ##
==========================================
- Coverage   58.56%   58.47%   -0.10%     
==========================================
  Files          72       71       -1     
  Lines        3367     3386      +19     
==========================================
+ Hits         1972     1980       +8     
- Misses       1111     1121      +10     
- Partials      284      285       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.47% <62.35%> (-0.10%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `75.00% <0.00%> (-25.00%)` | :arrow_down: |
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `76.47% <ø> (+8.04%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.33% <0.00%> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `71.42% <0.00%> (ø)` | |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.00% <14.28%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `44.44% <22.22%> (-23.36%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `70.37% <50.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `77.19% <77.19%> (ø)` | |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <100.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [92f43c2...6221ab5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/679?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on May 19, 2021 at 11:41 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/679*
