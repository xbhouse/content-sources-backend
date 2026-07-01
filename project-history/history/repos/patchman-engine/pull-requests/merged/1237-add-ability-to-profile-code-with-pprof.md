---
type: pull_request
number: 1237
title: "add ability to profile code with pprof"
state: merged
author: psegedy
created: 2023-06-07T17:48:05Z
updated: 2023-06-12T08:30:16Z
closed: 2023-06-12T08:30:16Z
merged: 2023-06-12T08:30:16Z
base_branch: master
head_branch: profile
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1237
---

# Pull Request #1237: add ability to profile code with pprof

**Author**: @psegedy
**Created**: June 07, 2023 at 05:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `profile`

## Description

Usage:
- local development in docker-compose
  - `go tool pprof http://localhost:{9000|9002|9003|9004}/debug/pprof/{heap|profile|block|mutex}`
  - `curl -o trace.out http://localhost:{9000|9002|9003|9004}/debug/pprof/trace?seconds=5 && go tool trace trace.out`
- admin (turnpike) api
  - `/api/patch/admin/pprof/{manager|listener|evaluator_upload|evaluator_recalc}/{heap|profile|block|mutex|trace}`
  - `go tool pprof <saved.file>`
  - `go tool trace <saved.file>`

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

### Comment by @jira-linking on June 07, 2023 at 05:48 PM UTC

Commits missing Jira IDs:
78f8d8338c8fd0241019e0218e09e256ceec4f87
58296748d9ea6d072460579a01ca84304f01bc4a
0557eeb19a3d7579ba4478c748428b774b1c486d


### Comment by @psegedy on June 09, 2023 at 09:15 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Commented on June 12, 2023 at 08:07 AM UTC

### Review by @MichaelMraka - Approved on June 12, 2023 at 08:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1237*
