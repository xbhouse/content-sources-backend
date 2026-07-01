---
type: pull_request
number: 511
title: "Fixes 3252: listing with limit=0 causes divide by 0 panic"
state: merged
author: xbhouse
created: 2023-12-19T21:38:08Z
updated: 2024-02-13T15:16:04Z
closed: 2024-01-02T11:06:22Z
merged: 2024-01-02T11:06:22Z
base_branch: main
head_branch: fix-0-limit
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/511
---

# Pull Request #511: Fixes 3252: listing with limit=0 causes divide by 0 panic

**Author**: @xbhouse
**Created**: December 19, 2023 at 09:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-0-limit`

## Description

## Summary

Updates limit in API request to `DefaultLimit` (100) if ever set to 0

## Testing steps

Test any listing API endpoint (`/repositories`, `/popular_repositories`, `/repositories/:uuid/rpms`, etc) and set the limit to 0. Should not panic and update limit to default of 100.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 19, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-3252

### Comment by @swadeley on January 02, 2024 at 08:53 AM UTC

/retest

### Comment by @swadeley on January 02, 2024 at 11:06 AM UTC

Hi

I tested by creating two repos, then:
`In [11]: app.content_sources.rest_client.repositories_api.list_repositories(offset=0,limit=0,sort_by='status')`

shows:
` 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0&sort_by=status',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0&sort_by=status'},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}
`

The limit is set back to the default value of 100.

If I test from a shell against stage_proxy I get `502 Bad Gateway`

Thank you



---

## Reviews

### Review by @rverdile - Approved on December 20, 2023 at 05:56 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/511*
