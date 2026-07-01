---
type: pull_request
number: 679
title: "Fixes 3656: add api to list template errata"
state: merged
author: xbhouse
created: 2024-05-28T16:31:56Z
updated: 2024-06-03T20:05:42Z
closed: 2024-06-03T20:05:41Z
merged: 2024-06-03T20:05:41Z
base_branch: main
head_branch: list-template-errata
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/679
---

# Pull Request #679: Fixes 3656: add api to list template errata

**Author**: @xbhouse
**Created**: May 28, 2024 at 04:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `list-template-errata`

## Description

## Summary

- Adds API for listing content template errata: `/templates/:uuid/errata` (similar to the list errata endpoint for snapshots)
- Includes optional query parameters `search`, `limit`, `type`, and `severity`
- Depends on this [tang PR](https://github.com/content-services/tang/pull/10) to add in the CVEs

## Testing steps

- Add a custom repository with errata and a RH repository and let them snapshot (this will snapshot a small RH repo: `OPTIONS_REPOSITORY_IMPORT_FILTER=small go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go nightly-jobs`)
- Create a template with those repositories, selecting the latest date 
- Make a request to the `/templates/:uuid/errata` endpoint with the template uuid and you should see the template's errata (should include errata from both of the snapshots in the template) in a response similar to this:
```
{
    "data": [
        {
            "id": "018fbf80-9d1e-7458-b10b-9b39b3b652a7",
            "errata_id": "RHSA-2020:3602",
            "title": "Important: Ansible security and bug fix update (2.9.13)",
            "summary": "An update for ansible is now available for Ansible Engine 2\n\nRed Hat Product Security has rated this update as having a security impact\nof Important. A Common Vulnerability Scoring System (CVSS) base score,\nwhich gives a detailed severity rating, is available for each vulnerability\nfrom the CVE link(s) in the References section.",
            "description": "Ansible is a simple model-driven configuration management, multi-node\ndeployment, and remote-task execution system. Ansible works over SSH and\ndoes not require any software or daemons to be installed on remote nodes.\nExtension modules can be written in any language and are transferred to\nmanaged machines automatically.\n\nThe following packages have been upgraded to a newer upstream version:\nansible (2.9.13)\n\nBug Fix(es):\n* CVE-2020-14365 ansible: dnf module install packages with no GPG signature\n\nSee:\nhttps://github.com/ansible/ansible/blob/v2.9.13/changelogs/CHANGELOG-v2.9.rst\nfor details on bug fixes in this release.",
            "issued_date": "2020-09-01 19:10:30",
            "updated_date": "2020-09-01 19:10:25",
            "type": "security",
            "severity": "Important",
            "reboot_suggested": false,
            "cves": [
                "CVE-2020-14365"
            ]
        },
        ...
```        

- To test that the correct snapshot's errata are shown, you can update the custom repository with a different URL and let it snapshot
- Edit the template date to include this latest snapshot and view the errata again. The errata list should reflect the latest snapshot's errata
- Edit the template to remove the custom repository and then view the errata again, you should only see the errata from the RH repository's snapshot

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @xbhouse on May 28, 2024 at 04:40 PM UTC

need to update this to return a list of CVEs, i think that would require a change in tang. 

also, this endpoint refers to errata which is consistent with the similar endpoint to list snapshot errata, but the task says to use advisories. should i refer to these as errata or advisories?

EDIT: based on our conversation in standup, will continue to use errata here and refer to them as advisories in the UI

### Comment by @jlsherrill on May 28, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-3656

### Comment by @swadeley on June 03, 2024 at 07:41 AM UTC

/retest

### Comment by @swadeley on June 03, 2024 at 06:36 PM UTC

Hi

ACK from me

API works[1], UI has problems not related to this PR[2]


@jlsherrill Merge on ACK



[1]
```
In [3]: app.content_sources.rest_client.templates_api.get_template('0bcf3753-e81f-48cc-b40c-64a7210f27d9')
2024-06-03 16:22:01.786 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=9c84f4e7-4f68-475b-b393-d4cff9d904d5, params=[]
Out[3]: 
{'arch': 'x86_64',
 'date': '2024-06-03T04:00:00Z',
 'description': 'my 1st template mk II',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['bd473a30-69f3-4ffe-b86e-c7c433a20540',
                      '6168cd88-f36d-4224-aec2-6dfd8eee7a37'],
 'rhsm_environment_id': '0bcf3753e81f48ccb40c64a7210f27d9',
 'uuid': '0bcf3753-e81f-48cc-b40c-64a7210f27d9',
 'version': '8'}

In [4]: app.content_sources.rest_client.templates_api.list_template_errata('0bcf3753-e81f-48cc-b40c-64a7210f27d9',limit=10,offset=30)
2024-06-03 16:22:14.893 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 10), ('offset', 30)]
Out[4]: 
{'data': [{'cves': [],
           'description': 'Ansible is a simple model-driven configuration '
                          'management, multi-node\n'
                          'deployment, and remote-task execution system. '
                          'Ansible works over SSH and\n'
                          'does\n'
                          'not require any software or daemons to be installed '
                          'on remote nodes.\n'
                          'Extension\n'
                          'modules can be written in any language and are '
                          'transferred to managed\n'
                          'machines\n'
                          'automatically.\n'
                          '\n'
                          'The following packages have been upgraded to a '
                          'newer upstream version:\n'
                          'ansible\n'
                          '(2.8.1)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.8.1/changelogs/CHANGELOG-v2.8.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHBA-2019:1433',
           'id': '018fddef-c58c-793e-92ce-8b2bf013864f',
           'issued_date': '2019-06-11 08:15:49',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'Ansible 2.8.1 release for Ansible Engine 2',
           'title': 'Ansible 2.8.1 release for Ansible Engine 2',
           'type': 'bugfix',
           'updated_date': '2019-06-11 08:15:49'},
          {'cves': [],
           'description': 'Ansible is a simple model-driven configuration '
                          'management, multi-node\n'
                          'deployment, and remote-task execution system. '
                          'Ansible works over SSH and does not require any '
                          'software or daemons to be installed on remote '
                          'nodes. Extension modules can be written in any '
                          'language and are transferred to managed machines '
                          'automatically.\n'
                          '\n'
                          'The following packages have been upgraded to a '
                          'newer upstream version: ansible (2.8.0)\n'
                          '\n'
                          'Bug Fix(es):\n'
                          '\n'
                          'See:\n'
                          'https://github.com/ansible/ansible/blob/v2.8.0/changelogs/CHANGELOG-v2.8.rst\n'
                          'for details on bug fixes in this release.',
           'errata_id': 'RHEA-2019:1251',
           'id': '018fddef-c5e8-7c20-aa84-7b5abeabb8d6',
           'issued_date': '2019-05-21 10:06:22',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'Ansible 2.8.0 release for Ansible Engine 2 on RHEL8',
           'title': 'Ansible 2.8.0 release for Ansible Engine 2 on RHEL8',
           'type': 'enhancement',
           'updated_date': '2019-05-21 10:06:22'},
          {'cves': [],
           'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fde45-ae93-7ff2-93b7-0ad017c2a0cc',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
<snip>
          {'cves': [],
           'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fde45-ae8d-7e8e-a92e-a615cfa29cf4',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''},
<snip>
          {'cves': [],
           'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fde45-ae96-7533-a918-f60e2f80d5ee',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'cves': [],
           'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fde45-ae90-7e2d-bc2b-c5193913bcf8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/templates/0bcf3753-e81f-48cc-b40c-64a7210f27d9/errata?limit=10&offset=0',
           'last': '/api/content-sources/v1/templates/0bcf3753-e81f-48cc-b40c-64a7210f27d9/errata?limit=10&offset=30',
           'prev': '/api/content-sources/v1/templates/0bcf3753-e81f-48cc-b40c-64a7210f27d9/errata?limit=10&offset=20'},
 'meta': {'count': 38, 'limit': 10, 'offset': 30}}
```


[2] https://issues.redhat.com/browse/HMS-4242


---

## Reviews

### Review by @rverdile - Approved on June 03, 2024 at 08:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/679*
