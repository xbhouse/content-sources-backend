---
type: pull_request
number: 591
title: "SPM-830 split Evaluate method to multiple ones"
state: closed
author: josef-hak
created: 2021-04-01T14:51:32Z
updated: 2021-04-12T10:39:07Z
closed: 2021-04-07T16:48:59Z
base_branch: master
head_branch: installed-pkgs
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/591
---

# Pull Request #591: SPM-830 split Evaluate method to multiple ones

**Author**: @josef-hak
**Created**: April 01, 2021 at 02:51 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `installed-pkgs`

## Description

- split evaluate.go to evaluate_advisories.go and evaluate_packages.go
- split huge package evaluation methods into smaller ones

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

## Discussion

### Comment by @josef-hak on April 07, 2021 at 04:48 PM UTC

Closing as a subPR of #593 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/591*
