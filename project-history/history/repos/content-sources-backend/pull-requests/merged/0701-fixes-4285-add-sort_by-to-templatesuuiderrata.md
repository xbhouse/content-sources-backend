---
type: pull_request
number: 701
title: "Fixes 4285: add sort_by to /templates/uuid/errata"
state: merged
author: xbhouse
created: 2024-06-11T20:19:48Z
updated: 2024-06-20T07:40:40Z
closed: 2024-06-13T12:39:00Z
merged: 2024-06-13T12:39:00Z
base_branch: main
head_branch: 4285
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/701
---

# Pull Request #701: Fixes 4285: add sort_by to /templates/uuid/errata

**Author**: @xbhouse
**Created**: June 11, 2024 at 08:19 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4285`

## Description

## Summary

Adds sort_by parameter to `/templates/:uuid/errata`

## Testing steps

- Add a repository with errata, let it snapshot, and create a template
- Make a request to list the template errata with a sort_by parameter. You can sort by issued_date, updated_date, type, or severity. Example request: `/templates/:uuid/errata?sort_by=type:asc`
- Response should list template errata sorted by the column and order specified (type and severity will be sorted alphabetically)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 11, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-4285

### Comment by @swadeley on June 13, 2024 at 12:11 PM UTC

Hi

Here with `sort_by="type:asc"` you can see `'type': 'bugfix',` first in the response:

```
In [8]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="type:asc",limit=4,offset=0)
2024-06-13 12:59:04.842 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'type:asc'), ('limit', 4), ('offset', 0)]
Out[8]: 
{'data': [{'cves': [],
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
                          'ansible (2.9.22)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.9.22/changelogs/CHANGELOG-v2.9.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHBA-2021:2138',
           'id': '01901135-d685-7a65-b9a8-9f533d7c7121',
           'issued_date': '2021-05-26 21:07:57',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'Ansible 2.9.22 release for Ansible Engine 2',
           'title': 'Ansible 2.9.22 release for Ansible Engine 2',
           'type': 'bugfix',
           'updated_date': '2021-05-26 21:07:50'},
```

whereas here with `"type:desc"` you can see  `'type': 'security'` is first:

```
In [9]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="type:desc",limit=4,offset=0)
2024-06-13 12:59:43.260 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'type:desc'), ('limit', 4), ('offset', 0)]
Out[9]: 
{'data': [{'cves': ['CVE-2021-3583'],
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
                          'ansible (2.9.23)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '* CVE-2021-3583 ansible: Template Injection through '
                          'yaml multi-line strings\n'
                          'with ansible facts used in template.\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.9.23/changelogs/CHANGELOG-v2.9.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHSA-2021:2664',
           'id': '01901135-d67d-7f0a-bcd9-d11517467a04',
           'issued_date': '2021-07-07 04:25:02',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': 'An update for ansible is now available for Ansible '
                      'Engine 2\n'
                      '\n'
                      'Red Hat Product Security has rated this update as '
                      'having a security impact\n'
                      'of Moderate. A Common Vulnerability Scoring System '
                      '(CVSS) base score, which\n'
                      'gives a detailed severity rating, is available for each '
                      'vulnerability from\n'
                      'the CVE link(s) in the References section.',
           'title': 'Important: Ansible security and bug fix update (2.9.23)',
           'type': 'security',
           'updated_date': '2021-07-07 04:24:44'},
```

### Comment by @swadeley on June 13, 2024 at 12:22 PM UTC

Hi

Lack of severity, `severity': ''` is sorted inversely to `'severity': 'None'`

```
In [18]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="severity:asc",limit=1,offset=0)['data'][0]['severity']
2024-06-13 13:19:00.597 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=5<>, params=[('sort_by', 'severity:asc'), ('limit', 1), ('offset', 0)]
Out[18]: ''

In [19]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="severity:desc",limit=1,offset=0)['data'][0]['severity']
2024-06-13 13:19:09.805 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'severity:desc'), ('limit', 1), ('offset', 0)]
Out[19]: 'None'
```

i.e. this is just sorting alphabetically 

### Comment by @swadeley on June 13, 2024 at 12:31 PM UTC

Hi

sorting by dates look good to me:

```
In [20]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="issued_date:asc",limit=1,offset=0)['data'][0]['issued_date']
2024-06-13 13:23:54.420 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'issued_date:asc'), ('limit', 1), ('offset', 0)]
Out[20]: '2009-05-20 00:00:00'

In [21]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="issued_date:desc",limit=1,offset=0)['data'][0]['issued_date']
2024-06-13 13:24:04.106 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'issued_date:desc'), ('limit', 1), ('offset', 0)]
Out[21]: '2021-10-14 19:31:51'

In [22]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="updated_date:asc",limit=1,offset=0)['data'][0]['updated_date']
2024-06-13 13:24:35.497 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'updated_date:asc'), ('limit', 1), ('offset', 0)]
Out[22]: '2019-05-21 10:06:22'

In [23]: app.content_sources.rest_client.templates_api.list_template_errata('246d8dda-4f3f-4184-951b-5218d2e06f85',sort_by="updated_date:desc",limit=1,offset=0)['data'][0]['updated_date']
2024-06-13 13:29:38.912 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=9<>, params=[('sort_by', 'updated_date:desc'), ('limit', 1), ('offset', 0)]
Out[23]: ''

```

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on June 12, 2024 at 03:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/701*
