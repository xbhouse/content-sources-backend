---
type: pull_request
number: 558
title: "Build: dependency update"
state: merged
author: jlsherrill
created: 2024-02-02T13:49:29Z
updated: 2024-02-02T15:59:31Z
closed: 2024-02-02T15:59:28Z
merged: 2024-02-02T15:59:28Z
base_branch: main
head_branch: dep_update_2_2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/558
---

# Pull Request #558: Build: dependency update

**Author**: @jlsherrill
**Created**: February 02, 2024 at 01:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `dep_update_2_2`

## Description

## Summary

Updates our dependencies with the exception of gorm v1.25.6 as we hit this bug: https://github.com/go-gorm/gorm/issues/6812  which should be fixed in the next gorm release

## Testing steps

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @rverdile - Approved on February 02, 2024 at 03:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/558*
