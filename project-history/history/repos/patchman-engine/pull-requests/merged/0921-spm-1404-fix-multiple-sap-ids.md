---
type: pull_request
number: 921
title: "SPM-1404: Fix multiple SAP IDs"
state: merged
author: michalslomczynski
created: 2022-03-11T15:03:49Z
updated: 2022-03-15T09:11:23Z
closed: 2022-03-15T09:10:38Z
merged: 2022-03-15T09:10:38Z
base_branch: master
head_branch: fix-sap-sids
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/921
---

# Pull Request #921: SPM-1404: Fix multiple SAP IDs

**Author**: @michalslomczynski
**Created**: March 11, 2022 at 03:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-sap-sids`

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

### Comment by @michalslomczynski on March 11, 2022 at 03:14 PM UTC

After these changes, multiple SID ID do not overwrite itself, but they stack up and are passed as an array to the database instead.

---

## Reviews

### Review by @josef-hak - Changes Requested on March 14, 2022 at 06:26 PM UTC

### Review by @michalslomczynski - Commented on March 15, 2022 at 08:20 AM UTC

### Review by @josef-hak - Approved on March 15, 2022 at 09:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/921*
