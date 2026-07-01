---
type: pull_request
number: 926
title: "HMS-4987: fix 500s on invalid input syntax"
state: merged
author: xbhouse
created: 2024-12-18T22:03:27Z
updated: 2025-01-07T01:57:45Z
closed: 2025-01-07T01:57:45Z
merged: 2025-01-07T01:57:45Z
base_branch: main
head_branch: 4987
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/926
---

# Pull Request #926: HMS-4987: fix 500s on invalid input syntax

**Author**: @xbhouse
**Created**: December 18, 2024 at 10:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `4987`

## Description

## Summary

- Fixes "invalid input syntax" errors when requesting to fetch the latest snapshot and list snapshots by date
- Adds validation to the repository uuids in requests to list snapshots by date

## Testing steps

1. Make a request to `/repositories/:uuid/config.repo` with the uuid set to something like "uuid". This should return a 404.
2. Make a request to `/snapshots/for_date/` with the repository_uuids set to something like ["uuid1", "uuid2"]. This should also return a 404.

---

## Discussion

### Comment by @jlsherrill on December 18, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-4987

### Comment by @swadeley on January 02, 2025 at 01:43 AM UTC

/retest

### Comment by @swadeley on January 02, 2025 at 05:47 AM UTC

Hi

I have a repo with a snapshot:
```
In [17]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('7cf1f7c2-d877-458a-a12d-375400b6951d')
2025-01-02 13:42:28.845 [    INFO] [iqe.base.rest_client] REST: GET https://ee-1u593mc2.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/7cf1f7c2-d877-458a-a12d-375400b6951d/snapshots/ with query params [] and x-rh-insights-request-id=<-9dcd-4012-b613->
Out[17]: 
{'data': [{'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'created_at': '2025-01-02T05:13:31.779967Z',
           'removed_counts': {},
           'repository_name': 'test',
           'repository_path': 'cs-0768d2ac43/7cf1f7c2-d877-458a-a12d-375400b6951d/29c69cdd-8374-4280-b48a-ae020ef3720b',
           'repository_uuid': '7cf1f7c2-d877-458a-a12d-375400b6951d',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-0768d2ac43/7cf1f7c2-d877-458a-a12d-375400b6951d/29c69cdd-8374-4280-b48a-ae020ef3720b/',
           'uuid': '6425d836-6120-400e-9b13-69285f1820e9'}],
 'links': {'first': '/api/content-sources/v1/repositories/7cf1f7c2-d877-458a-a12d-375400b6951d/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/7cf1f7c2-d877-458a-a12d-375400b6951d/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```
I can list the snapshots by date:

```
In [18]: app.content_sources.rest_client.snapshots_api.list_snapshots_by_date({"repository_uuids":['7cf1f7c2-d877-458a-a12d-375400b6951d'],"date": "2025-01-02T05:13:00Z"})
2025-01-02 13:43:11.890 [    INFO] [iqe.base.rest_client] REST: POST https://ee-1u593mc2.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/for_date/ with query params [] and x-rh-insights-request-id=<-706d-4e78-b11c->
Out[18]: 
{'data': [{'is_after': False,
           'match': {'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
                     'created_at': '2025-01-02T05:13:31.779967Z',
                     'removed_counts': {},
                     'repository_name': '',
                     'repository_path': 'cs-0768d2ac43/7cf1f7c2-d877-458a-a12d-375400b6951d/29c69cdd-8374-4280-b48a-ae020ef3720b',
                     'repository_uuid': '',
                     'url': 'http://pulp-content:8000/api/pulp-content/cs-0768d2ac43/7cf1f7c2-d877-458a-a12d-375400b6951d/29c69cdd-8374-4280-b48a-ae020ef3720b/',
                     'uuid': '6425d836-6120-400e-9b13-69285f1820e9'},
           'repository_uuid': '7cf1f7c2-d877-458a-a12d-375400b6951d'}]}

```

If I mangle the UUID (changed the last three characters):

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '116', 'content-type': 'application/json', 'date': 'Thu, 02 Jan 2025 05:45:35 GMT', 'server': 'Caddy', 'x-rh-identity': 'e<snip>=', 'x-rh-insights-request-id': '<snip>', 'set-cookie': '<snip>=<snip>; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error fetching snapshots","detail":"One or more repository uuids was invalid."}]}
```

Looks good, upper case UUID would be good in text.


### Comment by @swadeley on January 07, 2025 at 01:57 AM UTC

Thank you

---

## Reviews

### Review by @dominikvagner - Approved on December 19, 2024 at 12:23 PM UTC

looks good! and the appropriate "error" messages are now displayed 🥳👏🏼 

### Review by @swadeley - Commented on January 02, 2025 at 01:53 AM UTC

### Review by @swadeley - Commented on January 02, 2025 at 01:54 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/926*
