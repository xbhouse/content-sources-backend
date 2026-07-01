---
type: pull_request
number: 842
title: "Fixes 4821: expose URL on snapshot/for_date"
state: merged
author: jlsherrill
created: 2024-10-09T14:32:25Z
updated: 2024-10-09T17:30:23Z
closed: 2024-10-09T17:28:30Z
merged: 2024-10-09T17:28:30Z
base_branch: main
head_branch: 4821
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/842
---

# Pull Request #842: Fixes 4821: expose URL on snapshot/for_date

**Author**: @jlsherrill
**Created**: October 09, 2024 at 02:32 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4821`

## Description

## Summary

fills in the URL field on the snapshot response on the for_date snapshot api

## Testing steps

1.  create a repo, let it snapshot
2. call the for_date api:
```
####
POST http://localhost:8000/api/content-sources/v1.0/snapshots/for_date/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "repository_uuids": ["cace7ad3-e385-41a8-afc2-aae94bde6cf6"],
  "date": "2024-10-08T14:19:58Z"
}
```

verify the response has a url filled in:

```

{
  "data": [
    {
      "repository_uuid": "cace7ad3-e385-41a8-afc2-aae94bde6cf6",
      "is_after": true,
      "match": {
        "uuid": "ac856ef6-bd0f-473f-9574-26328a26a215",
        "created_at": "2024-10-09T14:17:31.029366Z",
        "repository_path": "95161890/cace7ad3-e385-41a8-afc2-aae94bde6cf6/2c1f5a50-0087-4368-bc9f-e1d62c5c836c",
        "content_counts": {
          "rpm.advisory": 4,
          "rpm.package": 32
        },
        "added_counts": {
          "rpm.advisory": 4,
          "rpm.package": 32
        },
        "removed_counts": {},
        "url": "http://pulp.content:8080/pulp/content/95161890/cace7ad3-e385-41a8-afc2-aae94bde6cf6/2c1f5a50-0087-4368-bc9f-e1d62c5c836c/",
        "repository_name": "test57",
        "repository_uuid": "cace7ad3-e385-41a8-afc2-aae94bde6cf6"
      }
    }
  ]
}
```



---

## Discussion

### Comment by @jlsherrill on October 09, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4821

### Comment by @swadeley on October 09, 2024 at 05:26 PM UTC

Hi

Created repo, let it snapshot, got the snapshot using repo UUID:

```
In [4]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('11f22fc1-ff6d-4f49-9542-bb6db3bd8df2')
2024-10-09 18:22:10.989 [    INFO] [iqe.base.rest_client] REST: GET https://ee-0s9kqucm.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'data': [{'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'created_at': '2024-10-09T17:18:58.911905Z',
           'removed_counts': {},
           'repository_name': 'multi_errata',
           'repository_path': '724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419',
           'repository_uuid': '11f22fc1-ff6d-4f49-9542-bb6db3bd8df2',
           'url': 'http://pulp-content:8000/api/pulp-content/724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419/',
           'uuid': '6666d5fa-4d91-48b6-90cc-3da8e95080d5'}],
 'links': {'first': '/api/content-sources/v1/repositories/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

Got the snapshot using date:

```
In [6]: app.content_sources.rest_client.snapshots_api.list_snapshots_by_date({"repository_uuids":['11f22fc1-ff6d-4f49-9542-bb6db3bd8df2'],"date": "2024-10-09T17:18:58Z"})
2024-10-09 18:23:25.283 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0s9kqucm.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/for_date/ with query params [] and x-rh-insights-request-id=<>
Out[6]: 
{'data': [{'is_after': False,
           'match': {'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'created_at': '2024-10-09T17:18:58.911905Z',
                     'removed_counts': {},
                     'repository_name': 'multi_errata',
                     'repository_path': '724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419',
                     'repository_uuid': '11f22fc1-ff6d-4f49-9542-bb6db3bd8df2',
                     'url': 'http://pulp-content:8000/api/pulp-content/724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419/',
                     'uuid': '6666d5fa-4d91-48b6-90cc-3da8e95080d5'},
           'repository_uuid': '11f22fc1-ff6d-4f49-9542-bb6db3bd8df2'}]}

```

### Comment by @swadeley on October 09, 2024 at 05:28 PM UTC

You can use full date time too:
```

In [7]: app.content_sources.rest_client.snapshots_api.list_snapshots_by_date({"repository_uuids":['11f22fc1-ff6d-4f49-9542-bb6db3bd8df2'],"date": "2024-10-09T17:18:58.911905Z"})
2024-10-09 18:27:08.373 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0s9kqucm.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/for_date/ with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'data': [{'is_after': False,
           'match': {'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'created_at': '2024-10-09T17:18:58.911905Z',
                     'removed_counts': {},
                     'repository_name': 'multi_errata',
                     'repository_path': '724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419',
                     'repository_uuid': '11f22fc1-ff6d-4f49-9542-bb6db3bd8df2',
                     'url': 'http://pulp-content:8000/api/pulp-content/724e3b9f/11f22fc1-ff6d-4f49-9542-bb6db3bd8df2/c1963811-9030-47d9-8be7-980c4da6d419/',
                     'uuid': '6666d5fa-4d91-48b6-90cc-3da8e95080d5'},
           'repository_uuid': '11f22fc1-ff6d-4f49-9542-bb6db3bd8df2'}]}
```


---

## Reviews

### Review by @xbhouse - Approved on October 09, 2024 at 03:29 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/842*
