---
type: pull_request
number: 1005
title: "SPM-1484: Add empty metadata object to bypass NPE further on"
state: merged
author: michalslomczynski
created: 2022-07-04T10:05:05Z
updated: 2022-07-07T13:41:06Z
closed: 2022-07-07T13:41:06Z
merged: 2022-07-07T13:41:06Z
base_branch: master
head_branch: notif-log
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1005
---

# Pull Request #1005: SPM-1484: Add empty metadata object to bypass NPE further on

**Author**: @michalslomczynski
**Created**: July 04, 2022 at 10:05 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `notif-log`

## Description

Notification format was valid in `Notification Validator` but lack of metadata key happened to raise NPE on stage. Here is a quick fix to avoid this issue.

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

### Comment by @jira-linking on July 04, 2022 at 10:05 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1484


### Comment by @digitronik on July 04, 2022 at 01:09 PM UTC

/restest

### Comment by @michalslomczynski on July 05, 2022 at 12:49 PM UTC

@psegedy @MichaelMraka Should be merged and deployed to stage if I will be absent. Thanks!

### Comment by @michalslomczynski on July 05, 2022 at 12:52 PM UTC

If full notification log is in your opinion too verbose, it can be reverted to previous format.

### Comment by @digitronik on July 07, 2022 at 07:17 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Changes Requested on July 07, 2022 at 08:42 AM UTC

### Review by @MichaelMraka - Approved on July 07, 2022 at 01:38 PM UTC

### Review by @psegedy - Approved on July 07, 2022 at 01:39 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1005*
