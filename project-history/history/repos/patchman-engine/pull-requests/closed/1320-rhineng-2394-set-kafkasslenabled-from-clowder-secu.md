---
type: pull_request
number: 1320
title: "RHINENG-2394: set KafkaSslEnabled from clowder SecurityProtocol"
state: closed
author: psegedy
created: 2023-10-05T10:30:23Z
updated: 2023-10-10T13:20:04Z
closed: 2023-10-10T13:20:04Z
base_branch: master
head_branch: fix_kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1320
---

# Pull Request #1320: RHINENG-2394: set KafkaSslEnabled from clowder SecurityProtocol

**Author**: @psegedy
**Created**: October 05, 2023 at 10:30 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix_kafka`

## Description

Set as mentioned in `Upcoming Stage Changes for Kafka Migration Testing` email thread
* if kafka.brokers[0].securityProtocol is defined and its value is "SSL" or "SASL_SSL", then SSL should be ON

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

### Comment by @jira-linking on October 05, 2023 at 10:30 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2394


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1320*
