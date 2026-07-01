---
type: pull_request
number: 785
title: "Fixes 4297: template API responds in UTC time"
state: merged
author: dominikvagner
created: 2024-08-26T12:37:38Z
updated: 2024-08-28T14:30:28Z
closed: 2024-08-28T14:19:51Z
merged: 2024-08-28T14:19:51Z
base_branch: main
head_branch: 4297
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/785
---

# Pull Request #785: Fixes 4297: template API responds in UTC time

**Author**: @dominikvagner
**Created**: August 26, 2024 at 12:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4297`

## Description

## Summary
This PR makes the template part of the API respond with times in UTC time format, which fixes weird timezone/location converted output of the zero date (and others). This has to be done, even though all of our `time.Time` data is stored as UTC in our DB, but that is converted when reading it to the local time of the system it's running on (and also output in that).
Also has a [sibling front-end PR](https://github.com/content-services/content-sources-frontend/pull/325) addressing [this issue](https://issues.redhat.com/browse/HMS-4297) there.
We should think about doing this for all of our endpoints so that the responses are clearer and less "confusing"/more standardized.

## Testing steps
Make a request to the endpoint which lists all templates and check date that the format of all dates is in UTC (doesn't have a time offset)
```bash
http :8000/api/content-sources/v1.0/templates/ "$( ./scripts/header.sh acme 1111)
```


---

## Discussion

### Comment by @jlsherrill on August 26, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4297

### Comment by @swadeley on August 27, 2024 at 12:19 PM UTC

Hi @dominikvagner , please rebase to pick up change in https://github.com/content-services/content-sources-backend/pull/788

Thank you

### Comment by @swadeley on August 28, 2024 at 10:29 AM UTC

Hi 

I deployed images: backend=pr-788-90abe20  frontend=pr-325-aa3f61e

I created a template at 10:13 UTC, which is 11:13 local time

In shell to ephemeral I see:

```
Out[1]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-08-28T10:13:02.554325Z',
           'created_by': 'ephemeral-user',
           'date': '2024-08-27T23:00:00Z',
           'description': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'template1',
           'org_id': '3340851',
           'repository_uuids': ['c262ef41-59e9-44a8-bb17-6140f3f7fc2e',
                                'f619ae84-ab11-40c0-b0cd-1bf887a91644'],
           'rhsm_environment_id': 'c6a604cf850e4ef3bf42b8c016c583c1',
           'updated_at': '2024-08-28T10:13:02.554325Z',
           'use_latest': False,
           'uuid': 'c6a604cf-850e-4ef3-bf42-b8c016c583c1',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

Previously the template target date would have been returned as:
'date': '0001-01-01T00:00:00Z',

Now it returned:
 'date': '2024-08-27T23:00:00Z',

So that is changed.

### Comment by @dominikvagner on August 28, 2024 at 12:47 PM UTC

> Hi
> 
> I deployed images: backend=pr-788-90abe20 frontend=pr-325-aa3f61e
> 
> I created a template at 10:13 UTC, which is 11:13 local time
> 
> In shell to ephemeral I see:
> 
> ```
> Out[1]: 
> {'data': [{'arch': 'x86_64',
>            'created_at': '2024-08-28T10:13:02.554325Z',
>            'created_by': 'ephemeral-user',
>            'date': '2024-08-27T23:00:00Z',
>            'description': '',
>            'last_updated_by': 'ephemeral-user',
>            'name': 'template1',
>            'org_id': '3340851',
>            'repository_uuids': ['c262ef41-59e9-44a8-bb17-6140f3f7fc2e',
>                                 'f619ae84-ab11-40c0-b0cd-1bf887a91644'],
>            'rhsm_environment_id': 'c6a604cf850e4ef3bf42b8c016c583c1',
>            'updated_at': '2024-08-28T10:13:02.554325Z',
>            'use_latest': False,
>            'uuid': 'c6a604cf-850e-4ef3-bf42-b8c016c583c1',
>            'version': '8'}],
>  'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
>            'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
>  'meta': {'count': 1, 'limit': 100, 'offset': 0}}
> ```
> 
> Previously the template target date would have been returned as: 'date': '0001-01-01T00:00:00Z',
> 
> Now it returned: 'date': '2024-08-27T23:00:00Z',
> 
> So that is changed.

How did you create the template? 😄 When I create it through the API with `use_latest=false` and don't supply the date, I get the template with the zero date.
```
❯ http :8000/api/content-sources/v1.0/templates/ "$( ./scripts/header.sh 18064594 1111)" arch='x86_64' version='9' repository_uuids:='[]' name='test' use_latest:=false
HTTP/1.1 201 Created
Content-Length: 392
Content-Type: application/json
Date: Wed, 28 Aug 2024 12:45:26 GMT
X-Rh-Identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxODA2NDU5NCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxMTExIn0sImFjY291bnRfbnVtYmVyIjoiMTgwNjQ1OTQiLCJpbnRlcm5hbCI6eyJvcmdfaWQiOiIxODA2NDU5NCJ9fX0K
X-Rh-Insights-Request-Id: e3a14e0a-0af1-4238-821c-7ee6300dab0e

{
    "arch": "x86_64",
    "created_at": "2024-08-28T12:45:26.82580689Z",
    "created_by": "1111",
    "date": "0001-01-01T00:00:00Z",
    "description": "",
    "last_updated_by": "1111",
    "name": "test",
    "org_id": "18064594",
    "repository_uuids": [],
    "rhsm_environment_id": "f1552d6e566e4c1d94148674d7d9e9ab",
    "updated_at": "2024-08-28T12:45:26.82580689Z",
    "use_latest": false,
    "uuid": "f1552d6e-566e-4c1d-9414-8674d7d9e9ab",
    "version": "9"
}
```

### Comment by @swadeley on August 28, 2024 at 02:14 PM UTC

Hi @dominikvagner 

that first (and second) template was created in the UI. Sorry for confusion.

Here is my post from the sibling PR #325:

```
In [6]:  app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-08-28 13:18:40.115 [    INFO] [iqe.base.rest_client] REST: POST https://ee-suq7uw68.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=<>
Out[6]: 
{'arch': 'x86_64',
 'created_at': '2024-08-28T12:18:40.060803083Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test zero',
 'org_id': '3340851',
 'repository_uuids': ['c262ef41-59e9-44a8-bb17-6140f3f7fc2e'],
 'rhsm_environment_id': 'c175d10e81c049839a1ec3c84c31a81d',
 'updated_at': '2024-08-28T12:18:40.060803083Z',
 'use_latest': False,
 'uuid': 'c175d10e-81c0-4983-9a1e-c3c84c31a81d',
 'version': '8'}

In [7]: 
```

![image](https://github.com/user-attachments/assets/16049746-8a2e-489d-b69c-41e8677ecb08)


lgtm

### Comment by @swadeley on August 28, 2024 at 02:18 PM UTC

@dominikvagner I can merge this and then deploy #791 with #325 for testing further.

---

## Reviews

### Review by @xbhouse - Approved on August 27, 2024 at 04:26 PM UTC

looks good, template datetimes are consistent! 🎉 nice job! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/785*
