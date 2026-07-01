---
type: pull_request
number: 727
title: "SPM-674 respect vmaas resource settings in ephemeral"
state: merged
author: psegedy
created: 2021-07-09T13:46:07Z
updated: 2021-07-09T14:28:47Z
closed: 2021-07-09T14:28:47Z
merged: 2021-07-09T14:28:47Z
base_branch: master
head_branch: patch-1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/727
---

# Pull Request #727: SPM-674 respect vmaas resource settings in ephemeral

**Author**: @psegedy
**Created**: July 09, 2021 at 01:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `patch-1`

## Description

vmaas-reposcan gets OOMKilled during sync if it does not have enough resources, default resource limit in ephemeral is 256Mb which is not enough, so we need to respect settings from vmaas clowdapp.yaml


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

### Review by @josef-hak - Approved on July 09, 2021 at 02:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/727*
