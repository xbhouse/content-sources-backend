---
type: pull_request
number: 340
title: "Fixes 2244,2280: simplify snap model and soft del repo_conf"
state: merged
author: jlsherrill
created: 2023-07-31T19:05:02Z
updated: 2023-08-08T18:56:57Z
closed: 2023-08-08T18:56:53Z
merged: 2023-08-08T18:56:53Z
base_branch: main
head_branch: snap_refactor
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/340
---

# Pull Request #340: Fixes 2244,2280: simplify snap model and soft del repo_conf

**Author**: @jlsherrill
**Created**: July 31, 2023 at 07:05 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `snap_refactor`

## Description

## Summary
This replaces the org_id & repository_uuid with a
repository_configuration_uuid column, and updates all
associated code.  This simplifies the association in order
to easily add RH repo snapshots in the future

This also switches repo_config to soft delete (via a 2nd commit).  This moves the actual row deletion into the DeleteSnapshot task.  This task now runs regardless of whether pulp is enabled or not


## Testing steps
Test a simple snapshot creation & repo deletion

no tests break

---

## Discussion

### Comment by @jlsherrill on July 31, 2023 at 07:05 PM UTC

Note this is based ontop of the domain work.  Will rebase once that is merged

### Comment by @jlsherrill on July 31, 2023 at 07:06 PM UTC

https://issues.redhat.com/browse/HMS-2244

### Comment by @jlsherrill on August 02, 2023 at 06:43 PM UTC

seeing this error on deletion, will fix:
```
 "errors": [
    {
      "status": 500,
      "title": "Error deleting repository",
      "detail": "ERROR: update or delete on table \"repository_configurations\" violates foreign key constraint \"fk_repository_config\" on table \"snapshots\" (SQLSTATE 23503)"
    }
  ]
}
```

### Comment by @jlsherrill on August 04, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-2280

### Comment by @jlsherrill on August 07, 2023 at 02:00 PM UTC

i realized a mistake on the soft delete, working on that now
EDIT: Fixed

---

## Reviews

### Review by @rverdile - Commented on August 08, 2023 at 03:43 PM UTC

### Review by @jlsherrill - Commented on August 08, 2023 at 03:52 PM UTC

### Review by @jlsherrill - Commented on August 08, 2023 at 03:56 PM UTC

### Review by @jlsherrill - Commented on August 08, 2023 at 04:01 PM UTC

### Review by @rverdile - Approved on August 08, 2023 at 06:23 PM UTC

tested and looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/340*
