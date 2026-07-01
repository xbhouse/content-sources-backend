---
type: pull_request
number: 1028
title: "Hotfix advisory detail cache + notifications orgID + force schema migration"
state: closed
author: michalslomczynski
created: 2022-07-19T10:47:37Z
updated: 2022-07-20T09:33:49Z
closed: 2022-07-20T09:33:46Z
base_branch: master
head_branch: hotfix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1028
---

# Pull Request #1028: Hotfix advisory detail cache + notifications orgID + force schema migration

**Author**: @michalslomczynski
**Created**: July 19, 2022 at 10:47 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `hotfix`

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

### Comment by @jira-linking on July 19, 2022 at 10:47 AM UTC

Commits missing Jira IDs:
179a160d371fa1fd300b6c82df1a4653eff5d953
1cd8a591a8725c4cb106144aad9197ddce00b3c1
53d738c3f5fbac8b75be0b0370766f8596d9afcf
38101170dc185f39dcc925e8fd9a5b3c40d5e90b
d65dd929c7633a6fa308fd815b4171e572c42b0f
02553483e4ac5ce2f5d91ef8581508fa0398f878
cdb31f072903a87f66734cfdb4ac5f8c69a83c2e
01301920234a451674d7e700ffd1720c61c573a9
c370f1c5ed6dd4b468fb344c8a337611c0c94d99
41ccb90af81f6bcd6f7dff3e52e4cc8efcdaafa3
96ae2a16be41747a54bb7aab52be610c557a7545
462fbb3ef7726a6cd246982e134e2fb0d808cae4
0d1842795c9c2c58aacdcaebe8711b8f1df1df6a
8025d3bae174abde044166c979d0fc38a80bdbb3
a11a6d98b70f3e8896c93e9a0fd708bd47f2841b
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1643
https://issues.redhat.com/browse/SPM-1630
https://issues.redhat.com/browse/SPM-1616
https://issues.redhat.com/browse/SPM-1636
https://issues.redhat.com/browse/SPM-1633
https://issues.redhat.com/browse/SPM-1626
https://issues.redhat.com/browse/SPM-1619
https://issues.redhat.com/browse/SPM-1576
https://issues.redhat.com/browse/SPM-1622


### Comment by @codecov-commenter on July 19, 2022 at 10:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1028](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (db2bd62) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fd466045239b9c4ed8afae2e6294685d32eee94a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fd46604) will **decrease** coverage by `0.69%`.
> The diff coverage is `65.11%`.

```diff
@@            Coverage Diff             @@
##           master    #1028      +/-   ##
==========================================
- Coverage   61.25%   60.56%   -0.70%     
==========================================
  Files          95       94       -1     
  Lines        5413     5219     -194     
==========================================
- Hits         3316     3161     -155     
+ Misses       1673     1646      -27     
+ Partials      424      412      -12     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.56% <65.11%> (-0.70%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.68% <0.00%> (+0.47%)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <ø> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `65.00% <ø> (-0.44%)` | :arrow_down: |
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `9.09% <ø> (+0.39%)` | :arrow_up: |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `27.69% <ø> (+2.34%)` | :arrow_up: |
| [base/utils/log.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9sb2cuZ28=) | `93.10% <ø> (+18.10%)` | :arrow_up: |
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <ø> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `41.66% <ø> (-10.84%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| ... and [31 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b6ea455...db2bd62](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1028?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1028*
