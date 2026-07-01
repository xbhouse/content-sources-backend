---
type: pull_request
number: 1393
title: "HMS-9885: turn on snapshot cleanup"
state: merged
author: rverdile
created: 2026-02-16T16:22:00Z
updated: 2026-02-17T17:17:52Z
closed: 2026-02-17T17:17:49Z
merged: 2026-02-17T17:17:49Z
base_branch: main
head_branch: cleanup
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1393
---

# Pull Request #1393: HMS-9885: turn on snapshot cleanup

**Author**: @rverdile
**Created**: February 16, 2026 at 04:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `cleanup`

## Description

## Summary
Re-enables snapshot cleanup in the deployment.yaml Limits the number of cleanup tasks that will be enqueued at once. Increases the frequency of the job to help get through the backlog of outdated snapshots.

Sets the limit to 37 tasks, which is a quarter of the workers and increases the frequency to 4x a day. This means that we'd get through the backlog of snapshot cleanups in around a week without fully saturating the workers with snapshot cleanup tasks.

## Testing steps
1. Use [this](https://gist.github.com/rverdile/8b92f8dc16cf5045cfd4fc523fdd839e) script to create some test outdated snapshots, e.g. `go run cmd/test/create_test_snapshot.go`
2. Run snapshot cleanup with a limit: `go run cmd/external-repos/main.go cleanup --type snapshot --snapshot-clean-batch-size 10`
3. Verify batch size matches number of tasks that are enqueued



---

## Discussion

### Comment by @xbhouse on February 16, 2026 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-9885

---

## Reviews

### Review by @TenSt - Approved on February 17, 2026 at 02:33 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1393*
