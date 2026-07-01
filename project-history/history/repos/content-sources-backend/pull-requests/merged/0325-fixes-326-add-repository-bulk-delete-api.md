---
type: pull_request
number: 325
title: "Fixes 326: Add repository bulk delete API"
state: merged
author: dpang314
created: 2023-06-30T15:39:09Z
updated: 2023-07-06T08:00:32Z
closed: 2023-07-06T07:50:25Z
merged: 2023-07-06T07:50:25Z
base_branch: main
head_branch: HMS-326-bulk-delete-repositories
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/325
---

# Pull Request #325: Fixes 326: Add repository bulk delete API

**Author**: @dpang314
**Created**: June 30, 2023 at 03:39 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-326-bulk-delete-repositories`

## Description

## Summary

Adds an endpoint at `POST /repositories/bulk_delete/` to bulk delete repository configs by passing a list of UUIDs in the request body:

```
{
    uuids: ["uuid-1", "uuid-2"]
}
```

If an error occurs in deleting any of the repositories, the entire call will be canceled and return a list of errors. The max number of repositories you can delete in a single call is 100

## Testing steps

The org_id must match the org_id of the repository_configurations you are trying to delete.

```
curl "http://localhost:8000/api/content-sources/v1/repositories/bulk_delete/"  -H "`./scripts/header.sh {org_id} user`" -H "Content-Type: application/json" -d '{"uuids": [ "uuid-1" ]} '
```

---

## Discussion

### Comment by @jlsherrill on June 30, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-326

### Comment by @rverdile on July 03, 2023 at 05:02 PM UTC

everything else is working great!!

### Comment by @swadeley on July 06, 2023 at 07:50 AM UTC

Hi

I created two repos. I could delete them like so:
`app.content_sources.rest_client.repositories_api.bulk_delete_repositories({"uuids":['35d5f4ea-3a50-47bf-91d4-b92b27c5cfb6','5b9be49c-3cf7-4566-8a50-364609deda81']})`

I created two new repos. I tried to delete one plus the previously deleted  two repos:
`HTTP response body: {"errors":[{"status":404,"title":"Error deleting repositories","detail":"Could not find repository with UUID 35d5f4ea-3a50-47bf-91d4-b92b27c5cfb6"},{"status":404,"title":"Error deleting repositories","detail":"Could not find repository with UUID 5b9be49c-3cf7-4566-8a50-364609deda81"},{}]}`
and the repos are still there, as expected.

I confirm empty list returns 400:
```
In [16]: app.content_sources.rest_client.repositories_api.bulk_delete_repositories({"uuids":[]})
HTTPError: 400 Client Error: Bad Request for url: https://<snip>
```

Thank you

---

## Reviews

### Review by @rverdile - Commented on July 03, 2023 at 03:46 PM UTC

### Review by @rverdile - Commented on July 03, 2023 at 04:49 PM UTC

### Review by @rverdile - Commented on July 03, 2023 at 05:00 PM UTC

### Review by @rverdile - Commented on July 05, 2023 at 02:40 PM UTC

Notifications are working now!! I found one more thing and then this should be good to go.

Right now if I make a request with no body, or with an empty list of UUIDs, the response returns 204 but nothing actually happens. I think we should catch that somewhere and ultimately return a 400.

### Review by @rverdile - Approved on July 05, 2023 at 06:10 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/325*
