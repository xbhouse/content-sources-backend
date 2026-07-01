---
type: pull_request
number: 597
title: "Build: zest update"
state: merged
author: jlsherrill
created: 2024-03-06T13:04:57Z
updated: 2024-03-06T13:31:54Z
closed: 2024-03-06T13:31:35Z
merged: 2024-03-06T13:31:35Z
base_branch: main
head_branch: zest-update
labels: ["bug"]
url: https://github.com/content-services/content-sources-backend/pull/597
---

# Pull Request #597: Build: zest update

**Author**: @jlsherrill
**Created**: March 06, 2024 at 01:04 PM UTC
**Status**: Merged
**Labels**: `bug`
**Base**: `main` ← **Head**: `zest-update`

## Description

## Summary

## Testing steps
tests pass
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on March 06, 2024 at 01:20 PM UTC

/retest

### Comment by @mshriver on March 06, 2024 at 01:31 PM UTC

The pr-check tests are going to fail until an iqe-core update is released - which is actually blocked by the app being broken in stage. The job at least built and deployed the content-sources-backend image successfully to ephemeral.

---

## Reviews

### Review by @mshriver - Approved on March 06, 2024 at 01:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/597*
