---
type: pull_request
number: 1500
title: "Build: bump yummy"
state: merged
author: xbhouse
created: 2026-05-20T18:44:39Z
updated: 2026-05-26T09:43:41Z
closed: 2026-05-20T20:04:03Z
merged: 2026-05-20T20:04:03Z
base_branch: main
head_branch: bump-yummy
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1500
---

# Pull Request #1500: Build: bump yummy

**Author**: @xbhouse
**Created**: May 20, 2026 at 06:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `bump-yummy`

## Description

## Summary

Bump yummy to include increased max XML size that can be parsed for introspection

## Testing steps

Import and introspect the RHEL 9.8 arm baseos repos:

```
make repos-import

go run cmd/external-repos/main.go introspect --url https://cdn.redhat.com/content/eus/rhel9/9.8/aarch64/baseos/os --url https://cdn.redhat.com/content/e4s/rhel9/9.8/aarch64/baseos/os
```

You shouldn't see an introspection error

---

## Reviews

### Review by @swadeley - Approved on May 20, 2026 at 06:55 PM UTC

ACK

### Review by @swadeley - Approved on May 20, 2026 at 07:56 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1500*
