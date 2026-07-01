---
type: pull_request
number: 510
title: "Fixes 2965: add create/fetch/list for templates"
state: merged
author: rverdile
created: 2023-12-19T17:46:34Z
updated: 2024-01-15T17:00:41Z
closed: 2024-01-15T16:58:03Z
merged: 2024-01-15T16:58:03Z
base_branch: main
head_branch: content-templates
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/510
---

# Pull Request #510: Fixes 2965: add create/fetch/list for templates

**Author**: @rverdile
**Created**: December 19, 2023 at 05:46 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `content-templates`

## Description

## Summary

- Adds create, list, and fetch endpoints for content templates
- Adds filters and sorting for name, version, and arch
- Adds search by name

## Testing steps
1. Create a repository
2. Create templates

```
POST http://localhost:8000/api/content-sources/v1.0/templates/
Content-Type: application/json
x-Rh-Identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE2Nzc5ODUyIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTY3Nzk4NTIifX19Cg==

{
  "name": "my template 5",
  "repository_uuids": ["8e94a78e-915e-4eee-8781-7592b5e3f6f4"],
  "description": "my new template",
  "arch": "x86_64",
  "version": "7"

}
```
3. List Templates with filtering, sorting, or searching
```
GET http://localhost:8000/api/content-sources/v1.0/templates/?version=7&sort_by=name
Content-Type: application/json
x-Rh-Identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE2Nzc5ODUyIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTY3Nzk4NTIifX19Cg==
```
4. Fetch Templates
```
GET http://localhost:8000/api/content-sources/v1.0/templates/7fc1ee4a-fad2-4e4c-9979-953136c6eb09
Content-Type: application/json
x-Rh-Identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE2Nzc5ODUyIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTY3Nzk4NTIifX19Cg==
```
5. Make sure to look at the JIRA card and check against the requirements there in case I missed something.


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 19, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-2965

### Comment by @swadeley on January 09, 2024 at 07:53 PM UTC

/retest

### Comment by @swadeley on January 10, 2024 at 01:50 PM UTC

@rverdile Please rebase. Thank you.

### Comment by @rverdile on January 10, 2024 at 03:50 PM UTC

rebased

### Comment by @swadeley on January 12, 2024 at 11:34 AM UTC

Hi

I get error:
`ApiAttributeError: ApiTemplateResponse has no attribute 'UUID' at ['received_data']['UUID']`

with a dict like so:
```
In [17]: app.content_sources.rest_client.templates_api.create_template(dict(
    ...:   name="my template 5",
    ...:   repository_uuids=["8b1851fc-276e-4cda-99c9-044cf384408d"],
    ...:   description="my new template",
    ...:   arch="x86_64",
    ...:   version="7"
    ...: ))
```

I checked the API, there is this under TemplateResponse:
```
                    },
                    "uuid": {
                        "type": "string"
                    },
```



### Comment by @xbhouse on January 12, 2024 at 09:05 PM UTC

i'm seeing this error message when trying to fetch a template with an invalid uuid: `error="error: code=404, title=Error fetching template, detail=Could not find repository with UUID 1f47bd1b-5785-4d41-a326-1c01d37247f3 \n" `

should probably say `Could not find template`, though i'm not sure where this error detail is set so i can't point to where it's happening :/ 

### Comment by @swadeley on January 15, 2024 at 04:56 PM UTC

Hi

looks good now, thank you all.

```
In [2]: app.content_sources.rest_client.repositories_api.create_repository(repo)
<snip>
2024-01-15 16:43:22.108 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=RaXiDaxHNlacmDGWWSpDQeEuGywYrDDC, params=[]
Out[2]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': 'https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'gpgtest-snapshot-false',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': False,
 'status': 'Pending',
 'url': 'http://jlsherrill.fedorapeople.org/fake-repos/signed/',
 'uuid': 'b259fcb1-0d24-4a41-8487-8577b68cf211'}

In [3]:  app.content_sources.rest_client.templates_api.create_template(dict(
   ...:     ...:   name="my template 5",
   ...:     ...:   repository_uuids=["b259fcb1-0d24-4a41-8487-8577b68cf211"],
   ...:     ...:   description="my new template",
   ...:     ...:   arch="x86_64",
   ...:     ...:   version="7"
   ...:     ...: ))
2024-01-15 16:44:32.633 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, <snip>
Out[3]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my new template',
 'name': 'my template 5',
 'org_id': '3340851',
 'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
 'version': '7'}

In [4]: app.content_sources.rest_client.templates_api.list_templates()
2024-01-15 16:51:13.245 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>
Out[4]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my new template',
           'name': 'my template 5',
           'org_id': '3340851',
           'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [5]: app.content_sources.rest_client.templates_api.list_templates.params_map
Out[5]: 
{'all': ['offset',
  'limit',
  'version',
  'arch',
  'name',
  'sort_by',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': [],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}

In [6]: app.content_sources.rest_client.templates_api.list_templates(arch='x86_64')
2024-01-15 16:52:03.721 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>
Out[6]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my new template',
           'name': 'my template 5',
           'org_id': '3340851',
           'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?arch=x86_64&limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?arch=x86_64&limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [7]: app.content_sources.rest_client.templates_api.list_templates(name='my template 5')
<snip>
Out[7]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my new template',
           'name': 'my template 5',
           'org_id': '3340851',
           'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&name=my+template+5&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&name=my+template+5&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [8]: app.content_sources.rest_client.templates_api.list_templates(version='7')
2024-01-15 16:52:52.916 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>
Out[8]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my new template',
           'name': 'my template 5',
           'org_id': '3340851',
           'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&version=7',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&version=7'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

<snip>
In [10]: app.content_sources.rest_client.templates_api.list_templates(sort_by='name')
2024-01-15 16:53:41.401 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('sort_by', 'name')]
Out[10]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my new template',
           'name': 'my template 5',
           'org_id': '3340851',
           'uuid': '5b586840-fde2-4a12-8101-b1b65fdae600',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [11]: 
```

---

## Reviews

### Review by @Andrewgdewar - Approved on January 05, 2024 at 04:06 PM UTC

Tested, integrated, ACK!

### Review by @xbhouse - Commented on January 12, 2024 at 09:00 PM UTC

### Review by @rverdile - Commented on January 15, 2024 at 02:46 PM UTC

### Review by @xbhouse - Approved on January 15, 2024 at 02:57 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/510*
