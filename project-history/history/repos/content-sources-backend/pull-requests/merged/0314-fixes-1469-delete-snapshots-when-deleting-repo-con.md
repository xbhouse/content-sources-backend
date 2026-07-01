---
type: pull_request
number: 314
title: "Fixes 1469: delete snapshots when deleting repo config"
state: merged
author: rverdile
created: 2023-06-21T18:37:34Z
updated: 2023-06-30T09:25:14Z
closed: 2023-06-28T14:08:37Z
merged: 2023-06-28T14:08:37Z
base_branch: main
head_branch: pulp-delete
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/314
---

# Pull Request #314: Fixes 1469: delete snapshots when deleting repo config

**Author**: @rverdile
**Created**: June 21, 2023 at 06:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `pulp-delete`

## Description

## Summary
Starts a task to delete snapshots and pulp references when deleting a repository config. If creating a snapshot is in progress when the delete request is made, the API will reject the request to delete the repository config.

## Testing steps
1. Create a repository
2. Wait for the sync task to complete
3. Make note of the pulp hrefs created for the snapshot (remote, repository, repository version, distribution).
4. Delete the repository configuration
5. Wait for the delete task to complete, then check the earlier pulp hrefs to see that they were deleted.
6. Create a repository again
7. Delete the repository before the sync task completes, your request to delete the repository should be rejected.

Note that this doesn't handle the delete snapshot task exiting early. If the server closes during this task, the snapshot/pulp objects may or may not be deleted.

Todo
- [x] Tests
- [x] Cleanup

---

## Discussion

### Comment by @jlsherrill on June 21, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-1469

### Comment by @jlsherrill on June 23, 2023 at 03:24 PM UTC

this needs a rebase now

### Comment by @jlsherrill on June 26, 2023 at 08:51 PM UTC

i got an error when trying to delete:
```
4:39PM ERR error enqueuing task error="Could not find repository with UUID 8ab52ad1-f9bf-41af-9116-8bf602ab31fd"
4:39PM ERR error enqueuing task error="error enqueuing task: ERROR: invalid input syntax for type uuid: \"\" (SQLSTATE 22P02)"
```

But the repository seemed to have been deleted, i'm guessing the snapshots and pulp bits didn't get cleaned up

### Comment by @rverdile on June 27, 2023 at 03:54 PM UTC

> i got an error when trying to delete:
> 
> ```
> 4:39PM ERR error enqueuing task error="Could not find repository with UUID 8ab52ad1-f9bf-41af-9116-8bf602ab31fd"
> 4:39PM ERR error enqueuing task error="error enqueuing task: ERROR: invalid input syntax for type uuid: \"\" (SQLSTATE 22P02)"
> ```
> 
> But the repository seemed to have been deleted, i'm guessing the snapshots and pulp bits didn't get cleaned up

This should be fixed now. I'm surprised I didn't hit this bug sooner. It seems like I should have hit it right away.

### Comment by @swadeley on June 28, 2023 at 08:28 AM UTC

Hi

How to check the Pulp hrefs?

I can find distribution path:

```
In [5]: app.content_sources.rest_client.repositories_api.list_snapshots('7d1cfb0f-b4a9-4ce9-b69c-8dd9991c177c')
2023-06-28 09:09:02.531 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=e7EmEyoTmmR8Rsvx7P5h9j7wyN3zLMDV, params=[]
Out[5]: 
{'data': [{'content_counts': {'rpm.package': 1},
           'created_at': '2023-06-28T08:00:15.906089Z',
           'distribution_path': '7d1cfb0f-b4a9-4ce9-b69c-8dd9991c177c/27c7f2f3-17b1-4368-902e-9b80e079d9e5'}],
 'links': {'first': '/api/content-sources/v1/repositories/7d1cfb0f-b4a9-4ce9-b69c-8dd9991c177c/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/7d1cfb0f-b4a9-4ce9-b69c-8dd9991c177c/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

### Comment by @swadeley on June 28, 2023 at 02:00 PM UTC

Hi

Justin explained how to get pulp's hrefs (see above on how to get distribution path):

```
iqe]$ oc rsh $POD


sh-4.4$ curl -L "http://pulp-web-svc:24880/pulp/content/7d1cfb0f-b4a9-4ce9-b69c-8dd9991c177c/27c7f2f3-17b1-4368-902e-9b80e079d9e5/repodata/repomd.xml" 
<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo" xmlns:rpm="http://linux.duke.edu/metadata/rpm">
  <revision>1687939202</revision>
  <data type="primary">
    <checksum type="sha256">08f849d83a8a4062ea1ce878585e805ddc961252c7b41e1922a525f0549b53a8</checksum>
    <open-checksum type="sha256">a308caa4c7558769e68d3ab894996446407b4e50d5c7e9c99471fc98a6baa2d5</open-checksum>
    <location href="repodata/08f849d83a8a4062ea1ce878585e805ddc961252c7b41e1922a525f0549b53a8-primary.xml.gz"/>
    <timestamp>1687939200</timestamp>
    <size>612</size>
    <open-size>1215</open-size>
  </data>
  <data type="filelists">
    <checksum type="sha256">9c8653246307e8214567da37c9655399b66c1522b2e6af7d3d7c788e10921b2e</checksum>
    <open-checksum type="sha256">86919c7c025a7d925a7d7f37ee0406e61737d6dbda8960e67fc7d463127c44a2</open-checksum>
    <location href="repodata/9c8653246307e8214567da37c9655399b66c1522b2e6af7d3d7c788e10921b2e-filelists.xml.gz"/>
    <timestamp>1687939200</timestamp>
    <size>245</size>
    <open-size>325</open-size>
  </data>
  <data type="other">
    <checksum type="sha256">718fb104dab8b92ccccb866b527fd74087a7c068126f0ba15e4dc4a095b7c7c0</checksum>
    <open-checksum type="sha256">986b235ee6d51fd1c1d681a0bf463e6a24554e5ceb6d41d724f436356b8d0283</open-checksum>
    <location href="repodata/718fb104dab8b92ccccb866b527fd74087a7c068126f0ba15e4dc4a095b7c7c0-other.xml.gz"/>
    <timestamp>1687939200</timestamp>
    <size>228</size>
    <open-size>287</open-size>
  </data>
</repomd>
sh-4.4$ 

```

You need `-L` to make curl follow redirects (because storage is in S3 bucket.)

### Comment by @swadeley on June 28, 2023 at 02:08 PM UTC

Hi

step 4 and 5 verified:
Pulp's hrefs were gone by the second time I ran the curl command.

step 6 & 7:
Created new repo, tried to delete ASAP in UI, got warning pop up:

Error deleting item from content list
Request failed with status code 400

Thank you.

---

## Reviews

### Review by @jlsherrill - Commented on June 21, 2023 at 08:18 PM UTC

### Review by @jlsherrill - Commented on June 21, 2023 at 08:30 PM UTC

### Review by @rverdile - Commented on June 22, 2023 at 02:10 PM UTC

### Review by @jlsherrill - Commented on June 22, 2023 at 02:21 PM UTC

### Review by @jlsherrill - Commented on June 22, 2023 at 02:22 PM UTC

### Review by @jlsherrill - Commented on June 22, 2023 at 02:30 PM UTC

### Review by @rverdile - Commented on June 22, 2023 at 02:35 PM UTC

### Review by @rverdile - Commented on June 22, 2023 at 02:35 PM UTC

### Review by @rverdile - Commented on June 22, 2023 at 02:41 PM UTC

### Review by @jlsherrill - Commented on June 22, 2023 at 05:53 PM UTC

### Review by @jlsherrill - Commented on June 22, 2023 at 06:41 PM UTC

### Review by @jlsherrill - Approved on June 27, 2023 at 08:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/314*
