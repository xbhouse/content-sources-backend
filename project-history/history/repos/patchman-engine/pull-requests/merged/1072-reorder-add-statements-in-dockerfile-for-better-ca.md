---
type: pull_request
number: 1072
title: "reorder add statements in dockerfile for better caching"
state: merged
author: psegedy
created: 2022-08-19T15:36:36Z
updated: 2022-08-22T16:56:34Z
closed: 2022-08-22T16:56:34Z
merged: 2022-08-22T16:56:34Z
base_branch: master
head_branch: reorder_dockerfile
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1072
---

# Pull Request #1072: reorder add statements in dockerfile for better caching

**Author**: @psegedy
**Created**: August 19, 2022 at 03:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `reorder_dockerfile`

## Description

these are the most frequently changed files
```
✗ git log --pretty=format: --name-only | sort | uniq -c | sort -rg | head -20

1781
 266 manager/manager.go
 265 docs/openapi.json
 248 base/metrics/metrics.go
 177 evaluator/evaluate.go
 128 listener/upload.go
  88 database_admin/schema/create_schema.sql
  87 manager/controllers/systems.go
  82 docs/v2/openapi.json
  77 go.mod
  75 manager/controllers/utils.go
  75 dev/test_data.sql
  74 manager/controllers/advisories.go
  71 vmaas_sync/advisory_sync.go
  71 manager/controllers/systems_test.go
  71 go.sum
  69 evaluator/evaluate_test.go
  68 docker-compose.yml
  68 base/models/models.go
  66 listener/upload_test.go
```
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

### Comment by @MichaelMraka on August 22, 2022 at 08:35 AM UTC

That's good idea how to reorder it. IMHO We should better compare usage of whole top directories not single files, i.e.
$ git log --pretty=format: --name-only | **sed 's,/.*$,,'** | sort | uniq -c |sort -rg | head -20
```
   1781 
   1689 manager
    831 base
    481 vmaas_sync
    417 listener
    400 evaluator
    393 docs
    363 database_admin
    214 dev
    189 conf
    182 prototypes
    178 database
    168 scripts
    113 platform
     90 openshift
     77 go.mod
     71 go.sum
     68 docker-compose.yml
     61 deploy
     43 Dockerfile.centos
```

### Comment by @psegedy on August 22, 2022 at 09:05 AM UTC

good suggestion, I've reordered it

---

## Reviews

### Review by @MichaelMraka - Approved on August 22, 2022 at 11:02 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1072*
