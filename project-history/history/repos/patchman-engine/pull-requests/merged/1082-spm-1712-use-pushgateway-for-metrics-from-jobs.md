---
type: pull_request
number: 1082
title: "SPM-1712: use pushgateway for metrics from jobs"
state: merged
author: psegedy
created: 2022-08-26T11:59:52Z
updated: 2022-08-30T12:51:09Z
closed: 2022-08-30T12:51:09Z
merged: 2022-08-30T12:51:08Z
base_branch: master
head_branch: jobs_metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1082
---

# Pull Request #1082: SPM-1712: use pushgateway for metrics from jobs

**Author**: @psegedy
**Created**: August 26, 2022 at 11:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `jobs_metrics`

## Description

for collecting metrics from jobs we need to use pushgateway, not sure how to test it locally or in ephemeral

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

### Review by @MichaelMraka - Approved on August 29, 2022 at 12:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1082*
