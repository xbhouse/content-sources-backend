---
type: pull_request
number: 1006
title: "HMS-5501: support new features api"
state: merged
author: rverdile
created: 2025-03-04T20:00:50Z
updated: 2025-03-05T16:28:53Z
closed: 2025-03-05T16:28:49Z
merged: 2025-03-05T16:28:49Z
base_branch: main
head_branch: new-features-api
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1006
---

# Pull Request #1006: HMS-5501: support new features api

**Author**: @rverdile
**Created**: March 04, 2025 at 08:00 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `new-features-api`

## Description

## Summary

- Makes a few small changes to migrate to the new feature services API
- Moves the clients to a shared folder
- Fixes a bug in the `/features/:feature_name/content/` API where high-availability features contents were returning "ha" for the distribution version

## Testing steps
1. Replace the feature_service in your config the new URL
2. Try the `/api/content-sources/v1/admin/features/` 
3. Try the `/api/content-sources/v1/admin/features/RHEL-HA-x86_64/content/` endpoint and verify the distribution version



---

## Discussion

### Comment by @jlsherrill on March 04, 2025 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5501

---

## Reviews

### Review by @jlsherrill - Approved on March 05, 2025 at 01:48 PM UTC

ACK works great, and good idea to move those two clients!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1006*
