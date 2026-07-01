---
type: pull_request
number: 536
title: "Fixes 2983: add api for searching snap comps groups"
state: merged
author: rverdile
created: 2024-01-17T22:04:42Z
updated: 2024-02-01T20:00:33Z
closed: 2024-02-01T19:57:30Z
merged: 2024-02-01T19:57:30Z
base_branch: main
head_branch: search-snaps-comps
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/536
---

# Pull Request #536: Fixes 2983: add api for searching snap comps groups

**Author**: @rverdile
**Created**: January 17, 2024 at 10:04 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `search-snaps-comps`

## Description

## Summary

- ~~Built on this PR: https://github.com/content-services/content-sources-backend/pull/517. Do not merge before this PR has merged.~~
- ~~Waiting for this dependency PR to finish and merge: https://github.com/content-services/tang/pull/3~~
- Adds endpoints to search environments and package groups of snapshots, by name.
- Also fixes issue where ID metadata was not being exposed for package groups and environments. 
- Functionality is the same as the existing searching endpoints for repositories. The only difference is you must send a snapshot UUID in the request body (instead of a repository UUID or URL).

## Testing steps
1. Create a snapshot
2. These two repos are good for testing: https://rverdile.fedorapeople.org/dummy-repos/comps/. So are epel and rh repos.
3. Search package groups by name using `POST snapshots/package_groups/names`
4. Search environments by name using `POST snapshots/environments/names/`
5. Body of the request looks something like `{"uuids":["abcd"],"search":"demo","limit":50}`, where uuid is the uuid of the snapshot.

## Checklist
- [x] Revisit when dependency PRs are finished
- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 17, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-2983

### Comment by @swadeley on January 25, 2024 at 01:54 PM UTC

/retest

### Comment by @swadeley on January 25, 2024 at 02:30 PM UTC

/retest

### Comment by @xbhouse on January 25, 2024 at 03:12 PM UTC

other than the nitpicky comment, this looks good to me and works well!! will wait for Justin's input on the test fix

### Comment by @swadeley on January 25, 2024 at 06:37 PM UTC

/retest

### Comment by @swadeley on January 26, 2024 at 03:17 PM UTC

/retest

### Comment by @swadeley on January 26, 2024 at 09:14 PM UTC

Hi @rverdile , I checked out :
 commit b571fac95e4213882885ec5ac8c1867afdaeea57 (HEAD -> pr/536)


In the git diff I see:
"description": "A comma separated list of uuids to control api response.",

which looks inconsistent compared to:
description": "List of Snapshot UUIDs to search",

I think it should be `UUIDs`.

But then searching the whole file I find:
 "List of RepositoryConfig UUIDs to search",
"Repository uuids to find snapshots for",

so, its inconsistent in other places too.

Thank you

### Comment by @swadeley on January 31, 2024 at 08:45 PM UTC

@rverdile Thanks for the update, please rebase.

### Comment by @rverdile on February 01, 2024 at 03:54 PM UTC

rebased

### Comment by @swadeley on February 01, 2024 at 04:19 PM UTC

Hi @rverdile 

I created a repo and a snapshot, but I don't see package group "birds":

```
In [3]: app.content_sources.rest_client.repositories_api.create_repository(repo)
2024-02-01 15:54:51.379 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[3]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_snapshot_task_uuid': '1747e6d6-e49d-41f5-b8f4-07b83675bce6',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'Hello Ryan',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'https://rverdile.fedorapeople.org/dummy-repos/comps/repo1/',
 'uuid': 'ac1122a5-64f0-4cae-b8e3-f2826861e0d1'}

In [4]: app.content_sources.rest_client.tasks_api.get_task('1747e6d6-e49d-41f5-b8f4-07b83675bce6')
2024-02-01 16:03:20.705 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[4]: 
{'created_at': '2024-02-01T15:54:51Z',
 'ended_at': '2024-02-01T15:55:05Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'Hello Ryan',
 'repository_uuid': 'ac1122a5-64f0-4cae-b8e3-f2826861e0d1',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '1747e6d6-e49d-41f5-b8f4-07b83675bce6'}

In [5]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(uuids=['1747e6d6-e49d-41f5-b8f4-07b83675bce6'],search='demo'))
2024-02-01 16:07:08.695 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[5]: []


In [7]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(uuids=['1747e6d6-e49d-41f5-b8f4-07b83675bce6'],search='birds'))
2024-02-01 16:08:02.487 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[7]: []
```


```
In [11]: app.content_sources.rest_client.snapshots_api.search_snapshot_environments(dict(uuids=['1747e6d6-e49d-41f5-b8f4-07b83675bce6'],search='avians'))
2024-02-01 16:14:00.509 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[11]: []
```


Are these parameter maps correct? The all have
 ` 'required': ['api_snapshot_search_rpm_request'],`:

```
In [12]: app.content_sources.rest_client.snapshots_api.search_snapshot_environments.params_map
Out[12]: 
{'all': ['api_snapshot_search_rpm_request',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['api_snapshot_search_rpm_request'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}

In [13]: app.content_sources.rest_client.snapshots_api.search_snapshot_rpms.params_map
Out[13]: 
{'all': ['api_snapshot_search_rpm_request',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['api_snapshot_search_rpm_request'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}

In [14]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups.params_map
Out[14]: 
{'all': ['api_snapshot_search_rpm_request',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['api_snapshot_search_rpm_request'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}
```

Thank you


### Comment by @rverdile on February 01, 2024 at 06:59 PM UTC

@swadeley I think it's because you're searching on the "last_snapshot_task_uuid" instead of the snapshot_uuid. To get the snapshot_uuid I use `GET repositories/:repo_uuid/snapshots/`.

### Comment by @swadeley on February 01, 2024 at 07:56 PM UTC

Hi @rverdile , your are correct, I had wrong UUID (must use UUID from snapshot itself not task)

```
In [3]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(search='birds', uuids=['a6ec82ce-d443-490b-b89f-6091169b6c13']))
2024-02-01 19:51:26.180 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[3]: 
[{'description': 'birds',
  'id': 'birds',
  'package_group_name': 'birds',
  'package_list': ['cockateel', 'penguin', 'stork']}]

In [4]: app.content_sources.rest_client.snapshots_api.search_snapshot_environments(dict(uuids=['a6ec82ce-d443-490b-b89f-6091169b6c13'],search='avians'))
2024-02-01 19:52:02.012 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[4]: [{'description': 'avians', 'environment_name': 'avians', 'id': 'avians'}]
```


To get the snapshot:

```
In [5]: app.content_sources.rest_client.repositories_api.list_snapshots('ac1122a5-64f0-4cae-b8e3-f2826861e0d1')
2024-02-01 19:55:02.130 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[5]: 
{'data': [{'added_counts': {'rpm.package': 3,
                            'rpm.packageenvironment': 1,
                            'rpm.packagegroup': 1,
                            'rpm.packagelangpacks': 1},
           'content_counts': {'rpm.package': 3,
                              'rpm.packageenvironment': 1,
                              'rpm.packagegroup': 1,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-02-01T15:55:05.752976Z',
           'removed_counts': {},
           'repository_path': 'e89bc8d9/ac1122a5-64f0-4cae-b8e3-f2826861e0d1/d5ef8122-04ab-4bc0-a6bd-e6311d32589a',
           'url': 'https://e<>.openshiftapps.com/pulp/content/e89bc8d9/ac1122a5-64f0-4cae-b8e3-f2826861e0d1/d5ef8122-04ab-4bc0-a6bd-e6311d32589a/',
           'uuid': 'a6ec82ce-d443-490b-b89f-6091169b6c13'}],
 'links': {'first': '/api/content-sources/v1/repositories/ac1122a5-64f0-4cae-b8e3-f2826861e0d1/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/ac1122a5-64f0-4cae-b8e3-f2826861e0d1/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

Thank you

---

## Reviews

### Review by @rverdile - Commented on January 23, 2024 at 08:45 PM UTC

### Review by @rverdile - Commented on January 23, 2024 at 08:46 PM UTC

### Review by @rverdile - Commented on January 23, 2024 at 08:49 PM UTC

### Review by @xbhouse - Commented on January 25, 2024 at 02:33 PM UTC

### Review by @xbhouse - Commented on January 25, 2024 at 02:36 PM UTC

### Review by @rverdile - Commented on January 25, 2024 at 02:56 PM UTC

### Review by @rverdile - Commented on January 25, 2024 at 03:00 PM UTC

### Review by @xbhouse - Commented on January 25, 2024 at 03:06 PM UTC

### Review by @xbhouse - Approved on January 26, 2024 at 02:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/536*
