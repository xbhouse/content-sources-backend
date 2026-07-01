---
type: pull_request
number: 824
title: "Fixes 4735: snapshot repos even if last task was deleted"
state: merged
author: jlsherrill
created: 2024-09-17T14:46:32Z
updated: 2024-09-18T16:40:38Z
closed: 2024-09-18T16:40:35Z
merged: 2024-09-18T16:40:35Z
base_branch: main
head_branch: 4735
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/824
---

# Pull Request #824: Fixes 4735: snapshot repos even if last task was deleted

**Author**: @jlsherrill
**Created**: September 17, 2024 at 02:46 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4735`

## Description

## Summary

If a repo's last_snapshot_task got cleaned up the repo was never considered for nightly snapshotting because:
1)  the query that queries for repos to snapshot assumed that either the last_snapshot_task_uuid would be null, or it would exist and have a finished date
2) we have code to snapshot repos even if not needed yet to help spread out the load, but apparently sorting in ascending order puts NULLs last, so they would never have been selected

## Testing steps
On main:
1.  Create a few repos, or import some red hat ones and trigger them to sync with ```go run cmd/external-repos/main.go nightly-jobs```
2.   After the repos have synced, run:
 ```go run cmd/external-repos/main.go nightly-jobs 6```

This should trigger at least 1 snapshot task:
```

```
3.  Now "clean" the task:

```
make db-cli-connect

delete from tasks;
```

4.  Try to run the nightly jobs again:
 ```go run cmd/external-repos/main.go nightly-jobs 6```

You won't see any of these snapshot task get enqueued 

5.  Checkout this pr and ```make db-migrate-up```
6. Re-run the nightly jobs:

 ```go run cmd/external-repos/main.go nightly-jobs 6```

You should see their snapshot task be enqueued.

---

## Discussion

### Comment by @jlsherrill on September 17, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4735

---

## Reviews

### Review by @jlsherrill - Commented on September 17, 2024 at 05:54 PM UTC

### Review by @rverdile - Approved on September 18, 2024 at 03:10 PM UTC

looks good, just needs rebase!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/824*
