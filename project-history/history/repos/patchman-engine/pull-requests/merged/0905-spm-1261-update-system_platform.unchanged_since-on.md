---
type: pull_request
number: 905
title: "SPM-1261: Update \"system_platform.unchanged_since\" on baseline update"
state: merged
author: josef-hak
created: 2022-02-09T21:53:31Z
updated: 2022-02-10T14:33:47Z
closed: 2022-02-10T14:33:47Z
merged: 2022-02-10T14:33:46Z
base_branch: master
head_branch: baseline-update
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/905
---

# Pull Request #905: SPM-1261: Update "system_platform.unchanged_since" on baseline update

**Author**: @josef-hak
**Created**: February 09, 2022 at 09:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `baseline-update`

## Description

- needed not to skip system evaluation when re-uploaded after baseline assignment
even if vmaas_json (json_checksum) is unchanged

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

### Review by @MichaelMraka - Approved on February 10, 2022 at 12:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/905*
