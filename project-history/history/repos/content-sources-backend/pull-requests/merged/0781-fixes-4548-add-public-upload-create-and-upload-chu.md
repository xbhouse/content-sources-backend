---
type: pull_request
number: 781
title: "Fixes 4548: add public upload create and upload chunk apis"
state: merged
author: xbhouse
created: 2024-08-19T16:00:06Z
updated: 2024-08-29T09:30:18Z
closed: 2024-08-29T09:23:38Z
merged: 2024-08-29T09:23:38Z
base_branch: main
head_branch: 4548
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/781
---

# Pull Request #781: Fixes 4548: add public upload create and upload chunk apis

**Author**: @xbhouse
**Created**: August 19, 2024 at 04:00 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4548`

## Description

## Summary

- Adds the public APIs for creating an upload and uploading a chunk: `POST /repositories/uploads/` and `POST /repositories/uploads/:upload_uuid/upload_chunk/`
- Modifies the existing addUploads API to use either an upload UUID or an upload href
- Modifies the python script for testing the upload process to use either the public or internal APIs

## Testing steps

- Create an upload repository with snapshot set to true
- Use the ./scripts/uploads.py script to use the public APIs to create an upload, upload the chunks, and add the uploads:
`python3 uploads.py <repo_uuid> <rpm1 rpm2 ...> --api public` (the same script should still be able to be used with the internal APIs: `python3 uploads.py <repo_uuid> <rpm1 rpm2 ...> --api internal`)
- Fetch the repository and verify the package count (should be 1)
- Fetch the repository's rpms and the snapshot's rpms and verify they include the one you just uploaded
- Access the pulp content via the pulp URL and verify the expected rpm is there
- Upload the same rpm and verify no new snapshot is created
- Upload a different rpm and verify a new snapshot is created with that rpm added
- For QE, the upload chunk API in particular may be difficult to test and I may need to change some things so this works with generated client. 

Here are some example commands in case it's helpful:

```
Create upload request:
curl -X POST 'localhost:8000/api/content-sources/v1.0/repositories/uploads/' \
-H 'x-rh-identity: <identity_header>' \
-H 'Content-Type: application/json' \
-d '{    
    "size": 7036
}'

IQE shell:
app.content_sources.rest_client.repositories_api.create_upload({"size": 7036})

Response:
{
    "upload_uuid": "01916bde-8e8a-76a6-aa34-17acdfb86ceb",
    "created": "2024-08-19T18:20:33.290458Z",
    "last_updated": "2024-08-19T18:20:33.290468Z",
    "size": 7036
}

Upload chunk request:
curl -X POST 'localhost:8000/api/content-sources/v1.0/repositories/uploads/01916bde-8e8a-76a6-aa34-17acdfb86ceb/upload_chunk/' \
-H 'x-rh-identity:  <identity_header>' \
-H 'Content-Range: bytes 0-7035/*' \
-F 'file=@"/Users/bhouse/rpms/rpm_chunkaa"' \
-F 'sha256="7e380bd4c000378f407c1a3663dedfda01248564fd7d4e33d56a1c691046bce1"'

IQE shell:
with open('/Users/bhouse/rpms/rpm_chunkaa', 'rb') as file:
   response = app.content_sources.rest_client.repositories_api.upload_chunk('01916bde-8e8a-76a6-aa34-17acdfb86ceb', 'bytes 0-7035/*', file, '7e380bd4c000378f407c1a3663dedfda01248564fd7d4e33d56a1c691046bce1')
response

Response:
{
    "upload_uuid": "01916bde-8e8a-76a6-aa34-17acdfb86ceb",
    "created": "2024-08-19T18:20:33.290458Z",
    "last_updated": "2024-08-19T18:20:33.290468Z",
    "size": 7036
}

Add uploads request:
curl -X POST 'localhost:8000/api/content-sources/v1.0/repositories/8a578521-50be-48df-b029-d7f240aa7dfa/add_uploads/' \
-H 'x-rh-identity: <identity_header>' \
-H 'Content-Type: application/json' \
-d '{
    "uploads": [
        {
            "uuid": "01916bde-8e8a-76a6-aa34-17acdfb86ceb",
            "sha256": "7e380bd4c000378f407c1a3663dedfda01248564fd7d4e33d56a1c691046bce1"
        }
    ]
}'

IQE shell:
body = { 'uploads': [{'uuid':'01916bde-8e8a-76a6-aa34-17acdfb86ceb', 'sha256':'7e380bd4c000378f407c1a3663dedfda01248564fd7d4e33d56a1c691046bce1'}], 'artifacts': [] }
app.content_sources.rest_client.repositories_api.add_upload('8a578521-50be-48df-b029-d7f240aa7dfa', body)


Response:
{
    "uuid": "29125d02-d11c-4d44-973d-878c9b963d64",
    "status": "pending",
    "created_at": "2024-08-22T11:55:13-04:00",
    "ended_at": "",
    "error": "",
    "org_id": "17684632",
    "type": "add-uploads-repository",
    "repository_name": "test",
    "repository_uuid": "8a578521-50be-48df-b029-d7f240aa7dfa",
    "dependencies": [],
    "dependents": []
}

```

---

## Discussion

### Comment by @jlsherrill on August 19, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-4548

### Comment by @xbhouse on August 26, 2024 at 12:07 PM UTC

/retest

### Comment by @swadeley on August 27, 2024 at 01:45 PM UTC

/restest

### Comment by @swadeley on August 29, 2024 at 09:22 AM UTC

Hi, its working now, thank you.

```
    In [1]:  app.content_sources.rest_client.repositories_api.create_upload({"size":7056})
<snip>
    2024-08-28 20:32:42.864 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ with query params [] and x-rh-insights-request-id=<>
    Out[1]: 
    {'created': '2024-08-28T19:32:42.773677Z',
     'last_updated': '2024-08-28T19:32:42.773695Z',
     'size': 7056,
     'upload_uuid': '01919a79-da94-78aa-a914-5e270cb1326f'}
     
     
     
    In [3]: ls -laZ /tmp/testaa
    -rw-r--r--. 1 xxx xxx unconfined_u:object_r:user_tmp_t:s0 7056 Aug 27 16:40 /tmp/testaa
     
    In [4]: with open('/tmp/testaa', 'rb') as file:
       ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('01919a79-da94-78aa-a914-5e270cb1326f', 'bytes 0-7055/*', file, 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef576770
       ...: 65')
       ...: 
    2024-08-28 20:42:46.201 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/01919a79-da94-78aa-a914-5e270cb1326f/upload_chunk/ with query params [] and x-rh-insights-request-id=<>
     
    In [5]: response
    Out[5]: 
    {'created': '2024-08-28T19:32:42.773677Z',
     'last_updated': '2024-08-28T19:32:42.773695Z',
     'size': 7056,
     'upload_uuid': '01919a79-da94-78aa-a914-5e270cb1326f'}
     
    In [6]: app.content_sources.rest_client.repositories_api.create_repository(dict(name='stephen-test1', origin='upload', snapshot=True))
    2024-08-28 20:48:42.207 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [] and x-rh-insights-request-id=<>
    Out[6]: 
    {'account_id': '0369233',
     'content_type': 'rpm',
     'distribution_arch': 'any',
     'distribution_versions': ['any'],
     'failed_introspections_count': 0,
     'gpg_key': '',
     'label': 'stephen_test1',
     'last_introspection_error': '',
     'last_introspection_status': 'Valid',
     'last_introspection_time': '2024-08-28T19:48:42Z',
     'last_snapshot_task_uuid': '4a75b82c-5b38-46d2-9495-2099d1e523c6',
     'last_success_introspection_time': '2024-08-28T19:48:42Z',
     'last_update_introspection_time': '2024-08-28T19:48:42Z',
     'metadata_verification': False,
     'module_hotfixes': False,
     'name': 'stephen-test1',
     'org_id': '3340851',
     'origin': 'upload',
     'package_count': 0,
     'snapshot': True,
     'status': 'Valid',
     'url': '',
     'uuid': 'f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680'}
     
    In [7]: body = { 'uploads': [{'uuid': '01919a79-da94-78aa-a914-5e270cb1326f', 'sha256':'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef57677065'}], 'artifacts': [] }
     
    In [8]: app.content_sources.rest_client.repositories_api.add_upload('f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680', body)
    2024-08-28 20:56:02.903 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/add_uploads/ with query params [] and x-rh-insights-request-id=<>
    Out[8]: 
    {'created_at': '2024-08-28T19:56:02Z',
     'dependencies': [],
     'dependents': [],
     'ended_at': '',
     'error': '',
     'org_id': '3340851',
     'repository_name': 'stephen-test1',
     'repository_uuid': 'f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680',
     'status': 'pending',
     'type': 'add-uploads-repository',
     'uuid': 'aec8cb88-489b-4814-9d15-b18af1311d48'}
     
    In [9]: app.content_sources.rest_client.rpms_api.list_repositories_rpms('f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680')
    2024-08-28 20:56:59.083 [    INFO] [iqe.base.rest_client] REST: GET https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/rpms with query params [] and x-rh-insights-request-id=<>
    Out[9]: 
    {'data': [{'arch': 'x86_64',
               'checksum': 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef57677065',
               'epoch': 0,
               'name': 'libreoffice24.8-writer',
               'release': '3',
               'summary': 'Writer brand module for LibreOffice 24.8.0.3',
               'uuid': 'b63cbbb3-58de-4193-864b-5d65b69d53f7',
               'version': '24.8.0.3'}],
     'links': {'first': '/api/content-sources/v1/repositories/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/rpms?limit=100&offset=0',
               'last': '/api/content-sources/v1/repositories/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/rpms?limit=100&offset=0'},
     'meta': {'count': 1, 'limit': 100, 'offset': 0}}
     
```

Use `orgin='upload'` to list the repo:

`In [16]: app.content_sources.rest_client.repositories_api.list_repositories(origin='upload')

Get the snapshot UUID

    ```
    In [12]: app.content_sources.rest_client.repositories_api.get_repository('f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680')
    2024-08-28 21:01:42.883 [    INFO] [iqe.base.rest_client] REST: GET https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680 with query params [] and x-rh-insights-request-id=<>
    Out[12]: 
    {'account_id': '0369233',
     'content_type': 'rpm',
     'distribution_arch': 'any',
     'distribution_versions': ['any'],
     'failed_introspections_count': 0,
     'gpg_key': '',
     'label': 'stephen_test1',
     'last_introspection_error': '',
     'last_introspection_status': 'Valid',
     'last_introspection_time': '2024-08-28T19:56:16Z',
     'last_snapshot': {'added_counts': {'rpm.package': 1},
                       'content_counts': {'rpm.package': 1},
                       'created_at': '2024-08-28T19:56:16.132616Z',
                       'removed_counts': {},
                       'repository_path': '928650e2/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/67bc49f7-8f30-47c0-b4ce-c33dde398068',
                       'url': 'http://pulp-content:8000/api/pulp-content/928650e2/f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680/67bc49f7-8f30-47c0-b4ce-c33dde398068/',
                       'uuid': 'fbf6ab46-2682-46bf-9e01-9ea6451587a3'},
     'last_snapshot_task': {'created_at': '2024-08-28T19:48:42Z',
                            'dependencies': [],
                            'dependents': [],
                            'ended_at': '2024-08-28T19:48:51Z',
                            'error': '',
                            'org_id': '3340851',
                            'repository_name': 'stephen-test1',
                            'repository_uuid': 'f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680',
                            'status': 'completed',
                            'type': 'snapshot',
                            'uuid': '4a75b82c-5b38-46d2-9495-2099d1e523c6'},
     'last_snapshot_task_uuid': '4a75b82c-5b38-46d2-9495-2099d1e523c6',
     'last_snapshot_uuid': 'fbf6ab46-2682-46bf-9e01-9ea6451587a3',
     'last_success_introspection_time': '2024-08-28T19:56:16Z',
     'last_update_introspection_time': '2024-08-28T19:56:16Z',
     'metadata_verification': False,
     'module_hotfixes': False,
     'name': 'stephen-test1',
     'org_id': '3340851',
     'origin': 'upload',
     'package_count': 1,
     'snapshot': True,
     'status': 'Valid',
     'url': '',
     'uuid': 'f957ad1b-d61b-4c19-b7c9-8a7a0f7ae680'}
     
    In [13]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms('fbf6ab46-2682-46bf-9e01-9ea6451587a3')
    2024-08-28 21:02:52.104 [    INFO] [iqe.base.rest_client] REST: GET https://ee-ju5heheo.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/fbf6ab46-2682-46bf-9e01-9ea6451587a3/rpms with query params [] and x-rh-insights-request-id=<>
    Out[13]: 
    {'data': [{'arch': 'x86_64',
               'epoch': '0',
               'name': 'libreoffice24.8-writer',
               'release': '3',
               'summary': 'Writer brand module for LibreOffice 24.8.0.3',
               'version': '24.8.0.3'}],
     'links': {'first': '/api/content-sources/v1/snapshots/fbf6ab46-2682-46bf-9e01-9ea6451587a3/rpms?limit=100&offset=0',
               'last': '/api/content-sources/v1/snapshots/fbf6ab46-2682-46bf-9e01-9ea6451587a3/rpms?limit=100&offset=0'},
     'meta': {'count': 1, 'limit': 100, 'offset': 0}}
     
    In [14]: 
     ```

---

## Reviews

### Review by @rverdile - Commented on August 21, 2024 at 02:06 PM UTC

### Review by @rverdile - Commented on August 22, 2024 at 03:34 PM UTC

one more small comment, otherwise looks good! tested with script and manually

### Review by @rverdile - Approved on August 22, 2024 at 04:58 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/781*
