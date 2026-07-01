---
type: pull_request
number: 696
title: "SPM-964: fix Kafka certificate issue adding var to test.env"
state: merged
author: josef-hak
created: 2021-06-07T19:59:41Z
updated: 2021-08-26T18:41:37Z
closed: 2021-06-08T07:25:19Z
merged: 2021-06-08T07:25:19Z
base_branch: master
head_branch: fix-cert
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/696
---

# Pull Request #696: SPM-964: fix Kafka certificate issue adding var to test.env

**Author**: @josef-hak
**Created**: June 07, 2021 at 07:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-cert`

## Description

- GODEBUG=x509ignoreCN=0
- error message: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
- added script to generate testing kafka certificates

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

### Review by @MichaelMraka - Approved on June 08, 2021 at 07:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/696*
