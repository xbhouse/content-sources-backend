---
type: pull_request
number: 909
title: "Fixes 5038: add uploads search endpoint"
state: closed
author: dominikvagner
created: 2024-11-26T12:08:48Z
updated: 2024-12-04T15:04:37Z
closed: 2024-12-04T15:04:36Z
base_branch: main
head_branch: 5038
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/909
---

# Pull Request #909: Fixes 5038: add uploads search endpoint

**Author**: @dominikvagner
**Created**: November 26, 2024 at 12:08 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `5038`

## Description

## Summary
This PR adds a new endpoint for checking if a file has already been uploaded (i.e. is present as an artifact in Pulp). The search can be done by a `sha256` hash of a file.

## Testing steps
1. Create an upload repo.
2. Get a hash of any rpm file you want to upload.
3. Search uploads for that hash by using the new endpoint:
    `http :8000/api/content-sources/v1.0/repositories/uploads/search "$( ./scripts/header.sh $YOUR_ORG_ID 1111)" sha256:='["$THE_HASH_OF_THE_RPM"]'`
4. Verify that the response contains the hash in the missing array.
5. Upload that rpm to the upload repo and wait for the upload to finish.
6. Repeat step 3 and verify the response has that rpm hash in the found array.


---

## Discussion

### Comment by @jlsherrill on November 26, 2024 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-5038

### Comment by @mayurilahane on November 28, 2024 at 10:26 AM UTC

/retest

### Comment by @swadeley on December 02, 2024 at 12:56 AM UTC

/retest

### Comment by @swadeley on December 02, 2024 at 12:17 PM UTC

Hi @dominikvagner 

Rebase please because I get strange results

Thank you

### Comment by @swadeley on December 03, 2024 at 06:44 AM UTC

Hi

Working now:
```
~$ sha256sum acme-package-1.0.1-1.noarch.rpm
d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866  acme-package-1.0.1-1.noarch.rpm

In [1]: app.content_sources.rest_client.repositories_api.search_uploads({'sha256':(["d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866"])})
2024-12-03 14:25:35.458 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-12-03 14:25:35.458 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-12-03 14:25:36.335 [    INFO] [root] Created RESTPluginService client for https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 with the following attributes: ['client', 'environments_api', 'features_api', 'gpg_key_api', 'package', 'packagegroups_api', 'popular_repositories_api', 'public_repositories_api', 'repositories_api', 'rpms_api', 'snapshots_api', 'tasks_api', 'templates_api']
2024-12-03 14:25:37.112 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/search with query params [] and x-rh-insights-request-id=<>
Out[1]: 
{'found': {},
 'missing': ['d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866']}

In [2]: repo_dict_1 = dict(
   ...:    ...:    name="test_upload_repo_1",
   ...:    ...:    origin="upload",
   ...:    ...:    snapshot=True,
   ...:    ...:    )

In [3]: myrepo = app.content_sources.rest_client.repositories_api.create_repository(repo_dict_1)
2024-12-03 14:26:20.378 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [] and x-rh-insights-request-id=<>

In [4]: app.content_sources.rest_client.repositories_api.create_upload({"size":8536})
2024-12-03 14:28:40.776 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'created': '2024-12-03T06:28:40.666674Z',
 'last_updated': '2024-12-03T06:28:40.66669Z',
 'size': 8536,
 'upload_uuid': '01938b34-e859-769d-814a-e998c49f97de'}

In [5]: with open('acme-package-1.0.1-1.noarch.rpm', 'rb') as file:
   ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('01938b34-e859-769d-814a-e998c49f97de', 'bytes 0-8535/*', file, 'd3c9d09b
   ...: ba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866')
   ...: 
2024-12-03 14:29:25.129 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/01938b34-e859-769d-814a-e998c49f97de/upload_chunk/ with query params [] and x-rh-insights-request-id=<>

In [6]: body = { 'uploads': [{'uuid':'01938b34-e859-769d-814a-e998c49f97de', 'sha256':'d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866'}], 'a
   ...: rtifacts': [] }

In [7]: app.content_sources.rest_client.repositories_api.add_upload(myrepo.uuid, body)
2024-12-03 14:38:17.158 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/c25683d9-1421-415f-96cd-61083597dff6/add_uploads/ with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'created_at': '2024-12-03T06:38:17Z',
 'dependents': ['6b9e95e5-91e6-4982-a647-ecbfeddef401'],
 'ended_at': '',
 'error': '',
 'object_name': 'test_upload_repo_1',
 'object_type': 'repository',
 'object_uuid': 'c25683d9-1421-415f-96cd-61083597dff6',
 'org_id': '3340851',
 'status': 'running',
 'type': 'add-uploads-repository',
 'uuid': 'c65a389c-86a3-4e4a-ab1c-d5420c70490c'}

In [8]: app.content_sources.rest_client.repositories_api.search_uploads({'sha256':(["d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866"])})
2024-12-03 14:38:37.551 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bz7ty4v9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/search with query params [] and x-rh-insights-request-id=<>
Out[8]: 
{'found': {'d3c9d09bba6bb097aeca402aacfea8e4765dbb8186be3d5099c54fba6d3e1866': '/api/pulp/cs-1658103b69/api/v3/artifacts/01938b3d-b98b-7454-b490-6c12809d33e8/'},
 'missing': []}

In [9]: 

```


### Comment by @dominikvagner on December 04, 2024 at 03:04 PM UTC

Closing this as we decided to this check for already uploaded files in a different way. 

@Andrewgdewar you can have a look at the changes in this PR to see if you want to reuse some parts of this ([this](https://github.com/content-services/content-sources-backend/pull/909/files#diff-1f86b548b27cf93ec91e4767056bd49ca63cad83c6c644dfb9298869eb407403R162) could be reused) 

---

## Reviews

### Review by @jlsherrill - Commented on November 26, 2024 at 12:45 PM UTC

### Review by @jlsherrill - Commented on November 26, 2024 at 12:47 PM UTC

### Review by @dominikvagner - Commented on November 28, 2024 at 10:48 AM UTC

### Review by @dominikvagner - Commented on November 28, 2024 at 10:49 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/909*
