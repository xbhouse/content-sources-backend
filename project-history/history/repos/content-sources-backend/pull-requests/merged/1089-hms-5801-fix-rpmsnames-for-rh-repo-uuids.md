---
type: pull_request
number: 1089
title: "HMS-5801: fix /rpms/names for RH repo uuids"
state: merged
author: xbhouse
created: 2025-04-23T19:01:27Z
updated: 2025-04-28T11:45:02Z
closed: 2025-04-28T11:45:02Z
merged: 2025-04-28T11:45:02Z
base_branch: main
head_branch: 5801
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1089
---

# Pull Request #1089: HMS-5801: fix /rpms/names for RH repo uuids

**Author**: @xbhouse
**Created**: April 23, 2025 at 07:01 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5801`

## Description

## Summary

Adds RH org to the query to fix searching for rpms by repo UUID

## Testing steps

1. Import RH repos and introspect at least one: `make repos-import-rhel9 && go run cmd/external-repos/main.go introspect https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/`
2. Searching rpms for that repo by UUID should work 

```
POST /rpms/names

{
  "search": "nodejs",
  "uuids": ["1d4d6b01-d5fd-466a-a56f-1863e0ce4d24"]
}
```
```
[
    {
        "package_name": "nodejs",
        "summary": "JavaScript runtime"
    },
    {
        "package_name": "nodejs-devel",
        "summary": "JavaScript runtime - development headers"
    },
    ...
]
``` 

---

## Discussion

### Comment by @jlsherrill on April 23, 2025 at 07:06 PM UTC

https://issues.redhat.com/browse/HMS-5801

### Comment by @swadeley on April 28, 2025 at 10:39 AM UTC

/retest

### Comment by @swadeley on April 28, 2025 at 11:44 AM UTC

Hi,

```
In [8]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://cdn.redhat.com/content/dist/rhel9/9/aarch64/codeready-builder/os/"], search='nginx',  include_package_sources=True))
2025-04-28 12:39:14.827 [    INFO] [iqe.base.rest_client] REST: POST https://ee-8iffymeu.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[8]: 
[{'package_name': 'nginx-mod-devel',
  'package_sources': [{'type': 'package'}],
  'summary': 'Nginx module development files'}]



In [10]: app.content_sources.rest_client.rpms_api.search_rpm(dict(uuids=[repo_uuid], search='nginx',  include_package_sources=True))
2025-04-28 12:42:19.968 [    INFO] [iqe.base.rest_client] REST: POST https://ee-8iffymeu.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=<>
Out[10]: 
[{'package_name': 'nginx-mod-devel',
  'package_sources': [{'type': 'package'}],
  'summary': 'Nginx module development files'}]

```

lgtm

---

## Reviews

### Review by @rverdile - Approved on April 24, 2025 at 06:16 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1089*
