---
type: pull_request
number: 846
title: "Fixes 4825: add template attribute for when env is created"
state: merged
author: jlsherrill
created: 2024-10-11T19:21:10Z
updated: 2024-10-18T10:55:20Z
closed: 2024-10-18T08:36:47Z
merged: 2024-10-18T08:36:47Z
base_branch: main
head_branch: 4825
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/846
---

# Pull Request #846: Fixes 4825: add template attribute for when env is created

**Author**: @jlsherrill
**Created**: October 11, 2024 at 07:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4825`

## Description

## Summary

adds an rhsm_environment_created db column and api attribute.  It is false when a template is first created.  Once it is created via the update-template-content task, it gets set to true.

## Testing steps

1.  create and snapshot one repository
2.  Create a template with the repository:
```
####
POST http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "name":"test7",
  "arch": "x86_64",
  "version": "9",
  "repository_uuids": ["270fabbd-ca28-4cd0-9cff-a274b4e797b9"],
  "date": "2024-09-10T15:09:43-04:00"
}
```

3.  check teh return:

  "rhsm_environment_created": false

4.  fetch the created template:
```
###
GET http://localhost:8000/api/content-sources/v1.0/templates/687c93e2-b146-428d-b65d-a972bed5edbe
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json
```

after a few seconds of creation, it should now be true:
```
  "rhsm_environment_created": true
```



---

## Discussion

### Comment by @jlsherrill on October 11, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-4825

### Comment by @xbhouse on October 14, 2024 at 08:34 PM UTC

just the one question, overall looks good and works well!

### Comment by @mayurilahane on October 15, 2024 at 02:57 PM UTC

/retest

### Comment by @mayurilahane on October 17, 2024 at 09:53 AM UTC

Created a repository 

```
In [1]: app.content_sources.rest_client.repositories_api.create_repository(dict(name='test', origin='upload', snapshot=True))

Out[1]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-10-17T09:16:08Z',
 'last_snapshot_task_uuid': '9263a30f-142a-4694-9316-9cb37b4f5209',
 'last_success_introspection_time': '2024-10-17T09:16:08Z',
 'last_update_introspection_time': '2024-10-17T09:16:08Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test',
 'org_id': '3340851',
 'origin': 'upload',
 'package_count': 0,
 'snapshot': True,
 'status': 'Valid',
 'url': '',
 'uuid': 'f5b987cc-aecb-4373-b6ad-f38b5c4f2cac'}
```

### Comment by @mayurilahane on October 17, 2024 at 09:54 AM UTC

Created a template and "rhsm_environment_created": false 

```
In [3]: app.content_sources.rest_client.templates_api.create_template(dict(name="test2", version="8", arch="x86_64", repository_uuids=["f5b987cc-aecb-4373-b6ad-f38b5
   ...: c4f2cac"], use_latest=True))

Out[3]: 
{'arch': 'x86_64',
 'created_at': '2024-10-17T09:16:39.914157359Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test2',
 'org_id': '3340851',
 'repository_uuids': ['f5b987cc-aecb-4373-b6ad-f38b5c4f2cac'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': '96bc0f7d3ce54c6a92d0d2afb535eeff',
 'updated_at': '2024-10-17T09:16:39.914157359Z',
 'use_latest': True,
 'uuid': '96bc0f7d-3ce5-4c6a-92d0-d2afb535eeff',
 'version': '8'}
```

### Comment by @mayurilahane on October 17, 2024 at 09:58 AM UTC

fetch the created template:

```
In [12]: app.content_sources.rest_client.templates_api.get_template("96bc0f7d-3ce5-4c6a-92d0-d2afb535eeff")
2024-10-17 15:26:07.302 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=4fe45e76-8510-4307-bf18-fcedb6fa0b39, params=[]
Out[12]: 
{'arch': 'x86_64',
 'created_at': '2024-10-17T09:16:39.914157Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test2',
 'org_id': '3340851',
 'repository_uuids': ['f5b987cc-aecb-4373-b6ad-f38b5c4f2cac'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': '96bc0f7d3ce54c6a92d0d2afb535eeff',
 'updated_at': '2024-10-17T09:16:39.914157Z',
 'use_latest': True,
 'uuid': '96bc0f7d-3ce5-4c6a-92d0-d2afb535eeff',
 'version': '8'}
```

The ` 'rhsm_environment_created': False,` is still set to false

### Comment by @mayurilahane on October 18, 2024 at 08:36 AM UTC

lgtm!

### Comment by @mayurilahane on October 18, 2024 at 10:55 AM UTC

Verified again on stage after merging this PR. 
The get_template function now returns 
`rhsm_environment_created: True` after some time 

---

## Reviews

### Review by @xbhouse - Commented on October 14, 2024 at 08:33 PM UTC

### Review by @jlsherrill - Commented on October 17, 2024 at 01:27 PM UTC

### Review by @jlsherrill - Commented on October 17, 2024 at 03:52 PM UTC

### Review by @xbhouse - Approved on October 17, 2024 at 04:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/846*
