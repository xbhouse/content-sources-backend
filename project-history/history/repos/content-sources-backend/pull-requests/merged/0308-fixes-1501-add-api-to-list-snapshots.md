---
type: pull_request
number: 308
title: "Fixes 1501: add API to list snapshots"
state: merged
author: rverdile
created: 2023-06-08T15:31:41Z
updated: 2023-06-23T15:30:32Z
closed: 2023-06-23T15:14:50Z
merged: 2023-06-23T15:14:50Z
base_branch: main
head_branch: snap-list
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/308
---

# Pull Request #308: Fixes 1501: add API to list snapshots

**Author**: @rverdile
**Created**: June 08, 2023 at 03:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `snap-list`

## Description

## Summary
Adds an API endpoint to list the snapshots of a repository: `repositories/:uuid/snapshots/`
## Testing steps
1. Using the API, create a repository, and in the body include the field `snapshot: true`
2. GET request `repositories/{new repository UUID}/snapshots/`
3. Should see the snapshot returned, including pagination metadata.


---

## Discussion

### Comment by @jlsherrill on June 12, 2023 at 01:36 PM UTC

https://issues.redhat.com/browse/HMS-1501

### Comment by @swadeley on June 19, 2023 at 12:01 PM UTC

/retest

### Comment by @swadeley on June 21, 2023 at 08:31 AM UTC

/retest

### Comment by @swadeley on June 22, 2023 at 10:37 AM UTC

/retest

### Comment by @swadeley on June 23, 2023 at 02:22 PM UTC

Hi

I created a repo with snapshot=true

I can see the tasks (one is for introspection, one for snapshot):

```
In [3]: app.content_sources.rest_client.tasks_api.list_tasks()
2023-06-23 15:05:48.753 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=o4K7ve1lVtam5a28d7wagxZQ8sH5yOSk, params=[]
Out[3]: 
{'data': [{'created_at': '2023-06-23T14:05:44Z',
           'ended_at': '2023-06-23T14:05:44Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': 'b8b72afd-0859-491b-9182-4de101fe0dab'},
          {'created_at': '2023-06-23T14:05:44Z',
           'ended_at': '',
           'error': '',
           'org_id': '3340851',
           'status': 'running',
           'uuid': 'e6fb679a-7eac-46b7-b2e6-78a077005c5e'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```

A bit later (see second task's "ended_at" and "status" values):
```
In [5]: app.content_sources.rest_client.tasks_api.list_tasks()
2023-06-23 15:06:47.160 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=Ijz8qGoQSYFmBNscX2zLXgQ8kVb6RSc7, params=[]
Out[5]: 
{'data': [{'created_at': '2023-06-23T14:05:44Z',
           'ended_at': '2023-06-23T14:05:44Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': 'b8b72afd-0859-491b-9182-4de101fe0dab'},
          {'created_at': '2023-06-23T14:05:44Z',
           'ended_at': '2023-06-23T14:06:19Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': 'e6fb679a-7eac-46b7-b2e6-78a077005c5e'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```



### Comment by @swadeley on June 23, 2023 at 02:44 PM UTC

hi
we can also GET the individual task info

```
In [1]: app.content_sources.rest_client.tasks_api.get_task('b8b72afd-0859-491b-9182-4de101fe0dab')
2023-06-23 15:42:39.996 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=BP84Ui30cLMXMCy1SmxmyzLxQzQD39tf, params=[]
Out[1]: 
{'created_at': '2023-06-23T14:05:44Z',
 'ended_at': '2023-06-23T14:05:44Z',
 'error': '',
 'org_id': '3340851',
 'status': 'completed',
 'uuid': 'b8b72afd-0859-491b-9182-4de101fe0dab'}

In [2]: app.content_sources.rest_client.tasks_api.get_task('e6fb679a-7eac-46b7-b2e6-78a077005c5e')
2023-06-23 15:43:03.491 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=JapXcX2MQZ6YMKo8Tz6cwmJAfuAKs4AA, params=[]
Out[2]: 
{'created_at': '2023-06-23T14:05:44Z',
 'ended_at': '2023-06-23T14:06:19Z',
 'error': '',
 'org_id': '3340851',
 'status': 'completed',
 'uuid': 'e6fb679a-7eac-46b7-b2e6-78a077005c5e'}
```

### Comment by @swadeley on June 23, 2023 at 02:50 PM UTC

Hi

Using the UUID of the repo, we can get info about snapshots of that repo:

```
In [5]: app.content_sources.rest_client.repositories_api.list_snapshots('60c3b07f-9f53-4c62-ac43-80c639edb30a')
2023-06-23 15:48:34.484 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=Lfv9BE5yCF9Z1CSNMk8t27eWNMD6mfgY, params=[]
Out[5]: 
{'data': [{'content_counts': {'rpm.package': 1},
           'created_at': '2023-06-23T14:06:19.383977Z',
           'distribution_path': '60c3b07f-9f53-4c62-ac43-80c639edb30a/2a3fea66-734f-4172-a8e2-b219c8a78990'}],
 'links': {'first': '/api/content-sources/v1/repositories/60c3b07f-9f53-4c62-ac43-80c639edb30a/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/60c3b07f-9f53-4c62-ac43-80c639edb30a/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

---

## Reviews

### Review by @jlsherrill - Commented on June 13, 2023 at 05:01 PM UTC

### Review by @jlsherrill - Commented on June 13, 2023 at 05:02 PM UTC

### Review by @rverdile - Commented on June 13, 2023 at 07:09 PM UTC

### Review by @rverdile - Commented on June 13, 2023 at 07:10 PM UTC

### Review by @jlsherrill - Commented on June 13, 2023 at 07:21 PM UTC

### Review by @jlsherrill - Commented on June 13, 2023 at 07:21 PM UTC

### Review by @rverdile - Commented on June 13, 2023 at 07:56 PM UTC

### Review by @jlsherrill - Approved on June 16, 2023 at 02:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/308*
