---
type: pull_request
number: 331
title: "Fixes 1939: pulp domain integration"
state: merged
author: jlsherrill
created: 2023-07-07T16:15:29Z
updated: 2023-08-04T20:00:52Z
closed: 2023-08-04T19:53:12Z
merged: 2023-08-04T19:53:12Z
base_branch: main
head_branch: domain
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/331
---

# Pull Request #331: Fixes 1939: pulp domain integration

**Author**: @jlsherrill
**Created**: July 07, 2023 at 04:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `domain`

## Description

## Summary

* Updates to use our own docker-compose for pulp
* Create a domain via a go routine at api calls to update/create if snapshotting is to be performed
* Split the pulp client into 2 interfaces, a global and a non-global.  Global doesn't require a domain.  Underlying struct is the same.

NOTE:  Before merging, we need to set  `DOMAIN_ENABLED: "true"` in the pulp deployment for ephemeral, within our app-inteface deploy.yaml:
```
  - namespace:
      $ref: /services/insights/ephemeral/namespaces/ephemeral-base.yml
    disable: true  # do not create an app-sre deploy job for ephemeral namespace
    ref: internal
    parameters:
      IMAGE_TAG: latest
```
just add below IMAGE_TAG:
```
DOMAIN_ENABLED: "true"
```

## Testing steps

 * make sure to be on main
 * make compose-clean
 * checkout this pr
 * make compose-up
 * create a repo:
 ```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqZG9lIn0sImludGVybmFsIjp7Im9yZ19pZCI6IjEyMyJ9fX0K" \
    -d "{\"name\":\"pulp 3.14\",\"url\":\"http://yum.theforeman.org/pulpcore/3.14/el8/x86_64/\",\"snapshot\": true}"
```
Use a bulk create with a different org:
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/bulk_create/"    -H "Content-Type: application/json"    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJhc2RmYXNkZiJ9LCJhY2NvdW50X251bWJlciI6ImFhc2RmIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiYWFzZGYifX19Cg=="    -d "[{\"name\":\"pulp 3.14\",\"url\":\"http://yum.theforeman.org/pulpcore/3.14/el8/x86_64/\",\"snapshot\": true}, {\"name\":\"pulp 3.16\",\"url\":\"http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/\",\"snapshot\": true}]"
```

List the snapshots for the repo and see that the new 'RepositoryPath' shows a domain (we still store the distribution path in the db, but thats not relevant to the user)

Also, do a bunch of bulk creates on a new org, verify that there are no errors.

### Ephemeral testing:
1.  Create a bonfire config:
```
appsFile:
  host: gitlab
  repo: insights-platform/cicd-common
  path: bonfire_configs/ephemeral_apps.yaml
apps:
 - name: content-sources
   components:
     - name: pulp
       host: github
       repo: pulp/pulp-clowder-deployments
       path: deploy/clowdapp.yaml
       parameters:
         DOMAIN_ENABLED: 'true'
     - name: content-sources-backend
       host: local
       repo: /home/jlsherri/git/content-sources-backend
       path: deployments/deployment.yaml
       parameters:
         IMAGE_TAG: pr-331-latest
         FEATURES_ADMIN_TASKS_ENABLED: true
         FEATURES_SNAPSHOTS_ENABLED: true
     - name: content-sources-frontend
       host: github
       repo: content-services/content-sources-frontend

```

reserve a namespace and deploy:
```bonfire deploy content-sources --frontends=true```

Once its up do a `namespace describe` to get the info and curl to create a repo (replacing with the correct info):
```
 curl -v -u jdoe:OK4dkBMZ7CKhZ8lH   -X POST https://env-ephemeral-bvyuo3-d7v745fk.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1.0/repositories/  -d '{"name":"pulp 3.18","url":"http://yum.theforeman.org/pulpcore/3.18/el8/x86_64/","snapshot": true}' -H 'Content-Type: application/json'
```
Monitor the admin tasks, and verify that the task is created.  You can curl the snapshots api to confirm the path:

```
curl -v -u jdoe:OSRFgOKvL5ZLdQbO https://env-ephemeral-7sp1sd-1o8n8w1w.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1.0/repositories/21202889-e9a2-4edf-87fa-b673c5f0339c/snapshots/ | jq
```

for example, my output shows a repository_path of: 
```
    "repository_path_path": "bf7f7fae/21202889-e9a2-4edf-87fa-b673c5f0339c/b959ce1f-7b6c-467c-857c-6afa3808abd6",
```
that first char set `bf7f7fae` is the domain name.

you might also try bulk creation here

### testing with local minion:
 * Login to minio at http://localhost:9001
 * create a bucket, named test
 * create a user (under identity) with test/password.  Give them the read/write role
 * add this to your config.yaml:
 ```  
pulp:
    server: http://localhost:8080
    username: admin
    password: password
    storage_type: object #object or local
    custom_repo_objects:
      url: http://minio:9000
      accessKey: test
      secretKey: password
      name: test
      region: rdu
```
* Create a repo in a new organization id (domains aren't currently updated).

---

## Discussion

### Comment by @jlsherrill on July 07, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-1939

### Comment by @jlsherrill on July 28, 2023 at 05:28 PM UTC

 Note, i had an error in my example config .yaml changes for minio, thats now corrected.

### Comment by @jlsherrill on August 02, 2023 at 06:52 PM UTC

working to rebase

### Comment by @swadeley on August 04, 2023 at 07:51 PM UTC

Hello

with the default ephemeral user jdoe I created some repos with snapshots:

```
In [2]:  app.user.to_dict()
Out[2]: 
{'username': 'jdoe',
 'password': '<redacted>',
 'idp_username': 'jdoe',
 'idp_password': '<redacted>',
 'identity': {'account_number': '0369233',
  'org_id': '3340851',                    <---- NOTE: ends in 51

          {'account_id': '0369233',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2023-08-04T16:18:42Z',
           'last_success_introspection_time': '2023-08-04T16:18:42Z',
           'last_update_introspection_time': '2023-08-04T16:18:42Z',
           'metadata_verification': False,
           'name': 'fedoratest-with-snapshot2',
           'org_id': '3340851',
           'package_count': 1,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://stephenw.fedorapeople.org/multirepos/2/repo03/',
           'uuid': '63e406b5-bc77-4cb5-93df-5d3a8d2d3c6e'},



In [3]: app.content_sources.rest_client.repositories_api.list_snapshots('63e406b5-bc77-4cb5-93df-5d3a8d2d3c6e')
2023-08-04 20:42:55.704 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=gH0YAmlkKR8tdJENtRzMSgU6ujA3ilKt, params=[]
Out[3]: 
{'data': [{'content_counts': {'rpm.package': 1},
           'created_at': '2023-08-04T16:23:44.642322Z',
           'repository_path': '2d18751c/63e406b5-bc77-4cb5-93df-5d3a8d2d3c6e/d0dd9813-717c-4506-b074-6471263cba54'}],
 'links': {'first': '/api/content-sources/v1/repositories/63e406b5-bc77-4cb5-93df-5d3a8d2d3c6e/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/63e406b5-bc77-4cb5-93df-5d3a8d2d3c6e/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

```

Notice in the above , the string  '`repository_path': '2d18751c/6`


I created a new user and bumped the org id and acc number by 1:
```

    In [10]: app.user.to_dict()
    Out[10]:
    {'username': 'zybijmjhejrslaqe',
     'password': '<redacted>',
     'idp_username': 'zybijmjhejrslaqe',
     'idp_password': '<redacted>',
     'identity': {'account_number': '0369234',  <-- bumped by 1
      'org_id': '3340852',  <-- bumped by 1
     
    In [1]: app.content_sources.rest_client.repositories_api.list_repositories(limit=6,offset=0)

    Out[1]:
    {'data': [],
     'links': {'first': '/api/content-sources/v1/repositories/?limit=6&offset=0',
               'last': '/api/content-sources/v1/repositories/?limit=6&offset=0'},
     'meta': {'count': 0, 'limit': 6, 'offset': 0}}
     
    In [2]:      repo = dict(
       ...:            snapshot=True,
       ...:            name='fedoratest-with-other-org',
       ...:            url='https://stephenw.fedorapeople.org/multirepos/3/repo04/',
       ...:         )
     
    In [3]: response =  app.content_sources.rest_client.repositories_api.create_repository(repo)

    In [4]: response
    Out[4]:
    {'account_id': '0369234',
     'distribution_arch': 'any',
     'distribution_versions': ['any'],
     'failed_introspections_count': 0,
     'gpg_key': '',
     'last_introspection_error': '',
     'last_introspection_time': '',
     'last_success_introspection_time': '',
     'last_update_introspection_time': '',
     'metadata_verification': False,
     'name': 'fedoratest-with-other-org',
     'org_id': '3340852',
     'package_count': 0,
     'snapshot': True,
     'status': 'Pending',
     'url': 'https://stephenw.fedorapeople.org/multirepos/3/repo04/',
     'uuid': '74eb9ca0-fd67-4109-9870-d64a5dd6d6f7'}
     
    In [5]: app.content_sources.rest_client.tasks_api.list_tasks()

    Out[5]:
    {'data': [{'created_at': '2023-08-04T19:27:29Z',
               'ended_at': '2023-08-04T19:27:29Z',
               'error': '',
               'org_id': '3340852',
               'status': 'completed',
               'uuid': '51f83cb5-2424-4186-b1a3-b6f1d2f75bb0'},
              {'created_at': '2023-08-04T19:27:29Z',
               'ended_at': '',
               'error': '',
               'org_id': '3340852',
               'status': 'running',
               'uuid': '8983348d-c1c7-4620-a740-151c6ef148d3'},
           
     
    In [6]: app.content_sources.rest_client.repositories_api.list_repositories(limit=6,offset=0)
    2023-08-04 20:29:39.559 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=pV8X69WAsKzG4zM11dLSL1JPHyvmMWxu, params=[('limit', 6), ('offset', 0)]
    Out[6]:
    {'data': [{'account_id': '0369234',
               'distribution_arch': 'any',
               'distribution_versions': ['any'],
               'failed_introspections_count': 0,
               'gpg_key': '',
               'last_introspection_error': '',
               'last_introspection_time': '2023-08-04T19:27:29Z',
               'last_success_introspection_time': '2023-08-04T19:27:29Z',
               'last_update_introspection_time': '2023-08-04T19:27:29Z',
               'metadata_verification': False,
               'name': 'fedoratest-with-other-org',
               'org_id': '3340852',
               'package_count': 1,
               'snapshot': True,
               'status': 'Valid',
               'url': 'https://stephenw.fedorapeople.org/multirepos/3/repo04/',
               'uuid': '74eb9ca0-fd67-4109-9870-d64a5dd6d6f7'}],
     'links': {'first': '/api/content-sources/v1/repositories/?limit=6&offset=0',
               'last': '/api/content-sources/v1/repositories/?limit=6&offset=0'},
     'meta': {'count': 1, 'limit': 6, 'offset': 0}}
     
    In [7]: app.content_sources.rest_client.repositories_api.get_repository('74eb9ca0-fd67-4109-9870-d64a5dd6d6f7')
    2023-08-04 20:32:28.502 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=tehE7mndUevuXBiamKrkYz2V63vtluye, params=[]
    Out[7]:
    {'account_id': '0369234',
     'distribution_arch': 'any',
     'distribution_versions': ['any'],
     'failed_introspections_count': 0,
     'gpg_key': '',
     'last_introspection_error': '',
     'last_introspection_time': '2023-08-04T19:27:29Z',
     'last_success_introspection_time': '2023-08-04T19:27:29Z',
     'last_update_introspection_time': '2023-08-04T19:27:29Z',
     'metadata_verification': False,
     'name': 'fedoratest-with-other-org',
     'org_id': '3340852',
     'package_count': 1,
     'snapshot': True,
     'status': 'Valid',
     'url': 'https://stephenw.fedorapeople.org/multirepos/3/repo04/',
     'uuid': '74eb9ca0-fd67-4109-9870-d64a5dd6d6f7'}
     
    In [8]: app.content_sources.rest_client.repositories_api.list_snapshots('74eb9ca0-fd67-4109-9870-d64a5dd6d6f7')
    2023-08-04 20:33:01.699 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=WrCK0ILPzP5NLjf8Y9vMJ1LZCx0caHsT, params=[]
    Out[8]:
    {'data': [{'content_counts': {'rpm.package': 1},
               'created_at': '2023-08-04T19:28:07.455003Z',
               'repository_path': '12d903db/74eb9ca0-fd67-4109-9870-d64a5dd6d6f7/92dff4ed-a041-4de1-9fd1-2bf1898a395c'}],
     'links': {'first': '/api/content-sources/v1/repositories/74eb9ca0-fd67-4109-9870-d64a5dd6d6f7/snapshots/?limit=100&offset=0',
               'last': '/api/content-sources/v1/repositories/74eb9ca0-fd67-4109-9870-d64a5dd6d6f7/snapshots/?limit=100&offset=0'},
     'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

In the above:   ` 'repository_path': '12d903db`

### Comment by @swadeley on August 04, 2023 at 07:52 PM UTC

If you check the admin UI

In this json,
```
{4 items
"SyncTaskHref":string"/pulp/2d18751c/api/v3/tasks/0189c158-1deb-7175-bf6a-6bcce110ec16/"
"SnapshotIdent":NULL
"PublicationTaskHref":string"/pulp/2d18751c/api/v3/tasks/0189c15c-34c5-7291-8f94-0e9e66c94055/"
"DistributionTaskHref":NULL
}
```

2d18751c  is the domain

---

## Reviews

### Review by @rverdile - Commented on July 25, 2023 at 07:39 PM UTC

### Review by @rverdile - Commented on July 25, 2023 at 07:47 PM UTC

### Review by @rverdile - Commented on July 25, 2023 at 07:56 PM UTC

### Review by @rverdile - Commented on July 25, 2023 at 08:20 PM UTC

### Review by @rverdile - Commented on July 27, 2023 at 03:08 PM UTC

### Review by @rverdile - Commented on July 27, 2023 at 03:19 PM UTC

### Review by @rverdile - Commented on July 27, 2023 at 06:46 PM UTC

The domain is still created and no actual functionality is disrupted, but I'm consistently getting ` Error creating domain error="500 Internal Server Error" body="\n<!doctype html>\n<html lang=\"en\">\n<head>\n  <title>Server Error (500)...` when using bulk create with a new orgID. 

It seems like you handled the "creation at the same time" edge case. Are you still seeing this 500 error?

### Review by @jlsherrill - Commented on July 28, 2023 at 05:11 PM UTC

### Review by @rverdile - Commented on July 28, 2023 at 07:18 PM UTC

### Review by @jlsherrill - Commented on July 28, 2023 at 08:06 PM UTC

### Review by @jlsherrill - Commented on July 28, 2023 at 08:12 PM UTC

### Review by @jlsherrill - Commented on July 28, 2023 at 08:37 PM UTC

### Review by @rverdile - Commented on July 31, 2023 at 04:19 PM UTC

### Review by @rverdile - Commented on July 31, 2023 at 04:20 PM UTC

### Review by @jlsherrill - Commented on July 31, 2023 at 04:24 PM UTC

### Review by @jlsherrill - Commented on July 31, 2023 at 04:24 PM UTC

### Review by @jlsherrill - Commented on July 31, 2023 at 04:36 PM UTC

### Review by @rverdile - Commented on July 31, 2023 at 04:38 PM UTC

### Review by @rverdile - Commented on July 31, 2023 at 08:40 PM UTC

### Review by @jlsherrill - Commented on July 31, 2023 at 08:48 PM UTC

### Review by @rverdile - Approved on July 31, 2023 at 08:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/331*
