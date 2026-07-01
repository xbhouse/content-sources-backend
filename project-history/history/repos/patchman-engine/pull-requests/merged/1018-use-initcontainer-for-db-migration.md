---
type: pull_request
number: 1018
title: "use initContainer for db migration"
state: merged
author: psegedy
created: 2022-07-11T15:46:03Z
updated: 2022-07-15T08:50:48Z
closed: 2022-07-15T08:50:48Z
merged: 2022-07-15T08:50:48Z
base_branch: master
head_branch: replace_db_admin
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1018
---

# Pull Request #1018: use initContainer for db migration

**Author**: @psegedy
**Created**: July 11, 2022 at 03:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `replace_db_admin`

## Description

Initially I wanted to use k8s job which runs migrations and add initContainer to all deployements to wait till migration is ready. However, using plain job suffers from clowder usage - it doesn't have the correct config and can't pull the image and ClowdJob triggered by ClowdJobInvoction can't be run if the app is not running so it can't be used for migrations.

The implementation is basically the same as in Vulnerability
1. manager pod has initContainer which runs migrations
2. to avoid concurrency we are using pg_advisory_lock
3. if the migration is not successful, init container panics which _should_ result in failure and new manager won't be deployed
4. other pods (evaluators, listeners,...) will start immediately and they may be in CrashLoop state until migrations are not successful

It is not the ideal solution, any recommendations?

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

### Comment by @jira-linking on July 11, 2022 at 03:46 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1616


---

## Reviews

### Review by @MichaelMraka - Commented on July 12, 2022 at 08:44 AM UTC

### Review by @psegedy - Commented on July 12, 2022 at 09:29 AM UTC

### Review by @MichaelMraka - Approved on July 14, 2022 at 08:11 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1018*
