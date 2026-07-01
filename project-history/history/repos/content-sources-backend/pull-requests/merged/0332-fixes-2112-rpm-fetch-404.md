---
type: pull_request
number: 332
title: "Fixes 2112: RPM fetch 404"
state: merged
author: dpang314
created: 2023-07-12T20:55:09Z
updated: 2023-07-17T14:30:17Z
closed: 2023-07-17T14:21:34Z
merged: 2023-07-17T14:21:34Z
base_branch: main
head_branch: HMS-2112-rpm-fetch-error
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/332
---

# Pull Request #332: Fixes 2112: RPM fetch 404

**Author**: @dpang314
**Created**: July 12, 2023 at 08:55 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2112-rpm-fetch-error`

## Description

## Summary

Changes RPM list endpoint to return a status 404 instead of 500 when the repository uuid is not found. 

~~GORM throws an error when an invalid uuid is passed in (i.e. invalid syntax) which is returned as a status 500. The 404 only occurs when a valid, but nonexistent uuid is passed.~~

The endpoint will return a 404 if an invalid uuid is passed in. I changed the `isOwnedRepository` method to compare the uuid as `text`, which is consistent with other endpoints.

## Testing steps

```
curl localhost:8000/api/content-sources/v1.0/repositories/{some nonexistent repository config uuid}/rpms  -H "`./scripts/header.sh 1 user`"
```

---

## Discussion

### Comment by @jlsherrill on July 12, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-2112

### Comment by @swadeley on July 17, 2023 at 01:11 PM UTC

/retest

### Comment by @swadeley on July 17, 2023 at 01:34 PM UTC

Hi

I have a repo:

```
In [4]: app.content_sources.rest_client.repositories_api.get_repository('c02ad914-b81c-4770-ac2e-b87557820616')
2023-07-17 14:32:14.314 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=0PbeIzWYkJ9azjAnSwyPBCrrQCtJEiO8, params=[]
Out[4]: 
{'account_id': '12345',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2023-07-17T13:31:33Z',
 'last_success_introspection_time': '2023-07-17T13:31:33Z',
 'last_update_introspection_time': '2023-07-17T13:31:33Z',
 'metadata_verification': False,
 'name': 'fedoratest-no-snapshot',
 'org_id': '3340851',
 'package_count': 1,
 'snapshot': False,
 'status': 'Valid',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo02/',
 'uuid': 'c02ad914-b81c-4770-ac2e-b87557820616'}

```

Change the last digit of the UUID:

```
In [5]: app.content_sources.rest_client.repositories_api.get_repository('c02ad914-b81c-4770-ac2e-b87557820617')
2023-07-17 14:32:21.245 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=ItBiRno5ysgTOWWpFQlEpFE1iwfg8YC3, params=[]

NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '148', 'content-type': 'application/json; tZ<snip>
HTTP response body: {"errors":[{"status":404,"title":"Error fetching repository","detail":"Could not find repository with UUID c02ad914-b81c-4770-ac2e-b87557820617"}]}
```

so that looks expected, i.e. new in this PR.

### Comment by @swadeley on July 17, 2023 at 01:38 PM UTC

Hi

with shortened UUID (i.e. invalid):

```
In [6]: app.content_sources.rest_client.repositories_api.get_repository('c02ad914-b81c-4770-ac2e-b8755782061')

HTTP response body: {"errors":[{"status":404,"title":"Error fetching repository","detail":"Could not find repository with UUID c02ad914-b81c-4770-ac2e-b8755782061"}]}
```

that is not expected.


### Comment by @swadeley on July 17, 2023 at 02:20 PM UTC

Hi

test listing RPMS in repo:

Before this PR, using a valid UUID for a non existent repo:

`HTTP response body: {"errors":[{"status":500,"title":"Error listing RPMs","detail":"repositoryConfigUUID = 31169221-61c2-4bd8-a5eb-10608e051da5 is not owned"}]}`


With this PR:

```
In [7]: app.content_sources.rest_client.repositories_api.list_repositories_rpms('c02ad914-b81c-4770-ac2e-b87557820616')
2023-07-17 15:17:55.235 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=TfokTfjyqCCOsXHreEZlNcPGEgbrOJAF, params=[]
Out[7]: 
{'data': [{'arch': 'noarch',
           'checksum': '9b3d22d05187810d8521d99ca2483232e7da80087691e5c1f8fa106a25118bdf',
           'epoch': 0,
           'name': 'cockateel',
           'release': '1',
           'summary': 'A dummy package of cockateel',
           'uuid': '6d8e7f42-5b5b-40e8-958d-55ffe9e55b75',
           'version': '3.1'}],
 'links': {'first': '/api/content-sources/v1/repositories/c02ad914-b81c-4770-ac2e-b87557820616/rpms?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/c02ad914-b81c-4770-ac2e-b87557820616/rpms?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```


now change last digit of UUID (so valid but non existing repo):

```
In [8]: app.content_sources.rest_client.repositories_api.list_repositories_rpms('c02ad914-b81c-4770-ac2e-b87557820617'

HTTP response body: {"errors":[{"status":404,"title":"Error listing RPMs","detail":"Could not find repository with UUID c02ad914-b81c-4770-ac2e-b87557820617"}]}
```

OK, good.

---

## Reviews

### Review by @rverdile - Commented on July 13, 2023 at 08:21 PM UTC

working nicely! one question

### Review by @dpang314 - Commented on July 13, 2023 at 08:57 PM UTC

### Review by @rverdile - Approved on July 13, 2023 at 09:24 PM UTC

nice!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/332*
