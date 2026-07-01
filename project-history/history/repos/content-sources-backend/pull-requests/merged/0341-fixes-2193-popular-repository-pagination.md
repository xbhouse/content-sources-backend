---
type: pull_request
number: 341
title: "Fixes 2193: popular repository pagination"
state: merged
author: dpang314
created: 2023-07-31T19:27:08Z
updated: 2023-08-02T19:09:48Z
closed: 2023-08-02T19:09:48Z
merged: 2023-08-02T19:09:48Z
base_branch: main
head_branch: HMS-2193-popular-repository-pagination
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/341
---

# Pull Request #341: Fixes 2193: popular repository pagination

**Author**: @dpang314
**Created**: July 31, 2023 at 07:27 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-2193-popular-repository-pagination`

## Description

## Summary

The popular repository table UI implements pagination and this PR adds backend support for pagination using the limit and offset URL parameters.

## Testing steps

The JSON that contains popular repositories currently has 3 entries, so it is easiest to test pagination with limits below 3.

```
curl "localhost:8000/api/content-sources/v1/popular_repositories/?offset=1&limit=1"  -H "`./scripts/header.sh 1 user`"
```

Negative or out of range offsets returns an empty array (consistent with other list APIs)


---

## Discussion

### Comment by @jlsherrill on July 31, 2023 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-2193

### Comment by @swadeley on August 02, 2023 at 06:39 PM UTC

Hi

looks good:
```

In [1]: app.content_sources.rest_client.popular_repositories_api.list_popular_repositories.params_map
<snip>
Out[1]: 
{'all': ['offset',
  'limit',
  'search',
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

### Comment by @swadeley on August 02, 2023 at 06:43 PM UTC

Testing

```
In [3]: app.content_sources.rest_client.popular_repositories_api.list_popular_repositories(limit=1,offset=1)
2023-08-02 19:40:17.568 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=b7cOhUGRMlQaCtxcluBWJL2PDxSFGrP5, params=[('limit', 1), ('offset', 1)]
Out[3]: 
<snip>
           'metadata_verification': False,
           'suggested_name': 'EPEL 8 Everything x86_64',
           'url': 'https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/',
           'uuid': ''}],
 'links': {'first': '/api/content-sources/v1/popular_repositories/?limit=1&offset=0',
           'last': '/api/content-sources/v1/popular_repositories/?limit=1&offset=2',
           'next': '/api/content-sources/v1/popular_repositories/?limit=1&offset=2',
           'prev': '/api/content-sources/v1/popular_repositories/?limit=1&offset=0'},
 'meta': {'count': 3, 'limit': 1, 'offset': 1}}

All good

thank you
```

---

## Reviews

### Review by @jlsherrill - Approved on August 01, 2023 at 07:50 PM UTC

Looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/341*
