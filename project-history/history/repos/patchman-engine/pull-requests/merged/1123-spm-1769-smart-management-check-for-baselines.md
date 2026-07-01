---
type: pull_request
number: 1123
title: "SPM-1769: smart management check for baselines"
state: merged
author: psegedy
created: 2022-10-11T13:20:45Z
updated: 2022-10-12T14:03:43Z
closed: 2022-10-12T14:03:43Z
merged: 2022-10-12T14:03:43Z
base_branch: master
head_branch: entitlementstoo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1123
---

# Pull Request #1123: SPM-1769: smart management check for baselines

**Author**: @psegedy
**Created**: October 11, 2022 at 01:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `entitlementstoo`

## Description

- use https://github.com/RedHatInsights/identity/blob/main/identity.go instead of https://github.com/RedHatInsights/platform-go-middlewares/blob/master/identity/identity.go per discussion on #forum-consoledot 
- add `middlewares.EntitlementsAuthenticator`
- drop dev mode bypass for authetication, it's better to use identity header with every local curl request
- common function for getting x-rh-identity from gin context

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

### Comment by @jira-linking on October 11, 2022 at 01:20 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1769


---

## Reviews

### Review by @MichaelMraka - Approved on October 12, 2022 at 12:55 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1123*
