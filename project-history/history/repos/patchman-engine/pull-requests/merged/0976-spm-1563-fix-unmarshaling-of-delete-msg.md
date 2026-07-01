---
type: pull_request
number: 976
title: "SPM-1563: fix unmarshaling of delete msg"
state: merged
author: psegedy
created: 2022-06-10T15:37:02Z
updated: 2022-06-13T08:51:43Z
closed: 2022-06-13T08:51:42Z
merged: 2022-06-13T08:51:42Z
base_branch: master
head_branch: fix_delete_msg
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/976
---

# Pull Request #976: SPM-1563: fix unmarshaling of delete msg

**Author**: @psegedy
**Created**: June 10, 2022 at 03:37 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_delete_msg`

## Description

fixes
```
time="2022-06-10T14:47:36Z" level=error msg="Invalid 'delete' message format" fields.msg="{\"insights_id\": \"4dbf09d0-848b-47f1-b60e-420cf0632c18\", \"metadata\": {\"request_id\": \"-1\"}, \"type\": \"delete\", \"account\": \"0369233\", \"timestamp\": \"2022-06-10T14:47:36.984126+00:00\", \"id\": \"13668aff-cc47-4120-bb1b-5e577813ffed\", \"org_id\": \"3340851\", \"request_id\": \"-1\"}" inventoryID=13668aff-cc47-4120-bb1b-5e577813ffed
```
`RequestIDs` have to be marshalled/unmarshalled as `request_ids`, not `request_id`

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

### Review by @MichaelMraka - Approved on June 13, 2022 at 08:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/976*
