---
type: pull_request
number: 1405
title: "HMS-9864: add EPEL for aarch64"
state: merged
author: rverdile
created: 2026-03-05T15:09:37Z
updated: 2026-03-05T18:24:54Z
closed: 2026-03-05T18:24:49Z
merged: 2026-03-05T18:24:49Z
base_branch: main
head_branch: add-epel-aarch
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1405
---

# Pull Request #1405: HMS-9864: add EPEL for aarch64

**Author**: @rverdile
**Created**: March 05, 2026 at 03:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-epel-aarch`

## Description

## Summary
Adds EPEL for aarch64 architecture

## Testing steps
1. Import the new repos `make repos-import`
2. Snapshot and introspect them `go run cmd/external-repos/main.go snapshot --url https://dl.fedoraproject.org/pub/epel/10/Everything/aarch64/ --url https://dl.fedoraproject.org/pub/epel/9/Everything/aarch64/ --url https://dl.fedoraproject.org/pub/epel/8/Everything/aarch64/`



---

## Discussion

### Comment by @xbhouse on March 05, 2026 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-9864

---

## Reviews

### Review by @TenSt - Approved on March 05, 2026 at 03:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1405*
