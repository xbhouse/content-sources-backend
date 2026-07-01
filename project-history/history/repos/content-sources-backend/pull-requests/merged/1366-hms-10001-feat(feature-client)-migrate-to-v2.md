---
type: pull_request
number: 1366
title: "HMS-10001: feat(feature-client): migrate to v2"
state: merged
author: TenSt
created: 2026-01-19T21:08:53Z
updated: 2026-01-21T10:51:30Z
closed: 2026-01-21T10:51:30Z
merged: 2026-01-21T10:51:30Z
base_branch: main
head_branch: stepan/HMS-10001-migrate-feature-service-to-v2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1366
---

# Pull Request #1366: HMS-10001: feat(feature-client): migrate to v2

**Author**: @TenSt
**Created**: January 19, 2026 at 09:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-10001-migrate-feature-service-to-v2`

## Description

## Summary
This PR migrates feature client from v1 to v2 and adds unit tests.

## Testing steps



---

## Discussion

### Comment by @xbhouse on January 19, 2026 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-10001

### Comment by @xbhouse on January 20, 2026 at 04:07 PM UTC

it looks like the `/features/v2` endpoint to list all features [isn't included](https://gitlab.cee.redhat.com/it-subscription-and-product/feature-service/-/blob/main/server/src/main/resources/FeatureService.openapi.yaml) in the migration. we call this via our `/admin/features` endpoint (this is used in the admin features UI):

```
GET /admin/features/

{
    "errors": [
        {
            "status": 404,
            "title": "Error listing features",
            "detail": "unexpected status code with body: {\"timestamp\":1768924758280,\"status\":404,\"error\":\"Not Found\",\"path\":\"/features/v2\"}"
        }
    ]
}
```

do you know if there's a plan to migrate this too? if not, we may have to just add either `/v1` or `/v2` depending on the endpoint we're using. 

### Comment by @TenSt on January 20, 2026 at 07:53 PM UTC

> it looks like the `/features/v2` endpoint to list all features [isn't included](https://gitlab.cee.redhat.com/it-subscription-and-product/feature-service/-/blob/main/server/src/main/resources/FeatureService.openapi.yaml) in the migration. we call this via our `/admin/features` endpoint (this is used in the admin features UI):
> 
> ```
> GET /admin/features/
> 
> {
>     "errors": [
>         {
>             "status": 404,
>             "title": "Error listing features",
>             "detail": "unexpected status code with body: {\"timestamp\":1768924758280,\"status\":404,\"error\":\"Not Found\",\"path\":\"/features/v2\"}"
>         }
>     ]
> }
> ```
> 
> do you know if there's a plan to migrate this too? if not, we may have to just add either `/v1` or `/v2` depending on the endpoint we're using.

Good catch.
I'll check it out, but seems like `v1` is completely deprecated and will be removed, so we'll need to find another endpoint/way to make this call. I'll reach out to the team and ask them.

### Comment by @TenSt on January 20, 2026 at 09:15 PM UTC

@xbhouse no, looks like only `featureStatus` endpoint is being deprecated. I'll make the necessary changes.

---

## Reviews

### Review by @xbhouse - Approved on January 20, 2026 at 10:43 PM UTC

lgtm! nice job :smile:  we'll just need to update the server url in stage once this is merged, and in prod after a prod push :rocket: 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1366*
