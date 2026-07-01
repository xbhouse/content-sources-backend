---
type: pull_request
number: 1120
title: "SPM-1745: refresh package cache in transactions"
state: merged
author: psegedy
created: 2022-10-06T16:12:59Z
updated: 2022-10-12T14:06:17Z
closed: 2022-10-12T14:06:17Z
merged: 2022-10-12T14:06:17Z
base_branch: master
head_branch: pkgcache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1120
---

# Pull Request #1120: SPM-1745: refresh package cache in transactions

**Author**: @psegedy
**Created**: October 06, 2022 at 04:12 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkgcache`

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

### Comment by @jira-linking on October 06, 2022 at 04:13 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1745


### Comment by @psegedy on October 10, 2022 at 11:54 AM UTC

@MichaelMraka do we need to do also something like SELECT rh_account_id FROM rh_account FOR UPDATE?

### Comment by @MichaelMraka on October 11, 2022 at 11:37 AM UTC

We don't need SELECT ... FOR UPDATE.

---

## Reviews

### Review by @MichaelMraka - Commented on October 11, 2022 at 09:31 AM UTC

### Review by @psegedy - Commented on October 11, 2022 at 01:32 PM UTC

### Review by @psegedy - Commented on October 11, 2022 at 01:33 PM UTC

### Review by @MichaelMraka - Approved on October 12, 2022 at 12:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1120*
