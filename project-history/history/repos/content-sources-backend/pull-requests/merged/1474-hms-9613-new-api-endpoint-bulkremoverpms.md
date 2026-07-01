---
type: pull_request
number: 1474
title: "HMS-9613: new api endpoint bulkRemoveRpms"
state: merged
author: Starle21
created: 2026-04-20T10:38:38Z
updated: 2026-05-05T15:55:52Z
closed: 2026-05-05T15:55:52Z
merged: 2026-05-05T15:55:52Z
base_branch: main
head_branch: api
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1474
---

# Pull Request #1474: HMS-9613: new api endpoint bulkRemoveRpms

**Author**: @Starle21
**Created**: April 20, 2026 at 10:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `api`

## Description

## Summary
Clients can now remove RPM packages from the latest snapshot of an upload repository.
A new snapshot is automatically created with the remaining packages.
Templates that use the latest snapshot are also updated.

To prevent race conditions, there is logic that blocks enqueuing bulkRemoveRpms task or addUploads task  if any of the relevant tasks are running.

## Testing steps
Create an upload repository and add uploads there.
Get the `repository_config uuid` and `[]rpm uuid` to be removed.
Do a post request to `/repositories/:uuid/rpms/bulk_remove/` with the RPMs to be removed.
Check that a new snapshot is created without the removed RPMs. Also the `down arrow` in the repository table on the frontend should show the number of removed RPMs. 


---

## Discussion

### Comment by @xbhouse on April 20, 2026 at 11:00 AM UTC

https://issues.redhat.com/browse/HMS-9613

### Comment by @xbhouse on April 21, 2026 at 05:22 PM UTC

this might be a good candidate for an integration test too :) could you add a test (or update an existing test) in `test/integration/pulp_upload_test.go` to cover removing an rpm from an upload repo?

### Comment by @Starle21 on April 22, 2026 at 01:27 PM UTC

I also wonder if we might want to call the endpoint something different? Since it is not truly deleting any RPMs from Pulp, it just actually creates a new snapshot with those RPMs removed. So `bulkRemoveRpms` might not be such a good name. :thinking: 
Maybe `removeRpmsAndCreateSnapshot` or `removeRpmsFromLatestSnapshot`? What do you think? 

### Comment by @xbhouse on April 22, 2026 at 06:06 PM UTC

> I also wonder if we might want to call the endpoint something different? Since it is not truly deleting any RPMs from Pulp, it just actually creates a new snapshot with those RPMs removed. So `bulkRemoveRpms` might not be such a good name. 🤔 Maybe `removeRpmsAndCreateSnapshot` or `removeRpmsFromLatestSnapshot`? What do you think?

hmm i think `bulkRemoveRpms` makes sense, the end result is a new snapshot without those rpms and this is probably what users would care about. the description you've added in the API spec explains how that's done, i think that's good enough :) `addUploads` also creates a new snapshot with added rpms, so i'd keep the naming consistent with that

---

## Reviews

### Review by @Starle21 - Commented on April 20, 2026 at 11:21 AM UTC

### Review by @xbhouse - Commented on April 21, 2026 at 04:45 PM UTC

### Review by @xbhouse - Commented on April 21, 2026 at 05:12 PM UTC

### Review by @xbhouse - Commented on April 21, 2026 at 05:45 PM UTC

### Review by @xbhouse - Commented on April 21, 2026 at 06:02 PM UTC

### Review by @Starle21 - Commented on April 22, 2026 at 12:37 PM UTC

### Review by @Starle21 - Commented on April 22, 2026 at 12:43 PM UTC

### Review by @Starle21 - Commented on April 22, 2026 at 12:46 PM UTC

### Review by @xbhouse - Commented on April 22, 2026 at 05:47 PM UTC

### Review by @rverdile - Commented on April 22, 2026 at 09:03 PM UTC

### Review by @rverdile - Commented on April 22, 2026 at 09:09 PM UTC

### Review by @Starle21 - Commented on April 23, 2026 at 11:43 AM UTC

### Review by @xbhouse - Commented on April 23, 2026 at 12:04 PM UTC

### Review by @xbhouse - Commented on April 28, 2026 at 09:47 PM UTC

### Review by @xbhouse - Commented on April 28, 2026 at 09:51 PM UTC

### Review by @Starle21 - Commented on April 30, 2026 at 11:02 AM UTC

### Review by @xbhouse - Commented on May 05, 2026 at 02:37 PM UTC

### Review by @Starle21 - Commented on May 05, 2026 at 03:16 PM UTC

### Review by @xbhouse - Approved on May 05, 2026 at 03:22 PM UTC

this looks great to me, awesome work!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1474*
