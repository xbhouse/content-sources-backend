---
type: pull_request
number: 567
title: "SPM-782: add additional info to systems endpoint"
state: merged
author: mkholjuraev
created: 2021-03-12T14:39:49Z
updated: 2021-03-23T15:15:39Z
closed: 2021-03-23T15:15:39Z
merged: 2021-03-23T15:15:39Z
base_branch: master
head_branch: expose-extra-system_platform-info
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/567
---

# Pull Request #567: SPM-782: add additional info to systems endpoint

**Author**: @mkholjuraev
**Created**: March 12, 2021 at 02:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `expose-extra-system_platform-info`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [ ] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [ ] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Discussion

### Comment by @josef-hak on March 12, 2021 at 02:50 PM UTC

@mkholjuraev there are some conflicts...

---

## Reviews

### Review by @josef-hak - Changes Requested on March 12, 2021 at 03:22 PM UTC

Thanks for contribution, some improvements and it will be perfect :wink: 

### Review by @mkholjuraev - Commented on March 15, 2021 at 12:05 PM UTC

### Review by @josef-hak - Changes Requested on March 18, 2021 at 10:50 AM UTC

Thanks for contribution, however there are some issues to solve:

- Failing on "too long line rule":
`manager/controllers/systems.go:51: line is 131 characters (lll)`

- Also there are some conflicting file.

### Review by @josef-hak - Changes Requested on March 19, 2021 at 12:52 PM UTC

Test fails `TestAdvisorySystemsDefault`.

### Review by @josef-hak - Changes Requested on March 23, 2021 at 11:31 AM UTC

Looks good now, just one more thing.

### Review by @josef-hak - Approved on March 23, 2021 at 01:26 PM UTC

Ok, we can add the test in a separate PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/567*
