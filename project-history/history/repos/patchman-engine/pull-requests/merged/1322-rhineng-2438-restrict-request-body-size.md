---
type: pull_request
number: 1322
title: "RHINENG-2438: restrict request body size"
state: merged
author: psegedy
created: 2023-10-11T14:57:28Z
updated: 2023-10-12T08:51:12Z
closed: 2023-10-12T08:51:11Z
merged: 2023-10-12T08:51:11Z
base_branch: master
head_branch: request_body_limit
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1322
---

# Pull Request #1322: RHINENG-2438: restrict request body size

**Author**: @psegedy
**Created**: October 11, 2023 at 02:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `request_body_limit`

## Description

set `http.MaxBytesReader` for request body, when the body is read by ShouldBindJSON and body size exceeds limit, `MaxBytesError` is returned and then we respond with HTTP code 400 with error like `{"error":"Invalid request body: http: request body too large"}`
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

### Comment by @jira-linking on October 11, 2023 at 02:57 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2438


---

## Reviews

### Review by @MichaelMraka - Approved on October 12, 2023 at 08:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1322*
