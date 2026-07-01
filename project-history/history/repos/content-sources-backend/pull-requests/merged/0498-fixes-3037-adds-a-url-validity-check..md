---
type: pull_request
number: 498
title: "Fixes 3037: Adds a url validity check."
state: merged
author: Andrewgdewar
created: 2023-12-06T23:03:57Z
updated: 2023-12-13T15:30:37Z
closed: 2023-12-13T15:27:51Z
merged: 2023-12-13T15:27:51Z
base_branch: main
head_branch: HMS-3037
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/498
---

# Pull Request #498: Fixes 3037: Adds a url validity check.

**Author**: @Andrewgdewar
**Created**: December 06, 2023 at 11:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3037`

## Description

## Summary

This adds a simple check on the model for repositories.URL to see if it is a valid. 

## Testing steps

Manually testing can be done by attempting to hit any of the repositories create/edit endpoints with an incorrect URL. 

Invalid URL example: repo_2_url  

Valid URL example: https://repo_2_url.org

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate 


---

## Discussion

### Comment by @jlsherrill on December 06, 2023 at 11:06 PM UTC

https://issues.redhat.com/browse/HMS-3037

### Comment by @swadeley on December 10, 2023 at 05:54 AM UTC

/retest

### Comment by @swadeley on December 10, 2023 at 12:15 PM UTC

Hi

Not working for me:


```
In [5]:      repo = dict(
   ...:            snapshot=False,
   ...:            name='example-repo-test3',
   ...:            url='http://test3',
   ...:            )

In [6]: app.content_sources.rest_client.repositories_api.create_repository(repo)
2023-12-10 20:09:01.746 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=GGZKPJoUyCbmDSrvKiEWgTIGuhOlBkzy, params=[]
Out[6]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'example-repo-test3',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': False,
 'status': 'Pending',
 'url': 'http://test3/',
 'uuid': 'c38b2293-42da-4c94-a18f-6e5050ae536b'}
```

EDIT: I deployed the PR image with `FEATURES_SNAPSHOTS_ENABLED=false` in case that matters.

### Comment by @jlsherrill on December 12, 2023 at 06:59 PM UTC

/retest

### Comment by @Andrewgdewar on December 13, 2023 at 02:50 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on December 07, 2023 at 01:25 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/498*
