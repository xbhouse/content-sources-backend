---
type: pull_request
number: 1406
title: "RHINENG-9238: check duplicate packageID before insert"
state: merged
author: psegedy
created: 2024-04-05T09:59:21Z
updated: 2024-04-08T14:47:46Z
closed: 2024-04-08T14:47:46Z
merged: 2024-04-08T14:47:46Z
base_branch: master
head_branch: duplicate_package_id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1406
---

# Pull Request #1406: RHINENG-9238: check duplicate packageID before insert

**Author**: @psegedy
**Created**: April 05, 2024 at 09:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `duplicate_package_id`

## Description

Storing system packages: ERROR: ON CONFLICT DO UPDATE command cannot affect row a second time

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

### Comment by @jira-linking on April 05, 2024 at 09:59 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9238
https://issues.redhat.com/browse/RHINENG-9077


### Comment by @psegedy on April 08, 2024 at 09:51 AM UTC

/retest


---

## Reviews

### Review by @MichaelMraka - Approved on April 08, 2024 at 11:57 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1406*
