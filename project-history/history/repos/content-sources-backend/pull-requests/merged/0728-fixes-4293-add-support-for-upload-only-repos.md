---
type: pull_request
number: 728
title: "Fixes 4293: add support for upload only repos"
state: merged
author: jlsherrill
created: 2024-07-02T13:15:17Z
updated: 2024-07-15T15:30:25Z
closed: 2024-07-15T15:08:34Z
merged: 2024-07-15T15:08:34Z
base_branch: main
head_branch: 4293
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/728
---

# Pull Request #728: Fixes 4293: add support for upload only repos

**Author**: @jlsherrill
**Created**: July 02, 2024 at 01:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4293`

## Description

## Summary

## Testing steps

Create an upload repo:

```
PATCH http://localhost:8000/api/content-sources/v1.0/repositories/c06a0512-0954-4b11-9b13-3358afe77f46
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json

{
  "name": "arg",
  "origin": "upload"
}
```

with batch:
```
POST http://localhost:8000/api/content-sources/v1.0/repositories/bulk_create/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==

[{
  "name": "test5",
  "gpg_key": "",
  "snapshot": true,
  "origin": "upload",
  "url": ""
},
  {
    "name": "test6",
    "gpg_key": "",
    "snapshot": true,
    "origin": "upload",
    "url": ""
  }]

```

try to:

1.  create/update an upload repo with a url, this should fail
2. try to change the origin of a repo, this should not succeed
3. try to introspect an upload repo, this should be rejected
4.  try to snapshot an upload repo, this should be rejected 
5.  run `go run cmd/external-repos/main.go nightly-jobs` and verify no errors, that things work as expected

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 02, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4293

### Comment by @swadeley on July 15, 2024 at 08:43 AM UTC

/retest

### Comment by @swadeley on July 15, 2024 at 03:06 PM UTC

Hi

I can create an upload repo:

```

In [7]:     repo_dict_3 = dict(
   ...:         name="test_upload_repo_2",
   ...:         origin="upload",
   ...:         snapshot=True,
   ...:      )

In [8]: created_repo3 = app.content_sources.rest_client.repositories_api.create_repository(repo_dict_3)
2024-07-15 15:10:58.648 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

In [9]: created_repo3
Out[9]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test_upload_repo_2',
 'last_introspection_error': '',
 'last_introspection_status': 'Pending',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test_upload_repo_2',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': '',
 'uuid': 'be85f598-62a2-427a-bfaa-e14becd4b4c0'}
```

To list upload repos at present it is required to use  `origin='upload'`:

```
In [20]: app.content_sources.rest_client.repositories_api.list_repositories(origin='upload')
2024-07-15 16:00:51.627 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('origin', 'upload')]
Out[20]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test_upload_repo_1',
           'last_introspection_error': '',
           'last_introspection_status': 'Pending',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test_upload_repo_1',
           'org_id': '3340851',
           'origin': 'upload',
           'package_count': 0,
           'snapshot': False,
           'status': 'Pending',
           'url': '',
           'uuid': 'e1474a4a-c4ba-48d8-8bbf-9b513d240711'},
          {'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test_upload_repo_2',
           'last_introspection_error': '',
           'last_introspection_status': 'Pending',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test_upload_repo_3',
           'org_id': '3340851',
           'origin': 'upload',
           'package_count': 0,
           'snapshot': True,
           'status': 'Pending',
           'url': '',
           'uuid': 'be85f598-62a2-427a-bfaa-e14becd4b4c0'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [21]: 
```

 

---

## Reviews

### Review by @jlsherrill - Commented on July 08, 2024 at 05:11 PM UTC

### Review by @rverdile - Commented on July 08, 2024 at 05:51 PM UTC

### Review by @rverdile - Commented on July 08, 2024 at 05:56 PM UTC

### Review by @rverdile - Commented on July 08, 2024 at 06:00 PM UTC

### Review by @jlsherrill - Commented on July 08, 2024 at 06:19 PM UTC

### Review by @rverdile - Commented on July 09, 2024 at 05:30 PM UTC

I found these two issues. Everything else seems to be working well.

### Review by @jlsherrill - Commented on July 10, 2024 at 03:17 PM UTC

### Review by @rverdile - Commented on July 10, 2024 at 05:34 PM UTC

### Review by @jlsherrill - Commented on July 12, 2024 at 02:33 PM UTC

### Review by @rverdile - Approved on July 12, 2024 at 05:47 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/728*
