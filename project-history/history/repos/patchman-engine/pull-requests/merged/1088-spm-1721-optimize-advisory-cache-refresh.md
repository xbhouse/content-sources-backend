---
type: pull_request
number: 1088
title: "SPM-1721: optimize advisory cache refresh"
state: merged
author: psegedy
created: 2022-08-31T12:54:27Z
updated: 2022-09-05T10:49:09Z
closed: 2022-09-05T10:49:09Z
merged: 2022-09-05T10:49:09Z
base_branch: master
head_branch: refresh_job
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1088
---

# Pull Request #1088: SPM-1721: optimize advisory cache refresh

**Author**: @psegedy
**Created**: August 31, 2022 at 12:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `refresh_job`

## Description

~~improvements:~~

~~- set max number of goroutines for refresh job, currently it runs as many as possible goroutines (for 80k accounts in prod) which results in locking all rows for accounts at the same moment~~
~~- query accounts which has inconsistencies in system counts and only refresh these accounts~~

~~the problem can be in the query which finds accounts, it works quite fast in stage but it timeouts in gabi in prod and it has quite high cost~~

~~Subquery Scan on sums  (cost=3722126.53..3723369.97 rows=199 width=4)"~~

~~any ideas how to improve it further?~~

After discussion with @MichaelMraka we've decided to refresh all caches at once by not supplying any account ids
```
refresh_advisory_caches(NULL, NULL)
```
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

## Reviews

### Review by @MichaelMraka - Approved on September 02, 2022 at 08:51 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1088*
