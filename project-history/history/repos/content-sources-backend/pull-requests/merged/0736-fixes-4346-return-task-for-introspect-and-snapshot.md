---
type: pull_request
number: 736
title: "Fixes 4346: return task for introspect and snapshot APIs"
state: merged
author: rverdile
created: 2024-07-09T17:38:05Z
updated: 2024-07-10T21:30:22Z
closed: 2024-07-10T21:00:24Z
merged: 2024-07-10T21:00:24Z
base_branch: main
head_branch: return-task
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/736
---

# Pull Request #736: Fixes 4346: return task for introspect and snapshot APIs

**Author**: @rverdile
**Created**: July 09, 2024 at 05:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `return-task`

## Description

## Summary
When using repository introspect and snapshot ends points, the response will now be the task info struct.

## Testing steps
1. Use the repository introspect API to introspect a repository, it should return a 200 and a json with the task info.
2. Use the repository snapshot API to snapshot a repository, it should return a 200 and a json with the task info.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 09, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-4346

### Comment by @swadeley on July 10, 2024 at 09:21 AM UTC

Hi

testing on master before deploying this image.

I created a repo:

```
In [4]: created_repo = app.content_sources.rest_client.repositories_api.create_repository(repo_dict_1)
2024-07-10 09:57:24.946 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

In [5]: created_repo
<snip>
 'last_introspection_status': 'Pending',
 'last_introspection_time': '',
```

I got the repo:

```
In [6]: app.content_sources.rest_client.repositories_api.get_repository('449fdf74-fe8d-4900-9248-a6445d8aa6a3')
2024-07-10 10:05:50.417 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[6]: 
<snip>
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-07-10T08:57:25Z'
```

I triggered an introspection:

```
In [9]: app.content_sources.rest_client.repositories_api.introspect('449fdf74-fe8d-4900-9248-a6445d8aa6a3')
2024-07-10 10:13:16.284 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
```
Nothing was returned, but we can see that the `last_introspection_time` time stamp changed:

```
In [10]: app.content_sources.rest_client.repositories_api.get_repository('449fdf74-fe8d-4900-9248-a6445d8aa6a3')
2024-07-10 10:13:24.728 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[10]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test_repo_name_1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-07-10T09:13:16Z',

```


### Comment by @swadeley on July 10, 2024 at 10:05 AM UTC

Hi @rverdile 

I deployed the pr-736-fa37e5d image to a new ephemeral namespace but I do not get anything returned for the introspection command.

I created new repo:
`'last_introspection_time': '2024-07-10T09:51:33Z',`

Trigged an introspection:
```
In [5]: app.content_sources.rest_client.repositories_api.introspect('ab7db788-b052-47b2-9f67-f4183a746e2c')
2024-07-10 10:57:14.150 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

In [6]: app.content_sources.rest_client.repositories_api.get_repository('ab7db788-b052-47b2-9f67-f4183a746e2c')['last_introspection_time']
2024-07-10 11:01:03.640 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[6]: '2024-07-10T09:57:14Z'
```

There are no changes to the API Docs. I wonder do I need to checkout the branch while I deploy the backend image?

Thank you

### Comment by @swadeley on July 10, 2024 at 07:00 PM UTC

Hi

Looks good now thank you:

```

In [4]: app.content_sources.rest_client.repositories_api.introspect(created_repo['uuid'])
2024-07-10 19:54:49.312 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[4]: 
{'created_at': '2024-07-10T18:54:49Z',
 'ended_at': '',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test_repo_name_1',
 'repository_uuid': '5482c878-c72a-4ff3-96e4-b69ab914b326',
 'status': 'pending',
 'type': 'introspect',
 'uuid': 'f92f128b-073e-4ea5-8367-e4f624df33ae'}
```

```
In [7]: app.content_sources.rest_client.repositories_api.create_snapshot(created_repo['uuid'])
2024-07-10 19:59:21.704 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[7]: 
{'created_at': '2024-07-10T18:59:21Z',
 'ended_at': '',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test_repo_name_1',
 'repository_uuid': '5482c878-c72a-4ff3-96e4-b69ab914b326',
 'status': 'pending',
 'type': 'snapshot',
 'uuid': '29093aeb-df51-4085-a013-9d10d092c8d8'}
```

---

## Reviews

### Review by @xbhouse - Approved on July 09, 2024 at 06:59 PM UTC

this looks good! i think you'll just need to rebase to fix the pulp image pull error

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/736*
