---
type: pull_request
number: 907
title: "Fixes 5037: add latest_snapshot_url to repo fetch"
state: merged
author: xbhouse
created: 2024-11-25T19:24:01Z
updated: 2024-11-26T22:00:28Z
closed: 2024-11-26T21:33:40Z
merged: 2024-11-26T21:33:40Z
base_branch: main
head_branch: 5037
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/907
---

# Pull Request #907: Fixes 5037: add latest_snapshot_url to repo fetch

**Author**: @xbhouse
**Created**: November 25, 2024 at 07:24 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `5037`

## Description

## Summary

Adds latest_snapshot_url to response when fetching a repository.

## Testing steps

1. Create a repo and let it snapshot
2. Fetch the repo, you should see the `latest_snapshot_url` in the response. It should be the same `latest_snapshot_url` seen when calling the list repos endpoint

---

## Discussion

### Comment by @jlsherrill on November 25, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5037

### Comment by @xbhouse on November 26, 2024 at 04:30 PM UTC

listing repos:

```
{'data': [{'account_id': '',
           'content_type': 'rpm',
           'distribution_arch': 'x86_64',
           'distribution_versions': ['8'],
           'failed_introspections_count': 0,
           'gpg_key': '-----BEGIN PGP PUBLIC KEY BLOCK-----\n'
                      '\n'
                      'mQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n'
                      '0HtzIG58qDRybwUe89FZprB1ffuUKzdE+HcL3FbNWSSOXVjZIersdXyH3NvnLLLF\n'
                      '0DNRB2ix3bXG9Rh/RXpFsNxDp2CEMdUvbYCzE79K1EnUTVh1L0Of023FtPSZXX0c\n'
                      'u7Pb5DI5lX5YeoXO6RoodrIGYJsVBQWnrWw4xNTconUfNPk0EGZtEnzvH2zyPoJh\n'
                      'XGF+Ncu9XwbalnYde10OCvSWAZ5zTCpoLMTvQjWpbCdWXJzCm6G+/hx9upke546H\n'
                      '5IjtYm4dTIVTnc3wvDiODgBKRzOl9rEOCIgOuGtDxRxcQkjrC+xvg5Vkqn7vBUyW\n'
                      '9pHedOU+PoF3DGOM+dqv+eNKBvh9YF9ugFAQBkcG7viZgvGEMGGUpzNgN7XnS1gj\n'
                      '/DPo9mZESOYnKceve2tIC87p2hqjrxOHuI7fkZYeNIcAoa83rBltFXaBDYhWAKS1\n'
                      'PcXS1/7JzP0ky7d0L6Xbu/If5kqWQpKwUInXtySRkuraVfuK3Bpa+X1XecWi24JY\n'
                      'HVtlNX025xx1ewVzGNCTlWn1skQN2OOoQTV4C8/qFpTW6DTWYurd4+fE0OJFJZQF\n'
                      'buhfXYwmRlVOgN5i77NTIJZJQfYFj38c/Iv5vZBPokO6mffrOTv3MHWVgQARAQAB\n'
                      'tDNSZWQgSGF0LCBJbmMuIChyZWxlYXNlIGtleSAyKSA8c2VjdXJpdHlAcmVkaGF0\n'
                      'LmNvbT6JAjYEEwECACAFAkrgSTsCGwMGCwkIBwMCBBUCCAMEFgIDAQIeAQIXgAAK\n'
                      'CRAZni+R/UMdUWzpD/9s5SFR/ZF3yjY5VLUFLMXIKUztNN3oc45fyLdTI3+UClKC\n'
                      '2tEruzYjqNHhqAEXa2sN1fMrsuKec61Ll2NfvJjkLKDvgVIh7kM7aslNYVOP6BTf\n'
                      'C/JJ7/ufz3UZmyViH/WDl+AYdgk3JqCIO5w5ryrC9IyBzYv2m0HqYbWfphY3uHw5\n'
                      'un3ndLJcu8+BGP5F+ONQEGl+DRH58Il9Jp3HwbRa7dvkPgEhfFR+1hI+Btta2C7E\n'
                      '0/2NKzCxZw7Lx3PBRcU92YKyaEihfy/aQKZCAuyfKiMvsmzs+4poIX7I9NQCJpyE\n'
                      'IGfINoZ7VxqHwRn/d5mw2MZTJjbzSf+Um9YJyA0iEEyD6qjriWQRbuxpQXmlAJbh\n'
                      '8okZ4gbVFv1F8MzK+4R8VvWJ0XxgtikSo72fHjwha7MAjqFnOq6eo6fEC/75g3NL\n'
                      'Ght5VdpGuHk0vbdENHMC8wS99e5qXGNDued3hlTavDMlEAHl34q2H9nakTGRF5Ki\n'
                      'JUfNh3DVRGhg8cMIti21njiRh7gyFI2OccATY7bBSr79JhuNwelHuxLrCFpY7V25\n'
                      'OFktl15jZJaMxuQBqYdBgSay2G0U6D1+7VsWufpzd/Abx1/c3oi9ZaJvW22kAggq\n'
                      'dzdA27UUYjWvx42w9menJwh/0jeQcTecIUd0d0rFcw/c1pvgMMl/Q73yzKgKYw==\n'
                      '=zbHE\n'
                      '-----END PGP PUBLIC KEY BLOCK-----\n'
                      '-----BEGIN PGP PUBLIC KEY BLOCK-----\n'
                      '\n'
                      'mQINBGIpIp4BEAC/o5e1WzLIsS6/JOQCs4XYATYTcf6B6ALzcP05G0W3uRpUQSrL\n'
                      'FRKNrU8ZCelm/B+XSh2ljJNeklp2WLxYENDOsftDXGoyLr2hEkI5OyK267IHhFNJ\n'
                      'g+BN+T5Cjh4ZiiWij6o9F7x2ZpxISE9M4iI80rwSv1KOnGSw5j2zD2EwoMjTVyVE\n'
                      '/t3s5XJxnDclB7ZqL+cgjv0mWUY/4+b/OoRTkhq7b8QILuZp75Y64pkrndgakm1T\n'
                      '8mAGXV02mEzpNj9DyAJdUqa11PIhMJMxxHOGHJ8CcHZ2NJL2e7yJf4orTj+cMhP5\n'
                      'LzJcVlaXnQYu8Zkqa0V6J1Qdj8ZXL72QsmyicRYXAtK9Jm5pvBHuYU2m6Ja7dBEB\n'
                      'Vkhe7lTKhAjkZC5ErPmANNS9kPdtXCOpwN1lOnmD2m04hks3kpH9OTX7RkTFUSws\n'
                      'eARAfRID6RLfi59B9lmAbekecnsMIFMx7qR7ZKyQb3GOuZwNYOaYFevuxusSwCHv\n'
                      '4FtLDIhk+Fge+EbPdEva+VLJeMOb02gC4V/cX/oFoPkxM1A5LHjkuAM+aFLAiIRd\n'
                      'Np/tAPWk1k6yc+FqkcDqOttbP4ciiXb9JPtmzTCbJD8lgH0rGp8ufyMXC9x7/dqX\n'
                      'TjsiGzyvlMnrkKB4GL4DqRFl8LAR02A3846DD8CAcaxoXggL2bJCU2rgUQARAQAB\n'
                      'tDVSZWQgSGF0LCBJbmMuIChhdXhpbGlhcnkga2V5IDMpIDxzZWN1cml0eUByZWRo\n'
                      'YXQuY29tPokCUgQTAQgAPBYhBH5GJCWMQGU11W1vE1BU5KRaY0CzBQJiKSKeAhsD\n'
                      'BQsJCAcCAyICAQYVCgkICwIEFgIDAQIeBwIXgAAKCRBQVOSkWmNAsyBfEACuTN/X\n'
                      'YR+QyzeRw0pXcTvMqzNE4DKKr97hSQEwZH1/v1PEPs5O3psuVUm2iam7bqYwG+ry\n'
                      'EskAgMHi8AJmY0lioQD5/LTSLTrM8UyQnU3g17DHau1NHIFTGyaW4a7xviU4C2+k\n'
                      'c6X0u1CPHI1U4Q8prpNcfLsldaNYlsVZtUtYSHKPAUcswXWliW7QYjZ5tMSbu8jR\n'
                      'OMOc3mZuf0fcVFNu8+XSpN7qLhRNcPv+FCNmk/wkaQfH4Pv+jVsOgHqkV3aLqJeN\n'
                      'kNUnpyEKYkNqo7mNfNVWOcl+Z1KKKwSkIi3vg8maC7rODsy6IX+Y96M93sqYDQom\n'
                      'aaWue2gvw6thEoH4SaCrCL78mj2YFpeg1Oew4QwVcBnt68KOPfL9YyoOicNs4Vuu\n'
                      'fb/vjU2ONPZAeepIKA8QxCETiryCcP43daqThvIgdbUIiWne3gae6eSj0EuUPoYe\n'
                      'H5g2Lw0qdwbHIOxqp2kvN96Ii7s1DK3VyhMt/GSPCxRnDRJ8oQKJ2W/I1IT5VtiU\n'
                      'zMjjq5JcYzRPzHDxfVzT9CLeU/0XQ+2OOUAiZKZ0dzSyyVn8xbpviT7iadvjlQX3\n'
                      'CINaPB+d2Kxa6uFWh+ZYOLLAgZ9B8NKutUHpXN66YSfe79xFBSFWKkJ8cSIMk13/\n'
                      'Ifs7ApKlKCCRDpwoDqx/sjIaj1cpOfLHYjnefg==\n'
                      '=UZd/\n'
                      '-----END PGP PUBLIC KEY BLOCK-----\n'
                      '\n',
           'label': 'ansible-2-for-rhel-8-x86_64-rpms',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-11-26T15:05:50Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 32,
                                              'rpm.package': 58,
                                              'rpm.repo_metadata_file': 1},
                             'content_counts': {'rpm.advisory': 32,
                                                'rpm.package': 58,
                                                'rpm.repo_metadata_file': 1},
                             'created_at': '2024-11-26T15:11:42.33839Z',
                             'removed_counts': {},
                             'repository_name': 'Red Hat Ansible Engine 2 for '
                                                'RHEL 8 x86_64 (RPMs)',
                             'repository_path': 'cs-bb4c4e1734/e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb/67493956-5039-4ed8-9d83-3141fa44f7f2',
                             'repository_uuid': 'e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb',
                             'url': 'http://pulp-content:8000/api/pulp-content/cs-bb4c4e1734/e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb/67493956-5039-4ed8-9d83-3141fa44f7f2/',
                             'uuid': 'ce973b19-4951-4edd-a709-96c597b80803'},
           'last_snapshot_task': {'created_at': '2024-11-26T15:11:06Z',
                                  'ended_at': '2024-11-26T15:11:42Z',
                                  'error': '',
                                  'object_name': 'Red Hat Ansible Engine 2 for '
                                                 'RHEL 8 x86_64 (RPMs)',
                                  'object_type': 'repository',
                                  'object_uuid': 'e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb',
                                  'org_id': '-1',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '13af6407-711a-4810-a1da-f5427899d46c'},
           'last_snapshot_task_uuid': '13af6407-711a-4810-a1da-f5427899d46c',
           'last_snapshot_uuid': 'ce973b19-4951-4edd-a709-96c597b80803',
           'last_success_introspection_time': '2024-11-26T15:05:50Z',
           'last_update_introspection_time': '2024-11-26T15:05:50Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/cs-bb4c4e1734/e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)',
           'org_id': '-1',
           'origin': 'red_hat',
           'package_count': 58,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/',
           'uuid': 'e1525da9-dbbb-4fb8-b3c1-71edcbbff5eb'},
          {'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'label': 'test1',
           'last_introspection_error': '',
           'last_introspection_status': 'Valid',
           'last_introspection_time': '2024-11-26T16:21:18Z',
           'last_snapshot': {'added_counts': {'rpm.advisory': 2,
                                              'rpm.package': 8,
                                              'rpm.packagecategory': 1,
                                              'rpm.packagegroup': 2},
                             'content_counts': {'rpm.advisory': 2,
                                                'rpm.package': 8,
                                                'rpm.packagecategory': 1,
                                                'rpm.packagegroup': 2},
                             'created_at': '2024-11-26T16:21:34.860343Z',
                             'removed_counts': {},
                             'repository_name': 'test1',
                             'repository_path': 'cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/90165459-3eb6-4d96-aa4a-ad5d05746468',
                             'repository_uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c',
                             'url': 'http://pulp-content:8000/api/pulp-content/cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/90165459-3eb6-4d96-aa4a-ad5d05746468/',
                             'uuid': 'f98030fa-9076-48a8-9354-df040bc3cf92'},
           'last_snapshot_task': {'created_at': '2024-11-26T16:21:18Z',
                                  'ended_at': '2024-11-26T16:21:34Z',
                                  'error': '',
                                  'object_name': 'test1',
                                  'object_type': 'repository',
                                  'object_uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c',
                                  'org_id': '3340851',
                                  'status': 'completed',
                                  'type': 'snapshot',
                                  'uuid': '0396a8a4-df11-4639-a988-fbb5919e7a1d'},
           'last_snapshot_task_uuid': '0396a8a4-df11-4639-a988-fbb5919e7a1d',
           'last_snapshot_uuid': 'f98030fa-9076-48a8-9354-df040bc3cf92',
           'last_success_introspection_time': '2024-11-26T16:21:18Z',
           'last_update_introspection_time': '2024-11-26T16:18:28Z',
           'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/latest/',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'test1',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 8,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://fedorapeople.org/groups/katello/fakerepos/zoo2/',
           'uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
```

fetching a repo:

```
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'label': 'test1',
 'last_introspection_error': '',
 'last_introspection_status': 'Valid',
 'last_introspection_time': '2024-11-26T16:21:18Z',
 'last_snapshot': {'added_counts': {'rpm.advisory': 2,
                                    'rpm.package': 8,
                                    'rpm.packagecategory': 1,
                                    'rpm.packagegroup': 2},
                   'content_counts': {'rpm.advisory': 2,
                                      'rpm.package': 8,
                                      'rpm.packagecategory': 1,
                                      'rpm.packagegroup': 2},
                   'created_at': '2024-11-26T16:21:34.860343Z',
                   'removed_counts': {},
                   'repository_name': 'test1',
                   'repository_path': 'cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/90165459-3eb6-4d96-aa4a-ad5d05746468',
                   'repository_uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c',
                   'url': 'http://pulp-content:8000/api/pulp-content/cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/90165459-3eb6-4d96-aa4a-ad5d05746468/',
                   'uuid': 'f98030fa-9076-48a8-9354-df040bc3cf92'},
 'last_snapshot_task': {'created_at': '2024-11-26T16:21:18Z',
                        'ended_at': '2024-11-26T16:21:34Z',
                        'error': '',
                        'object_name': 'test1',
                        'object_type': 'repository',
                        'object_uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c',
                        'org_id': '3340851',
                        'status': 'completed',
                        'type': 'snapshot',
                        'uuid': '0396a8a4-df11-4639-a988-fbb5919e7a1d'},
 'last_snapshot_task_uuid': '0396a8a4-df11-4639-a988-fbb5919e7a1d',
 'last_snapshot_uuid': 'f98030fa-9076-48a8-9354-df040bc3cf92',
 'last_success_introspection_time': '2024-11-26T16:21:18Z',
 'last_update_introspection_time': '2024-11-26T16:18:28Z',
 'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/latest/',
 'metadata_verification': False,
 'module_hotfixes': False,
 'name': 'test1',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 8,
 'snapshot': True,
 'status': 'Valid',
 'url': 'https://fedorapeople.org/groups/katello/fakerepos/zoo2/',
 'uuid': '6c15d657-14e0-4a2f-a599-c73ac56d433c'}
 ```
 
`'latest_snapshot_url': 'http://pulp-content:8000/api/pulp-content/cs-573ba53a4a/6c15d657-14e0-4a2f-a599-c73ac56d433c/latest/'` exists and is the same in both list and fetch

---

## Reviews

### Review by @jlsherrill - Commented on November 25, 2024 at 08:27 PM UTC

### Review by @xbhouse - Commented on November 25, 2024 at 08:49 PM UTC

### Review by @jlsherrill - Commented on November 25, 2024 at 08:55 PM UTC

### Review by @xbhouse - Commented on November 25, 2024 at 09:02 PM UTC

### Review by @xbhouse - Commented on November 25, 2024 at 09:12 PM UTC

### Review by @jlsherrill - Approved on November 25, 2024 at 10:17 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/907*
