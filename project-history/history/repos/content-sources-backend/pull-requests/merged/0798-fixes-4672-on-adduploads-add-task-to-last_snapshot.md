---
type: pull_request
number: 798
title: "Fixes 4672: on AddUploads add task to last_snapshot"
state: merged
author: jlsherrill
created: 2024-09-05T18:09:34Z
updated: 2024-09-10T19:00:37Z
closed: 2024-09-10T18:38:55Z
merged: 2024-09-10T18:38:55Z
base_branch: main
head_branch: 4672
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/798
---

# Pull Request #798: Fixes 4672: on AddUploads add task to last_snapshot

**Author**: @jlsherrill
**Created**: September 05, 2024 at 06:09 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4672`

## Description

## Summary

Adds the AddUploads task as the last_snapshot on the repository

## Testing steps

1.  create an upload repository


2. using its uuid, upload an rpm to it:
```
IDENTITY_HEADER="eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" python ./scripts/uploads.py  30d60f06-b45d-44fa-aece-869d55edf002  ~/rpms/zebra-0.1-2.noarch.rpm 
```

3.  Fetch the repository and see a new last_snapshot_task



---

## Discussion

### Comment by @jlsherrill on September 05, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4672

### Comment by @xbhouse on September 06, 2024 at 04:01 PM UTC

oh and i see the same unit test failures in another PR, are they related to the viper issues Ryan mentioned? 

### Comment by @mayurilahane on September 06, 2024 at 07:18 PM UTC

/retest

### Comment by @jlsherrill on September 06, 2024 at 07:59 PM UTC

test should be fixed now via https://github.com/content-services/content-sources-backend/pull/797

### Comment by @mayurilahane on September 09, 2024 at 02:09 PM UTC

/retest

### Comment by @swadeley on September 10, 2024 at 07:59 AM UTC

/retest

### Comment by @swadeley on September 10, 2024 at 09:57 AM UTC

/retest

### Comment by @swadeley on September 10, 2024 at 01:18 PM UTC

Hi

I followed steps from PR 781 to create the repo and chunk
https://github.com/content-services/content-sources-backend/pull/781#issuecomment-2317127995

```
In [22]: app.content_sources.rest_client.repositories_api.create_repository(dict(name='test1', origin='upload', snapshot=True))
    ...: 
2024-09-10 13:56:03.808 [    INFO] [iqe.base.rest_client] REST: POST https://console.stage.redhat.com/api/content-sources/v1/repositories/ with query params [] and x-rh-insights-request-id=<>
Out[22]: 
{'account_id': '11772588',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-09-10T12:56:03Z',
 'last_snapshot_task_uuid': '15e35d0f-ba3e-4eab-96db-760407417fad',
 'last_success_introspection_time': '2024-09-10T12:56:03Z',
 'last_update_introspection_time': '2024-09-10T12:56:03Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test1',
 'org_id': '17799759',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a'}

In [23]: body = { 'uploads': [{'uuid': '0191dba8-ab5e-7a99-800a-1a90b24633d1', 'sha256':'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef57677065'}], 'artifacts': [] }

In [24]: app.content_sources.rest_client.repositories_api.add_upload('6b390f55-6fa8-4132-ad80-a4edd000f28a', body)
2024-09-10 13:58:56.166 [    INFO] [iqe.base.rest_client] REST: POST https://console.stage.redhat.com/api/content-sources/v1/repositories/6b390f55-6fa8-4132-ad80-a4edd000f28a/add_uploads/ with query params [] and x-rh-insights-request-id=<>
Out[24]: 
{'created_at': '2024-09-10T12:58:55Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '',
 'error': '',
 'org_id': '17799759',
 'repository_name': 'test1',
 'repository_uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a',
 'status': 'pending',
 'type': 'add-uploads-repository',
 'uuid': '6e0e69bc-ae23-45b5-85ff-373fe6c3cca8'}


In [25]: app.content_sources.rest_client.repositories_api.get_repository('6b390f55-6fa8-4132-ad80-a4edd000f28a')
2024-09-10 13:59:25.537 [    INFO] [iqe.base.rest_client] REST: GET https://console.stage.redhat.com/api/content-sources/v1/repositories/6b390f55-6fa8-4132-ad80-a4edd000f28a with query params [] and x-rh-insights-request-id=<>
Out[25]: 
{'account_id': '11772588',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-09-10T12:59:09Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2024-09-10T12:59:08.655808Z',
                   'removed_counts': {},
                   'repository_path': 'd03ec5aa/6b390f55-6fa8-4132-ad80-a4edd000f28a/ad35564f-4e4c-4bbd-bf91-60e65bd9f96d',
                   'url': 'https://cert.console.stage.redhat.com/api/pulp-content/d03ec5aa/6b390f55-6fa8-4132-ad80-a4edd000f28a/ad35564f-4e4c-4bbd-bf91-60e65bd9f96d/',
                   'uuid': '318e73c1-970e-49a1-81c8-255b1b8f7253'},
 'last_snapshot_task': {'created_at': '2024-09-10T12:56:03Z',
                        'dependencies': [],
                        'dependents': [],
                        'ended_at': '2024-09-10T12:56:13Z',
                        'error': '',
                        'org_id': '17799759',
                        'repository_name': 'test1',
                        'repository_uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a',
                        'status': 'completed',
                        'type': 'snapshot',
                        'uuid': '15e35d0f-ba3e-4eab-96db-760407417fad'},
 'last_snapshot_task_uuid': '15e35d0f-ba3e-4eab-96db-760407417fad',
 'last_snapshot_uuid': '318e73c1-970e-49a1-81c8-255b1b8f7253',
 'last_success_introspection_time': '2024-09-10T12:59:09Z',
 'last_update_introspection_time': '2024-09-10T12:59:09Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test1',
 'org_id': '17799759',
 'origin': 'upload',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a'}

In [26]: 

In [26]: 

In [26]: app.content_sources.rest_client.repositories_api.get_repository('6b390f55-6fa8-4132-ad80-a4edd000f28a')
2024-09-10 14:05:34.643 [    INFO] [iqe.base.rest_client] REST: GET https://console.stage.redhat.com/api/content-sources/v1/repositories/6b390f55-6fa8-4132-ad80-a4edd000f28a with query params [] and x-rh-insights-request-id=<>
Out[26]: 
{'account_id': '11772588',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-09-10T12:59:09Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2024-09-10T12:59:08.655808Z',
                   'removed_counts': {},
                   'repository_path': 'd03ec5aa/6b390f55-6fa8-4132-ad80-a4edd000f28a/ad35564f-4e4c-4bbd-bf91-60e65bd9f96d',
                   'url': 'https://cert.console.stage.redhat.com/api/pulp-content/d03ec5aa/6b390f55-6fa8-4132-ad80-a4edd000f28a/ad35564f-4e4c-4bbd-bf91-60e65bd9f96d/',
                   'uuid': '318e73c1-970e-49a1-81c8-255b1b8f7253'},
 'last_snapshot_task': {'created_at': '2024-09-10T12:56:03Z',
                        'dependencies': [],
                        'dependents': [],
                        'ended_at': '2024-09-10T12:56:13Z',
                        'error': '',
                        'org_id': '17799759',
                        'repository_name': 'test1',
                        'repository_uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a',
                        'status': 'completed',
                        'type': 'snapshot',
                        'uuid': '15e35d0f-ba3e-4eab-96db-760407417fad'},
 'last_snapshot_task_uuid': '15e35d0f-ba3e-4eab-96db-760407417fad',
 'last_snapshot_uuid': '318e73c1-970e-49a1-81c8-255b1b8f7253',
 'last_success_introspection_time': '2024-09-10T12:59:09Z',
 'last_update_introspection_time': '2024-09-10T12:59:09Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test1',
 'org_id': '17799759',
 'origin': 'upload',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a'}

In [27]: app.content_sources.rest_client.repositories_api.get_repository('6b390f55-6fa8-4132-ad80-a4edd000f28a')['last_snapshot_task']
2024-09-10 14:13:07.508 [    INFO] [iqe.base.rest_client] REST: GET https://console.stage.redhat.com/api/content-sources/v1/repositories/6b390f55-6fa8-4132-ad80-a4edd000f28a with query params [] and x-rh-insights-request-id=<>
Out[27]: 
{'created_at': '2024-09-10T12:56:03Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '2024-09-10T12:56:13Z',
 'error': '',
 'org_id': '17799759',
 'repository_name': 'test1',
 'repository_uuid': '6b390f55-6fa8-4132-ad80-a4edd000f28a',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '15e35d0f-ba3e-4eab-96db-760407417fad'}

In [28]: 


```

The repo was created at
 'last_introspection_time': '2024-09-10T12:56:03Z',

The upload is at
{'created_at': '2024-09-10T12:58:55Z',

The last_snapshot_task shows:
{'created_at': '2024-09-10T12:56:03Z',

Two minutes before the updoad ?

### Comment by @swadeley on September 10, 2024 at 02:37 PM UTC

Hi

ignore the above. I think there was an environment issue, wrong shell command. Retesting with latest image.

### Comment by @swadeley on September 10, 2024 at 03:05 PM UTC

Hi

I tested again with frontend PR 327 deployed.
I created the upload repo in the UI.

```
In [3]: app.content_sources.rest_client.repositories_api.list_repositories(
   ...:         origin="upload"
   ...:      )
2024-09-10 15:50:20.806 [    INFO] [iqe.base.rest_client] REST: GET https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [('origin', 'upload')] and x-rh-insights-request-id=<>
Out[3]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'upload_test',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-09-10T14:50:13Z',
           'last_snapshot_task': {'created_at': '2024-09-10T14:50:13Z',
                                  'dependencies': [],
                                  'dependents': [],
                                  'ended_at': '',
                                  'error': '',
                                  'org_id': '3340851',
                                  'repository_name': 'upload-test',
                                  'repository_uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090',
                                  'status': 'running',
                                  'type': 'snapshot',
                                  'uuid': '5058606b-9228-4fbe-aa83-07e5e12d845b'},
           'last_snapshot_task_uuid': '5058606b-9228-4fbe-aa83-07e5e12d845b',
           'last_success_introspection_time': '2024-09-10T14:50:13Z',
           'last_update_introspection_time': '2024-09-10T14:50:13Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'upload-test',
           'org_id': '3340851',
           'origin': 'upload',
           'package_count': 0,
           'snapshot': True,
           'status': 'Pending',
           'url': '',
           'uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [4]: app.content_sources.rest_client.repositories_api.create_upload({"size": 7036})
   ...: 
2024-09-10 15:54:33.696 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'created': '2024-09-10T14:54:33.543576Z',
 'last_updated': '2024-09-10T14:54:33.543601Z',
 'size': 7036,
 'upload_uuid': '0191dc6d-de46-7624-9c76-c49ee0ee523a'}



In [7]: app.content_sources.rest_client.repositories_api.add_upload('a63e740e-648e-4cd2-aec3-c335ee2ac090', body)
2024-09-10 15:57:33.046 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/a63e740e-648e-4cd2-aec3-c335ee2ac090/add_uploads/ with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'created_at': '2024-09-10T14:57:32Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'upload-test',
 'repository_uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090',
 'status': 'pending',
 'type': 'add-uploads-repository',
 'uuid': '8fb57c11-8875-40a3-b236-272e40b7f270'}

In [8]: app.content_sources.rest_client.repositories_api.get_repository('a63e740e-648e-4cd2-aec3-c335ee2ac090')['last_snapshot_task']
2024-09-10 15:59:17.221 [    INFO] [iqe.base.rest_client] REST: GET https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/a63e740e-648e-4cd2-aec3-c335ee2ac090 with query params [] and x-rh-insights-request-id=<>
Out[8]: 
{'created_at': '2024-09-10T14:57:32Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '2024-09-10T14:57:34Z',
 'error': 'could not convert uploads to artifacts finish upload task failed '
          "unexpectedly : domain 39862113 had Pulp Task error 'traceback:   "
          'File '
          '"/usr/local/lib/pulp/lib64/python3.9/site-packages/pulpcore/tasking/tasks.py", '
          'line 75, in _execute_task\n'
          '    result = func(*args, **kwargs)\n'
          '  File '
          '"/usr/local/lib/pulp/lib64/python3.9/site-packages/pulpcore/app/tasks/upload.py", '
          'line 37, in commit\n'
          '    serializer.is_valid(raise_exception=True)\n'
          '  File '
          '"/usr/local/lib/pulp/lib64/python3.9/site-packages/rest_framework/serializers.py", '
          'line 231, in is_valid\n'
          '    raise ValidationError(self.errors)\n'
          ".  description: {'non_field_errors': [ErrorDetail(string='The "
          "sha256 checksum did not match.', code='invalid')]}.  '",
 'org_id': '3340851',
 'repository_name': 'upload-test',
 'repository_uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090',
 'status': 'failed',
 'type': 'add-uploads-repository',
 'uuid': '8fb57c11-8875-40a3-b236-272e40b7f270'}

In [9]: app.content_sources.rest_client.repositories_api.get_repository('a63e740e-648e-4cd2-aec3-c335ee2ac090')
2024-09-10 15:59:28.624 [    INFO] [iqe.base.rest_client] REST: GET https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/a63e740e-648e-4cd2-aec3-c335ee2ac090 with query params [] and x-rh-insights-request-id=<>
Out[9]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'upload_test',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-09-10T14:50:13Z',
 'last_snapshot': {'added_counts': {},
                   'content_counts': {},
                   'created_at': '2024-09-10T14:50:23.007637Z',
                   'removed_counts': {},
                   'repository_path': '39862113/a63e740e-648e-4cd2-aec3-c335ee2ac090/5c164b28-822e-499a-b41c-bb5f362bb8f3',
                   'url': 'http://pulp-content:8000/api/pulp-content/39862113/a63e740e-648e-4cd2-aec3-c335ee2ac090/5c164b28-822e-499a-b41c-bb5f362bb8f3/',
                   'uuid': '0245521f-b968-48c6-a6a2-a99124eb8efb'},
 'last_snapshot_task': {'created_at': '2024-09-10T14:57:32Z',
                        'dependencies': [],
                        'dependents': [],
                        'ended_at': '2024-09-10T14:57:34Z',
                        'error': 'could not convert uploads to artifacts '
                                 'finish upload task failed unexpectedly : '
                                 'domain 39862113 had Pulp Task error '
                                 "'traceback:   File "
                                 '"/usr/local/lib/pulp/lib64/python3.9/site-packages/pulpcore/tasking/tasks.py", '
                                 'line 75, in _execute_task\n'
                                 '    result = func(*args, **kwargs)\n'
                                 '  File '
                                 '"/usr/local/lib/pulp/lib64/python3.9/site-packages/pulpcore/app/tasks/upload.py", '
                                 'line 37, in commit\n'
                                 '    '
                                 'serializer.is_valid(raise_exception=True)\n'
                                 '  File '
                                 '"/usr/local/lib/pulp/lib64/python3.9/site-packages/rest_framework/serializers.py", '
                                 'line 231, in is_valid\n'
                                 '    raise ValidationError(self.errors)\n'
                                 ".  description: {'non_field_errors': "
                                 "[ErrorDetail(string='The sha256 checksum did "
                                 "not match.', code='invalid')]}.  '",
                        'org_id': '3340851',
                        'repository_name': 'upload-test',
                        'repository_uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090',
                        'status': 'failed',
                        'type': 'add-uploads-repository',
                        'uuid': '8fb57c11-8875-40a3-b236-272e40b7f270'},
 'last_snapshot_task_uuid': '8fb57c11-8875-40a3-b236-272e40b7f270',
 'last_snapshot_uuid': '0245521f-b968-48c6-a6a2-a99124eb8efb',
 'last_success_introspection_time': '2024-09-10T14:50:13Z',
 'last_update_introspection_time': '2024-09-10T14:50:13Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'upload-test',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Unavailable',
 'url': '',
 'uuid': 'a63e740e-648e-4cd2-aec3-c335ee2ac090'}

In [10]: 



```



### Comment by @swadeley on September 10, 2024 at 03:23 PM UTC

Hi

Further to the above, I see I forgot the file upload after creating the chunk:

```
In [10]: with open('/tmp/testaa', 'rb') as file:
    ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('0191dc6d-de46-7624-9c76-c49ee0ee523a', 'bytes 0-7055/*', file, 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef57677
    ...: 065')
    ...: 

```

but now I get ISE.

I will create new repo and test again.

### Comment by @swadeley on September 10, 2024 at 05:35 PM UTC

Hi @jlsherrill 

with backend=pr-798-541b56d   frontend=pr-327-7061d52

I created a new repo in the UI,  'label': 'test_two',  and tried again, now get ISE after file upload step:

```
In [2]: app.content_sources.rest_client.repositories_api.list_repositories(
   ...:         origin="upload"
   ...:      )
2024-09-10 18:25:22.737 [    INFO] [iqe.base.rest_client] REST: GET https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [('origin', 'upload')] and x-rh-insights-request-id=<>
Out[2]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test_two',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-09-10T17:25:17Z',
           'last_snapshot_task': {'created_at': '2024-09-10T17:25:17Z',
                                  'dependencies': [],
                                  'dependents': [],
                                  'ended_at': '',
                                  'error': '',
                                  'org_id': '3340851',
                                  'repository_name': 'test-two',
                                  'repository_uuid': '40f0fcf2-1744-4398-b3fd-fd1d8784b54d',
                                  'status': 'running',
                                  'type': 'snapshot',
                                  'uuid': 'a011e63a-f914-4bd5-b535-0e6174bc88bb'},
           'last_snapshot_task_uuid': 'a011e63a-f914-4bd5-b535-0e6174bc88bb',
           'last_success_introspection_time': '2024-09-10T17:25:17Z',
           'last_update_introspection_time': '2024-09-10T17:25:17Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test-two',
           'org_id': '3340851',
           'origin': 'upload',
           'package_count': 0,
           'snapshot': True,
           'status': 'Pending',
           'url': '',
           'uuid': '40f0fcf2-1744-4398-b3fd-fd1d8784b54d'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=upload'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [3]: app.content_sources.rest_client.repositories_api.create_upload({"size": 7036})
2024-09-10 18:26:42.026 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ with query params [] and x-rh-insights-request-id=<>
Out[3]: 
{'created': '2024-09-10T17:26:41.974331Z',
 'last_updated': '2024-09-10T17:26:41.974351Z',
 'size': 7036,
 'upload_uuid': '0191dcf9-2835-7075-9c14-06b6023e3915'}

In [4]: with open('/tmp/testaa', 'rb') as file:
   ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('0191dcf9-2835-7075-9c14-06b6023e3915', 'bytes 0-7055/*', file, 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef576770
   ...: 65')
   ...: 
2024-09-10 18:27:57.581 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/0191dcf9-2835-7075-9c14-06b6023e3915/upload_chunk/ with query params [] and x-rh-insights-request-id=93ac39fa-4da8-4790-8a73-5852b0130cb3
---------------------------------------------------------------------------
ServiceException                          Traceback (most recent call last)
```

oops, wrong byte count `7036`

### Comment by @swadeley on September 10, 2024 at 06:38 PM UTC

Works now:

```
In [1]: app.content_sources.rest_client.repositories_api.create_upload({"size": 7056})
Out[1]: 
{'created': '2024-09-10T18:21:22.424595Z',
 'last_updated': '2024-09-10T18:21:22.424613Z',
 'size': 7056,
 'upload_uuid': '0191dd2b-3678-7510-9558-6210d61c3ce5'}

In [2]: with open('/tmp/testaa', 'rb') as file:
   ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('0191dd2b-3678-7510-9558-6210d61c3ce5', 'bytes 0-7055/*', file, 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef576770
   ...: 65')
   ...: 
2024-09-10 19:21:58.226 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/0191dd2b-3678-7510-9558-6210d61c3ce5/upload_chunk/ with query params [] and x-rh-insights-request-id=<>

In [3]: body = { 'uploads': [{'uuid': '0191dd2b-3678-7510-9558-6210d61c3ce5', 'sha256':'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef57677065'}], 'artifacts': [] }

In [4]: app.content_sources.rest_client.repositories_api.add_upload('40f0fcf2-1744-4398-b3fd-fd1d8784b54d', body)
2024-09-10 19:23:17.490 [    INFO] [iqe.base.rest_client] REST: POST https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/40f0fcf2-1744-4398-b3fd-fd1d8784b54d/add_uploads/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'created_at': '2024-09-10T18:23:17Z',
 'dependencies': [],
 'dependents': [],
 'ended_at': '',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'test-two',
 'repository_uuid': '40f0fcf2-1744-4398-b3fd-fd1d8784b54d',
 'status': 'pending',
 'type': 'add-uploads-repository',
 'uuid': '7c5f68d1-dec8-481b-b98d-a5b48dddf5ed'}

In [5]: app.content_sources.rest_client.repositories_api.get_repository('40f0fcf2-1744-4398-b3fd-fd1d8784b54d')
2024-09-10 19:24:33.059 [    INFO] [iqe.base.rest_client] REST: GET https://ee-5yywk72x.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/40f0fcf2-1744-4398-b3fd-fd1d8784b54d with query params [] and x-rh-insights-request-id=cae13627-62b7-4763-9a44-24c4821d411e
Out[5]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test_two',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-09-10T18:23:31Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2024-09-10T18:23:30.986088Z',
                   'removed_counts': {},
                   'repository_path': '39862113/40f0fcf2-1744-4398-b3fd-fd1d8784b54d/c13da144-8b77-4e5e-b923-3f5b859f0c1c',
                   'url': 'http://pulp-content:8000/api/pulp-content/39862113/40f0fcf2-1744-4398-b3fd-fd1d8784b54d/c13da144-8b77-4e5e-b923-3f5b859f0c1c/',
                   'uuid': 'e74beb3f-7aae-41d3-a389-1b9765442f4b'},
 'last_snapshot_task': {'created_at': '2024-09-10T18:23:17Z',
                        'dependencies': [],
                        'dependents': [],
                        'ended_at': '2024-09-10T18:23:31Z',
                        'error': '',
                        'org_id': '3340851',
                        'repository_name': 'test-two',
                        'repository_uuid': '40f0fcf2-1744-4398-b3fd-fd1d8784b54d',
                        'status': 'completed',
                        'type': 'add-uploads-repository',
                        'uuid': '7c5f68d1-dec8-481b-b98d-a5b48dddf5ed'},
 'last_snapshot_task_uuid': '7c5f68d1-dec8-481b-b98d-a5b48dddf5ed',
 'last_snapshot_uuid': 'e74beb3f-7aae-41d3-a389-1b9765442f4b',
 'last_success_introspection_time': '2024-09-10T18:23:31Z',
 'last_update_introspection_time': '2024-09-10T18:23:31Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test-two',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '40f0fcf2-1744-4398-b3fd-fd1d8784b54d'}

In [6]: 

```

and we see :
last_snapshot_task': {'created_at': '2024-09-10T18:23:17Z',

which is after create_upload:
{'created': '2024-09-10T18:21:22.424595Z',

and repo was created at:
          'label': 'test_two',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-09-10T17:25:17Z',

so it is a new snapshot task.

thank you


---

## Reviews

### Review by @xbhouse - Approved on September 05, 2024 at 09:14 PM UTC

ack! tested via API and verified the status changes in the UI

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/798*
