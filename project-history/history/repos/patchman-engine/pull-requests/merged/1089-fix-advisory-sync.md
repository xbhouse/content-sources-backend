---
type: pull_request
number: 1089
title: "Fix advisory sync"
state: merged
author: psegedy
created: 2022-08-31T12:59:44Z
updated: 2022-09-01T13:35:53Z
closed: 2022-09-01T13:35:53Z
merged: 2022-09-01T13:35:53Z
base_branch: master
head_branch: fix_advisory_sync
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1089
---

# Pull Request #1089: Fix advisory sync

**Author**: @psegedy
**Created**: August 31, 2022 at 12:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_advisory_sync`

## Description

cannot send slice to gorm `Updates`
fixes
```
{"@timestamp":"2022-08-31T12:57:08.9Z","err":"unsupported data","level":"error","message":"couldn't update advisory_metadata"}
```
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

### Review by @MichaelMraka - Approved on September 01, 2022 at 01:34 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1089*
