---
type: pull_request
number: 1259
title: "RHINENG-1162: add missing read replica envs for package refresh job"
state: merged
author: psegedy
created: 2023-07-11T09:51:57Z
updated: 2023-07-12T20:37:22Z
closed: 2023-07-12T20:37:22Z
merged: 2023-07-12T20:37:22Z
base_branch: master
head_branch: job_read_replica
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1259
---

# Pull Request #1259: RHINENG-1162: add missing read replica envs for package refresh job

**Author**: @psegedy
**Created**: July 11, 2023 at 09:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `job_read_replica`

## Description

also change job schedule, some jobs take more than 1 hour which results in the next job to be scheduled just after the first on finishes. It causes jobs to start after 22:00 and interfere with upload evaluation
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

### Comment by @jira-linking on July 11, 2023 at 09:52 AM UTC

Commits missing Jira IDs:
3bad58e6462a2997608e2ad60a8ead65e4c35077
ee26fe19fbe7490ee199e61b7930316fe4160c83


---

## Reviews

### Review by @MichaelMraka - Approved on July 12, 2023 at 07:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1259*
