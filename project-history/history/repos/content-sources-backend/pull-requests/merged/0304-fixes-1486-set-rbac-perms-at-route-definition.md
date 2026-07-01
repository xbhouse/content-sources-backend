---
type: pull_request
number: 304
title: "Fixes 1486: set rbac perms at route definition"
state: merged
author: jlsherrill
created: 2023-06-05T11:38:23Z
updated: 2023-06-08T12:08:59Z
closed: 2023-06-08T12:08:59Z
merged: 2023-06-08T12:08:59Z
base_branch: main
head_branch: 1486
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/304
---

# Pull Request #304: Fixes 1486: set rbac perms at route definition

**Author**: @jlsherrill
**Created**: June 05, 2023 at 11:38 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1486`

## Description

this moves the permission declaration for routes to when the routes are defined so that we dont keep forgetting to define them.  This caused me to refactor a lot of the rbac files
* pkg/clients/ was renamed to pkg/rbac/
* rbac.go was renamed to client.go
* anything rbac related that wasn't middleware but was in /pkg/middlware was moved to pkg/rbac/
* a new type rbac.Resource was added, similar to rbac.Verb

In addition:
* parameter validate and external gpg key calls now require write permission, it didn't make sense for them to be read

## Testing steps
Enable rbac clients in  your configuration:
```
clients:
  rbac_enabled: true
  rbac_base_url: http://localhost:8800/api/rbac/v1
```
and configure mock service:
```
mocks:
  rbac:
    # Update with your account number for admin
    user_read_write: ["readwrite"]
    # Update with yout account number for viewer
    user_read: ["read"]
    user_no_permissions: ["noperms"]
```

then try to do various read/write tasks with the users above and very it works properly

---

## Discussion

### Comment by @jlsherrill on June 05, 2023 at 12:00 PM UTC

https://issues.redhat.com/browse/HMS-1486

### Comment by @jlsherrill on June 05, 2023 at 01:22 PM UTC

Reminder, before merge, i will need to rebase and integrate tasking changes here: https://github.com/content-services/content-sources-backend/pull/296

EDIT: This is done

---

## Reviews

### Review by @jlsherrill - Commented on June 05, 2023 at 11:46 AM UTC

### Review by @jlsherrill - Commented on June 05, 2023 at 11:49 AM UTC

### Review by @rverdile - Approved on June 07, 2023 at 08:10 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/304*
