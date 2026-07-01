---
type: pull_request
number: 675
title: "SPM-910: build image just once"
state: merged
author: MichaelMraka
created: 2021-05-17T10:44:59Z
updated: 2021-05-17T14:33:28Z
closed: 2021-05-17T14:33:28Z
merged: 2021-05-17T14:33:28Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/675
---

# Pull Request #675: SPM-910: build image just once

**Author**: @MichaelMraka
**Created**: May 17, 2021 at 10:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

improved speed of podman-compose up --build

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist
speeds up podman-compose build from 30s to 5s

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

### Review by @josef-hak - Approved on May 17, 2021 at 02:05 PM UTC

Thanks, works for me too, using docker.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/675*
