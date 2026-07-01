---
type: pull_request
number: 1351
title: "RHINENG-5394: system_package migration improvements"
state: merged
author: psegedy
created: 2023-12-07T10:25:43Z
updated: 2023-12-07T10:35:03Z
closed: 2023-12-07T10:35:03Z
merged: 2023-12-07T10:35:03Z
base_branch: master
head_branch: migration3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1351
---

# Pull Request #1351: RHINENG-5394: system_package migration improvements

**Author**: @psegedy
**Created**: December 07, 2023 at 10:25 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `migration3`

## Description

- increase batch size to 5000
- log every millionth insert
- truncate system_package2 table when migration fails
- don't fail when removal of constraints fails

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

### Comment by @jira-linking on December 07, 2023 at 10:25 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-5394


---

## Reviews

### Review by @MichaelMraka - Approved on December 07, 2023 at 10:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1351*
