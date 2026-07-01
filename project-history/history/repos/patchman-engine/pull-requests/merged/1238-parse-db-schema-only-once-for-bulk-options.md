---
type: pull_request
number: 1238
title: "parse db schema only once for bulk options"
state: merged
author: psegedy
created: 2023-06-08T15:10:54Z
updated: 2023-06-12T08:30:27Z
closed: 2023-06-12T08:30:26Z
merged: 2023-06-12T08:30:26Z
base_branch: master
head_branch: bulkInsert
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1238
---

# Pull Request #1238: parse db schema only once for bulk options

**Author**: @psegedy
**Created**: June 08, 2023 at 03:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bulkInsert`

## Description

~~-  [ ] test it properly to make sure it is inserting the same data~~  - will test in stage and compare with prod

it looks like 20% of evaluation time is spent on schema parsing (and functions called by schema.Parse) if I'm interpreting the results from profiler correctly
![image](https://github.com/RedHatInsights/patchman-engine/assets/14837841/1212c609-15c3-44ca-8225-12f01fee492f)

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

### Comment by @jira-linking on June 08, 2023 at 03:10 PM UTC

Commits missing Jira IDs:
b91b548bc67364a0e040a69a83929a777e6fe6cf


### Comment by @psegedy on June 09, 2023 at 11:30 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on June 12, 2023 at 08:25 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1238*
