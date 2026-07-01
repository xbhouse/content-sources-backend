---
type: pull_request
number: 742
title: "Fixes 4370: template create API allows blank arch/version"
state: merged
author: xbhouse
created: 2024-07-12T15:45:02Z
updated: 2024-07-16T20:30:19Z
closed: 2024-07-16T20:13:39Z
merged: 2024-07-16T20:13:39Z
base_branch: main
head_branch: 4370
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/742
---

# Pull Request #742: Fixes 4370: template create API allows blank arch/version

**Author**: @xbhouse
**Created**: July 12, 2024 at 03:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4370`

## Description

## Summary

Checks for arch/version before template creation and rejects request if either is missing or blank

## Testing steps

- Try to create a template with no arch or version 
- Request should be rejected and return a 400

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 12, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4370

### Comment by @mayurilahane on July 16, 2024 at 07:59 PM UTC

With blank version field

`app.content_sources.rest_client.templates_api.create_template(dict(name="my template 2", description="my new template", arch="x86_64"))`

output: 
```
ApiException: (400)
Reason: Bad Request
HTTP response headers: HTTPHeaderDict({'content-length': '98', 'content-type': 'application/json', 'date': 'Tue, 16 Jul 2024 20:04:32 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGVudGl0eZmFsc2V9fX0=', 'x-rh-insights-request-id': '355419f0-4842-41e6-aea2-112c093a6cd2', 'set-cookie': '3be893695e117cc440741f53bcf72900=931168c6d28f0eaa2f02371adcbb6a7f; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":400,"title":"Error creating template","detail":"Version cannot be blank."}]}

```

### Comment by @mayurilahane on July 16, 2024 at 08:09 PM UTC

with black arch type 

`app.content_sources.rest_client.templates_api.create_template(dict(name="my template 2", description="my new template", version="7"))`

output:

```
ApiException: (400)
Reason: Bad Request
HTTP response headers: HTTPHeaderDict({'content-length': '95', 'content-type': 'application/json', 'date': 'Tue, 16 Jul 2024 20:08:07 GMT', 'server': 'Caddy', 'x-rh-identity': 'eyJpZGVudGl2V9fX0=', 'x-rh-insights-request-id': '4c623fd2-e4ad-4edf-9832-8554278eb8d8', 'set-cookie': '3be893695e117cc440741f53bcf72900=931168c6d28f0eaa2f02371adcbb6a7f; path=/; HttpOnly; Secure; SameSite=None'})
HTTP response body: {"errors":[{"status":400,"title":"Error creating template","detail":"Arch cannot be blank."}]}
```

### Comment by @mayurilahane on July 16, 2024 at 08:13 PM UTC

/lgtm

---

## Reviews

### Review by @rverdile - Commented on July 15, 2024 at 07:56 PM UTC

### Review by @xbhouse - Commented on July 15, 2024 at 08:09 PM UTC

### Review by @xbhouse - Commented on July 15, 2024 at 08:17 PM UTC

### Review by @rverdile - Commented on July 16, 2024 at 03:17 PM UTC

### Review by @xbhouse - Commented on July 16, 2024 at 05:56 PM UTC

### Review by @rverdile - Approved on July 16, 2024 at 06:47 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/742*
