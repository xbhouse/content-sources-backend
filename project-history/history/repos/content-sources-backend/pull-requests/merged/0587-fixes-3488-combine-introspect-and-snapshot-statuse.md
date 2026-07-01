---
type: pull_request
number: 587
title: "Fixes 3488: combine introspect and snapshot statuses"
state: merged
author: xbhouse
created: 2024-02-26T19:50:26Z
updated: 2024-03-26T23:49:08Z
closed: 2024-03-26T23:48:06Z
merged: 2024-03-26T23:48:06Z
base_branch: main
head_branch: combine-statuses
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/587
---

# Pull Request #587: Fixes 3488: combine introspect and snapshot statuses

**Author**: @xbhouse
**Created**: February 26, 2024 at 07:50 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `combine-statuses`

## Description

## Summary

- Combines introspection and snapshot statuses into one field returned in the RepositoryResponse. Chart of combination logic can be found [here](https://issues.redhat.com/browse/HMS-3488)
- `status` removed from the db and replaced with `last_introspection_status`
- Handles filtering on the combined status
- Sorting is enabled on the `last_introspection_status` only. PR with the UI change to disable sorting on the new combined status can be found [here](https://github.com/content-services/content-sources-frontend/pull/219)

## Testing steps

- Introspect and snapshot a repository, the `status` returned in the response should reflect the combined status
- The different combinations can be difficult to simulate, one way to do this is to manually change the snapshot task status (in the tasks table) and the `last_introspection_status` (in the repositories table) in the db

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 26, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3488

### Comment by @swadeley on March 13, 2024 at 08:39 AM UTC

/retest

### Comment by @swadeley on March 14, 2024 at 02:29 PM UTC

/retest

### Comment by @swadeley on March 19, 2024 at 07:46 PM UTC

/retest

### Comment by @jlsherrill on March 19, 2024 at 08:40 PM UTC

overall looks good, just one small comment

### Comment by @jlsherrill on March 22, 2024 at 02:37 PM UTC

rebased against master

### Comment by @Andrewgdewar on March 25, 2024 at 07:56 PM UTC

I believe that you may need to update the API to handle manual snapshots.
See my comment about the front-end PR [here](https://github.com/content-services/content-sources-frontend/pull/234#issuecomment-2018637501).

### Comment by @mshriver on March 26, 2024 at 07:25 PM UTC

/retest

### Comment by @mshriver on March 26, 2024 at 11:49 PM UTC

We've got some suspicious test failures in ephemeral right now with this straightforward update, but its time to move to the stage environment with this and continue investigation there. Thanks for your patience @xbhouse! 

---

## Reviews

### Review by @jlsherrill - Commented on March 04, 2024 at 04:26 PM UTC

### Review by @jlsherrill - Commented on March 04, 2024 at 04:47 PM UTC

### Review by @jlsherrill - Commented on March 04, 2024 at 04:48 PM UTC

### Review by @xbhouse - Commented on March 04, 2024 at 04:55 PM UTC

### Review by @xbhouse - Commented on March 05, 2024 at 07:02 PM UTC

### Review by @jlsherrill - Commented on March 07, 2024 at 04:25 PM UTC

### Review by @jlsherrill - Commented on March 07, 2024 at 04:25 PM UTC

### Review by @xbhouse - Commented on March 07, 2024 at 04:49 PM UTC

### Review by @xbhouse - Commented on March 12, 2024 at 02:54 PM UTC

### Review by @xbhouse - Commented on March 13, 2024 at 03:50 PM UTC

### Review by @jlsherrill - Commented on March 13, 2024 at 04:46 PM UTC

### Review by @jlsherrill - Commented on March 19, 2024 at 07:46 PM UTC

### Review by @xbhouse - Commented on March 20, 2024 at 01:11 PM UTC

### Review by @xbhouse - Commented on March 20, 2024 at 01:17 PM UTC

### Review by @jlsherrill - Approved on March 20, 2024 at 08:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/587*
