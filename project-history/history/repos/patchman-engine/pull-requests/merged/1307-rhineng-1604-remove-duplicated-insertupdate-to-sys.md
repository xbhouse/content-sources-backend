---
type: pull_request
number: 1307
title: "RHINENG-1604: remove duplicated insert/update to system_platform"
state: merged
author: psegedy
created: 2023-09-06T16:32:08Z
updated: 2023-09-11T13:26:56Z
closed: 2023-09-11T13:26:56Z
merged: 2023-09-11T13:26:56Z
base_branch: master
head_branch: listener_commits
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1307
---

# Pull Request #1307: RHINENG-1604: remove duplicated insert/update to system_platform

**Author**: @psegedy
**Created**: September 06, 2023 at 04:32 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener_commits`

## Description

insert/update should be handled by the `storeOrUpdateSysPlatform` function which should avoid doing upserts
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

### Comment by @jira-linking on September 06, 2023 at 04:32 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1604


---

## Reviews

### Review by @MichaelMraka - Approved on September 07, 2023 at 09:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1307*
