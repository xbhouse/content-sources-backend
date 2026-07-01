---
type: pull_request
number: 1188
title: "SPM-1957: hotfix: fixing Storing system packages: extended protocol limited to 65535"
state: merged
author: MichaelMraka
created: 2023-03-09T14:11:47Z
updated: 2023-03-10T09:54:46Z
closed: 2023-03-10T09:54:46Z
merged: 2023-03-10T09:54:46Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1188
---

# Pull Request #1188: SPM-1957: hotfix: fixing Storing system packages: extended protocol limited to 65535

**Author**: @MichaelMraka
**Created**: March 09, 2023 at 02:11 PM UTC
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

### Comment by @jira-linking on March 09, 2023 at 02:11 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1957


### Comment by @psegedy on March 09, 2023 at 04:04 PM UTC

@MichaelMraka don't we need to base this PR on the version currently deployed in prod? checkout `v2.3.66`

### Comment by @MichaelMraka on March 10, 2023 at 07:54 AM UTC

This is actually hotfix change applied to master.
The "real" prod hotfix is here https://github.com/RedHatInsights/patchman-engine/tree/hotfix

### Comment by @psegedy on March 10, 2023 at 09:11 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on March 10, 2023 at 09:54 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1188*
