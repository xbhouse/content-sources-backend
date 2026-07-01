---
type: pull_request
number: 607
title: "Fixes 3668: remove err log for cache miss"
state: merged
author: rverdile
created: 2024-03-14T16:24:34Z
updated: 2024-03-19T14:38:56Z
closed: 2024-03-19T14:38:53Z
merged: 2024-03-19T14:38:53Z
base_branch: main
head_branch: redis-log
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/607
---

# Pull Request #607: Fixes 3668: remove err log for cache miss

**Author**: @rverdile
**Created**: March 14, 2024 at 04:24 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `redis-log`

## Description

## Summary
Corrects the method used to check the cache error is "NotFound". The "NotFound" error is wrapped within the return error, and so `errors.Is(...)` must be used because it checks the wrapped errors and not just the top error.
## Testing steps
1. In your config.yaml set `redis.expiration.rbac : 1s`.
2. Enable mock rbac. First, enable rbac in your config.yaml under `clients.rbac_enabled:  true`. Then, in your run arguments add `mock_rbac` i.e. `go run cmd/content-sources/main.go api consumer mock_rbac`.
3. Hit the repositories list endpoint. Without this change, you'll get the "not found" error.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 14, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-3668

### Comment by @xbhouse on March 18, 2024 at 04:03 PM UTC

confirmed, without this change i see: `cache error error="redis get error: not found in cache"`

no longer seeing that with this PR. lgtm! 

---

## Reviews

### Review by @xbhouse - Approved on March 18, 2024 at 04:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/607*
