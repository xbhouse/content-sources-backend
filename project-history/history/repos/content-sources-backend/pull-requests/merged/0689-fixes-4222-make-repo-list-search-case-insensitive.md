---
type: pull_request
number: 689
title: "Fixes 4222: make repo list search case insensitive"
state: merged
author: xbhouse
created: 2024-06-04T16:07:13Z
updated: 2024-06-07T15:00:26Z
closed: 2024-06-07T14:44:46Z
merged: 2024-06-07T14:44:46Z
base_branch: main
head_branch: 4222
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/689
---

# Pull Request #689: Fixes 4222: make repo list search case insensitive

**Author**: @xbhouse
**Created**: June 04, 2024 at 04:07 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4222`

## Description

## Summary

Changes repo list search query to be case insensitive

## Testing steps

- Add a repository
- Make a request to the list repos endpoint with a search term (i.e. `/repositories/?search=everything`)
- Should return repositories with that term regardless of case
- Can also check this in the UI with the Name/URL filter

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 04, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-4222

### Comment by @swadeley on June 06, 2024 at 06:26 PM UTC

/retest

### Comment by @swadeley on June 06, 2024 at 08:13 PM UTC

Hi

I created test repo called `test_repo_name_1`

In UI I can use capital letters to search for it:
![image](https://github.com/content-services/content-sources-backend/assets/11247237/856cce23-4f60-492b-9560-fe394cd93905)


also in shell with  API:

```
In [9]: app.content_sources.rest_client.repositories_api.list_repositories(search='TEST')['data'][0]['name']
2024-06-06 21:11:04.517 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=4<>, params=[('search', 'TEST')]
Out[9]: 'test_repo_name_1'
```

### Comment by @swadeley on June 06, 2024 at 08:16 PM UTC

Works with URL too:

![image](https://github.com/content-services/content-sources-backend/assets/11247237/4fa1f76f-b4e9-45d2-a88b-56d22c43d2a5)


```
In [11]: app.content_sources.rest_client.repositories_api.list_repositories(search='FEDORA')['data'][0]['url']
2024-06-06 21:14:20.936 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('search', 'FEDORA')]
Out[11]: 'https://fedorapeople.org/groups/katello/fakerepos/zoo/'
```

### Comment by @swadeley on June 06, 2024 at 08:19 PM UTC

QE says merge on dev ack.

---

## Reviews

### Review by @jlsherrill - Approved on June 07, 2024 at 12:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/689*
