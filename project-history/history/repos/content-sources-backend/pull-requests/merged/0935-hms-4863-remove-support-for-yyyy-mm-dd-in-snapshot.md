---
type: pull_request
number: 935
title: "HMS-4863: remove support for YYYY-MM-DD in snapshots/for_date"
state: merged
author: xbhouse
created: 2025-01-09T20:44:05Z
updated: 2025-01-13T14:10:03Z
closed: 2025-01-13T13:27:00Z
merged: 2025-01-13T13:27:00Z
base_branch: main
head_branch: 4863
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/935
---

# Pull Request #935: HMS-4863: remove support for YYYY-MM-DD in snapshots/for_date

**Author**: @xbhouse
**Created**: January 09, 2025 at 08:44 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `4863`

## Description

## Summary

- Changes the `/snapshots/for_date/` request to remove support for YYYY-MM-DD date formats
- Before this PR, we supported both RFC3339 and YYYY-MM-DD formats to maintain backward compatibility with Image Builder
- The needed changes in IB were made here:
    - UI: https://github.com/osbuild/image-builder-frontend/pull/2526
    - Backend: https://github.com/osbuild/image-builder/pull/1252, https://github.com/osbuild/image-builder/pull/1429
- To be safe, don't merge on a Friday in case anything was missed

## Testing steps

1. Try to make a request to /snapshots/for_date/ with a YYYY-MM-DD date, this should fail:
```
{
  "date": "2025-01-09",
  "repository_uuids": ["edb66d0f-1d44-498a-b55e-960183c42cff"]
}
```
2. A request like this should still work:
```
{
  "date": "2025-01-09T00:00:00Z",
  "repository_uuids": ["edb66d0f-1d44-498a-b55e-960183c42cff"]
}
```



---

## Discussion

### Comment by @jlsherrill on January 09, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4863

### Comment by @mayurilahane on January 10, 2025 at 01:25 AM UTC

This is failing as expected:

{
  "date": "2025-01-09",
  "repository_uuids": ["edb66d0f-1d44-498a-b55e-960183c42cff"]
}


```ApiException: (400)
Reason: Bad Request
HTTP response headers: HTTPHeaderDict({'content-length': '280', 'content-type': 'application/json', 'date': 'Fri, 10 Jan 2025 01:22:18 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZI6ZmFsc2V9fX0=', 'x-rh-insights-request-id': '4c23ded5-8ee6-4d42-ab22-df1a7a8458af', 'set-cookie': 'b56a84867487668ee1c4e4a3d76d072e=d1f77ed5f9c8ffdbb345e3fbe01bd6af; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":400,"title":"Error binding parameters","detail":"code=400, message=parsing time \"2025-01-09\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"T\", internal=parsing time \"2025-01-09\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"T\""}]}
```

This is passing :

{
  "date": "2025-01-09T00:00:00Z",
  "repository_uuids": ["edb66d0f-1d44-498a-b55e-960183c42cff"]
}

```
NotFoundException: (404)
Reason: Not Found
HTTP response headers: HTTPHeaderDict({'content-length': '116', 'content-type': 'application/json', 'date': 'Fri, 10 Jan 2025 01:22:36 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZG6ZmFsc2V9fX0=', 'x-rh-insights-request-id': 'c98bb1a9-3523-4b21-9f04-a151744177e6', 'set-cookie': 'b56a84867487668ee1c4e4a3d76d072e=d1f77ed5f9c8ffdbb345e3fbe01bd6af; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":404,"title":"Error fetching snapshots","detail":"One or more repository UUIDs was invalid."}]}

```






### Comment by @mayurilahane on January 10, 2025 at 01:29 AM UTC

/lgtm

### Comment by @swadeley on January 13, 2025 at 12:59 AM UTC

/retest

### Comment by @xbhouse on January 13, 2025 at 02:10 PM UTC

cc @regexowl for awareness, this was just merged. testing in stage all looks good so far 🎉 

---

## Reviews

### Review by @dominikvagner - Approved on January 13, 2025 at 09:39 AM UTC

looks good and works as expected 🎉 👍🏼 
and all other needed changes seem to be in place, good job 👏🏼 
approved ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/935*
