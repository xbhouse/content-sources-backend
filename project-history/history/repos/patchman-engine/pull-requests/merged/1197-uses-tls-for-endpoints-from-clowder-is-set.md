---
type: pull_request
number: 1197
title: "Uses TLS for Endpoints from Clowder is set"
state: merged
author: petrblaho
created: 2023-03-21T16:52:57Z
updated: 2023-03-22T13:00:04Z
closed: 2023-03-22T13:00:03Z
merged: 2023-03-22T13:00:03Z
base_branch: master
head_branch: RHINENG-100-tls-from-clowder
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1197
---

# Pull Request #1197: Uses TLS for Endpoints from Clowder is set

**Author**: @petrblaho
**Created**: March 21, 2023 at 04:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-100-tls-from-clowder`

## Description

RHINENG-100

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

### Comment by @petrblaho on March 21, 2023 at 06:02 PM UTC

> please change also clowdapp.yaml and add SSL_CERT_DIR env see [RedHatInsights/vmaas@`master`/deploy/clowdapp.yaml#L532-L533](https://github.com/RedHatInsights/vmaas/blob/master/deploy/clowdapp.yaml?rgh-link-date=2023-03-21T17%3A30%3A57Z#L532-L533)

You are right! I hope I managed to add it to all appropriate places.

---

## Reviews

### Review by @psegedy - Changes Requested on March 21, 2023 at 05:30 PM UTC

please change also clowdapp.yaml and add SSL_CERT_DIR env 
see https://github.com/RedHatInsights/vmaas/blob/master/deploy/clowdapp.yaml#L532-L533

### Review by @psegedy - Approved on March 22, 2023 at 12:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1197*
