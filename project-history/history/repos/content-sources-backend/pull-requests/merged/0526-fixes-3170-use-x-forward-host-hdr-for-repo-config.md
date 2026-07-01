---
type: pull_request
number: 526
title: "Fixes 3170: use x-forward-host hdr for repo config"
state: merged
author: jlsherrill
created: 2024-01-10T14:19:49Z
updated: 2024-01-10T20:17:20Z
closed: 2024-01-10T20:17:20Z
merged: 2024-01-10T20:17:20Z
base_branch: main
head_branch: 3170_2
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/526
---

# Pull Request #526: Fixes 3170: use x-forward-host hdr for repo config

**Author**: @jlsherrill
**Created**: January 10, 2024 at 02:19 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3170_2`

## Description

## Summary

To properly show the gpg key url from the gateway

## Testing steps

1.  create a repo with a gpg key set, and snapshotting enabled
2.  fetch your repos list to get uuids for the repo and latest snapshots: GET http://localhost:8000/api/content-sources/v1/repositories/
3.    fetch this url: http://localhost:8000/api/content-sources/v1/repositories/60dd106b-d08e-436a-bf34-53180713d4bc/snapshots/8a68fc2d-c8eb-4046-8c56-df44603f8206/config.repo

replacing the UUIDs for the repo & snapshot UUID. Notice the URL is likely https://localhost:8000/.

4.  now fetch the url with the header:   'X-Forwarded-Host: example.com'
you should see the url change to  https://example.com/

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 10, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-3170

### Comment by @xbhouse on January 10, 2024 at 06:22 PM UTC

tested and works well! only thing i noticed is i think there is a typo in the test, should be setting the `x-forwarded-host` header? 

### Comment by @jlsherrill on January 10, 2024 at 06:26 PM UTC

good catch! i think the default of 'example.com' was being picked up, changed that to see the failure and then fixed it

### Comment by @swadeley on January 10, 2024 at 08:17 PM UTC

All API tests pass.

---

## Reviews

### Review by @xbhouse - Approved on January 10, 2024 at 06:34 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/526*
