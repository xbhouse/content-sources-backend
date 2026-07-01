---
type: pull_request
number: 346
title: "Fixes 1963: nightly sync"
state: closed
author: dpang314
created: 2023-08-02T19:28:33Z
updated: 2023-10-04T13:44:43Z
closed: 2023-09-11T19:17:28Z
base_branch: main
head_branch: HMS-1963-nightly-cron-sync
labels: []
url: https://github.com/content-services/content-sources-backend/pull/346
---

# Pull Request #346: Fixes 1963: nightly sync

**Author**: @dpang314
**Created**: August 02, 2023 at 07:28 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `HMS-1963-nightly-cron-sync`

## Description

## Summary

Adds repository snapshotting to nightly tasks. When the snapshot feature is enabled in the config, nightly tasks will create a snapshot of all repositories with snapshot = true that haven't been snapshotted in the last 24 hours.

## Testing steps

Enable the snapshot feature in the config and create a repository with snapshotting enabled. 
`go run cmd/external-repos/main.go nightly-jobs` triggers the sync task. 
If there is a repository with snapshotting enabled that hasn't been snapshotted in the past 24 hours (including one that has no snapshots), then a snapshot task for that repository will be enqueued. Running the nightly-job will log that the snapshot task was enqueued.

---

## Discussion

### Comment by @jlsherrill on August 02, 2023 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-1963

### Comment by @swadeley on August 04, 2023 at 08:13 PM UTC

Hi, is this ready for testing? Please rebase.

### Comment by @jlsherrill on August 15, 2023 at 10:03 PM UTC

feel free to qe this, but we need to probably hold off on merging until snapshotting is fixed in stage

### Comment by @swadeley on August 16, 2023 at 06:10 AM UTC

> feel free to qe this, but we need to probably hold off on merging until snapshotting is fixed in stage

HI, I can test this on Friday, so I can leave the test ephemeral env running overnight.

### Comment by @jlsherrill on August 16, 2023 at 01:03 PM UTC

okay, stage is fixed!

### Comment by @swadeley on August 17, 2023 at 08:56 PM UTC

Hi

test started:

```
In [1]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='fedoratest-1',
   ...:            url='https://stephenw.fedorapeople.org/multirepos/5/repo01/'
   ...:       )

In [2]: app.content_sources.rest_client.repositories_api.create_repository(repo)
2023-08-17 21:54:54.336 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7f1417021a60> object....with url https://env-ephemeral-jo9225-wi6oqp1z.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
2023-08-17 21:54:54.337 [    INFO] [iqe.base.auth] Available auth types: ['basic', 'cert', 'identity', 'jwt', 'uht']
2023-08-17 21:54:54.337 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2023-08-17 21:54:56.670 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=nZUgqDVxGZOSdJIhsAY6GH7W7kMrM1wn, params=[]
Out[2]: 
{'account_id': '0369233',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'fedoratest-1',
 'org_id': '3340851',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo01/',
 'uuid': '5bad1ff3-2757-4b8b-8878-8947be7d8e96'}

In [3]: app.content_sources.rest_client.tasks_api.list_tasks()
2023-08-17 21:55:18.573 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=2hM2ZiIIIuRwT1TETyGInrOxtHT7ZM7R, params=[]
Out[3]: 
{'data': [{'created_at': '2023-08-17T20:54:56Z',
           'ended_at': '2023-08-17T20:54:56Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': '43016219-585f-4778-9ebb-3a50f980d0f3'},
          {'created_at': '2023-08-17T20:54:56Z',
           'ended_at': '2023-08-17T20:55:06Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': '95a44239-724c-4b3d-943b-0231a95f7f61'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [4]: 


```

### Comment by @swadeley on August 21, 2023 at 06:19 AM UTC

/retest


### Comment by @swadeley on August 22, 2023 at 09:21 AM UTC

Hi,

I see repos that have been deleted being introspected, not sure yet if that happens just once. Asked in slack.




### Comment by @swadeley on August 23, 2023 at 02:24 PM UTC

Hi @Andrewgdewar 

As discussed, I was seeing introspection never ending while testing this.

Please rebase for me to pull in latest Pulp improvements.

Thank you

### Comment by @swadeley on August 28, 2023 at 04:39 PM UTC

/retest

### Comment by @swadeley on August 31, 2023 at 06:32 PM UTC

/retest

### Comment by @swadeley on September 04, 2023 at 02:16 PM UTC

/retest

### Comment by @swadeley on September 06, 2023 at 08:07 AM UTC

   Need to test snapshot is created every 24 hours if no snapshot was already made
   We can check  'last_snapshot' values and  'last_snapshot_uuid' 
    --------------------------------------------------------------------------------------------------------------------------------


     
```
    In [1]:      repo = dict(
       ...:            snapshot=True,
       ...:            name='pulp-core',
       ...:            url='https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/'
       ...:       )
     
    In [2]: app.content_sources.rest_client.repositories_api.create_repository(repo)
    2023-09-04 19:21:59.837 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7f05a9983940> object....with url https://env-ephemeral-wfqmgj-ek8gfkpo.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
    2023-09-04 19:21:59.837 [    INFO] [iqe.base.auth] Available auth types: ['basic', 'cert', 'identity', 'jwt', 'uht']
    2023-09-04 19:21:59.837 [    INFO] [iqe.base.auth] Setting auth_type to jwt
    2023-09-04 19:22:01.075 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=m47XLg3re2D3ZEwmSHMYVsehBRyU89Xo, params=[]
    Out[2]:
    {'account_id': '0369233',
     'distribution_arch': 'any',
     'distribution_versions': ['any'],
     'failed_introspections_count': 0,
     'gpg_key': '',
     'last_introspection_error': '',
     'last_introspection_time': '',
     **last_snapshot_task_uuid**: '50e7f152-c65f-4591-8318-807762e204fc',
     'last_success_introspection_time': '',
     'last_update_introspection_time': '',
     'metadata_verification': False,
     'name': 'pulp-core',
     'org_id': '3340851',
     'package_count': 0,
     'snapshot': True,
     'status': 'Pending',
     'url': 'https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/',
     'uuid': '5ccad0eb-56e3-4ba1-a76f-60842bb9899f'}
     
     
    In [1]: app.content_sources.rest_client.repositories_api.list_repositories()
    2023-09-05 13:43:38.486 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7f0cf55e1790> object....with url https://env-ephemeral-wfqmgj-ek8gfkpo.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
    2023-09-05 13:43:38.486 [    INFO] [iqe.base.auth] Available auth types: ['basic', 'cert', 'identity', 'jwt', 'uht']
    2023-09-05 13:43:38.486 [    INFO] [iqe.base.auth] Setting auth_type to jwt
    2023-09-05 13:43:39.427 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=qw1eHltXTG5HaP1AwXkvQsqbGKBzFVrG, params=[]
    Out[1]:
    {'data': [{'account_id': '0369233',
               'distribution_arch': 'any',
               'distribution_versions': ['any'],
               'failed_introspections_count': 0,
               'gpg_key': '',
               'last_introspection_error': '',
               'last_introspection_time': '2023-09-04T18:46:01Z',
               'last_snapshot': {'added_counts': {'rpm.package': 121},
                                 'content_counts': {'rpm.package': 121},
                                 '**created_at**': '2023-09-04T18:46:00.633897Z',
                                 'removed_counts': {},
                                 'repository_path': ''},
               '**last_snapshot_task_uuid**': '50e7f152-c65f-4591-8318-807762e204fc',
               '**last_snapshot_uuid**': '679906d4-091e-4d26-bf61-3586e3796a3c',
               'last_success_introspection_time': '2023-09-04T18:46:01Z',
               'last_update_introspection_time': '2023-09-04T18:46:01Z',
               'metadata_verification': False,
               'name': 'pulp-core',
               'org_id': '3340851',
               'package_count': 121,
               'snapshot': True,
               'status': 'Valid',
               'url': 'https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/',
               'uuid': '5ccad0eb-56e3-4ba1-a76f-60842bb9899f'}],
     'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0',
               'last': '/api/content-sources/v1/repositories/?limit=100&offset=0'},
     'meta': {'count': 1, 'limit': 100, 'offset': 0}}```

----------------------------------------- the next day ----------------------

      ```{'account_id': '0369233',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2023-09-06T00:00:07Z',
           'last_snapshot': {'added_counts': {'rpm.package': 121},
                             'content_counts': {'rpm.package': 121},
                             '**created_at**': '2023-09-04T18:46:00.633897Z',
                             'removed_counts': {},
                             'repository_path': ''},
           **last_snapshot_task_uuid**: '50e7f152-c65f-4591-8318-807762e204fc',
           **last_snapshot_uuid**: '679906d4-091e-4d26-bf61-3586e3796a3c',
           'last_success_introspection_time': '2023-09-06T00:00:07Z',
           'last_update_introspection_time': '2023-09-04T18:46:01Z',
           'metadata_verification': False,
           'name': 'pulp-core',
           'org_id': '3340851',
           'package_count': 121,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://yum.theforeman.org/pulpcore/3.4/el7/x86_64/',
           'uuid': '5ccad0eb-56e3-4ba1-a76f-60842bb9899f'}]```


Does not seem to work, I see the same task ID:

**last_snapshot_task_uuid**: '50e7f152-c65f-4591-8318-807762e204fc',

and the **created_at** time did not change.

### Comment by @swadeley on September 06, 2023 at 02:09 PM UTC

Ok, my misunderstanding, in case there is no change to the repos then there should only be a task but no actual snapshot created.


### Comment by @swadeley on September 06, 2023 at 05:22 PM UTC

/retest

### Comment by @jlsherrill on September 11, 2023 at 12:28 PM UTC

i'm going to work to add the last_snapshot_uuid updating

### Comment by @jlsherrill on September 11, 2023 at 07:17 PM UTC

closing in favor of https://github.com/content-services/content-sources-backend/pull/387

---

## Reviews

### Review by @jlsherrill - Commented on August 04, 2023 at 03:30 PM UTC

### Review by @dpang314 - Commented on August 04, 2023 at 03:43 PM UTC

### Review by @jlsherrill - Commented on August 04, 2023 at 06:17 PM UTC

### Review by @jlsherrill - Approved on August 15, 2023 at 08:15 PM UTC

Thanks @dpang314 !!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/346*
