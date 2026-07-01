---
type: pull_request
number: 1149
title: "return 500 when app panics"
state: merged
author: psegedy
created: 2022-12-07T13:02:55Z
updated: 2022-12-08T12:28:57Z
closed: 2022-12-08T12:28:56Z
merged: 2022-12-08T12:28:56Z
base_branch: master
head_branch: 500
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1149
---

# Pull Request #1149: return 500 when app panics

**Author**: @psegedy
**Created**: December 07, 2022 at 01:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `500`

## Description

we need to return 500 when app panics, currently, the app will not respond to api call in case of panic
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

### Comment by @jira-linking on December 07, 2022 at 01:02 PM UTC

Commits missing Jira IDs:
dce29b443d095f61d0bd98bd08b68912bf555520


---

## Reviews

### Review by @MichaelMraka - Approved on December 08, 2022 at 10:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1149*
