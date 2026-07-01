---
type: pull_request
number: 676
title: "SPM-674 Add Clowder support 1 (db-admin, manager)"
state: merged
author: josef-hak
created: 2021-05-17T14:53:48Z
updated: 2021-05-26T11:14:10Z
closed: 2021-05-26T11:01:15Z
merged: 2021-05-26T11:01:15Z
base_branch: master
head_branch: clowder
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/676
---

# Pull Request #676: SPM-674 Add Clowder support 1 (db-admin, manager)

**Author**: @josef-hak
**Created**: May 17, 2021 at 02:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `clowder`

## Description

- Basic Clowder support works.
- Clowder params are exported to environment vars using `try_export_clowder_params.sh`.
- Two db-admin vars had to be renamed (DB_USER -> DB_ADMIN_USER, DB_PASSWD -> DB_ADMIN_PASSWD), to be consistent with Clowder params.
- Added `deploy/clowdapp.yml` template to deploy Patchman using Clowder to Ephemeral environemnt - used in new `ci.ext.devshift.net` PR pipeline. Only database-admin and manager for now.

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

### Comment by @codecov-commenter on May 18, 2021 at 09:06 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#676](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7277dfa) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ca738411d7eb2a9a0a7c0ff210209ac8c74ce822?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ca73841) will **decrease** coverage by `0.38%`.
> The diff coverage is `13.79%`.

> :exclamation: Current head 7277dfa differs from pull request most recent head a0b487f. Consider uploading reports for the commit a0b487f to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #676      +/-   ##
==========================================
- Coverage   58.50%   58.12%   -0.39%     
==========================================
  Files          72       73       +1     
  Lines        3379     3403      +24     
==========================================
+ Hits         1977     1978       +1     
- Misses       1118     1141      +23     
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.12% <13.79%> (-0.39%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/core/probes.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL3Byb2Jlcy5nbw==) | `46.66% <0.00%> (-11.67%)` | :arrow_down: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.69% <0.00%> (-0.31%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.16% <100.00%> (-0.18%)` | :arrow_down: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `72.72% <100.00%> (+6.06%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ca73841...a0b487f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/676?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on May 26, 2021 at 07:45 AM UTC

### Review by @MichaelMraka - Commented on May 26, 2021 at 08:04 AM UTC

### Review by @josef-hak - Commented on May 26, 2021 at 08:39 AM UTC

### Review by @MichaelMraka - Commented on May 26, 2021 at 08:48 AM UTC

### Review by @MichaelMraka - Approved on May 26, 2021 at 10:55 AM UTC

### Review by @semtexzv - Approved on May 26, 2021 at 11:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/676*
