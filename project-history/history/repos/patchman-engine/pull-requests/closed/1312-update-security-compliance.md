---
type: pull_request
number: 1312
title: "Update security compliance"
state: closed
author: SteveHNH
created: 2023-09-15T17:02:25Z
updated: 2023-10-05T09:54:54Z
closed: 2023-10-05T09:54:54Z
base_branch: security-compliance
head_branch: update_security_compliance
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1312
---

# Pull Request #1312: Update security compliance

**Author**: @SteveHNH
**Created**: September 15, 2023 at 05:02 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `security-compliance` ← **Head**: `update_security_compliance`

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

## Discussion

### Comment by @jira-linking on September 15, 2023 at 05:02 PM UTC

Commits missing Jira IDs:
d6906c606a1b10de39f483d32c8f6fe5abffb07d
e8b1e38d0f875a4efdf2ed29a3109bbf12c74c3a
61a96a9b89c129bf0d1178144d1c48b5cf3a4f2a
546a7d8cfd42955e1a7d9766c60ea2f266069d28
9ab64d8c8724d2cad8535c50b7871b491900c219
f0aa0d1d16ae0f4b9bb892a0ab2844c3ae87847b
a8dd28641b758df417ebe66057027296d5aea4af
70fea51a7680f37665f729d09edcd12f82e7c86d
9d128cf1bd6016c11a1e398ac61ee6d39f05b8d8
44dc62351e89b0dfa7e7b063bd61ac4231c904c6
56fa90283c6736a23c8c7276693e237f0f9df3ef
820d2bdada9396a3e5551c8b787a079534b6b993
bffd81ab577adc70beb5fb2ccfeb58070fca0897
7b4fb806a444c035148eeac9b77bad11fb09875e
3b0b343c1fb596768044ce1f0b93f9f721951899
2c48a453ad2b5350ca9abb1dcd408fa1b06c31ba
21ba8d9462970a7cfec992867ffba9f920748f94
be1dafbe5b3af1d450b980d448cced2f6b4c5826
bf83b9137bd7df92689511e0e18942f9e2689843
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1816
https://issues.redhat.com/browse/SPM-2024
https://issues.redhat.com/browse/SPM-2019
https://issues.redhat.com/browse/RHINENG-390
https://issues.redhat.com/browse/SPM-2071
https://issues.redhat.com/browse/SPM-2064
https://issues.redhat.com/browse/SPM-2077
https://issues.redhat.com/browse/SPM-2079
https://issues.redhat.com/browse/SPM-2080


### Comment by @MichaelMraka on September 18, 2023 at 09:57 AM UTC

Are you sure you want to merge master code which contains not yet verified code and enabled features?

### Comment by @SteveHNH on September 18, 2023 at 04:33 PM UTC

@MichaelMraka we've mainly been tracking with master for most repos because they use that as their release branch, but I _did_ just notice you have a stable branch. Would that be the better choice?

### Comment by @MichaelMraka on September 19, 2023 at 09:06 AM UTC

@SteveHNH Unfortunately no, `stable` branch is unused for about 2 years.
We don't have any prod tracking branch.
Deployments to the prod go from version tags (in`master`) after QEs verify them. 

### Comment by @SteveHNH on September 19, 2023 at 05:40 PM UTC

@MichaelMraka Okay, that makes sense. That's actually a good way for us to keep track of what's out there. I see how it's tied together in app interface as well. I think we can close this MR, and we might duplicate your current production commit into this branch

---

## Reviews

### Review by @casey-williams-rh - Approved on September 15, 2023 at 05:12 PM UTC

👍 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1312*
