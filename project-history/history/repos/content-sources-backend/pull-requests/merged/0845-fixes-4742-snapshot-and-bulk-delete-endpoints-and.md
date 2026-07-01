---
type: pull_request
number: 845
title: "Fixes 4742: snapshot and bulk delete endpoints and task"
state: merged
author: dominikvagner
created: 2024-10-11T12:31:42Z
updated: 2024-11-18T17:02:28Z
closed: 2024-11-18T16:59:13Z
merged: 2024-11-18T16:59:13Z
base_branch: main
head_branch: 4742
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/845
---

# Pull Request #845: Fixes 4742: snapshot and bulk delete endpoints and task

**Author**: @dominikvagner
**Created**: October 11, 2024 at 12:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4742`

## Description

## Summary
- Two endpoints for manual snapshot deletion added, one for deleting a single snapshot and a bulk delete one.
- Endpoints only soft-delete snapshots and then start a task cleaning those up fully along with Pulp content and updating templates that are using those.

## Testing steps
1. Create a repository, with snapshots enabled.
2. Add the repository to a template and make sure it uses the snapshots that's gonna be deleted.
3. After the repository has a snapshot created, change the URL and trigger a new snapshot.
4. Delete a snapshot from that repo with the new endpoint.
`http DELETE :8000/api/content-sources/v1.0/repositories/${REPO_UUID}/snapshots/${SNAP_UUID}`
5. Verify that the snapshot has been deleted (i.e. not present when listing snapshots for a repo).
6. Verify that the template snapshot has been updated to the other one. (Doesn't work yet.)
- For bulk delete, do the same with more snapshots.

---

## Discussion

### Comment by @jlsherrill on October 11, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4742

### Comment by @swadeley on November 13, 2024 at 09:49 AM UTC

Hi @dominikvagner , please rebase.

### Comment by @swadeley on November 14, 2024 at 11:31 AM UTC

Hi, I created a template and added a repo with one snapshot:

```
In [1]: repo_dict = dict(
   ...:     name='test',
   ...:     url='https://stephenw.fedorapeople.org/fakerepos/multiple_errata',
   ...:     snapshot=True,
   ...: )

In [2]: app.content_sources.rest_client.repositories_api.create_repository(repo_dict)
2024-11-14 10:59:19.574 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-11-14 10:59:19.574 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-11-14 10:59:20.147 [    INFO] [root] Created RESTPluginService client for https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 with the following attributes: ['client', 'environments_api', 'features_api', 'gpg_key_api', 'package', 'packagegroups_api', 'popular_repositories_api', 'public_repositories_api', 'repositories_api', 'rpms_api', 'snapshots_api', 'tasks_api', 'templates_api']
2024-11-14 10:59:20.548 [    INFO] [iqe.base.rest_client] REST: POST https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [] and x-rh-insights-request-id=<>
Out[2]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-11-14T10:54:05Z',
 'last_snapshot_task_uuid': '481f372f-751a-4157-81e7-0615f40d478f',
 'last_success_introspection_time': '2024-11-14T10:54:05Z',
 'last_update_introspection_time': '2024-11-14T10:54:05Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 27,
 'snapshot': True,
 'status': 'Valid',
 'url': 'https://stephenw.fedorapeople.org/fakerepos/multiple_errata/',
 'uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764'}
```

Then I changed repo URL and created second snapshot:

```
In [3]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('896965b1-0b2f-4e7a-b30e-a051b89dc764')
2024-11-14 11:10:16.256 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[3]: 
{'data': [{'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'},
          {'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'created_at': '2024-11-14T10:59:35.206916Z',
           'removed_counts': {},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/56eb81d8-5ed9-4e32-9a47-9d66af262592',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/56eb81d8-5ed9-4e32-9a47-9d66af262592/',
           'uuid': '3fa45d12-170f-4972-93d7-38d40244322e'}],
 'links': {'first': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```
Notice first snapshot was created at 10:59 with rpm.package': 27 and second snapshot was created at 11:06 with 'rpm.package': 34.

Deleted older snapshot:

```
In [4]: app.content_sources.rest_client.snapshots_api.delete_snapshot.params_map
Out[4]: 
{'all': ['repo_uuid',
  'snapshot_uuid',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['repo_uuid', 'snapshot_uuid'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}

In [5]: app.content_sources.rest_client.snapshots_api.delete_snapshot('896965b1-0b2f-4e7a-b30e-a051b89dc764','3fa45d12-170f-4972-93d7-38d40244322e')
2024-11-14 11:12:24.216 [    INFO] [iqe.base.rest_client] REST: DELETE https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/3fa45d12-170f-4972-93d7-38d40244322e with query params [] and x-rh-insights-request-id=<>

In [6]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('896965b1-0b2f-4e7a-b30e-a051b89dc764')
2024-11-14 11:12:54.683 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[6]: 
{'data': [{'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'}],
 'links': {'first': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [7]: app.content_sources.rest_client.templates_api.list_templates()['data'][0]['uuid']
2024-11-14 11:14:05.184 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=-<>-
Out[7]: '1a98c6d8-d929-4716-a208-4e89805e24f3'

In [8]: app.content_sources.rest_client.templates_api.get_template('1a98c6d8-d929-4716-a208-4e89805e24f3')
2024-11-14 11:14:25.005 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3 with query params [] and x-rh-insights-request-id=<>
Out[8]: 
{'arch': 'x86_64',
 'created_at': '2024-11-14T10:49:43.676697Z',
 'created_by': 'ephemeral-user',
 'date': '2024-11-14T10:49:44Z',
 'description': 'Edited the template',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'Test-template-edited-kKDA',
 'org_id': '3340851',
 'repository_uuids': ['0505cf94-379b-402f-8150-fd1b6a4f6425',
                      '896965b1-0b2f-4e7a-b30e-a051b89dc764'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': '1a98c6d8d9294716a2084e89805e24f3',
 'updated_at': '2024-11-14T11:05:04.068273Z',
 'use_latest': False,
 'uuid': '1a98c6d8-d929-4716-a208-4e89805e24f3',
 'version': '8'}

In [9]: 

In [9]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_template('1a98c6d8-d929-4716-a208-4e89805e24f3')
2024-11-14 11:26:05.366 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[9]: 
{'data': [{'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-11-14T10:07:38.999405Z',
           'removed_counts': {},
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_path': 'cs-3d34346b07/0505cf94-379b-402f-8150-fd1b6a4f6425/aa955816-273e-406c-8d27-6abd2fbeb9af',
           'repository_uuid': '0505cf94-379b-402f-8150-fd1b6a4f6425',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-3d34346b07/0505cf94-379b-402f-8150-fd1b6a4f6425/aa955816-273e-406c-8d27-6abd2fbeb9af/',
           'uuid': 'f92345d6-6efa-4a67-bdcc-e2f90775e04d'},
          {'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'}],
 'links': {'first': '/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

```

We can see template has newer snapshot created at 11:06 with  'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},

### Comment by @swadeley on November 14, 2024 at 12:56 PM UTC

Hi, my repo now has three snapshots:

```
In [11]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('896965b1-0b2f-4e7a-b30e-a051b89dc764')
2024-11-14 11:47:01.878 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[11]: 
{'data': [{'added_counts': {'rpm.advisory': 5, 'rpm.package': 23},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 23},
           'created_at': '2024-11-14T11:44:18.526663Z',
           'removed_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/11e6c076-653f-43f2-b4cf-79519d2db147',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/11e6c076-653f-43f2-b4cf-79519d2db147/',
           'uuid': '07d64a2e-875d-4690-92f1-0b7c004a70db'},
          {'added_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'content_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'created_at': '2024-11-14T11:43:56.250167Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/54222081-9160-4759-844f-c78930a176b4',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/54222081-9160-4759-844f-c78930a176b4/',
           'uuid': 'd1bf61de-03ec-4588-99d6-dfe9e89a774c'},
          {'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'}],
 'links': {'first': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}
```

Template is using the one created at 11:06:

```
In [12]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_template('1a98c6d8-d929-4716-a208-4e89805e24f3')
2024-11-14 11:48:16.333 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/ with query params [] and x-rh-insights-request-id=ea6d7be4-7637-4326-850e-da5837e42585
Out[12]: 
{'data': [{'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-11-14T10:07:38.999405Z',
           'removed_counts': {},
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_path': 'cs-3d34346b07/0505cf94-379b-402f-8150-fd1b6a4f6425/aa955816-273e-406c-8d27-6abd2fbeb9af',
           'repository_uuid': '0505cf94-379b-402f-8150-fd1b6a4f6425',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-3d34346b07/0505cf94-379b-402f-8150-fd1b6a4f6425/aa955816-273e-406c-8d27-6abd2fbeb9af/',
           'uuid': 'f92345d6-6efa-4a67-bdcc-e2f90775e04d'},
          {'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'}],
 'links': {'first': '/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/1a98c6d8-d929-4716-a208-4e89805e24f3/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```


But bulk delete does not work as I expected:

```
In [45]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('896965b1-0b2f-4e7a-b30e-a051b89dc764')
2024-11-14 12:52:15.984 [    INFO] [iqe.base.rest_client] REST: GET https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/ with query params []
Out[45]: 
{'data': [{'added_counts': {'rpm.advisory': 5, 'rpm.package': 23},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 23},
           'created_at': '2024-11-14T11:44:18.526663Z',
           'removed_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/11e6c076-653f-43f2-b4cf-79519d2db147',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/11e6c076-653f-43f2-b4cf-79519d2db147/',
           'uuid': '07d64a2e-875d-4690-92f1-0b7c004a70db'},
          {'added_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'content_counts': {'rpm.advisory': 24, 'rpm.package': 66},
           'created_at': '2024-11-14T11:43:56.250167Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/54222081-9160-4759-844f-c78930a176b4',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/54222081-9160-4759-844f-c78930a176b4/',
           'uuid': 'd1bf61de-03ec-4588-99d6-dfe9e89a774c'},
          {'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-11-14T11:06:34.033667Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'test',
           'repository_path': 'cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da',
           'repository_uuid': '896965b1-0b2f-4e7a-b30e-a051b89dc764',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-965c583b4c/896965b1-0b2f-4e7a-b30e-a051b89dc764/6b682dae-8ce4-4e20-b102-cd076026e8da/',
           'uuid': '28c72a4a-27f2-4a28-a1b6-27f880a0c239'}],
 'links': {'first': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/?limit=100&offset=0'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [46]: app.content_sources.rest_client.snapshots_api.bulk_delete_snapshots('896965b1-0b2f-4e7a-b30e-a051b89dc764',{'uuids':['28c72a4a-27f2-4a28-a1b6-27f880a0c239','d1bf61de-03ec-4588-99d6-dfe9e89a774c']})
2024-11-14 12:53:41.363 [    INFO] [iqe.base.rest_client] REST: DELETE https://ee-obgrrhxq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/896965b1-0b2f-4e7a-b30e-a051b89dc764/snapshots/bulk_delete/ with query params []
<snip>
HTTP response body: {"errors":[{"status":404,"title":"Error deleting snapshots","detail":"snapshot with this UUID does exist for the specified repository"}]}
```


### Comment by @dominikvagner on November 14, 2024 at 03:29 PM UTC

ahh, I put a wrong method in the godoc for the bulk delete handler so the openapi spec did generate incorrectly, it should be fixed now 🫣 😅  
@swadeley 

### Comment by @swadeley on November 15, 2024 at 08:46 PM UTC

/retest

### Comment by @swadeley on November 18, 2024 at 08:54 AM UTC

/retest

### Comment by @swadeley on November 18, 2024 at 04:55 PM UTC

Hi

It is working now as described:

```
In [1]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external')
2024-11-18 16:44:11.366 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-11-18 16:44:11.367 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-11-18 16:44:11.721 [    INFO] [root] Created RESTPluginService client for https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 with the following attributes: ['client', 'environments_api', 'features_api', 'gpg_key_api', 'package', 'packagegroups_api', 'popular_repositories_api', 'public_repositories_api', 'repositories_api', 'rpms_api', 'snapshots_api', 'tasks_api', 'templates_api']
2024-11-18 16:44:12.034 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/ with query params [('origin', 'external')] and x-rh-insights-request-id=<>
Out[1]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test_repo',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-11-18T16:37:37Z',
           'last_snapshot': {'added_counts': {'rpm.package': 35},
                             'content_counts': {'rpm.advisory': 4,
                                                'rpm.package': 35,
                                                'rpm.packagecategory': 1,
                                                'rpm.packagegroup': 2,
                                                'rpm.packagelangpacks': 1},
                             'created_at': '2024-11-18T16:37:50.204777Z',
                             'removed_counts': {'rpm.package': 35},
                             'repository_name': 'test-repo',
                             'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d',
                             'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
                             'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d/',
                             'uuid': '9def725e-964d-46b7-817b-143fd84189d8'},
           'last_snapshot_task': {'created_at': '2024-11-18T16:37:37Z',
                                  'ended_at': '2024-11-18T16:37:50Z',
                                  'error': '',
                                  'object_name': 'test-repo',
                                  'object_type': 'repository',
                                  'object_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
                                  'org_id': '3340851',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': 'a00ef41f-35c1-4c17-aca5-545e7ef5008f'},
           'last_snapshot_task_uuid': 'a00ef41f-35c1-4c17-aca5-545e7ef5008f',
           'last_snapshot_uuid': '9def725e-964d-46b7-817b-143fd84189d8',
           'last_success_introspection_time': '2024-11-18T16:37:37Z',
           'last_update_introspection_time': '2024-11-18T16:37:37Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test-repo',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 35,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://fixtures.pulpproject.org/rpm-with-sha-512/',
           'uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [2]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('9c13d9fd-0f34-473f-a106-cd326ae3b8be')
2024-11-18 16:44:41.528 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[2]: 
{'data': [{'added_counts': {'rpm.package': 35},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:37:50.204777Z',
           'removed_counts': {'rpm.package': 35},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d/',
           'uuid': '9def725e-964d-46b7-817b-143fd84189d8'},
          {'added_counts': {'rpm.package': 35},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:36:51.112913Z',
           'removed_counts': {'rpm.package': 35},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/8c41b8d8-1805-4c9d-b351-c126f12ee973',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/8c41b8d8-1805-4c9d-b351-c126f12ee973/',
           'uuid': 'e22b18c4-7bf7-412a-97f8-d7889b1a8441'},
          {'added_counts': {'rpm.advisory': 4,
                            'rpm.package': 35,
                            'rpm.packagecategory': 1,
                            'rpm.packagegroup': 2,
                            'rpm.packagelangpacks': 1},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:31:31.317636Z',
           'removed_counts': {},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/72feb2b5-c9ef-4e4e-8761-ab8cf6f88df2',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/72feb2b5-c9ef-4e4e-8761-ab8cf6f88df2/',
           'uuid': '5ca7a2e9-4329-4f8b-93be-9e72bfa42f40'}],
 'links': {'first': '/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/?limit=100&offset=0'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [3]: app.content_sources.rest_client.templates_api.list_templates()['data'][0]['uuid']
2024-11-18 16:48:08.145 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=<>
Out[3]: 'e3a93c75-91b9-46ff-afa2-6949379bdd31'

In [4]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_template('e3a93c75-91b9-46ff-afa2-6949379bdd31')
2024-11-18 16:48:34.275 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[4]: 
{'data': [{'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-11-18T16:15:50.326283Z',
           'removed_counts': {},
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_path': 'cs-c3038dd201/79fd1033-ae13-4947-9e0b-e6eeb318284b/da80af25-91e4-4b12-b6f2-14189d4919e0',
           'repository_uuid': '79fd1033-ae13-4947-9e0b-e6eeb318284b',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-c3038dd201/79fd1033-ae13-4947-9e0b-e6eeb318284b/da80af25-91e4-4b12-b6f2-14189d4919e0/',
           'uuid': 'c0cf7aca-d48c-490a-9957-bc46aed4f1ba'},
          {'added_counts': {'rpm.advisory': 4,
                            'rpm.package': 35,
                            'rpm.packagecategory': 1,
                            'rpm.packagegroup': 2,
                            'rpm.packagelangpacks': 1},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:31:31.317636Z',
           'removed_counts': {},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/72feb2b5-c9ef-4e4e-8761-ab8cf6f88df2',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/72feb2b5-c9ef-4e4e-8761-ab8cf6f88df2/',
           'uuid': '5ca7a2e9-4329-4f8b-93be-9e72bfa42f40'}],
 'links': {'first': '/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [5]: app.content_sources.rest_client.snapshots_api.bulk_delete_snapshots('9c13d9fd-0f34-473f-a106-cd326ae3b8be',{'uuids':['5ca7a2e9-4329-4f8b-93be-9e72bfa42f40','e22b18c4-7bf7-412a-97f8-d7889b1a8441']})
2024-11-18 16:51:33.532 [    INFO] [iqe.base.rest_client] REST: POST https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/bulk_delete/ with query params [] and x-rh-insights-request-id=<>

In [6]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_repo('9c13d9fd-0f34-473f-a106-cd326ae3b8be')
2024-11-18 16:51:44.909 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[6]: 
{'data': [{'added_counts': {'rpm.package': 35},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:37:50.204777Z',
           'removed_counts': {'rpm.package': 35},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d/',
           'uuid': '9def725e-964d-46b7-817b-143fd84189d8'}],
 'links': {'first': '/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/9c13d9fd-0f34-473f-a106-cd326ae3b8be/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [7]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_template('e3a93c75-91b9-46ff-afa2-6949379bdd31')
2024-11-18 16:52:30.071 [    INFO] [iqe.base.rest_client] REST: GET https://ee-gcc4jxg9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'data': [{'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-11-18T16:15:50.326283Z',
           'removed_counts': {},
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_path': 'cs-c3038dd201/79fd1033-ae13-4947-9e0b-e6eeb318284b/da80af25-91e4-4b12-b6f2-14189d4919e0',
           'repository_uuid': '79fd1033-ae13-4947-9e0b-e6eeb318284b',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-c3038dd201/79fd1033-ae13-4947-9e0b-e6eeb318284b/da80af25-91e4-4b12-b6f2-14189d4919e0/',
           'uuid': 'c0cf7aca-d48c-490a-9957-bc46aed4f1ba'},
          {'added_counts': {'rpm.package': 35},
           'content_counts': {'rpm.advisory': 4,
                              'rpm.package': 35,
                              'rpm.packagecategory': 1,
                              'rpm.packagegroup': 2,
                              'rpm.packagelangpacks': 1},
           'created_at': '2024-11-18T16:37:50.204777Z',
           'removed_counts': {'rpm.package': 35},
           'repository_name': 'test-repo',
           'repository_path': 'cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d',
           'repository_uuid': '9c13d9fd-0f34-473f-a106-cd326ae3b8be',
           'url': 'http://pulp-content:8000/api/pulp-content/cs-f9e1844bba/9c13d9fd-0f34-473f-a106-cd326ae3b8be/e35a5410-21dc-43e3-a846-0898fba5e91d/',
           'uuid': '9def725e-964d-46b7-817b-143fd84189d8'}],
 'links': {'first': '/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/e3a93c75-91b9-46ff-afa2-6949379bdd31/snapshots/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [8]: 

```

Thank you

---

## Reviews

### Review by @rverdile - Commented on October 17, 2024 at 07:55 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 03:45 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 03:45 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 05:43 PM UTC

I'm noticing an issue you can reproduce like this:

1. Create a repository, change the URL to get 2 snapshots.
2. Delete the more recent snapshot.
3. In the UI, if you try to view the snapshots you'll see "No snapshots yet". I believe because in the API, if you fetch the repository, `last_snapshot_uuid` is omitted from the response.

### Review by @rverdile - Commented on October 29, 2024 at 06:50 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 07:51 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 07:54 PM UTC

### Review by @dominikvagner - Commented on October 31, 2024 at 08:45 AM UTC

### Review by @dominikvagner - Commented on October 31, 2024 at 08:45 AM UTC

### Review by @dominikvagner - Commented on October 31, 2024 at 08:50 AM UTC

### Review by @dominikvagner - Commented on October 31, 2024 at 08:52 AM UTC

### Review by @dominikvagner - Commented on October 31, 2024 at 08:53 AM UTC

### Review by @rverdile - Commented on November 07, 2024 at 06:34 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 09:11 PM UTC

Two more small edge-cases in the handler.  Overall looking good. Templates and latest snapshot are being auto-updated. I'm going to take one more close look at the delete_snapshots task to make sure I didn't miss anything.

### Review by @rverdile - Commented on November 11, 2024 at 05:14 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 05:18 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 05:57 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 05:59 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 08:00 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 08:04 PM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 11:28 AM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 11:28 AM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 11:33 AM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 11:33 AM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 11:36 AM UTC

### Review by @rverdile - Commented on November 13, 2024 at 04:39 PM UTC

### Review by @dominikvagner - Commented on November 14, 2024 at 09:11 AM UTC

### Review by @rverdile - Approved on November 15, 2024 at 01:15 PM UTC

nice job!! :tada: 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/845*
