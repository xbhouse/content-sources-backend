---
type: pull_request
number: 706
title: "Fixes 3956: consider preexisting rh repos in dao test"
state: merged
author: rverdile
created: 2024-06-14T17:19:34Z
updated: 2024-06-20T20:14:14Z
closed: 2024-06-20T20:14:14Z
merged: 2024-06-20T20:14:14Z
base_branch: main
head_branch: 3956
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/706
---

# Pull Request #706: Fixes 3956: consider preexisting rh repos in dao test

**Author**: @rverdile
**Created**: June 14, 2024 at 05:19 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3956`

## Description

## Summary
- Fixes TaskInfo dao list tests to so they don't fail when red hat repos already exist locally
- Adds an "ExcludeRedHatOrg" filter to the tasks list filter to make the tests cleaner
- Seeds a red hat repo as part of the dao suite SetupTest() to help identify these types of bugs in the future

Although I wrote on the JIRA card that the FetchSnapshotsByDate tests fail, I was not able to reproduce that.

## Testing steps
1. Run `make repos-import-rhel9` then `go run cmd/external_repos/main.go nightly-jobs` to get snapshot tasks with the red hat org
2. `GET /tasks/?exclude_red_hat_org=true` would give a list without the red hat org tasks included.
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 14, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-3956

### Comment by @rverdile on June 18, 2024 at 06:11 PM UTC

updated to expose the new filter

### Comment by @swadeley on June 20, 2024 at 08:06 PM UTC

Hi

```
In [1]: app.content_sources.rest_client.tasks_api.list_tasks(exclude_red_hat_org=True)['meta']['count']
2024-06-20 21:03:40.614 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7fe9d45136a0> object....with url https://<>.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True

2024-06-20 21:03:41.202 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('exclude_red_hat_org', True)]
Out[1]: 0

In [2]: app.content_sources.rest_client.tasks_api.list_tasks(exclude_red_hat_org=False)['meta']['count']
2024-06-20 21:03:44.050 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('exclude_red_hat_org', False)]
Out[2]: 2

In [3]: app.content_sources.rest_client.tasks_api.list_tasks(exclude_red_hat_org=False)
2024-06-20 21:04:29.610 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('exclude_red_hat_org', False)]
Out[3]: 
{'data': [{'created_at': '2024-06-20T20:00:06Z',
           'ended_at': '',
           'error': '',
           'org_id': '-1',
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_uuid': '52bde2f1-525d-4b26-85bc-3e0ad5609ca7',
           'status': 'pending',
           'type': 'snapshot',
           'uuid': '98516c01-0021-4526-95f2-cf56154e9c8b'},
          {'created_at': '2024-06-20T19:58:11Z',
           'ended_at': '2024-06-20T19:58:50Z',
           'error': '',
           'org_id': '-1',
           'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                              '(RPMs)',
           'repository_uuid': '52bde2f1-525d-4b26-85bc-3e0ad5609ca7',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': '29c344f8-44f4-4371-ae0a-a69c0d057154'}],
 'links': {'first': '/api/content-sources/v1/tasks/?exclude_red_hat_org=False&limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?exclude_red_hat_org=False&limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [4]: 
```

---

## Reviews

### Review by @rverdile - Commented on June 14, 2024 at 05:22 PM UTC

### Review by @jlsherrill - Commented on June 17, 2024 at 03:46 PM UTC

### Review by @jlsherrill - Approved on June 20, 2024 at 05:28 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/706*
