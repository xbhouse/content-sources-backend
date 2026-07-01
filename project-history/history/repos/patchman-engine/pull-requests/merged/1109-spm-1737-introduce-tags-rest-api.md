---
type: pull_request
number: 1109
title: "SPM-1737: introduce tags REST API"
state: merged
author: vkrizan
created: 2022-09-20T13:52:16Z
updated: 2022-09-21T14:31:56Z
closed: 2022-09-21T14:31:56Z
merged: 2022-09-21T14:31:56Z
base_branch: master
head_branch: feat-effective-tags-1737
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1109
---

# Pull Request #1109: SPM-1737: introduce tags REST API

**Author**: @vkrizan
**Created**: September 20, 2022 at 01:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `feat-effective-tags-1737`

## Description

Introduces `/tags` which returns tags with host counts effective for this service.

Along the way it refactors following:
* moves `SystemTag` type
* use of common `ih` alias for `inventory.hosts` table within tags filter
* filter tags where clause split

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

### Comment by @jira-linking on September 20, 2022 at 01:52 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1737


---

## Reviews

### Review by @vkrizan - Commented on September 20, 2022 at 02:01 PM UTC

### Review by @MichaelMraka - Approved on September 21, 2022 at 09:14 AM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 10:41 AM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 10:50 AM UTC

### Review by @psegedy - Changes Requested on September 21, 2022 at 08:46 AM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 11:08 AM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 11:18 AM UTC

### Review by @psegedy - Commented on September 21, 2022 at 11:43 AM UTC

### Review by @psegedy - Changes Requested on September 21, 2022 at 12:05 PM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 12:45 PM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 12:52 PM UTC

### Review by @vkrizan - Commented on September 21, 2022 at 12:52 PM UTC

### Review by @psegedy - Commented on September 21, 2022 at 02:31 PM UTC

### Review by @psegedy - Approved on September 21, 2022 at 02:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1109*
