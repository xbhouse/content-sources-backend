---
type: pull_request
number: 766
title: "Fixes 4313: update use_latest templates after new snapshot"
state: merged
author: rverdile
created: 2024-08-07T14:12:04Z
updated: 2024-08-20T11:00:29Z
closed: 2024-08-20T10:34:42Z
merged: 2024-08-20T10:34:42Z
base_branch: main
head_branch: 4313
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/766
---

# Pull Request #766: Fixes 4313: update use_latest templates after new snapshot

**Author**: @rverdile
**Created**: August 07, 2024 at 02:12 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4313`

## Description

## Summary
- Adds a new task `update-latest-snapshot` that updates all use_latest templates to use the latest snapshot for a given repository. This task is enqueued as a dependent of every snapshot task, so that it runs after the snapshot is completed. 
- Adds support for dependencies in the tasking system. Adds the "dependencies" and "dependents" fields to TaskInfoResponse. If a parent task fails, dependent tasks are canceled. If a parent task is requeued (after failure), dependent tasks would also be requeued.
 ## Testing steps
1. Create a repository with snapshot=true
2. Create a template with use_latest=true and add the repository.
3. List the template RPMs to compare with later
4. Update the repository to change the URL
5. Wait for the snapshot and update-latest-snapshot tasks to complete
6. Fetching the task details for the snapshot task should return the UUID of the update-latest-snapshot task in the "dependents" field. Fetching the details for the update-latest-snapshot task should show the UUID of the snapshot task in the "dependencies" field.
7. List the template RPMs again, they should instead be the RPMs from the new URL

---

## Discussion

### Comment by @jlsherrill on August 07, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-4313

### Comment by @jlsherrill on August 08, 2024 at 07:37 PM UTC

overall code looks good, just one question, and needs pointy references removed.  I will start testing!

### Comment by @swadeley on August 09, 2024 at 11:06 AM UTC

build failed:
`{"level":"error","module":"pgx","database":"postgres","err":"failed to connect to `user=postgres database=postgres`:\n\t127.0.0.1:5433 (localhost): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)\n\t[::1]:5433 (localhost): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)","host":"localhost","port":5433,"time":9.739279,"time":"2024-08-09T08:46:32Z","message":"Connect"}`

### Comment by @jlsherrill on August 09, 2024 at 01:58 PM UTC

testing all looked good, just the one comment 

### Comment by @swadeley on August 13, 2024 at 10:06 AM UTC

/retest

### Comment by @swadeley on August 14, 2024 at 02:35 PM UTC

Hi @rverdile, can you check the type is set correctly everywhere for `dependencies` as I get this error:

`Required value type is list and passed type was NoneType at ['received_data']['last_snapshot_task']['dependencies']`

### Comment by @rverdile on August 14, 2024 at 10:26 PM UTC

/retest

### Comment by @swadeley on August 15, 2024 at 09:51 AM UTC

/retest

### Comment by @swadeley on August 15, 2024 at 11:07 AM UTC

Hi

thanks for the update, I notice no API docs changed so no rebuild (generate-api) is required. 

I deployed `backend=pr-766-6b119d3` for testing.

I created a repo and added it to a template.
I changed the repo URL to point to a different repo, but these fields are empty:

```
           'dependencies': [],
           'dependents': [],
```


```
Out[1]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test10',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-08-15T10:44:14Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 5,
                                              'rpm.package': 22},
                             'content_counts': {'rpm.advisory': 5,
                                                'rpm.package': 23},
                             'created_at': '2024-08-15T10:44:27.930229Z',
                             'removed_counts': {},
                             'repository_path': '6e5fa381/41b3f191-46cb-4782-8126-daf1db32968d/48edfa23-f6e6-4dec-a571-769eb6e2e9cf',
                             'url': 'http://pulp-content:8000/api/pulp-content/6e5fa381/41b3f191-46cb-4782-8126-daf1db32968d/48edfa23-f6e6-4dec-a571-769eb6e2e9cf/',
                             'uuid': 'bc3dca35-0c62-4c9e-856e-386bdb2944a0'},
           'last_snapshot_task': {'created_at': '2024-08-15T10:44:13Z',
                                  'dependencies': [],
                                  'dependents': [],
                                  'ended_at': '2024-08-15T10:44:27Z',
                                  'error': '',
                                  'org_id': '3340851',
                                  'repository_name': 'test10',
                                  'repository_uuid': '41b3f191-46cb-4782-8126-daf1db32968d',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': 'b96d1f31-5473-4903-b30a-c4c7d22e0e56'},
           'last_snapshot_task_uuid': 'b96d1f31-5473-4903-b30a-c4c7d22e0e56',
           'last_snapshot_uuid': 'bc3dca35-0c62-4c9e-856e-386bdb2944a0',
           'last_success_introspection_time': '2024-08-15T10:44:14Z',
           'last_update_introspection_time': '2024-08-15T10:44:14Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test10',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 23,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://stephenw.fedorapeople.org/fakerepos/fake_yum3/',
           'uuid': '41b3f191-46cb-4782-8126-daf1db32968d'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&status=Valid',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&status=Valid'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [2]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid="41b3f191-46cb-4782-8126-daf1db32968d")
2024-08-15 11:52:50.679 [    INFO] [iqe.base.rest_client] REST: GET https://ee-epgpoew8.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/ with query params [('repository_uuid', '41b3f191-46cb-4782-8126-daf1db32968d')] and x-rh-insights-request-id=<snip>
Out[2]: 
{'data': [{'created_at': '2024-08-15T10:44:13Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-15T10:44:13Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test10',
           'repository_uuid': '41b3f191-46cb-4782-8126-daf1db32968d',
           'status': 'completed',
           'type': 'update-repository',
           'uuid': '8f216eb5-c178-4018-8e13-ed944c60aca3'},
          {'created_at': '2024-08-15T10:44:13Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-15T10:44:14Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test10',
           'repository_uuid': '41b3f191-46cb-4782-8126-daf1db32968d',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '6c1ba72c-ada8-45ad-80aa-9c12c650779f'},
          {'created_at': '2024-08-15T10:44:13Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-15T10:44:27Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test10',
           'repository_uuid': '41b3f191-46cb-4782-8126-daf1db32968d',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': 'b96d1f31-5473-4903-b30a-c4c7d22e0e56'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=41b3f191-46cb-4782-8126-daf1db32968d',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=41b3f191-46cb-4782-8126-daf1db32968d'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [3]: 

```

Get snapshot task:
```
In [3]: app.content_sources.rest_client.tasks_api.get_task('b96d1f31-5473-4903-b30a-c4c7d22e0e56')
2024-08-15 13:41:28.918 [    INFO] [iqe.base.rest_client] REST: GET https://ee-epgpoew8.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/b96d1f31-5473-4903-b30a-c4c7d22e0e56 with query params [] and x-rh-insights-request-id=<-snip->
Out[3]: 
{'created_at': '2024-08-15T10:44:13Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '2024-08-15T10:44:27Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test10',
 'repository_uuid': '41b3f191-46cb-4782-8126-daf1db32968d',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': 'b96d1f31-5473-4903-b30a-c4c7d22e0e56'
```

### Comment by @rverdile on August 15, 2024 at 12:58 PM UTC

> I changed the repo URL to point to a different repo, but these fields are empty

I still have to look at this. It could also be addressed in a followup PR. This might be a lot harder to fix then it seems as it could be caused by a limitation of the ORM we use. You can see the dependencies and dependents if you fetch the task directly. 

### Comment by @jlsherrill on August 15, 2024 at 02:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on August 16, 2024 at 03:14 PM UTC

Hi @rverdile 

Please check bottom of my post above yours, I did fetch the task directly:

`2024-08-15 13:41:28.918 [    INFO] [iqe.base.rest_client] REST: GET https://ee-<snip>.com/api/content-sources/v1/tasks/b96d1f31-5473-4903-b30a-c4c7d22e0e56 with query params []`

Here is another one, I just deployed `backend=pr-766-243200b` and tested:

```
In [2]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid='f469806a-6dd8-4884-9d51-1d82c572b689')
2024-08-16 16:07:25.170 [    INFO] [iqe.base.rest_client] REST: GET https://ee-dhkvm7ol.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/ with query params [('repository_uuid', 'f469806a-6dd8-4884-9d51-1d82c572b689')] and x-rh-insights-request-id=<snip>
Out[2]: 
{'data': [{'created_at': '2024-08-16T15:03:33Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-16T15:03:33Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test1-2',
           'repository_uuid': 'f469806a-6dd8-4884-9d51-1d82c572b689',
           'status': 'completed',
           'type': 'update-repository',
           'uuid': 'eb3d6201-188a-4eaa-933c-47e4e59ddd83'},
          {'created_at': '2024-08-16T15:03:33Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-16T15:03:33Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test1-2',
           'repository_uuid': 'f469806a-6dd8-4884-9d51-1d82c572b689',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '2d146b73-8ff3-4c0a-9484-bb5f9dd8ea6a'},
          {'created_at': '2024-08-16T15:03:33Z',
           'dependencies': [],
           'dependents': [],
           'ended_at': '2024-08-16T15:03:46Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'test1-2',
           'repository_uuid': 'f469806a-6dd8-4884-9d51-1d82c572b689',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': 'faf0121a-cc23-417a-8099-02de7def6d8a'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=f469806a-6dd8-4884-9d51-1d82c572b689',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=f469806a-6dd8-4884-9d51-1d82c572b689'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [3]: app.content_sources.rest_client.tasks_api.get_task('faf0121a-cc23-417a-8099-02de7def6d8a')
2024-08-16 16:09:42.899 [    INFO] [iqe.base.rest_client] REST: GET https://ee-dhkvm7ol.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/faf0121a-cc23-417a-8099-02de7def6d8a with query params [] and x-rh-insights-request-id=<snip>
Out[3]: 
{'created_at': '2024-08-16T15:03:33Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '2024-08-16T15:03:46Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test1-2',
 'repository_uuid': 'f469806a-6dd8-4884-9d51-1d82c572b689',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': 'faf0121a-cc23-417a-8099-02de7def6d8a'}

In [4]: 

```

Point 7 is OK, I see the new packages from the new repo.

Thank you

### Comment by @swadeley on August 19, 2024 at 07:40 AM UTC

/retest

### Comment by @rverdile on August 19, 2024 at 07:00 PM UTC

@swadeley it should work now for listing or fetching tasks directly

### Comment by @swadeley on August 20, 2024 at 10:34 AM UTC

Hi @rverdile 

Thank you for the update. All looks good now:

From the repo I got:
` 'last_snapshot_task_uuid': '15e472f4-b626-40ff-b672-d084e89afeee',`

then I used that to get ast_snapshot_task details, then used dependants UUID to get other tasks details:
```
In [2]: app.content_sources.rest_client.tasks_api.get_task('15e472f4-b626-40ff-b672-d084e89afeee')
2024-08-20 11:27:27.562 [    INFO] [iqe.base.rest_client] REST: GET https://ee-agc7mp9w.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/15e472f4-b626-40ff-b672-d084e89afeee with query params [] and x-rh-insights-request-id=<snip>
Out[2]: 
{'created_at': '2024-08-20T10:25:10Z',
 'dependencies': [],
 'dependents': ['c479a57c-a66d-4c45-9f1b-ff5aa167c11b'],
 'ended_at': '2024-08-20T10:25:22Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test1.2',
 'repository_uuid': 'ec450365-7260-4aef-93d8-7d44693814c0',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '15e472f4-b626-40ff-b672-d084e89afeee'}

In [3]: app.content_sources.rest_client.tasks_api.get_task('c479a57c-a66d-4c45-9f1b-ff5aa167c11b')
2024-08-20 11:29:09.132 [    INFO] [iqe.base.rest_client] REST: GET https://ee-agc7mp9w.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/c479a57c-a66d-4c45-9f1b-ff5aa167c11b with query params [] and x-rh-insights-request-id=<snip>
Out[3]: 
{'created_at': '2024-08-20T10:25:10Z',
 'dependencies': ['15e472f4-b626-40ff-b672-d084e89afeee'],
 'dependents': [],
 'ended_at': '2024-08-20T10:25:22Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test1.2',
 'repository_uuid': 'ec450365-7260-4aef-93d8-7d44693814c0',
 'status': 'completed',
 'type': 'update-latest-snapshot',
 'uuid': 'c479a57c-a66d-4c45-9f1b-ff5aa167c11b'}

In [4]: 
```

---

## Reviews

### Review by @jlsherrill - Commented on August 07, 2024 at 06:56 PM UTC

### Review by @rverdile - Commented on August 08, 2024 at 06:08 PM UTC

### Review by @jlsherrill - Commented on August 08, 2024 at 06:56 PM UTC

### Review by @rverdile - Commented on August 08, 2024 at 08:03 PM UTC

### Review by @jlsherrill - Commented on August 08, 2024 at 09:44 PM UTC

### Review by @jlsherrill - Commented on August 08, 2024 at 09:46 PM UTC

### Review by @rverdile - Commented on August 12, 2024 at 06:49 PM UTC

### Review by @jlsherrill - Approved on August 13, 2024 at 02:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/766*
