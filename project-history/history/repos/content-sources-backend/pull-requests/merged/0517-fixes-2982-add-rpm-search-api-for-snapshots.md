---
type: pull_request
number: 517
title: "Fixes 2982: add rpm search api for snapshots"
state: merged
author: jlsherrill
created: 2024-01-03T17:30:11Z
updated: 2024-01-18T21:30:33Z
closed: 2024-01-18T21:13:24Z
merged: 2024-01-18T21:13:24Z
base_branch: main
head_branch: 2982
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/517
---

# Pull Request #517: Fixes 2982: add rpm search api for snapshots

**Author**: @jlsherrill
**Created**: January 03, 2024 at 05:30 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2982`

## Description

## Summary

Adds support for tang at startup of the api server and adds an api to search rpm names within a set of snapshots.

## Testing steps

1.  Copy the new example config section to your config.yaml
2.  Create 1 or more repos for snapshotting
3.   list the repos to get a snapshot UUID:
```
GET http://localhost:8000/api/content-sources/v1.0/repositories/
```
4.  Use the snapshot UUIDs to search for packages:

```
####
POST http://localhost:8000/api/content-sources/v1.0/snapshots/rpms/names
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
Content-Type: application/json

{
  "uuids": ["26181d5d-85fb-4fc6-b65e-3737c5c7dce5"],
  "search": "w"
}
```

Try with more than one uuid as well.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 08, 2024 at 05:23 PM UTC

https://issues.redhat.com/browse/HMS-2982

### Comment by @jlsherrill on January 08, 2024 at 06:56 PM UTC

/retest

### Comment by @jlsherrill on January 08, 2024 at 07:43 PM UTC

/retest

### Comment by @jlsherrill on January 09, 2024 at 03:27 PM UTC

updated

### Comment by @rverdile on January 10, 2024 at 07:29 PM UTC

Right now, if I disable the snapshot feature or unset the pulp server, I get a 500 when I make a request.
`"no tang configuration present"`

Do we want to catch that error in the handler and return a more helpful error message, or just let it error?

### Comment by @jlsherrill on January 10, 2024 at 09:01 PM UTC

we probably should return a bad request, i'll update!

### Comment by @jlsherrill on January 12, 2024 at 07:39 PM UTC

converted to draft until we get migrated to new pulp deployment

### Comment by @jlsherrill on January 17, 2024 at 04:44 PM UTC

adjusted to work with the new central pulp instance.  Will test this in ephemeral

### Comment by @swadeley on January 18, 2024 at 09:13 PM UTC

Hi

Created two repos, set them to make snapshots.

with one snapshot UUID:

```
In [10]: app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search='cockateel', uuids=['0c190ee8-9cfe-49f6-9e29-6a9e4dd8ad5d'],))
2024-01-18 18:41:24.721 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, <snip>, params=[]
Out[10]: [{'package_name': 'cockateel', 'summary': 'A dummy package of cockateel'}]
```

with two snapshot UUIDs:
`In [7]: app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search='cockateel', uuids=['0c190ee8-9cfe-49f6-9e29-6a9e4dd8ad5d','9b2370d3-f870-4cbb-87b8-d87f8824287d'],))
2024-01-18 21:10:48.780 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, <snip>, params=[]
Out[7]: [{'package_name': 'cockateel', 'summary': 'A dummy package of cockateel'}]
`

You can search on part of name, but not using package versions:

```
In [8]: app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search='co', uuids=['0c190ee8-9cfe-49f6-9e29-6a9e4dd8ad5d','9b2370d3-f870-4cbb-87b8-d87f8824287d'],))
2024-01-18 21:12:11.078 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, <snip>T, params=[]
Out[8]: 
[{'package_name': 'cockateel', 'summary': 'A dummy package of cockateel'},
 {'package_name': 'cow', 'summary': 'A dummy package of cow'}]
```


All good, thank you.

---

## Reviews

### Review by @rverdile - Commented on January 04, 2024 at 07:24 PM UTC

### Review by @jlsherrill - Commented on January 05, 2024 at 02:16 PM UTC

### Review by @jlsherrill - Commented on January 09, 2024 at 01:48 PM UTC

### Review by @jlsherrill - Commented on January 09, 2024 at 01:48 PM UTC

### Review by @jlsherrill - Commented on January 09, 2024 at 03:27 PM UTC

### Review by @rverdile - Commented on January 10, 2024 at 03:48 PM UTC

### Review by @rverdile - Commented on January 10, 2024 at 04:25 PM UTC

### Review by @jlsherrill - Commented on January 10, 2024 at 04:58 PM UTC

### Review by @rverdile - Approved on January 11, 2024 at 05:07 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/517*
