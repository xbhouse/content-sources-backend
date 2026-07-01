---
type: pull_request
number: 991
title: "Send only 1 payload status for skipped messages"
state: merged
author: psegedy
created: 2022-06-27T14:44:12Z
updated: 2022-06-28T12:36:35Z
closed: 2022-06-28T12:36:35Z
merged: 2022-06-28T12:36:35Z
base_branch: master
head_branch: send_payload_fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/991
---

# Pull Request #991: Send only 1 payload status for skipped messages

**Author**: @psegedy
**Created**: June 27, 2022 at 02:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `send_payload_fix`

## Description

We don;t need to send 2 messages - `received` and `error` when the message will not be evaluated (e.g. coming from yupana). Instead we will send only `received` msg with `StatusMsg` explaining the message was skipped.

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

### Review by @michalslomczynski - Approved on June 28, 2022 at 11:45 AM UTC

### Review by @MichaelMraka - Approved on June 28, 2022 at 11:46 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/991*
