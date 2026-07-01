---
type: pull_request
number: 878
title: "Fixes 4963: add package summary to rpm presence api"
state: closed
author: jlsherrill
created: 2024-11-07T18:43:39Z
updated: 2024-11-15T18:30:33Z
closed: 2024-11-13T21:16:17Z
base_branch: main
head_branch: 4963
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/878
---

# Pull Request #878: Fixes 4963: add package summary to rpm presence api

**Author**: @jlsherrill
**Created**: November 07, 2024 at 06:43 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4963`

## Description

## Summary


## Testing steps

Create a repo for introspection (such as "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/")


call presence api with a package in it:

```
####
POST http://localhost:8000/api/content-sources/v1.0/rpms/presence
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "urls": [ "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"],
  "rpm_names": ["walrus"]
}


```


see response:

```

{
  "found": [
    "ruby-solv"
  ],
  "missing": [],
  "found_packages": [
    {
      "package_name": "ruby-solv",
      "summary": "Ruby bindings for the libsolv library"
    }
  ]
}
```


---

## Discussion

### Comment by @jlsherrill on November 07, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4963

### Comment by @jlsherrill on November 11, 2024 at 01:03 AM UTC

[test]

### Comment by @swadeley on November 11, 2024 at 10:02 AM UTC

@jlsherrill Please check unit test fail in "build-test".

### Comment by @swadeley on November 13, 2024 at 11:23 AM UTC

Hi @jlsherrill 

```
In [6]: app.content_sources.rest_client.rpms_api.detect_rpm({"urls": ["https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"],
   ...: "rpm_names": ["walrus"],
   ...: })
2024-11-13 11:18:44.823 [    INFO] [iqe.base.rest_client] REST: POST https://ee-izcurwrv.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/presence with query params [] and x-rh-insights-request-id=<>--4815-<>-
Out[6]: 
{'found': ['walrus'],
 'found_packages': [{'package_name': 'walrus',
                     'summary': 'A dummy package of walrus'}],
 'missing': []}

In [7]: 
```

The response I got is the same as in comment 0, just the order is different.

```
In [7]: app.content_sources.rest_client.rpms_api.detect_rpm({"urls": ["https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"],
   ...: "rpm_names": ["walrus","blah"],
   ...: })
2024-11-13 13:48:40.026 [    INFO] [iqe.base.rest_client] REST: POST https://ee-izcurwrv.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/presence with query params [] and x-rh-insights-request-id=<>
Out[7]: 
{'found': ['walrus'],
 'found_packages': [{'package_name': 'walrus',
                     'summary': 'A dummy package of walrus'}],
 'missing': ['blah']}
```

### Comment by @swadeley on November 13, 2024 at 11:36 AM UTC

/retest

### Comment by @swadeley on November 13, 2024 at 02:39 PM UTC

Hi

if you use non-existent URL, you get error like so:

```
HTTP response body: {"errors":[{"status":404,"title":"Error detecting RPMs","detail":"Could not find repository with URL: https://jlsherrill.fedorapeople.org/fake-repos/blah-blah/"}]}
```

but, if you have no repos, you get:

```
ApiTypeError: Invalid type for variable 'found_packages'. Required value type is list and passed type was NoneType at ['received_data']['found_packages']
```

---

## Reviews

### Review by @xbhouse - Approved on November 07, 2024 at 08:44 PM UTC

lgtm!

### Review by @xbhouse - Commented on November 11, 2024 at 01:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/878*
