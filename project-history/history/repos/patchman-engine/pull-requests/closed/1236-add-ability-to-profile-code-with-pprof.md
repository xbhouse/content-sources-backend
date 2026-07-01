---
type: pull_request
number: 1236
title: "Add ability to profile code with pprof"
state: closed
author: psegedy
created: 2023-06-07T17:24:10Z
updated: 2023-06-07T17:47:27Z
closed: 2023-06-07T17:47:27Z
base_branch: master
head_branch: profile
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1236
---

# Pull Request #1236: Add ability to profile code with pprof

**Author**: @psegedy
**Created**: June 07, 2023 at 05:24 PM UTC
**Status**: Closed
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

### Comment by @jira-linking on June 07, 2023 at 05:24 PM UTC

Commits missing Jira IDs:
9aba057af4609f5a0c266f8e91f666d381bd87bb
f0984faec21259cac27dc2e747c8071100e7b254


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1236*
