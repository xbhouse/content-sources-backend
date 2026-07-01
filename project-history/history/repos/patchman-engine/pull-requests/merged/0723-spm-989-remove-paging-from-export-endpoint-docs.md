---
type: pull_request
number: 723
title: "SPM-989: Remove paging from export endpoint docs"
state: merged
author: semtexzv
created: 2021-06-30T11:48:21Z
updated: 2021-07-07T10:00:12Z
closed: 2021-07-07T10:00:12Z
merged: 2021-07-07T10:00:12Z
base_branch: master
head_branch: export-docs-fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/723
---

# Pull Request #723: SPM-989: Remove paging from export endpoint docs

**Author**: @semtexzv
**Created**: June 30, 2021 at 11:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `export-docs-fix`

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

### Comment by @josef-hak on June 30, 2021 at 11:56 AM UTC

@semtexzv ... and please add prefix to the commit: SPM-989, thanks

### Comment by @semtexzv on July 01, 2021 at 04:15 AM UTC

Alright, I have rebased this branch on top of master and re-created the docs repeatedly. I don't know how this error occurs.

---

## Reviews

### Review by @josef-hak - Commented on June 30, 2021 at 11:54 AM UTC

Tests failed when comparing actual and generated docs:
~~~
docs/openapi.json different from file generated with './scripts/generate_docs.sh'!
~~~

### Review by @josef-hak - Approved on July 07, 2021 at 10:00 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/723*
