---
type: pull_request
number: 579
title: "Fixes 3612: add sslclient options to yum config"
state: merged
author: jlsherrill
created: 2024-02-15T14:26:42Z
updated: 2024-02-16T19:50:05Z
closed: 2024-02-16T13:08:56Z
merged: 2024-02-16T13:08:56Z
base_branch: main
head_branch: 3612
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/579
---

# Pull Request #579: Fixes 3612: add sslclient options to yum config

**Author**: @jlsherrill
**Created**: February 15, 2024 at 02:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3612`

## Description

## Summary

adds  ssl options to repo config
```
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem
```

## Testing steps

snapshot a repo, copy/download the repo config, look for the piece above

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 15, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-3612

### Comment by @swadeley on February 16, 2024 at 01:08 PM UTC

Hi

`In [2]: app.content_sources.rest_client.repositories_api.list_repositories()['data'][0]
2024-02-16 12:58:56.211 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[2]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': '',
 'last_introspection_error': '',
 'last_introspection_time': '2024-02-16T12:02:00Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2024-02-16T12:02:13.227891Z',
                   'removed_counts': {},
                   'repository_path': '94288ffc/8c5f27de-3832-4e98-b919-3bd8d887b2eb/996bf0fa-4b97-4e3f-93ee-0a973b2dcd63',
                   'url': 'https://<>.openshiftapps.com/api/pulp-content/94288ffc/8c5f27de-3832-4e98-b919-3bd8d887b2eb/996bf0fa-4b97-4e3f-93ee-0a973b2dcd63/',
                   'uuid': 'a42d0095-67de-4913-a95b-0716451dbbaa'},
 'last_snapshot_task_uuid': '0f4a227f-9ee1-4de4-8702-a5a5c476b6c4',
 'last_snapshot_uuid': 'a42d0095-67de-4913-a95b-0716451dbbaa',
 'last_success_introspection_time': '2024-02-16T12:02:00Z',
 'last_update_introspection_time': '2024-02-16T12:02:00Z',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'fedoratest-snapshot-true-one',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': 'http://stephenw.fedorapeople.org/multirepos/5/repo05/',
 'uuid': '8c5f27de-3832-4e98-b919-3bd8d887b2eb'}

In [3]: print(app.content_sources.rest_client.repositories_api.get_repo_configuration_file('8c5f27de-3832-4e98-b919-3bd8d887b2eb','a42d0095-67de-4913-a95b-0716451dbbaa')
   ...: )
2024-02-16 13:05:24.137 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
[fedoratest_snapshot_true_one]
name=fedoratest-snapshot-true-one
baseurl=https://<>.openshiftapps.com/api/pulp-content/94288ffc/8c5f27de-3832-4e98-b919-3bd8d887b2eb/996bf0fa-4b97-4e3f-93ee-0a973b2dcd63/
module_hotfixes=0
gpgcheck=0
repo_gpgcheck=0
enabled=1
gpgkey=
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem


In [4]: 
`

lgtm, thank you

---

## Reviews

### Review by @rverdile - Approved on February 15, 2024 at 08:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/579*
