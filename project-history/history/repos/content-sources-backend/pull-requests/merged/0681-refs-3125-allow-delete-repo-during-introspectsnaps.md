---
type: pull_request
number: 681
title: "Refs 3125: allow delete repo during introspect/snapshot"
state: merged
author: rverdile
created: 2024-05-28T20:24:32Z
updated: 2024-05-30T12:55:02Z
closed: 2024-05-30T09:58:37Z
merged: 2024-05-30T09:58:37Z
base_branch: main
head_branch: delete-during-snapshot
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/681
---

# Pull Request #681: Refs 3125: allow delete repo during introspect/snapshot

**Author**: @rverdile
**Created**: May 28, 2024 at 08:24 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `delete-during-snapshot`

## Description

## Summary
This changes the behavior of the delete and bulk delelete repository API endpoints to allow deletion if a snapshot is in progress. If a snapshot is in progress, the snapshot task is canceled. The introspect task is also canceled.

The corresponding frontend PR is here: https://github.com/content-services/content-sources-frontend/pull/269

## Testing steps
1. Create or bulk create repositories. It's easiest to do this in the UI.
2. While the status is pending, delete the repositories.
3. You should find that the running or queued (introspect or snapshot) tasks are canceled, and the repositories are deleted.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 28, 2024 at 08:31 PM UTC

https://issues.redhat.com/browse/HMS-3125

### Comment by @jlsherrill on May 28, 2024 at 08:31 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @xbhouse on May 29, 2024 at 05:52 PM UTC

if i add a repository without snapshots enabled, and delete the repo during introspection, i'm seeing the delete-repository-snapshots task still be added to the queue and completed. is this expected?

### Comment by @xbhouse on May 29, 2024 at 06:00 PM UTC

other than that one question, this looks good and works well! 

### Comment by @rverdile on May 29, 2024 at 06:56 PM UTC

yep, that task runs whether there's snapshots or not

### Comment by @rverdile on May 29, 2024 at 07:08 PM UTC

If someone enabled snapshots, created a repo, then disabled snapshots, we'd still want to clean up the snapshots on deletion.

### Comment by @xbhouse on May 29, 2024 at 07:53 PM UTC

got it, makes sense! thanks for explaining

### Comment by @swadeley on May 30, 2024 at 09:58 AM UTC

It works, yay, thank you.

---

## Reviews

### Review by @xbhouse - Approved on May 29, 2024 at 07:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/681*
