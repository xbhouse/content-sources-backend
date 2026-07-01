---
type: pull_request
number: 576
title: "SPM-798 updated golangci-lint config"
state: merged
author: josef-hak
created: 2021-03-18T18:44:42Z
updated: 2021-03-19T09:32:34Z
closed: 2021-03-19T07:31:04Z
merged: 2021-03-19T07:31:04Z
base_branch: master
head_branch: lint
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/576
---

# Pull Request #576: SPM-798 updated golangci-lint config

**Author**: @josef-hak
**Created**: March 18, 2021 at 06:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `lint`

## Description

Resolve warnings printed by `golangci-lint run`:
~~~
WARN [runner] The linter 'interfacer' is deprecated due to: The repository of the linter has been archived by the owner. 
WARN [runner] The linter 'maligned' is deprecated due to: The repository of the linter has been archived by the owner. Use govet 'fieldalignment' instead.
~~~

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

### Review by @MichaelMraka - Approved on March 19, 2021 at 07:11 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/576*
