---
type: pull_request
number: 989
title: "SPM-1587: refresh cache after vmaas sync"
state: merged
author: psegedy
created: 2022-06-24T12:50:58Z
updated: 2022-06-28T12:37:52Z
closed: 2022-06-28T12:37:52Z
merged: 2022-06-28T12:37:52Z
base_branch: master
head_branch: sync_refresh_cache
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/989
---

# Pull Request #989: SPM-1587: refresh cache after vmaas sync

**Author**: @psegedy
**Created**: June 24, 2022 at 12:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sync_refresh_cache`

## Description

Currently, caches are created every 15 min. This means that if vmaas-sync finishes, we still need to wait for these jobs to trigger cache refresh to have latest data.
E.g. sync finishes, but /packages api is still not showing any packages since there wasn't a cache refresh

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

### Comment by @psegedy on June 24, 2022 at 02:57 PM UTC

I removed mutex since these functions are using DB transactions anyway

---

## Reviews

### Review by @michalslomczynski - Commented on June 24, 2022 at 01:01 PM UTC

### Review by @michalslomczynski - Approved on June 27, 2022 at 11:49 AM UTC

### Review by @MichaelMraka - Approved on June 28, 2022 at 11:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/989*
