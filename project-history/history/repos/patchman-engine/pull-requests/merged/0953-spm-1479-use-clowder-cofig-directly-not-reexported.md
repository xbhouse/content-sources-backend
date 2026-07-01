---
type: pull_request
number: 953
title: "SPM-1479: use clowder cofig directly not reexported to env"
state: merged
author: MichaelMraka
created: 2022-05-10T13:39:46Z
updated: 2022-05-19T15:57:21Z
closed: 2022-05-19T15:57:21Z
merged: 2022-05-19T15:57:21Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/953
---

# Pull Request #953: SPM-1479: use clowder cofig directly not reexported to env

**Author**: @MichaelMraka
**Created**: May 10, 2022 at 01:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @codecov-commenter on May 12, 2022 at 01:34 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#953](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9f65cb1) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f1a89b80985fe10be4b8a756c3dcf9face1e70cd?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f1a89b8) will **increase** coverage by `0.06%`.
> The diff coverage is `48.82%`.

```diff
@@            Coverage Diff             @@
##           master     #953      +/-   ##
==========================================
+ Coverage   60.60%   60.66%   +0.06%     
==========================================
  Files          92       92              
  Lines        4871     4955      +84     
==========================================
+ Hits         2952     3006      +54     
- Misses       1523     1550      +27     
- Partials      396      399       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.66% <48.82%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `27.69% <0.00%> (+16.40%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `76.47% <0.00%> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `71.42% <0.00%> (-7.15%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.08% <0.00%> (ø)` | |
| [base/utils/awscloudwatch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9hd3NjbG91ZHdhdGNoLmdv) | `13.79% <25.00%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `32.37% <38.79%> (+32.37%)` | :arrow_up: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `43.20% <71.42%> (+0.11%)` | :arrow_up: |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `65.00% <84.61%> (ø)` | |
| ... and [9 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f1a89b8...9f65cb1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/953?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on May 19, 2022 at 03:57 PM UTC

LGTM, works fine in ephemeral

---

## Reviews

### Review by @psegedy - Commented on May 18, 2022 at 11:00 AM UTC

### Review by @psegedy - Approved on May 19, 2022 at 03:56 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/953*
