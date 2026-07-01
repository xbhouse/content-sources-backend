---
type: pull_request
number: 1011
title: "Hotfix advisory detail cache"
state: closed
author: MichaelMraka
created: 2022-07-08T09:43:47Z
updated: 2022-07-08T14:46:17Z
closed: 2022-07-08T14:46:16Z
base_branch: master
head_branch: hotfix-advisory-detail-cache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1011
---

# Pull Request #1011: Hotfix advisory detail cache

**Author**: @MichaelMraka
**Created**: July 08, 2022 at 09:43 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `hotfix-advisory-detail-cache`

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

### Comment by @jira-linking on July 08, 2022 at 09:43 AM UTC

Commits missing Jira IDs:
98f19c210bba593971e82098f83e3fc3774f2ae7
6a1d5d464c64e3392267c6d99cc8a0f633e5e272
eb5675dbb9582b6352dd9750ac7a8c3863f158ac
06ec6fc984b45e499b933b07cee9126a08728591
e6383e2a614550767647ce694177da0f63e7b843
32476b0939aecf06c3313b27881e5d3b4f366cd4
3ac741de829f69e9b408d9dec16fd8e861b6bb8c
8a4acd1a8d4d713b44907279403b30895c157a09
c70ee7a2a696aaea6f051e432fac83af97f360e9
3d40fc507353905e74ed059c2e0e70e8fd8b3a0d
9d1775e10a971f75308d90b1619773ac0156c7b1
2fbcd0bfabafd909b4b7b65a43902be854b10889
5e88f461b5b807e304b07b480868adce8030ee67
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1619
https://issues.redhat.com/browse/SPM-1610
https://issues.redhat.com/browse/SPM-1484
https://issues.redhat.com/browse/SPM-1609
https://issues.redhat.com/browse/SPM-1600
https://issues.redhat.com/browse/SPM-1439
https://issues.redhat.com/browse/SPM-1501
https://issues.redhat.com/browse/SPM-1550


### Comment by @codecov-commenter on July 08, 2022 at 09:50 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1011](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c0c2da5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf619f46f8c3ad992a623e23e98cb50c81230a93?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf619f4) will **decrease** coverage by `0.34%`.
> The diff coverage is `55.85%`.

```diff
@@            Coverage Diff             @@
##           master    #1011      +/-   ##
==========================================
- Coverage   60.97%   60.63%   -0.35%     
==========================================
  Files          95       94       -1     
  Lines        5374     5213     -161     
==========================================
- Hits         3277     3161     -116     
+ Misses       1666     1643      -23     
+ Partials      431      409      -22     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.63% <55.85%> (-0.35%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.68% <ø> (+0.35%)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <ø> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `65.00% <ø> (-0.44%)` | :arrow_down: |
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <ø> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `41.66% <ø> (-10.84%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/refresh.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoLmdv) | `0.00% <0.00%> (-30.77%)` | :arrow_down: |
| [vmaas\_sync/refresh\_advisory\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoX2Fkdmlzb3J5X2NhY2hlcy5nbw==) | `38.88% <ø> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `43.20% <20.00%> (-1.52%)` | :arrow_down: |
| ... and [11 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [90d5e89...c0c2da5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1011?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 08, 2022 at 02:46 PM UTC

released to prod, closing

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1011*
