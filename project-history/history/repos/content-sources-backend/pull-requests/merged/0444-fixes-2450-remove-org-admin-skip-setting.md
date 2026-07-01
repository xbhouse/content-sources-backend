---
type: pull_request
number: 444
title: "Fixes 2450: remove org admin skip setting"
state: merged
author: rverdile
created: 2023-10-25T13:57:07Z
updated: 2023-10-30T13:46:24Z
closed: 2023-10-30T13:46:20Z
merged: 2023-10-30T13:46:20Z
base_branch: main
head_branch: rm-org-admin
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/444
---

# Pull Request #444: Fixes 2450: remove org admin skip setting

**Author**: @rverdile
**Created**: October 25, 2023 at 01:57 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `rm-org-admin`

## Description

## Summary
Removes the setting that enables/disables "org admin rbac skip", but keeps the functionality i.e. it is permanently enabled. This allows a user with the org admin field set to true to skip the rbac check.

## Testing steps
1. In your `config.yaml` set `clients: rbac_enabled` to true.
2. When your start the server, start it with the `mock_rbac` argument. i.e. `go run cmd/content-sources/main.go api consumer instrumentation mock_rbac`
1. GET repositories with a user that does not have read permissions and the org admin field is _false_. You should get a 401. For example `./scripts/header.sh <org_id> usernoperms` will generate the identity.
3. GET repositories with a user that does not have read permissions and the org admin field is _true_. You should get a 200 response. For example `./scripts/header_org_admin.sh <org_id> usernoperms` will generate the identity.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on October 25, 2023 at 02:01 PM UTC

https://issues.redhat.com/browse/HMS-2450

### Comment by @bsquizz on October 27, 2023 at 07:06 PM UTC

/retest

### Comment by @jlsherrill on October 27, 2023 at 07:37 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on October 26, 2023 at 07:29 PM UTC

LGTM! tested and works as expected. someone else should probably double check the code though :)

### Review by @jlsherrill - Approved on October 27, 2023 at 01:29 PM UTC

ACK from a code perspective, and we'll want to to remove the RBAC_ORG_ADMIN_SKIP setttings from our app-interface deploy.yaml   for stage and prod when this goes to prod

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/444*
