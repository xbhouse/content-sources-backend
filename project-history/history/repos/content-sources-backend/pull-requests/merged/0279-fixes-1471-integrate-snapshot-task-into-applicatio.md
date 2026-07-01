---
type: pull_request
number: 279
title: "Fixes 1471: Integrate snapshot task into application"
state: merged
author: jlsherrill
created: 2023-05-23T17:48:54Z
updated: 2023-06-03T15:30:29Z
closed: 2023-06-03T15:12:59Z
merged: 2023-06-03T15:12:59Z
base_branch: main
head_branch: 1471
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/279
---

# Pull Request #279: Fixes 1471: Integrate snapshot task into application

**Author**: @jlsherrill
**Created**: May 23, 2023 at 05:48 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1471`

## Description

## Summary

This adds snapshotting to repositories, integrated with the api and tasking system.  When creating a repository, if you pass snapshot=true, a task is spawned to snapshot the repository.  This snapshot task should save information within the payload such as sync task href, publication task href, etc..., such that if the worker dies, the task should be able to be resumed.

## Testing steps

Run a create or bulk_create request as normal, but add the parameter 'snapshot=true'.

Try killing the worker (with ctrl-c or kill -9) and restarting it.  In all cases you should be able to see the task resumed and a snapshot created.  For now you'll have to check the db to see if snapshot got created:

```select distribution_path from snapshots;```

you can then curl the pulp distribution path:

```curl  localhost:8080/pulp/content/73e2d9fa-387d-4df1-ac58-fe898aebb30f/f4e2f625-4cba-4d5c-9ea5-3de353bb72f1/```

For QE:

We'll need to update automation to support the snapshot parameter, but we can't really test snapshots yet, as there is no way to monitor the snapshot or confirm its created.

---

## Discussion

### Comment by @jlsherrill on May 23, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-1471

### Comment by @swadeley on June 01, 2023 at 07:05 AM UTC

Hi, I have tested this and updated automation:

[Add snapshot bool (!180) · Merge requests · insights-qe / iqe-content-sources-plugin · GitLab](https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/merge_requests/180)


### Comment by @jlsherrill on June 01, 2023 at 06:52 PM UTC

i'll want to disable the pulpserver in stage and prod before we merge this, openeing an MR to do that.

EDIT:  This is done

---

## Reviews

### Review by @swadeley - Commented on May 30, 2023 at 05:13 PM UTC

### Review by @Andrewgdewar - Dismissed on May 30, 2023 at 06:51 PM UTC

Ack for functionality. 

### Review by @rverdile - Commented on May 30, 2023 at 07:08 PM UTC

### Review by @rverdile - Commented on May 31, 2023 at 07:23 PM UTC

### Review by @rverdile - Commented on May 31, 2023 at 08:15 PM UTC

those are my only two comments!

### Review by @jlsherrill - Commented on June 01, 2023 at 06:47 PM UTC

### Review by @jlsherrill - Commented on June 01, 2023 at 06:49 PM UTC

### Review by @rverdile - Approved on June 02, 2023 at 04:23 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/279*
