---
type: pull_request
number: 588
title: "Fixes 3654: Remove UUID from public repos"
state: merged
author: Andrewgdewar
created: 2024-02-26T21:12:43Z
updated: 2024-02-27T20:00:30Z
closed: 2024-02-27T19:43:47Z
merged: 2024-02-27T19:43:47Z
base_branch: main
head_branch: HMS-3654
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/588
---

# Pull Request #588: Fixes 3654: Remove UUID from public repos

**Author**: @Andrewgdewar
**Created**: February 26, 2024 at 09:12 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3654`

## Description

## Summary

Simply removes the UUID from the public_repos response.

## Testing steps

- Get /public_repositories/ list, check that the UUID is no longer present.


---

## Discussion

### Comment by @jlsherrill on February 26, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-3654

### Comment by @xbhouse on February 27, 2024 at 05:28 PM UTC

lgtm! thanks for the addition to the readme too :) 

### Comment by @swadeley on February 27, 2024 at 06:47 PM UTC

Hi

lgtm

```
In [3]: app.content_sources.rest_client.public_repositories_api.list_public_repositories(limit=4)
2024-02-27 18:45:26.209 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 4)]
Out[3]: 
{'data': [{'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/'},
          {'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/'},
          {'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.centos.org/centos/8-stream/extras/x86_64/os/'},
          {'last_introspection_error': '',
           'last_introspection_time': '',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'package_count': 0,
           'status': 'Pending',
           'url': 'http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/'}],
 'links': {'first': '/api/content-sources/v1/public_repositories/?limit=4&offset=0',
           'last': '/api/content-sources/v1/public_repositories/?limit=4&offset=52',
           'next': '/api/content-sources/v1/public_repositories/?limit=4&offset=4'},
 'meta': {'count': 56, 'limit': 4, 'offset': 0}}

In [4]: 


```

### Comment by @swadeley on February 27, 2024 at 07:10 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on February 27, 2024 at 05:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/588*
