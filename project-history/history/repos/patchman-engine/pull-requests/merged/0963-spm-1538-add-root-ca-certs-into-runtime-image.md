---
type: pull_request
number: 963
title: "SPM-1538: add root ca certs into runtime image"
state: merged
author: MichaelMraka
created: 2022-05-26T09:42:23Z
updated: 2022-05-26T12:17:29Z
closed: 2022-05-26T12:17:29Z
merged: 2022-05-26T12:17:29Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/963
---

# Pull Request #963: SPM-1538: add root ca certs into runtime image

**Author**: @MichaelMraka
**Created**: May 26, 2022 at 09:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

## Description

fixes
{"@timestamp":"2022-05-24T11:26:48.842Z","err":"RequestError: send request failed\ncaused by: Post \"https://logs.us-east-1.amazonaws.com/\": x509: certificate signed by unknown authority","level":"error","message":"unable to setup CloudWatch logging"}

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

### Review by @psegedy - Approved on May 26, 2022 at 12:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/963*
