---
type: pull_request
number: 1389
title: "RHINENG-8389: make /ids responses consistent"
state: merged
author: Dugowitch
created: 2024-03-05T16:35:53Z
updated: 2024-03-07T17:06:23Z
closed: 2024-03-07T17:06:23Z
merged: 2024-03-07T17:06:23Z
base_branch: master
head_branch: RHINENG-8389
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1389
---

# Pull Request #1389: RHINENG-8389: make /ids responses consistent

**Author**: @Dugowitch
**Created**: March 05, 2024 at 04:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-8389`

## Description

Currently some */ids* endpoints respond with a list of ids in `ids` section, other endpoints respond with `ids` and `data` sections. After the change, all endpoint responses should have both `ids` and `data` sections, even if the data section consists only of `IDPlain`s.

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

### Comment by @jira-linking on March 05, 2024 at 04:35 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8389


### Comment by @psegedy on March 07, 2024 at 03:27 PM UTC

/retest

---

## Reviews

### Review by @psegedy - Commented on March 06, 2024 at 09:46 AM UTC

### Review by @Dugowitch - Commented on March 07, 2024 at 10:37 AM UTC

### Review by @Dugowitch - Commented on March 07, 2024 at 10:37 AM UTC

### Review by @psegedy - Approved on March 07, 2024 at 05:06 PM UTC

looks great, thanks!

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1389*
