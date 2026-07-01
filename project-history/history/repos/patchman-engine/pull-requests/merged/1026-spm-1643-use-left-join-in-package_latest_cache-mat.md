---
type: pull_request
number: 1026
title: "SPM-1643: use left join in package_latest_cache mat view"
state: merged
author: psegedy
created: 2022-07-15T13:21:35Z
updated: 2022-07-15T14:04:35Z
closed: 2022-07-15T14:04:35Z
merged: 2022-07-15T14:04:35Z
base_branch: master
head_branch: mat_view_join
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1026
---

# Pull Request #1026: SPM-1643: use left join in package_latest_cache mat view

**Author**: @psegedy
**Created**: July 15, 2022 at 01:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `mat_view_join`

## Description

needed to solve empty string vs null issue
basically the same kind of migration as 066
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

### Comment by @jira-linking on July 15, 2022 at 01:21 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1643


---

## Reviews

### Review by @MichaelMraka - Approved on July 15, 2022 at 02:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1026*
