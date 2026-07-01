---
type: pull_request
number: 545
title: "Build: Update to use new pulp api path"
state: merged
author: jlsherrill
created: 2024-01-23T14:49:11Z
updated: 2024-01-23T15:52:26Z
closed: 2024-01-23T15:52:26Z
merged: 2024-01-23T15:52:26Z
base_branch: main
head_branch: pulp_api_root
labels: []
url: https://github.com/content-services/content-sources-backend/pull/545
---

# Pull Request #545: Build: Update to use new pulp api path

**Author**: @jlsherrill
**Created**: January 23, 2024 at 02:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `pulp_api_root`

## Description

## Summary

Updates our compose files to configure pulp to use /pulp/api

## Testing steps

1.  checkout PR
2. make compose-clean compose-up
3. make run
4. create a repo with snapshotting turned on, verify the snapshot gets created properly.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @rverdile - Approved on January 23, 2024 at 03:49 PM UTC

tested and works!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/545*
