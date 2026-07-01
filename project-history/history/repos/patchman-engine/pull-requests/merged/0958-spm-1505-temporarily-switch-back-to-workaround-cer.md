---
type: pull_request
number: 958
title: "SPM-1505: temporarily switch back to workaround cert issue in stage"
state: merged
author: MichaelMraka
created: 2022-05-13T11:58:53Z
updated: 2022-05-13T13:12:51Z
closed: 2022-05-13T13:12:51Z
merged: 2022-05-13T13:12:51Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/958
---

# Pull Request #958: SPM-1505: temporarily switch back to workaround cert issue in stage

**Author**: @MichaelMraka
**Created**: May 13, 2022 at 11:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

## Description

RequestError: send request failed\ncaused by: Post \"https://.../\": x509: certificate signed by unknown authority: unable to setup CloudWatch logging
failed to initialize database, got error failed to connect to `host=patchman-stage...`: failed to write startup message (x509: certificate relies on legacy Common Name field, use SANs instead

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

### Review by @psegedy - Approved on May 13, 2022 at 01:12 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/958*
