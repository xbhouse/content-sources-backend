---
type: pull_request
number: 1507
title: "HMS-10698: improve rbac 401 logging"
state: merged
author: rverdile
created: 2026-05-28T17:39:09Z
updated: 2026-05-29T16:10:39Z
closed: 2026-05-29T16:10:35Z
merged: 2026-05-29T16:10:35Z
base_branch: main
head_branch: 401-logs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1507
---

# Pull Request #1507: HMS-10698: improve rbac 401 logging

**Author**: @rverdile
**Created**: May 28, 2026 at 05:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `401-logs`

## Description

## Summary
Fixes the logging so that 401 is returned in the status field. Also adds the org, resources, and verb to the log for additional debugging information.

## Testing steps
1. Enable mock_rbac
2. Make sure the user name in your identity header has no permissions
3. Make a request to an endpoint, you should get a 401
4. Check the logs, you should see the logs say `status:401` and also show the org, resource, and verb for the denied permission request


---

## Discussion

### Comment by @xbhouse on May 28, 2026 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-10698

---

## Reviews

### Review by @TenSt - Approved on May 29, 2026 at 08:12 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1507*
