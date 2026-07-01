---
type: pull_request
number: 830
title: "Fixes 4643: return object name for red hat task fetch"
state: merged
author: rverdile
created: 2024-09-23T13:52:31Z
updated: 2024-09-24T10:30:25Z
closed: 2024-09-24T10:29:29Z
merged: 2024-09-24T10:29:29Z
base_branch: main
head_branch: 4643
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/830
---

# Pull Request #830: Fixes 4643: return object name for red hat task fetch

**Author**: @rverdile
**Created**: September 23, 2024 at 01:52 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4643`

## Description

## Summary
object name and UUID was not being returned for the fetch endpoint

## Testing steps
1. Fetch a task that has the red hat org is the orgID
2. The object UUID and name should appear. Previously, those would be blank.



---

## Discussion

### Comment by @jlsherrill on September 23, 2024 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-4643

### Comment by @swadeley on September 24, 2024 at 10:28 AM UTC

Hi

testing in stage first for comparison:

```
In [2]: app.content_sources.rest_client.tasks_api.get_task('2ed3a1ba-60e3-4493-bcbf-3e1109588ba3')
2024-09-24 11:24:34.244 [    INFO] [iqe.base.rest_client] REST: GET https://console.stage.redhat.com/api/content-sources/v1/tasks/2ed3a1ba-60e3-4493-bcbf-3e1109588ba3 with query params [] and x-rh-insights-request-id= <>
Out[2]: 
{'created_at': '2024-09-23T19:00:21Z',
 'dependents': ['ff546ca3-50a1-4554-a56b-56b9cc237325'],
 'ended_at': '2024-09-23T19:13:48Z',
 'error': '',
 'object_name': '',
 'object_type': 'repository',
 'object_uuid': '',
 'org_id': '-1',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '2ed3a1ba-60e3-4493-bcbf-3e1109588ba3'}
```

now in ephemeral:
```
In [3]: app.content_sources.rest_client.tasks_api.get_task('64084a97-3b0f-45f2-9846-d28eea04ab26')
2024-09-24 11:25:49.416 [    INFO] [iqe.base.rest_client] REST: GET https://ee-pmuempf2.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/tasks/64084a97-3b0f-45f2-9846-d28eea04ab26 with query params [] and x-rh-insights-request-id=<>
Out[3]: 
{'created_at': '2024-09-24T10:10:49Z',
 'dependents': ['cf5c614c-aedd-460c-9d08-43d7eb8f47c4'],
 'ended_at': '2024-09-24T10:11:25Z',
 'error': '',
 'object_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)',
 'object_type': 'repository',
 'object_uuid': 'b1a2f212-41f3-43c8-8cf6-718b5a46e31b',
 'org_id': '-1',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': '64084a97-3b0f-45f2-9846-d28eea04ab26'}

```


lgtm

---

## Reviews

### Review by @xbhouse - Approved on September 23, 2024 at 04:13 PM UTC

looks good! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/830*
