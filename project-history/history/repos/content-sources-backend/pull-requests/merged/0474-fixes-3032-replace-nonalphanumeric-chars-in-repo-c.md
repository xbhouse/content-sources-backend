---
type: pull_request
number: 474
title: "Fixes 3032: replace nonalphanumeric chars in repo config ID"
state: merged
author: rverdile
created: 2023-11-16T21:05:58Z
updated: 2023-11-30T18:49:42Z
closed: 2023-11-29T02:49:17Z
merged: 2023-11-29T02:49:17Z
base_branch: main
head_branch: repo-id
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/474
---

# Pull Request #474: Fixes 3032: replace nonalphanumeric chars in repo config ID

**Author**: @rverdile
**Created**: November 16, 2023 at 09:05 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `repo-id`

## Description

## Summary
We want to remove any nonalphanumeric characters from the label the config.repo file and replace with an underscore.

## Testing steps
1. Sync a repository with a name that has nonalphanumeric characters, like parentheses. 
2. Fetch the config.repo file from the `/respositories/:uuid/snapshots/:snapshot_uuid/config.repo` endpoint.
3. The label (the title in the square brackets) should have all the nonalphanumeric characters replaced with an underscore.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 16, 2023 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-3032

### Comment by @xbhouse on November 21, 2023 at 04:24 PM UTC

lgtm! tested with epel 8 repo named `(test-8)` and config.repo label was changed to `[_test_8_]`. tested with epel 9 repo named test-9* and label was changed to `[test_9_]`. checked multiple nonalphanumeric characters in the repo name in the dao snapshot-test, all passed

### Comment by @swadeley on November 23, 2023 at 06:55 AM UTC

Hi @rverdile, rebase please

### Comment by @swadeley on November 29, 2023 at 02:48 AM UTC

Hi,

I created repo:

```
In [7]: app.content_sources.rest_client.repositories_api.create_repository(repo)
Out[7]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2023-11-29T02:04:49Z',
 'last_snapshot_task_uuid': '4772f371-1f4c-43e9-a60e-95295f60c190',
 'last_success_introspection_time': '2023-11-29T02:04:49Z',
 'last_update_introspection_time': '2023-11-29T01:58:36Z',
 'metadata_verification': False,
 'name': 'fedoratest-not-normal-name-2',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 34,
 'snapshot': True,
 'status': 'Valid',
 'url': 'https://stephenw.fedorapeople.org/fakerepos/name%5bnot%5dnormal/',
 'uuid': 'ae34eb51-5c26-4d39-bb87-09c3c0e0e66d'}
```

`In [15]: response = app.content_sources.rest_client.repositories_api.get_repo_configuration_file('ae34eb51-5c26-4d39-bb87-09c3c0e0e66d', 'bcff1b39-6444-4ef6-a6fb-be06ad14d96e')`

```
In [19]: pprint(response)
('[fedoratest_not_normal_name_2]\n'
 'name=fedoratest-not-normal-name-2\n'
 'baseurl=https://env-ephemeral-6eflgk.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/pulp/content/f4b45f50/ae34eb51-5c26-4d39-bb87-09c3c0e0e66d/ffb36f81-bacd-4422-83da-a3e67b38c1f4/\n'
 'gpgcheck=0\n'
 'repo_gpgcheck=0\n'
 'enabled=1\n'
 'gpgkey=\n')
```

Dash became underscore.
lgtm.

---

## Reviews

### Review by @xbhouse - Approved on November 21, 2023 at 04:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/474*
