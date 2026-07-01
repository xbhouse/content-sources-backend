---
type: pull_request
number: 669
title: "Fixes 3948: fix resp codes if resource exists/not found"
state: merged
author: xbhouse
created: 2024-05-13T20:21:02Z
updated: 2024-05-28T21:00:35Z
closed: 2024-05-28T20:49:48Z
merged: 2024-05-28T20:49:48Z
base_branch: main
head_branch: 3948
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/669
---

# Pull Request #669: Fixes 3948: fix resp codes if resource exists/not found

**Author**: @xbhouse
**Created**: May 13, 2024 at 08:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3948`

## Description

## Summary

- Changes response code when searching for a resource that is not found from 400 to 404
- Changes response code when creating a resource that already exists from 400 to 409
- Also updates the error returned in the template update method. If the template wasn't found when trying to PUT or PATCH, this was returning a 500

## Testing steps

- Send request with an invalid uuid (or an invalid URL if applicable) to the following endpoints, they should return a 404:
```
search repository environments --> POST /environments/names
search snapshot environments --> POST /snapshots/environments/names
search repository package groups --> POST /package_groups/names
search snapshot package groups --> POST /snapshots/package_groups/names
search repository rpms --> POST /rpms/names
search snapshot rpms --> POST /snapshots/rpms/names
fully update template --> PUT /templates/:uuid
partially update template --> PATCH /templates/:uuid
list snapshot rpms --> GET /snapshots/:uuid/rpms
list snapshot errata --> GET /snapshots/:uuid/errata
```
- Try to create a repository with the same name or URL as one that already exists, this should return a 409
- Try to create a template with the same name as one that already exists, this should also return a 409


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 13, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3948

### Comment by @mayurilahane on May 15, 2024 at 06:37 PM UTC

/retest

### Comment by @swadeley on May 16, 2024 at 02:11 PM UTC

/retest

### Comment by @swadeley on May 16, 2024 at 05:03 PM UTC

@xbhouse this needs a rebase to pull in the fix to `"Red Hat Enterprise Linux 9 for ARM 64 - AppStream (RPMs)"` label. 

Thank you

### Comment by @mayurilahane on May 16, 2024 at 08:51 PM UTC

/retest


### Comment by @mayurilahane on May 16, 2024 at 11:48 PM UTC

1. **search repository environments --> POST /environments/names**

>> app.content_sources.rest_client.environments_api.list_repositories_environments(uuid="12344")

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '118', 'content-type': 'application/json', 'date': 'Thu, 16 May 2024 23:23:17 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGVudGl0190cmlhbCI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': 'ba4d42df-22be-44dc-8bef-9b98a6dd0e28', 'set-cookie': '773bb68ee176cccf9c5308b57a675a82=7008831994368d0aeff17f96f1fabc39; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error listing environments","detail":"Could not find repository with UUID 12344"}]}
```


2. **search snapshot environments --> POST /snapshots/environments/names**
>> app.content_sources.rest_client.snapshots_api.search_snapshot_environments(dict(uuids=['zzzzzzzz-af91-4fa7-be24-000b233969a5'],search='avians'))
```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '150', 'content-type': 'application/json', 'date': 'Wed, 22 May 2024 16:46:06 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJ106efabb87050d3f02a06e8f16; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error searching environments","detail":"Could not find snapshot with UUID: zzzzzzzz-af91-4fa7-be24-000b233969a5"}]}
```

3.  **search repository package groups --> POST /package_groups/names**

>> app.content_sources.rest_client.packagegroups_api.list_repositories_package_groups("123")

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '118', 'content-type': 'application/json', 'date': 'Thu, 16 May 2024 23:29:02 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGc190cmlhbCI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': 'b324ab50-9940-4b70-b69b-08e335612501', 'set-cookie': '773bb68ee176cccf9c5308b57a675a82=7008831994368d0aeff17f96f1fabc39; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error listing package groups","detail":"Could not find repository with UUID 123"}]}
```

4. **search snapshot package groups --> POST /snapshots/package_groups/names**
 >>  app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(search="KDE", uuids=["fd3a7bef-2a64-48e0-ac4a-d38c04cd6357"]))

```NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '152', 'content-type': 'application/json', 'date': 'Wed, 22 May 2024 21:36:59 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyZmFsc2V9fX0=', 'x-rh-insights-request-id': '605b9a41-d2a3-4a93-9e26-3ea944157492', 'set-cookie': '03e47ae7a48d978a970b3367e91cedb6=33a2cb106efabb87050d3f02a06e8f16; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error searching package groups","detail":"Could not find snapshot with UUID: fd3a7bef-2a64-48e0-ac4a-d38c04cd6357"}]}
```

5. **search repository rpms --> POST /rpms/names**

>>app.content_sources.rest_client.repositories_api.search_rpm(dict(search="bear", urls=["https://fixtures.pulpproject.org/rpm-unsigned/"]))

```
In [36]: app.content_sources.rest_client.repositories_api.search_rpm(dict(search="hakfj", urls=["http://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"]))
2024-05-22 17:50:19.231 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=5ab8cb45-b715-49e1-95c0-f2b5eeb1c367, params=[]
Out[36]: []

```
6. **search snapshot rpms --> POST /snapshots/rpms/names**

>>app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search="asjdgaksf", uuids=["1e5b786e-fdcc-4198-95f5-dba4903aaad8"]))

```
In [48]: app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search="asjdgaksf", uuids=["1e5b786e-fdcc-4198-95f5-dba4903aaad8"]))
2024-05-22 18:03:19.559 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=ea30345a-2323-4103-892a-d669b370d9d6, params=[]
Out[48]: []
```

7. **fully update template --> PUT /templates/:uuid**
>>app.content_sources.rest_client.templates_api.full_update_template('e5793156-dc2e-4f5e-ab44-e18b3f2723be',dict(repository_uuids=["b115914c-782c-4252-a5d6-98a79ab9488f"], description= 'Full update', date="2023-12-01T15:04:05Z",))

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '144', 'content-type': 'application/json', 'date': 'Tue, 28 May 2024 19:49:22 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJ50fb', 'set-cookie': 'eb0bac5d9937780cb7a9414bcd323e79=1a1280cbc7f49015291e4fd295dbda22; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error updating template","detail":"Could not find template with UUID e5793156-dc2e-4f5e-ab44-e18b3f2723be"}]}
```

8. **partially update template --> PATCH /templates/:uuid**
>>app.content_sources.rest_client.templates_api.partial_update_template('b898c98b-a61c-4ed4-b467-a3adef312e9d', dict(description="myfoo",))

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '144', 'content-type': 'application/json', 'date': 'Tue, 28 May 2024 19:45:44 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJp0bdGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': '4f14ae35-50fb-4425-91fb-e5eb4925f5b0', 'set-cookie': 'eb0bac5d9937780cb7a9414bcd323e79=1a1280cbc7f49015291e4fd295dbda22; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error updating template","detail":"Could not find template with UUID b898c98b-a61c-4ed4-b467-a3adef312e9d"}]}
```

9. **list snapshot rpms --> GET /snapshots/:uuid/rpms**

>> app.content_sources.rest_client.snapshots_api.list_snapshot_rpms(uuid="ahf")

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '107', 'content-type': 'application/json', 'date': 'Thu, 23 May 2024 18:46:16 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGVudGl0ehbCI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': '402677e0-5632-4831-9ee0-02677d074ac6', 'set-cookie': 'b58642189c08a7fd2efcc58ddb02f46b=0b504b6c8528cd1c574f12944852f0bc; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error listing RPMs","detail":"Could not find snapshot with UUID: ahf"}]}
```

10. **list snapshot errata --> GET /snapshots/:uuid/errata**

>>app.content_sources.rest_client.snapshots_api.list_snapshot_errata(uuid="a")

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '107', 'content-type': 'application/json', 'date': 'Thu, 23 May 2024 18:48:40 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGVudGl0eSI6eyJhY2NvdW50X251bWJlciI6IjAzNjkyMzMiLCJvcmdfaWQiOiIzMzQwODUxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMzM0MDg1MSJ9LCJ1c2VyIjp7InVzZXJuYW1lIjoiZXBoZW1lcmFsLXVzZXI2f0bc; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error listing Errata","detail":"Could not find snapshot with UUID: a"}]}
```

11.  **Try to create a repository with the same name or URL as one that already exists, this should return a 409**

>>app.content_sources.rest_client.repositories_api.create_repository({"snapshot": True, "name": "test", "url": fake_data.url})

```
ApiException: (409)
Reason: Conflict
HTTP response headers: HTTPHeaderDict({'content-length': '132', 'content-type': 'application/json', 'date': 'Thu, 16 May 2024 23:41:09 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGcmlhbCI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': '2f62a4e0cea043', 'set-cookie': '773bb6a82=7008831994f1fabc39; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":409,"title":"Error creating repository","detail":"Repository with this URL already belongs to organization"}]}
```

12. **Try to create a template with the same name as one that already exists, this should also return a 409**

>>app.content_sources.rest_client.templates_api.create_template(dict(name="my template 1", repository_uuids=["e93b6af3-187d-4f4d-82a8-0a569a515066"], description="my new template", arch="x86_64", version="7")) 

```

ApiException: (409)
Reason: Conflict
HTTP response headers: HTTPHeaderDict({'content-length': '129', 'content-type': 'application/json', 'date': 'Tue, 28 May 2024 20:21:49 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGV3f8468', 'set-cookie': '00faa4ebdc53c9c06668c6e017223c50=95c8d3f2de47b0a7879e452e6641e235; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":409,"title":"Error creating template","detail":"Template with this name already belongs to organization"}]}

```


### Comment by @xbhouse on May 17, 2024 at 03:49 PM UTC

@mayurilahane it looks like `app.content_sources.rest_client.snapshots_api.list_snapshot_rpms("1234")` is calling this endpoint `GET /snapshots/:uuid/rpms` and not `POST /snapshots/rpms/names`, but the list snapshot rpms endpoint should also return a 404 if given an invalid uuid. i missed changing that one :) thank you 

### Comment by @swadeley on May 22, 2024 at 10:01 AM UTC

/retest

### Comment by @mayurilahane on May 28, 2024 at 03:10 PM UTC

/retest

### Comment by @mayurilahane on May 28, 2024 at 07:52 PM UTC

when I am trying to create a template with the wrong repo uuid it returns 400 

Expected is 404.

Re: @xbhouse added update for this case 

---

## Reviews

### Review by @rverdile - Commented on May 14, 2024 at 06:08 PM UTC

### Review by @xbhouse - Commented on May 14, 2024 at 06:39 PM UTC

### Review by @rverdile - Commented on May 14, 2024 at 07:03 PM UTC

### Review by @rverdile - Commented on May 14, 2024 at 07:04 PM UTC

### Review by @xbhouse - Commented on May 14, 2024 at 07:12 PM UTC

### Review by @rverdile - Commented on May 14, 2024 at 07:14 PM UTC

### Review by @rverdile - Commented on May 14, 2024 at 07:34 PM UTC

### Review by @rverdile - Commented on May 14, 2024 at 07:35 PM UTC

### Review by @rverdile - Approved on May 14, 2024 at 08:25 PM UTC

looks good! Up to you if you want to add the change I suggested

### Review by @xbhouse - Commented on May 15, 2024 at 01:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/669*
