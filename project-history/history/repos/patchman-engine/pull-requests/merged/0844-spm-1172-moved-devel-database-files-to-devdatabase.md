---
type: pull_request
number: 844
title: "SPM-1172: moved devel database files to dev/database"
state: merged
author: MichaelMraka
created: 2021-10-08T09:43:56Z
updated: 2021-10-08T11:43:34Z
closed: 2021-10-08T11:43:34Z
merged: 2021-10-08T11:43:34Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/844
---

# Pull Request #844: SPM-1172: moved devel database files to dev/database

**Author**: @MichaelMraka
**Created**: October 08, 2021 at 09:43 AM UTC
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

## Reviews

### Review by @josef-hak - Changes Requested on October 08, 2021 at 10:05 AM UTC

... failed tests. It seems some path was not updated.
~~~
ADD failed: file not found in build context or excluded by .dockerignore: stat database/pgca.crt: file does not exist
Service 'platform' failed to build : Build failed
~~~

### Review by @josef-hak - Approved on October 08, 2021 at 11:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/844*
