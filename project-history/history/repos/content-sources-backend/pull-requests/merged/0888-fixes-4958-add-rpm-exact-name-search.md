---
type: pull_request
number: 888
title: "Fixes 4958: add rpm exact name search"
state: merged
author: rverdile
created: 2024-11-13T20:45:21Z
updated: 2024-11-20T14:28:48Z
closed: 2024-11-15T20:16:59Z
merged: 2024-11-15T20:16:59Z
base_branch: main
head_branch: search-exact-names
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/888
---

# Pull Request #888: Fixes 4958: add rpm exact name search

**Author**: @rverdile
**Created**: November 13, 2024 at 08:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `search-exact-names`

## Description

## Summary
Adds new parameter to the request body of name search for rpms, package groups, and environments, "exact_names", to search for content units that exactly match the given names.

Also, the RPM presence API will now show as deprecated in the API spec

## Testing steps
Example request
```
POST http://localhost:8000/api/content-sources/v1.0/rpms/names
Content-Type: application/json
x-Rh-Identity: {{identity-17791560}}

{
  "urls": ["https://rverdile.fedorapeople.org/dummy-repos/comps/repo1/", "https://rverdile.fedorapeople.org/dummy-repos/comps/repo2/"],
  "search": "",
  "exact_names": ["penguin", "bear"]
}
```
1. With ["penguin", "bear"] as the value for "exact_names", both will be returned by the response
2. With ["penguin", "be"] as the value for "exact_names", only penguin will show, because "be" is not an exact match


---

## Discussion

### Comment by @jlsherrill on November 13, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4958

### Comment by @swadeley on November 15, 2024 at 02:05 PM UTC

/retest

### Comment by @swadeley on November 15, 2024 at 07:40 PM UTC

/retest

### Comment by @swadeley on November 15, 2024 at 08:08 PM UTC

Hi

I works:

```
In [3]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"Bea", "exact_names":["bear"]})
2024-11-15 20:06:48.963 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=4c35da60-9f1e-49e1-9796-94c23b8d1bf6
Out[3]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [4]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"", "exact_names":["bear","crow"]})
2024-11-15 20:07:15.419 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=c26e0075-b03e-4819-9238-53e28c22f5ba
Out[4]: 
[{'package_name': 'bear', 'summary': 'A dummy package of bear'},
 {'package_name': 'crow', 'summary': 'A dummy package of crow'}]

In [5]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"", "exact_names":["bear","cro"]})
2024-11-15 20:07:26.403 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=e279f3de-c816-45d8-8bcc-c48d063c329e
Out[5]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [6]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"cro", "exact_names":["bear","cro"]})
2024-11-15 20:07:37.589 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=995f4cdd-ea4a-4ae7-a2d2-ee0784c13fd6
Out[6]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [7]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"cro", "exact_names":["bear"]})
2024-11-15 20:07:47.038 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=7d490a98-e7c3-40ef-8605-af4067736047
Out[7]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [8]: app.content_sources.rest_client.rpms_api.search_rpm({"urls": ["https://stephenw.fedorapeople.org/fakerepos/multiple_errata/"],
   ...: "search":"cro"})
2024-11-15 20:07:56.180 [    INFO] [iqe.base.rest_client] REST: POST https://ee-0zaq6yjn.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/rpms/names with query params [] and x-rh-insights-request-id=a092131d-1a91-4be6-9dab-a4931c9cf6d8
Out[8]: [{'package_name': 'crow', 'summary': 'A dummy package of crow'}]

In [9]: 
```

---

## Reviews

### Review by @jlsherrill - Approved on November 13, 2024 at 09:17 PM UTC

ACK! tested all 3 apis with RHEL content and got expected results

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/888*
