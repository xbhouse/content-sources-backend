---
type: pull_request
number: 1145
title: "SPM-1807: run package refresh in 1 thread"
state: merged
author: psegedy
created: 2022-11-22T11:46:49Z
updated: 2022-11-22T16:51:51Z
closed: 2022-11-22T16:51:51Z
merged: 2022-11-22T16:51:51Z
base_branch: master
head_branch: jobs_procs
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1145
---

# Pull Request #1145: SPM-1807: run package refresh in 1 thread

**Author**: @psegedy
**Created**: November 22, 2022 at 11:46 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `jobs_procs`

## Description

try to run package refresh job in 1 golang thread to see if it is not causing any issues

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

### Comment by @jira-linking on November 22, 2022 at 11:46 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1807


---

## Reviews

### Review by @MichaelMraka - Approved on November 22, 2022 at 12:02 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1145*
