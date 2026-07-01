---
type: pull_request
number: 678
title: "SPM-911 Add kafka authentication support"
state: merged
author: josef-hak
created: 2021-05-19T09:13:05Z
updated: 2021-08-26T18:41:44Z
closed: 2021-06-08T08:44:18Z
merged: 2021-06-08T08:44:18Z
base_branch: master
head_branch: kafka-conf
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/678
---

# Pull Request #678: SPM-911 Add kafka authentication support

**Author**: @josef-hak
**Created**: May 19, 2021 at 09:13 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka-conf`

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

### Comment by @codecov-commenter on May 19, 2021 at 09:22 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#678](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (66c8616) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f361e5e4c9e5a56956ed2ea70a4a51f4a837b802?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f361e5e) will **decrease** coverage by `0.89%`.
> The diff coverage is `14.13%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #678      +/-   ##
==========================================
- Coverage   57.45%   56.55%   -0.90%     
==========================================
  Files          74       75       +1     
  Lines        3455     3517      +62     
==========================================
+ Hits         1985     1989       +4     
- Misses       1187     1245      +58     
  Partials      283      283              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.55% <14.13%> (-0.90%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `56.66% <0.00%> (+5.38%)` | :arrow_up: |
| [base/mqueue/mqueue\_impl\_confluentic.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfY29uZmx1ZW50aWMuZ28=) | `0.00% <0.00%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.33% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `63.15% <29.41%> (-15.69%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `70.37% <50.00%> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `78.57% <100.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `41.02% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f361e5e...66c8616](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/678?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 07, 2021 at 02:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/678*
