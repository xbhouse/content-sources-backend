---
type: pull_request
number: 1195
title: "SPM-1962: fix 500 for filter"
state: merged
author: psegedy
created: 2023-03-17T15:01:57Z
updated: 2023-03-17T16:27:29Z
closed: 2023-03-17T16:27:29Z
merged: 2023-03-17T16:27:29Z
base_branch: master
head_branch: 500_sap_sids
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1195
---

# Pull Request #1195: SPM-1962: fix 500 for filter

**Author**: @psegedy
**Created**: March 17, 2023 at 03:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `500_sap_sids`

## Description

such query returned 500 `http://localhost:8080/api/patch/v3/systems?filter%5Bstale%5D=True&filter%5Bsystem_profile%5D%5Bsap_sids%5D%5Bin%5D=insights-client%2Fdemo%3D%22%7B%7B%20slug%20%7D%7D%22`

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

### Review by @MichaelMraka - Approved on March 17, 2023 at 03:42 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1195*
