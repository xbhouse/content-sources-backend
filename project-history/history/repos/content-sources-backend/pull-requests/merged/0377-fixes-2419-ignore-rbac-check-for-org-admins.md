---
type: pull_request
number: 377
title: "Fixes 2419: ignore rbac check for org admins"
state: merged
author: jlsherrill
created: 2023-08-29T12:26:15Z
updated: 2023-08-29T14:50:38Z
closed: 2023-08-29T14:50:34Z
merged: 2023-08-29T14:50:34Z
base_branch: main
head_branch: 2419
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/377
---

# Pull Request #377: Fixes 2419: ignore rbac check for org admins

**Author**: @jlsherrill
**Created**: August 29, 2023 at 12:26 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2419`

## Description

## Summary
if rbac_org_admin_skip is set to true, then don't check rbac if a user identity header tells us the user is an org_admin


## Testing steps
1.  turn on rbac_enabled in config.yaml
2. run the app with skip turned on and the mock rbac service
```
RBAC_ORG_ADMIN_SKIP=true go run cmd/content-sources/main.go api mock_rbac
```
3. use this header script to generate org_admin headers:  https://gist.github.com/jlsherrill/880bb18c288be3916c8cda94db0efacd

4.  List/create repos via the api.  Any request generated with the new script should work.  Using the older script, only those users listed in 'mocks' -> 'rbac' section of config.yaml would be allowed, all others would be rejected

---

## Discussion

### Comment by @jlsherrill on August 29, 2023 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-2419

---

## Reviews

### Review by @rverdile - Approved on August 29, 2023 at 02:16 PM UTC

tested and working for read and write permissions

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/377*
