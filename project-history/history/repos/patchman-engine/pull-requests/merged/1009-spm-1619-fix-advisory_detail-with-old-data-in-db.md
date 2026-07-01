---
type: pull_request
number: 1009
title: "SPM-1619: fix advisory_detail with old data in db"
state: merged
author: psegedy
created: 2022-07-07T16:20:38Z
updated: 2022-07-08T09:09:50Z
closed: 2022-07-08T09:09:49Z
merged: 2022-07-08T09:09:49Z
base_branch: master
head_branch: fix-spm-1619
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1009
---

# Pull Request #1009: SPM-1619: fix advisory_detail with old data in db

**Author**: @psegedy
**Created**: July 07, 2022 at 04:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-spm-1619`

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

### Comment by @jira-linking on July 07, 2022 at 04:20 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1619


### Comment by @MichaelMraka on July 08, 2022 at 08:20 AM UTC

I've added some cleanup. Please review it.

### Comment by @psegedy on July 08, 2022 at 08:52 AM UTC

looks way better with `parseJSONList`

---

## Reviews

### Review by @MichaelMraka - Approved on July 08, 2022 at 08:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1009*
