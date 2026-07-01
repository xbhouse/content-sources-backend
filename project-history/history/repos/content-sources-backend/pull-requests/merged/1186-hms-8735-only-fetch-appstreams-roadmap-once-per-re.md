---
type: pull_request
number: 1186
title: "HMS-8735: only fetch appstreams roadmap once per req"
state: merged
author: jlsherrill
created: 2025-08-27T21:15:07Z
updated: 2025-09-03T13:33:35Z
closed: 2025-09-03T13:32:36Z
merged: 2025-09-03T13:32:36Z
base_branch: main
head_branch: 8735
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1186
---

# Pull Request #1186: HMS-8735: only fetch appstreams roadmap once per req

**Author**: @jlsherrill
**Created**: August 27, 2025 at 09:15 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `8735`

## Description

## Summary

A major source for slow package searches (especially when many packages are returned) is that we are fetching the appstream from cache twice for each package.  Each fetch is relatively small (4-5ms), but when this is done up to 100 times, it can quickly add up.  

This changes it to only fetch once for the whole request

## Testing steps

in config.yaml clients section add:

```
  roadmap:
    server: https://console.stage.redhat.com/api/roadmap/v1
    username: USERNAME
    password: PASSWORD
    proxy: STAGE_PROXY
```

Restart the server.

snaphot the rhel x86_64 repos

and perform a package search:
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiZm9vIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K" \
    -H "x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae" \
    -H "Content-Type: application/json" \
    -d '{"urls":["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/baseos/os","https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os"],
          "include_package_sources": true,
          "search":"n"}'
```

when passing 1 or two letters, the search is much slower without this change, but even with passing something like 'node', the difference is noticeable (30-40ms vs 80-130ms)


---

## Discussion

### Comment by @jlsherrill on August 27, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-8735

---

## Reviews

### Review by @rverdile - Approved on September 02, 2025 at 12:46 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1186*
