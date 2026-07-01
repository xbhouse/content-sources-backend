---
type: pull_request
number: 839
title: "Fixes 4151: remove support for YYYY-MM-DD date formats"
state: merged
author: xbhouse
created: 2024-10-07T14:06:16Z
updated: 2024-10-10T09:57:15Z
closed: 2024-10-10T09:57:14Z
merged: 2024-10-10T09:57:14Z
base_branch: main
head_branch: 4151
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/839
---

# Pull Request #839: Fixes 4151: remove support for YYYY-MM-DD date formats

**Author**: @xbhouse
**Created**: October 07, 2024 at 02:06 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4151`

## Description

## Summary

- Changes /snapshots/for_date/ request to remove support for YYYY-MM-DD date formats
    - We currently support both RFC3339 and YYYY-MM-DD formats to maintain backward compatibility with IB. We can remove support for YYYY-MM-DD now as they've made the necessary changes [here](https://github.com/osbuild/image-builder/pull/1252)
- [ Frontend PR](https://github.com/content-services/content-sources-frontend/pull/352) should be merged first

## Testing steps

1. Try to make a request to /snapshots/for_date/ with a YYYY-MM-DD date, this should fail:
```
{
  "date": "2024-10-02",
  "repository_uuids": ["30050c99-e904-4bf3-a31d-c213e203e19d"]
}
```
2. A request like this should still work:
```
{
  "date": "2024-10-02T00:00:00Z",
  "repository_uuids": ["30050c99-e904-4bf3-a31d-c213e203e19d"]
}
```




---

## Discussion

### Comment by @jlsherrill on October 07, 2024 at 05:36 PM UTC

https://issues.redhat.com/browse/HMS-4151

### Comment by @swadeley on October 09, 2024 at 08:00 PM UTC

/retest

### Comment by @swadeley on October 10, 2024 at 09:40 AM UTC

Hi

this gives error:
```
In [2]: app.content_sources.rest_client.snapshots_api.list_snapshots_by_date({"repository_uuids":['35e82994-04ac-465e-8b72-0318a2cfbc7f'],"date": "2024-10-10"})

2024-10-10 10:33:31.797 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ceta08o4.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/for_date/ with query params [] and x-rh-insights-request-id=<>
<snip>
HTTP response body: {"errors":[{"status":400,"title":"Error binding parameters","detail":"code=400, message=parsing time \"2024-10-10\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"T\", internal=parsing time \"2024-10-10\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"T\""}]}
```

This works:
```
In [4]: app.content_sources.rest_client.snapshots_api.list_snapshots_by_date({"repository_uuids":['35e82994-04ac-465e-8b72-0318a2cfbc7f'],"date": "2024-10-10T00:00:00Z"})
2024-10-10 10:38:12.722 [    INFO] [iqe.base.rest_client] REST: POST https://ee-ceta08o4.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/for_date/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'data': [{'is_after': False,
           'match': {'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'created_at': '2024-10-10T09:20:25.921784Z',
                     'removed_counts': {},
                     'repository_name': 'multiple_errata',
                     'repository_path': 'ac1cf75a/35e82994-04ac-465e-8b72-0318a2cfbc7f/b9475114-8714-4851-8016-2f83e6bb1fc5',
                     'repository_uuid': '35e82994-04ac-465e-8b72-0318a2cfbc7f',
                     'url': 'http://pulp-content:8000/api/pulp-content/ac1cf75a/35e82994-04ac-465e-8b72-0318a2cfbc7f/b9475114-8714-4851-8016-2f83e6bb1fc5/',
                     'uuid': 'f6f12432-7efe-4ea2-b08d-d163d4b2f68d'},
           'repository_uuid': '35e82994-04ac-465e-8b72-0318a2cfbc7f'}
```

---

## Reviews

### Review by @swadeley - Approved on October 10, 2024 at 09:56 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/839*
