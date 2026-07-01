---
type: pull_request
number: 1513
title: "HMS-10486: increase introspect timeout, enqueue once per job"
state: merged
author: xbhouse
created: 2026-06-03T19:19:00Z
updated: 2026-06-11T12:26:47Z
closed: 2026-06-11T12:26:47Z
merged: 2026-06-11T12:26:46Z
base_branch: main
head_branch: 10486
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1513
---

# Pull Request #1513: HMS-10486: increase introspect timeout, enqueue once per job

**Author**: @xbhouse
**Created**: June 03, 2026 at 07:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `10486`

## Description

## Summary

- Increases CDN request timeout during introspection from 6 to 10 minutes. In stage, some larger repos are getting close to hitting the 6 minute limit:
```
{"level":"debug","task_type":"introspect","task_id":"c965a698-1d90-483d-b632-e1dcfeb85b4b","request_id":"","url":"https://cdn.redhat.com/content/eus/rhel9/9.8/aarch64/baseos/os/","phase":"parse_packages","duration":331.864419797,"time":"2026-06-01T14:05:39Z","severity":"debug","message":"introspection phase parse_packages completed"}
```
-  Adjusts the `process-repos` job and snapshot command to only enqueue one introspect task per repo and prevent new introspect or snapshot tasks if any are already active for a repo
- Adds an optional `--introspect` flag to the snapshot command and updates `make repos-minimal` to use this
- Fixes a couple methods used during introspection to return proper values

## Testing steps

1. Run `make repos-import && make process-repos`. All introspect tasks should succeed and there should only be one per repo
2. Running `make process-repos` twice in a row shouldn't add multiple introspects for a repo
2. `make repos-minimal` should still introspect and snapshot the "small" repo (RHEL9 ARM codeready)

---

## Discussion

### Comment by @xbhouse on June 03, 2026 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-10486

### Comment by @jlsherrill on June 03, 2026 at 09:26 PM UTC

out of curiosity, do you know which files are taking so long to download?  Is it the larger metadata files?

### Comment by @xbhouse on June 04, 2026 at 12:24 PM UTC

> out of curiosity, do you know which files are taking so long to download? Is it the larger metadata files?

@jlsherrill yea it's mostly the primary, processing the packages for RHEL baseos repos takes the longest usually. for RHEL8, the module streams seem to take quite a bit to process too, the longest i've seen so far for those is about 90 seconds. i'm guessing it takes longer when multiple baseos repos get updates at the same time

### Comment by @rverdile on June 08, 2026 at 02:52 PM UTC

> Adjusts the process-repos job to only enqueue one introspect task per repo and prevent new ones if any are active

wanna change this for snapshots too? i.e. don't enqueue a new snapshot task for a repo when one is already active

### Comment by @xbhouse on June 10, 2026 at 05:53 PM UTC

/retest

### Comment by @rverdile on June 10, 2026 at 06:55 PM UTC

/retest

---

## Reviews

### Review by @Starle21 - Commented on June 09, 2026 at 07:17 AM UTC

### Review by @Starle21 - Commented on June 09, 2026 at 07:23 AM UTC

Testing locally, I confirm that `make process-repos` does not add subsequent tasks if there are still active tasks 👍 

### Review by @xbhouse - Commented on June 09, 2026 at 06:10 PM UTC

### Review by @swadeley - Commented on June 10, 2026 at 10:37 AM UTC

### Review by @Starle21 - Commented on June 10, 2026 at 03:05 PM UTC

### Review by @rverdile - Commented on June 10, 2026 at 04:09 PM UTC

### Review by @xbhouse - Commented on June 10, 2026 at 05:14 PM UTC

### Review by @rverdile - Approved on June 10, 2026 at 05:48 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1513*
