---
type: pull_request
number: 924
title: "HMS-5064: fix template list filtering by uuids"
state: merged
author: dominikvagner
created: 2024-12-17T14:20:38Z
updated: 2024-12-19T13:47:53Z
closed: 2024-12-19T13:47:53Z
merged: 2024-12-19T13:47:53Z
base_branch: main
head_branch: 5064
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/924
---

# Pull Request #924: HMS-5064: fix template list filtering by uuids

**Author**: @dominikvagner
**Created**: December 17, 2024 at 02:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5064`

## Description

## Summary
This PR fixes an bug with filtering listed templates by both `repository_uuids` and `snapshot_uuids` (the `templates_repository_configurations` was being joined twice).

## Testing steps
- Create repos, let them snapshot and use them in templates.
- List templates and filter them by `repository_uuids` and `snapshot_uuids`.
- Verify that the correct templates are shown (if both of these filters are used together they are in a `OR`)



---

## Discussion

### Comment by @jlsherrill on December 17, 2024 at 02:33 PM UTC

https://issues.redhat.com/browse/HMS-5064

### Comment by @mayurilahane on December 18, 2024 at 11:24 AM UTC

/retest

### Comment by @jlsherrill on December 18, 2024 at 06:12 PM UTC

https://issues.redhat.com/browse/Fixes 5064

### Comment by @mayurilahane on December 19, 2024 at 01:12 PM UTC

template filtering with repository_uuids

```In [17]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="a3d00123-9030-4d88-b3a9-be4626d923a7")
2024-12-19 18:39:14.530 [    INFO] [iqe.base.rest_client] REST: GET https://ee-75ivd2q7.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [('repository_uuids', 'a3d00123-9030-4d88-b3a9-be4626d923a7')] and x-rh-insights-request-id=51eb0eac-8457-40c1-b0b7-fb48b13774a7
Out[17]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-12-19T13:06:40.677875Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_update_snapshot_error': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'demo2_temp',
           'org_id': '3340851',
           'repository_uuids': ['a3d00123-9030-4d88-b3a9-be4626d923a7',
                                '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b'],
           'rhsm_environment_created': False,
           'rhsm_environment_id': '84280beaf5da4305991e7bd835a94f93',
           'snapshots': [{'added_counts': {'rpm.advisory': 2, 'rpm.package': 8},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8},
                          'created_at': '2024-12-19T13:02:35.509361Z',
                          'removed_counts': {},
                          'repository_name': 'demo1',
                          'repository_path': 'cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d',
                          'repository_uuid': 'a3d00123-9030-4d88-b3a9-be4626d923a7',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d/',
                          'uuid': '301df057-829e-4e8e-8a5b-8f9e0e7073da'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:02:47.75184Z',
                          'removed_counts': {},
                          'repository_name': 'demo2',
                          'repository_path': 'cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1',
                          'repository_uuid': '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1/',
                          'uuid': '172ad302-fb92-46b9-9b39-a490755094e4'}],
           'updated_at': '2024-12-19T13:06:40.677875Z',
           'use_latest': True,
           'uuid': '84280bea-f5da-4305-991e-7bd835a94f93',
           'version': '8'},
          {'arch': 'x86_64',
           'created_at': '2024-12-19T13:06:55.135078Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_update_snapshot_error': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'demo3_temp',
           'org_id': '3340851',
           'repository_uuids': ['a3d00123-9030-4d88-b3a9-be4626d923a7'],
           'rhsm_environment_created': False,
           'rhsm_environment_id': '72fc01ae68494d5f8322295a889a541d',
           'snapshots': [{'added_counts': {'rpm.advisory': 2, 'rpm.package': 8},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8},
                          'created_at': '2024-12-19T13:02:35.509361Z',
                          'removed_counts': {},
                          'repository_name': 'demo1',
                          'repository_path': 'cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d',
                          'repository_uuid': 'a3d00123-9030-4d88-b3a9-be4626d923a7',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d/',
                          'uuid': '301df057-829e-4e8e-8a5b-8f9e0e7073da'}],
           'updated_at': '2024-12-19T13:06:55.135078Z',
           'use_latest': True,
           'uuid': '72fc01ae-6849-4d5f-8322-295a889a541d',
           'version': '8'},
          {'arch': 'x86_64',
           'created_at': '2024-12-19T13:06:06.819933Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_update_snapshot_error': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'demo_temp',
           'org_id': '3340851',
           'repository_uuids': ['a3d00123-9030-4d88-b3a9-be4626d923a7',
                                '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                                'fa1d26d3-6832-495e-bdc4-14f5d4fe6068'],
           'rhsm_environment_created': False,
           'rhsm_environment_id': '2a584f164f21467cb6b89cd1cb05d577',
           'snapshots': [{'added_counts': {'rpm.advisory': 2, 'rpm.package': 8},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8},
                          'created_at': '2024-12-19T13:02:35.509361Z',
                          'removed_counts': {},
                          'repository_name': 'demo1',
                          'repository_path': 'cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d',
                          'repository_uuid': 'a3d00123-9030-4d88-b3a9-be4626d923a7',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d/',
                          'uuid': '301df057-829e-4e8e-8a5b-8f9e0e7073da'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:02:47.75184Z',
                          'removed_counts': {},
                          'repository_name': 'demo2',
                          'repository_path': 'cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1',
                          'repository_uuid': '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1/',
                          'uuid': '172ad302-fb92-46b9-9b39-a490755094e4'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.distribution_tree': 1,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.distribution_tree': 1,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:03:53.635076Z',
                          'removed_counts': {},
                          'repository_name': 'demo3',
                          'repository_path': 'cs-404f0bfd89/fa1d26d3-6832-495e-bdc4-14f5d4fe6068/42644334-db3d-498a-b318-883b9b494fdc',
                          'repository_uuid': 'fa1d26d3-6832-495e-bdc4-14f5d4fe6068',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/fa1d26d3-6832-495e-bdc4-14f5d4fe6068/42644334-db3d-498a-b318-883b9b494fdc/',
                          'uuid': '60ecaf6e-1dde-401c-96fd-6f1523c2e145'}],
           'updated_at': '2024-12-19T13:06:06.819933Z',
           'use_latest': True,
           'uuid': '2a584f16-4f21-467c-b6b8-9cd1cb05d577',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=a3d00123-9030-4d88-b3a9-be4626d923a7',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids=a3d00123-9030-4d88-b3a9-be4626d923a7'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}
```

### Comment by @mayurilahane on December 19, 2024 at 01:25 PM UTC


Filtering with snapshot_uuids 
```

In [24]: app.content_sources.rest_client.templates_api.list_templates(snapshot_uuids="172ad302-fb92-46b9-9b39-a490755094e4")
2024-12-19 18:51:36.616 [    INFO] [iqe.base.rest_client] REST: GET https://ee-75ivd2q7.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [('snapshot_uuids', '172ad302-fb92-46b9-9b39-a490755094e4')] and x-rh-insights-request-id=c052e39b-4367-4e37-8a1f-dc1d2daee825
Out[24]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-12-19T13:06:40.677875Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_update_snapshot_error': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'demo2_temp',
           'org_id': '3340851',
           'repository_uuids': ['a3d00123-9030-4d88-b3a9-be4626d923a7',
                                '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b'],
           'rhsm_environment_created': False,
           'rhsm_environment_id': '84280beaf5da4305991e7bd835a94f93',
           'snapshots': [{'added_counts': {'rpm.advisory': 2, 'rpm.package': 8},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8},
                          'created_at': '2024-12-19T13:02:35.509361Z',
                          'removed_counts': {},
                          'repository_name': 'demo1',
                          'repository_path': 'cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d',
                          'repository_uuid': 'a3d00123-9030-4d88-b3a9-be4626d923a7',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d/',
                          'uuid': '301df057-829e-4e8e-8a5b-8f9e0e7073da'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:02:47.75184Z',
                          'removed_counts': {},
                          'repository_name': 'demo2',
                          'repository_path': 'cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1',
                          'repository_uuid': '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1/',
                          'uuid': '172ad302-fb92-46b9-9b39-a490755094e4'}],
           'updated_at': '2024-12-19T13:06:40.677875Z',
           'use_latest': True,
           'uuid': '84280bea-f5da-4305-991e-7bd835a94f93',
           'version': '8'},
          {'arch': 'x86_64',
           'created_at': '2024-12-19T13:06:06.819933Z',
           'created_by': 'ephemeral-user',
           'date': '0001-01-01T00:00:00Z',
           'description': '',
           'last_update_snapshot_error': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'demo_temp',
           'org_id': '3340851',
           'repository_uuids': ['a3d00123-9030-4d88-b3a9-be4626d923a7',
                                '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                                'fa1d26d3-6832-495e-bdc4-14f5d4fe6068'],
           'rhsm_environment_created': False,
           'rhsm_environment_id': '2a584f164f21467cb6b89cd1cb05d577',
           'snapshots': [{'added_counts': {'rpm.advisory': 2, 'rpm.package': 8},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8},
                          'created_at': '2024-12-19T13:02:35.509361Z',
                          'removed_counts': {},
                          'repository_name': 'demo1',
                          'repository_path': 'cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d',
                          'repository_uuid': 'a3d00123-9030-4d88-b3a9-be4626d923a7',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/a3d00123-9030-4d88-b3a9-be4626d923a7/cc2553fb-77dc-4181-9a9c-4367288ccf3d/',
                          'uuid': '301df057-829e-4e8e-8a5b-8f9e0e7073da'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:02:47.75184Z',
                          'removed_counts': {},
                          'repository_name': 'demo2',
                          'repository_path': 'cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1',
                          'repository_uuid': '9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/9e0f6e0e-a01b-4f5a-8da0-2e77b8a4008b/50c2d117-6f73-4873-a5f8-0980709bf7d1/',
                          'uuid': '172ad302-fb92-46b9-9b39-a490755094e4'},
                         {'added_counts': {'rpm.advisory': 2,
                                           'rpm.distribution_tree': 1,
                                           'rpm.package': 8,
                                           'rpm.packagecategory': 1,
                                           'rpm.packagegroup': 2},
                          'content_counts': {'rpm.advisory': 2,
                                             'rpm.distribution_tree': 1,
                                             'rpm.package': 8,
                                             'rpm.packagecategory': 1,
                                             'rpm.packagegroup': 2},
                          'created_at': '2024-12-19T13:03:53.635076Z',
                          'removed_counts': {},
                          'repository_name': 'demo3',
                          'repository_path': 'cs-404f0bfd89/fa1d26d3-6832-495e-bdc4-14f5d4fe6068/42644334-db3d-498a-b318-883b9b494fdc',
                          'repository_uuid': 'fa1d26d3-6832-495e-bdc4-14f5d4fe6068',
                          'url': 'http://pulp-content:8000/api/pulp-content/cs-404f0bfd89/fa1d26d3-6832-495e-bdc4-14f5d4fe6068/42644334-db3d-498a-b318-883b9b494fdc/',
                          'uuid': '60ecaf6e-1dde-401c-96fd-6f1523c2e145'}],
           'updated_at': '2024-12-19T13:06:06.819933Z',
           'use_latest': True,
           'uuid': '2a584f16-4f21-467c-b6b8-9cd1cb05d577',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&snapshot_uuids=172ad302-fb92-46b9-9b39-a490755094e4',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&snapshot_uuids=172ad302-fb92-46b9-9b39-a490755094e4'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```


### Comment by @mayurilahane on December 19, 2024 at 01:28 PM UTC

Hey @dominikvagner 
is this call supposed to work?

```
In [26]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="a3d00123-9030-4d88-b3a9-be4626d923a7", snapshot_uuids="172ad302-fb92-46b9-9b3
    ...: 9-a490755094e4")
```


### Comment by @dominikvagner on December 19, 2024 at 01:32 PM UTC

> Hey @dominikvagner is this call supposed to work?
> 
> ```
> In [26]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="a3d00123-9030-4d88-b3a9-be4626d923a7", snapshot_uuids="172ad302-fb92-46b9-9b3
>     ...: 9-a490755094e4")
> ```

@mayurilahane yes it is 😅 

### Comment by @mayurilahane on December 19, 2024 at 01:41 PM UTC

> > Hey @dominikvagner is this call supposed to work?
> > ```
> > In [26]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="a3d00123-9030-4d88-b3a9-be4626d923a7", snapshot_uuids="172ad302-fb92-46b9-9b3
> >     ...: 9-a490755094e4")
> > ```
> 
> @mayurilahane yes it is 😅

it's working i just wanted to confirm it.
Thanks !

### Comment by @mayurilahane on December 19, 2024 at 01:41 PM UTC

/lgtm merging it 

---

## Reviews

### Review by @xbhouse - Approved on December 17, 2024 at 06:40 PM UTC

looks good! TestListBySnapshot was panicking locally for me, this fixes that too :D 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/924*
