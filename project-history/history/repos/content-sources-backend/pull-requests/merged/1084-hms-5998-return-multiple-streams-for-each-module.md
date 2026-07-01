---
type: pull_request
number: 1084
title: "HMS-5998: return multiple streams for each module"
state: merged
author: jlsherrill
created: 2025-04-17T18:46:33Z
updated: 2025-04-25T16:01:28Z
closed: 2025-04-25T09:44:13Z
merged: 2025-04-25T09:44:13Z
base_branch: main
head_branch: 5998
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1084
---

# Pull Request #1084: HMS-5998: return multiple streams for each module

**Author**: @jlsherrill
**Created**: April 17, 2025 at 06:46 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5998`

## Description

## Summary

Previously modules were being returned with only one stream per module.  In reality we want all the stream but only 1 version per stream.  

## Testing steps
```
make repos-import-rhel9
go cmd/external-repos/main.go snapshot https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/
```

run this query:
```
POST http://localhost:8000/api/content-sources/v1.0/rpms/names
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "urls": ["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"],
  "search": "nginx",
  "include_package_sources": true
}
```

with this change:

```
{
    "package_name": "nginx",
    "summary": "A high performance web server and reverse proxy server",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "x86_64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "x86_64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      }
    ]
  },
```



---

## Discussion

### Comment by @jlsherrill on April 17, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-5998

### Comment by @jlsherrill on April 22, 2025 at 12:23 PM UTC

/retest

### Comment by @jlsherrill on April 22, 2025 at 04:50 PM UTC

/retest

### Comment by @swadeley on April 25, 2025 at 09:44 AM UTC

Hi

In ephemeral I can only do a simple check:

```
In [2]: In [1]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/"], search='walrus',  include_package_sources=True))
2025-04-25 10:27:31.976 [    INFO] [iqe.base.rest_client] REST: POST https://ee-2j9c91iq.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[2]: 
[{'package_name': 'walrus',
  'package_sources': [{'arch': 'x86_64',
                       'context': 'c0ffee42',
                       'description': 'A module for the walrus 0.71 package',
                       'name': 'walrus',
                       'stream': '0.71',
                       'type': 'module',
                       'version': '20180707144203'},
                      {'arch': 'x86_64',
                       'context': 'deadbeef',
                       'description': 'A module for the walrus 5.21 package',
                       'name': 'walrus',
                       'stream': '5.21',
                       'type': 'module',
                       'version': '20180704144203'}],
  'summary': 'A dummy package of walrus'}]
```

In stage, before emerging, I get only  'stream': '1.24' and 'version': '9050020250324055038':

```
In [2]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"], search='nginx',  include_package_sources=True))
2025-04-25 10:38:04.323 [    INFO] [iqe.base.rest_client] REST: POST https://console.stage.redhat.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[2]: 
[{'package_name': 'nginx',
  'package_sources': [{'arch': 'x86_64',
                       'context': '9',
                       'description': 'nginx 1.24 webserver module',
                       'name': 'nginx',
                       'stream': '1.24',
                       'type': 'module',
                       'version': '9050020250324055038'}],
```

### Comment by @swadeley on April 25, 2025 at 10:00 AM UTC

Hi

Testing stage again after merging. I do not get modules:
```
In [1]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"], search='nginx',  include_package_sources=True))
<snip>
2025-04-25 10:56:11.308 [    INFO] [iqe.base.rest_client] REST: POST https://console.stage.redhat.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[1]: 
[{'package_name': 'nginx',
  'package_sources': [{'type': 'package'}],
  'summary': 'A high performance web server and reverse proxy server'},
 {'package_name': 'nginx-all-modules',
  'package_sources': [{'type': 'package'}],
  'summary': 'A meta package that installs all available Nginx modules'},
```

### Comment by @jlsherrill on April 25, 2025 at 04:01 PM UTC

https://github.com/content-services/content-sources-backend/pull/1091  should resolve this issue

---

## Reviews

### Review by @xbhouse - Approved on April 21, 2025 at 05:16 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1084*
