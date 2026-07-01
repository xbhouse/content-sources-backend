---
type: pull_request
number: 685
title: "Swap RBAC and account check order"
state: merged
author: semtexzv
created: 2021-05-25T08:02:55Z
updated: 2021-05-27T03:42:18Z
closed: 2021-05-27T03:42:18Z
merged: 2021-05-27T03:42:18Z
base_branch: master
head_branch: check-swap
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/685
---

# Pull Request #685: Swap RBAC and account check order

**Author**: @semtexzv
**Created**: May 25, 2021 at 08:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `check-swap`

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

### Comment by @MichaelMraka on May 26, 2021 at 08:19 AM UTC

What does this fixes? Does it  belong to any SPM task?

### Comment by @semtexzv on May 26, 2021 at 09:36 AM UTC

When we have an empty account, and remove the RBAC permissions from this account, we keep showing the "how to add your systems" page instead of "you can't see this" page

---

## Reviews

### Review by @josef-hak - Approved on May 27, 2021 at 03:42 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/685*
