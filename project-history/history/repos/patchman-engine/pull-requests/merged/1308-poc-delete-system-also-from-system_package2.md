---
type: pull_request
number: 1308
title: "POC: delete system also from system_package2"
state: merged
author: psegedy
created: 2023-09-12T10:24:01Z
updated: 2023-09-12T11:10:41Z
closed: 2023-09-12T11:10:41Z
merged: 2023-09-12T11:10:41Z
base_branch: poc
head_branch: delete_system
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1308
---

# Pull Request #1308: POC: delete system also from system_package2

**Author**: @psegedy
**Created**: September 12, 2023 at 10:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `poc` ← **Head**: `delete_system`

## Description

error when deleting system in listener
```
time="2023-09-12T10:14:30Z" level=error msg="Try failed" attempt=3 err="Could not opt_out system: ERROR: update or delete on table \"system_platform_8\" violates foreign key constraint \"system_package2_rh_account_id_system_id_fkey9\" on table \"system_package2\" (SQLSTATE 23503)"
time="2023-09-12T10:14:31Z" level=error msg="Could not delete system" err="ERROR: update or delete on table \"system_platform_8\" violates foreign key constraint \"system_package2_rh_account_id_system_id_fkey9\" on table \"system_package2\" (SQLSTATE 23503)" inventoryID=803eea15-79a5-4b9f-863b-631409911ed0
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

## Discussion

### Comment by @jira-linking on September 12, 2023 at 10:24 AM UTC

Commits missing Jira IDs:
b077f58d3d0faa8e57c9a00343cef64d2f92814d


---

## Reviews

### Review by @MichaelMraka - Approved on September 12, 2023 at 11:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1308*
