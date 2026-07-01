---
type: pull_request
number: 573
title: "Fixes 3577: creation prevented on previously deleted item"
state: merged
author: Andrewgdewar
created: 2024-02-12T21:45:56Z
updated: 2024-02-16T14:00:32Z
closed: 2024-02-16T13:48:20Z
merged: 2024-02-16T13:48:20Z
base_branch: main
head_branch: HMS-3577
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/573
---

# Pull Request #573: Fixes 3577: creation prevented on previously deleted item

**Author**: @Andrewgdewar
**Created**: February 12, 2024 at 09:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3577`

## Description

## Summary

This change updates the unique key value on the templates table to only encompass items that have their "deleted_at" set to null, IE not "soft-deleted". 

Thus preventing a dupe-key error that would present when a previously "soft-deleted" template shared a name with a template that was being created.

## Testing steps

- Create a template, 
- Delete it, 
- Create another template with the same name, 
- Prior to this change, the api would error with "Template with this name already exists".

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 12, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-3577

### Comment by @xbhouse on February 14, 2024 at 08:50 PM UTC

reproduced the error with these steps:

- import RH repos and snapshot all of them (`make repos-import && go run cmd/external-repos/main.go nightly-jobs`)

while those tasks are running:
- create template
- delete template
- create template with name of the one that was just deleted

not seeing this error anymore with your changes :) looks good! 

### Comment by @swadeley on February 15, 2024 at 04:06 PM UTC

Hi

I deployed:
content-sources-backend=pr-573-latest 

but I get error:
`HTTP response body: {"errors":[{"status":400,"title":"Error creating template","detail":"Template with this name already belongs to organization"}]}`


My steps:

```
In [4]: app.content_sources.rest_client.templates_api.create_template(dict(
   ...:    name="my template one",
   ...:    repository_uuids=["daa4e8e4-412b-4052-9654-b7f557b95a72"],
   ...:    description="my first draft",
   ...:    arch="x86_64",
   ...:    version="7"
   ...:    ))
2024-02-15 16:01:00.089 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, <>, params=[]
Out[4]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my first draft',
 'name': 'my template one',
 'org_id': '16781244',
 'repository_uuids': ['daa4e8e4-412b-4052-9654-b7f557b95a72'],
 'uuid': '7bee45e6-d632-460e-96b9-bd0985a8715e',
 'version': '7'}

In [5]: app.content_sources.rest_client.templates_api.delete_template('7bee45e6-d632-460e-96b9-bd0985a8715e')
2024-02-15 16:01:18.987 [    INFO] [iqe.base.rest_client] REST: METHOD=DELETE, request_id=<>, <>, params=[]

In [6]: app.content_sources.rest_client.templates_api.list_templates(sort_by='name')
2024-02-15 16:01:42.706 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, <>, params=[('sort_by', 'name')]
Out[6]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name'},
 'meta': {'count': 0, 'limit': 100, 'offset': 0}}

In [7]: app.content_sources.rest_client.templates_api.create_template(dict(
   ...:    name="my template one",
   ...:    repository_uuids=["daa4e8e4-412b-4052-9654-b7f557b95a72"],
   ...:    description="my first draft",
   ...:    arch="x86_64",
   ...:    version="7"
   ...:    ))
2024-02-15 16:01:52.448 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, <>, params=[]

<traceback>

```
```
ApiException: (400)
Reason: Bad Request
```

### Comment by @swadeley on February 16, 2024 at 10:14 AM UTC

/retest

### Comment by @swadeley on February 16, 2024 at 12:00 PM UTC

Hi, going to retest soon.

### Comment by @swadeley on February 16, 2024 at 01:47 PM UTC

Hi

it works!

```In [1]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='fedoratest-snapshot-true',
   ...:            module_hotfixes=False,
   ...:            url='http://stephenw.fedorapeople.org/multirepos/2/repo02'
   ...:       )

In [2]: app.content_sources.rest_client.repositories_api.create_repository(repo)
2024-02-16 13:43:56.909 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7f7291319940> object....with url https://e<>.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
2024-02-16 13:43:56.910 [    INFO] [iqe.base.auth] Available auth types: ['basic', 'cert', 'identity', 'jwt', 'uht']
2024-02-16 13:43:56.910 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-02-16 13:43:58.176 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[2]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_snapshot_task_uuid': '01b1f6f5-062f-4487-b5a2-d609c39276a3',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'fedoratest-snapshot-true',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'http://stephenw.fedorapeople.org/multirepos/2/repo02/',
 'uuid': '61d1501c-f17d-4512-9571-86f867bc845b'}

In [3]: app.content_sources.rest_client.templates_api.create_template(dict(
   ...:    ...:    name="my template one",
   ...:    ...:    repository_uuids=["61d1501c-f17d-4512-9571-86f867bc845b"],
   ...:    ...:    description="my first draft",
   ...:    ...:    arch="x86_64",
   ...:    ...:    version="7"
   ...:    ...:    ))
2024-02-16 13:44:20.838 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[3]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my first draft',
 'name': 'my template one',
 'org_id': '3340851',
 'repository_uuids': ['61d1501c-f17d-4512-9571-86f867bc845b'],
 'uuid': '8f7e5401-3806-468d-b8ca-4ab29913d0e5',
 'version': '7'}

In [4]: app.content_sources.rest_client.templates_api.delete_template('8f7e5401-3806-468d-b8ca-4ab29913d0e5')
2024-02-16 13:44:46.911 [    INFO] [iqe.base.rest_client] REST: METHOD=DELETE, request_id=<>, params=[]

In [5]: app.content_sources.rest_client.templates_api.create_template(dict(
   ...:    ...:    name="my template one",
   ...:    ...:    repository_uuids=["61d1501c-f17d-4512-9571-86f867bc845b"],
   ...:    ...:    description="my first draft",
   ...:    ...:    arch="x86_64",
   ...:    ...:    version="7"
   ...:    ...:    ))
2024-02-16 13:45:00.949 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[5]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my first draft',
 'name': 'my template one',
 'org_id': '3340851',
 'repository_uuids': ['61d1501c-f17d-4512-9571-86f867bc845b'],
 'uuid': '35f79f68-b1e1-4566-99a6-ca7cdc6b6881',
 'version': '7'}

In [6]: app.content_sources.rest_client.templates_api.list_templates(sort_by='name')
2024-02-16 13:45:21.526 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'name')]
Out[6]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my first draft',
           'name': 'my template one',
           'org_id': '3340851',
           'repository_uuids': ['61d1501c-f17d-4512-9571-86f867bc845b'],
           'uuid': '35f79f68-b1e1-4566-99a6-ca7cdc6b6881',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [7]: 
```

---

## Reviews

### Review by @xbhouse - Approved on February 14, 2024 at 08:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/573*
