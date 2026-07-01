---
type: pull_request
number: 836
title: "Fixes 4741: associate snapshots to templates directly"
state: merged
author: xbhouse
created: 2024-10-01T14:51:54Z
updated: 2024-10-21T12:20:24Z
closed: 2024-10-21T12:20:24Z
merged: 2024-10-21T12:20:24Z
base_branch: main
head_branch: 4741
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/836
---

# Pull Request #836: Fixes 4741: associate snapshots to templates directly

**Author**: @xbhouse
**Created**: October 01, 2024 at 02:51 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4741`

## Description

## Summary

- Associates snapshots to templates by adding snapshot_uuid to the templates_repository_configurations table
- Migrates any existing templates' snapshots to this table
- Simplifies the query for listing a template's snapshots to use the template-snapshot association
- Checks that repos added to a template have a snapshot and rejects request if they do not

## Testing steps

To test the migration:

1. In main, create 1 custom repo and 1 RH repo
2. Change the URL of the custom repo so you have 2 snapshots, and update the snapshot's created_at dates to be on different days
3. Change the URL one more time to get another snapshot and don't update this date
4. Create 3 templates: 1 with a date set to be between the dates of the snapshots, 1 with a date before the snapshots, and 1 using latest content
5. Check out this PR and run the migration. Verify the correct snapshot uuids were added to the templates_repository_configurations table (you can either check the table directly or check via the `GET /templates/:uuid/snapshots/` endpoint)

To test a template's snapshots are updated correctly:

1. Create a custom repo and a RH repo, let them snapshot, and create a template with these repos with use_latest set to true
2. Listing the template's snapshots (`GET /templates/:uuid/snapshots/`) should still work as expected
3. Add / remove one of the template's repos, listing the template's snapshots show updated snapshots
4. Trigger another snapshot of one of the repos in template. Once this task completes you should see that latest snapshot when listing the template's snapshots


---

## Discussion

### Comment by @jlsherrill on October 07, 2024 at 05:36 PM UTC

https://issues.redhat.com/browse/HMS-4741

### Comment by @jlsherrill on October 18, 2024 at 05:48 PM UTC

overall this seemed to work fine.  I did see one error though if i:

1.  create a repo
2. create a template and add the repo
3. delete the repo

i get an error:
```
1:46PM WRN [Finished Task] task failed error="ERROR: update or delete on table \"snapshots\" violates foreign key constraint \"fk_templates_repository_configurations_snapshots\" on table \"templates_repository_configurations\" (SQLSTATE 23503)" request_id=4c0220f1-94d8-44cf-819d-6ef213c4c11c task_id=5c876540-ea47-4703-813e-bcd312e7ff87 task_type=delete-repository-snapshots
```




### Comment by @xbhouse on October 18, 2024 at 07:10 PM UTC

> overall this seemed to work fine. I did see one error though if i:
> 
> 1. create a repo
> 2. create a template and add the repo
> 3. delete the repo
> 
> i get an error:
> 
> ```
> 1:46PM WRN [Finished Task] task failed error="ERROR: update or delete on table \"snapshots\" violates foreign key constraint \"fk_templates_repository_configurations_snapshots\" on table \"templates_repository_configurations\" (SQLSTATE 23503)" request_id=4c0220f1-94d8-44cf-819d-6ef213c4c11c task_id=5c876540-ea47-4703-813e-bcd312e7ff87 task_type=delete-repository-snapshots
> ```

ah good catch. because we're restricting deletion we'll need to delete any rows from template_repo_configs with that repo's snapshot uuids before the repo is deleted. then the snapshot integration test updates are probably not necessary. fixing!

---

## Reviews

### Review by @jlsherrill - Commented on October 07, 2024 at 08:28 PM UTC

### Review by @jlsherrill - Commented on October 07, 2024 at 08:30 PM UTC

### Review by @jlsherrill - Commented on October 07, 2024 at 08:39 PM UTC

### Review by @xbhouse - Commented on October 08, 2024 at 02:31 PM UTC

### Review by @xbhouse - Commented on October 08, 2024 at 02:37 PM UTC

### Review by @rverdile - Commented on October 08, 2024 at 03:02 PM UTC

### Review by @jlsherrill - Commented on October 14, 2024 at 03:27 PM UTC

### Review by @xbhouse - Commented on October 14, 2024 at 04:29 PM UTC

### Review by @xbhouse - Commented on October 16, 2024 at 04:03 PM UTC

### Review by @xbhouse - Commented on October 16, 2024 at 04:05 PM UTC

### Review by @jlsherrill - Commented on October 16, 2024 at 08:01 PM UTC

### Review by @xbhouse - Commented on October 17, 2024 at 01:52 PM UTC

### Review by @jlsherrill - Approved on October 18, 2024 at 08:44 PM UTC

ACK!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/836*
