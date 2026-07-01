---
type: pull_request
number: 767
title: "SPM-1089: Fix setting of thid_party flag for repos"
state: merged
author: semtexzv
created: 2021-08-06T11:40:14Z
updated: 2021-08-11T12:50:21Z
closed: 2021-08-11T12:50:21Z
merged: 2021-08-11T12:50:21Z
base_branch: master
head_branch: repos-fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/767
---

# Pull Request #767: SPM-1089: Fix setting of thid_party flag for repos

**Author**: @semtexzv
**Created**: August 06, 2021 at 11:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `repos-fix`

## Description

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

### Review by @josef-hak - Changes Requested on August 06, 2021 at 12:21 PM UTC

1) Tests fail:
~~~
 zookeeper    | ===> Check if /var/lib/zookeeper/log is writable ...
test         | 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100  118k    0 36708  100 84892  86780   195k --:--:-- --:--:-- --:--:--  280k
test         | 1762,1764c1762
test         | <             }
test         | <         },
test         | <         "/api/patch/v1/systems/{inventory_id}/": {
test         | ---
test         | >             },
test         | docs/openapi.json different from file generated with './scripts/generate_docs.sh'!
test exited with code 1
~~~
2) Please associate PR (and commit) to some ticket.

### Review by @josef-hak - Changes Requested on August 10, 2021 at 11:08 AM UTC

### Review by @semtexzv - Commented on August 10, 2021 at 12:13 PM UTC

### Review by @semtexzv - Commented on August 10, 2021 at 12:13 PM UTC

### Review by @josef-hak - Approved on August 11, 2021 at 12:50 PM UTC

Great, I've tested that in CI and it works. EPEL repos are updated to third_party=true.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/767*
