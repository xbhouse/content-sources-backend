---
type: pull_request
number: 1087
title: "HMS-5571: add upload cleanup"
state: merged
author: rverdile
created: 2025-04-22T20:00:29Z
updated: 2025-04-25T19:05:27Z
closed: 2025-04-25T19:05:26Z
merged: 2025-04-25T19:05:26Z
base_branch: main
head_branch: upload-cleanup
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1087
---

# Pull Request #1087: HMS-5571: add upload cleanup

**Author**: @rverdile
**Created**: April 22, 2025 at 08:00 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `upload-cleanup`

## Description

## Summary
Adds upload cleanup to `process-repos` and also as it's own command `go run cmd/external_repos/main.go upload-cleanup`

## Testing steps
1. Create an upload repo
2. Add two RPMs
3. Verify the two RPMs are listed in pulp: `pulp upload list`
4. Change the created_at of the first one: `UPDATE uploads SET created_at = '2025-04-20' WHERE upload_uuid = <upload_uuid>`
5. Run upload cleanup: `go run cmd/external_repos/main.go upload-cleanup`
6. Verify in pulp and the content sources database that the uploads have been removed.
7. Make sure it works with uploads that have been finalized or have not been finalized
8. Make sure it works with process-repos as well



---

## Discussion

### Comment by @jlsherrill on April 22, 2025 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5571

---

## Reviews

### Review by @xbhouse - Approved on April 24, 2025 at 09:01 PM UTC

looks and works great!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1087*
