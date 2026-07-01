---
type: pull_request
number: 673
title: "Fixes 4009: snapshot errata should be sortable"
state: merged
author: xbhouse
created: 2024-05-20T16:03:20Z
updated: 2024-05-28T17:00:38Z
closed: 2024-05-28T16:22:43Z
merged: 2024-05-28T16:22:43Z
base_branch: main
head_branch: 4009
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/673
---

# Pull Request #673: Fixes 4009: snapshot errata should be sortable

**Author**: @xbhouse
**Created**: May 20, 2024 at 04:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4009`

## Description

## Summary

- Enables sorting on issued_date, updated_date, type, and severity for snapshot errata
- Dependent on this [tang PR](https://github.com/content-services/tang/pull/9) and a new tang release

## Testing steps

- Create a repository and let it snapshot
- Make a request to list the snapshot errata and sort by issued_date, updated_date, type, and severity (asc or desc). Example request:`/snapshots/:uuid/errata?sort_by=type:desc`
- This should return a list of errata sorted by the specified field and direction
- Listing the snapshot errata with no sort_by parameter should still default to sorting by issued_date descending

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate

---

## Discussion

### Comment by @jlsherrill on May 20, 2024 at 04:37 PM UTC

https://issues.redhat.com/browse/HMS-4009

### Comment by @swadeley on May 28, 2024 at 07:17 AM UTC

/retest

### Comment by @swadeley on May 28, 2024 at 10:45 AM UTC

Hi

default (issued date):
```
In [13]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('bab20333-0a11-46b0-8719-8c36178cb087')
2024-05-28 11:38:45.588 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[13]: 
{'data': [{'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fbec1-1ed3-7e5e-9d6c-3813a045705b',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018fbec1-1ec8-7479-8d25-0ff1fba8efeb',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fbec1-1ecd-70ba-b9c6-7e4de615911c',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018fbec1-1ec3-76d6-a46e-bdaa096f892b',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fbec1-1ed6-7088-a4ff-c7379a246ecc',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fbec1-1ed0-75bc-857c-27f73d7162f8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0',
           'last': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [14]: 

```
by desc:
```
In [15]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('bab20333-0a11-46b0-8719-8c36178cb087', sort_by='desc')
2024-05-28 11:40:48.025 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'desc')]
Out[15]: 
{'data': [{'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fbec1-1ed3-7e5e-9d6c-3813a045705b',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018fbec1-1ec8-7479-8d25-0ff1fba8efeb',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fbec1-1ecd-70ba-b9c6-7e4de615911c',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018fbec1-1ec3-76d6-a46e-bdaa096f892b',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fbec1-1ed6-7088-a4ff-c7379a246ecc',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fbec1-1ed0-75bc-857c-27f73d7162f8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=desc',
           'last': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=desc'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [16]: 

```
by updated_date:

```
In [16]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('bab20333-0a11-46b0-8719-8c36178cb087', sort_by='updated_date')
2024-05-28 11:42:34.226 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'updated_date')]
Out[16]: 
{'data': [{'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fbec1-1ed0-75bc-857c-27f73d7162f8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018fbec1-1ec8-7479-8d25-0ff1fba8efeb',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018fbec1-1ec3-76d6-a46e-bdaa096f892b',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fbec1-1ed6-7088-a4ff-c7379a246ecc',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fbec1-1ed3-7e5e-9d6c-3813a045705b',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fbec1-1ecd-70ba-b9c6-7e4de615911c',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=updated_date',
           'last': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=updated_date'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [17]: 

```

by type:

```
In [17]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('bab20333-0a11-46b0-8719-8c36178cb087', sort_by='type')
2024-05-28 11:43:44.006 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'type')]
Out[17]: 
{'data': [{'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fbec1-1ed0-75bc-857c-27f73d7162f8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018fbec1-1ec8-7479-8d25-0ff1fba8efeb',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018fbec1-1ec3-76d6-a46e-bdaa096f892b',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fbec1-1ecd-70ba-b9c6-7e4de615911c',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fbec1-1ed3-7e5e-9d6c-3813a045705b',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fbec1-1ed6-7088-a4ff-c7379a246ecc',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=type',
           'last': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=type'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [18]: 


```

by severity:

```
In [18]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('bab20333-0a11-46b0-8719-8c36178cb087', sort_by='severity')
2024-05-28 11:44:30.496 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('sort_by', 'severity')]
Out[18]: 
{'data': [{'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018fbec1-1ecd-70ba-b9c6-7e4de615911c',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'acme-package_Erratum description',
           'errata_id': 'RHSA-2009:1425',
           'id': '018fbec1-1ed0-75bc-857c-27f73d7162f8',
           'issued_date': '2009-05-20 00:00:00',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': '',
           'title': 'acme-package_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018fbec1-1ec8-7479-8d25-0ff1fba8efeb',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': 'Important',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018fbec1-1ec3-76d6-a46e-bdaa096f892b',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018fbec1-1ed6-7088-a4ff-c7379a246ecc',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018fbec1-1ed3-7e5e-9d6c-3813a045705b',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=severity',
           'last': '/api/content-sources/v1/snapshots/bab20333-0a11-46b0-8719-8c36178cb087/errata?limit=100&offset=0&sort_by=severity'},
 'meta': {'count': 6, 'limit': 100, 'offset': 0}}

In [19]: 


```




### Comment by @swadeley on May 28, 2024 at 10:51 AM UTC

@xbhouse Does the sorting for severity look wrong to you in the previous comment? I see "Moderate", "Low", "Important", "Critical".

Thank you

EDIT: I expect "Low" before "Moderate"

### Comment by @xbhouse on May 28, 2024 at 01:58 PM UTC

@swadeley sorting is only available on these fields: issued_date, updated_date, type, and severity. so not for name or description. and for the severity, those values look correct since the default order is descending when not specified, which would be ordered alphabetically in descending order (Moderate, Low, Important, Critical). 

### Comment by @swadeley on May 28, 2024 at 02:17 PM UTC

> @swadeley sorting is only available on these fields: issued_date, updated_date, type, and severity. so not for name or description. 

Sorry, I copy pasted `desc` from the example, wich was probably meant to be generic.

>and for the severity, those values look correct since the default order is descending when not specified, which would be ordered **alphabetically** in descending order (Moderate, Low, Important, Critical).

Ahh, only **alphabetical** sort, not on the meaning of the terms. OK.

Thank you



### Comment by @xbhouse on May 28, 2024 at 02:18 PM UTC

oh yes desc in the example is meant to be descending order, not description. sorry for the confusion! 

### Comment by @swadeley on May 28, 2024 at 02:20 PM UTC

@jlsherrill ACK from me, merge when ready. Thank you.

### Comment by @swadeley on May 28, 2024 at 04:39 PM UTC

Thank you

---

## Reviews

### Review by @jlsherrill - Dismissed on May 22, 2024 at 09:11 PM UTC

### Review by @jlsherrill - Commented on May 22, 2024 at 09:13 PM UTC

### Review by @xbhouse - Commented on May 22, 2024 at 09:20 PM UTC

### Review by @jlsherrill - Approved on May 28, 2024 at 02:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/673*
