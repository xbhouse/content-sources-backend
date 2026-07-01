---
type: pull_request
number: 893
title: "SPM-1349: (1) Replace generated vmaas client"
state: merged
author: josef-hak
created: 2022-01-27T17:33:12Z
updated: 2022-02-02T12:55:01Z
closed: 2022-02-02T12:55:00Z
merged: 2022-02-02T12:55:00Z
base_branch: master
head_branch: rm-vmaas-client
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/893
---

# Pull Request #893: SPM-1349: (1) Replace generated vmaas client

**Author**: @josef-hak
**Created**: January 27, 2022 at 05:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rm-vmaas-client`

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

### Comment by @josef-hak on January 27, 2022 at 10:31 PM UTC

/retest

### Comment by @codecov-commenter on January 28, 2022 at 08:15 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#893](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (10b43a8) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4d1f24e7a6654d7f07b1586895cc0008eadb50a2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4d1f24e) will **decrease** coverage by `0.11%`.
> The diff coverage is `51.06%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #893      +/-   ##
==========================================
- Coverage   59.26%   59.15%   -0.12%     
==========================================
  Files          88       88              
  Lines        4608     4639      +31     
==========================================
+ Hits         2731     2744      +13     
- Misses       1493     1510      +17     
- Partials      384      385       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.15% <51.06%> (-0.12%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | `34.61% <0.00%> (-12.76%)` | :arrow_down: |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.33% <ø> (ø)` | |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `73.93% <ø> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `78.57% <ø> (ø)` | |
| [manager/middlewares/swagger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9zd2FnZ2VyLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `41.66% <0.00%> (ø)` | |
| [docs/docs.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZG9jcy9kb2NzLmdv) | `65.21% <71.42%> (+65.21%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.16% <100.00%> (-0.29%)` | :arrow_down: |
| [evaluator/evaluate\_baseline.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | `83.33% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b148ecf...10b43a8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/893?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on January 31, 2022 at 06:44 PM UTC

/retest

### Comment by @josef-hak on February 01, 2022 at 07:09 PM UTC

/retest

### Comment by @digitronik on February 02, 2022 at 06:58 AM UTC

/retest

### Comment by @digitronik on February 02, 2022 at 09:19 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on February 02, 2022 at 12:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/893*
