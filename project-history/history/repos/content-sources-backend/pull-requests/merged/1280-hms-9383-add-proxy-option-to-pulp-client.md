---
type: pull_request
number: 1280
title: "HMS-9383: add proxy option to pulp client"
state: merged
author: rverdile
created: 2025-11-03T21:15:33Z
updated: 2025-11-05T16:11:18Z
closed: 2025-11-05T16:11:14Z
merged: 2025-11-05T16:11:14Z
base_branch: main
head_branch: pulp-stage-proxy
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1280
---

# Pull Request #1280: HMS-9383: add proxy option to pulp client

**Author**: @rverdile
**Created**: November 03, 2025 at 09:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `pulp-stage-proxy`

## Description

## Summary
Adds an option to the pulp client to specify the stage proxy

## Testing steps
In your config, set:
  - the pulp server to `https://mtls.internal.cloud.stage.redhat.com`
  - the pulp client cert
  - the pulp client key
  - the pulp proxy to the stage proxy

1. Start the server, make a request to list repositories (it should work)
2. To verify the proxy is working, set the stage proxy to something incorrect, you'll see an error

---

## Discussion

### Comment by @xbhouse on November 03, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-9383

---

## Reviews

### Review by @TenSt - Changes Requested on November 04, 2025 at 11:06 AM UTC

Could you please add some tests for the `GetTransport` function?

### Review by @TenSt - Approved on November 05, 2025 at 02:53 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1280*
