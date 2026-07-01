---
type: pull_request
number: 414
title: "Fixes 2615: add config.repo endpoint"
state: merged
author: rverdile
created: 2023-10-03T14:08:31Z
updated: 2023-10-18T19:00:28Z
closed: 2023-10-18T18:57:50Z
merged: 2023-10-18T18:57:50Z
base_branch: main
head_branch: repo-config
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/414
---

# Pull Request #414: Fixes 2615: add config.repo endpoint

**Author**: @rverdile
**Created**: October 03, 2023 at 02:08 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `repo-config`

## Description

## Summary
- Adds endpoint, `repositories/:repo_uuid/snapshots/:snapshot_uuid`/config.repo that returns config.repo file for a particular snapshot.
- Adds snapshot UUID and content URL to snapshot response.
- Adds pulp content path base to cache after first query to pulp.

## Testing steps
1. Create a repository with a snapshot
2. Once the snapshot is created, list the snapshots for that repo to get the UUID of the snapshot.
3. GET the new endpoint.
4. It should return plain text containing the config.repo as described in the JIRA card.


---

## Discussion

### Comment by @jlsherrill on October 04, 2023 at 01:41 PM UTC

https://issues.redhat.com/browse/HMS-2615

### Comment by @jlsherrill on October 06, 2023 at 09:53 PM UTC

when you fetch a repository (or list), you also get the last_snapshot, but url is empty string.  I guess we'd want to add that there to?

### Comment by @swadeley on October 13, 2023 at 10:34 AM UTC

@rverdile Rebase please.

### Comment by @jlsherrill on October 17, 2023 at 05:25 PM UTC

/retest

### Comment by @swadeley on October 18, 2023 at 08:00 AM UTC

Hi, we can see the new 'last_snapshot'  'uuid' and  'url':

```
In [9]: app.content_sources.rest_client.repositories_api.list_repositories(search='fedorapeople.org')['data'][0]['last_snapshot']
<snip>
Out[9]: 
{'added_counts': {'rpm.package': 1},
 'content_counts': {'rpm.package': 1},
 'created_at': '2023-10-17T19:35:30.865733Z',
 'removed_counts': {},
 'repository_path': '',
 'url': 'http://cs-pulp-content-svc.local/pulp/content//',
 'uuid': '680b4d3f-284e-4da7-a86c-f6e65cb7fb1a'}
```

In the  'last_snapshot'  'url' I was expecting to see path to repo version.

### Comment by @swadeley on October 18, 2023 at 09:41 AM UTC

Here is the path to repo version:
```
In [13]: app.content_sources.rest_client.repositories_api.get_repository('d6d49d52-2b64-4e40-a78d-adf4191979b1')['last_snapshot']['url']
<snip>
Out[13]: 'http://cs-pulp-content-svc.local/pulp/content/c7a88bc0/d6d49d52-2b64-4e40-a78d-adf4191979b1/b19ca00a-7efe-4951-a2ef-15d1b7a8f731/'
```

Using repo 'uuid' and 'snapshot_uuid' we can get the config file
```
In [15]: app.content_sources.rest_client.repositories_api.get_repo_configuration_file('d6d49d52-2b64-4e40-a78d-adf4191979b1', '680b4d3f-284e-4da7-a86c-f6e65cb7fb1a')
<snip>
Out[15]: '[fedoratest1]\nname=fedoratest1\nbaseurl=http://cs-pulp-content-svc.local/pulp/content/c7a88bc0/d6d49d52-2b64-4e40-a78d-adf4191979b1/b19ca00a-7efe-4951-a2ef-15d1b7a8f731/\ngpgcheck=0\nrepo_gpgcheck=0\nenabled=1\n'
```



### Comment by @jlsherrill on October 18, 2023 at 12:09 PM UTC

filed https://github.com/pulp/pulp-clowder-deployments/issues/94  for the incorrect pulp content_origin

### Comment by @rverdile on October 18, 2023 at 02:03 PM UTC

> In the 'last_snapshot' 'url' I was expecting to see path to repo version.

this is fixed now

### Comment by @swadeley on October 18, 2023 at 06:49 PM UTC

Hi, I see fully formed url now:
```
In [3]: app.content_sources.rest_client.repositories_api.list_repositories(search='fedorapeople.org')['data'][0]['last_snapshot']
<snip>
Out[3]: 
{'added_counts': {'rpm.package': 1},
 'content_counts': {'rpm.package': 1},
 'created_at': '2023-10-18T18:36:28.277468Z',
 'removed_counts': {},
 'repository_path': 'a7b352c3/4730399f-c570-48fb-a85d-9a8fe565a0f5/141a1ec0-51df-4ea2-b3af-bbe1ea4b36df',
 'url': 'https://env-ephemeral-4n8ce1.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/pulp/content/a7b352c3/4730399f-c570-48fb-a85d-9a8fe565a0f5/141a1ec0-51df-4ea2-b3af-bbe1ea4b36df/',
 'uuid': '3cfc86fd-18ed-40a0-aa11-7a7a6895eb4c'}


```

and  'repository_path':  is now showing a value

### Comment by @swadeley on October 18, 2023 at 06:57 PM UTC

['last_snapshot']['url'] looks good:
```
In [6]: app.content_sources.rest_client.repositories_api.get_repository('4730399f-c570-48fb-a85d-9a8fe565a0f5')['last_snapshot']['url']
<snip>
Out[6]: 'https://env-ephemeral-4n8ce1.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/pulp/content/a7b352c3/4730399f-c570-48fb-a85d-9a8fe565a0f5/141a1ec0-51df-4ea2-b3af-bbe1ea4b36df/'
```

```
In [8]: app.content_sources.rest_client.repositories_api.get_repo_configuration_file('4730399f-c570-48fb-a85d-9a8fe565a0f5', '3cfc86fd-18ed-40a0-aa11-7a7a6895eb4c')
<snip>
Out[8]: '[fedoratest1]\nname=fedoratest1\nbaseurl=https://env-ephemeral-4n8ce1.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/pulp/content/a7b352c3/4730399f-c570-48fb-a85d-9a8fe565a0f5/141a1ec0-51df-4ea2-b3af-bbe1ea4b36df/\ngpgcheck=0\nrepo_gpgcheck=0\nenabled=1\n'
```


### Comment by @swadeley on October 18, 2023 at 06:57 PM UTC

Thank you

---

## Reviews

### Review by @jlsherrill - Commented on October 03, 2023 at 06:18 PM UTC

### Review by @rverdile - Commented on October 03, 2023 at 07:55 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 08:04 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 08:04 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 08:06 PM UTC

### Review by @rverdile - Commented on October 05, 2023 at 03:51 PM UTC

### Review by @jlsherrill - Commented on October 05, 2023 at 05:32 PM UTC

### Review by @jlsherrill - Commented on October 06, 2023 at 09:38 PM UTC

### Review by @jlsherrill - Commented on October 06, 2023 at 09:40 PM UTC

### Review by @swadeley - Commented on October 09, 2023 at 01:25 PM UTC

I think we don't need capitals for repo and snapshot. 

### Review by @jlsherrill - Approved on October 17, 2023 at 03:50 PM UTC

### Review by @rverdile - Commented on October 18, 2023 at 03:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/414*
