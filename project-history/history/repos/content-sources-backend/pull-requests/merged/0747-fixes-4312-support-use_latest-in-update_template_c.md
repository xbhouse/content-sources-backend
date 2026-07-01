---
type: pull_request
number: 747
title: "Fixes 4312: support use_latest in update_template_content"
state: merged
author: rverdile
created: 2024-07-17T13:17:32Z
updated: 2024-07-26T14:00:25Z
closed: 2024-07-26T13:39:12Z
merged: 2024-07-26T13:39:12Z
base_branch: main
head_branch: use_latest_fetch
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/747
---

# Pull Request #747: Fixes 4312: support use_latest in update_template_content

**Author**: @rverdile
**Created**: July 17, 2024 at 01:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `use_latest_fetch`

## Description

## Summary
After template creation, supports fetching the latest snapshot for each of the repositories in the template, if the use_latest parameter is set.

## Testing steps
1. Create a repository with multiple snapshots by changing the URL after creating the first snapshot
2. Add that repository to a template (should work on create and update for the template)
3. If you navigate to where the distribution servers the content, should be `pulp.content:8080/pulp/content/<domain>/templates/<template_uuid>` locally, you should see only the content for the most recent snapshot.
4. A call to `GET templates/:uuid/rpms` should return only the content of the most recent snapshot

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 17, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4312

### Comment by @xbhouse on July 23, 2024 at 06:31 PM UTC

just the one question, otherwise looks and works great! 

### Comment by @mayurilahane on July 24, 2024 at 02:38 PM UTC

/retest

### Comment by @mayurilahane on July 25, 2024 at 09:03 PM UTC

Hi,

So i created a repo with 2 snapshots (33) and (32) packages 

and the latest one is with the (32) packages 

```
In [31]:  app.content_sources.rest_client.repositories_api.list_repositories()
2024-07-25 16:49:36.412 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=c554018e-d6fc-4a19-97ec-2d26c217e4fd, params=[]
Out[31]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'Test_package_count_first_snapshot',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-07-25T20:21:07Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 4,
                                              'rpm.package': 32},
                             'content_counts': {'rpm.advisory': 4,
                                                'rpm.package': 32},
                             'created_at': '2024-07-25T20:21:21.034074Z',
                             'removed_counts': {'rpm.package': 33},
                             'repository_path': '15038db9/5c7aa33b-00aa-4c68-96a7-bae76a8de1e2/845af66c-b5f3-493c-882a-ec27cef1e42e',
                             'url': 'http://pulp-content-svc:24816/api/pulp-content/15038db9/5c7aa33b-00aa-4c68-96a7-bae76a8de1e2/845af66c-b5f3-493c-882a-ec27cef1e42e/',
                             'uuid': 'aa58c3ce-e3df-4e48-a0dd-20e180f4667f'},
           'last_snapshot_task': {'created_at': '2024-07-25T20:21:07Z',
                                  'ended_at': '2024-07-25T20:21:21Z',
                                  'error': '',
                                  'org_id': '3340851',
                                  'repository_name': 'Test package count '
                                                     'second snapshot',
                                  'repository_uuid': '5c7aa33b-00aa-4c68-96a7-bae76a8de1e2',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '65b44ce5-f86b-4185-a4a8-6d1999373907'},
           'last_snapshot_task_uuid': '65b44ce5-f86b-4185-a4a8-6d1999373907',
           'last_snapshot_uuid': 'aa58c3ce-e3df-4e48-a0dd-20e180f4667f',
           'last_success_introspection_time': '2024-07-25T20:21:07Z',
           'last_update_introspection_time': '2024-07-25T20:21:07Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'Test package count second snapshot',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 32,
           'snapshot': True,
           'status': 'Valid',
           'url': 'http://jlsherrill.fedorapeople.org/fake-repos/needed-errata/',
           'uuid': '5c7aa33b-00aa-4c68-96a7-bae76a8de1e2'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

### Comment by @mayurilahane on July 25, 2024 at 09:05 PM UTC

Created a template with the repo by setting use_latest parameter to "True"

```
In [32]: app.content_sources.rest_client.templates_api.create_template(dict(name="edited_repo_template", version="8", arch="x86_64", repository_uuids=["5c7aa33b-00aa
    ...: -4c68-96a7-bae76a8de1e2"], use_latest=True))
2024-07-25 16:54:08.561 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=37d9e701-9b24-42c8-9fc4-075685404a0d, params=[]
Out[32]: 
{'arch': 'x86_64',
 'created_at': '2024-07-25T20:54:08.541720202Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'edited_repo_template',
 'org_id': '3340851',
 'repository_uuids': ['5c7aa33b-00aa-4c68-96a7-bae76a8de1e2'],
 'rhsm_environment_id': 'a1a0a6aa0a714b1682fc7bc123c6fafe',
 'updated_at': '2024-07-25T20:54:08.541720202Z',
 'use_latest': True,
 'uuid': 'a1a0a6aa-0a71-4b16-82fc-7bc123c6fafe',
 'version': '8'}
```

### Comment by @mayurilahane on July 25, 2024 at 09:08 PM UTC

It can successfully fetch the latest snapshot for the repository in the template.

```

In [35]:  app.content_sources.rest_client.rpms_api.list_template_rpms("a1a0a6aa-0a71-4b16-82fc-7bc123c6fafe")
2024-07-25 17:01:12.133 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=a17a171f-f674-4257-92a2-a1c5214c8745, params=[]
Out[35]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'bear',
           'release': '1',
           'summary': 'A dummy package of bear',
           'version': '4.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'camel',
           'release': '1',
           'summary': 'A dummy package of camel',
           'version': '0.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'cat',
           'release': '1',
           'summary': 'A dummy package of cat',
           'version': '1.0'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'cheetah',
           'release': '5',
           'summary': 'A dummy package of cheetah',
           'version': '1.25.3'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'chimpanzee',
           'release': '1',
           'summary': 'A dummy package of chimpanzee',
           'version': '0.21'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'cockateel',
           'release': '1',
           'summary': 'A dummy package of cockateel',
           'version': '3.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'cow',
           'release': '3',
           'summary': 'A dummy package of cow',
           'version': '2.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'crow',
           'release': '1',
           'summary': 'A dummy package of crow',
           'version': '0.8'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'dog',
           'release': '1',
           'summary': 'A dummy package of dog',
           'version': '4.23'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'dolphin',
           'release': '1',
           'summary': 'A dummy package of dolphin',
           'version': '3.10.232'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'duck',
           'release': '1',
           'summary': 'A dummy package of duck',
           'version': '0.6'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'elephant',
           'release': '1',
           'summary': 'A dummy package of elephant',
           'version': '8.3'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'fox',
           'release': '2',
           'summary': 'A dummy package of fox',
           'version': '1.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'frog',
           'release': '1',
           'summary': 'A dummy package of frog',
           'version': '0.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'giraffe',
           'release': '2',
           'summary': 'A dummy package of giraffe',
           'version': '0.67'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'gorilla',
           'release': '1',
           'summary': 'A dummy package of gorilla',
           'version': '0.62'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'horse',
           'release': '2',
           'summary': 'A dummy package of horse',
           'version': '0.22'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'kangaroo',
           'release': '1',
           'summary': 'A dummy package of kangaroo',
           'version': '0.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'lion',
           'release': '1',
           'summary': 'A dummy package of lion',
           'version': '0.4'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'mouse',
           'release': '1',
           'summary': 'A dummy package of mouse',
           'version': '0.1.12'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'penguin',
           'release': '1',
           'summary': 'A dummy package of penguin',
           'version': '0.9.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'pike',
           'release': '1',
           'summary': 'A dummy package of pike',
           'version': '2.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'shark',
           'release': '1',
           'summary': 'A dummy package of shark',
           'version': '0.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'squirrel',
           'release': '1',
           'summary': 'A dummy package of squirrel',
           'version': '0.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'stork',
           'release': '2',
           'summary': 'A dummy package of stork',
           'version': '0.12'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'tiger',
           'release': '4',
           'summary': 'A dummy package of tiger',
           'version': '1.0'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'trout',
           'release': '1',
           'summary': 'A dummy package of trout',
           'version': '0.12'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'walrus',
           'release': '1',
           'summary': 'A dummy package of walrus',
           'version': '0.71'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'walrus',
           'release': '1',
           'summary': 'A dummy package of walrus',
           'version': '5.21'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'whale',
           'release': '1',
           'summary': 'A dummy package of whale',
           'version': '0.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'wolf',
           'release': '2',
           'summary': 'A dummy package of wolf',
           'version': '9.4'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'zebra',
           'release': '2',
           'summary': 'A dummy package of zebra',
           'version': '0.1'}],
 'links': {'first': '/api/content-sources/v1/templates/a1a0a6aa-0a71-4b16-82fc-7bc123c6fafe/rpms?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/a1a0a6aa-0a71-4b16-82fc-7bc123c6fafe/rpms?limit=100&offset=0'},
 'meta': {'count': 32, 'limit': 100, 'offset': 0}}
```


### Comment by @mayurilahane on July 25, 2024 at 09:09 PM UTC

/lgtm 

---

## Reviews

### Review by @xbhouse - Commented on July 23, 2024 at 06:26 PM UTC

### Review by @rverdile - Commented on July 23, 2024 at 06:33 PM UTC

### Review by @rverdile - Commented on July 23, 2024 at 07:27 PM UTC

### Review by @xbhouse - Approved on July 23, 2024 at 07:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/747*
