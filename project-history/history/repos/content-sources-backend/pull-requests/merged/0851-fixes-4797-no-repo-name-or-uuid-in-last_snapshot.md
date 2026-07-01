---
type: pull_request
number: 851
title: "Fixes 4797: no repo name or uuid in last_snapshot"
state: merged
author: xbhouse
created: 2024-10-16T18:45:24Z
updated: 2024-10-18T08:30:30Z
closed: 2024-10-18T08:03:17Z
merged: 2024-10-18T08:03:17Z
base_branch: main
head_branch: 4797
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/851
---

# Pull Request #851: Fixes 4797: no repo name or uuid in last_snapshot

**Author**: @xbhouse
**Created**: October 16, 2024 at 06:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4797`

## Description

## Summary

Fixes a bug where the repo UUID and name were missing from `last_snapshot` in the repository response

## Testing steps

1. Create a repository and let it snapshot
2. Fetch that repository. The `repository_uuid` and `repository_name` should not be empty in `last_snapshot`
```
...
"last_snapshot": {
        "uuid": "911a76a9-2942-42b4-85aa-406222e543d2",
        "created_at": "2024-10-16T18:22:58.496151Z",
        "repository_path": "34aadb80/30d9d7b9-d49a-45b5-8a75-a527c35bd821/9eb85075-6c1e-426e-aff2-8454eb5f8c80",
        "content_counts": {
            "rpm.advisory": 2,
            "rpm.package": 8,
            "rpm.packagecategory": 1,
            "rpm.packagegroup": 2
        },
        "added_counts": {
            "rpm.advisory": 2,
            "rpm.package": 8,
            "rpm.packagecategory": 1,
            "rpm.packagegroup": 2
        },
        "removed_counts": {},
        "url": "http://pulp.content:8080/pulp/content/34aadb80/30d9d7b9-d49a-45b5-8a75-a527c35bd821/9eb85075-6c1e-426e-aff2-8454eb5f8c80/",
        "repository_name": "test1",
        "repository_uuid": "30d9d7b9-d49a-45b5-8a75-a527c35bd821"
    },
...   
```

---

## Discussion

### Comment by @jlsherrill on October 16, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-4797

### Comment by @swadeley on October 17, 2024 at 09:56 AM UTC

Hi, looks good:

```
In [3]: app.content_sources.rest_client.repositories_api.get_repository('9d41e962-3439-415e-921a-30f9f8e2ca76')['last_snapshot']
2024-10-17 10:55:41.531 [    INFO] [iqe.base.rest_client] REST: GET https://ee-bwgsl3he.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/9d41e962-3439-415e-921a-30f9f8e2ca76 with query params [] and x-rh-insights-request-id=<>
Out[3]: 
{'added_counts': {'rpm.package': 1},
 'content_counts': {'rpm.package': 1},
 'created_at': '2024-10-17T09:54:14.143551Z',
 'removed_counts': {},
 'repository_name': 'Repo1-MshYVXdy',
 'repository_path': '754f4feb/9d41e962-3439-415e-921a-30f9f8e2ca76/822a624c-c71b-482c-8c96-0853aaa0964f',
 'repository_uuid': '9d41e962-3439-415e-921a-30f9f8e2ca76',
 'url': 'http://pulp-content:8000/api/pulp-content/754f4feb/9d41e962-3439-415e-921a-30f9f8e2ca76/822a624c-c71b-482c-8c96-0853aaa0964f/',
 'uuid': '26e54d51-f10b-4e34-adcc-fd6cb6267f58'}
```

---

## Reviews

### Review by @dominikvagner - Approved on October 17, 2024 at 12:12 PM UTC

tested and works as expected 🎉 
good job, approved! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/851*
