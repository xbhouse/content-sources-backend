---
type: pull_request
number: 371
title: "Fixes 1968: Update RepositoryResponse with snapshot taskID"
state: merged
author: Andrewgdewar
created: 2023-08-23T20:11:26Z
updated: 2023-09-04T13:39:15Z
closed: 2023-09-04T13:39:14Z
merged: 2023-09-04T13:39:14Z
base_branch: main
head_branch: HMS-1968
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/371
---

# Pull Request #371: Fixes 1968: Update RepositoryResponse with snapshot taskID

**Author**: @Andrewgdewar
**Created**: August 23, 2023 at 08:11 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1968`

## Description

## Summary

Add a new parameter "last_snapshot_task_uuid" to the RepositoryResponse (repositoryList/Fetch/Edit responses)

When a repository enques a task to snapshot, the api will now return the above parameter.
<img width="475" alt="Screen Shot 2023-08-23 at 2 11 47 PM" src="https://github.com/content-services/content-sources-backend/assets/38083295/ad3ae1d4-16da-4bc2-b137-e83be78416d3">

This UUID can be confirmed by checking the taskInfo in the Admin tasks view under the ViewDetails Modal  > UUID (for the associated task of course).

<img width="416" alt="Screen Shot 2023-08-23 at 2 12 07 PM" src="https://github.com/content-services/content-sources-backend/assets/38083295/d4e71890-432e-48ea-b577-9a3f69f86b9c">

## Suggested Testing steps
- Create or Edit a repository with snapshotting enabled, check the response in the network for the "last_snapshot_task_uuid". 
- Copy the ID from the network tab above and check it against the  in the tasksInfo modal in the associated task.

---

## Discussion

### Comment by @jlsherrill on August 23, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-1968

### Comment by @swadeley on August 29, 2023 at 09:54 AM UTC

Hi @Andrewgdewar 

Please rebase

thank you

### Comment by @swadeley on August 29, 2023 at 09:57 AM UTC

> Which part of the UI is for? Specifically, where are we going to use the last snapshot task ID?

I think this is just adding "last_snapshot_task_uuid" to the API response and not adding it to the UI as its already displayed in the UI.

### Comment by @Andrewgdewar on August 31, 2023 at 09:15 PM UTC

> 

I'll walk through this with you tomorrow, good catch.

### Comment by @Andrewgdewar on September 01, 2023 at 09:19 PM UTC

Pushed the fix!

### Comment by @swadeley on September 04, 2023 at 12:57 PM UTC

Hi

I built new API doc. Ran all UI and API tests.

 Created repo. 
In shell I listed repo to see last_snapshot_task_uuid

```
Out[1]: 
{'data': [{'account_id': '0369233',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2023-09-04T11:23:09Z',
           'last_snapshot': {'added_counts': {'rpm.package': 1},
                             'content_counts': {'rpm.package': 1},
                             'created_at': '2023-09-04T11:23:26.829725Z',
                             'removed_counts': {},
                             'repository_path': ''},
           'last_snapshot_task_uuid': '25743c9b-1547-4e07-9c55-18bec6ee5a21',
```

I confirm that same UUID is visible in Task details UI.:



![Screenshot 2023-09-04 at 13-50-00 Repositories console redhat com](https://github.com/content-services/content-sources-backend/assets/11247237/14f179f6-9ccc-465b-8d48-4f387148c184)







---

## Reviews

### Review by @rverdile - Commented on August 24, 2023 at 08:19 PM UTC

### Review by @rverdile - Commented on August 24, 2023 at 08:25 PM UTC

### Review by @rverdile - Commented on August 24, 2023 at 08:28 PM UTC

### Review by @Andrewgdewar - Commented on August 24, 2023 at 08:44 PM UTC

### Review by @Andrewgdewar - Commented on August 24, 2023 at 08:50 PM UTC

### Review by @Andrewgdewar - Commented on August 24, 2023 at 08:51 PM UTC

### Review by @rverdile - Commented on August 24, 2023 at 09:38 PM UTC

### Review by @rverdile - Commented on August 24, 2023 at 09:50 PM UTC

### Review by @rverdile - Commented on August 24, 2023 at 09:51 PM UTC

Which part of the UI is for? Specifically, where are we going to use the last snapshot task ID?

### Review by @Andrewgdewar - Commented on August 25, 2023 at 03:59 PM UTC

### Review by @Andrewgdewar - Commented on August 25, 2023 at 04:06 PM UTC

### Review by @rverdile - Commented on August 29, 2023 at 03:54 PM UTC

### Review by @rverdile - Commented on August 29, 2023 at 03:58 PM UTC

### Review by @rverdile - Commented on August 29, 2023 at 05:27 PM UTC

### Review by @rverdile - Commented on August 29, 2023 at 05:40 PM UTC

### Review by @Andrewgdewar - Commented on August 30, 2023 at 06:19 PM UTC

### Review by @rverdile - Commented on August 31, 2023 at 03:24 PM UTC

When I create a repository, I don't see the `last_snapshot_task_uuid`. And when I update a repository, I see the 2nd to last uuid, not the last uuid.

I suspect the response in the handler needs to be updated. And I would add a test around that too.

### Review by @swadeley - Approved on September 04, 2023 at 01:39 PM UTC

lgtm now

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/371*
