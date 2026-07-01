---
type: pull_request
number: 543
title: "Build: switch to zest v2024 and new pulp-ubi"
state: merged
author: jlsherrill
created: 2024-01-22T14:51:15Z
updated: 2024-01-22T15:32:38Z
closed: 2024-01-22T15:32:38Z
merged: 2024-01-22T15:32:38Z
base_branch: main
head_branch: zestv2024
labels: []
url: https://github.com/content-services/content-sources-backend/pull/543
---

# Pull Request #543: Build: switch to zest v2024 and new pulp-ubi

**Author**: @jlsherrill
**Created**: January 22, 2024 at 02:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `zestv2024`

## Description

## Summary

Switches to new zest build that is built with pulp-ubi.

## Testing steps

``` 
make compose-clean
make compose-up
go get -u  github.com/content-services/zest/release/v2024
make run
```

then try to curl the repositories listing:  GET http://localhost:8000/api/content-sources/v1.0/repositories/


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @swadeley - Approved on January 22, 2024 at 03:32 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/543*
