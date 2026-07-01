---
type: pull_request
number: 906
title: "Fixes 5052: refactor jobs and temp/repo association"
state: merged
author: jlsherrill
created: 2024-11-25T14:04:51Z
updated: 2024-11-26T17:31:36Z
closed: 2024-11-26T16:46:36Z
merged: 2024-11-26T16:46:35Z
base_branch: main
head_branch: 5052
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/906
---

# Pull Request #906: Fixes 5052: refactor jobs and temp/repo association

**Author**: @jlsherrill
**Created**: November 25, 2024 at 02:04 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5052`

## Description

## Summary

This pulls 2 refactors out of https://github.com/content-services/content-sources-backend/pull/865
* refactors jobs into a single command  ./cmd/jobs/main.go
 * removes old unused jobs
* refactors the template_repository_configurations association into an official model

## Testing steps

1.  Tests pass
2.  running `go run cmd/jobs/main.go retry_failed_task`  seems to run the appropriate code


---

## Discussion

### Comment by @jlsherrill on November 25, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-5052

---

## Reviews

### Review by @xbhouse - Commented on November 25, 2024 at 07:57 PM UTC

### Review by @jlsherrill - Commented on November 26, 2024 at 12:57 PM UTC

### Review by @jlsherrill - Commented on November 26, 2024 at 12:59 PM UTC

### Review by @xbhouse - Commented on November 26, 2024 at 02:15 PM UTC

### Review by @jlsherrill - Commented on November 26, 2024 at 03:27 PM UTC

### Review by @jlsherrill - Commented on November 26, 2024 at 03:28 PM UTC

### Review by @xbhouse - Approved on November 26, 2024 at 04:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/906*
