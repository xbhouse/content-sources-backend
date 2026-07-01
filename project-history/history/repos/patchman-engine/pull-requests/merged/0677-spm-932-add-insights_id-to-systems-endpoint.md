---
type: pull_request
number: 677
title: "SPM-932: add insights_id to /systems endpoint"
state: merged
author: mkholjuraev
created: 2021-05-18T15:38:12Z
updated: 2021-05-31T06:00:07Z
closed: 2021-05-31T06:00:07Z
merged: 2021-05-31T06:00:07Z
base_branch: master
head_branch: add-insigths-id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/677
---

# Pull Request #677: SPM-932: add insights_id to /systems endpoint

**Author**: @mkholjuraev
**Created**: May 18, 2021 at 03:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-insigths-id`

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

### Comment by @josef-hak on May 27, 2021 at 04:06 AM UTC

@mkholjuraev it's not rebased with master, right? Could you do it, pls? It should fix the 4th testing pipeline.

### Comment by @mkholjuraev on May 30, 2021 at 10:32 PM UTC

@Josca sorry, my PRs taking too much time and effort :). I have rebased with master, but 4th check did not pass again

### Comment by @josef-hak on May 31, 2021 at 04:55 AM UTC

/retest


---

## Reviews

### Review by @MichaelMraka - Approved on May 20, 2021 at 08:45 AM UTC

looks great!

### Review by @josef-hak - Changes Requested on May 24, 2021 at 12:09 PM UTC

Thanks, I like it, just one thing.

### Review by @josef-hak - Changes Requested on May 27, 2021 at 03:58 AM UTC

Small change requested. And pls, fix the typo in commit message (SOM -> SPM), and add SPM-932 also as a prefix of PR title. It will be automatically assigned to the Jira ticket, thanks.

### Review by @josef-hak - Approved on May 31, 2021 at 05:41 AM UTC

@mkholjuraev thanks. Reg. failing clowder pipeline - it's not your fault. Probably something is wrong in different place. Other PRs have it's failing too.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/677*
