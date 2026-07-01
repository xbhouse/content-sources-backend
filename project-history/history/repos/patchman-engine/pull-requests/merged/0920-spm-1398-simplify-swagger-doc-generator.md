---
type: pull_request
number: 920
title: "SPM-1398: Simplify swagger doc generator"
state: merged
author: josef-hak
created: 2022-03-09T17:15:48Z
updated: 2022-03-10T22:26:59Z
closed: 2022-03-10T22:26:58Z
merged: 2022-03-10T22:26:58Z
base_branch: master
head_branch: simplify-generator
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/920
---

# Pull Request #920: SPM-1398: Simplify swagger doc generator

**Author**: @josef-hak
**Created**: March 09, 2022 at 05:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `simplify-generator`

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

### Comment by @michalslomczynski on March 09, 2022 at 05:32 PM UTC

Maybe it would be good to add message and exit if `swag` is not found in scope of this PR. I can do that tomorrow if you agree.

---

## Reviews

### Review by @michalslomczynski - Approved on March 10, 2022 at 08:55 AM UTC

Previously it was writing unrelated curl error to openapi.json but it works great now :+1: 
```
./scripts/generate_docs.sh: line 10: swag: command not found
```

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/920*
