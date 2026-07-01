---
type: pull_request
number: 608
title: "Fixes 2941: Add listSnapshotErrata to api"
state: merged
author: Andrewgdewar
created: 2024-03-15T21:27:14Z
updated: 2024-04-23T17:13:13Z
closed: 2024-04-22T13:50:47Z
merged: 2024-04-22T13:50:47Z
base_branch: main
head_branch: HMS-2941
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/608
---

# Pull Request #608: Fixes 2941: Add listSnapshotErrata to api

**Author**: @Andrewgdewar
**Created**: March 15, 2024 at 09:27 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2941`

## Description

## Summary

- Adds an endpoint to list snapshot errata: `GET /snapshots/{uuid}/errata`
- Support for filtering on name / synopsis, type, severity
- Dependent on this [tang PR](https://github.com/content-services/tang/pull/6)

## Testing steps

- Add a repository with package errata (https://dl.fedoraproject.org/pub/epel/7/x86_64/ is a good one but takes a bit of time) and let it snapshot
- Make a request, will look something like this:
`GET /snapshots/597779df-e5f9-4db1-97e0-ced3074c8beb/errata`
- Can also make a request that includes optional filtering, which will look something like this: `GET /snapshots/597779df-e5f9-4db1-97e0-ced3074c8beb/errata?type=security&severity=Critical&search=epel`
- You can filter on types `security`, `enhancement`, `bugfix`, and `other`
- You can filter on severities `Important`, `Critical`, `Moderate`, `Low`, and `Unknown`
- You can filter on multiple types and severities (`/snapshots/:uuid/errata?type=security,enhancement&severity=Critical,Moderate`)

- Should return an array of results, here's one example:

```
[
     {
            "id": "018ec320-5315-7174-ae8b-fbeff290a222",
            "errata_id": "FEDORA-EPEL-2023-6569f44fa5",
            "title": "moin-1.9.11-1.el7",
            "summary": "moin-1.9.11-1.el7 security update",
            "description": "Security update",
            "issued_date": "2023-01-06 17:28:59",
            "updated_date": "2023-01-15 00:45:13",
            "type": "security",
            "severity": "Critical",
            "reboot_suggested": false
     },
     {
            "id": "018ec31f-a55f-7cf4-b908-1b5f4c2a0384",
            "errata_id": "FEDORA-EPEL-2022-db09048bde",
            "title": "nbd-3.24-1.el7",
            "summary": "nbd-3.24-1.el7 security update",
            "description": "Update to 3.24: fix CVE-2022-26495, CVE-2022-26496\n",
            "issued_date": "2022-03-08 02:53:28",
            "updated_date": "2022-03-16 15:31:09",
            "type": "security",
            "severity": "Critical",
            "reboot_suggested": false
        },
]     
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 15, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-2941

### Comment by @swadeley on April 19, 2024 at 03:27 PM UTC

Hi

configured epel7 repo

with limit=4:
```
In [2]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=4)
2024-04-19 14:27:15.959 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 4)]
Out[2]: 
{'data': [{'description': 'New upstream version',
           'errata_id': 'FEDORA-EPEL-2014-2316',
           'id': '018ef63b-1258-7498-8e12-272423599fb6',
           'issued_date': '2014-08-29 22:51:32',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'pynag-0.9.1-1.el7 bugfix update',
           'title': 'pynag-0.9.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-25 00:20:43'},
          {'description': 'Update to lastest upstream.',
           'errata_id': 'FEDORA-EPEL-2014-2319',
           'id': '018ef63b-8cba-7a4d-9988-3123159549f3',
           'issued_date': '2014-08-30 02:02:16',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'weasyprint-0.22-1.el7 newpackage update',
           'title': 'weasyprint-0.22-1.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:48:49'},
          {'description': 'New package',
           'errata_id': 'FEDORA-EPEL-2014-2320',
           'id': '018ef63b-6cc5-7b5a-a3fb-197df3069d45',
           'issued_date': '2014-08-30 05:22:39',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'ckermit-9.0.302-7.el7 newpackage update',
           'title': 'ckermit-9.0.302-7.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:47:18'},
          {'description': 'mod_xsendfile is a small Apache2 module that '
                          'processes X-SENDFILE headers registered by the '
                          'original output handler.\r\n'
                          '\r\n'
                          'If it encounters the presence of such header it '
                          'will discard all output and send the file specified '
                          'by that header instead using Apache internals '
                          'including all optimizations like caching-headers '
                          'and sendfile or mmap if configured.\r\n'
                          '\r\n'
                          'It is useful for processing script-output of e.g. '
                          'php, perl or any cgi.\r\n',
           'errata_id': 'FEDORA-EPEL-2014-2321',
           'id': '018ef63b-61b8-7374-85ea-f56215981022',
           'issued_date': '2014-08-29 19:21:23',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'mod_xsendfile-0.12-10.el7 newpackage update',
           'title': 'mod_xsendfile-0.12-10.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:45:49'}],
 'links': {'first': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=4&offset=0',
           'last': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=4&offset=4824',
           'next': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=4&offset=4'},
 'meta': {'count': 4825, 'limit': 4, 'offset': 0}}

```

I cannot match the search example in comment 0:
```
In [5]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=2,offset=1,search="newpackage")
2024-04-19 16:03:23.752 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 1), ('search', 'newpackage')]
Out[5]: 
{'data': [{'description': 'New package',
           'errata_id': 'FEDORA-EPEL-2014-2320',
           'id': '018ef63b-6cc5-7b5a-a3fb-197df3069d45',
           'issued_date': '2014-08-30 05:22:39',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'ckermit-9.0.302-7.el7 newpackage update',
           'title': 'ckermit-9.0.302-7.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:47:18'},
          {'description': 'mod_xsendfile is a small Apache2 module that '
                          'processes X-SENDFILE headers registered by the '
                          'original output handler.\r\n'
                          '\r\n'
                          'If it encounters the presence of such header it '
                          'will discard all output and send the file specified '
                          'by that header instead using Apache internals '
                          'including all optimizations like caching-headers '
                          'and sendfile or mmap if configured.\r\n'
                          '\r\n'
                          'It is useful for processing script-output of e.g. '
                          'php, perl or any cgi.\r\n',
           'errata_id': 'FEDORA-EPEL-2014-2321',
           'id': '018ef63b-61b8-7374-85ea-f56215981022',
           'issued_date': '2014-08-29 19:21:23',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'mod_xsendfile-0.12-10.el7 newpackage update',
           'title': 'mod_xsendfile-0.12-10.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:45:49'}],
 'links': {'first': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=2&offset=0&search=newpackage',
           'last': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=2&offset=1714&search=newpackage',
           'next': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=2&offset=3&search=newpackage'},
 'meta': {'count': 1715, 'limit': 2, 'offset': 1}}
```

but simple string search on name of package or title of errata works:

```In [19]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=99,offset=0,search='FEDORA-EPEL-2014-2316')
2024-04-19 16:24:33.262 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 99), ('offset', 0), ('search', 'FEDORA-EPEL-2014-2316')]
Out[19]: 
{'data': [{'description': 'New upstream version',
           'errata_id': 'FEDORA-EPEL-2014-2316',
           'id': '018ef63b-1258-7498-8e12-272423599fb6',
           'issued_date': '2014-08-29 22:51:32',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'pynag-0.9.1-1.el7 bugfix update',
           'title': 'pynag-0.9.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-25 00:20:43'}],
 'links': {'first': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=99&offset=0&search=FEDORA-EPEL-2014-2316',
           'last': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=99&offset=0&search=FEDORA-EPEL-2014-2316'},
 'meta': {'count': 1, 'limit': 99, 'offset': 0}}

In [20]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=99,offset=0,search='pynag-0.9.1-1.el7')
2024-04-19 16:25:44.966 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 99), ('offset', 0), ('search', 'pynag-0.9.1-1.el7')]
Out[20]: 
{'data': [{'description': 'New upstream version',
           'errata_id': 'FEDORA-EPEL-2014-2316',
           'id': '018ef63b-1258-7498-8e12-272423599fb6',
           'issued_date': '2014-08-29 22:51:32',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'pynag-0.9.1-1.el7 bugfix update',
           'title': 'pynag-0.9.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-25 00:20:43'}],
 'links': {'first': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=99&offset=0&search=pynag-0.9.1-1.el7',
           'last': '/api/content-sources/v1/snapshots/3e72fb36-acab-43d3-83a7-a907f5c860d9/errata?limit=99&offset=0&search=pynag-0.9.1-1.el7'},
 'meta': {'count': 1, 'limit': 99, 'offset': 0}}

```

### Comment by @xbhouse on April 19, 2024 at 03:34 PM UTC

@swadeley if i understand correctly i think you'd want to filter on `type` instead of `search` when looking for type of `security`, `enhancement`, or `bugfix`. to find type of `newpackage`, since we don't filter on this in the UI, you can filter on type `other` to get any types that are not `security`, `enhancement`, or `bugfix`. so i think the request would look something like this to get all eratta of type `newpackage`:

app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=2,offset=1,type="other")


### Comment by @swadeley on April 19, 2024 at 03:40 PM UTC

@xbhouse I did try things like:

```
In [60]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('3e72fb36-acab-43d3-83a7-a907f5c860d9',limit=2,offset=1,type='newpackage')
---------------------------------------------------------------------------
ApiTypeError
```

### Comment by @xbhouse on April 19, 2024 at 03:41 PM UTC

just edited my comment ^

### Comment by @xbhouse on April 19, 2024 at 03:42 PM UTC

@swadeley sorry, i should have added better details in the description. adding them now

### Comment by @xbhouse on April 19, 2024 at 04:06 PM UTC

@swadeley you can also pass in multiple types / severities, so you may need to change that request to take an array

EDIT: nope it is a string, api docs were missing the parameters :) 

### Comment by @swadeley on April 19, 2024 at 04:35 PM UTC

Hi

```
In [6]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('f9240c46-f340-48f4-af37-cc9ef4b285be',limit=4,offset=0,type='other')

ApiTypeError: Got an unexpected parameter 'type' to method `list_snapshot_errata`
```

```
In [9]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('f9240c46-f340-48f4-af37-cc9ef4b285be',limit=4,offset=0,severity="")
ApiTypeError: Got an unexpected parameter 'severity' to method `list_snapshot_errata`

```


### Comment by @xbhouse on April 19, 2024 at 04:44 PM UTC

ohhhh wow i see the api docs actually are missing type and severity:

                "parameters": [
                    {
                        "type": "string",
                        "description": "Snapshot ID.",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of items to include in response. Use it to control the number of items, particularly when dealing with large datasets. Default value: ` + "`" + `100` + "`" + `.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Starting point for retrieving a subset of results. Determines how many items to skip from the beginning of the result set. Default value:` + "`" + `0` + "`" + `.",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Term to filter and retrieve items that match the specified search criteria. Search term can include name.",
                        "name": "search",
                        "in": "query"
                    }
                ],
                
                
let me fix this 

### Comment by @swadeley on April 19, 2024 at 06:21 PM UTC

Hi @xbhouse 

Now I get:

`ApiTypeError: Invalid type for variable 'updated_date'. Required value type is str and passed type was NoneType at ['received_data']['data'][0]['updated_date']`

EDIT: testing with `https://stephenw.fedorapeople.org/fakerepos/multiple_errata/` this time

### Comment by @xbhouse on April 19, 2024 at 06:25 PM UTC

@swadeley ooof i just saw this too when testing with your repo in my iqe env, let me try adding `omitempty` to the go docs

### Comment by @swadeley on April 19, 2024 at 06:42 PM UTC

Hi

severity
working well with epel repo:

```
In [12]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,severity="Moderate")
2024-04-19 19:37:34.299 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('severity', 'Moderate')]
Out[12]: 
{'data': [{'description': 'This adjusts gwenhywfar to use the system copy of '
                          'ca-certificates, instead of a bundled copy.',
           'errata_id': 'FEDORA-EPEL-2015-0272adfe4b',
           'id': '018ef7a2-152f-7dc5-b881-4d12606c917d',
           'issued_date': '2015-12-12 04:45:25',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': 'gwenhywfar-4.13.1-2.el7 security update',
           'title': 'gwenhywfar-4.13.1-2.el7',
           'type': 'security',
           'updated_date': '2015-12-30 17:33:18'},
          {'description': 'Fix runtime dependencies for el7.2',
           'errata_id': 'FEDORA-EPEL-2015-06f5c58464',
           'id': '018ef7a1-a377-7646-b8e1-db76cb1e3f9f',
           'issued_date': '2015-12-18 04:44:48',
           'reboot_suggested': False,
           'severity': 'Moderate',
           'summary': 'kmahjongg-4.10.5-2.el7 bugfix update',
           'title': 'kmahjongg-4.10.5-2.el7',
           'type': 'bugfix',
           'updated_date': '2016-01-04 20:55:07'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=Moderate',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=358&severity=Moderate',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&severity=Moderate'},
 'meta': {'count': 360, 'limit': 2, 'offset': 0}}

In [13]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,severity="High")
2024-04-19 19:37:42.701 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('severity', 'High')]
Out[13]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=High',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=High'},
 'meta': {'count': 0, 'limit': 2, 'offset': 0}}

In [14]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,severity="Critical")
2024-04-19 19:37:51.312 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('severity', 'Critical')]
Out[14]: 
{'data': [{'description': 'Update package to support newer version of npm(d), '
                          'to fix unresolved dependencies',
           'errata_id': 'FEDORA-EPEL-2016-298d867ed8',
           'id': '018ef7a0-648f-77f0-af68-fcf3c5bcc737',
           'issued_date': '2016-01-12 15:33:06',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': 'nodejs-es6-iterator-2.0.0-2.el7 bugfix update',
           'title': 'nodejs-es6-iterator-2.0.0-2.el7',
           'type': 'bugfix',
           'updated_date': '2016-01-27 15:18:08'},
          {'description': 'Allow newer version of npm(d), to fix unresolved '
                          'dependencies in the repo',
           'errata_id': 'FEDORA-EPEL-2016-cc40ad47de',
           'id': '018ef7a0-5ea8-7265-81cb-f435263992eb',
           'issued_date': '2016-01-12 15:35:56',
           'reboot_suggested': False,
           'severity': 'Critical',
           'summary': 'nodejs-es6-symbol-3.0.2-2.el7 bugfix update',
           'title': 'nodejs-es6-symbol-3.0.2-2.el7',
           'type': 'bugfix',
           'updated_date': '2016-01-27 15:18:08'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=Critical',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=8&severity=Critical',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&severity=Critical'},
 'meta': {'count': 10, 'limit': 2, 'offset': 0}}

In [15]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,severity="Low")
2024-04-19 19:37:59.900 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('severity', 'Low')]
Out[15]: 
{'data': [{'description': 'File::Type uses magic numbers (typically at the '
                          'start of a file) to determine the MIME type of that '
                          'file.',
           'errata_id': 'FEDORA-EPEL-2015-1c187681bb',
           'id': '018ef7a0-a64b-7f73-9221-17af476544cd',
           'issued_date': '2015-12-19 00:19:09',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': 'perl-File-Type-0.22-23.el7 newpackage update',
           'title': 'perl-File-Type-0.22-23.el7',
           'type': 'newpackage',
           'updated_date': '2016-01-04 20:55:07'},
          {'description': '\n'
                          '\n'
                          'perl-BZ-Client-1.04-12.el7\n'
                          '\n'
                          '- Perl 5.20 rebuild\n',
           'errata_id': 'FEDORA-EPEL-2015-2e6e8e31a7',
           'id': '018ef7a0-d703-7413-9588-3a76599958d2',
           'issued_date': '2015-11-16 14:47:27',
           'reboot_suggested': False,
           'severity': 'Low',
           'summary': 'perl-BZ-Client-1.04-12.el7 newpackage update',
           'title': 'perl-BZ-Client-1.04-12.el7',
           'type': 'newpackage',
           'updated_date': '2015-12-03 03:29:46'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=Low',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=460&severity=Low',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&severity=Low'},
 'meta': {'count': 461, 'limit': 2, 'offset': 0}}

In [16]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,severity="Unknown")
2024-04-19 19:38:23.055 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('severity', 'Unknown')]
Out[16]: 
{'data': [{'description': 'New upstream version',
           'errata_id': 'FEDORA-EPEL-2014-2316',
           'id': '018ef7a0-acbc-731b-b3e2-c2984291837a',
           'issued_date': '2014-08-29 22:51:32',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'pynag-0.9.1-1.el7 bugfix update',
           'title': 'pynag-0.9.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-25 00:20:43'},
          {'description': 'Update to lastest upstream.',
           'errata_id': 'FEDORA-EPEL-2014-2319',
           'id': '018ef7a1-389e-792d-938a-29d3a6fb4d82',
           'issued_date': '2014-08-30 02:02:16',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'weasyprint-0.22-1.el7 newpackage update',
           'title': 'weasyprint-0.22-1.el7',
           'type': 'newpackage',
           'updated_date': '2014-09-24 03:48:49'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&severity=Unknown',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=3918&severity=Unknown',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&severity=Unknown'},
 'meta': {'count': 3919, 'limit': 2, 'offset': 0}}
```



### Comment by @swadeley on April 19, 2024 at 06:53 PM UTC

and for type:

```
In [18]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,type="bugfix")
2024-04-19 19:50:18.804 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('type', 'bugfix')]
Out[18]: 
{'data': [{'description': 'New upstream version',
           'errata_id': 'FEDORA-EPEL-2014-2316',
           'id': '018ef7a0-acbc-731b-b3e2-c2984291837a',
           'issued_date': '2014-08-29 22:51:32',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'pynag-0.9.1-1.el7 bugfix update',
           'title': 'pynag-0.9.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-25 00:20:43'},
          {'description': 'Update to official release, instead of git '
                          'snapshot. No change.',
           'errata_id': 'FEDORA-EPEL-2014-2350',
           'id': '018ef7a1-d971-793a-80e2-16ce5a73c244',
           'issued_date': '2014-08-30 09:53:01',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'php-pecl-igbinary-1.2.1-1.el7 bugfix update',
           'title': 'php-pecl-igbinary-1.2.1-1.el7',
           'type': 'bugfix',
           'updated_date': '2014-09-24 03:41:46'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&type=bugfix',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=1416&type=bugfix',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&type=bugfix'},
 'meta': {'count': 1417, 'limit': 2, 'offset': 0}}

In [19]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,type="security")
2024-04-19 19:50:30.102 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('type', 'security')]
Out[19]: 
{'data': [{'description': 'The qs module has the ability to create sparse '
                          'arrays during parsing. By specifying a high index '
                          'it is possible to create a large array that will '
                          'eventually take up all the allocated memory of the '
                          'running process, resulting in a crash.\r\n'
                          '\r\n'
                          'More information: '
                          'https://github.com/visionmedia/node-querystring/issues/104\r\n',
           'errata_id': 'FEDORA-EPEL-2014-2861',
           'id': '018ef7a0-641f-7298-87f9-27b34bd59ebf',
           'issued_date': '2014-09-24 12:21:50',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'nodejs-qs-0.6.6-3.el7 security update',
           'title': 'nodejs-qs-0.6.6-3.el7',
           'type': 'security',
           'updated_date': '2014-10-31 01:22:25'},
          {'description': 'When relying on the root option to restrict file '
                          'access it may be possible for an application '
                          'consumer to escape out of the restricted directory '
                          'and access files in a similarly named directory. '
                          "For example, static(_dirname + '/public') would "
                          "allow access to _dirname + '/public-restricted'.\r\n"
                          '\r\n'
                          'https://nodesecurity.io/advisories/send-directory-traversal',
           'errata_id': 'FEDORA-EPEL-2014-2870',
           'id': '018ef7a0-fd46-7ef5-86a6-da2a465e0224',
           'issued_date': '2014-09-24 12:52:19',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'nodejs-send-0.3.0-4.el7 security update',
           'title': 'nodejs-send-0.3.0-4.el7',
           'type': 'security',
           'updated_date': '2014-10-31 01:25:17'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&type=security',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=298&type=security',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&type=security'},
 'meta': {'count': 299, 'limit': 2, 'offset': 0}}

In [20]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9162a11b-7918-4f39-9a54-ce29bb8a6471',limit=2,offset=0,type="enhancement")
2024-04-19 19:50:55.501 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('limit', 2), ('offset', 0), ('type', 'enhancement')]
Out[20]: 
{'data': [{'description': '- Latest upstream',
           'errata_id': 'FEDORA-EPEL-2014-2349',
           'id': '018ef7a2-453f-7932-a5ed-2b89bce27a7e',
           'issued_date': '2014-09-01 10:23:03',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'python-taskflow-0.3.21-1.el7 enhancement update',
           'title': 'python-taskflow-0.3.21-1.el7',
           'type': 'enhancement',
           'updated_date': '2014-09-24 03:43:33'},
          {'description': 'Added new module for managing OpenStack flavours: '
                          'nova_flavor New package New package',
           'errata_id': 'FEDORA-EPEL-2014-2392',
           'id': '018ef7a1-3961-7482-806e-503eb3bccd20',
           'issued_date': '2014-09-02 09:59:21',
           'reboot_suggested': False,
           'severity': 'None',
           'summary': 'ansible-openstack-modules-0-20140902git79d751a.el7 '
                      'enhancement update',
           'title': 'ansible-openstack-modules-0-20140902git79d751a.el7',
           'type': 'enhancement',
           'updated_date': '2015-02-17 19:01:48'}],
 'links': {'first': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=0&type=enhancement',
           'last': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=1242&type=enhancement',
           'next': '/api/content-sources/v1/snapshots/9162a11b-7918-4f39-9a54-ce29bb8a6471/errata?limit=2&offset=2&type=enhancement'},
 'meta': {'count': 1244, 'limit': 2, 'offset': 0}}

In [21]: 

```

### Comment by @swadeley on April 19, 2024 at 07:09 PM UTC

> @swadeley ooof i just saw this too when testing with your repo in my iqe env, let me try adding `omitempty` to the go docs

Hi @xbhouse , I see your update to `pkg/api/rpms.go`, but I think we need corresponding change in `api/openapi.json`.

### Comment by @xbhouse on April 19, 2024 at 09:42 PM UTC

@swadeley ok with the latest changes, package errata with no updated date can be fetched (this is using your repo):
```
In [4]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('4f14f268-1185-4eba-b97e-2cecc8793217')
2024-04-19 17:39:20.975 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=zdOfzpgaoKESdzLXgIUJMsfbTcshLsAx, params=[]
Out[4]: 
{'data': [{'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018ef788-b174-7887-9b75-47ae78233bfd',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': 'low',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018ef788-b171-7e40-9cc8-b7e692b8f9be',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018ef788-b164-7ad6-b4ed-ef8c0c8c8de5',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018ef788-b169-770c-9173-f13a1c7e3126',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018ef788-b16e-78a1-b6e1-07f202d6f733',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/4f14f268-1185-4eba-b97e-2cecc8793217/errata?limit=100&offset=0',
           'last': '/api/content-sources/v1/snapshots/4f14f268-1185-4eba-b97e-2cecc8793217/errata?limit=100&offset=0'},
 'meta': {'count': 5, 'limit': 100, 'offset': 0}}
 ```
 thanks for finding all these issues and for your patience :) 


### Comment by @swadeley on April 20, 2024 at 07:29 AM UTC

Hi @xbhouse 

thanks for the updates. Those recent changes do not change the API docs file, so now I realise I need to build a new image and redeploy.

thank you

### Comment by @swadeley on April 20, 2024 at 07:29 AM UTC

/retest

### Comment by @swadeley on April 20, 2024 at 05:23 PM UTC

Hi

testing with:  `https://stephenw.fedorapeople.org/fakerepos/multiple_errata/`

I could not get this:
<description>Sea_Erratum</description>
<severity>Critical</severity>

with:
`In [6]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9454fb71-fc55-4aa8-80d9-7d45c869ad80',limit=2,offset=0,severity="Critical")
`

I see the APi downlands this (severity is all missing):
```
In [13]: app.content_sources.rest_client.snapshots_api.list_snapshot_errata('9454fb71-fc55-4aa8-80d9-7d45c869ad80')

Out[13]: 
{'data': [{'description': 'Kangaroo_Erratum',
           'errata_id': 'RHBA-2012:1030',
           'id': '018efc40-6710-7a95-b5d8-eff8229e7476',
           'issued_date': '2012-01-08 10:40:38',
           'reboot_suggested': False,
           'severity': 'low',
           'summary': '',
           'title': 'Kangaroo_Erratum',
           'type': 'bugfix',
           'updated_date': ''},
          {'description': 'Gorilla_Erratum',
           'errata_id': 'RHEA-2012:0058',
           'id': '018efc40-670d-70b9-901f-6788e4103af7',
           'issued_date': '2013-01-27 16:08:09',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Gorilla_Erratum',
           'type': 'enhancement',
           'updated_date': ''},
          {'description': 'Sea_Erratum',
           'errata_id': 'RHSA-2012:0055',
           'id': '018efc40-6700-79db-9889-0669ec4a4828',
           'issued_date': '2012-01-27 16:08:06',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Sea_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bird_Erratum',
           'errata_id': 'RHSA-2012:0056',
           'id': '018efc40-6705-7dab-a17f-6d6d1cbb60ef',
           'issued_date': '2013-01-27 16:08:08',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Bird_Erratum',
           'type': 'security',
           'updated_date': ''},
          {'description': 'Bear_Erratum',
           'errata_id': 'RHSA-2012:0057',
           'id': '018efc40-670a-7152-8dd6-95965fc28a59',
           'issued_date': '2013-01-27 16:08:05',
           'reboot_suggested': False,
           'severity': '',
           'summary': '',
           'title': 'Bear_Erratum',
           'type': 'security',
           'updated_date': ''}],
 'links': {'first': '/api/content-sources/v1/snapshots/9454fb71-fc55-4aa8-80d9-7d45c869ad80/errata?limit=100&offset=0',
           'last': '/api/content-sources/v1/snapshots/9454fb71-fc55-4aa8-80d9-7d45c869ad80/errata?limit=100&offset=0'},
 'meta': {'count': 5, 'limit': 100, 'offset': 0}}
```

I am happy to merge this and raise new bug as your changes work with epel repo.
Let me know.

Thank you

EDIT: Test repo was confirmed to be at fault, now fixed.

### Comment by @swadeley on April 22, 2024 at 08:56 AM UTC

/retest

---

## Reviews

### Review by @swadeley - Commented on April 12, 2024 at 06:26 PM UTC

### Review by @swadeley - Commented on April 12, 2024 at 06:27 PM UTC

### Review by @swadeley - Commented on April 12, 2024 at 06:40 PM UTC

### Review by @swadeley - Commented on April 12, 2024 at 06:40 PM UTC

### Review by @swadeley - Commented on April 12, 2024 at 06:40 PM UTC

### Review by @swadeley - Commented on April 12, 2024 at 06:43 PM UTC

### Review by @rverdile - Commented on April 16, 2024 at 07:47 PM UTC

### Review by @rverdile - Commented on April 16, 2024 at 07:47 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 01:55 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 01:57 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 02:03 PM UTC

### Review by @rverdile - Commented on April 17, 2024 at 02:07 PM UTC

### Review by @rverdile - Commented on April 17, 2024 at 02:08 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 02:10 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 02:15 PM UTC

### Review by @rverdile - Commented on April 17, 2024 at 06:22 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 06:30 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 07:46 PM UTC

### Review by @rverdile - Commented on April 18, 2024 at 12:52 PM UTC

### Review by @rverdile - Approved on April 18, 2024 at 04:05 PM UTC

lgtm!

### Review by @swadeley - Commented on April 19, 2024 at 09:02 AM UTC

### Review by @rverdile - Commented on April 19, 2024 at 07:19 PM UTC

### Review by @rverdile - Commented on April 19, 2024 at 07:21 PM UTC

### Review by @xbhouse - Commented on April 19, 2024 at 07:30 PM UTC

### Review by @rverdile - Commented on April 19, 2024 at 07:37 PM UTC

### Review by @rverdile - Commented on April 19, 2024 at 07:50 PM UTC

### Review by @xbhouse - Commented on April 19, 2024 at 09:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/608*
