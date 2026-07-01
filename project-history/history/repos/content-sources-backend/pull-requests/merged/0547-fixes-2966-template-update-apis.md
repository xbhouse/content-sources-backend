---
type: pull_request
number: 547
title: "Fixes 2966: Template update apis"
state: merged
author: jlsherrill
created: 2024-01-23T18:01:38Z
updated: 2024-02-01T15:00:38Z
closed: 2024-02-01T14:49:46Z
merged: 2024-02-01T14:49:46Z
base_branch: main
head_branch: temp_patch
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/547
---

# Pull Request #547: Fixes 2966: Template update apis

**Author**: @jlsherrill
**Created**: January 23, 2024 at 06:01 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `temp_patch`

## Description

## Summary

## Testing steps

After creating a repo, create a template:
```
####
POST http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
Content-Type: application/json

{
  "name": "needed",
  "repository_uuids": [
    "5e6937b5-bd6f-4c5e-8338-f62a231728ff"
  ],
  "date": "2023-12-31T15:04:05Z",
  "version": "8",
  "arch": "x86_64"
}

# Fetch the templates and get the uuid

####
GET http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
Content-Type: application/json

# partially  update the template
###
PATCH http://localhost:8000/api/content-sources/v1.0/templates/7a69e234-f21a-4072-a15f-72e3836fb667
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
Content-Type: application/json

{
  "description": "myfoo",
}



# fully update the template

###
PUT http://localhost:8000/api/content-sources/v1.0/templates/7a69e234-f21a-4072-a15f-72e3836fb667
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
Content-Type: application/json

{
  "name": "needed",
  "repository_uuids": [
    "5e6937b5-bd6f-4c5e-8338-f62a231728ff"
  ],
  "date": "2023-12-01T15:04:05Z",
  "version": "8",
  "arch": "x86_64"
}
```
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 23, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-2966

### Comment by @xbhouse on January 26, 2024 at 03:51 PM UTC

if we have a template for a rhel8 repository, should we be able to update that template version to 7 or 9? i know we can create a template based on a rhel8 repo with a template version of 7/9, so maybe i just don't understand what template version is referring to

### Comment by @xbhouse on January 26, 2024 at 03:52 PM UTC

i'm able to send PUT and PATCH requests, but i don't see the PATCH template in the api docs. was that intentional? 

### Comment by @xbhouse on January 26, 2024 at 03:59 PM UTC

update code looks good to me! 

i see you added in repo uuids in the response, and those are listed in the response after updating and fetching templates, but repo uuids are blank when listing templates:
`
{
  "uuid": "b0022b97-ff01-4d2e-8357-d0872864d06c",
  "name": "test-template4",
  "org_id": "16799052",
  "description": "",
  "arch": "x86_64",
  "version": "8",
  "date": "2023-12-01T10:04:05-05:00",
  "repository_uuids": []
}
`

### Comment by @jlsherrill on January 26, 2024 at 04:55 PM UTC

you know, thats a good point, we should likely block arch/version updates.  I'll update with that

### Comment by @Andrewgdewar on January 26, 2024 at 10:01 PM UTC

> you know, thats a good point, we should likely block arch/version updates. I'll update with that

hmm I just finished the front-end "edit modal" using this code, let me know when you have updated this, I'll make those fields read only in the modal, and test again.

### Comment by @swadeley on January 29, 2024 at 08:46 AM UTC

/retest

### Comment by @jlsherrill on January 29, 2024 at 02:13 PM UTC

updated:
* make version & arch not able to be updated
* fixed list response not returning repositoryUUIDs
* fixed patch api not showing up in api docs

Nice review @xbhouse !

### Comment by @xbhouse on January 29, 2024 at 02:56 PM UTC

lgtm! 

### Comment by @swadeley on February 01, 2024 at 10:13 AM UTC

Hi

built APIdocs

Created a repo

Created this template:

```
In [5]: app.content_sources.rest_client.templates_api.create_template(dict(
   ...:    name="my template 1",
   ...:    repository_uuids=["0db4752c-97bc-48ca-a31f-40cc4db8b7ca"],
   ...:    description="my new template",
   ...:    arch="x86_64",
   ...:    version="7"
   ...:    ))
2024-02-01 09:45:40.947 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[5]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my new template',
 'name': 'my template 1',
 'org_id': '3340851',
 'repository_uuids': ['0db4752c-97bc-48ca-a31f-40cc4db8b7ca'],
 'uuid': 'b898c98b-a61c-4ed4-b467-a3adef312e9d',
 'version': '7'}

```

### partially  update the template

```
In [8]: app.content_sources.rest_client.templates_api.partial_update_template('b898c98b-a61c-4ed4-b467-a3adef312e9d', dict(description="myfoo",))
2024-02-01 10:00:40.172 [    INFO] [iqe.base.rest_client] REST: METHOD=PATCH, request_id=<>, params=[]
Out[8]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'myfoo',
 'name': 'my template 1',
 'org_id': '3340851',
 'repository_uuids': ['0db4752c-97bc-48ca-a31f-40cc4db8b7ca'],
 'uuid': 'b898c98b-a61c-4ed4-b467-a3adef312e9d',
 'version': '7'}
```


### fully update the template

I tried this:

```
In [10]: app.content_sources.rest_client.templates_api.full_update_template('b898c98b-a61c-4ed4-b467-a3adef312e9d',dict(
    ...:   name="needed",
    ...:   repository_uuids=[
    ...:     "5e6937b5-bd6f-4c5e-8338-f62a231728ff"
    ...:   ],
    ...:   date="2023-12-01T15:04:05Z",
    ...:   version="8",
    ...:   arch="x86_64"
    ...: ))
```

but got error

`ApiAttributeError: ApiTemplateUpdateRequest has no attribute 'name' at ['api_template_update_request']['name']
`

EDIT: I see now I should not try to update version and arch, will test again.





### Comment by @swadeley on February 01, 2024 at 01:31 PM UTC

/retest

### Comment by @swadeley on February 01, 2024 at 02:40 PM UTC

Hi

Full update works with these keys:

```

In [16]: app.content_sources.rest_client.templates_api.full_update_template('e5793156-dc2e-4f5e-ab44-e18b3f2723be',dict(
    ...:   repository_uuids=[
    ...:     "b115914c-782c-4252-a5d6-98a79ab9488f"
    ...:   ],
    ...:   description= 'Full update',
    ...:   date="2023-12-01T15:04:05Z",
    ...:   ))
2024-02-01 14:39:42.748 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>Z, params=[]
Out[16]: 
{'arch': 'x86_64',
 'date': '2023-12-01T15:04:05Z',
 'description': 'Full update',
 'name': 'my template one',
 'org_id': '3340851',
 'repository_uuids': ['b115914c-782c-4252-a5d6-98a79ab9488f'],
 'uuid': 'e5793156-dc2e-4f5e-ab44-e18b3f2723be',
 'version': '7'}

```
You can add more repo UUIDs provide the repo UUID is valid and the repo exists:


```
In [20]: app.content_sources.rest_client.templates_api.full_update_template('e5793156-dc2e-4f5e-ab44-e18b3f2723be',dict(
    ...:   repository_uuids=[
    ...:     "b115914c-782c-4252-a5d6-98a79ab9488f","d3da09d7-3572-4592-8cc8-dcbf9b8b445c"
    ...:   ],
    ...:   description= 'Full update',
    ...:   date="2023-12-01T15:04:05Z",
    ...:   ))
2024-02-01 14:46:04.974 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=<>, params=[]
Out[20]: 
{'arch': 'x86_64',
 'date': '2023-12-01T15:04:05Z',
 'description': 'Full update',
 'name': 'my template one',
 'org_id': '3340851',
 'repository_uuids': ['b115914c-782c-4252-a5d6-98a79ab9488f',
                      'd3da09d7-3572-4592-8cc8-dcbf9b8b445c'],
 'uuid': 'e5793156-dc2e-4f5e-ab44-e18b3f2723be',
 'version': '7'}

```

---

## Reviews

### Review by @jlsherrill - Commented on January 23, 2024 at 07:52 PM UTC

### Review by @xbhouse - Approved on January 29, 2024 at 02:56 PM UTC

### Review by @Andrewgdewar - Approved on January 30, 2024 at 02:29 PM UTC

### Review by @swadeley - Commented on February 01, 2024 at 08:36 AM UTC

### Review by @swadeley - Commented on February 01, 2024 at 08:36 AM UTC

### Review by @swadeley - Commented on February 01, 2024 at 08:38 AM UTC

### Review by @swadeley - Commented on February 01, 2024 at 08:38 AM UTC

### Review by @swadeley - Commented on February 01, 2024 at 09:05 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/547*
