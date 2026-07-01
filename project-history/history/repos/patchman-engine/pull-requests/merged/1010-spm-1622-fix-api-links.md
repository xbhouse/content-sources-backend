---
type: pull_request
number: 1010
title: "SPM-1622: fix api links"
state: merged
author: psegedy
created: 2022-07-08T08:40:04Z
updated: 2022-07-08T14:27:48Z
closed: 2022-07-08T14:27:47Z
merged: 2022-07-08T14:27:47Z
base_branch: master
head_branch: fix_links
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1010
---

# Pull Request #1010: SPM-1622: fix api links

**Author**: @psegedy
**Created**: July 08, 2022 at 08:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_links`

## Description

Some links in response metadata were like `/api/patch/v1/advisories`, others were just `/packages`. Either way, they were hardcoded and didn't show the correct path for v2 api. Use `gin.Context.FullPath` to get the correct path for links
 
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

### Comment by @jira-linking on July 08, 2022 at 08:40 AM UTC

Commits missing Jira IDs:
adfc8525a0ceaa2fefdf85caef922b0a38c36b2e
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1622


---

## Reviews

### Review by @MichaelMraka - Approved on July 08, 2022 at 02:16 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1010*
