---
type: pull_request
number: 700
title: "Fixes 4256: snapshot every hour"
state: merged
author: jlsherrill
created: 2024-06-10T20:52:55Z
updated: 2024-06-13T15:23:39Z
closed: 2024-06-13T15:23:36Z
merged: 2024-06-13T15:23:36Z
base_branch: main
head_branch: 4256
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/700
---

# Pull Request #700: Fixes 4256: snapshot every hour

**Author**: @jlsherrill
**Created**: June 10, 2024 at 08:52 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4256`

## Description

## Summary
start snapshotting every hour.  On each hourly run:
* Snapshot at least # of repos/24 on each run:
  * snapshots all repos that have not been snapshotted in the last 24 hours
  * if this amount is less than # of repos/24, then grab more repos to snapshot up to that number prioritizing those that have been snapshotted last.

## Testing steps

* on a fresh db:  make repos-import
* `go run cmd/external-repos/main.go nightly-jobs`
* wait until all repos are done snapshotting
* create one or more custom repos

from a `make db-cli-connect` console run:

```
select count(queued_at), max(name), tasks.repository_uuid from tasks inner join repository_configurations rc on rc.repository_uuid = tasks.repository_uuid  where type = 'snapshot'  group by tasks.repository_uuid ;
```
this will show you the number of snapshot tasks per repository.  

Then run:

 `go run cmd/external-repos/main.go nightly-jobs  5`

re-run the query, and on every cmd run you should see (# of repos/5) + 1 snapshots run.  You shouldn't see any counts go more than 1 more than any other repository  (they should all slowly go to '2' and then once they are all '2' you will see some go to '3' tasks)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 10, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4256

### Comment by @jlsherrill on June 10, 2024 at 09:01 PM UTC

/retest

### Comment by @swadeley on June 11, 2024 at 12:41 PM UTC

/retest

### Comment by @jlsherrill on June 11, 2024 at 03:31 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on June 11, 2024 at 06:54 PM UTC

### Review by @rverdile - Commented on June 11, 2024 at 06:54 PM UTC

### Review by @jlsherrill - Commented on June 12, 2024 at 01:13 PM UTC

### Review by @rverdile - Approved on June 12, 2024 at 07:19 PM UTC

works and looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/700*
