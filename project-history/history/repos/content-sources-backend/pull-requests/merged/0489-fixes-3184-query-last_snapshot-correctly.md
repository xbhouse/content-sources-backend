---
type: pull_request
number: 489
title: "Fixes 3184: query last_snapshot correctly"
state: merged
author: jlsherrill
created: 2023-11-28T17:18:59Z
updated: 2023-11-30T16:00:29Z
closed: 2023-11-30T15:48:36Z
merged: 2023-11-30T15:48:36Z
base_branch: main
head_branch: 3184
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/489
---

# Pull Request #489: Fixes 3184: query last_snapshot correctly

**Author**: @jlsherrill
**Created**: November 28, 2023 at 05:18 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3184`

## Description

## Summary

We weren't specifying the association foreign key correctly, so we were getting a 'random' snapshot, and not the latest

## Testing steps

Generate a repo with multiple snapshots, view the json when fetching a repo and make sure that 'last_snapshot_uuid' matches 'last_snapshot.uuid'

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 28, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-3184

### Comment by @swadeley on November 29, 2023 at 01:57 PM UTC

Hi, I could not get snapshotting to work in the usual way of manipulating the backends of the repo. Will try again.

### Comment by @jlsherrill on November 30, 2023 at 03:48 PM UTC

I'm going to go ahead and merge this, as we need it to run a repair operation in stage.  Please reach out if you find an issue with it

---

## Reviews

### Review by @rverdile - Approved on November 28, 2023 at 09:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/489*
