---
type: pull_request
number: 552
title: "Fixes 3492,3371: properly join repo to task"
state: merged
author: jlsherrill
created: 2024-01-29T22:13:59Z
updated: 2024-01-31T20:00:31Z
closed: 2024-01-31T20:00:03Z
merged: 2024-01-31T20:00:03Z
base_branch: main
head_branch: 3492
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/552
---

# Pull Request #552: Fixes 3492,3371: properly join repo to task

**Author**: @jlsherrill
**Created**: January 29, 2024 at 10:13 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3492`

## Description

## Summary

Fixes a query that didn't properly limit joined repo configs to the proper org id.

## Testing steps

Create two repos pointing to the same URL in two different orgs.
Fetch the resulting introspect task for each repo in each org
verify the repo uuid and name matches from the expected org

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 29, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-3492

### Comment by @swadeley on January 30, 2024 at 10:17 AM UTC

Hi

/lgtm


New test running against stage:

```
>               assert task.repository_name == edited_repo.name
E               AssertionError: assert 'RBAC-pass-TwUGNkAE' == 'Snapshot-func-test-xhhzQeXH'
E                 - Snapshot-func-test-xhhzQeXH
E                 + RBAC-pass-TwUGNkAE
```

New snapshot tests passes with this PR branch. 
Thank you.

### Comment by @jlsherrill on January 30, 2024 at 05:27 PM UTC

realized the list tasks api had the same issue, so i updated that

### Comment by @jlsherrill on January 30, 2024 at 08:58 PM UTC

yep, i agree.  I'll update the title and let @swadeley verify that too.  Good catch @rverdile !

### Comment by @jlsherrill on January 30, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-3371

### Comment by @swadeley on January 31, 2024 at 09:02 AM UTC

Hi

tested with latest image. It seems repo name and UUID is missing:

```
In [4]: app.content_sources.rest_client.tasks_api.list_tasks(type='snapshot',limit=8)
2024-01-31 08:58:22.078 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('type', 'snapshot'), ('limit', 8)]
Out[4]: 
{'data': [{'created_at': '2024-01-31T08:50:52Z',
           'ended_at': '2024-01-31T08:51:04Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '8fe8b66b-6d13-4766-a324-c62b5f222bb4'},
          {'created_at': '2024-01-31T08:48:33Z',
           'ended_at': '2024-01-31T08:48:38Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '4d2a2515-ce48-4512-bd68-0b937a7dea3e'},
          {'created_at': '2024-01-31T08:47:19Z',
           'ended_at': '2024-01-31T08:47:29Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '0531675f-34e5-4f33-9f22-12a9ac1e860f'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=8&offset=0&type=snapshot',
           'last': '/api/content-sources/v1/tasks/?limit=8&offset=0&type=snapshot'},
 'meta': {'count': 3, 'limit': 8, 'offset': 0}}


```

The above was while running a CRUD test.


Now, I enabled EPEL 9 with snapshot disabled, then enabled snapshots:

```
Out[13]: 
{'data': [{'created_at': '2024-01-31T09:07:21Z',
           'ended_at': '',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'EPEL 9 Everything x86_64',
           'repository_uuid': '5192af12-8483-425f-b367-7224e5835d2e',
           'status': 'running',
           'type': 'snapshot',
           'uuid': 'fa5fd3f3-4ba3-4de4-8b86-68d41d08f540'},
```
```
In [18]: app.content_sources.rest_client.repositories_api.list_repositories(name='EPEL 9 Everything x86_64')['data'][0]
 'last_introspection_error': '',
 'last_introspection_time': '2024-01-31T09:07:21Z',
 'last_snapshot_task_uuid': 'fa5fd3f3-4ba3-4de4-8b86-68d41d08f540',
 'last_success_introspection_time': '2024-01-31T09:07:21Z',
 'last_update_introspection_time': '2024-01-31T08:00:27Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'EPEL 9 Everything x86_64',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 19870,
 'snapshot': True,
 'status': 'Valid',
 'url': 'https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/',
 'uuid': '5192af12-8483-425f-b367-7224e5835d2e'}

```

UUIDs match in task and repo 

### Comment by @swadeley on January 31, 2024 at 10:42 AM UTC

Hi, I was just single stepping through my updated snapshot test to make sure the asserts were working. 

with this repo:
```
(Pdb) repo
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2024-01-31T10:25:40Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2024-01-31T10:25:52.0968Z',
                   'removed_counts': {},
                   'repository_path': '029e4b9a/1f3beed6-3151-410d-99f7-bf464144753f/6acc9a33-fa71-47b2-99d2-5c04da08d9ca',
                   'url': 'https://<snip>.openshiftapps.com/pulp/content/029e4b9a/1f3beed6-3151-410d-99f7-bf464144753f/6acc9a33-fa71-47b2-99d2-5c04da08d9ca/',
                   'uuid': '5dc0b009-4ba6-4454-8bb2-3afd11c6a123'},
 'last_snapshot_task_uuid': '3ecc3116-9b1e-48a3-bbd9-6ac57dba8de7',
 'last_snapshot_uuid': '5dc0b009-4ba6-4454-8bb2-3afd11c6a123',
 'last_success_introspection_time': '2024-01-31T10:25:40Z',
 'last_update_introspection_time': '2024-01-31T08:47:19Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'Snapshot-func-test-FwIweqOc',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': 'http://stephenw.fedorapeople.org/multirepos/1/repo01/',
 'uuid': '1f3beed6-3151-410d-99f7-bf464144753f'}
```

The task looks good here (UUIDs and name match):

```
(Pdb) task
{'created_at': '2024-01-31T10:25:40Z',
 'ended_at': '2024-01-31T10:25:52Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'Snapshot-func-test-FwIweqOc',
 'repository_uuid': '1f3beed6-3151-410d-99f7-bf464144753f',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '3ecc3116-9b1e-48a3-bbd9-6ac57dba8de7'}
(Pdb)
```
Then, in a shell, I did:

```
In [1]: app.content_sources.rest_client.tasks_api.list_tasks(type='snapshot')

2024-01-31 10:29:11.571 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('type', 'snapshot')]
Out[1]: 
{'data': [{'created_at': '2024-01-31T10:25:40Z',
           'ended_at': '2024-01-31T10:25:52Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'Snapshot-func-test-FwIweqOc',
           'repository_uuid': '1f3beed6-3151-410d-99f7-bf464144753f',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '3ecc3116-9b1e-48a3-bbd9-6ac57dba8de7'},
          {'created_at': '2024-01-31T10:18:51Z',
           'ended_at': '2024-01-31T10:19:05Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'Snapshot-func-test-FwIweqOc',
           'repository_uuid': '1f3beed6-3151-410d-99f7-bf464144753f',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '736c518c-fc7f-4ec3-ae6f-86aaa3813cf3'},
          {'created_at': '2024-01-31T09:07:21Z',
           'ended_at': '2024-01-31T09:08:52Z',
           'error': 'reason: Killed by signal 9..  ',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'failed',
           'type': 'snapshot',
           'uuid': 'fa5fd3f3-4ba3-4de4-8b86-68d41d08f540'},
          {'created_at': '2024-01-31T08:50:52Z',
           'ended_at': '2024-01-31T08:51:04Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'Snapshot-func-test-FwIweqOc',
           'repository_uuid': '1f3beed6-3151-410d-99f7-bf464144753f',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '8fe8b66b-6d13-4766-a324-c62b5f222bb4'},
          {'created_at': '2024-01-31T08:48:33Z',
           'ended_at': '2024-01-31T08:48:38Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '4d2a2515-ce48-4512-bd68-0b937a7dea3e'},
          {'created_at': '2024-01-31T08:47:19Z',
           'ended_at': '2024-01-31T08:47:29Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'Snapshot-func-test-FwIweqOc',
           'repository_uuid': '1f3beed6-3151-410d-99f7-bf464144753f',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '0531675f-34e5-4f33-9f22-12a9ac1e860f'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [2]: 

```

Notice that the task `0531675f-34e5-4f33-9f22-12a9ac1e860f` I pasted earlier, in a previous comment, now has name and repo UUID.

### Comment by @swadeley on January 31, 2024 at 10:52 AM UTC

Hi

I just ran the snapshot test again. In a shell:

```
In [1]: app.content_sources.rest_client.tasks_api.list_tasks(type='snapshot')

2024-01-31 10:47:31.417 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('type', 'snapshot')]
Out[1]: 
{'data': [{'created_at': '2024-01-31T10:44:35Z',
           'ended_at': '2024-01-31T10:44:45Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': 'f886d704-56e7-48b3-83aa-ee765ad72c24'},
          {'created_at': '2024-01-31T10:25:40Z',
           'ended_at': '2024-01-31T10:25:52Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '3ecc3116-9b1e-48a3-bbd9-6ac57dba8de7'},
          {'created_at': '2024-01-31T10:18:51Z',
           'ended_at': '2024-01-31T10:19:05Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '736c518c-fc7f-4ec3-ae6f-86aaa3813cf3'},
          {'created_at': '2024-01-31T09:07:21Z',
           'ended_at': '2024-01-31T09:08:52Z',
           'error': 'reason: Killed by signal 9..  ',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'failed',
           'type': 'snapshot',
           'uuid': 'fa5fd3f3-4ba3-4de4-8b86-68d41d08f540'},
          {'created_at': '2024-01-31T08:50:52Z',
           'ended_at': '2024-01-31T08:51:04Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '8fe8b66b-6d13-4766-a324-c62b5f222bb4'},
          {'created_at': '2024-01-31T08:48:33Z',
           'ended_at': '2024-01-31T08:48:38Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '4d2a2515-ce48-4512-bd68-0b937a7dea3e'},
          {'created_at': '2024-01-31T08:47:19Z',
           'ended_at': '2024-01-31T08:47:29Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': '',
           'repository_uuid': '',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '0531675f-34e5-4f33-9f22-12a9ac1e860f'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot'},
 'meta': {'count': 7, 'limit': 100, 'offset': 0}}

In [2]: 

```

Notice the change in task `0531675f-34e5-4f33-9f22-12a9ac1e860f`

### Comment by @jlsherrill on January 31, 2024 at 04:05 PM UTC

@swadeley is it possible that these repos were deleted?  I think those would be blank if it occured after deletion

### Comment by @swadeley on January 31, 2024 at 04:15 PM UTC

> @swadeley is it possible that these repos were deleted? I think those would be blank if it occured after deletion

Yes, the clean up function would delete them.
That explains why the name and UUID is not there in the last, i.e. third, task list, but not why its not there in the first one but is there in the second one. But its not an issue for this PR. 

I still need to do the user-in-a-different-org test mentioned in comment 0.

### Comment by @swadeley on January 31, 2024 at 07:59 PM UTC

> 
> I still need to do the user-in-a-different-org test mentioned in comment 0.

with same url
       	'url': 'https://stephenw.fedorapeople.org/multirepos/1/repo10/',

I created repos in two accounts:

In [35]: app.user.auth
Out[35]: <Box: {'username': 'jdoe',

```
           'url': 'https://stephenw.fedorapeople.org/multirepos/1/repo10/',
           'uuid': '8c06233b-19c0-4cd1-836c-d230b803cc6e'}],
```

rw_app.user.auth
Out[40]: <Box: {'username': 'ETlVXdeDZCfaVCAc',
```
           'url': 'https://stephenw.fedorapeople.org/multirepos/1/repo10/',
           'uuid': 'c2893a62-f8df-4627-aa16-1090dd284bc1'}],
```


Here is tasks shown by jdoe user:

```
In [32]: app.content_sources.rest_client.tasks_api.list_tasks(status="completed")

Out[32]:
{'data': [{'created_at': '2024-01-31T19:06:21Z',
       	'ended_at': '2024-01-31T19:06:22Z',
       	'error': '',
       	'org_id': '3340851',
       	'repository_name': 'test-user-jdoe',
       	'repository_uuid': '8c06233b-19c0-4cd1-836c-d230b803cc6e',
       	'status': 'completed',
       	'type': 'introspect',
       	'uuid': '335bde32-418d-4bd7-bbfa-4777537ff3d4'

```

Here is task show by another user in another org:

```
In [38]: rw_app.content_sources.rest_client.tasks_api.list_tasks(status="completed")

Out[38]:
{'data': [{'created_at': '2024-01-31T19:37:34Z',
       	'ended_at': '2024-01-31T19:37:34Z',
       	'error': '',
       	'org_id': '45',
       	'repository_name': 'some-other-user',
       	'repository_uuid': 'c2893a62-f8df-4627-aa16-1090dd284bc1',
       	'status': 'completed',
       	'type': 'introspect',
       	'uuid': 'a4c8e075-ab64-439e-b05a-6a171f7a3187'},

```


Task and repository_uuids are unique 

---

## Reviews

### Review by @rverdile - Approved on January 30, 2024 at 07:46 PM UTC

tested and working!

Also, as far as I can tell, this also fixes this issue: https://issues.redhat.com/browse/HMS-3371. If you agree, I think we should add that to the title. 

You should be able to verify by following your testing steps, but querying the list endpoint instead of fetching each task.

For example
1. Create two repos in different orgs with same url
2. Filter task list by type snapshot with the org ID of the first repo.
3. You'll see both snapshot tasks, but only the org ID of the first repo. This is misleading because the org ID for the second repo's task is wrong. That task should not be shown at all.
4. Testing against your PR, I only see the tasks for the first repo.


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/552*
