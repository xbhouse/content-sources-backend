---
type: pull_request
number: 1045
title: "HMS-5640: add module info to rpm search api"
state: merged
author: xbhouse
created: 2025-03-27T15:23:17Z
updated: 2025-04-02T17:16:25Z
closed: 2025-04-02T17:16:25Z
merged: 2025-04-02T17:16:25Z
base_branch: main
head_branch: 5640
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1045
---

# Pull Request #1045: HMS-5640: add module info to rpm search api

**Author**: @xbhouse
**Created**: March 27, 2025 at 03:23 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5640`

## Description

## Summary

- Extends the rpms search endpoints (`/rpm/names` and `/snapshots/rpms/names`) with module information if requested
- The case where rpms exist both inside and outside of a module is not handled in this PR. That will be achieved by a couple additional tasks: https://issues.redhat.com/browse/HMS-5802, https://issues.redhat.com/browse/HMS-5803

## Testing steps

1. Add a couple repos with modules and let them snapshot (including a RH one to test searching those rpms). You'll need the snapshot UUIDs when hitting the `/snapshots/rpms/names` endpoint
2. Use the `/rpms/names` endpoint to search for rpms in the RH repo (using the URL, there is a [bug](https://issues.redhat.com/browse/HMS-5801) when searching by UUID 😅 ) and set `include_package_sources` to `true` to include module info:
```
{
  "search": "apache",
  "uuids": [],
  "urls": ["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"],
  "include_package_sources": true
}
```
3. With this example request, you should see the apache* packages returned, some included in a module and some that are outside. If they are inside a module, you should see a response like this:

```
{
        "package_name": "apache-commons-cli",
        "summary": "Command Line Interface Library for Java",
        "package_sources": [
            {
                "type": "module",
                "name": "maven",
                "stream": "3.8",
                "context": "470dcefd",
                "arch": "x86_64",
                "version": "9010020220803061628",
                "description": "Maven is a software project management and comprehension tool. Based on the concept of a project object model (POM), Maven can manage a project's build, reporting and documentation from a central piece of information."
            },
...  
```

If the package is outside a module, the response should look like this:
```
{
        "package_name": "apache-commons-logging",
        "summary": "Apache Commons Logging",
        "package_sources": [
            {
                "type": "package"
            }
        ]
},
```
4. Do the same for custom repos and with the `/snapshots/rpms/names` endpoint (using snapshot UUIDs). You should see similar responses

---

## Discussion

### Comment by @jlsherrill on March 27, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-5640

### Comment by @jlsherrill on April 01, 2025 at 01:32 PM UTC

one small comment/question but otherwise looks and runs great!

### Comment by @swadeley on April 02, 2025 at 02:50 PM UTC

Hi

I built API docs (MR 619)

Using this for test repo: `https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/`

```
In [20]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms('5885628c-265e-4c1b-af69-04e0c208b17c', search="kangaroo")
2025-04-02 15:41:28.922 [    INFO] [iqe.base.rest_client] REST: GET https://ee-q3oo02bc.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms with query params [('search', 'kangaroo')] and x-rh-insights-request-id=4c1fc0f8-8cc9-4268-8c4d-8128756a4822
Out[20]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'kangaroo',
           'release': '1',
           'summary': 'A dummy package of kangaroo',
           'version': '0.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'kangaroo',
           'release': '1',
           'summary': 'hop like a kangaroo in Australia',
           'version': '0.3'}],
 'links': {'first': '/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms?limit=100&offset=0&search=kangaroo',
           'last': '/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms?limit=100&offset=0&search=kangaroo'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [21]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms('5885628c-265e-4c1b-af69-04e0c208b17c', search="kangaroo", include_package_sources=true)
---------------------------------------------------------------------------
NameError                                 Traceback (most recent call last)
<snip>
ApiTypeError: Got an unexpected parameter 'include_package_sources' to method `list_snapshot_rpms`

In [23]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms.attribute_map
Out[23]: {'uuid': 'uuid', 'limit': 'limit', 'offset': 'offset', 'search': 'search'}
```

What have I done wrong? Maybe forgot to add that in IQE ? I tried adding that value, but that only helps if its returned in a response. This error is on trying to search, so its before I think.

### Comment by @xbhouse on April 02, 2025 at 03:14 PM UTC

> Hi
> 
> I built API docs (MR 619)
> 
> Using this for test repo: `https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/`
> 
> ```
> In [20]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms('5885628c-265e-4c1b-af69-04e0c208b17c', search="kangaroo")
> 2025-04-02 15:41:28.922 [    INFO] [iqe.base.rest_client] REST: GET https://ee-q3oo02bc.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms with query params [('search', 'kangaroo')] and x-rh-insights-request-id=4c1fc0f8-8cc9-4268-8c4d-8128756a4822
> Out[20]: 
> {'data': [{'arch': 'noarch',
>            'epoch': '0',
>            'name': 'kangaroo',
>            'release': '1',
>            'summary': 'A dummy package of kangaroo',
>            'version': '0.2'},
>           {'arch': 'noarch',
>            'epoch': '0',
>            'name': 'kangaroo',
>            'release': '1',
>            'summary': 'hop like a kangaroo in Australia',
>            'version': '0.3'}],
>  'links': {'first': '/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms?limit=100&offset=0&search=kangaroo',
>            'last': '/api/content-sources/v1/snapshots/5885628c-265e-4c1b-af69-04e0c208b17c/rpms?limit=100&offset=0&search=kangaroo'},
>  'meta': {'count': 2, 'limit': 100, 'offset': 0}}
> 
> In [21]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms('5885628c-265e-4c1b-af69-04e0c208b17c', search="kangaroo", include_package_sources=true)
> ---------------------------------------------------------------------------
> NameError                                 Traceback (most recent call last)
> <snip>
> ApiTypeError: Got an unexpected parameter 'include_package_sources' to method `list_snapshot_rpms`
> 
> In [23]: app.content_sources.rest_client.rpms_api.list_snapshot_rpms.attribute_map
> Out[23]: {'uuid': 'uuid', 'limit': 'limit', 'offset': 'offset', 'search': 'search'}
> ```
> 
> What have I done wrong? Maybe forgot to add that in IQE ? I tried adding that value, but that only helps if its returned in a response. This error is on trying to search, so its before I think.

hi @swadeley, these updates are for the search endpoints, so i think the endpoints in IQE would be:
```
app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=[<repo_url>], search='abc',  include_package_sources=True))
app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=[<snapshot_uuid>], search='abc',  include_package_sources=True))
```


### Comment by @swadeley on April 02, 2025 at 05:08 PM UTC

Thank you @xbhouse 

Looks good:
```
In [6]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/"], search='kangaroo',  include_package_sources=True))
2025-04-02 18:06:00.742 [    INFO] [iqe.base.rest_client] REST: POST https://ee-q3oo02bc.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[6]: 
[{'package_name': 'kangaroo',
  'package_sources': [{'arch': 'noarch',
                       'context': 'deadbeef',
                       'description': 'A module for the kangaroo 0.2 package',
                       'name': 'kangaroo',
                       'stream': '0',
                       'type': 'module',
                       'version': '20180704111719'},
                      {'arch': 'noarch',
                       'context': 'deadbeef',
                       'description': 'A module for the kangaroo 0.3 package',
                       'name': 'kangaroo',
                       'stream': '0',
                       'type': 'module',
                       'version': '20180730223407'}],
  'summary': 'hop like a kangaroo in Australia'}]

In [7]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/"], search='kangaroo'))
2025-04-02 18:06:22.214 [    INFO] [iqe.base.rest_client] REST: POST https://ee-q3oo02bc.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[7]: [{'package_name': 'kangaroo', 'summary': 'hop like a kangaroo in Australia'}]

In [8]: 
```

---

## Reviews

### Review by @jlsherrill - Commented on April 01, 2025 at 01:30 PM UTC

### Review by @xbhouse - Commented on April 01, 2025 at 01:56 PM UTC

### Review by @jlsherrill - Approved on April 01, 2025 at 09:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1045*
