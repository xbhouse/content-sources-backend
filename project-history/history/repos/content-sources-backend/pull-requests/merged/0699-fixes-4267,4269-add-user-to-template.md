---
type: pull_request
number: 699
title: "Fixes 4267,4269: add user to template"
state: merged
author: rverdile
created: 2024-06-10T19:40:21Z
updated: 2024-06-12T13:30:26Z
closed: 2024-06-12T13:23:16Z
merged: 2024-06-12T13:23:16Z
base_branch: main
head_branch: template-user
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/699
---

# Pull Request #699: Fixes 4267,4269: add user to template

**Author**: @rverdile
**Created**: June 10, 2024 at 07:40 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `template-user`

## Description

## Summary
Adds created_by and last_updated_by fields to template, which is the username if the user that is modifying the template

Adds created_at and updated_at fields to template response

## Testing steps
1. To test through the API, you need an identity header the the username. You can use `./scripts/headers.sh <org_id> [username]`. By default the username is "snapUser".
2. Testing through the UI will already have the username in the header
3. Creating a template should set the created_by and last_updated_by fields
4. Updating a template should set the last_updated_by field
5. Also check for created_at and updated_at in template GET requests
 

## Checklist
- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 10, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4267

### Comment by @jlsherrill on June 10, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-4269

### Comment by @swadeley on June 11, 2024 at 12:43 PM UTC

/retest

### Comment by @swadeley on June 12, 2024 at 11:02 AM UTC

Hi

Creating:

```
In [4]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-06-12 11:06:03.538 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[4]: 
{'arch': 'x86_64',
 'created_at': '2024-06-12T10:06:03.492091Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my 1st template',
 'last_updated_by': 'ephemeral-user',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
 'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
 'updated_at': '2024-06-12T10:06:03.492091Z',
 'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
 'version': '8'}
```

updating:

```
In [12]: template_update_dict = dict(repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template")


In [15]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
2024-06-12 11:34:08.173 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
Out[15]: 
{'arch': 'x86_64',
 'created_at': '2024-06-12T10:06:03.492091Z',
 'created_by': 'ephemeral-user',
 'date': '2024-06-12T10:34:08.120882Z',
 'description': 'my edited template',
 'last_updated_by': 'ephemeral-user',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
 'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
 'updated_at': '2024-06-12T10:34:08.124216Z',
 'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
 'version': '8'}
```

but I cannot change the name.
`ApiAttributeError: ApiTemplateUpdateRequest has no attribute 'name' at ['api_template_update_request']['name']
`

```
    attribute_map = {
        'date': 'date',  # noqa: E501
        'description': 'description',  # noqa: E501
        'repository_uuids': 'repository_uuids',  # noqa: E501
    }
```
 @rverdile ,  I thought it should now be possible to change the name?

### Comment by @swadeley on June 12, 2024 at 11:12 AM UTC

Hi

update the date for repo snapshots:

```
In [18]: template_update_dict = dict(date="2024-06-13T12:00:00",repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template with date change")

In [19]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
2024-06-12 12:10:24.392 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]

HTTP response body: {"errors":[{"status":400,"title":"Error binding parameters","detail":"code=400, message=parsing time \"2024-06-13T12:00:00\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"Z07:00\", internal=parsing time \"2024-06-13T12:00:00\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"Z07:00\""}]}
```

### Comment by @xbhouse on June 12, 2024 at 01:17 PM UTC

@swadeley looks like this happens on the main branch too with that datetime format. can you try with a datetime in either of these formats: `2024-06-12T00:00:00-04:00` or `2024-06-12T00:00:00.000000Z`?  

### Comment by @xbhouse on June 12, 2024 at 01:18 PM UTC

> Hi
> 
> Creating:
> 
> ```
> In [4]: app.content_sources.rest_client.templates_api.create_template(template_dict)
> 2024-06-12 11:06:03.538 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
> Out[4]: 
> {'arch': 'x86_64',
>  'created_at': '2024-06-12T10:06:03.492091Z',
>  'created_by': 'ephemeral-user',
>  'date': '0001-01-01T00:00:00Z',
>  'description': 'my 1st template',
>  'last_updated_by': 'ephemeral-user',
>  'name': 'test1 el8',
>  'org_id': '3340851',
>  'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
>  'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
>  'updated_at': '2024-06-12T10:06:03.492091Z',
>  'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
>  'version': '8'}
> ```
> 
> updating:
> 
> ```
> In [12]: template_update_dict = dict(repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template")
> 
> 
> In [15]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
> 2024-06-12 11:34:08.173 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
> Out[15]: 
> {'arch': 'x86_64',
>  'created_at': '2024-06-12T10:06:03.492091Z',
>  'created_by': 'ephemeral-user',
>  'date': '2024-06-12T10:34:08.120882Z',
>  'description': 'my edited template',
>  'last_updated_by': 'ephemeral-user',
>  'name': 'test1 el8',
>  'org_id': '3340851',
>  'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
>  'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
>  'updated_at': '2024-06-12T10:34:08.124216Z',
>  'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
>  'version': '8'}
> ```
> 
> but I cannot change the name. `ApiAttributeError: ApiTemplateUpdateRequest has no attribute 'name' at ['api_template_update_request']['name'] `
> 
> ```
>     attribute_map = {
>         'date': 'date',  # noqa: E501
>         'description': 'description',  # noqa: E501
>         'repository_uuids': 'repository_uuids',  # noqa: E501
>     }
> ```
> 
> @rverdile , I thought it should now be possible to change the name?

@swadeley the [PR](https://github.com/content-services/content-sources-backend/pull/694) to make the template name editable hasn't been merged in yet, it's dependent on an addition to the candlepin API

### Comment by @swadeley on June 12, 2024 at 01:19 PM UTC

Hi

this works:

```
In [28]: template_update_dict = dict(date='2024-06-13T12:00:00.120882Z',repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template with date change")

In [29]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
2024-06-12 14:17:26.222 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
Out[29]: 
{'arch': 'x86_64',
 'created_at': '2024-06-12T10:06:03.492091Z',
 'created_by': 'ephemeral-user',
 'date': '2024-06-13T12:00:00.120882Z',
 'description': 'my edited template with date change',
 'last_updated_by': 'ephemeral-user',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
 'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
 'updated_at': '2024-06-12T13:17:26.052978Z',
 'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
 'version': '8'}

In [30]:
```

### Comment by @swadeley on June 12, 2024 at 01:22 PM UTC

> 2024-06-12T00:00:00.000000Z

Hi, yes, those work:

```
In [30]: template_update_dict = dict(date='2024-06-12T00:00:00-04:00',repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template with date change")

In [31]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
2024-06-12 14:20:55.089 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
Out[31]: 
{'arch': 'x86_64',
 'created_at': '2024-06-12T10:06:03.492091Z',
 'created_by': 'ephemeral-user',
 'date': '2024-06-12T04:00:00Z',
 'description': 'my edited template with date change',
 'last_updated_by': 'ephemeral-user',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
 'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
 'updated_at': '2024-06-12T13:20:54.941332Z',
 'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
 'version': '8'}

In [32]: template_update_dict = dict(date='2024-06-12T00:00:00.000000Z',repository_uuids=['fc6f77da-819b-4d0e-914b-ab359013b6dd'],description="my edited template with date change")

In [33]: app.content_sources.rest_client.templates_api.full_update_template('867a8af5-2110-4a4f-a969-f48814eed816',template_update_dict)
2024-06-12 14:21:22.465 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
Out[33]: 
{'arch': 'x86_64',
 'created_at': '2024-06-12T10:06:03.492091Z',
 'created_by': 'ephemeral-user',
 'date': '2024-06-12T00:00:00Z',
 'description': 'my edited template with date change',
 'last_updated_by': 'ephemeral-user',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['fc6f77da-819b-4d0e-914b-ab359013b6dd'],
 'rhsm_environment_id': '867a8af521104a4fa969f48814eed816',
 'updated_at': '2024-06-12T13:21:22.345812Z',
 'uuid': '867a8af5-2110-4a4f-a969-f48814eed816',
 'version': '8'}

In [34]: 
```

thank you

---

## Reviews

### Review by @xbhouse - Approved on June 11, 2024 at 02:36 PM UTC

works and looks great! ack! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/699*
