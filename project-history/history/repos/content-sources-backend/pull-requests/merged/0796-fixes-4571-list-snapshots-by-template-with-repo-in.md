---
type: pull_request
number: 796
title: "Fixes 4571: list snapshots by template with repo information"
state: merged
author: dominikvagner
created: 2024-09-04T09:59:00Z
updated: 2024-09-13T08:30:32Z
closed: 2024-09-13T08:11:35Z
merged: 2024-09-13T08:11:35Z
base_branch: main
head_branch: 4571
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/796
---

# Pull Request #796: Fixes 4571: list snapshots by template with repo information

**Author**: @dominikvagner
**Created**: September 04, 2024 at 09:59 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4571`

## Description

## Summary
This PR adds a new endpoint for listing snapshots by template and adds 2 new fields `repository_name` and `repository_uuid` to API snapshot response (for all endpoints).
The new added fields required adding preloading to and modifying queries in other endpoints.
New endpoint had to be added, rather than modifying the existing list one, because the process of querying and filtering the snapshots is completely different.

## Testing steps
1. Create a new template with some repositories that have snapshots.
2. Query the snapshots for the template. `/templates/{uuid}/snapshots/` and verify the shown snapshots are actually closest one for the template date.
3. Check the filtering of snapshots with the query param `repository_search` and sorting by repository name (`sort_by=repository_name:asc|desc`) or by created_at.


---

## Discussion

### Comment by @jlsherrill on September 04, 2024 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-4571

### Comment by @swadeley on September 10, 2024 at 07:57 AM UTC

/retest

### Comment by @rverdile on September 10, 2024 at 05:45 PM UTC

/retest

### Comment by @swadeley on September 12, 2024 at 11:27 AM UTC

@dominikvagner Please rebase

### Comment by @swadeley on September 12, 2024 at 03:38 PM UTC

Hi, I built API bindings (MR499) but I get:
`AttributeError: 'SnapshotsApi' object has no attribute 'list_snapshots'`
when I run API tests.
I did reinstall the plugin, no change.

EDIT:
"operationId": "listSnapshotsForRepo",

let me change that

tests are passing again

### Comment by @swadeley on September 12, 2024 at 05:56 PM UTC

/retest

### Comment by @swadeley on September 12, 2024 at 06:24 PM UTC

Hi

I created to repos, let them snapshot, change their URLs, let them snapshot again.
Added them to a test template.

I confirm in the output below that the `created_at` time stamp is from the 2nd of the two snapshots.

```
In [5]: app.content_sources.rest_client.templates_api.get_template('d2453d87-ab08-425e-9842-2615f8872161')
2024-09-12 19:15:08.981 [    INFO] [iqe.base.rest_client] REST: GET https://ee-d66czvt5.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/d2453d87-ab08-425e-9842-2615f8872161 with query params [] and x-rh-insights-request-id=<>
Out[5]: 
{'arch': 'x86_64',
 'created_at': '2024-09-12T18:05:34.797859Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test',
 'org_id': '3340851',
 'repository_uuids': ['d3d7028e-ec1d-4044-89a4-e37e75cb67b7',
                      '16367cdf-094f-49ed-bb5d-91819316eb58',
                      'a722d70c-4ac6-4627-a299-10169e05d9e4'],
 'rhsm_environment_id': 'd2453d87ab08425e98422615f8872161',
 'updated_at': '2024-09-12T18:11:52.033289Z',
 'use_latest': True,
 'uuid': 'd2453d87-ab08-425e-9842-2615f8872161',
 'version': '8'}

In [6]: app.content_sources.rest_client.snapshots_api.list_snapshots_for_template('d2453d87-ab08-425e-9842-2615f8872161')
2024-09-12 19:16:20.370 [    INFO] [iqe.base.rest_client] REST: GET https://ee-d66czvt5.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/d2453d87-ab08-425e-9842-2615f8872161/snapshots/ with query params [] and x-rh-insights-request-id=<>
Out[6]: 
{'data': [{'added_counts': {},
           'content_counts': {'rpm.package': 1},
           'created_at': '2024-09-12T18:12:41.386629Z',
           'removed_counts': {'rpm.advisory': 6, 'rpm.package': 26},
           'repository_name': 'anothermulti',
           'repository_path': '2d505269/a722d70c-4ac6-4627-a299-10169e05d9e4/b341f7ee-2535-47d8-b83d-2d6eb49440a0',
           'repository_uuid': 'a722d70c-4ac6-4627-a299-10169e05d9e4',
           'url': 'http://pulp-content:8000/api/pulp-content/2d505269/a722d70c-4ac6-4627-a299-10169e05d9e4/b341f7ee-2535-47d8-b83d-2d6eb49440a0/',
           'uuid': '043ad8ed-4a90-4c0f-a5ff-5ae096ef8f5a'},
          {'added_counts': {'rpm.advisory': 4, 'rpm.package': 19},
           'content_counts': {'rpm.advisory': 5, 'rpm.package': 34},
           'created_at': '2024-09-12T18:09:03.76252Z',
           'removed_counts': {'rpm.advisory': 5, 'rpm.package': 12},
           'repository_name': 'needed-test',
           'repository_path': '2d505269/16367cdf-094f-49ed-bb5d-91819316eb58/57ad2e8d-0aa9-4060-ba06-290dcf9fac86',
           'repository_uuid': '16367cdf-094f-49ed-bb5d-91819316eb58',
           'url': 'http://pulp-content:8000/api/pulp-content/2d505269/16367cdf-094f-49ed-bb5d-91819316eb58/57ad2e8d-0aa9-4060-ba06-290dcf9fac86/',
           'uuid': 'b0e75a88-a12a-406b-9c9a-a26335034239'},
          {'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-09-12T11:38:45.388764Z',
           'removed_counts': {},
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_path': '4816cc76/d3d7028e-ec1d-4044-89a4-e37e75cb67b7/7d003060-05d5-4fa0-9ed8-901f10b47bab',
           'repository_uuid': 'd3d7028e-ec1d-4044-89a4-e37e75cb67b7',
           'url': 'http://pulp-content:8000/api/pulp-content/4816cc76/d3d7028e-ec1d-4044-89a4-e37e75cb67b7/7d003060-05d5-4fa0-9ed8-901f10b47bab/',
           'uuid': '745fad01-5372-41d6-bb92-8fac8c9c2d0e'}],
 'links': {'first': '/api/content-sources/v1/templates/d2453d87-ab08-425e-9842-2615f8872161/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/d2453d87-ab08-425e-9842-2615f8872161/snapshots/?limit=100&offset=0'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [7]: 

```




### Comment by @swadeley on September 13, 2024 at 08:07 AM UTC

Hi

These work:
`In [25]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external', sort_by="name:asc")`

`In [26]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external', sort_by="name:desc")`

So all is well.

Thank you


---

## Reviews

### Review by @rverdile - Commented on September 06, 2024 at 04:59 PM UTC

### Review by @rverdile - Commented on September 06, 2024 at 05:08 PM UTC

### Review by @rverdile - Commented on September 06, 2024 at 05:19 PM UTC

### Review by @rverdile - Commented on September 06, 2024 at 05:22 PM UTC

### Review by @swadeley - Commented on September 09, 2024 at 11:36 AM UTC

### Review by @dominikvagner - Commented on September 09, 2024 at 11:58 AM UTC

### Review by @dominikvagner - Commented on September 09, 2024 at 11:59 AM UTC

### Review by @dominikvagner - Commented on September 09, 2024 at 12:07 PM UTC

### Review by @dominikvagner - Commented on September 09, 2024 at 12:11 PM UTC

### Review by @rverdile - Commented on September 09, 2024 at 02:21 PM UTC

### Review by @dominikvagner - Commented on September 09, 2024 at 03:43 PM UTC

### Review by @rverdile - Approved on September 10, 2024 at 05:44 PM UTC

nice job!!

### Review by @Andrewgdewar - Approved on September 10, 2024 at 08:40 PM UTC

Confirmed functionality.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/796*
