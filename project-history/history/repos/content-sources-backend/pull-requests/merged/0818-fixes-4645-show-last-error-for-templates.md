---
type: pull_request
number: 818
title: "Fixes 4645: show last error for templates"
state: merged
author: xbhouse
created: 2024-09-16T18:12:03Z
updated: 2024-09-23T14:00:37Z
closed: 2024-09-23T13:33:01Z
merged: 2024-09-23T13:33:01Z
base_branch: main
head_branch: 4645
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/818
---

# Pull Request #818: Fixes 4645: show last error for templates

**Author**: @xbhouse
**Created**: September 16, 2024 at 06:12 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4645`

## Description

## Summary

- Adds new fields to the template response to track errors coming from the UpdateTemplateContent and UpdateLatestSnapshot tasks
- `last_update_snapshot_error` holds the last error coming from UpdateLatestSnapshot or an empty string if there isn't one
- `last_update_task` holds the response from the last UpdateTemplateContent task
- `last_update_task_uuid` holds the uuid from the last UpdateTemplateContent task

## Testing steps

1. Create a custom repository and let it snapshot
2. Create a template with that repo and `use_latest` set to true
3. Fetch the template. `last_update_snapshot_error` should be empty and `last_update_task` should hold the response from the latest UpdateTemplateContent task
4. Change the pulp URL in config.yaml to an invalid URL and edit the template. Fetching the template after this task completes all retries should show an error in `last_update_task` and a failed status
5. Force the UpdateLatestSnapshot task to fail (by either adding code to the Run method in update_latest_snapshot.go or making the pulp URL invalid as soon as a snapshot is completed). Fetch the template again and you should see an error in `last_update_snapshot_error` 
6. Change the pulp URL back to the valid URL and edit the template again to trigger the UpdateTemplateContent task. The `last_update_snapshot_error` field should now be cleared


---

## Discussion

### Comment by @jlsherrill on September 16, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4645

### Comment by @jlsherrill on September 19, 2024 at 05:41 PM UTC

but this now needs a rebase :) 

### Comment by @swadeley on September 20, 2024 at 11:32 AM UTC

Hi

I created a custom repo, created template using that repo, fetched the template :

```
In [7]: app.content_sources.rest_client.templates_api.get_template('5985098d-a7fa-4570-afb0-f1a4ce7c2b53')
2024-09-20 12:26:34.862 [    INFO] [iqe.base.rest_client] REST: GET https://ee-905m0cya.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/5985098d-a7fa-4570-afb0-f1a4ce7c2b53 with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'arch': 'x86_64',
 'created_at': '2024-09-20T11:07:13.155861Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'pr-818',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'Test-a-template',
 'org_id': '3340851',
 'repository_uuids': ['d15f4647-2413-40d1-900a-f0d5fce50aa0',
                      '5c97d0cd-025d-49ee-a2aa-0fc747817787'],
 'rhsm_environment_id': '5985098da7fa4570afb0f1a4ce7c2b53',
 'updated_at': '2024-09-20T11:07:13.155861Z',
 'use_latest': True,
 'uuid': '5985098d-a7fa-4570-afb0-f1a4ce7c2b53',
 'version': '8'}

```
I don't see `last_update_task`

I changed the URL for the custom repo to an invalid one. Opened the template in the UI and saved it. Fetched template again:

```
In [8]: app.content_sources.rest_client.templates_api.get_template('5985098d-a7fa-4570-afb0-f1a4ce7c2b53')
2024-09-20 12:27:03.837 [    INFO] [iqe.base.rest_client] REST: GET https://ee-905m0cya.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/5985098d-a7fa-4570-afb0-f1a4ce7c2b53 with query params [] and x-rh-insights-request-id=<>
Out[8]: 
{'arch': 'x86_64',
 'created_at': '2024-09-20T11:07:13.155861Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'pr-818',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'Test-a-template',
 'org_id': '3340851',
 'repository_uuids': ['d15f4647-2413-40d1-900a-f0d5fce50aa0',
                      '5c97d0cd-025d-49ee-a2aa-0fc747817787'],
 'rhsm_environment_id': '5985098da7fa4570afb0f1a4ce7c2b53',
 'updated_at': '2024-09-20T11:26:52.343137Z',
 'use_latest': True,
 'uuid': '5985098d-a7fa-4570-afb0-f1a4ce7c2b53',
 'version': '8'}
```

No change in:
` 'last_update_snapshot_error': '',`

### Comment by @swadeley on September 23, 2024 at 10:35 AM UTC

Hi @xbhouse 

IIRC you said `last_update_task` would only work in stage. Does that mean `last_update_snapshot_error` will also not work as expected in ephemeral?

I tested again:

```
Out[10]: 
{'arch': 'x86_64',
 'created_at': '2024-09-23T10:14:15.996499Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'save after repo broken using API',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test-template',
 'org_id': '3340851',
 'repository_uuids': ['486f4d5e-ec66-4a34-9956-b1df6d1cc292',
                      'c2c10656-5ded-4fb4-8f81-57c78812fba8'],
 'rhsm_environment_id': '431e9ba8d8004c159b112af97a1bd2e0',
 'updated_at': '2024-09-23T10:30:25.337018Z',
 'use_latest': True,
 'uuid': '431e9ba8-d800-4c15-9b11-2af97a1bd2e0',
 'version': '8'}
```

### Comment by @xbhouse on September 23, 2024 at 11:40 AM UTC

hi @swadeley, `last_update_task` will not work in ephemeral because it depends on candlepin. this task response is from the update_template_content task, which is also how we clear any errors in `last_update_snapshot_error`.

i'm not sure that simulating an error in `last_update_snapshot_error` would work because you'd have to successfully create a snapshot first (so changing the repo URL wouldn't work) and then somehow stop the pulp server immediately after the snapshot was created so the update_latest_snapshot task doesn't run. so we may only be able to test the happy path in stage

### Comment by @swadeley on September 23, 2024 at 01:08 PM UTC

Ok, thank you @xbhouse 

I will merge updated API bindings and then merge this and test in stage.

---

## Reviews

### Review by @jlsherrill - Commented on September 17, 2024 at 07:30 PM UTC

### Review by @jlsherrill - Commented on September 17, 2024 at 08:00 PM UTC

### Review by @xbhouse - Commented on September 18, 2024 at 04:18 PM UTC

### Review by @jlsherrill - Approved on September 19, 2024 at 05:40 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/818*
