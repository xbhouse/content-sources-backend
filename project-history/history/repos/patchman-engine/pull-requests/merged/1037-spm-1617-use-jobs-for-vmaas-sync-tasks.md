---
type: pull_request
number: 1037
title: "SPM-1617: use jobs for vmaas sync tasks"
state: merged
author: psegedy
created: 2022-07-26T11:23:45Z
updated: 2022-08-10T09:55:48Z
closed: 2022-08-10T09:55:48Z
merged: 2022-08-10T09:55:48Z
base_branch: master
head_branch: jobs
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1037
---

# Pull Request #1037: SPM-1617: use jobs for vmaas sync tasks

**Author**: @psegedy
**Created**: July 26, 2022 at 11:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `jobs`

## Description

- create jobs for each task which used to be run in vmaas_sync container
- run adminApi in manager container on privatePort
- move timestamp types to separate package app/base/types to avoid cyclic imports
- drop websocket connection, use vmaas `/dbchange` endpoint instead to check if vmaas has any new data compared to our timestamp in DB in `isSyncNeeded`, run the job every 5 minutes

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

### Comment by @jira-linking on July 26, 2022 at 11:23 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1617


### Comment by @codecov-commenter on August 09, 2022 at 08:38 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1037](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4c39aa9) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **increase** coverage by `0.23%`.
> The diff coverage is `41.10%`.

```diff
@@            Coverage Diff             @@
##           master    #1037      +/-   ##
==========================================
+ Coverage   61.25%   61.48%   +0.23%     
==========================================
  Files          95       98       +3     
  Lines        5448     5424      -24     
==========================================
- Hits         3337     3335       -2     
+ Misses       1682     1662      -20     
+ Partials      429      427       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.48% <41.10%> (+0.23%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `8.69% <0.00%> (+0.36%)` | :arrow_up: |
| [tasks/caches/refresh\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY2FjaGVzL3JlZnJlc2hfcGFja2FnZXMuZ28=) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.40% <0.00%> (ø)` | |
| [tasks/vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `61.90% <ø> (ø)` | |
| [tasks/vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `59.25% <ø> (ø)` | |
| [tasks/vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `62.00% <ø> (ø)` | |
| [tasks/vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `70.31% <ø> (ø)` | |
| ... and [19 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1037/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


---

## Reviews

### Review by @MichaelMraka - Approved on August 10, 2022 at 08:23 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1037*
