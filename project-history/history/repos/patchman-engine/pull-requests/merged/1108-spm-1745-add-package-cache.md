---
type: pull_request
number: 1108
title: "SPM-1745: add package cache"
state: merged
author: psegedy
created: 2022-09-19T12:24:05Z
updated: 2022-09-21T12:52:37Z
closed: 2022-09-21T12:52:37Z
merged: 2022-09-21T12:52:37Z
base_branch: master
head_branch: cache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1108
---

# Pull Request #1108: SPM-1745: add package cache

**Author**: @psegedy
**Created**: September 19, 2022 at 12:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `cache`

## Description

- use cached counts in `/packages` api, enable/disable it by env var
- cache is in `package_account_data` table
- cache validity for particular account is stored in `rh_account` table
  - invalidated after evaluation
- `refresh_packages_caches` procedure can refresh cache of a single account or for all accounts which have invalid cache
- refresh is triggered every 5 min by a job,  no trigger is used to run cache refresh
- there is also an admin api for cache refresh 

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

### Comment by @jira-linking on September 19, 2022 at 12:24 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1745


---

## Reviews

### Review by @MichaelMraka - Approved on September 21, 2022 at 09:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1108*
