---
type: pull_request
number: 493
title: "Fixes 3215: Add module_hotfixes flag for repositories"
state: merged
author: ezr-ondrej
created: 2023-12-02T13:31:51Z
updated: 2023-12-06T16:02:39Z
closed: 2023-12-04T23:28:36Z
merged: 2023-12-04T23:28:36Z
base_branch: main
head_branch: module_hotfixes
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/493
---

# Pull Request #493: Fixes 3215: Add module_hotfixes flag for repositories

**Author**: @ezr-ondrej
**Created**: December 02, 2023 at 01:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `module_hotfixes`

## Description

## Summary

Add flag for disabling modularity filtering on repository basis.

## Testing steps

Create a repository with 'module_hotfix' set to true or false,  Fetching the repository should show it set to true or false depending on what you set it to.  Existing repos should be set to false, and new repos should default to false.

A snapshot's generated repo config should have it set to 0 or 1 depending on if the repo is set to false or true



---

## Discussion

### Comment by @jlsherrill on December 02, 2023 at 01:32 PM UTC

https://issues.redhat.com/browse/HMS-3205

### Comment by @jlsherrill on December 04, 2023 at 09:30 AM UTC

https://issues.redhat.com/browse/HMS-3215

### Comment by @jlsherrill on December 04, 2023 at 09:30 AM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on December 04, 2023 at 08:20 PM UTC

[test]

### Comment by @jlsherrill on December 04, 2023 at 09:34 PM UTC

good point, it probably should!  i'll add it

### Comment by @swadeley on December 04, 2023 at 11:19 PM UTC

Hi, I created a repo without setting the new flag. 

`In [1]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='fedoratest-snapshot-true',
   ...:            url='https://stephenw.fedorapeople.org/multirepos/6/repo08'
   ...:       )`

The repo was created with flag set to false as expected:
`In [3]: app.content_sources.rest_client.repositories_api.get_repository('db230aac-5c31-4715-a6a5-2fc35309b212')['module_hotfixes']
Out[3]: False`

Created second repo:

`In [5]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='fedoratest-snapshot-true-two',
   ...:            module_hotfixes=True,
   ...:            url='https://stephenw.fedorapeople.org/multirepos/7/repo09'
   ...:       )
`

`In [7]: app.content_sources.rest_client.repositories_api.get_repository('276a6717-07cb-4987-a3cc-346d60985b10')['module_hotfixes']
Out[7]: True
`


```
In [12]: response = app.content_sources.rest_client.repositories_api.get_repo_configuration_file('276a6717-07cb-4987-a3cc-346d60985b10', 'b8541748
    ...: -4a9d-49d2-bc5e-cbe675eff514')

In [13]: import json
    ...: from pprint import pprint

In [14]: pprint(response)
('[fedoratest_snapshot_true_two]\n'
 'name=fedoratest-snapshot-true-two\n'
 'baseurl=https://env-test-<xxx>.openshiftapps.com/pulp/content/0206b201/276a6717-07cb-4987-a3cc-346d60985b10/509a7aa6-592a-44f8-831f-69db8741b92e/\n'
 'module_hotfixes=1\n'
 'gpgcheck=0\n'
 'repo_gpgcheck=0\n'
 'enabled=1\n'
 'gpgkey=\n')
```



### Comment by @ezr-ondrej on December 05, 2023 at 01:21 PM UTC

Thank you all for taking it over the finish line! :tada: 

### Comment by @jlsherrill on December 06, 2023 at 04:02 PM UTC

this is now in prod!

---

## Reviews

### Review by @rverdile - Commented on December 04, 2023 at 09:17 PM UTC

~~I think [this needs to be updated as well](https://github.com/content-services/content-sources-backend/blob/main/pkg/models/repository_configuration.go#L34-L47), in order to update ModulesHotfixes from true -> false~~

Never mind! You did that I just missed it

### Review by @rverdile - Commented on December 04, 2023 at 09:28 PM UTC

should this also be included in the response for the config.repo file?

### Review by @rverdile - Approved on December 04, 2023 at 10:05 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/493*
