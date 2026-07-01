---
type: pull_request
number: 1337
title: "RHINENG-2867: fix store to advisory metadata and system advisories when update does not have an erratum"
state: merged
author: psegedy
created: 2023-11-06T17:08:22Z
updated: 2023-11-07T08:57:34Z
closed: 2023-11-07T08:57:34Z
merged: 2023-11-07T08:57:34Z
base_branch: master
head_branch: fix_advisory_store
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1337
---

# Pull Request #1337: RHINENG-2867: fix store to advisory metadata and system advisories when update does not have an erratum

**Author**: @psegedy
**Created**: November 06, 2023 at 05:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_advisory_store`

## Description

When vmaas (or yum_updates) reports an update without erratum, then this erratum is returned as `''` by `getReportedAdvisories` and in turn `getAdvisoriesFromDB` creates empty `AdvisoryMetadata{}` due to an issue with slice initialization. Empty `AdvisoryMetadata` has `ID: 0` which cannot be stored to `system_advisories` since advisory with id=0 does not exist in `advisory_metadata` table.

```
2023/11/06 16:26:10 /go/src/app/evaluator/evaluate_advisories.go:149
[0.916ms] [rows:0] SELECT id, name FROM "advisory_metadata" WHERE name IN ('')
{"@timestamp":"2023-11-06T16:26:10.118Z","installable":1,"inventoryID":"b46d27c1-a747-4638-8c3a-1362bd5efa22","level":"info","message":"installable advisories"}

2023/11/06 16:26:10 /go/src/app/base/database/batch.go:72 ERROR: insert or update on table "system_advisories_24" violates foreign key constraint "advisory_metadata_id" (SQLSTATE 23503)
[1.183ms] [rows:-] INSERT INTO "system_advisories" ("advisory_id", "first_reported", "rh_account_id", "status_id", "system_id") VALUES (0, NULL, 1, 0, 4200239) ON CONFLICT (rh_account_id, system_id, advisory_id) DO UPDATE SET status_id = excluded.status_id RETURNING *
{"@timestamp":"2023-11-06T16:26:10.124Z","err":"unable to evaluate in database: evaluation with vmaas failed: Unable to evaluate and store results: Advisory analysis failed: Unable to store advisory data: Unable to update system advisories: ERROR: insert or update on table \"system_advisories_24\" violates foreign key constraint \"advisory_metadata_id\" (SQLSTATE 23503)","evalLabel":"upload","inventoryID":"b46d27c1-a747-4638-8c3a-1362bd5efa22","level":"error","message":"Eval message handling"}
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

### Comment by @jira-linking on November 06, 2023 at 05:08 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2867


---

## Reviews

### Review by @MichaelMraka - Approved on November 07, 2023 at 08:53 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1337*
