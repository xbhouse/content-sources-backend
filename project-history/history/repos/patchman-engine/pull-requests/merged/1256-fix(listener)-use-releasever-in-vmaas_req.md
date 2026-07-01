---
type: pull_request
number: 1256
title: "fix(listener): use releasever in vmaas_req"
state: merged
author: vkrizan
created: 2023-07-03T14:26:18Z
updated: 2023-07-04T11:16:57Z
closed: 2023-07-04T11:16:56Z
merged: 2023-07-04T11:16:56Z
base_branch: master
head_branch: fix-relesever-use-SPM-2137
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1256
---

# Pull Request #1256: fix(listener): use releasever in vmaas_req

**Author**: @vkrizan
**Created**: July 03, 2023 at 02:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-relesever-use-SPM-2137`

## Description

Utilize relelasever from a system profile if RHSM version isn't set. This should improve reporting from RHUI systems.

SPM-2137

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

### Comment by @jira-linking on July 03, 2023 at 02:26 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2137


### Comment by @vkrizan on July 03, 2023 at 02:28 PM UTC

@MichaelMraka @psegedy heya, would you guys be able to take this over from me? I'm not sure whether you'd like to extend the APIs as well to include `releasever`.

### Comment by @MichaelMraka on July 04, 2023 at 11:15 AM UTC

@vkrizan there's no issue with API. This is correct way to fix it.

---

## Reviews

### Review by @MichaelMraka - Approved on July 04, 2023 at 11:16 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1256*
