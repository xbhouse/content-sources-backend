---
type: pull_request
number: 1427
title: "RHINENG-9873: rework pod configuration"
state: merged
author: MichaelMraka
created: 2024-05-16T13:50:05Z
updated: 2024-05-21T14:06:29Z
closed: 2024-05-21T14:06:29Z
merged: 2024-05-21T14:06:29Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1427
---

# Pull Request #1427: RHINENG-9873: rework pod configuration

**Author**: @MichaelMraka
**Created**: May 16, 2024 at 01:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @jira-linking on May 16, 2024 at 01:50 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9873


### Comment by @psegedy on May 20, 2024 at 10:05 AM UTC

looks like there is something wrong with some env var for migrations
```
DB migration in progress, waiting for schema=-1
current version: 0, expected: -1
```

### Comment by @psegedy on May 20, 2024 at 04:53 PM UTC

/retest

---

## Reviews

### Review by @psegedy - Commented on May 20, 2024 at 12:52 PM UTC

### Review by @psegedy - Approved on May 21, 2024 at 11:51 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1427*
