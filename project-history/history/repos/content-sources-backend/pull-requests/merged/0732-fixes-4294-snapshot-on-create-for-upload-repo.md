---
type: pull_request
number: 732
title: "Fixes 4294: snapshot on create for upload repo"
state: merged
author: jlsherrill
created: 2024-07-08T14:53:53Z
updated: 2024-07-18T11:30:21Z
closed: 2024-07-18T11:07:24Z
merged: 2024-07-18T11:07:24Z
base_branch: main
head_branch: 4294
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/732
---

# Pull Request #732: Fixes 4294: snapshot on create for upload repo

**Author**: @jlsherrill
**Created**: July 08, 2024 at 02:53 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4294`

## Description

## Summary

Run the snapshot task at creation of upload repositories.  
Modify the snapshot task to not create a remote and not perform a sync during the snapshot task.  Instead it will grab the latest version href from the repository in pulp (which should be empty version 0)

## Testing steps

1. create an upload repo:
```
####
POST http://localhost:8000/api/content-sources/v1.0/repositories/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json

{
  "name": "test",
  "url": "",
  "origin": "upload",
  "snapshot": true
}
```

2.  Fetch the repository (GET /repositories/?origin=upload), you should see a snapshot task be created.  Once that is done, you can fetch the snapshots for the repository (GET /repositories/UUID/snapshots)
3. From the repository data, grab the latest snapshot url, and fetch that:
```
curl http://localhost:8080/pulp/content/f5cc5c96/4d19b9f2-ecb2-41d5-8844-cea5f9a701fe/8419ed86-9e49-4ebc-a34e-0158e67d79be/
```


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 08, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4294

### Comment by @jlsherrill on July 09, 2024 at 01:10 PM UTC

requires https://github.com/content-services/content-sources-backend/pull/728 to be merged first, will rebase once that happens

### Comment by @xbhouse on July 10, 2024 at 01:56 PM UTC

it looks like the combined status on upload repos will always be pending since it takes into account introspection and that will never occur. maybe introspection should have a different initial status than pending for upload repos? either way the combined status logic in dao/repository_configs.go will need to be updated. let me know if you want me to file a ticket for any of these things

### Comment by @jlsherrill on July 12, 2024 at 04:46 PM UTC

good find!  Updated and added to a test 

### Comment by @jlsherrill on July 12, 2024 at 05:00 PM UTC

oh and i realized the api docs for the list repositories status filter was wrong,  It had all lowercase examples (valid) instead of the camel case that is required:  (Valid).  So i fixed that

### Comment by @xbhouse on July 12, 2024 at 06:18 PM UTC

nice, status looks good now and filtering still works! is it expected to only see upload repos listed if i explicitly set origin=upload? eg `/repositories/?status=Valid` just returns repos with external origin, so i have to use `/repositories/?status=Valid&origin=upload`

### Comment by @xbhouse on July 12, 2024 at 06:22 PM UTC

do we expect users to bulk create upload repos? when i bulk create a couple upload repos, in the response i see a blank introspection status so the combined status is Unknown. it eventually switches to Valid somewhere, because listing the repos a few seconds later shows them all as Valid

EDIT: origin is also set to empty in the response on bulk create, but this is filled in as well somewhere further down in the creation process

### Comment by @xbhouse on July 12, 2024 at 06:23 PM UTC

everything else is working as expected and looks good :) 

### Comment by @jlsherrill on July 12, 2024 at 06:42 PM UTC

> nice, status looks good now and filtering still works! is it expected to only see upload repos listed if i explicitly set origin=upload? eg `/repositories/?status=Valid` just returns repos with external origin, so i have to use `/repositories/?status=Valid&origin=upload`

yes, this is controlled by the NewRepositoryFiltering 'feature'.  We should be able to just get rid of that now and default to showing all repos when hitting /repositories/.  Let me file a task!

### Comment by @swadeley on July 15, 2024 at 03:50 PM UTC

@jlsherrill Rebase please

### Comment by @jlsherrill on July 16, 2024 at 03:24 PM UTC

squashed and rebased

### Comment by @swadeley on July 17, 2024 at 08:37 AM UTC

Hi @jlsherrill 
I see:
```make: *** [mk/go-rules.mk:58: test-db-migrations] Error 1```
is that unrelated to your changes?

### Comment by @swadeley on July 17, 2024 at 11:33 AM UTC

Hi

I deployed this PR's image. I generated new API docs.

I created upload repo:
```
In [6]:     repo_dict_2 = dict(
   ...:         name="test_upload_repo_1",
   ...:         origin="upload",
   ...:         snapshot= True,
   ...:      )

In [7]: created_repo2 = app.content_sources.rest_client.repositories_api.create_repository(repo_dict_2)
2024-07-17 11:21:45.068 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>----<>, params=[]

In [8]: created_repo2
Out[8]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test_upload_repo_1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-07-17T10:21:44Z',
 'last_snapshot_task_uuid': '5e13b12d-a535-44ef-83cd-4f71d2d55235',
 'last_success_introspection_time': '2024-07-17T10:21:44Z',
 'last_update_introspection_time': '2024-07-17T10:21:44Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test_upload_repo_1',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '98216dcd-4c6b-4a12-9b1f-6fd31b59324f'}

In [9]: app.content_sources.rest_client.repositories_api.get_repository(created_repo2.uuid)
2024-07-17 11:23:12.515 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[9]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test_upload_repo_1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-07-17T10:21:44Z',
 'last_snapshot': {'added_counts': {},
                   'content_counts': {},
                   'created_at': '2024-07-17T10:21:54.054727Z',
                   'removed_counts': {},
                   'repository_path': '5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa',
                   'url': 'http://pulp-content-svc:24816/api/pulp-content/5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa/',
                   'uuid': '1ac26016-070d-435f-927f-71d82d51f3c5'},
 'last_snapshot_task': {'created_at': '2024-07-17T10:21:45Z',
                        'ended_at': '2024-07-17T10:21:54Z',
                        'error': '',
                        'org_id': '3340851',
                        'repository_name': 'test_upload_repo_1',
                        'repository_uuid': '98216dcd-4c6b-4a12-9b1f-6fd31b59324f',
                        'status': 'completed',
                        'type': 'snapshot',
                        'uuid': '5e13b12d-a535-44ef-83cd-4f71d2d55235'},
 'last_snapshot_task_uuid': '5e13b12d-a535-44ef-83cd-4f71d2d55235',
 'last_snapshot_uuid': '1ac26016-070d-435f-927f-71d82d51f3c5',
 'last_success_introspection_time': '2024-07-17T10:21:44Z',
 'last_update_introspection_time': '2024-07-17T10:21:44Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test_upload_repo_1',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': '98216dcd-4c6b-4a12-9b1f-6fd31b59324f'}

In [10]: app.content_sources.rest_client.snapshots_api.list_snapshots(created_repo2.uuid)
2024-07-17 11:26:07.611 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[10]: 
{'data': [{'added_counts': {},
           'content_counts': {},
           'created_at': '2024-07-17T10:21:54.054727Z',
           'removed_counts': {},
           'repository_path': '5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa',
           'url': 'http://pulp-content-svc:24816/api/pulp-content/5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa/',
           'uuid': '1ac26016-070d-435f-927f-71d82d51f3c5'}],
 'links': {'first': '/api/content-sources/v1/repositories/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

```

I logged into another pod and tried to get the snapshot:

```
oc rsh -n $(oc project -q) pulp-content-d66f5b6bb-dqgsb
Defaulted container "pulp-content" out of: pulp-content, crcauth, wait-on-migrations-init (init)
sh-5.1$ curl 'http://pulp-content-svc:24816/api/pulp-content/5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa/'

<html>
<head><title>Index of /api/pulp-content/5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa/</title></head>
<body bgcolor="white">
<h1>Index of /api/pulp-content/5786a863/98216dcd-4c6b-4a12-9b1f-6fd31b59324f/8d241c98-4f9d-4b93-a13f-bf30eb66c4aa/</h1>
<hr><pre><a href="../">../</a>
<a href="repodata/">repodata/</a>                                                                                           17-Jul-2024 10:21  1.5 kB
</pre><hr></body>

```

### Comment by @swadeley on July 18, 2024 at 11:06 AM UTC

Hi

Checking for zero packages as expected in newly created upload repo. 

I had to do this in a debug pod (because you need write access to store the file):
```
sh-4.4$ curl -L -O 'http://pulp-content-svc:24816/api/pulp-content/ba64179c/53b88156-d117-4729-bf37-67e98f164011/dd8de72a-4043-45f4-acf6-ce1ecceabad2/repodata/1cb61ea996355add02b1426ed4c1780ea75ce0c04c5d1107c025c3fbd7d8bcae-primary.xml.gz'

sh-4.4$ gunzip 1cb61ea996355add02b1426ed4c1780ea75ce0c04c5d1107c025c3fbd7d8bcae-primary.xml.gz

sh-4.4$ grep packages 1cb61ea996355add02b1426ed4c1780ea75ce0c04c5d1107c025c3fbd7d8bcae-primary.xml
<metadata xmlns="http://linux.duke.edu/metadata/common" xmlns:rpm="http://linux.duke.edu/metadata/rpm" packages="0">
```

---

## Reviews

### Review by @xbhouse - Approved on July 15, 2024 at 01:46 PM UTC

this looks good! will just need a rebase once [this](https://github.com/content-services/content-sources-backend/pull/728) is merged in as you mentioned

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/732*
