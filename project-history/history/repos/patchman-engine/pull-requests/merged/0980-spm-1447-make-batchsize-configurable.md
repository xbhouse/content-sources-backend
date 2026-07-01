---
type: pull_request
number: 980
title: "SPM-1447: Make BatchSize configurable"
state: merged
author: psegedy
created: 2022-06-17T14:59:19Z
updated: 2022-06-22T14:29:32Z
closed: 2022-06-22T14:29:31Z
merged: 2022-06-22T14:29:31Z
base_branch: master
head_branch: configure_batch_size
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/980
---

# Pull Request #980: SPM-1447: Make BatchSize configurable

**Author**: @psegedy
**Created**: June 17, 2022 at 02:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `configure_batch_size`

## Description

**NOTE**: we need to be careful when changing `BatchSize`. I experienced some OOMKilled errors in ephemeral. Probably due to usage of 8 consumers (default in clowdapp.yaml) in evaluator-recalc. 

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

### Review by @MichaelMraka - Approved on June 22, 2022 at 07:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/980*
