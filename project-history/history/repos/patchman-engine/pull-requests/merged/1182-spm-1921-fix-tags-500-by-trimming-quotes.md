---
type: pull_request
number: 1182
title: "SPM-1921: fix tags 500 by trimming quotes"
state: merged
author: psegedy
created: 2023-03-02T11:52:46Z
updated: 2023-03-02T15:47:37Z
closed: 2023-03-02T15:47:36Z
merged: 2023-03-02T15:47:36Z
base_branch: master
head_branch: 500_2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1182
---

# Pull Request #1182: SPM-1921: fix tags 500 by trimming quotes

**Author**: @psegedy
**Created**: March 02, 2023 at 11:52 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `500_2`

## Description

example of another query which caused 500 on tags filter
```sh
# tags='insights-client/demo="{{slug}}"'
tags=%27insights-client%2Fdemo%3D%22%7B%7Bslug%7D%7D%22%27
```
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

### Comment by @jira-linking on March 02, 2023 at 11:52 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1921


### Comment by @psegedy on March 02, 2023 at 02:02 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on March 02, 2023 at 03:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1182*
