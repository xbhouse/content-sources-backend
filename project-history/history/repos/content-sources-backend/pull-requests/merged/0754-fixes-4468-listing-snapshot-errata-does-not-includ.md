---
type: pull_request
number: 754
title: "Fixes 4468: listing snapshot errata does not include CVEs"
state: merged
author: xbhouse
created: 2024-07-24T14:36:16Z
updated: 2024-08-01T17:30:30Z
closed: 2024-08-01T17:22:56Z
merged: 2024-08-01T17:22:56Z
base_branch: main
head_branch: fix-snapshot-errata-cves
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/754
---

# Pull Request #754: Fixes 4468: listing snapshot errata does not include CVEs

**Author**: @xbhouse
**Created**: July 24, 2024 at 02:36 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-snapshot-errata-cves`

## Description

## Summary

This adds CVEs to the response when listing snapshot errata.

## Testing steps

- Add 2 repos with and without CVEs and let them snapshot (the small RH repo https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/ has CVEs)
- Make a request to list the snapshot errata (`GET /snapshots/:uuid/errata`). You should see the CVEs listed in the response for the one with CVEs and an empty array for the CVEs for the one without
- Without this PR you would see `"cves": null`

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 24, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4468

### Comment by @mayurilahane on July 26, 2024 at 02:00 PM UTC

/retest

### Comment by @mayurilahane on July 30, 2024 at 03:02 PM UTC

/retest

### Comment by @mayurilahane on August 01, 2024 at 02:50 PM UTC

This is the snapshot content for the public repo.
```
In [13]: app.content_sources.rest_client.snapshots_api.list_snapshots("a4944f7b-4aed-4d4f-8a40-d86c74d76a17")
2024-08-01 10:37:52.436 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=c92e8189-cbc7-4efd-828a-326d4b75a11a, params=[]
Out[13]: 
{'data': [{'added_counts': {'rpm.advisory': 32,
                            'rpm.package': 58,
                            'rpm.repo_metadata_file': 1},
           'content_counts': {'rpm.advisory': 32,
                              'rpm.package': 58,
                              'rpm.repo_metadata_file': 1},
           'created_at': '2024-08-01T14:03:23.012786Z',
           'removed_counts': {},
           'repository_path': 'ed3aa925/a4944f7b-4aed-4d4f-8a40-d86c74d76a17/6d5ec1de-054f-4aae-a7d7-94e074e71534',
           'url': 'http://pulp-content-svc:24816/api/pulp-content/ed3aa925/a4944f7b-4aed-4d4f-8a40-d86c74d76a17/6d5ec1de-054f-4aae-a7d7-94e074e71534/',
           'uuid': '610221a8-5341-4fc0-ad7f-018cf2a74136'}],
 'links': {'first': '/api/content-sources/v1/repositories/a4944f7b-4aed-4d4f-8a40-d86c74d76a17/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/a4944f7b-4aed-4d4f-8a40-d86c74d76a17/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

### Comment by @mayurilahane on August 01, 2024 at 02:51 PM UTC

This is the list of CVE's from that snapshot (with some empty CVE's)

pasting first couple of  CVE's  entries to avoid the clutter

```

In [14]: app.content_sources.rest_client.rpms_api.list_snapshot_errata("610221a8-5341-4fc0-ad7f-018cf2a74136")
2024-08-01 10:38:25.473 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=f3ca83c6-d187-4531-b63c-6f75950d13a1, params=[]
Out[14]: 
{'data': [{'cves': ['CVE-2021-3620'],
           'description': 'Ansible is a simple model-driven configuration '
                          'management, multi-node\n'
                          'deployment, and remote-task execution system. '
                          'Ansible works over SSH and\n'
                          'does not require any software or daemons to be '
                          'installed on remote nodes.\n'
                          'Extension modules can be written in any language '
                          'and are transferred to\n'
                          'managed machines automatically.\n'
                          '\n'
                          'The following packages have been upgraded to a '
                          'newer upstream version:\n'
                          'ansible (2.9.27)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '* CVE-2021-3620 Ansible: ansible-connection module '
                          'discloses sensitive info\n'
                          'in traceback error message\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.9.27/changelogs/CHANGELOG-v2.9.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHSA-2021:3872',
           'id': '01910e40-2f34-7d37-9439-694f1b4a19dc',
           'issued_date': '2021-10-14 19:31:51',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': 'An update for ansible is now available for Ansible '
                      'Engine 2\n'
                      '\n'
                      'Red Hat Product Security has rated this update as '
                      'having a security impact\n'
                      'of Important. A Common Vulnerability Scoring System '
                      '(CVSS) base score,\n'
                      'which gives a detailed severity rating, is available '
                      'for each vulnerability\n'
                      'from the CVE link(s) in the References section.',
           'title': 'Important: Ansible security and bug fix update (2.9.27)',
           'type': 'security',
           'updated_date': '2021-10-14 19:31:29'},
          {'cves': [],
           'description': 'Ansible is a simple model-driven configuration '
                          'management, multi-node\n'
                          'deployment, and remote-task execution system. '
                          'Ansible works over SSH and\n'
                          'does not require any software or daemons to be '
                          'installed on remote nodes.\n'
                          'Extension modules can be written in any language '
                          'and are transferred to\n'
                          'managed machines automatically.\n'
                          '\n'
                          'The following packages have been upgraded to a '
                          'newer upstream version:\n'
                          'ansible (2.9.26)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.9.26/changelogs/CHANGELOG-v2.9.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHBA-2021:3544',
           'id': '01910e40-2eb7-7493-a10b-9dd2a9a1c69e',
           'issued_date': '2021-09-14 18:38:27',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'Ansible 2.9.26 release for Ansible Engine 2',
           'title': 'Ansible 2.9.26 release for Ansible Engine 2',
           'type': 'bugfix',
           'updated_date': '2021-09-14 18:38:22'},
```





### Comment by @mayurilahane on August 01, 2024 at 05:21 PM UTC

Tried with non CVE repo 
```
 'url': 'https://stephenw.fedorapeople.org/fakerepos/multiple_errata/',
 'uuid': 'a8ccf83f-f826-4bd8-af03-9f3542add67e'}

In [27]: app.content_sources.rest_client.snapshots_api.list_snapshots("a8ccf83f-f826-4bd8-af03-9f3542add67e")
2024-08-01 11:05:59.415 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=3c50fb06-bb6d-4830-b542-d85f348d52d6, params=[]
Out[27]: 
{'data': [{'added_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'content_counts': {'rpm.advisory': 6, 'rpm.package': 27},
           'created_at': '2024-08-01T15:05:52.646756Z',
           'removed_counts': {},
           'repository_path': '7f173ad3/a8ccf83f-f826-4bd8-af03-9f3542add67e/c3489552-face-4f3f-b3cf-ae1f73d9e86a',
           'url': 'http://pulp-content-svc:24816/api/pulp-content/7f173ad3/a8ccf83f-f826-4bd8-af03-9f3542add67e/c3489552-face-4f3f-b3cf-ae1f73d9e86a/',
           'uuid': 'bc8612cc-874c-4688-83c1-9b6fdbc5959c'}],
 'links': {'first': '/api/content-sources/v1/repositories/a8ccf83f-f826-4bd8-af03-9f3542add67e/snapshots/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/a8ccf83f-f826-4bd8-af03-9f3542add67e/snapshots/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

### Comment by @mayurilahane on August 01, 2024 at 05:22 PM UTC

And it does have null value of CVE's 
```
In [28]: app.content_sources.rest_client.rpms_api.list_snapshot_errata("bc8612cc-874c-4688-83c1-9b6fdbc5959c")
2024-08-01 11:06:34.952 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=53d170f2-674f-468f-afcd-14a0f1979f02, params=[]
Out[28]: 
{'data': [{'cves': [],
           'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '01910e79-c086-7488-858b-3346113acf0a',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'cves': [],
           'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '01910e79-c07b-7881-8af2-05b597237156',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
```

### Comment by @mayurilahane on August 01, 2024 at 05:22 PM UTC

LGTM!!!

---

## Reviews

### Review by @rverdile - Approved on July 24, 2024 at 07:34 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/754*
