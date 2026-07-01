---
type: pull_request
number: 359
title: "Fixes 1971: Task info list filtering and new fields"
state: merged
author: rverdile
created: 2023-08-14T17:45:50Z
updated: 2023-08-28T20:35:21Z
closed: 2023-08-17T09:19:53Z
merged: 2023-08-17T09:19:53Z
base_branch: main
head_branch: task-list-filter
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/359
---

# Pull Request #359: Fixes 1971: Task info list filtering and new fields

**Author**: @rverdile
**Created**: August 14, 2023 at 05:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `task-list-filter`

## Description

## Summary
This adds 3 new fields to the task fetch and list API responses: type, repository name, repository (config) UUID. It also adds filtering by type and repository UUID to the task list endpoint.
 
## Testing steps
1. Create a repository with snapshot enabled so that a couple of tasks are produced.
2. List the tasks without filtering, filtering by repository UUID, and filtering by name. Check for new fields.
3. Fetch a task and check for new fields

---

## Discussion

### Comment by @jlsherrill on August 14, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-1971

### Comment by @jlsherrill on August 16, 2023 at 01:48 PM UTC

Everything else looks good!  I'll wait to test in case you update those queries

### Comment by @swadeley on August 17, 2023 at 09:11 AM UTC

Hi

I ran apigen.

I created a repo:

```
Out[4]: 
{'account_id': '0369233',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'fedoratest-1',
 'org_id': '3340851',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo01/',
 'uuid': '54e92977-9c58-43de-be31-bcacc92ce721'}
```


Looks good:

```
In [5]: app.content_sources.rest_client.tasks_api.list_tasks()
2023-08-17 09:32:04.557 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=XZilDYrmQ2jGx0bZ11iia5bwciCYanWt, params=[]
Out[5]: 
{'data': [{'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:22Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '48bda3b4-aa25-4183-bc2e-a49009624f89'},
          {'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:32Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '674b6996-8e8d-45b3-9532-e60b1c9257e3'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

```

### Comment by @swadeley on August 17, 2023 at 09:18 AM UTC

Test filter for tasks:

```
In [9]: app.content_sources.rest_client.tasks_api.list_tasks.attribute_map
Out[9]: 
{'offset': 'offset',
 'limit': 'limit',
 'status': 'status',
 'type': 'type',
 'repository_uuid': 'repository_uuid'}
```

By UUID of the repo:

```
In [10]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid='54e92977-9c58-43de-be31-bcacc92ce721')
2023-08-17 10:11:22.121 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=JXHeBcq9kMdBS5v1kbUZgtxfkUQK4goZ, params=[('repository_uuid', '54e92977-9c58-43de-be31-bcacc92ce721')]
Out[10]: 
{'data': [{'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:22Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '48bda3b4-aa25-4183-bc2e-a49009624f89'},
          {'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:32Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '674b6996-8e8d-45b3-9532-e60b1c9257e3'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=54e92977-9c58-43de-be31-bcacc92ce721',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=54e92977-9c58-43de-be31-bcacc92ce721'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

```

By type:

```
In [12]: app.content_sources.rest_client.tasks_api.list_tasks(type='snapshot')
2023-08-17 10:14:42.443 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=p0dFltTVLEOjWYrcwP7y7zOLOpntgw66, params=[('type', 'snapshot')]
Out[12]: 
{'data': [{'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:32Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '674b6996-8e8d-45b3-9532-e60b1c9257e3'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=snapshot'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [13]: app.content_sources.rest_client.tasks_api.list_tasks(type='introspect')
2023-08-17 10:16:54.933 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=5odjHScgCyDZ4h4cuow45ZFhqpuzK8oZ, params=[('type', 'introspect')]
Out[13]: 
{'data': [{'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:22Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '48bda3b4-aa25-4183-bc2e-a49009624f89'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=introspect',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&type=introspect'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

```

By status:

```
In [14]: app.content_sources.rest_client.tasks_api.list_tasks(status='completed')
2023-08-17 10:17:52.293 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=96gTBMtJVseW6EeZVuEOjY4d6snTKs3P, params=[('status', 'completed')]
Out[14]: 
{'data': [{'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:22Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'introspect',
           'uuid': '48bda3b4-aa25-4183-bc2e-a49009624f89'},
          {'created_at': '2023-08-17T08:23:23Z',
           'ended_at': '2023-08-17T08:29:32Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'fedoratest-1',
           'repository_uuid': '54e92977-9c58-43de-be31-bcacc92ce721',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '674b6996-8e8d-45b3-9532-e60b1c9257e3'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&status=completed',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&status=completed'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```



### Comment by @swadeley on August 17, 2023 at 09:20 AM UTC

All good.

---

## Reviews

### Review by @jlsherrill - Commented on August 16, 2023 at 01:38 PM UTC

### Review by @jlsherrill - Commented on August 16, 2023 at 01:40 PM UTC

### Review by @rverdile - Commented on August 16, 2023 at 03:05 PM UTC

### Review by @jlsherrill - Approved on August 16, 2023 at 03:29 PM UTC

Works as advertised!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/359*
