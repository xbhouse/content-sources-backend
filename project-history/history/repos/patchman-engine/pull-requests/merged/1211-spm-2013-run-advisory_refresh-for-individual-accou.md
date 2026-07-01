---
type: pull_request
number: 1211
title: "SPM-2013: run advisory_refresh for individual accounts in 4 goroutines"
state: merged
author: psegedy
created: 2023-04-14T12:33:29Z
updated: 2023-04-17T08:22:08Z
closed: 2023-04-17T08:22:08Z
merged: 2023-04-17T08:22:08Z
base_branch: master
head_branch: fix_advisory_refresh
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1211
---

# Pull Request #1211: SPM-2013: run advisory_refresh for individual accounts in 4 goroutines

**Author**: @psegedy
**Created**: April 14, 2023 at 12:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_advisory_refresh`

## Description

- run refresh for single account at a time
- run in 4 goroutines so we don't need to wait for refresh of bigger accounts
- add `SKIP_N_ACCOUNTS_REFRESH` to skip accounts refresh if something goes wrong and we want to continue with refresh

previous solution didn't finish in prod even after 6 hours with evaluator, listener, and all other jobs disabled
this finished in 9 minutes with evaluator, listener and jobs disabled

it is already pushed to prod using hotfix branch
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

### Comment by @jira-linking on April 14, 2023 at 12:33 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2013


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1211*
