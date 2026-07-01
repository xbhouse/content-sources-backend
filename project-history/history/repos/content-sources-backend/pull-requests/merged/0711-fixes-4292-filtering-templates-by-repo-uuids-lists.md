---
type: pull_request
number: 711
title: "Fixes 4292: filtering templates by repo uuids lists dups"
state: merged
author: xbhouse
created: 2024-06-18T15:37:39Z
updated: 2024-06-20T15:30:27Z
closed: 2024-06-20T15:25:54Z
merged: 2024-06-20T15:25:54Z
base_branch: main
head_branch: 4292
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/711
---

# Pull Request #711: Fixes 4292: filtering templates by repo uuids lists dups

**Author**: @xbhouse
**Created**: June 18, 2024 at 03:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4292`

## Description

## Summary

Fixes a bug where filtering templates by multiple repo uuids lists duplicate templates

## Testing steps

- Add 2 repositories, let them snapshot, and create a template with both of those repositories
- Make a request to list templates filtering by those 2 repo uuids (`templates/?repository_uuids=f572ddd9-2aaa-4503-bd87-9be723574f33,3a379621-e5af-40ad-a7a6-1b0c59934194`)
- You should see that template listed once and count should be 1:
```
{
    "data": [
        {
            "uuid": "8f52f555-eef5-4b06-96ce-c0afdd260dd6",
            "name": "template1",
            "org_id": "17684632",
            "description": "",
            "arch": "x86_64",
            "version": "8",
            "date": "2024-06-17T00:00:00-04:00",
            "repository_uuids": [
                "f572ddd9-2aaa-4503-bd87-9be723574f33",
                "3a379621-e5af-40ad-a7a6-1b0c59934194",
            ],
            "rhsm_environment_id": "8f52f555eef54b0696cec0afdd260dd6",
            "created_by": "bryttaniehouse",
            "last_updated_by": "bryttaniehouse",
            "created_at": "2024-06-17T11:40:12.335712-04:00",
            "updated_at": "2024-06-17T11:40:12.335712-04:00"
        }
    ],
    "meta": {
        "limit": 100,
        "offset": 0,
        "count": 1
    },
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 18, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4292

### Comment by @xbhouse on June 18, 2024 at 06:16 PM UTC

@rverdile oh woops! just pushed up a test

### Comment by @swadeley on June 19, 2024 at 10:37 AM UTC

Hi

Is the API spec correct?

```
                    {
                        "description": "Filter templates by associated repositories using a comma separated list of repository UUIDs",
                        "in": "query",
                        "name": "repository_uuids",
                        "schema": {
                            "type": "string"
                        }
                    },

```

I can pass in one UUID

```
In [44]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids='62960863-afd6-44be-bf64-98cade0f747e')
2024-06-19 10:53:44.839 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('repository_uuids', '62960863-afd6-44be-bf64-98cade0f747e')]
Out[44]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-06-19T08:29:47.363217Z',
           'created_by': 'ephemeral-user',
           'date': '2024-06-18T23:00:00Z',
           'description': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'test-template1',
           'org_id': '3340851',
           'repository_uuids': ['f4ec14e3-3ddf-441d-be83-49a1719a2f55',
                                '8648d5c4-c7f6-4df8-b70a-91b27a662b32',
                                '27834f5c-5f30-48e7-a40c-7b2ed8eec8e3',
                                '05c84992-bc18-4a1c-832f-b52e0eec2cb9',
                                '62960863-afd6-44be-bf64-98cade0f747e',
                                'f4b7b620-318d-467b-9736-2a2702fcb63a',
                                '2de4abd1-6006-4582-8345-4a62a27a679b'],
```

but I cannot figure out how to pass in a list:
```
In [65]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids=['f4ec14e3-3ddf-441d-be83-49a1719a2f55','8648d5c4-c7f6-4df8-b70a-91b27a662b32'])
ApiTypeError: Invalid type for variable 'repository_uuids'. Required value type is str and passed type was list at ['repository_uuids']
```

```
In [66]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids='f4ec14e3-3ddf-441d-be83-49a1719a2f55','8648d5c4-c7f6-4df8-b70a-91b27a662b32')
  Cell In[66], line 1
    app.content_sources.rest_client.templates_api.list_templates(repository_uuids='f4ec14e3-3ddf-441d-be83-49a1719a2f55','8648d5c4-c7f6-4df8-b70a-91b27a662b32')
                                                                                                                                                               ^
SyntaxError: positional argument follows keyword argument
```

### Comment by @xbhouse on June 20, 2024 at 01:33 PM UTC

@swadeley can you try with a comma-separated string? like this: `app.content_sources.rest_client.templates_api.list_templates(repository_uuids="some-uuid1,some-uuid2")`. the api spec should be correct, it's similar to how you'd filter by multiple uuids in the list repositories API

```
In [13]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="4acc07e6-19a3-423f-a619-d6fd706850,69311439-a36a-4c8a-a552-ef276d501cfc")
2024-06-20 09:43:55.424 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=a0d4f983-6f8a-483d-aea4-46c053cc37f7, params=[('repository_uuids', '4acc07e6-19a3-423f-a619-d6d6fd706850,69311439-a36a-4c8a-a552-ef276d501cfc')]
Out[13]: 
{'data': [{'arch': '',
           'created_at': '2024-06-20T13:43:11.646918Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'test',
           'org_id': '3340851',
           'repository_uuids': ['4acc07e6-19a3-423f-a619-d6d6fd706850',
                                '69311439-a36a-4c8a-a552-ef276d501cfc'],
           'rhsm_environment_id': '4787679034b34943b3d75836db2e2c32',
           'updated_at': '2024-06-20T13:43:11.646918Z',
           'uuid': '47876790-34b3-4943-b3d7-5836db2e2c32',
           'version': ''}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=4acc07e6-19a3-423f-a619-d6d6fd706850,69311439-a36a-4c8a-a552-ef276d501cfc',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=4acc07e6-19a3-423f-a619-d6d6fd706850,69311439-a36a-4c8a-a552-ef276d501cfc'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
 ```

### Comment by @swadeley on June 20, 2024 at 03:14 PM UTC

Hi @xbhouse 

that was the second method I tried , see comment above yours, but I got error:
`SyntaxError: positional argument follows keyword argument`

EDIT: oh, wait, i see now what you mean

`(repository_uuids="some-uuid1,some-uuid2")`




### Comment by @swadeley on June 20, 2024 at 03:23 PM UTC

Ok, it works, thank you


```
In [12]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="3b4b971b-c140-4997-8345-03b94b0d00ca,5916447d-5dcd-42a7-9d08-8b3f8bc1fd50")
2024-06-20 16:22:28.913 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('repository_uuids', '3b4b971b-c140-4997-8345-03b94b0d00ca,5916447d-5dcd-42a7-9d08-8b3f8bc1fd50')]
Out[12]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-06-20T15:22:25.343057Z',
           'created_by': 'ephemeral-user',
           'date': '2024-06-19T23:00:00Z',
           'description': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'test-template1',
           'org_id': '3340851',
           'repository_uuids': ['3b4b971b-c140-4997-8345-03b94b0d00ca',
                                '5916447d-5dcd-42a7-9d08-8b3f8bc1fd50',
                                '1bec0599-dc56-4bd9-9dd4-8fd5420e72e2'],
           'rhsm_environment_id': '6d6a2f0904584ae4a1447692ded00d68',
           'updated_at': '2024-06-20T15:22:25.343057Z',
           'uuid': '6d6a2f09-0458-4ae4-a144-7692ded00d68',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=3b4b971b-c140-4997-8345-03b94b0d00ca,5916447d-5dcd-42a7-9d08-8b3f8bc1fd50',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=3b4b971b-c140-4997-8345-03b94b0d00ca,5916447d-5dcd-42a7-9d08-8b3f8bc1fd50'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [13]: 

```
count is 1 as expected.

---

## Reviews

### Review by @rverdile - Commented on June 18, 2024 at 05:57 PM UTC

looks good! can you just add a test? you can probably just modify the existing one for repository uuids

### Review by @rverdile - Approved on June 18, 2024 at 06:35 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/711*
