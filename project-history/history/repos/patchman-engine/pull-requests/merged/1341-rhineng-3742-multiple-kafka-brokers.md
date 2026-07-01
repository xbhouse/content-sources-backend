---
type: pull_request
number: 1341
title: "RHINENG-3742: Multiple kafka brokers"
state: merged
author: psegedy
created: 2023-11-20T10:26:26Z
updated: 2023-11-20T12:31:27Z
closed: 2023-11-20T12:31:27Z
merged: 2023-11-20T12:31:27Z
base_branch: master
head_branch: multiple_brokers
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1341
---

# Pull Request #1341: RHINENG-3742: Multiple kafka brokers

**Author**: @psegedy
**Created**: November 20, 2023 at 10:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `multiple_brokers`

## Description

- we need to use multiple kafka brokers - use KafkaServers for that
- we probably no longer need to overwrite KAFKA_ADDRESS, it doesn't seem to be used anywhere
- handle Endpoint and PrivateEndpoint separately, structs are no longer equivalent after latest changes in Clowder (Enfpoint has extra ApiPath field and we can't convert PrivateEndpoint to Endpoint 😞 )
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

### Comment by @jira-linking on November 20, 2023 at 10:26 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-3742


---

## Reviews

### Review by @MichaelMraka - Approved on November 20, 2023 at 12:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1341*
