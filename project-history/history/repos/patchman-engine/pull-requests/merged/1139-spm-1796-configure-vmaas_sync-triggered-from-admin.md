---
type: pull_request
number: 1139
title: "SPM-1796: configure vmaas_sync triggered from admin api"
state: merged
author: MichaelMraka
created: 2022-11-02T10:34:24Z
updated: 2022-11-03T10:52:35Z
closed: 2022-11-03T10:52:35Z
merged: 2022-11-03T10:52:34Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1139
---

# Pull Request #1139: SPM-1796: configure vmaas_sync triggered from admin api

**Author**: @MichaelMraka
**Created**: November 02, 2022 at 10:34 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

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

## Reviews

### Review by @psegedy - Changes Requested on November 02, 2022 at 12:03 PM UTC

please change `configure` -> `Configure` in other sync tasks, tests are failing with
```
tasks/vmaas_sync/advisory_sync_test.go:170:2: undeclared name: `configure` (typecheck)
test         | 	configure()
test         | 	^
test         | tasks/vmaas_sync/advisory_sync_test.go:189:2: undeclared name: `configure` (typecheck)
test         | 	configure()
test         | 	^
test         | tasks/vmaas_sync/advisory_sync_test.go:217:2: undeclared name: `configure` (typecheck)
test         | 	configure()
test         | 	^
```

### Review by @psegedy - Approved on November 02, 2022 at 01:03 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1139*
