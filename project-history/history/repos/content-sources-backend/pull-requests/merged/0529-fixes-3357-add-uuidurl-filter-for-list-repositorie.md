---
type: pull_request
number: 529
title: "Fixes 3357: Add UUID/URL filter for list repositories"
state: merged
author: Andrewgdewar
created: 2024-01-11T18:14:17Z
updated: 2024-01-19T10:00:31Z
closed: 2024-01-19T09:55:01Z
merged: 2024-01-19T09:55:01Z
base_branch: main
head_branch: HMS-3357
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/529
---

# Pull Request #529: Fixes 3357: Add UUID/URL filter for list repositories

**Author**: @Andrewgdewar
**Created**: January 11, 2024 at 06:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3357`

## Description

## Summary

- Adds 1 new variable  (UUID)  to the FilterData struct for repoconfigs.
- Updates 1 variable (URL) to now accept a comma separated string list.

## Testing steps




---

## Discussion

### Comment by @jlsherrill on January 11, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-3357

### Comment by @swadeley on January 16, 2024 at 07:06 AM UTC

/retest

### Comment by @swadeley on January 16, 2024 at 09:16 AM UTC

Hi

Looks good (url and uuid)

`In [5]: app.content_sources.rest_client.repositories_api.list_repositories(uuid='87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8,7999300c-6063-43d6-a685-054c9b700d87')
2024-01-16 08:49:10.965 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('uuid', '87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8,7999300c-6063-43d6-a685-054c9b700d87')]
Out[5]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2024-01-16T08:33:35Z',
           'last_success_introspection_time': '2024-01-16T08:33:35Z',
           'last_update_introspection_time': '2024-01-16T08:00:22Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'comps-env-test',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 10003,
           'snapshot': False,
           'status': 'Valid',
           'url': 'https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/',
           'uuid': '87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8'},
          {'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': 'https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg',
           'last_introspection_error': '',
           'last_introspection_time': '2024-01-16T08:48:28Z',
           'last_success_introspection_time': '2024-01-16T08:48:28Z',
           'last_update_introspection_time': '2024-01-16T08:48:28Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'gpgtest-snapshot-false',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 33,
           'snapshot': False,
           'status': 'Valid',
           'url': 'http://jlsherrill.fedorapeople.org/fake-repos/signed/',
           'uuid': '7999300c-6063-43d6-a685-054c9b700d87'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&uuid=87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8,7999300c-6063-43d6-a685-054c9b700d87',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&uuid=87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8,7999300c-6063-43d6-a685-054c9b700d87'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [6]: app.content_sources.rest_client.repositories_api.list_repositories(url='https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/,http://jlsherrill.fedorapeople.org/fake-repos/signed/')
2024-01-16 09:11:33.043 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('url', 'https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/,http://jlsherrill.fedorapeople.org/fake-repos/signed/')]
Out[6]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2024-01-16T08:33:35Z',
           'last_success_introspection_time': '2024-01-16T08:33:35Z',
           'last_update_introspection_time': '2024-01-16T08:00:22Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'comps-env-test',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 10003,
           'snapshot': False,
           'status': 'Valid',
           'url': 'https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/',
           'uuid': '87c6c5d6-fb59-4b6f-ba3e-81446c5efbe8'},
          {'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': 'https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg',
           'last_introspection_error': '',
           'last_introspection_time': '2024-01-16T08:48:28Z',
           'last_success_introspection_time': '2024-01-16T08:48:28Z',
           'last_update_introspection_time': '2024-01-16T08:48:28Z',
           'metadata_verification': False,
           'module_hotfixes': False,
           'name': 'gpgtest-snapshot-false',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 33,
           'snapshot': False,
           'status': 'Valid',
           'url': 'http://jlsherrill.fedorapeople.org/fake-repos/signed/',
           'uuid': '7999300c-6063-43d6-a685-054c9b700d87'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&url=https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/,http://jlsherrill.fedorapeople.org/fake-repos/signed/',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&url=https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/,http://jlsherrill.fedorapeople.org/fake-repos/signed/'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
`


### Comment by @swadeley on January 16, 2024 at 09:19 AM UTC

Hi @jlsherrill 

I see pulp start up failed, so pr checks failed.

### Comment by @swadeley on January 17, 2024 at 03:46 PM UTC

/retest

### Comment by @jlsherrill on January 18, 2024 at 06:23 PM UTC

small comment, otherwise looks good

### Comment by @swadeley on January 19, 2024 at 09:54 AM UTC

Hi, apidoc for QE updated with latest changes. Thank you.

---

## Reviews

### Review by @swadeley - Commented on January 18, 2024 at 09:06 AM UTC

### Review by @swadeley - Commented on January 18, 2024 at 09:08 AM UTC

### Review by @swadeley - Commented on January 18, 2024 at 09:09 AM UTC

### Review by @jlsherrill - Commented on January 18, 2024 at 06:20 PM UTC

### Review by @jlsherrill - Approved on January 18, 2024 at 06:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/529*
