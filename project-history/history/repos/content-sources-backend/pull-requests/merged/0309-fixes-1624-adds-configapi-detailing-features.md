---
type: pull_request
number: 309
title: "Fixes 1624: adds config/api detailing features"
state: merged
author: jlsherrill
created: 2023-06-13T15:14:48Z
updated: 2023-06-27T17:30:33Z
closed: 2023-06-27T17:23:03Z
merged: 2023-06-27T17:23:03Z
base_branch: main
head_branch: 1624
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/309
---

# Pull Request #309: Fixes 1624: adds config/api detailing features

**Author**: @jlsherrill
**Created**: June 13, 2023 at 03:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1624`

## Description

## Summary
Adds a setting to determine if snapshotting is enabled
Adds a per account and per user setting to determine if the user can snapshot repositories
Adds an api to return if the feature is 1) enabled 2) accessible to the current user
Updated the header.sh to provide the ability to pass in an account id optionally 

## Testing steps

copy new config options from the example config file
```/bin/bash
# should show as accessible
 curl localhost:8000/api/content-sources/v1.0/features/  -H "`./scripts/header.sh 1 snapUser`" 
# should show not as accessible
 curl localhost:8000/api/content-sources/v1.0/features/  -H "`./scripts/header.sh 1 noAccess`" 
# should show as accessible
 curl localhost:8000/api/content-sources/v1.0/features/  -H "`./scripts/header.sh 1 noAccess snapAccount`" 
# should show not as accessible
 curl localhost:8000/api/content-sources/v1.0/features/  -H "`./scripts/header.sh 1 noAccess NoAccess`" 
```



---

## Discussion

### Comment by @jlsherrill on June 13, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-1624

### Comment by @jlsherrill on June 14, 2023 at 01:10 PM UTC

/retest

### Comment by @Andrewgdewar on June 15, 2023 at 09:14 PM UTC

Ummn hi! Looking great! 

But can you change the response from:
`[{"name":"snapshot", "enabled":true, "accessible":"true"}]`
to 
`{"snapshot":{enabled:true, accessible:true}}`
🤣 



### Comment by @jlsherrill on June 20, 2023 at 02:10 PM UTC

/retest

### Comment by @swadeley on June 26, 2023 at 10:32 AM UTC

Hi

More info on the example config file please.

Testing in a shell:
```
In [3]: app.content_sources.rest_client.features_api.list_features()
2023-06-26 10:47:37.287 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=d4Gq5Ka9M0luh4wJGEc9TiGHBUpECszq, params=[]
Out[3]: {'snapshots': {'accessible': True, 'enabled': False}}

```
I created a repo
```

In [2]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='fedoratest',
   ...:            url='https://stephenw.fedorapeople.org/multirepos/1/repo01/',
   ...:         )
```

but task status shows failed:

```
In [6]: app.content_sources.rest_client.tasks_api.list_tasks()
2023-06-26 11:14:46.364 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=qOLccxIsigHWWwCYNo2SDlUS2TEW0V4n, params=[]
Out[6]: 
{'data': [{'created_at': '2023-06-26T10:11:25Z',
           'ended_at': '2023-06-26T10:11:25Z',
           'error': '',
           'org_id': '3340851',
           'status': 'completed',
           'uuid': 'bd4d6846-35e2-4a7d-8ed9-d5d635cee746'},
          {'created_at': '2023-06-26T10:11:25Z',
           'ended_at': '2023-06-26T10:11:25Z',
           'error': '404 Not Found',
           'org_id': '3340851',
           'status': 'failed',
           'uuid': 'a843356c-6090-497b-958f-ea730f3e5dad'}],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```


I think we need a rebase to pull in changes from PR 308. 
I should be able to do that locally for testing.

EDIT: No, I see `Fixes 1501: add API to list snapshots (#308)` is in the backend repo PR 309 branch from which I updated the API.


### Comment by @swadeley on June 26, 2023 at 10:54 AM UTC

/retest

### Comment by @jlsherrill on June 26, 2023 at 01:02 PM UTC



To control from ephemeral, you can set the following env variables:

FEATURES_SNAPSHOTS_ENABLED="true" 
FEATURES_SNAPSHOTS_ACCOUNTS="123,456"

If FEATURES_SNAPSHOTS_ACCOUNTS is not set or empty string, then all accounts can access it.

I think we probably need deployment.yaml updates for it to work in ephemeral, let me do that and very myself there

### Comment by @jlsherrill on June 26, 2023 at 03:38 PM UTC

okay, this should work in ephemeral now, the way i tested, my bonfire conifg: (note i don't have pulp here, it doesn't matter, this was just quicker):
```
apps:
  - name: content-sources
    components:
      - name: content-sources-backend
        host: local
        repo: /home/jlsherri/git/content-sources-backend
        path: deployments/deployment.yaml
        parameters:
          IMAGE_TAG: pr-309-latest
      - name: content-sources-frontend
        host: github
        repo: content-services/content-sources-frontend
        path: /deploy/frontend.yaml
```

This config would show that the feature is disabled (accessibility doesn't really matter).

adding to the parameter list:
 
```
          FEATURES_SNAPSHOTS_ENABLED: true
```

would show that its enabled, and accessible to all users

Redploying with these params:
 
```
          FEATURES_SNAPSHOTS_ENABLED: true
          FEATURES_SNAPSHOTS_ACCOUNTS: "1111"
```
Would show it enabled, but not accessible to the jdoe user account (account # 12345).



Redploying with these params:
 
```
          FEATURES_SNAPSHOTS_ENABLED: true
          FEATURES_SNAPSHOTS_ACCOUNTS: "12345"
```

Would show its enable and accessible to the jdoe user.

### Comment by @swadeley on June 27, 2023 at 08:42 AM UTC

Hi

with config like so:

```
apps:
  - name: content-sources
    components:
      - name: content-sources-backend
        host: github
        repo: content-services/content-sources-backend
        path: deployments/deployment.yaml
        parameters:
          IMAGE_TAG: pr-309-latest
          FEATURES_SNAPSHOTS_ENABLED: true
      - name: content-sources-frontend
        host: github
        repo: content-services/content-sources-frontend
        path: /deploy/frontend.yaml
```

I still get
```

In [6]: app.content_sources.rest_client.features_api.list_features()
2023-06-27 09:38:25.886 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=8kLjQSqMel3kVG7uRAiojixGe2ez1tLH, params=[]
Out[6]: {'snapshots': {'accessible': True, 'enabled': False}}
```


### Comment by @swadeley on June 27, 2023 at 11:51 AM UTC

Hi

I also tried deploying with:
`--set-parameter content-sources-backend/FEATURES_SNAPSHOTS_ENABLED=true`

but still get:
`Out[1]: {'snapshots': {'accessible': True, 'enabled': False}}`


### Comment by @jlsherrill on June 27, 2023 at 12:46 PM UTC

you'll need to specify a custom deployment.yaml from this branch:

```
- name: content-sources-backend
  host: local
  repo: /home/jlsherri/git/content-sources-backend
  path: deployments/deployment.yaml
```

### Comment by @swadeley on June 27, 2023 at 04:34 PM UTC

Hi

Thanks for the info, progress now.


with
        parameters:
          IMAGE_TAG: pr-309-latest
          FEATURES_SNAPSHOTS_ENABLED: true
I get:
Out[1]: {'snapshots': {'accessible': True, 'enabled': True}}

--------------------

With
          FEATURES_SNAPSHOTS_ENABLED: true
          FEATURES_SNAPSHOTS_ACCOUNTS: "1111"

I get:

Out[1]: {'snapshots': {'accessible': False, 'enabled': True}}

--------------------------

with
          FEATURES_SNAPSHOTS_ENABLED: true
          FEATURES_SNAPSHOTS_ACCOUNTS: "<jdoe_acc_num>"

I get
Out[1]: {'snapshots': {'accessible': True, 'enabled': True}}




---

## Reviews

### Review by @rverdile - Commented on June 15, 2023 at 06:37 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 01:31 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 01:33 PM UTC

### Review by @jlsherrill - Commented on June 16, 2023 at 01:47 PM UTC

### Review by @jlsherrill - Commented on June 16, 2023 at 01:48 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 01:55 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 01:59 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 02:44 PM UTC

### Review by @rverdile - Commented on June 16, 2023 at 03:34 PM UTC

### Review by @rverdile - Approved on June 20, 2023 at 01:58 PM UTC

### Review by @swadeley - Commented on June 27, 2023 at 12:09 PM UTC

### Review by @jlsherrill - Commented on June 27, 2023 at 02:51 PM UTC

### Review by @swadeley - Commented on June 27, 2023 at 03:22 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/309*
