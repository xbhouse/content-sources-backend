---
type: pull_request
number: 645
title: "Refs 2422: Add pending task metrics"
state: merged
author: jlsherrill
created: 2024-04-23T15:57:19Z
updated: 2024-05-07T13:53:45Z
closed: 2024-05-07T13:53:42Z
merged: 2024-05-07T13:53:42Z
base_branch: main
head_branch: metrics
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/645
---

# Pull Request #645: Refs 2422: Add pending task metrics

**Author**: @jlsherrill
**Created**: April 23, 2024 at 03:57 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `metrics`

## Description

## Summary
    Adds new metrics:
    * Average queue time of pending task
    * Length in queue for the oldest pending task
    * Number of pending tasks
    Also:
    * Increases default metric calculations to 30 seconds from 5
    * makes interval for metric calculation configurable
    * adds log level override for metrics so you can set the rest of the app to trace while metrics are set to debugging

## Testing steps

on a new db:

```
make run
```
in another tab:
```
make repos-import
go run cmd/external-repos/main.go nightly-jobs
```
This will kick off a ton of tasks, you can monitor with:
```
 curl localhost:9000/metrics | grep task_stats
```

This will update every ~30 seconds.  

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 23, 2024 at 03:57 PM UTC

note this is currently built ontop of https://github.com/content-services/content-sources-backend/pull/637

will rebase once merged.

### Comment by @jlsherrill on April 23, 2024 at 07:58 PM UTC

https://issues.redhat.com/browse/HMS-2422

### Comment by @jlsherrill on April 24, 2024 at 12:55 PM UTC

/retest

### Comment by @swadeley on April 26, 2024 at 02:52 PM UTC

/retest

### Comment by @swadeley on April 26, 2024 at 04:31 PM UTC

/retest

### Comment by @swadeley on April 26, 2024 at 06:32 PM UTC

/retest

### Comment by @jlsherrill on May 01, 2024 at 08:25 PM UTC

added a test to check for queued_at updates.

Also i realized that the 'number of days' until the cdn cert expires was only calculated at startup, so i changed it to be calculated as part of the go routine that updates the metrics.  At startup it was set to zero, so to keep that from firing an alert, i set it to calculate that one metric at startup (and only that metric, because others require the db to be up, which may not actually have happened yet?) 

### Comment by @jlsherrill on May 07, 2024 at 12:21 PM UTC

added

---

## Reviews

### Review by @rverdile - Commented on April 29, 2024 at 01:47 PM UTC

### Review by @rverdile - Commented on May 01, 2024 at 05:04 PM UTC

### Review by @rverdile - Commented on May 02, 2024 at 08:42 PM UTC

### Review by @rverdile - Commented on May 02, 2024 at 08:46 PM UTC

tested and the numbers look right! just the one small comment

### Review by @jlsherrill - Commented on May 07, 2024 at 12:13 PM UTC

### Review by @rverdile - Approved on May 07, 2024 at 01:44 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/645*
