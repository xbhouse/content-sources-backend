---
type: pull_request
number: 1437
title: "RHINENG-10747: limit=-1 is not allowed"
state: merged
author: psegedy
created: 2024-06-17T10:13:15Z
updated: 2024-06-17T11:19:19Z
closed: 2024-06-17T11:17:42Z
merged: 2024-06-17T11:17:42Z
base_branch: master
head_branch: limitminusone
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1437
---

# Pull Request #1437: RHINENG-10747: limit=-1 is not allowed

**Author**: @psegedy
**Created**: June 17, 2024 at 10:13 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `limitminusone`

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

### Comment by @jira-linking on June 17, 2024 at 10:13 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10747


### Comment by @MichaelMraka on June 17, 2024 at 11:14 AM UTC

I'd also change
`base/utils/gin.go:76:           return errors.New("limit must not be less than 1, or should be -1 to return all items")`

---

## Reviews

### Review by @MichaelMraka - Approved on June 17, 2024 at 11:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1437*
