---
type: pull_request
number: 1022
title: "SPM-1559: Refactor manager/controllers tests"
state: closed
author: Mischulee
created: 2022-07-12T16:25:05Z
updated: 2022-08-01T11:55:15Z
closed: 2022-08-01T11:55:15Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1022
---

# Pull Request #1022: SPM-1559: Refactor manager/controllers tests

**Author**: @Mischulee
**Created**: July 12, 2022 at 04:25 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `master`

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

### Comment by @jira-linking on July 12, 2022 at 04:25 PM UTC

Commits missing Jira IDs:
1301f8593a42ceac2c596faba29a02f2b7874dcc
a034f09352052e6741b3b34ffa5e30d367626da0
f3ddc16f1f0f6a0d6cc5889f9cd4a4ab300e71d1
45f6f28f6ed985b8b77b8f28750fdf732198bb5a
667c3964d3a09830f0c14b6444b6c7705a5aee1a
4c6cac210b236a7bca7b739452610c9b2fdc5027
ee643a8bc01d37764cca97ffc6b51734cab04e13
d45650c310a24677863bfd48ee8112c4cf0887f8
71344c1d47cf37f7c8a829a9dae64146a46cefd9
4d6e928f8fc7753984ea23a2f76e6c1e3ac360c9
fd7bb8f3e8f0017a4695ac567bad63c3ff86a45d
364d28941f566f31295ebaddac4a297016ce41d7
202edcb69de5da674f46426c47fef1febe091435
60acb85ecc06eebfef4aa712e1c05ff36cbe6b8d
56c80403c1958f7a4655b2470b2323777fd7853e
0cde227a54c9d60da87d16b727ee4fa978ae2cc3
469a8f00613cc85d8f3818fa371a0927e3efcc29
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1559
https://issues.redhat.com/browse/SPM-1626
https://issues.redhat.com/browse/SPM-1633
https://issues.redhat.com/browse/SPM-1636
https://issues.redhat.com/browse/SPM-1616
https://issues.redhat.com/browse/SPM-1630
https://issues.redhat.com/browse/SPM-1643
https://issues.redhat.com/browse/SPM-1638
https://issues.redhat.com/browse/SPM-1648
https://issues.redhat.com/browse/SPM-1649
https://issues.redhat.com/browse/SPM-1655


### Comment by @codecov-commenter on July 12, 2022 at 04:44 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1022](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2c9980c) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/af7aeaa935907531686a4c89f0c742338f3b5028?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (af7aeaa) will **decrease** coverage by `0.02%`.
> The diff coverage is `65.00%`.

```diff
@@            Coverage Diff             @@
##           master    #1022      +/-   ##
==========================================
- Coverage   61.32%   61.30%   -0.03%     
==========================================
  Files          95       95              
  Lines        5396     5411      +15     
==========================================
+ Hits         3309     3317       +8     
- Misses       1661     1670       +9     
+ Partials      426      424       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.30% <65.00%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `25.35% <0.00%> (-2.35%)` | :arrow_down: |
| [base/utils/log.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9sb2cuZ28=) | `75.00% <0.00%> (-18.11%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `78.88% <87.50%> (+0.13%)` | :arrow_up: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `79.87% <100.00%> (+0.62%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `93.84% <100.00%> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `71.79% <100.00%> (-0.71%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `73.33% <100.00%> (-0.58%)` | :arrow_down: |
| [manager/controllers/baseline\_systems\_remove.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX3JlbW92ZS5nbw==) | `72.72% <100.00%> (+21.00%)` | :arrow_up: |
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `90.69% <100.00%> (ø)` | |
| ... and [11 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2651d06...2c9980c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1022?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 20, 2022 at 07:56 PM UTC

a few more things:
- rebase it, please. There has been a change in `baseline_create_test` in the meantime.
- squash the commits a bit, the commits below could be a single commit. To avoid these `Merge commits` use `git rebase master` instead of merging master into your branch.
<img width="500" alt="image" src="https://user-images.githubusercontent.com/14837841/180070757-4808674d-8cba-4241-b000-fe95ae679c72.png">
- don't forget to prefix commits with `SPM-XXXX` if it's associate with some Jira task. This way the commit will be linked to Jira card. You don't need to do it for the commit which fixes the readme though.



---

## Reviews

### Review by @psegedy - Changes Requested on July 20, 2022 at 07:10 PM UTC

It looks great, thank you for making tests more readable and maintainable. I just left some comments to make it a bit clearer.
Sorry it took so long to review it

### Review by @Mischulee - Commented on July 27, 2022 at 11:48 AM UTC

### Review by @Mischulee - Commented on July 27, 2022 at 11:49 AM UTC

### Review by @Mischulee - Commented on July 27, 2022 at 11:49 AM UTC

### Review by @Mischulee - Commented on July 27, 2022 at 11:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1022*
