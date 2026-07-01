---
type: pull_request
number: 375
title: "Fixes 2018: add task cleanup"
state: merged
author: jlsherrill
created: 2023-08-28T13:47:03Z
updated: 2023-09-11T12:30:29Z
closed: 2023-09-11T12:08:11Z
merged: 2023-09-11T12:08:11Z
base_branch: main
head_branch: task_cleanup
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/375
---

# Pull Request #375: Fixes 2018: add task cleanup

**Author**: @jlsherrill
**Created**: August 28, 2023 at 01:47 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `task_cleanup`

## Description

## Summary
according to the following rules:
* repo delete tasks that are older than 10 days and not failed
* introspect tasks that are older than 10 days
* snapshot tasks where the repository config has been deleted

Previously orphan cleanup was also being run multiple times, this reduces it to just once per nightly job (before starting the introspection jobs)

## Testing steps

On a freshly cleaned db, use this main script:   https://gist.github.com/jlsherrill/efe0bd375a46f32ed3a562e130bb8371
to create some tasks.

This creates similar tasks to the unit test i wrote. 
You can then use 'make db-cli-connect' to look at the tasks.

run: ```go run cmd/external-repos/main.go  nightly-jobs```  to cleanup the tasks.


## QE testing steps
QE can only really easily test the 3rd of the rules above. 
* Create a repository with snapshotting enabled.  
* Wait for it to finish snapshotting.  
* Fetch its snapshot tasks via the tasks api. 
*  Delete the repository. 
* wait ~8 hours
* try to fetch the task again, should be a 404

---

## Discussion

### Comment by @jlsherrill on August 28, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-2018

### Comment by @jlsherrill on September 06, 2023 at 06:13 PM UTC

updated

### Comment by @swadeley on September 11, 2023 at 07:31 AM UTC

/retest

### Comment by @swadeley on September 11, 2023 at 10:36 AM UTC

Hi

I deployed with `NIGHTLY_CRON_JOB='*/20 * * * *'`

I created repo:

```
In [2]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='pulp-core',
   ...:            url='https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/'
   ...:       )

In [3]: app.content_sources.rest_client.repositories_api.create_repository(repo)
2023-09-11 10:46:31.312 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 
2023-09-11 10:46:32.794 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, 
Out[3]: 
{'account_id': '0369233',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_snapshot_task_uuid': 'ce423b55-3cab-44f4-bbc6-ce63e809a1a4',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'pulp-core',
 'org_id': '3340851',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/',
 'uuid': '26324257-692a-47d5-9bc9-cac72b188b06'}

In [4]: app.content_sources.rest_client.tasks_api.list_tasks(limit=20)
2023-09-11 10:59:01.524 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, 
Out[4]: 
{'data': [{'created_at': '2023-09-11T09:46:32Z',
           'ended_at': '2023-09-11T09:46:33Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'pulp-core',
           'repository_uuid': '26324257-692a-47d5-9bc9-cac72b188b06',
           'status': 'completed',
           'type': 'introspect',
           'uuid': 'c34af680-1bc7-4762-9234-26bcd391464e'},
          {'created_at': '2023-09-11T09:46:32Z',
           'ended_at': '2023-09-11T09:46:53Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'pulp-core',
           'repository_uuid': '26324257-692a-47d5-9bc9-cac72b188b06',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': 'ce423b55-3cab-44f4-bbc6-ce63e809a1a4'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=20&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=20&offset=0'},
 'meta': {'count': 2, 'limit': 20, 'offset': 0}}

In [5]: # now we get the tasks, then delete repo, wait for 20 mins, check again 


In [10]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid='26324257-692a-47d5-9bc9-cac72b188b06')
2023-09-11 11:06:37.730 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=T8HAd6RbWRaT2cwVE2o3v4pYMifzKh4l, params=[('repository_uuid', '26324257-692a-47d5-9bc9-cac72b188b06')]
Out[10]: 
{'data': [{'created_at': '2023-09-11T10:06:02Z',
           'ended_at': '',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'pulp-core',
           'repository_uuid': '26324257-692a-47d5-9bc9-cac72b188b06',
           'status': 'pending',
           'type': 'delete-repository-snapshots',
           'uuid': 'bf66fe71-5b5a-4520-89c5-680e4baf948b'},
          {'created_at': '2023-09-11T09:46:32Z',
           'ended_at': '2023-09-11T09:46:33Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'pulp-core',
           'repository_uuid': '26324257-692a-47d5-9bc9-cac72b188b06',
           'status': 'completed',
           'type': 'introspect',
           'uuid': 'c34af680-1bc7-4762-9234-26bcd391464e'},
          {'created_at': '2023-09-11T09:46:32Z',
           'ended_at': '2023-09-11T09:46:53Z',
           'error': '',
           'org_id': '3340851',
           'repository_name': 'pulp-core',
           'repository_uuid': '26324257-692a-47d5-9bc9-cac72b188b06',
           'status': 'completed',
           'type': 'snapshot',
           'uuid': 'ce423b55-3cab-44f4-bbc6-ce63e809a1a4'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=26324257-692a-47d5-9bc9-cac72b188b06',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=26324257-692a-47d5-9bc9-cac72b188b06'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

# Note time above 

```


Some time later after NIGHTLY_CRON_JOB


```
In [12]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid='26324257-692a-47d5-9bc9-cac72b188b06')
2023-09-11 11:29:16.753 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, 
Out[12]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=26324257-692a-47d5-9bc9-cac72b188b06',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=26324257-692a-47d5-9bc9-cac72b188b06'},
 'meta': {'count': 0, 'limit': 100, 'offset': 0}}

```

Try to get the task:

`In [1]: app.content_sources.rest_client.tasks_api.get_task('ce423b55-3cab-44f4-bbc6-ce63e809a1a4')
2023-09-11 13:04:23.224 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, 
<snip>

HTTP response body: {"errors":[{"status":404,"title":"Error fetching task","detail":"Could not find task with UUID ce423b55-3cab-44f4-bbc6-ce63e809a1a4"}]}`

lgtm, thank you

---

## Reviews

### Review by @rverdile - Commented on September 06, 2023 at 02:05 PM UTC

### Review by @rverdile - Commented on September 06, 2023 at 02:07 PM UTC

working well. only found a typo.

### Review by @rverdile - Commented on September 06, 2023 at 02:07 PM UTC

### Review by @rverdile - Approved on September 06, 2023 at 06:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/375*
