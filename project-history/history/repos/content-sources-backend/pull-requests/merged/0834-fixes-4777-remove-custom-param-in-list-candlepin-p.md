---
type: pull_request
number: 834
title: "Fixes 4777: remove custom param in list candlepin products"
state: merged
author: rverdile
created: 2024-09-27T18:43:15Z
updated: 2024-09-27T20:06:59Z
closed: 2024-09-27T20:06:56Z
merged: 2024-09-27T20:06:56Z
base_branch: main
head_branch: 4777
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/834
---

# Pull Request #834: Fixes 4777: remove custom param in list candlepin products

**Author**: @rverdile
**Created**: September 27, 2024 at 06:43 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4777`

## Description

## Summary
We don't need the custom param defined to correctly query products here. I think I misunderstood how the parameter worked when I first implemented this. 

Without this change, only custom products are included in the response, but we want to return custom and red hat products (which is the default behavior).

## Testing steps
I tested the API call in stage so it should work there

To test locally:
1. Turn off devel_org in the config file
2. `GET /api/content-sources/v1/subscription_check/` should return false
3. Turn on devel_org
4.  `GET /api/content-sources/v1/subscription_check/` should return true

---

## Discussion

### Comment by @jlsherrill on September 27, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4777

---

## Reviews

### Review by @jlsherrill - Approved on September 27, 2024 at 07:20 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/834*
