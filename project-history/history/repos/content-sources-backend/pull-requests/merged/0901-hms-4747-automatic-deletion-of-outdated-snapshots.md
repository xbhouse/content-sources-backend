---
type: pull_request
number: 901
title: "HMS-4747: automatic deletion of outdated snapshots"
state: merged
author: dominikvagner
created: 2024-11-21T10:39:02Z
updated: 2024-12-19T14:25:50Z
closed: 2024-12-19T14:25:50Z
merged: 2024-12-19T14:25:50Z
base_branch: main
head_branch: 4747
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/901
---

# Pull Request #901: HMS-4747: automatic deletion of outdated snapshots

**Author**: @dominikvagner
**Created**: November 21, 2024 at 10:39 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `4747`

## Description

## Summary
This PR adds automatic deletion of outdated snapshots to the `nightly-jobs` command. The date after which the snapshots should be deleted is settable in the config.
Also adds a list of snapshots which are going to be deleted in the next 14 days to the fetch/list template endpoints.

## Testing steps
1. Create a repository and let it snapshot.
2. Create a template that uses the repo and is using the latest snap.
3. Connect to the DB CLI `make db-cli-connect` and set the snapshot date to the past.
    - `UPDATE snapshots SET created_at = (NOW() - CAST('400 days' AS INTERVAL)) WHERE uuid = '$SNAPSHOT_UUID';`
4. Run nightly jobs (`go run cmd/external-repos/main.go nightly-jobs`) and verify that there wasn't a "delete-snapshots" task enqueued.
5. Verify the outdated snaphot is listed in both the "snapshots" and "to_be_deleted_snapshots" when querying templates.
6. Edit the repo so that it snapshots again and change the template to use the old one.
7. Run the nightly jobs again and verify the snapshot has been deleted and the template updated to use the other one.

---

## Discussion

### Comment by @jlsherrill on November 28, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4747

### Comment by @swadeley on December 02, 2024 at 01:27 AM UTC

/retest

### Comment by @swadeley on December 02, 2024 at 01:55 PM UTC

/retest

### Comment by @swadeley on December 03, 2024 at 11:32 AM UTC

Hi, I am stuck with this:
```
E               iqe_content_sources_api.exceptions.ApiTypeError: Invalid type for variable 'to_be_deleted_snapshots'. Required value type is list and passed type was NoneType at ['received_data']['to_be_deleted_snapshots']
```

### Comment by @swadeley on December 04, 2024 at 12:30 PM UTC

> Hi, I am stuck with this:
> 
> ```
> E               iqe_content_sources_api.exceptions.ApiTypeError: Invalid type for variable 'to_be_deleted_snapshots'. Required value type is list and passed type was NoneType at ['received_data']['to_be_deleted_snapshots']
> ```
 
Hi, thanks for the update. IQE tests now pass.

### Comment by @xbhouse on December 10, 2024 at 07:26 PM UTC

i'm seeing this error when running nightly-jobs if there is a template using an outdated snapshot:

```
"error deleting rpm repository versions: 400 Bad Request: [\"The repository version cannot be deleted because it (or its publications) are currently being used to distribute content. Please update the necessary distributions first.\"]"
```

it seems like maybe the template isn't being updated to the next snapshot before the snapshot is deleted? though on the next requeue it does seem like the distribution was removed (getting a 404 when trying to view the content url and seeing resource not found errors), but the snapshot was removed from the template and still exists in our db. 

i may just be missing something when testing 😄 to reproduce, i'm using these steps:

1. create a template (set to the current date) with a RH repo and a custom repo
2. update the custom repo's snapshot's created_at field to be more than 365 days in the past 
3. edit the custom repo's url to get a new snapshot 
4. run nightly-jobs

### Comment by @xbhouse on December 16, 2024 at 09:39 PM UTC

> i'm seeing this error when running nightly-jobs if there is a template using an outdated snapshot:
> 
> ```
> "error deleting rpm repository versions: 400 Bad Request: [\"The repository version cannot be deleted because it (or its publications) are currently being used to distribute content. Please update the necessary distributions first.\"]"
> ```
> 
> it seems like maybe the template isn't being updated to the next snapshot before the snapshot is deleted? though on the next requeue it does seem like the distribution was removed (getting a 404 when trying to view the content url and seeing resource not found errors), but the snapshot was removed from the template and still exists in our db.
> 
> i may just be missing something when testing 😄 to reproduce, i'm using these steps:
> 
> 1. create a template (set to the current date) with a RH repo and a custom repo
> 2. update the custom repo's snapshot's created_at field to be more than 365 days in the past
> 3. edit the custom repo's url to get a new snapshot
> 4. run nightly-jobs

your latest commit seems to have fixed this! 🎉 

### Comment by @jlsherrill on December 17, 2024 at 10:13 PM UTC

(using your PR to test some automation real quick)

### Comment by @jlsherrill on December 17, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-HMS-4747

### Comment by @jlsherrill on December 18, 2024 at 07:05 PM UTC

https://issues.redhat.com/browse/Fixes HMS-4747

---

## Reviews

### Review by @jlsherrill - Commented on November 25, 2024 at 08:48 PM UTC

### Review by @xbhouse - Commented on December 16, 2024 at 09:36 PM UTC

### Review by @dominikvagner - Commented on December 17, 2024 at 09:36 AM UTC

### Review by @xbhouse - Approved on December 17, 2024 at 06:31 PM UTC

this looks good to me! @jlsherrill may want to take another look?

### Review by @jlsherrill - Commented on December 17, 2024 at 07:04 PM UTC

### Review by @dominikvagner - Commented on December 18, 2024 at 07:56 AM UTC

### Review by @jlsherrill - Approved on December 18, 2024 at 09:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/901*
