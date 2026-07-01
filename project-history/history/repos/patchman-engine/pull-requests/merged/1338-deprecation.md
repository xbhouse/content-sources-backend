---
type: pull_request
number: 1338
title: "Deprecation"
state: merged
author: psegedy
created: 2023-11-08T18:00:15Z
updated: 2023-11-10T18:30:53Z
closed: 2023-11-10T18:30:53Z
merged: 2023-11-10T18:30:53Z
base_branch: master
head_branch: deprecation
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1338
---

# Pull Request #1338: Deprecation

**Author**: @psegedy
**Created**: November 08, 2023 at 06:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `deprecation`

## Description

- [x] Generic Deprecation interface for future deprecations
- [x] API Deprecation struct for deprecation of whole api version or single endpoint
  - Set `Deprecation` and `Sunset` headers
  - 301 Redirect after `redirectTimestamp`, redirect location can be set either by providing `redirectLocation` or by modifying the current location with `redirectReplacer`
  - 401 Gone after `deprecationTimestamp`
- [x] Deprecation middleware
- [x] Deprecation of /v1 and /v2 api versions
  - Deprecation and Sunset headers set to 2024-01-31
  - 301 Redirect to /v3 since 2024-01-01
  - 401 Gone returned since  2024-01-31
- [x] Set v1 and v2 as deprecated in docs/{v1|v2}/openapi.json
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

### Comment by @jira-linking on November 08, 2023 at 06:00 PM UTC

Commits missing Jira IDs:
a6141150e20614e294792dafc6a92d1516f90e58
ddd1e98391749d840acc4dbaabd71b5dcef11fb0
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-3104


---

## Reviews

### Review by @MichaelMraka - Approved on November 10, 2023 at 10:19 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1338*
