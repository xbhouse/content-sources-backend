---
type: pull_request
number: 1232
title: "HMS-9513: skip consistent org ID check for admin tasks"
state: merged
author: rverdile
created: 2025-10-09T20:37:28Z
updated: 2025-10-13T13:20:06Z
closed: 2025-10-13T13:20:03Z
merged: 2025-10-13T13:20:03Z
base_branch: main
head_branch: skip-admin
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1232
---

# Pull Request #1232: HMS-9513: skip consistent org ID check for admin tasks

**Author**: @rverdile
**Created**: October 09, 2025 at 08:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `skip-admin`

## Description

## Summary
The EnforceConsistentOrgID middleware ensures data returned matches the org of the requester. This cannot apply to the admin tasks endpoints because they return data for all orgs.

## Testing steps
1. Without this PR
2. Make a request to `/admin/tasks/` and `/admin/tasks/:uuid`. Make sure you use an org different from the one you've created repos with
3. You will get a 500 from the EnforceConsistentOrgID middleware
4. With this PR
5. Make the same requests, it should work now



---

## Discussion

### Comment by @xbhouse on October 09, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-9513

---

## Reviews

### Review by @xbhouse - Approved on October 10, 2025 at 03:07 PM UTC

nice catch, thanks for fixing this! lgtm :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1232*
