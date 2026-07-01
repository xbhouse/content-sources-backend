---
type: pull_request
number: 800
title: "Fixes 4445: add apis to bulk import/export custom repos"
state: merged
author: xbhouse
created: 2024-09-05T20:24:08Z
updated: 2024-09-24T19:30:27Z
closed: 2024-09-24T19:24:41Z
merged: 2024-09-24T19:24:41Z
base_branch: main
head_branch: 4445
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/800
---

# Pull Request #800: Fixes 4445: add apis to bulk import/export custom repos

**Author**: @xbhouse
**Created**: September 05, 2024 at 08:24 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4445`

## Description

## Summary

Adds bulk import and export apis for custom repos

## Testing steps

1. Add a couple custom repos and let them snapshot
2. Make a request to export the repos and send the repository uuids in the body: `POST /repositories/bulk_export/`
3. Use the response you get from the bulk export (try changing some of the fields) to make a request to import the repos into the same org (`POST /repositories/bulk_import/`). You should see warnings in the response that the repo urls already exist and any mismatches between fields of existing and new repos
4. Use the same response to import the repos into a different org that doesn't have those repos. The repos should be created and introspection and snapshot tasks should be triggered

---

## Discussion

### Comment by @jlsherrill on September 05, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-4445

### Comment by @xbhouse on September 09, 2024 at 03:50 PM UTC

cc @avitova

### Comment by @xbhouse on September 10, 2024 at 07:01 PM UTC

/retest

### Comment by @avitova on September 17, 2024 at 02:00 PM UTC

Thank you, looks good:)

### Comment by @xbhouse on September 19, 2024 at 08:06 PM UTC

/retest

### Comment by @xbhouse on September 23, 2024 at 02:49 PM UTC

/retest

### Comment by @swadeley on September 24, 2024 at 09:45 AM UTC

Hi @xbhouse 

I can export but not import.

```
In [5]: app.content_sources.rest_client.repositories_api.bulk_export_repositories({"repository_uuids":['d5bddf59-d1f0-4fed-8c68-25049eb0858d','2267ea59-b2a1-45b5-9a02-155273bd3f59']})
2024-09-24 09:04:31.175 [    INFO] [iqe.base.rest_client] REST: POST https://ee-xbcrifi9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/bulk_export/ with query params [] and x-rh-insights-request-id=<>
Out[5]: 
[{'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'gpg_key': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'repo09',
  'origin': 'external',
  'snapshot': True,
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo09/'},
 {'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'gpg_key': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'repo10',
  'origin': 'external',
  'snapshot': True,
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo10/'}]
```

```
In [32]: my_repo_list = [{'distribution_arch': 'any',
    ...:   'distribution_versions': ['any'],
    ...:   'gpg_key': '',
    ...:   'metadata_verification': False,
    ...:   'module_hotfixes': False,
    ...:   'name': 'repo07',
    ...:   'origin': 'external',
    ...:   'snapshot': True,
    ...:   'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo07/'},
    ...:  {'distribution_arch': 'any',
    ...:   'distribution_versions': ['any'],
    ...:   'gpg_key': '',
    ...:   'metadata_verification': False,
    ...:   'module_hotfixes': False,
    ...:   'name': 'repo08',
    ...:   'origin': 'external',
    ...:   'snapshot': True,
    ...:   'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo08/'}]

In [33]: app.content_sources.rest_client.repositories_api.bulk_import_repositories(my_repo_list)
<snip traceback>
ApiTypeError: Invalid type for variable 'warnings'. Required value type is list and passed type was NoneType at ['received_data'][0]['warnings']
```

The same list works with bulk create:

```
In [37]: app.content_sources.rest_client.repositories_api.bulk_create_repositories(my_repo_list)
2024-09-24 10:19:56.431 [    INFO] [iqe.base.rest_client] REST: POST https://ee-xbcrifi9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/bulk_create/ with query params [] and x-rh-insights-request-id=42f0ce21-59f7-4375-900a-e63950b196e7
Out[37]: 
[{'account_id': '0369233',
  'content_type': 'rpm',
  'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'failed_introspections_count': 0,
  'gpg_key': '',
  'label': 'repo07',
  'last_introspection_error': '',
  'last_introspection_status': 'Valid',
  'last_introspection_time': '2024-09-24T09:17:54Z',
  'last_snapshot_task_uuid': '6a5dff64-73e3-4968-89a5-b6288fc857b4',
  'last_success_introspection_time': '2024-09-24T09:17:54Z',
  'last_update_introspection_time': '2024-09-24T09:15:46Z',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'repo07',
  'org_id': '3340851',
  'origin': 'external',
  'package_count': 1,
  'snapshot': True,
  'status': 'Valid',
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo07/',
  'uuid': '2b6c6218-759a-4267-bc55-2effa37cd14f'},
 {'account_id': '0369233',
  'content_type': 'rpm',
  'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'failed_introspections_count': 0,
  'gpg_key': '',
  'label': 'repo08',
  'last_introspection_error': '',
  'last_introspection_status': 'Valid',
  'last_introspection_time': '2024-09-24T09:17:54Z',
  'last_snapshot_task_uuid': '39d0fef6-9267-45f0-a835-3b70c80d7d1d',
  'last_success_introspection_time': '2024-09-24T09:17:54Z',
  'last_update_introspection_time': '2024-09-24T09:15:46Z',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'repo08',
  'org_id': '3340851',
  'origin': 'external',
  'package_count': 1,
  'snapshot': True,
  'status': 'Valid',
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo08/',
  'uuid': 'af982bdd-f738-4948-b56b-7d34a23ca2f6'}]
```


### Comment by @xbhouse on September 24, 2024 at 01:20 PM UTC

hi @swadeley, pushed a fix to change the warnings to an empty list if they are null. seems to be working now with the new image:

```
In [8]: app.content_sources.rest_client.repositories_api.bulk_export_repositories({'repository_uuids': ['f885b2f3-5e9e-4779-9b22-dc66ab4feedf', '34e7cf0f-ab1f-4d16-9864-bc706c6bb4a1']})
2024-09-24 09:15:01.529 [    INFO] [iqe.base.rest_client] REST: POST https://ee-re4oc2ae.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/bulk_export/ with query params [] and x-rh-insights-request-id=0984e2b2-a53d-493c-828a-514d96bd2f89

Out[8]: 
[{'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'gpg_key': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'test1',
  'origin': 'external',
  'snapshot': True,
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo09/'},
 {'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'gpg_key': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'test2',
  'origin': 'external',
  'snapshot': True,
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo10/'}]

In [9]: repo_list = [{'distribution_arch': 'any',
   ...:   'distribution_versions': ['any'],
   ...:   'gpg_key': '',
   ...:   'metadata_verification': False,
   ...:   'module_hotfixes': False,
   ...:   'name': 'test4',
   ...:   'origin': 'external',
   ...:   'snapshot': True,
   ...:   'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo07/'},
   ...:  {'distribution_arch': 'any',
   ...:   'distribution_versions': ['any'],
   ...:   'gpg_key': '',
   ...:   'metadata_verification': False,
   ...:   'module_hotfixes': False,
   ...:   'name': 'test3',
   ...:   'origin': 'external',
   ...:   'snapshot': True,
   ...:   'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo08/'}]
   
In [10]: app.content_sources.rest_client.repositories_api.bulk_import_repositories(repo_list)
2024-09-24 09:16:29.845 [    INFO] [iqe.base.rest_client] REST: POST https://ee-re4oc2ae.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/bulk_import/ with query params [] and x-rh-insights-request-id=9e3b5c7c-4c8d-4d58-9945-289ce76cebe3

Out[10]: 
[{'account_id': '0369233',
  'content_type': 'rpm',
  'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'failed_introspections_count': 0,
  'gpg_key': '',
  'label': 'test4',
  'last_introspection_error': '',
  'last_introspection_status': 'Pending',
  'last_introspection_time': '',
  'last_success_introspection_time': '',
  'last_update_introspection_time': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'test4',
  'org_id': '3340851',
  'origin': 'external',
  'package_count': 0,
  'snapshot': True,
  'status': 'Pending',
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo07/',
  'uuid': '2539b981-d9fe-475b-b3b2-c3d5ebf35e50',
  'warnings': []},
 {'account_id': '0369233',
  'content_type': 'rpm',
  'distribution_arch': 'any',
  'distribution_versions': ['any'],
  'failed_introspections_count': 0,
  'gpg_key': '',
  'label': 'test3',
  'last_introspection_error': '',
  'last_introspection_status': 'Pending',
  'last_introspection_time': '',
  'last_success_introspection_time': '',
  'last_update_introspection_time': '',
  'metadata_verification': False,
  'module_hotfixes': False,
  'name': 'test3',
  'org_id': '3340851',
  'origin': 'external',
  'package_count': 0,
  'snapshot': True,
  'status': 'Pending',
  'url': 'https://stephenw.fedorapeople.org/multirepos/10/repo08/',
  'uuid': 'ee2c5508-6586-4545-8658-6c5560fcf015',
  'warnings': []}]
```

### Comment by @swadeley on September 24, 2024 at 07:24 PM UTC

Ok, API itself works as expected. There is a known  issue in apigen but that only effects error handling when attempting to create duplicate repos in tests.

---

## Reviews

### Review by @avitova - Commented on September 11, 2024 at 10:25 AM UTC

Thank you again for looking into this! :heart_hands: This is a lot of work, I hope I did not miss anything, but overall it seems good, and this will be very useful for us!
I left a few comments with concerns to think about.

### Review by @xbhouse - Commented on September 11, 2024 at 06:45 PM UTC

### Review by @xbhouse - Commented on September 11, 2024 at 07:10 PM UTC

### Review by @xbhouse - Commented on September 11, 2024 at 10:23 PM UTC

### Review by @avitova - Commented on September 12, 2024 at 09:16 AM UTC

### Review by @rverdile - Approved on September 19, 2024 at 05:04 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/800*
