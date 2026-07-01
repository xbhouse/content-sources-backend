---
type: pull_request
number: 369
title: "Fixes 2398: Fix nil pointer error due to no pulp server in production"
state: merged
author: Andrewgdewar
created: 2023-08-21T18:18:54Z
updated: 2023-09-04T14:00:28Z
closed: 2023-08-21T21:12:21Z
merged: 2023-08-21T21:12:21Z
base_branch: main
head_branch: pulpPropFix
labels: ["bug", "qe-testing-needed", "qe-approved"]
url: https://github.com/content-services/content-sources-backend/pull/369
---

# Pull Request #369: Fixes 2398: Fix nil pointer error due to no pulp server in production

**Author**: @Andrewgdewar
**Created**: August 21, 2023 at 06:18 PM UTC
**Status**: Merged
**Labels**: `bug`, `qe-testing-needed`, `qe-approved`
**Base**: `main` ← **Head**: `pulpPropFix`

## Description

## Summary

A nil reference error was causing repositories not to complete introspection in production.
This was caused due to pulp related changes being pushed to production, without pulp being present. 

This PR adds a check to the two locations that were causing this error and the related bug.

## Testing steps


---

## Discussion

### Comment by @jlsherrill on August 21, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-2398

### Comment by @jlsherrill on August 21, 2023 at 08:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @mshriver on August 21, 2023 at 08:36 PM UTC

marking draft to block merge while unit test is resolved. 

### Comment by @mshriver on August 21, 2023 at 09:22 PM UTC

This introduced a submodule for pulp-oci-images inadvertently 

---

## Reviews

### Review by @mshriver - Approved on August 21, 2023 at 08:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/369*
