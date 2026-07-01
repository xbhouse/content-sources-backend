---
type: pull_request
number: 1166
title: "SPM-1868: fix baselines 500 with empty description or with whitespaces"
state: merged
author: psegedy
created: 2023-01-30T16:19:57Z
updated: 2023-01-31T08:45:21Z
closed: 2023-01-31T08:45:21Z
merged: 2023-01-31T08:45:20Z
base_branch: master
head_branch: baseline_udpate_500
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1166
---

# Pull Request #1166: SPM-1868: fix baselines 500 with empty description or with whitespaces

**Author**: @psegedy
**Created**: January 30, 2023 at 04:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline_udpate_500`

## Description

- empty description should be stored as NULL in DB
- 400 for strings containing only whitespaces 
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

### Comment by @jira-linking on January 30, 2023 at 04:19 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1868


---

## Reviews

### Review by @MichaelMraka - Approved on January 31, 2023 at 08:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1166*
