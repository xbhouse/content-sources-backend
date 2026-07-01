---
type: pull_request
number: 1099
title: "Konflux build pipeline service account migration"
state: merged
author: red-hat-konflux
created: 2025-04-30T04:56:00Z
updated: 2025-04-30T13:05:10Z
closed: 2025-04-30T12:14:06Z
merged: 2025-04-30T12:14:06Z
base_branch: main
head_branch: konflux-sa-migration-content-sources-backend
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1099
---

# Pull Request #1099: Konflux build pipeline service account migration

**Author**: @red-hat-konflux
**Created**: April 30, 2025 at 04:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux-sa-migration-content-sources-backend`

## Description


## Build pipeline Service Account migration

This PR changes Service Account used by build pipeline from "appstudio-pipeline" to dedicated to the Component Service Account.
Please merge the Service Account update to avoid broken builds when deprected "appstudio-pipeline" Service Account is removed.


---

## Reviews

### Review by @jlsherrill - Approved on April 30, 2025 at 12:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1099*
