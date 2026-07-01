---
type: pull_request
number: 844
title: "Fixes 4716: Add \"latest\" snapshot distribution"
state: merged
author: Andrewgdewar
created: 2024-10-10T17:40:39Z
updated: 2024-11-01T18:30:31Z
closed: 2024-11-01T18:14:06Z
merged: 2024-11-01T18:14:06Z
base_branch: main
head_branch: HMS-4716
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/844
---

# Pull Request #844: Fixes 4716: Add "latest" snapshot distribution

**Author**: @Andrewgdewar
**Created**: October 10, 2024 at 05:40 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-4716`

## Description

## Summary

- When a new snapshot is taken either via a sync, or via an upload, a new distribution there is a new distribution in pulp pointing to /latest.

- The full URL to the above 'latest' snapshot distribution is returned with each snapshotted & upload repository, as the attribute:  "latest_snapshot_url"

- New route to download the repo-config with associated 'latest' url @/repositories/UUID/config.repo  

## Testing steps

- Create a custom repository > wait for snapshot > check the response for the "latest_snapshot_url"
- Call the /repositories/UUID/config.repo endpoint with the above repository uuid, see the latest_snapshot_url represented there as "url"
- Check that you can download the above distribution.


---

## Discussion

### Comment by @jlsherrill on October 10, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-4716

### Comment by @swadeley on October 16, 2024 at 09:45 AM UTC

Hi

build test is failing and if I try deploy I get:

`chrome-service-api-69f784949-pxv7v                               0/2     Init:ImagePullBackOff   0             36m`




### Comment by @jlsherrill on October 18, 2024 at 07:40 PM UTC

getting a task error when i create an upload repository and upload files to it:
```
3:33PM WRN [Finished Task] task failed error="could not find a distribution href in task: 0xc000b5d470" request_id=aa0b853f-72ff-4421-a6b1-52752578a2d2 task_id=24422f4b-965f-4de1-9082-f79885dbcf58 task_type=add-uploads-repository
```

### Comment by @Andrewgdewar on October 21, 2024 at 04:14 PM UTC

> getting a task error when i create an upload repository and upload files to it:
> 
> ```
> 3:33PM WRN [Finished Task] task failed error="could not find a distribution href in task: 0xc000b5d470" request_id=aa0b853f-72ff-4421-a6b1-52752578a2d2 task_id=24422f4b-965f-4de1-9082-f79885dbcf58 task_type=add-uploads-repository
> ```

Looking now

### Comment by @Andrewgdewar on October 21, 2024 at 05:42 PM UTC

Updated based on a few comments, rebased. 
Checking upload repository functionality.

### Comment by @swadeley on October 22, 2024 at 10:12 AM UTC

Hi

I generated new API, reinstalled the plugin, created a repo with snapshot:

```
Out[1]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'ppc64le',
           'distribution_versions': ['8'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'Filter_MjVFomZe',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-10-22T09:42:22Z',
           'last_snapshot': {'added_counts': {'rpm.package': 1},
                             'content_counts': {'rpm.package': 1},
                             'created_at': '2024-10-22T09:42:37.971781Z',
                             'removed_counts': {},
                             'repository_name': 'Filter-MjVFomZe',
                             'repository_path': 'eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/ed094119-4e50-4336-8d32-2145ec127313',
                             'repository_uuid': 'fc37f436-78ee-4256-a0b5-28417db119bf',
                             'url': 'http://pulp-content:8000/api/pulp-content/eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/ed094119-4e50-4336-8d32-2145ec127313/',
                             'uuid': 'f232dff7-9c9c-4574-bf16-9d93f2029c52'},
           'last_snapshot_task': {'created_at': '2024-10-22T09:42:23Z',
                                  'ended_at': '2024-10-22T09:42:37Z',
                                  'error': '',
                                  'object_name': 'Filter-MjVFomZe',
                                  'object_type': 'repository',
                                  'object_uuid': 'fc37f436-78ee-4256-a0b5-28417db119bf',
                                  'org_id': '3340851',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '25a56e63-978d-460a-9c62-9e80ed2b7da0'},
           'last_snapshot_task_uuid': '25a56e63-978d-460a-9c62-9e80ed2b7da0',
           'last_snapshot_uuid': 'f232dff7-9c9c-4574-bf16-9d93f2029c52',
           'last_success_introspection_time': '2024-10-22T09:42:22Z',
           'last_update_introspection_time': '2024-10-22T09:14:19Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'Filter-MjVFomZe',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 1,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://stephenw.fedorapeople.org/multirepos/1/repo01/',
           'uuid': 'fc37f436-78ee-4256-a0b5-28417db119bf'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [2]: 
```

The new parameter `latest_snapshot_url` is there.

### Comment by @swadeley on October 22, 2024 at 10:16 AM UTC

Hi @Andrewgdewar 

Using the `last_snapshot_uuid` I get this:

```
In [3]: app.content_sources.rest_client.repositories_api.get_repo_configuration_file('f232dff7-9c9c-4574-bf16-9d93f2029c52')
2024-10-22 11:11:15.071 [    INFO] [iqe.base.rest_client] REST: GET https://ee-ebfik0w2.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/f232dff7-9c9c-4574-bf16-9d93f2029c52/config.repo with query params [] and x-rh-insights-request-id=<>
Out[3]: '[Filter_MjVFomZe]\nname=Filter-MjVFomZe\nbaseurl=http://pulp-content:8000/api/pulp-content/eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/ed094119-4e50-4336-8d32-2145ec127313/\nmodule_hotfixes=0\ngpgcheck=1\nrepo_gpgcheck=0\nenabled=1\ngpgkey=https://env-ephemeral-vyelyw.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/fc37f436-78ee-4256-a0b5-28417db119bf\nsslclientcert=/etc/pki/consumer/cert.pem\nsslclientkey=/etc/pki/consumer/key.pem\n'
```

I see `baseurl` has the same value as `latest_snapshot_url`.

Thank you

### Comment by @swadeley on October 22, 2024 at 10:16 AM UTC

/retest

### Comment by @swadeley on October 22, 2024 at 11:58 AM UTC

/retest

### Comment by @Andrewgdewar on October 22, 2024 at 07:08 PM UTC

> Hi @Andrewgdewar
> 
> Using the `last_snapshot_uuid` I get this:
> 
> ```
> In [3]: app.content_sources.rest_client.repositories_api.get_repo_configuration_file('f232dff7-9c9c-4574-bf16-9d93f2029c52')
> 2024-10-22 11:11:15.071 [    INFO] [iqe.base.rest_client] REST: GET https://ee-ebfik0w2.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/f232dff7-9c9c-4574-bf16-9d93f2029c52/config.repo with query params [] and x-rh-insights-request-id=<>
> Out[3]: '[Filter_MjVFomZe]\nname=Filter-MjVFomZe\nbaseurl=http://pulp-content:8000/api/pulp-content/eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/ed094119-4e50-4336-8d32-2145ec127313/\nmodule_hotfixes=0\ngpgcheck=1\nrepo_gpgcheck=0\nenabled=1\ngpgkey=https://env-ephemeral-vyelyw.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/fc37f436-78ee-4256-a0b5-28417db119bf\nsslclientcert=/etc/pki/consumer/cert.pem\nsslclientkey=/etc/pki/consumer/key.pem\n'
> ```
> 
> I see `baseurl` has the same value as `latest_snapshot_url`.
> 
> Thank you

If you mean the publication (the data shown from the url) yes they should be the same. 
If the urls themselves are the same that would be a bug.

### Comment by @swadeley on October 22, 2024 at 07:32 PM UTC

Hi @Andrewgdewar 

Sorry for confusion, I was trying to understand and respond to this test requirement:
'see the latest_snapshot_url represented there as "url" '

So I took `http://pulp-content:8000/api/pulp-content/eef4f7da/fc37f436-78ee-4256-a0b5-28417db119bf/latest/` and dropped the `/latest/` and saw it matched `baseurl`

In `config.repo`  there is only `baseurl` not `url`


Thank you

### Comment by @swadeley on October 23, 2024 at 11:00 AM UTC

Hi, I deployed the latest image ` pr-844-fb5a86c `

I created a new repo with snapshot:
```
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
           'last_introspection_time': '2024-10-23T10:20:19Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 6,
                                              'rpm.package': 27},
                             'content_counts': {'rpm.advisory': 6,
                                                'rpm.package': 27},
                             'created_at': '2024-10-23T10:20:35.051075Z',
                             'removed_counts': {},
                             'repository_name': 'test-repo',
                             'repository_path': 'e673e554/8a853a4b-ea97-4b28-bc96-ce8c82572182/d2c76c99-0b07-430b-a7f5-ee2e7c756b4c',
                             'repository_uuid': '8a853a4b-ea97-4b28-bc96-ce8c82572182',
                             'url': 'http://pulp-content:8000/api/pulp-content/e673e554/8a853a4b-ea97-4b28-bc96-ce8c82572182/d2c76c99-0b07-430b-a7f5-ee2e7c756b4c/',
                             'uuid': 'd2d8a6ab-c50a-4083-8739-a6635737bcb6'},
           'last_snapshot_task': {'created_at': '2024-10-23T10:20:19Z',
                                  'ended_at': '2024-10-23T10:20:35Z',
                                  'error': '',
                                  'object_name': 'test-repo',
                                  'object_type': 'repository',
                                  'object_uuid': '8a853a4b-ea97-4b28-bc96-ce8c82572182',
                                  'org_id': '3340851',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '8a316ab7-b361-4a89-8d94-b515ab61e9ea'},
           'last_snapshot_task_uuid': '8a316ab7-b361-4a89-8d94-b515ab61e9ea',
           'last_snapshot_uuid': 'd2d8a6ab-c50a-4083-8739-a6635737bcb6',
           'last_success_introspection_time': '2024-10-23T10:20:19Z',
           'last_update_introspection_time': '2024-10-23T10:20:19Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/e673e554/8a853a4b-ea97-4b28-bc96-ce8c82572182/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test-repo',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 27,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://stephenw.fedorapeople.org/fakerepos/multiple_errata/',
           'uuid': '8a853a4b-ea97-4b28-bc96-ce8c82572182'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

```


I can get repo config file with "normal" path:
```
In [1]: repodata = app.content_sources.rest_client.repositories_api.get_repo_configuration_file('d2d8a6ab-c50a-4083-8739-a6635737bcb6')
2024-10-23 11:49:35.721 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-10-23 11:49:35.721 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-10-23 11:49:36.055 [    INFO] [root] Created RESTPluginService client for https://ee-2vi7oqo0.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 with the following attributes: ['client', 'environments_api', 'features_api', 'gpg_key_api', 'package', 'packagegroups_api', 'popular_repositories_api', 'public_repositories_api', 'repositories_api', 'rpms_api', 'snapshots_api', 'tasks_api', 'templates_api']
2024-10-23 11:49:36.355 [    INFO] [iqe.base.rest_client] REST: GET https://ee-2vi7oqo0.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/d2d8a6ab-c50a-4083-8739-a6635737bcb6/config.repo with query params [] and x-rh-insights-request-id=<>

In [2]: print(repodata)
[test_repo]
name=test-repo
baseurl=http://pulp-content:8000/api/pulp-content/e673e554/8a853a4b-ea97-4b28-bc96-ce8c82572182/d2c76c99-0b07-430b-a7f5-ee2e7c756b4c/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=https://env-ephemeral-c6tulb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/8a853a4b-ea97-4b28-bc96-ce8c82572182
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem
```

but not using `get_latest_repo_configuration_file`:

```
In [3]: app.content_sources.rest_client.repositories_api.get_latest_repo_configuration_file('d2d8a6ab-c50a-4083-8739-a6635737bcb6')
<snip>
HTTP response body: {"errors":[{"status":500,"title":"Error fetching latest snapshot","detail":"record not found"}]}

```

### Comment by @rverdile on October 24, 2024 at 05:53 PM UTC

@swadeley It looks like you're using the snapshot UUID for `get_latest_repo_configuration_file('d2d8a6ab-c50a-4083-8739-a6635737bcb6')`. If you use the repository UUID hopefully that will resolve the error. 

However, a not found error should be returning a 404 instead of a 500 error, so there is still something to fix there

### Comment by @swadeley on October 25, 2024 at 10:05 AM UTC

Hi @rverdile 

you are correct

I created new repo:
```
Out[1]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test_snap',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-10-25T08:55:27Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 4,
                                              'rpm.package': 32},
                             'content_counts': {'rpm.advisory': 4,
                                                'rpm.package': 32},
                             'created_at': '2024-10-25T08:55:42.563075Z',
                             'removed_counts': {},
                             'repository_name': 'test-snap',
                             'repository_path': '8ed19f2c/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/e8077abc-2511-462f-8003-15f1934f08ab',
                             'repository_uuid': '7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4',
                             'url': 'http://pulp-content:8000/api/pulp-content/8ed19f2c/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/e8077abc-2511-462f-8003-15f1934f08ab/',
                             'uuid': '3fe393f3-3d9d-4140-8a46-320c32c15dd3'},
           'last_snapshot_task': {'created_at': '2024-10-25T08:55:27Z',
                                  'ended_at': '2024-10-25T08:55:42Z',
                                  'error': '',
                                  'object_name': 'test-snap',
                                  'object_type': 'repository',
                                  'object_uuid': '7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4',
                                  'org_id': '3340851',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '9c7ac932-4f8d-44ae-be56-81a1cceeb408'},
           'last_snapshot_task_uuid': '9c7ac932-4f8d-44ae-be56-81a1cceeb408',
           'last_snapshot_uuid': '3fe393f3-3d9d-4140-8a46-320c32c15dd3',
           'last_success_introspection_time': '2024-10-25T08:55:27Z',
           'last_update_introspection_time': '2024-10-25T08:55:27Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/8ed19f2c/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test-snap',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 32,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/',
           'uuid': '7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&origin=external'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

With last_snapshot_uuid:
```
In [4]: print(app.content_sources.rest_client.repositories_api.get_repo_configuration_file('3fe393f3-3d9d-4140-8a46-320c32c15dd3'))
2024-10-25 10:58:30.276 [    INFO] [iqe.base.rest_client] REST: GET https://ee-631gze7a.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/3fe393f3-3d9d-4140-8a46-320c32c15dd3/config.repo with query params [] and x-rh-insights-request-id=--4047--
[test_snap]
name=test-snap
baseurl=http://pulp-content:8000/api/pulp-content/8ed19f2c/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/e8077abc-2511-462f-8003-15f1934f08ab/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=https://env-ephemeral-wlyio8.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem
```

With repo uuid:
```
In [5]: print(app.content_sources.rest_client.repositories_api.get_latest_repo_configuration_file('7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4'))
2024-10-25 11:00:20.771 [    INFO] [iqe.base.rest_client] REST: GET https://ee-631gze7a.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/config.repo with query params [] and x-rh-insights-request-id=--4974--
[test_snap]
name=test-snap
baseurl=http://pulp-content:8000/api/pulp-content/8ed19f2c/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4/latest/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=https://env-ephemeral-wlyio8.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/7b528bbf-9137-4e8a-a0d0-3f9b64b8a7d4
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem

```





### Comment by @Andrewgdewar on October 30, 2024 at 01:49 PM UTC

> @swadeley It looks like you're using the snapshot UUID for `get_latest_repo_configuration_file('d2d8a6ab-c50a-4083-8739-a6635737bcb6')`. If you use the repository UUID hopefully that will resolve the error.
> 
> However, a not found error should be returning a 404 instead of a 500 error, so there is still something to fix there

Looking for this now!

### Comment by @Andrewgdewar on October 30, 2024 at 03:48 PM UTC

Okay this is a 500:
![Screenshot 2024-10-30 at 9 47 59 AM](https://github.com/user-attachments/assets/f5ccc392-b24c-415a-b137-57c8930aa697)

This is a 500:
![Screenshot 2024-10-30 at 9 44 52 AM](https://github.com/user-attachments/assets/49c126f4-b079-4668-bb1e-d238c27c40cf)

And this is a 500:
![Screenshot 2024-10-30 at 9 50 32 AM](https://github.com/user-attachments/assets/5d29e69e-f3ef-4db6-9a03-ad04f2c33633)


🤔



---

## Reviews

### Review by @jlsherrill - Commented on October 15, 2024 at 11:24 AM UTC

### Review by @Andrewgdewar - Commented on October 15, 2024 at 03:24 PM UTC

### Review by @jlsherrill - Commented on October 15, 2024 at 03:25 PM UTC

### Review by @swadeley - Commented on October 16, 2024 at 10:05 AM UTC

### Review by @rverdile - Commented on October 24, 2024 at 06:01 PM UTC

### Review by @Andrewgdewar - Commented on October 28, 2024 at 04:46 PM UTC

### Review by @Andrewgdewar - Commented on October 29, 2024 at 08:54 PM UTC

### Review by @rverdile - Commented on October 29, 2024 at 10:31 PM UTC

### Review by @Andrewgdewar - Commented on October 30, 2024 at 01:49 PM UTC

### Review by @rverdile - Commented on October 31, 2024 at 03:10 PM UTC

this lgtm, but @jlsherrill was reviewing initially so he might want to take another look

### Review by @jlsherrill - Approved on November 01, 2024 at 04:11 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/844*
