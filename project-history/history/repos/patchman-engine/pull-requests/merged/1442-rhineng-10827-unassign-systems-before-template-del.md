---
type: pull_request
number: 1442
title: "RHINENG-10827: unassign systems before template delete"
state: merged
author: MichaelMraka
created: 2024-06-20T10:04:57Z
updated: 2024-06-24T12:46:26Z
closed: 2024-06-24T12:46:25Z
merged: 2024-06-24T12:46:25Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1442
---

# Pull Request #1442: RHINENG-10827: unassign systems before template delete

**Author**: @MichaelMraka
**Created**: June 20, 2024 at 10:04 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @jira-linking on June 20, 2024 at 10:05 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10827


### Comment by @MichaelMraka on June 21, 2024 at 02:35 PM UTC

> it seems like error is always nil in this function if I didn't miss something so we might not need to return error at all, otherwise LGTM

Well, let's try once more... let `TemplateDelete()` return errors and put error handling into `TemplatesMessageHandler()`. This seems to be cleaner.

---

## Reviews

### Review by @psegedy - Commented on June 21, 2024 at 10:55 AM UTC

### Review by @psegedy - Approved on June 21, 2024 at 10:59 AM UTC

it seems like error is always nil in this function if I didn't miss something so we might not need to return error at all, otherwise LGTM

### Review by @MichaelMraka - Commented on June 21, 2024 at 02:32 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1442*
