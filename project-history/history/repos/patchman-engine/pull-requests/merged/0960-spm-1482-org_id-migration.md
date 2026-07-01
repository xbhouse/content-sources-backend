---
type: pull_request
number: 960
title: "SPM-1482: org_id migration"
state: merged
author: psegedy
created: 2022-05-13T16:44:29Z
updated: 2022-05-31T10:17:49Z
closed: 2022-05-31T10:17:49Z
merged: 2022-05-31T10:17:49Z
base_branch: master
head_branch: org_id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/960
---

# Pull Request #960: SPM-1482: org_id migration

**Author**: @psegedy
**Created**: May 13, 2022 at 04:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `org_id`

## Description

- add org_id column
- create new accounts using both account_number and org_id
- assign org_id to existing accounts
- handle case if account_number is deleted from model
- make sure we won't overwrite org_id if we have more records with null account_number and receive identity header with empty account_number
- create accounts using account_number from kafka messages because they are not sending org_id in the msg yet

~~todo:~~
~~- build AccountIDCache correctly if we receive only account_number (from kafka msg)~~
~~- make sure we won't leak data to wrong user. It can happen that we will have same ID in cache when some account_number of one user is the same as org_id of other user~~ should not happen, we are not using findAccount in listener

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

### Comment by @codecov-commenter on May 16, 2022 at 12:08 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#960](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b91b146) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4bc1c162bf6de549434a3563561b2036dbd1392d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4bc1c16) will **decrease** coverage by `0.02%`.
> The diff coverage is `61.76%`.

```diff
@@            Coverage Diff             @@
##           master     #960      +/-   ##
==========================================
- Coverage   61.20%   61.18%   -0.03%     
==========================================
  Files          91       91              
  Lines        4823     4846      +23     
==========================================
+ Hits         2952     2965      +13     
- Misses       1475     1482       +7     
- Partials      396      399       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.18% <61.76%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/identity.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9pZGVudGl0eS5nbw==) | `33.33% <20.00%> (-16.67%)` | :arrow_down: |
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `35.71% <65.38%> (+9.29%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.81% <100.00%> (+0.18%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [036bd63...b91b146](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/960?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on May 17, 2022 at 08:28 AM UTC

Please rebase it on top of current master.

---

## Reviews

### Review by @MichaelMraka - Commented on May 17, 2022 at 08:22 AM UTC

### Review by @psegedy - Commented on May 18, 2022 at 08:31 AM UTC

### Review by @psegedy - Commented on May 18, 2022 at 08:32 AM UTC

### Review by @MichaelMraka - Changes Requested on May 25, 2022 at 10:02 AM UTC

### Review by @MichaelMraka - Approved on May 26, 2022 at 02:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/960*
