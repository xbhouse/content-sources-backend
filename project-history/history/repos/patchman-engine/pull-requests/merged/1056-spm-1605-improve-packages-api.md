---
type: pull_request
number: 1056
title: "SPM-1605: improve /packages API"
state: merged
author: MichaelMraka
created: 2022-08-15T10:56:04Z
updated: 2022-08-18T13:12:28Z
closed: 2022-08-18T13:12:28Z
merged: 2022-08-18T13:12:28Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1056
---

# Pull Request #1056: SPM-1605: improve /packages API

**Author**: @MichaelMraka
**Created**: August 15, 2022 at 10:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

Moves summary from `package_latest_cache` to `package_name` which simplifies query behind /packages.
Also removes then unused `package_latest_cache`.

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

### Comment by @psegedy on August 15, 2022 at 11:24 AM UTC

please remove cache refresh from https://github.com/RedHatInsights/patchman-engine/blob/master/dev/test_data.sql#L168

### Comment by @psegedy on August 16, 2022 at 04:35 PM UTC

Here's my observation about thirdparty packages. When you upload the system with thirdparty package, the package is inserted in package_name but then it is not found by https://github.com/RedHatInsights/patchman-engine/blob/9de1aeca7ff735dd3c86f1f7165df3cc20e70231/evaluator/package_cache.go#L251-L254. With the second upload (after removing the system), it is evaluated correctly. Are we missing db.commit when storing lazy packages?

Anyways, based on the code changes it does not seem to be caused by this PR, therefore it LGTM.

---

## Reviews

### Review by @psegedy - Approved on August 16, 2022 at 04:36 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1056*
