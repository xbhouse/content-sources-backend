---
type: pull_request
number: 1285
title: "HMS-9677: Configure Tang for all processes"
state: merged
author: rverdile
created: 2025-11-04T18:59:51Z
updated: 2025-11-05T18:25:07Z
closed: 2025-11-05T18:25:03Z
merged: 2025-11-05T18:25:03Z
base_branch: main
head_branch: configure-tang
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1285
---

# Pull Request #1285: HMS-9677: Configure Tang for all processes

**Author**: @rverdile
**Created**: November 04, 2025 at 06:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `configure-tang`

## Description

## Summary
Seeing this error in stage: `could not set detected OS release: error querying redhat-release package: no tang configuration present`

Tang is only configured for the api process, should be configured for workers as well.

## Testing steps
1. `make compose-clean compose-up`
2. `make repo-imports-rhel9`
3. `make process-repos`
4. `go run cmd/content-sources/main.go consumer`
5. Without this PR, you should see the above error pop up. With this PR, just running the consumer process should also configure Tang and there will be no error.



---

## Discussion

### Comment by @xbhouse on November 04, 2025 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-9677

### Comment by @rverdile on November 04, 2025 at 09:34 PM UTC

/retest

### Comment by @xbhouse on November 05, 2025 at 02:05 PM UTC

/retest

### Comment by @rverdile on November 05, 2025 at 03:38 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on November 04, 2025 at 08:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1285*
