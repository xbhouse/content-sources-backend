---
type: pull_request
number: 704
title: "SPM-984: use postgresql sslrootcert if available"
state: merged
author: MichaelMraka
created: 2021-06-17T14:08:37Z
updated: 2021-06-18T08:13:00Z
closed: 2021-06-18T08:13:00Z
merged: 2021-06-18T08:12:59Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/704
---

# Pull Request #704: SPM-984: use postgresql sslrootcert if available

**Author**: @MichaelMraka
**Created**: June 17, 2021 at 02:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

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

### Comment by @MichaelMraka on June 18, 2021 at 08:01 AM UTC

> test         | base/mqueue/event.go:73:14:

I did not touch that file :). 

---

## Reviews

### Review by @josef-hak - Changes Requested on June 18, 2021 at 06:48 AM UTC

tests failed due to lint rule:
~~~
test         | base/mqueue/event.go:73:14: ST1023: should omit type int from declaration; it will be inferred from the right-hand side (stylecheck)
test         | 	var batches int = 0
~~~

### Review by @josef-hak - Approved on June 18, 2021 at 08:10 AM UTC

Looks good, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/704*
