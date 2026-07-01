---
type: pull_request
number: 315
title: "Fixes 1579: update url on remote if url changed"
state: merged
author: jlsherrill
created: 2023-06-22T15:28:45Z
updated: 2023-07-05T18:00:29Z
closed: 2023-07-05T17:55:53Z
merged: 2023-07-05T17:55:52Z
base_branch: main
head_branch: update_2
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/315
---

# Pull Request #315: Fixes 1579: update url on remote if url changed

**Author**: @jlsherrill
**Created**: June 22, 2023 at 03:28 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `update_2`

## Description

## Summary

Supports updating url of a repository with snapshotting enabled.
Will update the remote to the new url as part of the snapshot
If the url is being updated, and a snapshot is in progress, we reject the request

## Testing steps

make this request:
```
POST http://localhost:8000/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqZG9lIn0sImludGVybmFsIjp7Im9yZ19pZCI6IjEyMyJ9fX0K

{"name":"sample zoo","url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/","snapshot": true}
```

Then update the repo (change the uuid to match what is returned):

```
PATCH http://localhost:8000/api/content-sources/v1.0/repositories/1538f16f-324e-4d8e-95d4-483d73b3cebd
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqZG9lIn0sImludGVybmFsIjp7Im9yZ19pZCI6IjEyMyJ9fX0K
```

Then check the snapshots table:

```
make db-cli-connect
select * from snapshots
```

you should see 2 created snapshots now.  

---

## Discussion

### Comment by @jlsherrill on June 22, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-1579

### Comment by @jlsherrill on June 29, 2023 at 03:49 PM UTC

/retest

### Comment by @swadeley on June 30, 2023 at 09:55 AM UTC

Hi

Would changing `Restrict architecture` value be enough to trigger a new snapshot?

In which pod would the `make db-cli-connect` command work?

Thank you



### Comment by @swadeley on July 03, 2023 at 01:38 PM UTC

> Hi
> 
> Would changing `Restrict architecture` value be enough to trigger a new snapshot?
> 
> In which pod would the `make db-cli-connect` command work?
> 
> Thank you

Justin Says:
Edit repo URL (i.e. change it), list snapshots, see the change

### Comment by @swadeley on July 03, 2023 at 01:40 PM UTC

/retest

### Comment by @swadeley on July 05, 2023 at 05:46 PM UTC

Hi

Created a repo, listed its snapshots:

In [9]: app.content_sources.rest_client.repositories_api.list_snapshots('80fc43a0-520c-45e0-b755-e814e3f619ad')
2023-07-05 18:31:14.070 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=v5HNU8WEjxMukwwzmVq7Ygmy6Ot7zvmZ, params=[]
Out[9]: 
{'data': [{'content_counts': {'rpm.package': 1},
           'created_at': '2023-07-05T17:26:34.957859Z',
           'distribution_path': '80fc43a0-520c-45e0-b755-e814e3f619ad/c2fbb229-af22-476d-bebd-037d050c44ba'}],
 'links': {'first': '/api/content-sources/v1/repositories/80fc43a0-520c-45e0-b755-e814e3f619ad/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/80fc43a0-520c-45e0-b755-e814e3f619ad/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}


Change URL to point to new repo with different packages:

Out[18]: 
{'data': [{'content_counts': {'rpm.package': 1},
           'created_at': '2023-07-05T17:26:34.957859Z',
           'distribution_path': '80fc43a0-520c-45e0-b755-e814e3f619ad/c2fbb229-af22-476d-bebd-037d050c44ba'},
          {'content_counts': {'rpm.advisory': 4, 'rpm.package': 32},
           'created_at': '2023-07-05T17:44:57.23853Z',
           'distribution_path': '80fc43a0-520c-45e0-b755-e814e3f619ad/5259c272-6770-4e66-bfd6-5c55201224d3'}],
 'links': {'first': '/api/content-sources/v1/repositories/80fc43a0-520c-45e0-b755-e814e3f619ad/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/80fc43a0-520c-45e0-b755-e814e3f619ad/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

There we see :  {'rpm.advisory': 4, 'rpm.package': 32},
and : {'count': 2

Thank you


---

## Reviews

### Review by @rverdile - Commented on June 30, 2023 at 08:01 PM UTC

### Review by @rverdile - Commented on June 30, 2023 at 08:05 PM UTC

### Review by @jlsherrill - Commented on June 30, 2023 at 08:18 PM UTC

### Review by @jlsherrill - Commented on June 30, 2023 at 08:31 PM UTC

### Review by @rverdile - Approved on June 30, 2023 at 08:47 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/315*
