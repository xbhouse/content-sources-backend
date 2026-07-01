---
type: pull_request
number: 981
title: "SPM-1580: don't block websocket client with msg handler"
state: merged
author: psegedy
created: 2022-06-20T12:28:32Z
updated: 2022-06-22T14:30:17Z
closed: 2022-06-22T14:30:16Z
merged: 2022-06-22T14:30:16Z
base_branch: master
head_branch: websocket-fail
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/981
---

# Pull Request #981: SPM-1580: don't block websocket client with msg handler

**Author**: @psegedy
**Created**: June 20, 2022 at 12:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `websocket-fail`

## Description

fixing 
```
time="2022-06-20T10:27:03Z" level=error msg="Failed to retrieve VMaaS websocket message" err="websocket: close 1005 (no status)"
time="2022-06-20T10:27:03Z" level=error msg="Websocket error occurred, waiting" err="websocket: close 1005 (no status)"
```

vmaas closes websocket connection if subscribed client won't respond to `ping` message in predetermined timeout. Patch is not responding because it's syncing data which blocks websocket client.

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

### Comment by @psegedy on June 20, 2022 at 04:07 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on June 22, 2022 at 07:11 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/981*
