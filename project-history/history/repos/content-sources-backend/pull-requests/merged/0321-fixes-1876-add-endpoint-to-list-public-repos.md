---
type: pull_request
number: 321
title: "Fixes 1876: add endpoint to list public repos"
state: merged
author: rverdile
created: 2023-06-27T19:56:21Z
updated: 2023-08-28T20:55:34Z
closed: 2023-07-06T15:50:20Z
merged: 2023-07-06T15:50:20Z
base_branch: main
head_branch: pub-repo-endpoint
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/321
---

# Pull Request #321: Fixes 1876: add endpoint to list public repos

**Author**: @rverdile
**Created**: June 27, 2023 at 07:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `pub-repo-endpoint`

## Description

## Summary
Adds a `GET .../v1/public_repositories/` endpoint to return a paginated list all the "public" repositories (i.e. the repositories on the repositories table where "public" is true). These are also the ones used by image builder.
## Testing steps
1. Run `make repos-import` to make sure you've imported all the public repos
2. GET `http://localhost:8000/api/content-sources/v1/public_repositories/?limit=100&offset=0`. Also try using the offset and limit params.
3. You should get a paginated list of results.


---

## Discussion

### Comment by @jlsherrill on June 27, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1876

### Comment by @jlsherrill on June 29, 2023 at 06:56 PM UTC

everything else looks fine and works great

### Comment by @swadeley on July 06, 2023 at 09:31 AM UTC

/retest

### Comment by @swadeley on July 06, 2023 at 12:26 PM UTC

Hi

API works:
```
In [3]: app.content_sources.rest_client.public_repositories_api.list_public_repositories()
2023-07-06 11:54:15.694 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=kCSdGLU5GrHCny3MqOaPFuaDyIrWB4cc, params=[]
Out[3]: 
{'data': [{'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/',
           'uuid': '719d9a7a-613b-470f-1234-45348fc6621c'},
          {'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/',
           'uuid': '21f5da47-cc2b-4de7-bc91-4ed3efa3ab4f'},

```

but I can't get the limit and offset to work yet.


### Comment by @swadeley on July 06, 2023 at 12:31 PM UTC

Hi

compare:

In [4]: app.content_sources.rest_client.tasks_api.list_tasks.params_map
```
Out[4]: 
{'all': ['offset',
  'limit',
  'status',
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
```
and

```
In [5]: app.content_sources.rest_client.public_repositories_api.list_public_repositories.params_map
Out[5]: 
{'all': ['async_req',
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


```

### Comment by @swadeley on July 06, 2023 at 01:07 PM UTC

Hi

it works OK if I use curl:

content-sources-backend]$ curl  -u "$BASIC_AUTH" "https://${URL}/api/content-sources/v1.0/public_repositories/?limit=2&offset=2" | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   867  100   867    0     0   1353      0 --:--:-- --:--:-- --:--:--  1354
{
  "data": [
    {
      "uuid": "6d3c8483-1234-4daa-1234-ead366c478e6",
      "url": "http://mirror.centos.org/centos/8-stream/extras/x86_64/os/",
      "status": "Pending",
      "last_introspection_time": "",
      "last_success_introspection_time": "",
      "last_update_introspection_time": "",
      "last_introspection_error": "",
      "package_count": 0
    },
    {
      "uuid": "f36123a8-1234-4d0d-1234-7c3cfd0519e5",
      "url": "http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/",
      "status": "Pending",
      "last_introspection_time": "",
      "last_success_introspection_time": "",
      "last_update_introspection_time": "",
      "last_introspection_error": "",
      "package_count": 0
    }
  ],
  "meta": {
    "limit": 2,
    "offset": 2,
    "count": 0
  },
  "links": {
    "first": "/api/content-sources/v1.0/public_repositories/?limit=2&offset=0",
    "prev": "/api/content-sources/v1.0/public_repositories/?limit=2&offset=0",
    "last": "/api/content-sources/v1.0/public_repositories/?limit=2&offset=0"
  }
}



### Comment by @rverdile on July 06, 2023 at 03:02 PM UTC

limit and offset were missing from the docs. it should work now!

### Comment by @swadeley on July 06, 2023 at 03:50 PM UTC

Hi

```
In [1]: app.content_sources.rest_client.public_repositories_api.list_public_repositories.params_map
<snip>
Out[1]: 
{'all': ['offset',
  'limit',
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

```

```
In [2]: app.content_sources.rest_client.public_repositories_api.list_public_repositories(limit=2, offset=2)
2023-07-06 16:47:47.046 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=YlqVZDPou0hwsmStDI9dZxHFNGDYwVZu, params=[('limit', 2), ('offset', 2)]
Out[2]: 
{'data': [{'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/extras/x86_64/os/',
           'uuid': '6d3c8483-9492-1234-1234-ead366c478e6'},
          {'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/',
           'uuid': 'f36123a8-1234-1234-9980-7c3cfd0519e5'}],
 'links': {'first': '/api/content-sources/v1/public_repositories/?limit=2&offset=0',
           'last': '/api/content-sources/v1/public_repositories/?limit=2&offset=0',
           'prev': '/api/content-sources/v1/public_repositories/?limit=2&offset=0'},
 'meta': {'count': 0, 'limit': 2, 'offset': 2}}
```

Thank you

---

## Reviews

### Review by @jlsherrill - Commented on June 29, 2023 at 03:13 PM UTC

### Review by @jlsherrill - Commented on June 29, 2023 at 06:44 PM UTC

### Review by @rverdile - Commented on June 30, 2023 at 08:23 PM UTC

### Review by @jlsherrill - Approved on July 05, 2023 at 03:02 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/321*
