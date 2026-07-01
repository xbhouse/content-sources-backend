---
type: pull_request
number: 364
title: "Refs 2344: Add a job name stub for integration job"
state: merged
author: mshriver
created: 2023-08-16T16:50:59Z
updated: 2023-08-16T17:19:39Z
closed: 2023-08-16T17:19:39Z
merged: 2023-08-16T17:19:39Z
base_branch: main
head_branch: fix-job-name
labels: []
url: https://github.com/content-services/content-sources-backend/pull/364
---

# Pull Request #364: Refs 2344: Add a job name stub for integration job

**Author**: @mshriver
**Created**: August 16, 2023 at 04:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-job-name`

## Description

the IQE_PLUGINS parameter has invalid chars for a job name

## Summary

Auto promotion is failing because my updated job name doesn't meet job name specs with the plugin names.
https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/77500

## Testing steps


---

## Discussion

### Comment by @jlsherrill on August 16, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2344

### Comment by @mshriver on August 16, 2023 at 05:19 PM UTC

Bonfire and Openshift are happy when I run this job deployment locally, although the ephemeral cluster doesn't like our memory limit and request ratio.

Not changing that ratio as the intended target for this job is the stage and prod clusters.

---

## Reviews

### Review by @jlsherrill - Approved on August 16, 2023 at 05:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/364*
