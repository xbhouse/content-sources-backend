---
type: pull_request
number: 739
title: "Build: set new resource limits per deployment"
state: merged
author: jlsherrill
created: 2024-07-10T15:23:29Z
updated: 2024-07-10T16:07:56Z
closed: 2024-07-10T16:07:56Z
merged: 2024-07-10T16:07:56Z
base_branch: main
head_branch: limits
labels: []
url: https://github.com/content-services/content-sources-backend/pull/739
---

# Pull Request #739: Build: set new resource limits per deployment

**Author**: @jlsherrill
**Created**: July 10, 2024 at 03:23 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `limits`

## Description

## Summary

sets resource limits per deployment, as the api doesn't need as much as the workers.  Adapts to a new format to support the newest bonfire.

## Testing steps

Tests pass

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @Andrewgdewar - Approved on July 10, 2024 at 03:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/739*
