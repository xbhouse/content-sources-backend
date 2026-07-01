---
type: pull_request
number: 1126
title: "SPM-1775: fix developer setup"
state: merged
author: MichaelMraka
created: 2022-10-12T12:27:13Z
updated: 2022-10-13T13:31:13Z
closed: 2022-10-13T13:31:12Z
merged: 2022-10-13T13:31:12Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1126
---

# Pull Request #1126: SPM-1775: fix developer setup

**Author**: @MichaelMraka
**Created**: October 12, 2022 at 12:27 PM UTC
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

### Comment by @jira-linking on October 12, 2022 at 12:27 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1775


### Comment by @MichaelMraka on October 12, 2022 at 02:44 PM UTC

I can add it to the `admin` container.
Yes, evaluators still report group issues but they connect after while.

---

## Reviews

### Review by @psegedy - Commented on October 12, 2022 at 02:03 PM UTC

shouldn't we use `./dev/scripts/docker-compose-entrypoint.sh` also for `admin` container?

I'm getting the following error in `evaluator_upload` but it might be an issue with my local kafka
```
time="2022-10-12T12:55:11Z" level=error msg="Unable to establish connection to consumer group coordinator for group patchman: [15] Group Coordinator Not Available: the broker returns this error code for group coordinator requests, offset commits, and most group management requests if the offsets topic has not yet been created, or if the group coordinator is not active" type=kafka
time="2022-10-12T12:55:11Z" level=error msg="[15] Group Coordinator Not Available: the broker returns this error code for group coordinator requests, offset commits, and most group management requests if the offsets topic has not yet been created, or if the group coordinator is not active" type=kafka
```

Otherwise it looks good when I tried `docker-compose down -v` and `docker-compose up --build`

### Review by @psegedy - Approved on October 13, 2022 at 12:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1126*
