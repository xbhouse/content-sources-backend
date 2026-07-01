---
type: pull_request
number: 680
title: "SPM-910: image build improvements"
state: merged
author: MichaelMraka
created: 2021-05-20T14:54:13Z
updated: 2021-05-21T13:57:25Z
closed: 2021-05-21T13:57:25Z
merged: 2021-05-21T13:57:25Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/680
---

# Pull Request #680: SPM-910: image build improvements

**Author**: @MichaelMraka
**Created**: May 20, 2021 at 02:54 PM UTC
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

### Comment by @codecov-commenter on May 20, 2021 at 06:09 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#680](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (035631e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/92f43c2e6aca29d2cee9dca6488003402a435421?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (92f43c2) will **decrease** coverage by `0.06%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #680      +/-   ##
==========================================
- Coverage   58.56%   58.50%   -0.07%     
==========================================
  Files          72       72              
  Lines        3367     3379      +12     
==========================================
+ Hits         1972     1977       +5     
- Misses       1111     1118       +7     
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.50% <ø> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `75.00% <0.00%> (-25.00%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `51.28% <0.00%> (-16.52%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.00% <0.00%> (ø)` | |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <0.00%> (ø)` | |
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/encryption.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZW5jcnlwdGlvbi5nbw==) | | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `78.84% <0.00%> (ø)` | |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `76.47% <0.00%> (+8.04%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [92f43c2...035631e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/680?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on May 21, 2021 at 01:16 PM UTC

compose version fixed

---

## Reviews

### Review by @josef-hak - Changes Requested on May 21, 2021 at 01:09 PM UTC

Please change `version: '3'` in both docker-compose configs to `version: '3.4'`. With docker-compose it fails with the error:
~~~
services.platform.build contains unsupported option: 'target'
~~~
Thanks.

### Review by @josef-hak - Approved on May 21, 2021 at 01:29 PM UTC

great, smart, beautiful, brief, thanks :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/680*
