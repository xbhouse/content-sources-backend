---
type: pull_request
number: 590
title: "Fixes 2940: add snapshot rpm list api"
state: merged
author: jlsherrill
created: 2024-02-28T20:51:24Z
updated: 2024-03-13T19:00:33Z
closed: 2024-03-13T18:36:43Z
merged: 2024-03-13T18:36:43Z
base_branch: main
head_branch: 2940
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/590
---

# Pull Request #590: Fixes 2940: add snapshot rpm list api

**Author**: @jlsherrill
**Created**: February 28, 2024 at 08:51 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2940`

## Description

## Summary
adds a new snapshot rpm list api:

```// @Router       /repositories/{uuid}/rpms [get]```

## Testing steps

Create a repo and let it snapshot, then fetch the rpms:

```
GET http://localhost:8000/api/content-sources/v1.0/snapshots/0e6c9e7f-e021-4710-8021-f59cb9c56d82/rpms?search=foo&limit=1
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 28, 2024 at 08:57 PM UTC

requires https://github.com/content-services/tang/pull/5 and an updated tang

You can test locally by checking out that tang pr, and adjusting go.mod to point to your local checkout

### Comment by @jlsherrill on February 28, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-2940

### Comment by @swadeley on March 06, 2024 at 02:47 PM UTC

@jlsherrill Please rebase

### Comment by @jlsherrill on March 07, 2024 at 05:33 PM UTC

updated, rebased, and set to use the new tang

### Comment by @swadeley on March 08, 2024 at 07:06 PM UTC

Hi

the title has the word "list" so when I saw ` /repositories/{uuid}/rpms `  I thought this was the new route:
`app.content_sources.rest_client.rpms_api.list_repositories_rpms`
and made the test below.

but when I look at the example with `api/content-sources/v1.0/snapshots/` I think it might be:
`app.content_sources.rest_client.rpms_api.search_snapshot_rpms`  (search not list).

```In [4]: app.content_sources.rest_client.rpms_api.list_repositories_rpms.params_map
Out[4]: 
{'all': ['uuid',
  'limit',
  'offset',
  'search',
  'sort_by',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['uuid'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}

In [5]: created_repo['uuid']
Out[5]: '45c35a10-420c-4507-92c1-b762b5635a6f'

In [6]: app.content_sources.rest_client.rpms_api.list_repositories_rpms('45c35a10-420c-4507-92c1-b762b5635a6f')
2024-03-08 19:01:26.677 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[6]: 
{'data': [{'arch': 'noarch',
           'checksum': '7a831f9f90bf4d21027572cb503d20b702de8e8785b02c0397445c2e481d81b3',
           'epoch': 0,
           'name': 'bear',
           'release': '1',
           'summary': 'A dummy package of bear',
           'uuid': 'ba75ff00-4785-4347-82a8-ef17dde8397a',
           'version': '4.1'},
```

```
In [9]: app.content_sources.rest_client.rpms_api.list_repositories_rpms('45c35a10-420c-4507-92c1-b762b5635a6f', limit=2, offset=2)
2024-03-08 19:03:36.129 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 2)]
Out[9]: 
{'data': [{'arch': 'noarch',
           'checksum': '43e77adb7f51b5542b813024a8ee3e477175c142f35982ab5ae2b2978486239f',
           'epoch': 0,
           'name': 'cat',
           'release': '1',
           'summary': 'A dummy package of cat',
           'uuid': '57d987f6-b369-4076-90aa-1000aa0bb924',
           'version': '1.0'},
          {'arch': 'noarch',
           'checksum': '2189ac4bf059f98a8d48136eb72b46415f3aa263026443ebd8879d415eacaf42',
           'epoch': 0,
           'name': 'cheetah',
           'release': '5',
           'summary': 'A dummy package of cheetah',
           'uuid': 'eb869ce3-a17f-4625-8442-46ee0ba696d0',
           'version': '1.25.3'}],
 'links': {'first': '/api/content-sources/v1/repositories/45c35a10-420c-4507-92c1-b762b5635a6f/rpms?limit=2&offset=0',
           'last': '/api/content-sources/v1/repositories/45c35a10-420c-4507-92c1-b762b5635a6f/rpms?limit=2&offset=30',
           'next': '/api/content-sources/v1/repositories/45c35a10-420c-4507-92c1-b762b5635a6f/rpms?limit=2&offset=4',
           'prev': '/api/content-sources/v1/repositories/45c35a10-420c-4507-92c1-b762b5635a6f/rpms?limit=2&offset=0'},
 'meta': {'count': 32, 'limit': 2, 'offset': 2}}

In [10]: 
```



### Comment by @swadeley on March 08, 2024 at 07:36 PM UTC

Hows this:

```
In [18]: app.content_sources.rest_client.repositories_api.get_repository('45c35a10-420c-4507-92c1-b762b5635a6f')['last_snapshot']['uuid']
2024-03-08 19:26:50.680 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[18]: '4f919c63-e47f-470a-9578-8828a89eecda'
In [19]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['4f919c63-e47f-470a-9578-8828a89eecda']))
2024-03-08 19:32:16.327 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[19]: 
[{'package_name': 'bear', 'summary': 'A dummy package of bear'},
 {'package_name': 'camel', 'summary': 'A dummy package of camel'},
 {'package_name': 'cat', 'summary': 'A dummy package of cat'},
 {'package_name': 'cheetah', 'summary': 'A dummy package of cheetah'},
 {'package_name': 'chimpanzee', 'summary': 'A dummy package of chimpanzee'},
 {'package_name': 'cockateel', 'summary': 'A dummy package of cockateel'},
 {'package_name': 'cow', 'summary': 'A dummy package of cow'},
 {'package_name': 'crow', 'summary': 'A dummy package of crow'},
 {'package_name': 'dog', 'summary': 'A dummy package of dog'},
 {'package_name': 'dolphin', 'summary': 'A dummy package of dolphin'},
 {'package_name': 'duck', 'summary': 'A dummy package of duck'},
 {'package_name': 'elephant', 'summary': 'A dummy package of elephant'},
 {'package_name': 'fox', 'summary': 'A dummy package of fox'},
 {'package_name': 'frog', 'summary': 'A dummy package of frog'},
 {'package_name': 'giraffe', 'summary': 'A dummy package of giraffe'},
 {'package_name': 'gorilla', 'summary': 'A dummy package of gorilla'},
 {'package_name': 'horse', 'summary': 'A dummy package of horse'},
 {'package_name': 'kangaroo', 'summary': 'A dummy package of kangaroo'},
 {'package_name': 'lion', 'summary': 'A dummy package of lion'},
 {'package_name': 'mouse', 'summary': 'A dummy package of mouse'},
 {'package_name': 'penguin', 'summary': 'A dummy package of penguin'},
 {'package_name': 'pike', 'summary': 'A dummy package of pike'},
 {'package_name': 'shark', 'summary': 'A dummy package of shark'},
 {'package_name': 'squirrel', 'summary': 'A dummy package of squirrel'},
 {'package_name': 'stork', 'summary': 'A dummy package of stork'},
 {'package_name': 'tiger', 'summary': 'A dummy package of tiger'},
 {'package_name': 'trout', 'summary': 'A dummy package of trout'},
 {'package_name': 'walrus', 'summary': 'A dummy package of walrus'},
 {'package_name': 'whale', 'summary': 'A dummy package of whale'},
 {'package_name': 'wolf', 'summary': 'A dummy package of wolf'},
 {'package_name': 'zebra', 'summary': 'A dummy package of zebra'}]

```

### Comment by @swadeley on March 11, 2024 at 09:24 AM UTC

Hi, in master branch I can already list rpms and search for one in the list for a repo in stage:
```

In [8]: app.content_sources.rest_client.rpms_api.list_repositories_rpms('07a581c4-e586-4eee-b853-042f77301bdc', search='acme-package')
2024-03-11 09:20:11.941 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, 892fcf8fad22485cb62b4d217953ad80, params=[('search', 'acme-package')]
Out[8]: 
{'data': [{'arch': 'noarch',
           'checksum': '2f7db1e01ccb0e0c0c4a6ef4981e7398927e5e558671c5989c6b667a8c702e56',
           'epoch': 0,
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'uuid': '258a5f0a-10f6-40f8-8350-188218ee08d5',
           'version': '1.0.2'},
          {'arch': 'noarch',
           'checksum': 'da1bdb74dc9687adfa7e1a85026c2373ed8d91c03639446d9b865c3bf3e9c67d',
           'epoch': 0,
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'uuid': '603d140d-5093-4f89-a61c-ad9d5605c67a',
           'version': '1.1.2'},
          {'arch': 'noarch',
           'checksum': 'e5fae88537edadbda567359686451986b7947fa26b7a0bdc68fbf0373dcecbd1',
           'epoch': 0,
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'uuid': 'de386858-310d-4611-b13c-179977319cb0',
           'version': '1.0.1'}],
 'links': {'first': '/api/content-sources/v1/repositories/07a581c4-e586-4eee-b853-042f77301bdc/rpms?limit=100&offset=0&search=acme-package',
           'last': '/api/content-sources/v1/repositories/07a581c4-e586-4eee-b853-042f77301bdc/rpms?limit=100&offset=0&search=acme-package'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}
```


### Comment by @swadeley on March 11, 2024 at 09:39 AM UTC

Hi, on master branch, I reinstalled the plugin, opened shell to stage, created repo, got snapshot UUID:

`   'last_snapshot_uuid': 'c2eedc6f-d57f-46bd-a933-89b351de6f3`

Ran this:
In [2]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['c2eedc6f-d57f-46bd-a933-89b351de6f31']))

**ServiceException: (500)
Reason: Internal Server Error**

So search_snapshot_rpms is the new path, **search** not **list**. OK?

### Comment by @swadeley on March 11, 2024 at 10:04 AM UTC

back with shell to ephemeral venv, on the PR branch, plugin reinstalled.

More with search and limit.

I used this repo:
`https://mmccune.fedorapeople.org/repos/small/`
which has three acme packages but the snapshot only has one, so testing with **limit** is limited in this case:
```

In [5]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['417eaa78-d12c-488e-a1f8-12f5dc359168']))
2024-03-11 09:54:54.122 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[5]: 
[{'package_name': 'acme-package', 'summary': 'acme-package package'},
 {'package_name': 'beautifulness', 'summary': 'beautifulness package'},
 {'package_name': 'diphtheric', 'summary': 'diphtheric package'},
 {'package_name': 'frugalist', 'summary': 'frugalist package'},
 {'package_name': 'maleability', 'summary': 'maleability package'}]

In [6]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['417eaa78-d12c-488e-a1f8-12f5dc359168'], search='frugalist'))
2024-03-11 09:56:41.442 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[6]: [{'package_name': 'frugalist', 'summary': 'frugalist package'}]

In [7]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['417eaa78-d12c-488e-a1f8-12f5dc359168'], search='acme-package'))
2024-03-11 09:59:26.052 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[7]: [{'package_name': 'acme-package', 'summary': 'acme-package package'}]


In [9]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['417eaa78-d12c-488e-a1f8-12f5dc359168'], search='acme-package', limit=1))
2024-03-11 10:01:17.090 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[9]: [{'package_name': 'acme-package', 'summary': 'acme-package package'}]
```


### Comment by @jlsherrill on March 12, 2024 at 12:44 PM UTC

search is an existing api for searching for rpm names, this api is 'list', i'd expect it to be called list_snapshot_rpms or something similar

### Comment by @swadeley on March 12, 2024 at 01:23 PM UTC

/retest

### Comment by @swadeley on March 12, 2024 at 07:57 PM UTC

/retest

### Comment by @swadeley on March 12, 2024 at 08:44 PM UTC

> list_snapshot_rpms

Hi, looking in my plugin MR 342, I found the definition `def __list_snapshot_rpms`  in the `iqe_content_sources_api/api/rpms_api.py` file and in the ``iqe_content_sources_api/api/repositories_api.py`` file.

I'll tag you there.

But when I check, in a shell, both paths (to match file names):
app.content_sources.rest_client.repositories_api
 app.content_sources.rest_client.rpms_api

I don't see `list_snapshot_rpms`  only `list_repositories_rpms`.

Maybe something is missing in API docs? I can try generate them again.

### Comment by @jlsherrill on March 13, 2024 at 01:29 PM UTC

I didn't realize it would be duplicated into each 'tag', so i've updated here to only use the snapshot tag.  IDK if that will make a difference, but it really doesn't need to be in 3 places in the generated client.  Maybe try regenerating and see if it shows up in the snapshot_api?

### Comment by @swadeley on March 13, 2024 at 04:00 PM UTC

Hi @jlsherrill 

Thanks for the update. I can now find the full path, but does not work:

```
In [3]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms.params_map
Out[3]: 
{'all': ['uuid',
  'limit',
  'offset',
  'search',
  'async_req',
  '_host_index',
  '_preload_content',
  '_request_timeout',
  '_return_http_data_only',
  '_check_input_type',
  '_check_return_type'],
 'required': ['uuid'],
 'nullable': ['_request_timeout'],
 'enum': [],
 'validation': []}
```

```
In [10]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms('b342c4f3-c0bd-4eb4-9b2d-aad494928d2d', limit=2)
ApiTypeError: Invalid type for variable 'epoch'. Required value type is int and passed type was str at ['received_data']['data'][0]['epoch']


```

```
In [4]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms('b342c4f3-c0bd-4eb4-9b2d-aad494928d2d')
ApiTypeError: Invalid type for variable 'epoch'. Required value type is int and passed type was str at ['received_data']['data'][0]['epoch']

```



### Comment by @swadeley on March 13, 2024 at 06:36 PM UTC

Hi, all working now:


`In [2]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms('b342c4f3-c0bd-4eb4-9b2d-aad494928d2d', limit=4)
2024-03-13 18:34:30.499 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 4)]
Out[2]: 
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
           'version': '1.25.3'}],
 'links': {'first': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=0',
           'last': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=32',
           'next': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=4'},
 'meta': {'count': 35, 'limit': 4, 'offset': 0}}

In [3]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms('b342c4f3-c0bd-4eb4-9b2d-aad494928d2d', search='cat')
2024-03-13 18:34:50.674 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('search', 'cat')]
Out[3]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'cat',
           'release': '1',
           'summary': 'A dummy package of cat',
           'version': '1.0'}],
 'links': {'first': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=100&offset=0&search=cat',
           'last': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=100&offset=0&search=cat'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [4]: app.content_sources.rest_client.snapshots_api.list_snapshot_rpms('b342c4f3-c0bd-4eb4-9b2d-aad494928d2d', limit=4, offset=2)
2024-03-13 18:35:03.202 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 4), ('offset', 2)]
Out[4]: 
{'data': [{'arch': 'noarch',
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
           'version': '3.1'}],
 'links': {'first': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=0',
           'last': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=32',
           'next': '/api/content-sources/v1/snapshots/b342c4f3-c0bd-4eb4-9b2d-aad494928d2d/rpms?limit=4&offset=6'},
 'meta': {'count': 35, 'limit': 4, 'offset': 2}}

In [5]: 

`

Thank you

---

## Reviews

### Review by @rverdile - Commented on March 06, 2024 at 08:23 PM UTC

I was testing this alongside the Tang PR. Should be good once it's rebased and updated with the tang release.

### Review by @rverdile - Commented on March 06, 2024 at 08:24 PM UTC

### Review by @rverdile - Commented on March 06, 2024 at 08:33 PM UTC

### Review by @rverdile - Approved on March 07, 2024 at 08:41 PM UTC

nice!

### Review by @rverdile - Commented on March 13, 2024 at 05:30 PM UTC

### Review by @jlsherrill - Commented on March 13, 2024 at 05:44 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/590*
