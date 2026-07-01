---
type: pull_request
number: 387
title: "Fixes 1963: nightly sync "
state: merged
author: jlsherrill
created: 2023-09-11T19:17:14Z
updated: 2023-09-27T14:12:13Z
closed: 2023-09-27T14:12:13Z
merged: 2023-09-27T14:12:13Z
base_branch: main
head_branch: HMS-1963-nightly-cron-sync
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/387
---

# Pull Request #387: Fixes 1963: nightly sync 

**Author**: @jlsherrill
**Created**: September 11, 2023 at 07:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-1963-nightly-cron-sync`

## Description


## Summary

Adds repository snapshotting to nightly tasks. When the snapshot feature is enabled in the config, nightly tasks will create a snapshot of all repositories with snapshot = true that haven't been snapshotted in the last 24 hours.

## Testing steps

Enable the snapshot feature in the config and create a repository with snapshotting enabled.
go run cmd/external-repos/main.go nightly-jobs triggers the sync task.
If there is a repository with snapshotting enabled that hasn't been snapshotted in the past 24 hours (including one that has no snapshots), then a snapshot task for that repository will be enqueued. Running the nightly-job will log that the snapshot task was enqueued.


---

## Discussion

### Comment by @jlsherrill on September 11, 2023 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-1963

### Comment by @jlsherrill on September 11, 2023 at 08:47 PM UTC

It should, i'll dig into why not

### Comment by @swadeley on September 14, 2023 at 07:03 AM UTC

/retest

### Comment by @swadeley on September 17, 2023 at 03:49 PM UTC

/retest

### Comment by @swadeley on September 21, 2023 at 01:58 PM UTC

/retest

### Comment by @jlsherrill on September 26, 2023 at 03:25 PM UTC

rebased and added a new parameter: OPTIONS_ALWAYS_RUN_CRON_TASKS.  When set to true: `OPTIONS_ALWAYS_RUN_CRON_TASKS=true`  this will cause all repos to be introspected and snapshotted anytime the cron job runs, ignoring any time constraint normally in place

### Comment by @swadeley on September 27, 2023 at 02:12 PM UTC

Thank you.

---

## Reviews

### Review by @rverdile - Commented on September 11, 2023 at 08:38 PM UTC

### Review by @rverdile - Commented on September 11, 2023 at 08:42 PM UTC

If the snapshot fails, it should try again in 8 hours, right? I'm seeing that only snapshots older than 20 hours are queued.

### Review by @jlsherrill - Commented on September 11, 2023 at 08:46 PM UTC

### Review by @rverdile - Approved on September 11, 2023 at 09:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/387*
